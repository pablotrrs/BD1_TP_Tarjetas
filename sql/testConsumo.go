package sql

import (
	"log"
)

func generarConsumos() {
	_, err = db.Query(`
						INSERT INTO consumo VALUES(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(501 AS int), CAST(1200 AS decimal(7,2))); --compra						
						INSERT INTO consumo VALUES(CAST(9000001234567899 AS char(16)), CAST(733 AS char(4)), CAST(666 AS int), CAST(1802 AS decimal(7,2))); --tarjeta inexistente
						INSERT INTO consumo VALUES(CAST(4033002233062344 AS char(16)), CAST(202 AS char(4)), CAST(999 AS int), CAST(1500 AS decimal(7,2))); --tarjeta no vigente
						INSERT INTO consumo VALUES(CAST(4034006634262869 AS char(16)), CAST(097 AS char(4)), CAST(222 AS int), CAST(3012 AS decimal(7,2))); --tarjeta suspendida
						INSERT INTO consumo VALUES(CAST(4000001234567899 AS char(16)), CAST(111 AS char(4)), CAST(500 AS int), CAST(12501 AS decimal(7,2))); --codigo incorrecto
						INSERT INTO consumo	VALUES(CAST(4032002134557009 AS char(16)), CAST(070 AS char(4)), CAST(569 AS int), CAST(25000 AS decimal(7,2))); -- para que la siguiente query sea efectiva
						INSERT INTO consumo VALUES(CAST(4032002134557009 AS char(16)), CAST(070 AS char(4)), CAST(569 AS int), CAST(50001 AS decimal(7,2))); --supera monto primera vez
						INSERT INTO consumo VALUES(CAST(4000001234567899 AS char(16)), CAST(733 AS char(4)), CAST(400 AS int), CAST(22500 AS decimal(7,2))); --vencida
						INSERT INTO consumo VALUES(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(501 AS int), CAST(1500 AS decimal(7,2))); -- alerta de 1 min				
						INSERT INTO consumo VALUES(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(123 AS int), CAST(300 AS decimal(7,2))); -- alerta de 1 min										
						INSERT INTO consumo VALUES(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(501 AS int), CAST(3200 AS decimal(7,2))); -- alerta de 5 min		
						INSERT INTO consumo VALUES(CAST(4003300224374894 AS char(16)), CAST(284 AS char(4)), CAST(666 AS int), CAST(3400 AS decimal(7,2))); -- alerta de 5 min			
						INSERT INTO consumo VALUES(CAST(4032002134557009 AS char(16)), CAST(070 AS char(4)), CAST(569 AS int), CAST(50001 AS decimal(7,2))); --supera monto segunda vez 
						INSERT INTO consumo VALUES(CAST(4032002134557009 AS char(16)), CAST(070 AS char(4)), CAST(569 AS int), CAST(50001 AS decimal(7,2))); --supera monto tercera vez y creo alerta de cambio de estado
						
						INSERT INTO consumo VALUES(CAST(4040071730767070 AS char(16)), CAST(810 AS char(4)), CAST(678 AS int), CAST(750 AS decimal(7,2)));
						INSERT INTO consumo VALUES(CAST(4040071730767070 AS char(16)), CAST(810 AS char(4)), CAST(345 AS int), CAST(800 AS decimal(7,2)));
						INSERT INTO consumo VALUES(CAST(4040071730767070 AS char(16)), CAST(810 AS char(4)), CAST(500 AS int), CAST(120 AS decimal(7,2)));				
						INSERT INTO consumo VALUES(CAST(4032003238867044 AS char(16)), CAST(379 AS char(4)), CAST(888 AS int), CAST(2000 AS decimal(7,2)));
						INSERT INTO consumo VALUES(CAST(4032003238867044 AS char(16)), CAST(379 AS char(4)), CAST(333 AS int), CAST(140 AS decimal(7,2)));
						INSERT INTO consumo VALUES(CAST(4032003238867044 AS char(16)), CAST(379 AS char(4)), CAST(538 AS int), CAST(250 AS decimal(7,2)));
						INSERT INTO consumo VALUES(CAST(4032003238867044 AS char(16)), CAST(379 AS char(4)), CAST(500 AS int), CAST(1200 AS decimal(7,2)));`)
	if err != nil {
		log.Fatal(err)
	}
}

func testFunciones() {
	consumir()
	testCompra()
	testAutorizaciones()
	testAlertas()
	testAll()
}

func testAll() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION test_all() RETURNS boolean AS $$
		DECLARE
			compras boolean;
			autorizaciones boolean;
			alertas boolean;	
			ret boolean;		
		BEGIN
		
			SELECT test_compras() INTO compras;
			SELECT test_autorizaciones() INTO autorizaciones;
			SELECT test_alertas() INTO alertas;			
			
			ret := compras and autorizaciones and alertas;
			
			return ret;
		END
		$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}

func consumir() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION consumir() RETURNS void AS $$
		DECLARE
			v_consumo record;
		BEGIN
			
			FOR v_consumo IN SELECT * FROM consumo LOOP
				PERFORM autorizacion_compra(v_consumo.nrotarjeta, v_consumo.codseguridad, v_consumo.nrocomercio, v_consumo.monto);
				PERFORM * FROM compra WHERE nrotarjeta = '4032002134557009' and monto = 25000;	
				IF found THEN
					UPDATE compra SET pagado = true where nrotarjeta = '4032002134557009' and monto = 25000;   
				END IF;
			END LOOP;
			
		END
		$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(`SELECT consumir();`)
	if err != nil {
		log.Fatal(err)
	}
}

func testCompra() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION test_compras() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			v_consumo consumo%rowtype;
		BEGIN						
			ret := true;
			
			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '4003300224374894';
			PERFORM * FROM compra WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto;
				
			IF not found THEN
				ret := ret and false;				
			END IF;	
			
			return ret;
			
		END
		$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}

func testAutorizaciones() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION test_autorizaciones() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			cant_rechazos_limite_compra int;
			v_consumo consumo%rowtype;
		BEGIN			
			ret := true;
			cant_rechazos_limite_compra	:= 0;
			
			-- Tarjeta no valida o no vigente
			
			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '9000001234567899';
			PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'tarjeta no valida o no vigente';
			
			IF not found THEN
				ret := ret and false;				
			END IF;	

			-- Tarjeta no valida o no vigente
			
			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '4033002233062344';
			PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'tarjeta no valida o no vigente';
			
			IF not found THEN
				ret := ret and false;				
			END IF;			

			-- Tarjeta suspendida
			
			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '4034006634262869';
			PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'la tarjeta se encuentra suspendida';
			
			IF not found THEN
				ret := ret and false;				
			END IF;				

			-- Codigo de seguridad invalido

			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '4000001234567899' and codseguridad = '111';
			PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'codigo de seguridad invalido';
			
			IF not found THEN
				ret := ret and false;				
			END IF;				

			-- Plazo de vigencia expirado
			
			SELECT * INTO v_consumo FROM consumo WHERE nrotarjeta = '4000001234567899' and codseguridad = '733';
			PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'plazo de vigencia expirado';
			
			IF not found THEN
				ret := ret and false;				
			END IF;					
			
			-- Supera limite de tarjeta
					
			FOR v_consumo IN SELECT * FROM consumo WHERE nrotarjeta = '4032002134557009' and monto != 25000 LOOP
				PERFORM * FROM rechazo WHERE nrotarjeta = v_consumo.nrotarjeta and nrocomercio = v_consumo.nrocomercio and monto = v_consumo.monto and motivo = 'supera limite de tarjeta';
				IF not found THEN
					ret := ret and false;
				ELSE
					cant_rechazos_limite_compra := cant_rechazos_limite_compra + 1;
				END IF;										
			END LOOP;			
			
			IF cant_rechazos_limite_compra != 3 THEN
				ret := ret and false;
			END IF;
		
			return ret;
		END
		$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}

func testAlertas() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION test_alertas() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			v_rechazo rechazo%rowtype;
		BEGIN			
			ret := true;
			
			-- Alertas contiene a los rechazos
			
			FOR v_rechazo IN SELECT * FROM rechazo LOOP
				PERFORM * FROM alerta WHERE nrotarjeta = v_rechazo.nrotarjeta and descripcion = v_rechazo.motivo;
				
				IF not found THEN
					ret := ret and false;
				END IF;										
			END LOOP;			
			
			-- Alerta por menos de 1 minuto
		
			PERFORM * FROM alerta WHERE nrotarjeta = '4003300224374894' and codalerta = 1 and descripcion = 'compra en menos de 1 minuto en una misma zona';
			
			IF not found THEN
				ret := ret and false;				
			END IF;			
						
			-- Alerta por menos de 5 minutos

			PERFORM * FROM alerta WHERE nrotarjeta = '4003300224374894' and codalerta = 5 and descripcion = 'compra en menos de 5 minutos en diferentes zonas';
			
			IF not found THEN
				ret := ret and false;				
			END IF;			

			-- Alerta por dos rechazos en limite de compra excedido, seguida del cambio de estado de la tarjeta
			
			PERFORM * FROM alerta WHERE nrotarjeta = '4032002134557009' and codalerta = 32 and descripcion = 'suspencion preventiva';
			
			IF not found THEN
				ret := ret and false;				
			END IF;	
			
			PERFORM * FROM tarjeta WHERE nrotarjeta = '4032002134557009' and estado = 'suspendida';

			IF not found THEN
				ret := ret and false;				
			END IF;	
		
			return ret;
		END
		$$ LANGUAGE PLPGSQL;`)
	if err != nil {
		log.Fatal(err)
	}
}
