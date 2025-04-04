package libs

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

var (
	TONWallet  bool
	TONAccount bool
)

type TONUtils struct {
	Client *liteclient.ConnectionPool
	Api    *ton.APIClient
	Ctx    context.Context
}

func NewTONUtils() *TONUtils {
	client := liteclient.NewConnectionPool()

	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json"
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		panic(err)
	}

	api := ton.NewAPIClient(client).WithRetry().(*ton.APIClient)
	// bound all requests to single ton node
	ctx := client.StickyContext(context.Background())

	return &TONUtils{
		Client: client,
		Api:    api,
		Ctx:    ctx,
	}
}

/*
	Buka bot Telegram ini:
🔗 @testgiver_ton_bot

	Kirim alamat wallet testnet kamu (misalnya: UQC7XVoYjUc9I8Cb-xTo1XUgHNxr29p7NEqBRXxOETGkfZZJ).

	Tunggu beberapa saat, lalu cek saldo di wallet kamu.
*/

// ! This use testnet for test for check balance and send transaction
func (t *TONUtils) TONWallet() {
	// seed words of account, you can generate them with any wallet or using wallet.NewSeed() method
	words := strings.Split("diet diet attack autumn expose honey skate lounge holiday opinion village priority major enroll romance famous motor pact hello rubber express warfare rose whisper", " ")

	w, err := wallet.FromSeed(t.Api, words, wallet.ConfigV5R1Final{
		NetworkGlobalID: wallet.TestnetGlobalID,
	})
	if err != nil {
		log.Fatalln("FromSeed err:", err.Error())
		return
	}

	log.Println("wallet address:", w.WalletAddress())

	log.Println("fetching and checking proofs since config init block, it may take near a minute...")
	block, err := t.Api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		log.Fatalln("get masterchain info err: ", err.Error())
		return
	}
	log.Println("master proof checks are completed successfully, now communication is 100% safe!")

	balance, err := w.GetBalance(t.Ctx, block)
	if err != nil {
		log.Fatalln("GetBalance err:", err.Error())
		return
	}

	if balance.Nano().Uint64() >= 3000000 {
		addr := address.MustParseAddr("UQC7XVoYjUc9I8Cb-xTo1XUgHNxr29p7NEqBRXxOETGkfZZJ")

		log.Println("sending transaction and waiting for confirmation...")

		// if destination wallet is not initialized (or you don't care)
		// you should set bounce to false to not get money back.
		// If bounce is true, money will be returned in case of not initialized destination wallet or smart-contract error
		bounce := false

		transfer, err := w.BuildTransfer(addr, tlb.MustFromTON("0.003"), bounce, "Hello from tonutils-go!")
		if err != nil {
			log.Fatalln("Transfer err:", err.Error())
			return
		}

		tx, block, err := w.SendWaitTransaction(t.Ctx, transfer)
		if err != nil {
			log.Fatalln("SendWaitTransaction err:", err.Error())
			return
		}

		balance, err = w.GetBalance(t.Ctx, block)
		if err != nil {
			log.Fatalln("GetBalance err:", err.Error())
			return
		}

		log.Printf("transaction confirmed at block %d, hash: %s balance left: %s", block.SeqNo,
			base64.StdEncoding.EncodeToString(tx.Hash), balance.String())

		return
	}

	log.Println("not enough balance:", balance.String())
}

func (t *TONUtils) TONAccount() {
	// TON Foundation account
	addr := address.MustParseAddr("0QC7XVoYjUc9I8Cb-xTo1XUgHNxr29p7NEqBRXxOETGkfS3D")

	block, err := t.Api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		log.Fatalln("get masterchain info err: ", err.Error())
		return
	}

	account, err := t.Api.GetAccount(context.Background(), block, addr)
	if err != nil {
		log.Fatalln("get account err:", err.Error())
		return
	}

	// Balance: ACTIVE
	fmt.Printf("Status: %s\n", account.State.Status)
	// Balance: 66559946.09 TON
	fmt.Printf("Balance: %s TON\n", account.State.Balance.String())
	if account.Data != nil { // Can be nil if account is not active
		// Data: [0000003829a9a31772c9ed6b62a6e2eba14a93b90462e7a367777beb8a38fb15b9f33844d22ce2ff]
		fmt.Printf("Data: %s\n", account.Data.Dump())
	}

	// load last 15 transactions
	list, err := t.Api.ListTransactions(context.Background(), addr, 15, account.LastTxLT, account.LastTxHash)
	if err != nil {
		// In some cases you can get error:
		// lite server error, code XXX: cannot compute block with specified transaction: lt not in db
		// it means that current lite server does not store older data, you can query one with full history
		log.Printf("send err: %s", err.Error())
		return
	}

	// oldest = first in list
	for _, t := range list {
		// Out: 620.9939549 TON, To [EQCtiv7PrMJImWiF2L5oJCgPnzp-VML2CAt5cbn1VsKAxLiE]
		// In: 494.521721 TON, From EQB5lISMH8vLxXpqWph7ZutCS4tU4QdZtrUUpmtgDCsO73JR
		// ....
		fmt.Println(t.String())
	}
}
