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
			INSERT INTO rechazo VALUES(nextval('seq_nrorechazo'), numtarjeta, numcomercio, CURRENT_TIMESTAMP, montocompra, mensaje);
			
		END
		$$ LANGUAGE PLPGSQL;`)	

	if err != nil {
		log.Fatal(err)
	}	
	
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION chequear_cantidad_rechazos(numtarjeta char(16)) RETURNS void AS $$
		DECLARE
			cantidad_rechazos int;
		
		BEGIN
			SELECT COUNT(numtarjeta) INTO cantidad_rechazos FROM rechazo WHERE nrotarjeta = numtarjeta AND motivo ='supera limite de tarjeta' AND DATE_PART('day', fecha) = DATE_PART('day', CURRENT_TIMESTAMP);
				
			IF cantidad_rechazos > 1 THEN
				UPDATE tarjeta SET estado = 'suspendida' where nrotarjeta = numtarjeta;   
				INSERT INTO alerta VALUES(nextval('seq_nroalerta'), numtarjeta, CURRENT_TIMESTAMP, null, 32, 'suspencion preventiva'); 
			
			END IF;
			
		END
		$$ LANGUAGE PLPGSQL;`)	

	if err != nil {
		log.Fatal(err)
	}	

	_, err = db.Query(`
	CREATE OR REPLACE FUNCTION autorizacion_compra(numtarjeta char(16), codseg char(4), numcomercio int, montocompra decimal(7,2)) RETURNS boolean AS $$
	DECLARE
		tarj record;
		monto_compras_pendientes int;
		monto_total int;
		ano_actual char(6);
		mes_actual char(6);
		fecha_actual char(6);
	
	BEGIN
		
		------------------
		--    Caso 1    --
		
		--Numero tarjeta inexistente--
		SELECT * INTO tarj FROM tarjeta WHERE nrotarjeta = numtarjeta;
		
		IF not found THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'tarjeta no valida o no vigente');
			return false;
		END IF;
		
		--Tarjeta no esta vigente--
		
		IF tarj.estado != 'vigente' AND tarj.estado != 'suspendida' THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'tarjeta no valida o no vigente');
			return false;
		END IF;
		
		--              --
		------------------

		------------------
		--    Caso 5    --
		
		--Tarjeta suspendida--
		
		IF tarj.estado = 'suspendida' THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'la tarjeta se encuentra suspendida');
			return false;
		END IF;	
		
		--              --
		------------------
		
		------------------
		--    Caso 2    --
		
		-- Codigo de seguridad incorrecto --
		
		IF tarj.codseguridad != codseg THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'codigo de seguridad invalido');
			return false;
		END IF;
		
		--              --
		------------------

		------------------
		--    Caso 3    --
		
		------------------
		--    Caso 4    --
		
		-- Tarjeta vencida --
		
		SELECT DATE_PART('year', (SELECT CURRENT_DATE)) INTO ano_actual; 
		SELECT DATE_PART('month', (SELECT CURRENT_DATE)) INTO mes_actual;
		fecha_actual := ano_actual || mes_actual;
		
		IF tarj.validahasta < fecha_actual THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'plazo de vigencia expirado');
			return false;
		END IF;
		
		--              --
		------------------

		-- Limite de compra superado --
		
		SELECT SUM(monto) INTO monto_compras_pendientes FROM compra WHERE tarj.nrotarjeta = numtarjeta AND pagado = false;
		monto_total := monto_compras_pendientes + montocompra;
		
		IF tarj.limitecompra < monto_total THEN
			PERFORM cargar_rechazo(CAST(numtarjeta AS char(16)), CAST(numcomercio AS int), CAST(montocompra AS decimal(7,2)), 'supera limite de tarjeta');
			PERFORM chequear_cantidad_rechazos(CAST(numtarjeta AS char(16)));
			return false;
		END IF;

		--              --
		------------------
			
		------------------
		--Compra exitosa--
		
		INSERT INTO compra VALUES(nextval('seq_nrocompra'), numtarjeta, numcomercio, CURRENT_TIMESTAMP, montocompra, true);
		
		--              --
		------------------
		return true;
	END
	$$ LANGUAGE PLPGSQL;`)

	if err != nil {
		log.Fatal(err)
	}
}

func crearTriggers(){
	cargar_alerta()
	triggerstiempo()
}

func cargar_alerta(){
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION cargar_alerta() RETURNS trigger AS $$
		BEGIN

			INSERT INTO alerta VALUES(nextval('seq_nroalerta'), new.nrotarjeta, new.fecha, new.nrorechazo, 0, new.motivo);
			
		return new;			
		END
		$$ LANGUAGE PLPGSQL;`)	

	if err != nil {
		log.Fatal(err)
	}		
	
	trgCargarAlerta();
}

func triggerstiempo(){
		_, err = db.Query(`
		CREATE OR REPLACE FUNCTION compras_tiempo() RETURNS trigger AS $$
		DECLARE
			ultima_compra record;
			diferencia_tiempo decimal;
			cod_postal_anterior int;
			cod_postal_actual int;
			
		BEGIN
			SELECT * INTO ultima_compra FROM compra WHERE nrotarjeta = new.nrotarjeta ORDER BY nrooperacion DESC LIMIT 1;
			
			IF not found THEN
				return new;
			END IF;
						
			SELECT INTO diferencia_tiempo EXTRACT(EPOCH FROM (new.fecha - ultima_compra.fecha)) / 60;
			SELECT codigopostal INTO cod_postal_anterior FROM comercio WHERE nrocomercio = ultima_compra.nrocomercio;
			SELECT codigopostal INTO cod_postal_actual FROM comercio WHERE nrocomercio = new.nrocomercio;
			
			--Alerta por compras en menos de 1 minuto comercios con el mismo codigo postal
			
			IF diferencia_tiempo < 1 and ultima_compra.nrocomercio != new.nrocomercio and cod_postal_anterior = cod_postal_actual THEN
				INSERT INTO alerta VALUES(nextval('seq_nroalerta'), new.nrotarjeta, CURRENT_TIMESTAMP, null, 1, 'compra en menos de 1 minuto en una misma zona');
				return new;
			END IF;

			--Alerta por compras en menos de 5 minutos en comercios con diferentes codigos postales
			
			IF diferencia_tiempo < 5 and ultima_compra.nrocomercio != new.nrocomercio and cod_postal_anterior != cod_postal_actual THEN
				INSERT INTO alerta VALUES(nextval('seq_nroalerta'),new.nrotarjeta, CURRENT_TIMESTAMP, null, 5, 'compra en menos de 5 minutos en diferentes zonas');
				return new;
			END IF;
			
			
		return new;			
		END
		$$ LANGUAGE PLPGSQL;`)	

	if err != nil {
		log.Fatal(err)
	}		
	
	trgTiempoCompras();
}

func generarResumen(){
		_, err = db.Query(`
	CREATE OR REPLACE FUNCTION generarresumen(numCliente int, mesIN int) RETURNS void AS $$

	DECLARE
		clienteDEC record;
		tarjetaDEC record;
		contResumen int;
		nomComercioDEC record;
		compraDEC record;
		contLinea int;
		montofinal decimal(9,2);
		cierreTarjetaDEC record;
		
		
	BEGIN
			contLinea := 1;
			montofinal := 0;
	
			SELECT * INTO clienteDEC FROM cliente WHERE nrocliente = numCLiente;
			SELECT * INTO tarjetaDEC FROM tarjeta WHERE nrocliente = numCLiente;	
			SELECT * INTO cierreTarjetaDEC FROM cierre WHERE mes = mesIN and terminacion = substring(tarjetaDEC.nrotarjeta,16)::int;
			
			contResumen := 0;
			contResumen := contResumen + count(*) from cabecera;
				
			INSERT INTO cabecera VALUES (contResumen + 1, 
										clienteDEC.nombre,
										clienteDEC.apellido,
										clienteDEC.domicilio,
										tarjetaDEC.nrotarjeta,
										cierreTarjetaDEC.fechainicio, 
										cierreTarjetaDEC.fechacierre, 
										cierreTarjetaDEC.fechavto,
										montofinal
										);														
			
									
			FOR compraDEC IN SELECT * FROM compra WHERE nrotarjeta = tarjetaDEC.nrotarjeta AND pagado = false AND fecha::date >=  cierreTarjetaDEC.fechainicio AND fecha::date <= cierreTarjetaDEC.fechacierre
			LOOP	 
				SELECT * INTO nomComercioDEC FROM comercio WHERE nrocomercio = compraDEC.nrocomercio;
				INSERT INTO detalle VALUES (contResumen + 1, 
											contLinea,
											compraDEC.fecha,			
											nomComercioDEC.nombre, 
											compraDEC.monto
											);	
				contLinea := contLinea + 1;	
				montofinal := montofinal + compraDEC.monto;							
			END LOOP;	
			
			UPDATE cabecera SET total = montofinal WHERE nrotarjeta = tarjetaDEC.nrotarjeta	AND desde =	cierreTarjetaDEC.fechainicio AND hasta = cierreTarjetaDEC.fechacierre;										
							
			
	END
$$ LANGUAGE PLPGSQL;`)

	if err != nil {
		log.Fatal(err)
	}		
		
}
