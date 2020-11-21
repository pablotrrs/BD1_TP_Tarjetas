package main

import (
	"fmt"

	//noSQL "./no-sql"
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

func mostrarMenu() {
	fmt.Println("-------------------------------------------\n|Seleccione una opción y presione enter   |\n-------------------------------------------")
	fmt.Println("|1. Crear base de datos                   |\n-------------------------------------------")
	fmt.Println("|2. Borrar base de datos                  |\n-------------------------------------------")
	fmt.Println("|3. Crear tablas                          |\n-------------------------------------------")
	fmt.Println("|4. Borrar tablas                         |\n-------------------------------------------")	
	fmt.Println("|5. Crear PK's & FK's                     |\n-------------------------------------------")
	fmt.Println("|6. Borrar PK's & FK's                    |\n-------------------------------------------")
	fmt.Println("|7. Cargar todos los datos                |\n-------------------------------------------")
	fmt.Println("|8. Borrar todos los datos                |\n-------------------------------------------")	
	fmt.Println("|9. Probar consumo                        |\n-------------------------------------------")	
	fmt.Println("|10. Salir                                |\n-------------------------------------------")
	
}
func manejarOpciones(opcion int) bool {
	switch {
	case opcion == 1:
		sql.CrearDB()
		fmt.Println("Base de datos creada.")
	case opcion == 2:
		sql.BorrarBD()
		fmt.Println("Base de datos borrada.")
	case opcion == 3:
		sql.CrearTablas()
		fmt.Println("Tablas creadas.")
	case opcion == 4:
		sql.BorrarTablas()
		fmt.Println("Tablas borradas.")		
	case opcion == 5:
		sql.CrearPKsyFKs()
		fmt.Println("PK's y FK's creadas.")
	case opcion == 6:
		sql.BorrarPKsyFKs()
		fmt.Println("PK's y FK's borradas.")
	case opcion == 7:
		sql.CargarDatos()
		fmt.Println("Todos los datos fueron cargados.")
	case opcion == 8:
		sql.BorrarDatos()
		fmt.Println("Todos los datos fueron borrados.")		
	case opcion == 9:
		sql.ProbarConsumo()
		fmt.Println("Probado.")		
	case opcion == 10:
		return false
	default:
		fmt.Println("Ingrese un numero valido.")
	}
	return true
}
