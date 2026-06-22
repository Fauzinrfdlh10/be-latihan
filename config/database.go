package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("env tidak ditemukan")
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan di dalam file .env")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal terhubung ke database", err)
	}

	fmt.Println("Koneksi ke postgresSQL berhasil")

}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("database belum diinisialisasi")
	}
	return DB
}
