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

func ProbarConsumo() {
	probarConsumo()
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

func probarConsumo() {
	autorizacionCompra()
	crearTriggers()
	//probar()	
}

func probar() {
	_, err = db.Query(`
	SELECT autorizacion_compra(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --compra
	SELECT autorizacion_compra(CAST(9000001234567899 AS char(16)), CAST(733 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --tarjeta inexistente
	SELECT autorizacion_compra(CAST(4033002233062344 AS char(16)), CAST(202 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --tarjeta no vigente
	SELECT autorizacion_compra(CAST(4034006634262869 AS char(16)), CAST(097 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --tarjeta suspendida
	SELECT autorizacion_compra(CAST(4000001234567899 AS char(16)), CAST(111 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --codigo incorrecto
	INSERT INTO compra VALUES(999, '4032002134557009', 501, CURRENT_DATE, 25000, false); -- para que la siguiente query sea efectiva
	SELECT autorizacion_compra(CAST(4032002134557009 AS char(16)), CAST(070 AS char(4)), CAST(501 AS int), CAST(50001 AS decimal(7,2))); --supera monto
	SELECT autorizacion_compra(CAST(4000001234567899 AS char(16)), CAST(733 AS char(4)), CAST(501 AS int), CAST(12 AS decimal(7,2))); --vencida
	`)	
	if err != nil {
		log.Fatal(err)
	}	
}
