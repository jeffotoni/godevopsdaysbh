//
// conexao com postgresql
//
package pg

///
import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once

var (
	dbLocal *sql.DB
	err     error
)

// singleton versao
// resumida e eficiente
func ConnectNew() (db *sql.DB) {

	if dbLocal != nil {
		return dbLocal
	} else {

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_BANCO, DB_SSL)

		once.Do(func() {
			dbLocal, err = sql.Open(DB_SORCE, DBINFO)
		})

		if err != nil {
			log.Println("Erro ao tentar conectar Postgres: ", err)
			return dbLocal
		}

		return dbLocal
	}
}
