package libs

import (
	"time"

	"go.uber.org/zap"
)

var (
	IsZap            bool
	IsZapDevelopment bool
	IsZapProduction  bool
)

type Zap struct {
}

func NewZap() *Zap {
	return &Zap{}
}

func (z *Zap) Zap() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

func (z *Zap) ZapDevelopment() {
	logger, _ := zap.NewDevelopment()
	logger.Debug("Ini log mode development")
}

func (z *Zap) ZapProduction() {
	logger, _ := zap.NewProduction() // Logger untuk mode produksi
	defer logger.Sync()              // Pastikan buffer dikosongkan sebelum keluar

	logger.Info("Halo, ini log pertama dengan zap!",
		zap.String("user", "alice"),
		zap.Int("age", 30),
	)
}
