package sql

import(
	"log"
)

//generate_series() genera series según el argumento pasado. Para cada ciclo del for genera valores distintos.
func insertarCierres() {
	_, err = db.Query(`
	CREATE OR REPLACE FUNCTION insertcierres() RETURNS void AS $$
	BEGIN
		FOR i in 0..9 LOOP
			INSERT INTO cierre VALUES(2020,generate_series(1,12),i,
			generate_series('2020/01/01'::date,'2020/12/31','1 month'),
			generate_series('2020/01/10'::date,'2020/12/31','1 month'),
			generate_series('2020/01/25'::date,'2020/12/31','1 month')
			);
		END LOOP;
	END
	$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}	
}

func autorizacionCompra(){
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION cargar_rechazo(numtarjeta char(16), numcomercio int, montocompra decimal(7,2), mensaje text) RETURNS void AS $$
		BEGIN
			INSERT INTO rechazo VALUES(nextval('seq_nrorechazo'), numtarjeta, numcomercio, CURRENT_DATE, montocompra, mensaje);
		END
		$$ LANGUAGE PLPGSQL;`)	

	if err != nil {
		log.Fatal(err)
	}	
	
	_, err = db.Query(`
	CREATE OR REPLACE FUNCTION autorizacion_compra(numtarjeta char(16), codseg char(4), numcomercio int, montocompra decimal(7,2)) RETURNS boolean AS $$
	DECLARE
		ret boolean;
		tarj record;
	
	BEGIN
		ret := true;
		
		------------------
		--    Caso 1    --
		
		--Numero tarjeta inexistente--
		SELECT * INTO tarj FROM tarjeta WHERE nrotarjeta = numtarjeta;
		
		IF not found THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'tarjeta no valida o no vigente');
			ret := false;
		END IF;
		
		--Tarjeta no esta vigente--
		
		IF tarj.estado != 'vigente' THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'tarjeta no valida o no vigente');
			ret := false;
		END IF;
		
		--              --
		------------------
		
		------------------
		--    Caso 2    --
		
		-- Codigo de seguridad incorrecto --
		
		IF tarj.codseguridad != codseg THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'codigo de seguridad invalido');			
			ret := false;
		END IF;
		
		--              --
		------------------
		
		------------------
		--    Caso 5    --
		
		--Tarjeta suspendida--
		
		IF tarj.estado = 'suspendida' THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'la tarjeta se encuentra suspendida');
			ret := false;
		END IF;	
		
		--              --
		------------------
		
		------------------
		--Compra exitosa--
		
		IF ret = true THEN
			INSERT INTO compra VALUES(nextval('seq_nrocompra'), numtarjeta, numcomercio, CURRENT_DATE, montocompra, true);
		END IF;
		
		--              --
		------------------
		return ret;
	END
	$$ LANGUAGE PLPGSQL;`)

	if err != nil {
		log.Fatal(err)
	}
}
