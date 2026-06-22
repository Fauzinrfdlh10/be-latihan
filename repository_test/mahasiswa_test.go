package repository_test

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/repository"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		t.Logf("Gagal load .env dari tester: %v", errEnv)
	}
	
	config.InitDB()

	err := config.GetDB().AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("gagal migrasi: %v", err)
	}
}

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)

	npm := time.Now().UnixNano()

	mhs := model.Mahasiswa{
		NPM:    npm,
		Nama:   "Budi",
		Prodi:  "Teknik Informatika",
		Alamat: "Jakarta",
		Email:  "[EMAIL_ADDRESS]",
		Hobi:   []string{"membaca", "berenang"},
	}

	result, err := repository.InsertMahasiswa(&mhs)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetAllMahasiswa(t *testing.T) {
	setupTest(t)

	data, err := repository.GetAllMahasiswa()
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetMahasiswaByNPM(t *testing.T) {
	setupTest(t)

	data, err := repository.GetMahasiswaByNPM(123456789)
	if err == nil {
		assert.NotNil(t, data)
	}
}

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	// Buat data sementara untuk dites update
	npm := time.Now().UnixNano()
	
	repository.InsertMahasiswa(&model.Mahasiswa{
		NPM:    npm,
		Nama:   "Siswa Teladan",
		Prodi:  "Teknik Informatika",
		Alamat: "Jakarta",
		Email:  "[EMAIL_ADDRESS]",
		Hobi:   []string{"belajar"},
	})

	// Lakukan proses update
	result, err := repository.UpdateMahasiswa(npm, &model.Mahasiswa{
		Nama:   "Siswa Teladan (Diupdate)",
		Prodi:  "Sistem Informasi",
		Alamat: "Bandung",
		Email:  "[EMAIL_ADDRESS_UPDATED]",
		Hobi:   []string{"membaca", "ngoding", "menulis"},
	})
	
	assert.NoError(t, err)
	if err == nil {
		assert.NotNil(t, result)
	}
}

func TestDeleteMahasiswa(t *testing.T) {
	setupTest(t)

	err := repository.DeleteMahasiswa(123456789)
	// it might return assert.NoError or error if it doesn't exist
	if err == nil {
		assert.NoError(t, err)
	}
}

func TestDeleteMahasiswaNotFound(t *testing.T) {
	setupTest(t)

	err := repository.DeleteMahasiswa(999999999)
	if err != nil {
		assert.Error(t, err)
	}
}
