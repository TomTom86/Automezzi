package main

import (
	"automezzi/models"
	_ "automezzi/routers"
    "automezzi/controllers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "automezzi.db")
	orm.Debug = true
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	inizializeDBapp()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
    beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

func inizializeDBapp() {
	//inizialize db automezzi
	o := orm.NewOrm()

	//CONDITION
	var maps []orm.Params
	num, err := o.QueryTable("condizioni").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Condizioni")
		return
	}
	if num == 0 {
		condizioniArray := []string{"Buono stato", "Cattivo Stato", "Discreto Stato", "In Attesa di Alienazione", "In Attesa di Assegnazione", "In attesa di Riparazione", "Non utilizzabile", "Rubato", "Alienato"}

		for i := range condizioniArray {
			condizioni := models.Condizioni{ID: i + 1, Descrizione: condizioniArray[i]}

			_, err = o.Insert(&condizioni)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil
	//SPESA
	num, err = o.QueryTable("tipi_spesa").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella spesa")
		return
	}
	if num == 0 {

		tipiSpesaArray := []string{"Alienazione", "Assicurazione", "Bollo", "Contratto Canone", "Contratto varie", "Lavaggio", "Manutenzione ordinaria", "Pneumatici", "Revisione", "Riparazione per sinistro", "Riparazione straordinaria", "Varie"}
		for i := range tipiSpesaArray {

			tipiSpesa := models.TipiSpesa{ID: i + 1, Descrizione: tipiSpesaArray[i]}

			_, err = o.Insert(&tipiSpesa)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//TIPO VEICOLO

	num, err = o.QueryTable("tipi_veicolo").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella vehicle_type")
		return
	}
	if num == 0 {
		tipiVeicoloArray := []string{"Autoveicolo", "Camion", "Bollo", "Ciclomotore", "Furgone", "Pullman", "Motoveicolo", "Altro"}

		for i := range tipiVeicoloArray {

			tipiVeicolo := models.TipiVeicolo{ID: i + 1, Descrizione: tipiVeicoloArray[i]}
			_, err = o.Insert(&tipiVeicolo)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Sector

	num, err = o.QueryTable("settori").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Settori")
		return
	}
	if num == 0 {

		settoriArray := []string{"Food", "Lavanderia", "Pulizia", "Marketing", "Officina", "Agenti", "Direzione"}

		for i := range settoriArray {

			settori := models.Settori{ID: i + 1, Descrizione: settoriArray[i]}
			_, err = o.Insert(&settori)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Employment

	num, err = o.QueryTable("impieghi").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Employment")
		return
	}
	if num == 0 {

		impiegoArray := []string{"Aziendale", "Aziendale + Personale", "Personale"}

		for i := range impiegoArray {

			impiego := models.Impieghi{ID: i + 1, Descrizione: impiegoArray[i]}
			_, err = o.Insert(&impiego)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Responsabilita

	num, err = o.QueryTable("responsabilita").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Responsabilita")
		return
	}
	if num == 0 {

		responsabilitaArray := []string{"Concorso di colpa", "Da accertare", "Della controparte", "Propria"}

		for i := range responsabilitaArray {

			responsabilita := models.Responsabilita{ID: i + 1, Descrizione: responsabilitaArray[i]}
			_, err = o.Insert(&responsabilita)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//tipo_infrazione

	num, err = o.QueryTable("tipi_infrazione").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella tipo_infrazione")
		return
	}
	if num == 0 {

		tipiInfrazioneArray := []string{"Accesso in senso vietato", "Accesso in senso vietato", "Cinture di sicurezza non allacciate", "Eccesso di velocità", "Guida contromano", "Guida pericolosa", "Positivo a controlllo alcool", "Precedenza non rispettata", "Semaforo rosso", "Sosta vietata", "Utilizzo di telefono cellulare", "Violazione di corsia preferenziale", "Violazione di ztl", "Altro"}

		for i := range tipiInfrazioneArray {

			tipiInfrazione := models.TipiInfrazione{ID: i + 1, Descrizione: tipiInfrazioneArray[i]}
			_, err = o.Insert(&tipiInfrazione)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

	//Carburante

	num, err = o.QueryTable("carburante").Values(&maps, "ID")
	if err != nil {
		fmt.Println("Errore inizializzazione tabella Carburante")
		return
	}
	if num == 0 {

		carburanteArray := []string{"Benzina", "Diesel", "Gas", "Metano"}

		for i := range carburanteArray {

			carburante := models.Carburante{ID: i + 1, Descrizione: carburanteArray[i]}
			_, err = o.Insert(&carburante)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
	num = 0
	err = nil

}
