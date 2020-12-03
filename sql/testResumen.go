package sql

import (
	"log"
)

func testGenResumen() {

	_, err = db.Query(`

	SELECT generarresumen(14348789, 12, 2020); 
	SELECT generarresumen(12228777, 12, 2020);
	SELECT generarresumen(11448979, 12, 2020);  
	SELECT generarresumen(11732790, 12, 2020);  

	`)
	if err != nil {
		log.Fatal(err)
	}

	testResultCabecera()
	testResultDetalle()
}

func testResultCabecera() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION testCabecera() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			cierreTarjetaDEC record;
			cabeceraDEC record;
		BEGIN						
			ret := true;
			
			SELECT * INTO cabeceraDEC FROM cabecera WHERE nrotarjeta = '4032002134557009';
			SELECT * INTO cierreTarjetaDEC FROM cierre WHERE terminacion = 9 AND mes = 12;
			PERFORM * FROM compra WHERE nrotarjeta = cabeceraDEC.nrotarjeta AND fecha::date >=  cierreTarjetaDEC.fechainicio AND fecha::date <= cierreTarjetaDEC.fechacierre;
				
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

func testResultDetalle() {
	_, err = db.Query(`
		CREATE OR REPLACE FUNCTION testDetalle() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			cierreTarjetaDEC record;
			detalleDEC record;
			cabeceraDEC record;
			cabecera2DEC record;
			
		BEGIN						
			ret := true;
			
			SELECT * INTO cabeceraDEC FROM cabecera WHERE nrotarjeta = '4003300224374894';
			SELECT * INTO detalleDEC FROM detalle WHERE nroresumen = cabeceraDEC.nroresumen;
			SELECT * INTO cierreTarjetaDEC FROM cierre WHERE terminacion = 4 AND mes = 12;
			
			PERFORM * FROM compra WHERE nrotarjeta = cabeceraDEC.nrotarjeta AND cabeceraDEC.nroresumen = detalleDEC.nroresumen AND fecha::date >=  cierreTarjetaDEC.fechainicio AND fecha::date <= cierreTarjetaDEC.fechacierre AND monto = detalleDEC.monto;
			IF not found THEN
				ret := ret and false;				
			END IF;	
			
			----Persona que tiene todo pagado---
			
			SELECT * INTO cabecera2DEC FROM cabecera WHERE nrotarjeta = '4032002134557009';
			PERFORM * FROM compra WHERE nrotarjeta = '4032002134557009' AND pagado = true AND cabecera2DEC.total = 0;
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
