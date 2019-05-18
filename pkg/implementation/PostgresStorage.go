package implementation

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/SankeProds/jwtservice/pkg/usecases"
	_ "github.com/lib/pq"
)

type PostgresConf interface {
	GetConnStr() string
}

type authUserPostgresStorage struct {
	connStr string
}

func (aups *authUserPostgresStorage) getDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", aups.connStr)
	return db, err
}

func NewAuthUserPostgresStorage(conf PostgresConf) *authUserPostgresStorage {
	return &authUserPostgresStorage{
		connStr: conf.GetConnStr(),
	}
}

type userRow struct {
	id         string
	data       string
	authMethod string
	authData   string
}

func userRowToAuthUser(ur userRow) (*usecases.AuthUser, error) {
	var data interface{}
	var authData interface{}
	err := json.Unmarshal([]byte(ur.data), &data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(ur.authData), &authData)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Printf("lol %+v", ur)
		return nil, err
	}
	return usecases.NewAuthUser(ur.id, data, ur.authMethod, authData), nil
}

func getUserFromRows(rows *sql.Rows) (*usecases.AuthUser, error) {
	for rows.Next() {
		var ur userRow
		err := rows.Scan(&ur.id, &ur.data, &ur.authMethod, &ur.authData)
		if err != nil {
			log.Printf("Error scaning rows")
			return nil, err
		}
		user, err := userRowToAuthUser(ur)
		return user, err
	}
	return nil, nil
}

func (aups *authUserPostgresStorage) Get(id string) (*usecases.AuthUser, error) {
	db, err := aups.getDB()
	if err != nil {
		log.Printf("Error getting db conn")
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Printf("Error quering db")
		return nil, err
	}
	defer rows.Close()
	user, err := getUserFromRows(rows)
	return user, err
}

func (aups *authUserPostgresStorage) Save(user *usecases.AuthUser) error {
	db, err := aups.getDB()
	if err != nil {
		return err
	}
	b_dataStr, err := json.Marshal(user.GetData())
	if err != nil {
		return err
	}
	b_authDataStr, err := json.Marshal(user.GetAuthData())
	if err != nil {
		return err
	}
	_, err = db.Query(fmt.Sprintf(
		`INSERT INTO users(id, data, authMethod, authData) VALUES('%s', '%s', '%s', '%s') RETURNING id`,
		user.GetId(),
		string(b_dataStr),
		user.GetAuthMethod(),
		string(b_authDataStr)))
	return err
}
