package sql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func DbConnection() {
	db, err = sql.Open("postgres", "user=postgres host=localhost dbname=tarjeta sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func CrearDB() {
	crearDB()
}

func BorrarBD() {
	BorrarDB();
}

func CrearTablas(){
	crearTablas();
}

func BorrarTablas(){
	borrarTablas();
}

func CrearPKsyFKs() {
	crearPKs()
	crearFKs()
}

func BorrarPKsyFKs() {
	borrarFKs()
	borrarPKs()
}

func CargarDatos() {
	cargarDatos()
	InsertarCierres()
}

func BorrarDatos() {
	borrarDatos()
}

func crearDB() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

    _, err = db.Exec(`CREATE DATABASE tarjeta`)
	if err != nil {
		log.Fatal(err)
	}	
}

func BorrarDB(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
    defer db.Close()
   
    _, err = db.Exec(`DROP DATABASE tarjeta`)
	if err != nil {
		log.Fatal(err)
	}
}
