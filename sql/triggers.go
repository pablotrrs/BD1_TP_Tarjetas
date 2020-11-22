package sql

import(
	"log"
)

func trgCargarAlerta() {
	_, err = db.Query(
		`	DROP TRIGGER IF EXISTS cargaralerta_trg ON rechazo;
		
			CREATE trigger cargaralerta_trg
			AFTER INSERT ON rechazo
			FOR EACH ROW
			EXECUTE PROCEDURE cargar_alerta();`)
	if err != nil {
		log.Fatal(err)
	}
}

func trgTiempoCompras() {
	_, err = db.Query(
		`	DROP TRIGGER IF EXISTS tiempo_compras_trg ON compra;
		
			CREATE trigger tiempo_compras_trg
			BEFORE INSERT ON compra
			FOR EACH ROW
			EXECUTE PROCEDURE compras_tiempo();`)
	if err != nil {
		log.Fatal(err)
	}
}
