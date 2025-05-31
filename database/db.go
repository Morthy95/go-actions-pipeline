package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || pass == "" || name == "" || port == "" {
		log.Panic("Alguma variável de ambiente está vazia!")
	}

	stringDeConexao := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + name + " port=" + port + " sslmode=disable"

	//print para verificar a string de conexão
	fmt.Println("String de Conexão:", stringDeConexao)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	_ = DB.AutoMigrate(&models.Aluno{})
}
