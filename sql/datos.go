package sql

import(
	"log"

	_"github.com/lib/pq"
)

func cargarDatos() {
	_, err = db.Exec(`INSERT INTO cliente VALUES(11348773,'Rocío', 'Losada','Av. Presidente Perón 1530',1151102983);
					  INSERT INTO cliente VALUES(12349972,'María Estela', 'Martínez','Belgrano 1830',1150006655);
					  INSERT INTO cliente VALUES(22648991,'Laura', 'Santos','Italia 220',1153399452);
					  INSERT INTO cliente VALUES(11341003,'Graciela', 'Chasco','Tribulato 1340',1258579091);
					  INSERT INTO cliente VALUES(51558783,'Gabriela', 'Troncoso','Muñoz 1820',112234667);
					  INSERT INTO cliente VALUES(21347800,'Marta', 'Carbajo','San Luis 873',111998340);
					  INSERT INTO cliente VALUES(11448979,'Belén', 'Ferraris','Echeverría 780',113229087);
					  INSERT INTO cliente VALUES(44349773,'Abril', 'Hernández','Av. Sourdeaux 1700',115598342);
					  INSERT INTO cliente VALUES(33348679,'Sofía', 'Godoy','Av. SenadOR Morón 1221',114558004);
					  INSERT INTO cliente VALUES(25348533,'Adriana', 'Golluscio','Misiones 725',112111558);
					  INSERT INTO cliente VALUES(12228777,'Juan Carlos', 'Leguizamon','Serrano 120',1151101182);
					  INSERT INTO cliente VALUES(32680014,'Alberto', 'Ferrero','Pardo 990',1159944558);
					  INSERT INTO cliente VALUES(21545800,'Roberto', 'Ubertalli','Santa Fé 160',110076548);
					  INSERT INTO cliente VALUES(23679022,'Mario', 'Valdéz','Tucumán 550',116690874);
					  INSERT INTO cliente VALUES(12795452,'Ivana', 'Coronel','Río Diamante 186',113678652);
					  INSERT INTO cliente VALUES(11732790,'Bautista', 'Bello','Río Cuarto 191',111451419);
					  INSERT INTO cliente VALUES(29546643,'Diego', 'Fagnani','Av. Gaspar Campos 122',111009070);
					  INSERT INTO cliente VALUES(18397552,'Pedro', 'Tomarello','Av. San Martín 1511',110887547);
					  INSERT INTO cliente VALUES(13348765,'José', 'Mengarelli','Guido Spano 244',110044332);
					  INSERT INTO cliente VALUES(14348789,'Ricardo', 'Llanos','Corrientes 183',119034572);

					  INSERT INTO comercio VALUES(501,'Kevingston', 'Av. Tte. Gral. Ricchieri 965', 1661 ,46666181);
					  INSERT INTO comercio VALUES(523,'47 street', 'Paunero 1575', 1663 ,47597581);
					  INSERT INTO comercio VALUES(513,'Garbarino', 'Av. Bartolomé Mitre 1198', 1661 ,08104440018);
					  INSERT INTO comercio VALUES(521,'Bella Vista Hogar', 'Av. SenadOR Morón 1094', 1661 ,46661544);
					  INSERT INTO comercio VALUES(578,'Panadería y Confitería: La Princesa', 'Av. SenadOR Morón 1200', 1661 ,46681339);
					  INSERT INTO comercio VALUES(564,'FOX', 'Av Pres. Juan Domingo Perón 907', 1663 ,46676777);
					  INSERT INTO comercio VALUES(569,'La Pata Loca', 'Av. Moisés Lebensohn 98', 1661 ,46660861);
					  INSERT INTO comercio VALUES(545,'Frávega', 'Av. Pres. Juan Domingo Perón 1127', 1663 ,44512063);
					  INSERT INTO comercio VALUES(543,'Spit Bella Vista', 'Av. SenadOR Morón 1452', 1661 ,1153519765);
					  INSERT INTO comercio VALUES(527,'Óptica Cristal', 'Av. Dr. Ricardo Balbín 1125', 1663 ,46649400);
					  INSERT INTO comercio VALUES(508,'Óptica Mattaldi', 'Av. Mattaldi 1141', 1661 ,46683911);
					  INSERT INTO comercio VALUES(509,'Estancia San Francisco San Miguel', 'Concejal Tribulato 1265', 1663 ,5446676082);
					  INSERT INTO comercio VALUES(500,'Rabelia heladería', 'San José 972', 1663 ,46649352);
					  INSERT INTO comercio VALUES(520,'Heladería Ciwe', 'San José 785', 1663 ,46646003);
					  INSERT INTO comercio VALUES(588,'Rever Pass', 'Paunero 1447,', 1663 ,44513921);
					  INSERT INTO comercio VALUES(582,'Rapsodia', 'Av. Pres. Arturo Umberto Illia 3770', 1663 ,1160911581);
					  INSERT INTO comercio VALUES(530,'Grimoldi', 'Paunero 1415', 1663 ,44517343);
					  INSERT INTO comercio VALUES(596,'Umma', 'Paunero 1476', 1663 ,44519267);
					  INSERT INTO comercio VALUES(538,'COTO', 'Ohiggins 1280', 1661 ,46682636);
					  INSERT INTO comercio VALUES(553,'Disco', 'Av. SenadOR Morón 960', 1661 ,08107778888);

					  INSERT INTO tarjeta VALUES(4000001234567899,11348773, 201502, 202002, 733 ,50000,'vigente');
					  INSERT INTO tarjeta VALUES(4037001554363655,12349972, 201501, 202001, 332 ,55000,'vigente');
					  INSERT INTO tarjeta VALUES(4000001355435322,22648991, 201507, 202007, 201 ,60000,'vigente');
					  INSERT INTO tarjeta VALUES(4032011233774494,11341003, 201509, 202009, 204 ,90000,'vigente');
					  INSERT INTO tarjeta VALUES(4035055234867402,51558783, 201510, 202010, 108 ,50000,'vigente');
					  INSERT INTO tarjeta VALUES(4060001234507040,21347800, 201510, 202010, 909 ,10000,'vigente');
					  INSERT INTO tarjeta VALUES(4040071730767070,11448979, 201704, 202204, 810 ,57000,'vigente');
					  INSERT INTO tarjeta VALUES(4032002224865843,44349773, 201704, 202204, 327 ,64000,'vigente');
					  INSERT INTO tarjeta VALUES(4034006634262869,33348679, 201708, 202208, 097 ,99000,'suspendida');
					  INSERT INTO tarjeta VALUES(4034001232557669,25348533, 201708, 202208, 653 ,84000,'suspendida');
					  INSERT INTO tarjeta VALUES(4032002134557009,12228777, 201801, 202301, 070 ,99900,'vigente');
					  INSERT INTO tarjeta VALUES(4033002233062344,32680014, 201801, 202301, 202,90000,'anulada');
					  INSERT INTO tarjeta VALUES(4000006877865030,21545800, 201801, 202301, 115 ,80000,'vigente');
					  INSERT INTO tarjeta VALUES(4000001223567822,23679022, 201604, 202104, 559 ,70000,'vigente');
					  INSERT INTO tarjeta VALUES(4000001244532899,12795452, 201604, 202104, 842 ,59000,'vigente');
					  INSERT INTO tarjeta VALUES(4032003238867044,11732790, 201602, 202102, 379 ,73000,'vigente');
					  INSERT INTO tarjeta VALUES(4000002440217199,29546643, 201601, 202101, 794 ,62000,'vigente');
					  INSERT INTO tarjeta VALUES(4032000435566909,18397552, 201701, 202201, 621 ,59000,'suspendida');
					  INSERT INTO tarjeta VALUES(4037055274760805,13348765, 201712, 202212, 109 ,69000,'anulada');
					  INSERT INTO tarjeta VALUES(4000632234361811,13348765, 201709, 202209, 195 ,53000,'suspendida');
					  INSERT INTO tarjeta VALUES(4000000203465800,14348789, 201808, 202308, 290 ,78000,'vigente');
					  INSERT INTO tarjeta VALUES(4003300224374894,14348789, 201809, 202309, 284 ,84000,'vigente');`)
	if err != nil {
		log.Fatal(err)
	}
}

func borrarDatos(){
	_, err = db.Exec(`DELETE FROM tarjeta;
					  DELETE FROM cliente;
					  DELETE FROM comercio;
					  DELETE FROM cierre;`)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertarCierres() {
	insertarCierres()
	_, err = db.Query(`SELECT insertCierres();`)
	if err != nil {
		log.Fatal(err)
	}
}
