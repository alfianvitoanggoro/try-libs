package test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

// Struct untuk Asynq
type Asynq struct {
	RedisOpt  asynq.RedisClientOpt
	Client    *asynq.Client
	Scheduler *asynq.Scheduler
}

func NewAsynq() *Asynq {
	return &Asynq{
		RedisOpt:  asynq.RedisClientOpt{Addr: "localhost:6379"},
		Client:    asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"}),
		Scheduler: asynq.NewScheduler(asynq.RedisClientOpt{Addr: "localhost:6379"}, nil),
	}
}

// Handler untuk job email:send
func handleEmailTask(ctx context.Context, t *asynq.Task) error {
	email := string(t.Payload()) // Ambil email dari payload
	fmt.Printf("📨 Mengirim email ke: %s...\n", email)

	// Simulasi error (50% kemungkinan gagal)
	if time.Now().Unix()%2 == 0 {
		return fmt.Errorf("❌ Gagal mengirim email ke %s", email)
	}

	fmt.Println("✅ Email berhasil dikirim!")
	return nil
}

// Handler untuk memproses task
func handleEmailTaskDelayedJob(ctx context.Context, t *asynq.Task) error {
	email := string(t.Payload()) // Ambil email dari payload
	fmt.Printf("📩 Mengirim email ke: %s...\n", email)
	return nil
}

// Worker (Harus dijalankan dalam proses terpisah)
func (a *Asynq) Worker() {
	// Konfigurasi Worker
	server := asynq.NewServer(a.RedisOpt, asynq.Config{
		Concurrency: 10, // Bisa menjalankan 10 job sekaligus
		Queues: map[string]int{
			"critical": 3, // Prioritas tinggi (lebih banyak worker)
			"default":  2,
			"low":      1, // Prioritas rendah
		},
	})

	// Router untuk task yang ditangani Worker
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", handleEmailTask)
	mux.HandleFunc("email:delayed", handleEmailTaskDelayedJob)
	mux.HandleFunc("email:daily_report", handleEmailTaskDelayedJob)

	fmt.Println("🚀 Worker berjalan...")
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

// Client untuk menambahkan job ke antrian
func (a *Asynq) Send() {
	defer a.Client.Close()

	// Buat task untuk mengirim email
	email := "user@example.com"
	task := asynq.NewTask("email:send", []byte(email))

	// Tambahkan ke antrian dengan opsi retry & delay
	info, err := a.Client.Enqueue(task,
		asynq.MaxRetry(5),              // Retry maksimal 5 kali
		asynq.ProcessIn(5*time.Second), // Delay 5 detik sebelum dieksekusi
		asynq.Queue("default"),         // Masukkan ke queue "default"
	)

	if err != nil {
		log.Fatalf("❌ Gagal menambahkan job: %v", err)
	}

	log.Printf("✅ Job berhasil ditambahkan: ID=%s", info.ID)
}

func (a *Asynq) SendDelayedJob() {
	defer a.Client.Close()

	// Buat task baru untuk kirim email
	email := "user@example.com"
	task := asynq.NewTask("email:delayed", []byte(email))

	// Enqueue task dengan delay 1 menit
	info, err := a.Client.Enqueue(task, asynq.ProcessIn(1*time.Minute))
	if err != nil {
		log.Fatalf("❌ Gagal enqueue task: %v", err)
	}

	fmt.Printf("✅ Task akan dieksekusi dalam 1 menit. ID=%s\n", info.ID)
}

func (a *Asynq) SendCronJob() {
	defer a.Client.Close()

	email := "usercronjob@example.com"
	// Jadwalkan job untuk dijalankan setiap hari jam 12 siang
	_, err := a.Scheduler.Register("0 12 * * *", asynq.NewTask("email:daily_report", []byte(email)))
	if err != nil {
		log.Fatal("Gagal menambahkan jadwal:", err)
	}

	log.Println("📅 Scheduler aktif, cek di AsynqMon!")
	if err := a.Scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}

// Implement to cobra
var (
	asynqWorker         bool
	asynqSend           bool
	asynqSendDelayedJob bool
	asyncSendCronJob    bool
)

// command
var asynqCmd = &cobra.Command{
	Use:   "asynq",
	Short: "asynq",
	Long:  `asynq for using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		a := NewAsynq()

		switch {
		default:
			cmd.Help()
		case asynqWorker:
			a.Worker()
		case asynqSend:
			a.Send()
		case asynqSendDelayedJob:
			a.SendDelayedJob()
		case asyncSendCronJob:
			a.SendCronJob()
		}
	},
}
