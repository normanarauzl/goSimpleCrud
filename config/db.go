package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "root"
const password string = "1234"
const host string = "localhost"
const port string = "3306"
const driver string = "mysql"
const database string = "project_go_web"

func CrearConexion() (db *sql.DB) {

	if conexion, err := sql.Open(driver, GenerarUrl()); err != nil {
		panic(err.Error())
	} else {
		db = conexion
	}
	return db
}

func GenerarUrl() string {
	return fmt.Sprint("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}
