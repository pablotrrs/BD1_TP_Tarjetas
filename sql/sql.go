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

func BorrarDB() {
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

func ProbarConsumo() {
	autorizacionCompra()
	crearTriggers()
	generarConsumos()	
	testFunciones()
}

func ProbarResumen(){
	crearCompras()
	generarResumen()
	
	_, err = db.Query(`
	
	SELECT generarresumen(22648991, 11); 

	`)	
	if err != nil {
		log.Fatal(err)
	}
	
	
}

func crearCompras() {
	_, err = db.Query(`
	INSERT INTO compra VALUES(001, 4000001355435322, 538, '2020/11/10'::date, 100, false); 
	INSERT INTO compra VALUES(002, 4000001355435322, 222, CURRENT_DATE, 500, false); 
	INSERT INTO compra VALUES(003, 4000001355435322, 345, CURRENT_DATE, 50, false); 
	INSERT INTO compra VALUES(004, 4000001355435322, 999, '2020/11/30'::date, 50, false); 
	INSERT INTO compra VALUES(005, 4000001355435322, 588, '2020/11/05'::date, 80, false); 
	`)	
	if err != nil {
		log.Fatal(err)
	}	
}

