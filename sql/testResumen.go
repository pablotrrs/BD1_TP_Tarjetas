package sql

import(
	"log"
)

func testGenResumen(){
	generarcompras()
	
	_, err = db.Query(`
	
	SELECT generarresumen(22648991, 09); 
	SELECT generarresumen(11341003, 11); 
	SELECT generarresumen(51558783, 06); 
	SELECT generarresumen(21347800, 01); 
	SELECT generarresumen(11448979, 11); 
	

	`)	
	if err != nil {
		log.Fatal(err)
	}
	
	testResultCabecera()
	testResultDetalle()
}


func generarcompras() {
	_, err = db.Query(`
	
	--CLiente 22648991 -> Fretes Matias--
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4000001355435322, 538, '2020/09/10'::date, 20, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4000001355435322, 222, '2020/09/10'::date, 10, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4000001355435322, 345, '2020/09/10'::date, 50, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4000001355435322, 999, '2020/09/30'::date, 50, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4000001355435322, 588, '2020/09/05'::date, 80, false); 
	
	
	--CLiente 11341003 -> Roman Chasco--
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4032011233774494, 538, '2020/11/10'::date, 135, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4032011233774494, 222, '2020/11/27'::date, 280, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4032011233774494, 345, '2020/11/27'::date, 300, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4032011233774494, 999, '2020/11/30'::date, 1000, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4032011233774494, 588, '2020/11/05'::date, 700, false); 
	
	--CLiente 51558783 -> Diego Troncoso--
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4035055234867402, 400, '2020/06/06'::date, 1200, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4035055234867402, 527, '2020/06/28'::date, 510, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4035055234867402, 345, '2020/07/01'::date, 370, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4035055234867402, 999, '2020/06/30'::date, 100, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4035055234867402, 588, '2020/06/05'::date, 250, false); 
	
	--CLiente 21347800 -> Marta Carbajo--
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4060001234507040, 538, '2020/01/01'::date, 500, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4060001234507040, 222, '2020/01/02'::date, 350, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4060001234507040, 345, '2020/01/05'::date, 800, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4060001234507040, 999, '2020/01/08'::date, 140, true); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4060001234507040, 588, '2020/01/09'::date, 700, true); 
		
	--CLiente 11448979 -> Belen Ferraris--
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4040071730767070, 333, '2020/11/10'::date, 200, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4040071730767070, 888, '2020/11/11'::date, 100, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4040071730767070, 500, '2020/11/01'::date, 500, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4040071730767070, 582, '2020/11/30'::date, 500, false); 
	INSERT INTO compra VALUES(nextval('seq_nrocompra'), 4040071730767070, 530, '2020/11/05'::date, 800, false);	
		
		
	`)	
	if err != nil {
		log.Fatal(err)
	}	
}

func testResultCabecera(){
		_, err = db.Query(`
		CREATE OR REPLACE FUNCTION testCabecera() RETURNS boolean AS $$
		DECLARE
			ret boolean;
			cierreTarjetaDEC record;
			cabeceraDEC record;
		BEGIN						
			ret := true;
			
			SELECT * INTO cabeceraDEC FROM cabecera WHERE nrotarjeta = '4000001355435322';
			SELECT * INTO cierreTarjetaDEC FROM cierre WHERE terminacion = 2 AND mes = 09;
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


func testResultDetalle(){
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
			
			SELECT * INTO cabeceraDEC FROM cabecera WHERE nrotarjeta = '4032011233774494';
			SELECT * INTO detalleDEC FROM detalle WHERE nroresumen = cabeceraDEC.nroresumen;
			SELECT * INTO cierreTarjetaDEC FROM cierre WHERE terminacion = 4 AND mes = 11;
			
			PERFORM * FROM compra WHERE nrotarjeta = cabeceraDEC.nrotarjeta AND cabeceraDEC.nroresumen = detalleDEC.nroresumen AND fecha::date >=  cierreTarjetaDEC.fechainicio AND fecha::date <= cierreTarjetaDEC.fechacierre AND monto = detalleDEC.monto and pagado = false;
			IF not found THEN
				ret := ret and false;				
			END IF;	
			
			
			
			--Persona que pago todo--
			
			SELECT * INTO cabecera2DEC FROM cabecera WHERE nrotarjeta = '4060001234507040';
			PERFORM * FROM compra WHERE nrotarjeta = cabecera2DEC.nrotarjeta AND pagado = true AND cabecera2DEC.total = 0;	
			
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
