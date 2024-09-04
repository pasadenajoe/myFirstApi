package datasources

import (
	"database/sql"
	"fmt"
	"log"
)

// --------------------------------------------------
// Constant values for connection to the DB.
// * * * * * * * * * *
// valores constantes de conexión.
// --------------------------------------------------
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "calificacion"
)

// --------------------------------------------------
// In formal API, is very common the use of a pool of
// connections to the database.  However, I havent't
// reach that point in GO, and tried to make it simple.
//
// This Function is meant to be the base to get a connection
// "objects" to the database.
// * * * * * * * * * *
// En API productivos, es bastante común el uso de un
// pool de conexiones.  No obstante, todavía no he llegado
// a ese punto en Go, y tratamos de hacerlo simple.
//
// Esta función está diseñada para que devuelva un objeto de
// conexión a la base de datos.
// --------------------------------------------------
func connectToDb() *sql.DB {
	// --------------------------------------------------
	// Use of constants to db connection.  In production
	// environment, the password will not be here.
	// * * * * * * * * * *
	// Uso de las constantes de conexión.
	// En ambientes productivos, la contraseña no puede
	// estar en este código.
	// --------------------------------------------------
	psqlInfo := fmt.Sprintf(
		"host=%s "+
			"port=%d "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		host, port, user, password, dbname)

	// --------------------------------------------------
	// Open the connection
	// * * * * * * * * * *
	// Abriendo la conexión.
	// --------------------------------------------------
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("No se pudo conectar a la base dd datos")
		panic(err)
	}

	// --------------------------------------------------
	// Confirm a successful connection.
	// * * * * * * * * * *
	// Confirmando que haya conexión.
	// --------------------------------------------------
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// --------------------------------------------------
	// Returning the connection as an "object"
	// * * * * * * * * * *
	// regresando la conexión con un "objeto"
	// --------------------------------------------------
	return db
}
