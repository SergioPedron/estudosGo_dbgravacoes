package dbgravacoes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

//-----------------------------------------------------------------------------------------------------------------------------------

// Conecta ao banco mySQl gravacoes.  Executado no Init da aplicação
func Conectadbgravacoes() *sql.DB {
	var db *sql.DB
	cfg := mysql.Config{
		User:   "sergio",
		Passwd: "sergio",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gravacoes",
	}

	// Pega o identificar do banco
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err) // Rever o uso do Fatal, pois não tem como tratar o erro fora do pacote e irá parar o programa
	}

	// Verifica a conexão
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conectado ao banco de gravações!")
	return db
}

//-----------------------------------------------------------------------------------------------------------------------------------
