package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	types "github.com/Arch-4ng3l/Bank/Types"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db     *sql.DB
	dbName string
	host   string
	port   uint16
}

func NewPostgres(port uint16, host string, dbName string) *Postgres {
	return &Postgres{
		dbName: dbName,
		host:   host,
		port:   port,
	}
}

func (psql *Postgres) Connect() error {
	password := os.Getenv("PSQLPW")
	username := os.Getenv("PSQLUN")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		psql.host, psql.port, username, password, psql.dbName)

	if db, err := sql.Open("postgres", connStr); err != nil {
		return err
	} else {
		psql.Db = db
	}

	_, err := psql.Db.Exec(`CREATE TABLE IF NOT EXISTS accounts(
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE,
		password TEXT UNIQUE,
		email TEXT UNIQUE,
		created_at TIMESTAMP, 
		value INT8
	)`)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (psql *Postgres) SignUp(req *types.SignUpRequest) *types.Account {
	query := `INSERT INTO accounts 
		(username, email, password, created_at, value)
		VALUES($1, $2, $3, $4, $5)`

	now := time.Now()
	_, err := psql.Db.Exec(query, req.Username, req.EmailAddress, req.Password, now, 100)
	if err != nil {
		log.Println(err)
		//return nil
	}
	acc := &types.Account{
		Username:     req.Username,
		EmailAddress: req.EmailAddress,
		Password:     req.Password,
		CreatedAt:    now,
		Value:        100,
	}
	acc.Print()
	return acc
}

func (psql *Postgres) Login(req *types.LoginRequest, encryptedPw string) *types.Account {

	query := `SELECT * FROM accounts WHERE username=$1`

	rows, err := psql.Db.Query(query, req.Username)
	if err != nil {
		log.Printf(err.Error() + "\n")
		return nil
	}
	acc := &types.Account{}

	defer rows.Close()

	if rows.Next() {
		var id int
		err := rows.Scan(&id, &acc.Username, &acc.Password, &acc.EmailAddress, &acc.CreatedAt, &acc.Value)
		if err != nil {
			log.Println(err)
			return nil
		}
		fmt.Println(encryptedPw)
		fmt.Println(acc.Password)
		if acc.Password == encryptedPw {
			return acc
		}

		return nil
	}

	return nil
}

func (psql *Postgres) Transaction(tr *types.Transaction) {
	query := `
		UPDATE accounts
			SET value = value + $1
		WHERE username = $2
	`

	if _, err := psql.Db.Exec(query, tr.Value, tr.Receiver); err != nil {
		log.Println(err)
	}
	if _, err := psql.Db.Exec(query, -tr.Value, tr.Sender); err != nil {
		log.Println(err)
	}

}
