package dbmanager

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type DBManager struct {
	Instance *sql.DB
}

//DB指针

func init() {

}

//获取实例
func GetConn(username string, ip string, port int, tablename string, password string) DBManager {
	var dbManager DBManager
	var err error

	t := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, ip, port, tablename)
	log.Println(t)

	dbManager.Instance, err = sql.Open("mysql", t)
	if err != nil {
		dbManager.Instance.Close()
		log.Fatal(err)
	}
	dbManager.Instance.SetConnMaxLifetime(2 * 60)
	dbManager.Instance.SetMaxIdleConns(1000)
	dbManager.Instance.SetMaxOpenConns(2000)

	err = dbManager.Instance.Ping()
	if err != nil {
		dbManager.Instance.Close()
		log.Fatal(err)
	}

	return dbManager
}

//查询
func (db *DBManager) Query() (*sql.Rows, error) {
	if db == nil {
		return nil, errors.New("sql: DBManager is nil")
	}

	return db.Instance.Query(string(""))
}

//释放
func (db *DBManager) Close(query string) error {
	if db == nil {
		return errors.New("sql: DBManager is nil")
	}

	return db.Instance.Close()
}
