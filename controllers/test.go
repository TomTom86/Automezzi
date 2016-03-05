package controllers

import (
	//"automezzi/models"
	//pk "automezzi/utilities/pbkdf2wrapper"
	//"encoding/hex"
	"fmt"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	//"github.com/twinj/uuid"
	//"html/template"
	"strconv"
	//"strings"
	//"time"
    //"reflect"
)
//Test is for experiment
func (c *MainController) Test() {
	c.activeContent("test/test")

	parms,  err:= strconv.Atoi(c.Ctx.Input.Param(":parms"))
	fmt.Println(parms)   
    if err != nil {
		fmt.Println(err)
	}

    o := orm.NewOrm()
    o.Using("default")
 /*  
   //CONDITION
    var maps []orm.Params
    num, err := o.QueryTable("condition").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella Condition")
        return
    }
    fmt.Println("Condition ",num)    
    if num == 0 {
        conditionArray := []string{"Buono stato","Cattivo Stato", "Discreto Stato", "In Attesa di Alienazione","In Attesa di Assegnazione", "In attesa di Riparazione","Non utilizzabile", "Rubato", "Alienato"}

        for i := range conditionArray { 
            condition := models.Condition{Id: i+1,Description : conditionArray[i]}
            
            _, err = o.Insert(&condition)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }
    }
    num = 0 
    err = nil
    //SPESA
    num, err = o.QueryTable("spesa").Values(&maps, "id")
    fmt.Println(num)
    if err != nil {
        fmt.Println("Errore inizializzazione tabella spesa")
        return
    }
    if num == 0 {
     
        spesaArray := []string{ 1: "Alienazione", "Assicurazione", "Bollo", "Contratto Canone", "Contratto varie", "Lavaggio", "Manutenzione ordinaria", "Pneumatici", "Revisione", "Riparazione per sinistro", "Riparazione straordinaria", "Varie"}    
        for i := range spesaArray {
        
            spesa := models.Spesa{Id: i+1,Descrizione : spesaArray[i]}
            
            _, err = o.Insert(&spesa)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }  
    }        
    num = 0 
    err = nil
    
    //TIPO VEICOLO

    num, err = o.QueryTable("vehicle_type").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella vehicle_type")
        return
    }
    if num == 0 {    
        tipoVeicoloArray := []string{ "Autoveicolo", "Camion", "Bollo", "Ciclomotore", "Furgone", "Pullman", "Motoveicolo","Altro"}    
                                            
        for i := range tipoVeicoloArray {
        
            tipoVeicolo := models.VehicleType{Id: i+1,Description : tipoVeicoloArray[i]}
            _, err = o.Insert(&tipoVeicolo)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        } 
    }        
    num = 0 
    err = nil
    
    //Sector

    num, err = o.QueryTable("sector").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella Sector")
        return
    }
    if num == 0 {                 
    
        sectorArray := []string{ "Food", "Lavanderia", "Pulizia", "Marketing", "Officina", "Agenti", "Direzione"}    
                                            
        for i := range sectorArray {
        
            sector := models.Sector{Id: i+1,Description : sectorArray[i]}
            _, err = o.Insert(&sector)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }    
    }        
    num = 0 
    err = nil
    
    //Employment

    num, err = o.QueryTable("employment").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella Employment")
        return
    }
    if num == 0 {          

        impiegoArray := []string{ "Aziendale", "Aziendale + Personale", "Personale"}    
                                            
        for i := range impiegoArray {
        
            impiego := models.Employment{Id: i+1,Description : impiegoArray[i]}
            _, err = o.Insert(&impiego)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }
    }        
    num = 0 
    err = nil
    
    //Responsabilita

    num, err = o.QueryTable("responsabilita").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella Responsabilita")
        return
    }
    if num == 0 {              

        responsabilitaArray := []string{ "Concorso di colpa", "Da accertare", "Della controparte", "Propria"}    
                                            
        for i := range responsabilitaArray {
        
            responsabilita := models.Responsabilita{Id: i+1,Descrizione : responsabilitaArray[i]}
            _, err = o.Insert(&responsabilita)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }  
    }        
    num = 0 
    err = nil
    
    //tipo_infrazione

    num, err = o.QueryTable("tipo_infrazione").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella tipo_infrazione")
        return
    }
    if num == 0 {            

        tipoInfrazioneArray := []string{ 1: "Accesso in senso vietato", "Accesso in senso vietato", "Cinture di sicurezza non allacciate", "Eccesso di velocità", "Guida contromano", "Guida pericolosa", "Positivo a controlllo alcool", "Precedenza non rispettata", "Semaforo rosso", "Sosta vietata", "Utilizzo di telefono cellulare", "Violazione di corsia preferenziale", "Violazione di ztl", "Altro"}    
        
        for i := range tipoInfrazioneArray {
        
            tipoInfrazione := models.TipoInfrazione{Id: i+1,Descrizione : tipoInfrazioneArray[i]}
            _, err = o.Insert(&tipoInfrazione)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        }
    }        
    num = 0 
    err = nil
    
    //Carburante

    num, err = o.QueryTable("carburante").Values(&maps, "id")
    if err != nil {
        fmt.Println("Errore inizializzazione tabella Carburante")
        return
    }
    if num == 0 {              

        carburanteArray := []string{ 1: "Benzina", "Diesel", "Gas", "Metano"}  
        
        for i := range carburanteArray {
        
            carburante := models.Carburante{Id: i+1,Descrizione : carburanteArray[i]}
            _, err = o.Insert(&carburante)
            if err != nil {
                fmt.Println("Errore autorizzazioni applicazioni")
                return
            }
            
        } 
    }        
    num = 0 
    err = nil       

/*   
    
    VehicleDG 
	VehicleDT 
	Carburante 
	VehicleType
	Sector
	Condition
	Employment
	Assegnatari
	Conducenti
	Allegato
	ContrAcquisto
	ContrLeasing
	ContrNoleggio
	Fornitore
	Incidente
	Responsabilita
	ControparteIncidente
	Movimenti
	Multe 
	TipoInfrazione 
	Rifornimento
	Spesa
	TipoSpesa
    
//VEICOLO
//Veicolo
type VehicleDG struct {
	Id     int `orm:"auto;unique"`
	Targa          string `orm:"unique"`
	DataInFlotta   time.Time    `orm:"type(date)"`
	DataFineFlotta time.Time    `orm:"type(date)"`
	Note           string       `orm:"null;size(100)"`
	VehicleDT      *VehicleDT   `orm:"rel(one)"`
	VehicleType    *VehicleType `orm:"rel(fk)"`
	Sector         *Sector      `orm:"rel(fk)"`
	Condition      *Condition   `orm:"rel(fk)"`
	Employment     *Employment  `orm:"rel(fk)"`
	Conducenti     *Conducenti  `orm:"rel(fk)"`
	Movimenti		[]*Movimenti `orm:"reverse(many)"`
	Spesa			[]*Spesa `orm:"reverse(many)"`
	Incidenti		[]*Incidenti `orm:"reverse(many)"`
	Rifornimenti	[]*Rifornimenti `orm:"reverse(many)"`
	
}

//Specifiche tecniche veicolo
type VehicleDT struct {
	Id         int              `orm:"pk;not null;auto;unique"`
	MatriculationYear time.Time        `orm:"type(date)"`
	NLibretto         int              `orm:"null"`
	NTelaio           int              `orm:"null"`
	Marca             string           `orm:"size(7)"`
	Modello           string           `orm:"size(40)"`
	NorEuro           int              `orm:"null"`
	Kw                int              `orm:"null"`
	Cilindrata        int              `orm:"null"`
	ConsumoTeorico    int              `orm:"null"`
	KmAnno            int              `orm:"null"`
	CostoKm           int              `orm:"null;digits(12);decimals(4)"`
	Pneumatici        string           `orm:"null;size(20)"`
	VehicleDG         *VehicleDG       `orm:"reverse(one)"`
	Allegati          []*Allegati      `orm:"rel(m2m)"`
	Carburante        *Carburante      `orm:"rel(fk)"`
	ContrAcquisto     *ContrAcquisto   `orm:"rel(fk)"`
	ContrLeasing      []*ContrLeasing  `orm:"rel(m2m)"`
	ContrNoleggio     []*ContrNoleggio `orm:"rel(m2m)"`
}






type Conducenti struct {
	Id  int             `orm:"pk;not null;auto;unique"`
	Nome          string          `orm:"size(20)"`
	Cognome       string          `orm:"size(20)"`
	CodiceFiscale string          `orm:"null;size(16)"`
	VehicleDG     []*VehicleDG    `orm:"reverse(many)"`
	Incidenti     []*Incidenti    `orm:"reverse(many)"`
	Movimenti     []*Movimenti    `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti `orm:"reverse(many)"`
	Spesa         []*Spesa        `orm:"reverse(many)"`
}

//File ALlegati
type Allegati struct {
	Id    int              `orm:"pk;not null;auto;unique"`
	Percorso      string           `orm:"size(100);not null;"`
	Descrizione   string           `orm:"size(100);not null;"`
	ContrAcquisto []*ContrAcquisto `orm:"reverse(many)"`
	ContrLeasing  []*ContrLeasing  `orm:"reverse(many)"`
	ContrNoleggio []*ContrNoleggio `orm:"reverse(many)"`
	Incidenti     []*Incidenti     `orm:"reverse(many)"`
	Movimenti     []*Movimenti     `orm:"reverse(many)"`
	Multe         []*Multe         `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti  `orm:"reverse(many)"`
	Spesa         []*Spesa         `orm:"reverse(many)"`
	VehicleDT     []*VehicleDT     `orm:"reverse(many)"`
}

//Contratto di acquisto
type ContrAcquisto struct {
	Id   int          `orm:"pk;not null;auto;unique"`
	NContratto        string       `orm:"unique;not null;size(20)"`
	DataAcq           time.Time    `orm:"null;type(date)"`
	Importo           float64      `orm:"null;digits(12);decimals(4)"`
	AmmortamentoAnnuo int          `orm:"null"`
	FineGaranzia      time.Time    `orm:"null;auto_now_add;type(date)"`
	KmAcquisto        int          `orm:"null"`
	KmInizioGest      int          `orm:"null"`
	Note              string       `orm:"null;size(100)"`
	Allegati          []*Allegati  `orm:"rel(m2m)"`
	Fornitori         *Fornitori   `orm:"rel(fk)"`
	VehicleDT         []*VehicleDT `orm:"reverse(many)"`
}

//Contratto di leasing
type ContrLeasing struct {
	Id int          `orm:"pk;not null;auto;unique"`
	NContratto      string       `orm:"unique;not null;size(20)"`
	DataCont        time.Time    `orm:"auto_now_add;type(date)"`
	PrimaRata       float64      `orm:"null;digits(12);decimals(4)"`
	RataSucc        float64      `orm:"null;digits(12);decimals(4)"`
	NRate           int          `orm:"null"`
	Riscatto        float64      `orm:"null;digits(12);decimals(4)"`
	DataRiscatto    time.Time    `orm:"null;type(date)"`
	ImportoTot      float64      `orm:"null;digits(12);decimals(4)"`
	FineCont        time.Time    `orm:"null;type(date)"`
	FineGaranzia    time.Time    `orm:"null;type(date)"`
	KmInizioGest    int          `orm:"null"`
	KmFineGest      int          `orm:"null"`
	Note            string       `orm:"null;size(100)"`
	Allegati        []*Allegati  `orm:"rel(m2m)"`
	Fornitori       *Fornitori   `orm:"rel(fk)"`
	VehicleDT       []*VehicleDT `orm:"reverse(many)"`
}

//Contratto di noleggio
type ContrNoleggio struct {
	Id      int          `orm:"pk;not null;auto;unique"`
	NContratto           string       `orm:"unique;not null;size(20)"`
	DataCont             time.Time    `orm:"null;type(date)"`
	DataInizio           time.Time    `orm:"null;type(date)"`
	DataFine             time.Time    `orm:"null;type(date)"`
	Riparamentrizzazione int          `orm:"null"`
	NRate                int          `orm:"null"`
	CanoneBase           float64      `orm:"null;digits(12);decimals(4)"`
	CanoneServizi        float64      `orm:"null;digits(12);decimals(4)"`
	CanoneAltro          float64      `orm:"null;digits(12);decimals(4)"`
	CanoneTot            float64      `orm:"null;digits(12);decimals(4)"`
	KmContrattuali       int          `orm:"null"`
	AddebitoKmExtra      int          `orm:"null"`
	ImportoKm            float64      `orm:"null;digits(12);decimals(4)"`
	ImportoTot           float64      `orm:"null;digits(12);decimals(4)"`
	KmInizioGest         int          `orm:"null"`
	KmFineGest           int          `orm:"null"`
	Note                 string       `orm:"null;size(100)"`
	Allegati             []*Allegati  `orm:"rel(m2m)"`
	Fornitori            *Fornitori   `orm:"rel(fk)"`
	VehicleDT            []*VehicleDT `orm:"reverse(many)"`
}

//Fornitori
type Fornitori struct {
	Id   int              `orm:"pk;not null;auto;unique"`
	Descrizione   string           `orm:"size(100)"`
	PI            string           `orm:"null"`
	ContrAcquisto []*ContrAcquisto `orm:"reverse(many)"`
	ContrLeasing  []*ContrLeasing  `orm:"reverse(many)"`
	ContrNoleggio []*ContrNoleggio `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti  `orm:"reverse(many)"`
}

//Incidenti
type Incidenti struct {
	Id						int						`orm:"pk;not null;auto;unique"`
	Data					time.Time				`orm:"type(datetime)"`
	Assicurazione			string					`orm:"size(100)"`
	ImportoDanno			float64					`orm:"null;digits(12);decimals(4)"`
	FranchigiaPagata		float64					`orm:"null;digits(12);decimals(4)"`
	ImportoLiquidato		float64					`orm:"null;digits(12);decimals(4)"`
	DataChiusura			time.Time				`orm:"null;type(datetime)"`
	Feriti					bool					`orm:"null"`
	AddebitoConducente		bool					`orm:"null"`
	Note					string					`orm:"null;size(100)"`
	Descrizione				string					`orm:"null;size(100)"`
	ControparteIncidenti	*ControparteIncidenti	`orm:"rel(one)"`
	Conducenti				*Conducenti				`orm:"rel(fk)"`
	Allegati				[]*Allegati				`orm:"rel(m2m)"`
	VehicleDG				[]*VehicleDG			`orm:"rel(m2m)"`
	Responsabilita			*Responsabilita			`orm:"rel(fk)"`
}


type ControparteIncidenti struct {
	Id int        `orm:"pk;not null;auto;unique"`
	Assicurazione string     `orm:"null;size(100)"`
	Targa         string     `orm:"null;size(7)"`
	Marca         string     `orm:"null;size(30)"`
	Modello       string     `orm:"null;size(30)"`
	Proprietario  string     `orm:"null;size(100)"`
	Conducente    string     `orm:"null;size(100)"`
	Riferimento   string     `orm:"null;size(100)"`
	Incidenti     *Incidenti `orm:"reverse(one)"`
}

//MOVIMENTI

type Movimenti struct {
	Id  int			`orm:"pk;not null;auto;unique"`
	DataInizio   time.Time		`orm:"type(datetime)"`
	KmInizio     int			`orm:"unique;not null"`
	Destinazione string			`orm:"not null;size(100)"`
	DataFine     time.Time		`orm:"type(datetime)"`
	KmFine       int			`orm:"not null"`
	Note         string			`orm:"null;size(100)"`
	Conducenti   *Conducenti	`orm:"rel(fk)"`
	Allegati     []*Allegati	`orm:"rel(m2m)"`
	VehicleDG    []*VehicleDG   `orm:"rel(m2m)"`
}

//MULTE
type Multe struct {
	Id            int       `orm:"pk;not null;auto;unique"`
	Data               time.Time `orm:"type(datetime)"`
	Importo            float64   `orm:"digits(12);decimals(4)"`
	AddebitoConducente bool      `orm:"null"`
	AutoritaSanzione   string    `orm:"size(100)"`
	NVerbale           int
	DataNotifica       time.Time         `orm:"type(datetime)"`
	ScadenzaPagamento  time.Time         `orm:"null;type(datetime)"`
	DataPagamento      time.Time         `orm:"null;type(datetime)"`
	Note               string            `orm:"null;size(100)"`
	Conducenti         *Conducenti       `orm:"rel(fk)"`
	Allegati           []*Allegati       `orm:"rel(m2m)"`
	TipoInfrazione     []*TipoInfrazione `orm:"rel(m2m)"`
	VehicleDG         	[]*VehicleDG       		`orm:"rel(m2m)"`
}


//RIFORNIMENTO

type Rifornimenti struct {
	Id int       `orm:"pk;not null;auto;unique"`
	Data           time.Time `orm:"type(datetime)"`
	Km             int
	Importo        float64 `orm:"digits(12);decimals(4)"`
	CostoLitro     float64 `orm:"digits(12);decimals(4)"`
	Litri          int
	Note           string      `orm:"null;size(100)"`
	Fornitori      *Fornitori  `orm:"rel(fk)"`
	Conducenti     *Conducenti `orm:"rel(fk)"`
	Allegati       []*Allegati `orm:"rel(m2m)"`
	VehicleDG         	[]*VehicleDG       		`orm:"rel(m2m)"`
}

//SPESA



type Spesa struct {
	Id      		 int       		`orm:"pk;not null;auto;unique"`
	Data             time.Time 		`orm:"auto_now_add;type(datetime)"`
	Km               int
	Importo          float64     	`orm:"digits(12);decimals(4)"`
	Descrizione      string      	`orm:"size(100)"`
	NDocumento       string      	`orm:"null;size(20)"`
	DataDocu         time.Time   	`orm:"null;type(datetime)"`
	DataProsScadenza time.Time   	`orm:"null;type(datetime)"`
	KmProxScadenza   int         	`orm:"null"`
	Note             string      	`orm:"null;size(100)"`
	TipoSpesa        *TipoSpesa  	`orm:"rel(fk)"`
	Fornitori        *Fornitori  	`orm:"rel(fk)"`
	Conducenti       *Conducenti 	`orm:"rel(fk)"`
	Allegati         []*Allegati 	`orm:"rel(m2m)"`
	VehicleDG        []*VehicleDG  	`orm:"rel(m2m)"`

}    
    
    //vehicle := models.VehicleDG.
    
	user := models.AuthUser{Id:parms}

	
	err = o.Read(&user, "Id")

	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.First)
	}
			

	err = o.QueryTable("auth_user").Filter("Id", parms).RelatedSel().One(&user)
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user)
		fmt.Println(user.AuthApp.Automezzi)
	}

*/	
	
	


}