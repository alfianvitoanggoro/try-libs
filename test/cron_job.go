package test

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	c := cron.New(cron.WithSeconds()) // Gunakan WithSeconds() agar format cron dengan detik bisa digunakan

	// Menjalankan tugas setiap 10 detik
	c.AddFunc("*/10 * * * * *", func() {
		log.Println("Cron Job Running:", time.Now().Format("15:04:05"))
	})

	// Memulai cron job
	c.Start()

	// Gunakan channel untuk memastikan program tetap berjalan
	// Gunakan select {} jika hanya butuh cara sederhana untuk menjaga program tetap berjalan.
	// select{}

	// Gunakan channel (chan) jika ingin graceful shutdown atau lebih banyak kontrol atas kapan program berhenti.
	// menggunakan channel
	done := make(chan bool)
	<-done
}

func CronJobWithoutLibrary() {
	ticker := time.NewTicker(10 * time.Second)

	// Loop yang akan terus berjalan
	for t := range ticker.C {
		fmt.Println("Task running at:", t.Format("15:04:05"))
	}
}
