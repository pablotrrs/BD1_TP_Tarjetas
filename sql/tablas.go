package sql

import(
	"log"

	_"github.com/lib/pq"
)

func crearTablas() {
	borrarTablas()
	
	_, err = db.Exec(`CREATE SCHEMA public`) // Creo el schema tablas, contiene objetos (tablas, funciones, operadores, ...)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE cliente  (nrocliente 	int,
											 nombre 		text,
											 apellido 		text,
											 domicilio 		text,
											 telefono 		char(12));
											
					  CREATE TABLE tarjeta  (nrotarjeta 	char(16),
										 	 nrocliente 	int,
											 validadesde 	char(6),
											 validahasta 	char(6),
											 codseguridad 	char(4),
											 limitecompra 	decimal(8,2),
											 estado 		char(10));
											
 					  CREATE TABLE comercio (nrocomercio 	int,
											 nombre 		text,
											 domicilio 		text,
											 codigopostal 	char(8),
											 telefono 		char(12));
											
					  CREATE TABLE compra   (nrooperacion 	int,
										 	 nrotarjeta 	char(16),
											 nrocomercio 	int,
											 fecha 			timestamp,
											 monto 			decimal(7,2),
											 pagado 		boolean);
											
					  CREATE TABLE rechazo  (nrorechazo 	int,
											 nrotarjeta 	char(16),
											 nrocomercio 	int,
											 fecha 			timestamp,
											 monto 			decimal(7,2),
											 motivo 		text);
											
					  CREATE TABLE cierre   (año 			int,
											 mes 			int,
											 terminacion 	int,
											 fechainicio 	date,
											 fechacierre 	date,
											 fechavto 		date);
											
					  CREATE TABLE cabecera (nroresumen 	int,
											 nombre 		text,
											 apellido 		text,
											 domicilio 		text,
											 nrotarjeta 	char(16),
											 desde 			date,
											 hasta 			date,
											 vence 			date,
											 total 			decimal(8,2));
											
					  CREATE TABLE detalle  (nroresumen 	int,
											 nrolinea 		int,
											 fecha 			date,
											 nombrecomercio text,
											 monto 			decimal(7,2));
											
					  CREATE TABLE alerta   (nroalerta 		int,
										 	 nrotarjeta 	char(16),
											 fecha 			timestamp,
											 nrorechazo 	int,
											 codalerta 		int,
											 descripcion 	text);
											
					  CREATE TABLE consumo  (nrotarjeta 	char(16),
											 codseguridad 	char(4),
											 nrocomercio 	int,
											 monto 			decimal(7,2));`)
	if err != nil {
		log.Fatal(err)
	}
}

func borrarTablas() {
	_, err = db.Exec(`DROP SCHEMA IF EXISTS public CASCADE`) // Elimino todo lo que contenga el schema, tablas, funciones, etc.
	if err != nil {
		log.Fatal(err)
	}
}

func crearPKs() {
	_, err = db.Exec(`ALTER TABLE cliente ADD CONSTRAINT cliente_pk PRIMARY KEY (nrocliente);
					  ALTER TABLE tarjeta ADD CONSTRAINT tarjeta_pk PRIMARY KEY (nrotarjeta);
					  ALTER TABLE comercio ADD CONSTRAINT comercio_pk PRIMARY KEY (nrocomercio);
	                  ALTER TABLE compra ADD CONSTRAINT compra_pk PRIMARY KEY (nrooperacion);
	                  ALTER TABLE rechazo ADD CONSTRAINT rechazo_pk PRIMARY KEY (nrorechazo);
	                  ALTER TABLE cierre ADD CONSTRAINT cierre_pk PRIMARY KEY (año, mes, terminacion);
	                  ALTER TABLE cabecera ADD CONSTRAINT cabecera_pk PRIMARY KEY (nroresumen);
	                  ALTER TABLE detalle ADD CONSTRAINT detalle_pk PRIMARY KEY (nroresumen, nrolinea);
					  ALTER TABLE alerta ADD CONSTRAINT alerta_pk PRIMARY KEY (nroalerta);`)
	if err != nil {
		log.Fatal(err)
	}
}

func crearFKs() {
	_, err = db.Exec(`ALTER TABLE tarjeta ADD CONSTRAINT tarjeta_nrocliente_fk FOREIGN KEY (nrocliente) REFERENCES cliente(nrocliente);
					  ALTER TABLE compra ADD CONSTRAINT compra_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE compra ADD CONSTRAINT compra_nrocomercio_fk FOREIGN KEY (nrocomercio) REFERENCES comercio(nrocomercio);
					  ALTER TABLE rechazo ADD CONSTRAINT rechazo_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE rechazo ADD CONSTRAINT rechazo_nrocomercio_fk FOREIGN KEY (nrocomercio) REFERENCES comercio(nrocomercio);
					  ALTER TABLE cabecera ADD CONSTRAINT cabecera_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE detalle ADD CONSTRAINT detalle_cabecera_fk FOREIGN KEY (nroresumen) REFERENCES cabecera(nroresumen);
					  ALTER TABLE alerta ADD CONSTRAINT alerta_nrotarjeta_fk FOREIGN KEY (nrotarjeta) REFERENCES tarjeta(nrotarjeta);
					  ALTER TABLE alerta ADD CONSTRAINT alerta_nrorechazo_fk FOREIGN KEY (nrorechazo) REFERENCES rechazo(nrorechazo);`)
	if err != nil {
		log.Fatal(err)
	}
}

func borrarPKs() {
	_, err = db.Exec(`ALTER TABLE cliente DROP CONSTRAINT cliente_pk;
					  ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_pk;
					  ALTER TABLE comercio DROP CONSTRAINT comercio_pk;
	                  ALTER TABLE compra DROP CONSTRAINT compra_pk;
	                  ALTER TABLE rechazo DROP CONSTRAINT rechazo_pk;
	                  ALTER TABLE cierre DROP CONSTRAINT cierre_pk;
	                  ALTER TABLE cabecera DROP CONSTRAINT cabecera_pk;
	                  ALTER TABLE detalle DROP CONSTRAINT detalle_pk;
	                  ALTER TABLE alerta DROP CONSTRAINT alerta_pk;`)
	if err != nil {
		log.Fatal(err)
	}
}

func borrarFKs() {
	_, err = db.Exec(`ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_nrocliente_fk;
					  ALTER TABLE compra DROP CONSTRAINT compra_nrotarjeta_fk;
					  ALTER TABLE compra DROP CONSTRAINT compra_nrocomercio_fk;
					  ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrotarjeta_fk;
					  ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrocomercio_fk;
					  ALTER TABLE cabecera DROP CONSTRAINT cabecera_nrotarjeta_fk;
					  ALTER TABLE detalle DROP CONSTRAINT detalle_cabecera_fk;
					  ALTER TABLE alerta DROP CONSTRAINT alerta_nrotarjeta_fk;
					  ALTER TABLE alerta DROP CONSTRAINT alerta_nrorechazo_fk;`)
	if err != nil {
		log.Fatal(err)
	}
}
