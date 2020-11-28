package main

import (
	"fmt"
	"log"
	"strconv"
	bolt "github.com/coreos/bbolt"
	"encoding/json"
	"./sql"
)

func main() {

	sql.DbConnection()

	running := true
	var opcion int
	
	mostrarMenu()
	
	for running {
		if ret, _ := fmt.Scanln(&opcion); ret == 1 { //Scanea y guarda 1 en ret si el dato que leyo es del tipo de opcion. Esto restringe el scan a ints
			running = manejarOpciones(opcion)
		}
	}
}

func verificarError(err error){
		if err != nil {
        log.Fatal(err)
    }
}

func mostrarMenu() {
	fmt.Println("-------------------------------------------\n|Seleccione una opción y presione enter   |\n-------------------------------------------")
	fmt.Println("|0. Salir                                 |\n-------------------------------------------")	
	fmt.Println("|1. Crear base de datos                   |\n-------------------------------------------")
	fmt.Println("|2. Crear tablas                          |\n-------------------------------------------")	
	fmt.Println("|3. Crear PK's & FK's                     |\n-------------------------------------------")	
	fmt.Println("|4. Cargar todos los datos                |\n-------------------------------------------")
	fmt.Println("|5. Probar consumo                        |\n-------------------------------------------")			
	fmt.Println("|6. Generar resumen                       |\n-------------------------------------------")
	fmt.Println("|7. Crear base de datos no SQL            |\n-------------------------------------------")
	fmt.Println("|8. Borrar base de datos                  |\n-------------------------------------------")	
	fmt.Println("|9. Borrar tablas                         |\n-------------------------------------------")	
	fmt.Println("|10. Borrar PK's & FK's                   |\n-------------------------------------------")
	fmt.Println("|11. Borrar todos los datos               |\n-------------------------------------------")	
}

func manejarOpciones(opcion int) bool {
	switch {
	case opcion == 0:
		return false		
	case opcion == 1:
		sql.CrearDB()
		fmt.Println("Base de datos creada.")
	case opcion == 2:
		sql.CrearTablas()
		fmt.Println("Tablas creadas.")		
	case opcion == 3:
		sql.CrearPKsyFKs()
		fmt.Println("PK's y FK's creadas.")
	case opcion == 4:
		sql.CargarDatos()
		fmt.Println("Todos los datos fueron cargados.")		
	case opcion == 5:
		sql.ProbarConsumo()
		fmt.Println("Probado.")		
	case opcion == 6:
		sql.ProbarResumen()
		fmt.Println("Ok.")						
	case opcion == 7:
		BDnoSQL()
		fmt.Println("Base de datos no SQL creada.")		
	case opcion == 8:
		sql.BorrarBD()
		fmt.Println("Base de datos borrada.")
	case opcion == 9:
		sql.BorrarTablas()
		fmt.Println("Tablas borradas.")		
	case opcion == 10:
		sql.BorrarPKsyFKs()
		fmt.Println("PK's y FK's borradas.")
	case opcion == 11:
		sql.BorrarDatos()
		fmt.Println("Todos los datos fueron borrados.")	
	default:
		fmt.Println("Ingrese un numero valido.")
	}
	return true
}

type Cliente struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjeta struct {
	Nrotarjeta   string
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra int
	Estado       string
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compra struct {
	Nrooperacion int
	Nrotarjeta   string
	Nrocomercio  int
	Fecha        string
	Monto        int
	Pagado       bool
}


func BDnoSQL() {

    db, err := bolt.Open("testBolt.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	cliente1 := Cliente{33348679, "Sofía", "Godoy", "Av. Senador Morón 1221", "115598342"}
	cliente2 := Cliente{44349773, "Abril", "Hernández", "Av. Sourdeaux 1700", "115598342"}
	cliente3 := Cliente{14348789, "Ricardo", "Llanos", "Corrientes 183", "119034572"}

	comercio1 := Comercio{564, "FOX", "Av Pres. Juan Domingo Perón 907", "1663", "46676777"}
	comercio2 := Comercio{523, "47 street", "Paunero 1575", "1663", "47597581"}
	comercio3 := Comercio{553, "Disco", "Av. Senador Morón 960", "1661", "08107778888"}
	
	tarjeta1 := Tarjeta{"4000001234567899", 11348773, "201508", "202008", "733", 50000, "vigente"}
	tarjeta2 := Tarjeta{"4037001554363655", 12349972, "201507", "202007", "332", 55000, "vigente"}
	tarjeta3 := Tarjeta{"4000001355435322", 22648991, "201507", "202007", "201", 60000, "vigente"}

	compra1 := Compra{1, "4000001234567899", 501, "2020-04-25 00:00:00", 1500.00, true}
	compra2 := Compra{2, "4000001234567899", 513, "2020-04-27 00:00:00", 4500.00, true}
	compra3 := Compra{3, "4000001234567899", 523, "2020-04-30 00:00:00", 850.00, true}
	
	dataCliente1, err := json.Marshal(cliente1)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)), dataCliente1 )
	     resultadoCliente1, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente1)
    
    
	dataCliente2, err := json.Marshal(cliente2)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)), dataCliente2 )
	     resultadoCliente2, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente2)
    
    
	dataCliente3, err := json.Marshal(cliente3)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)), dataCliente3 )
	     resultadoCliente3, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente3)
    
    
	dataComercio1, err := json.Marshal(comercio1)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)), dataComercio1 )
	     resultadoComercio1, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio1)
    
    
	dataComercio2, err := json.Marshal(comercio2)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)), dataComercio2 )
	     resultadoComercio2, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio2)
    
    
	dataComercio3, err := json.Marshal(comercio3)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)), dataComercio3 )
	     resultadoComercio3, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio3)
    
    
	dataTarjeta1, err := json.Marshal(tarjeta1)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta1.Nrotarjeta), dataTarjeta1 )
	     resultadoTarjeta1, err := ReadUnique(db, "tarjeta", []byte(tarjeta1.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta1)
    
    
	dataTarjeta2, err := json.Marshal(tarjeta2)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta2.Nrotarjeta), dataTarjeta2 )
	     resultadoTarjeta2, err := ReadUnique(db, "tarjeta", []byte(tarjeta2.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta2)
    
    
	dataTarjeta3, err := json.Marshal(tarjeta3)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta3.Nrotarjeta), dataTarjeta3 )
	     resultadoTarjeta3, err := ReadUnique(db, "tarjeta", []byte(tarjeta3.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta3)
    
    
	dataCompra1, err := json.Marshal(compra1)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra1.Nrooperacion)), dataCompra1 )
	     resultadoCompra1, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra1.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra1)
    
    
	dataCompra2, err := json.Marshal(compra2)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra2.Nrooperacion)), dataCompra2 )
	     resultadoCompra2, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra2.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra2)
    
    
	dataCompra3, err := json.Marshal(compra3)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra3.Nrooperacion)), dataCompra3 )
	     resultadoCompra3, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra3.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra3)	
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, value []byte) error {

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, value)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {

	var buf []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}
