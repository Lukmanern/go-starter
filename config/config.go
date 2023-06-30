package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfg               Config
	PublicKey         *[]byte
	PrivateKey        *[]byte
	PublicKeyReadOne  sync.Once
	PrivateKeyReadOne sync.Once
	cfgOnce           sync.Once
	envFile           *string
	databaseURI       string
)

type Config struct {
	AccessTokenTTL time.Duration `env:"ACCESS_TOKEN_TTL"`
	Port           int           `env:"PORT"`
	DBConncetion   string        `env:"DB_CONNECTION"`
	DBPort         int           `env:"DB_PORT"`
	DBUser         string        `env:"DB_USER"`
	DBPassword     string        `env:"DB_PASSWORD"`
	DBName         string        `env:"DB_DBNAME"`
	DBRootHost     string        `env:"DB_ROOT_HOST"`
	DBRootPassword string        `env:"DB_ROOT_PASSWORD"`
	SMTPServer     string        `env:"SMTP_SERVER"`
	SMTPPort       int           `env:"SMTP_PORT"`
	SMTPEmail      string        `env:"SMTP_EMAIL"`
	SMTPPassword   string        `env:"SMTP_PASSWORD"`
	ClientURL      string        `env:"CLIENT_URL"`
	PublicKey      string        `env:"PUBLIC_KEY"`
	PrivateKey     string        `env:"PRIVATE_KEY"`
	SecretBytes    string        `env:"SECRET_BYTES"`
}

// ReadConfig untuk membaca file konfigurasi.
// Jika file tidak ditemukan, fungsi ini akan
// menampilkan pesan kesalahan. Fungsi ini juga
// dapat membaca variabel lingkungan jika file
// konfigurasi tidak dapat dibaca atau tidak ditemukan.
func ReadConfig(file string) *Config {
	// checking file is exist
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Config file not found: %s", file)
		} else {
			log.Fatalf("Error checking config file: %s", err.Error())
		}
	}

	cfgOnce.Do(func() {
		envFile = &file
		log.Printf(`Reading config file: "%s"`, *envFile)
		err := cleanenv.ReadConfig(file, &cfg)
		if err != nil {
			err := cleanenv.ReadEnv(&cfg)
			if err != nil {
				log.Fatalf("Config error %s", err.Error())
			}
		}
	})
	return &cfg
}

// Configuration untuk mengembalikan konfigurasi
// yang telah dibaca. Jika file konfigurasi belum
// diatur, fungsi ini akan menampilkan pesan kesalahan.
func Configuration() Config {
	if envFile == nil {
		log.Panic(`configuration file is not set. Call ReadConfig("path_to_file") first`)
	}
	err := cleanenv.UpdateEnv(&cfg)
	if err != nil {
		log.Fatalf("Config error %s", err.Error())
	}
	return cfg
}

func (c Config) GetDatabaseURI() string {
	params := "?charset=utf8mb4&multiStatements=true&parseTime=true"
	databaseURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s"+params, c.DBUser,
		c.DBPassword, c.DBRootHost, c.DBPort, c.DBName)

	return databaseURI
}

func (c Config) GetPublicKey() []byte {
	PublicKeyReadOne.Do(func() {
		signKey, err := os.ReadFile(c.PublicKey)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		PublicKey = &signKey
	})
	return *PublicKey
}

func (c Config) GetPrivateKey() []byte {
	PrivateKeyReadOne.Do(func() {
		signKey, err := os.ReadFile(c.PrivateKey)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		PrivateKey = &signKey
	})
	return *PrivateKey
}
