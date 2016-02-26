package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AuthUser struct {
	Id              int    `orm:"auto"`
	First           string `orm:"size(20)"`
	Last            string `orm:"size(20)"`
	Email           string `orm:"unique"`
	Password        string `orm:"size(60)"`
	Is_approved     bool
	Id_key          string    `orm:"size(20)"`
	Reg_date        time.Time `orm:"auto_now_add;type(datetime)"`
	Last_login_date time.Time `orm:"auto_now_add;type(datetime)"`
	Last_edit_date  time.Time `orm:"auto_now_add;type(datetime)"`
	Reset_key       string    `orm:"size(20)"`
	Block_controll  int
	Group           int
	Note            string   `orm:"size(100)"`
	AuthApp         *AuthApp `orm:"rel(one)"`
}

type AuthApp struct {
	Id        int 
	Automezzi bool
	Servizi   bool
	AuthUser  *AuthUser `orm:"reverse(one)"`
}

//***********DB AUTOMEZZI**************

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

type Carburante struct {
	Id int          `orm:"pk;not null;auto;unique"`
	Descrizione  string       `orm:"size(30)"`
	VehicleDT    []*VehicleDT `orm:"reverse(many)"`
}

//Tipo di veicolo
/*
Autoveicolo
Camion
Ciclomotore
Furgone
Pullman
Motoveicolo
Altro
*/
type VehicleType struct {
	Id   int          `orm:"pk;not null;auto;unique"`
	Description string       `orm:"size(100)"`
	VehicleDG   []*VehicleDG `orm:"reverse(many)"`
}

//Tipo di settore aziendale
/*
Food
Lavanderia
Pulizia
*/
type Sector struct {
	Id   int          `orm:"pk;not null;auto;unique"`
	Description string       `orm:"size(100)"`
	VehicleDG   []*VehicleDG `orm:"reverse(many)"`
}

//Condizione veicolo
/*
Buono stato
Cattivo Stato
Discreto Stato
In Attesa di Alienazione
In Attesa di Assegnazione
In attesa di Riparazione
Non utilizzabile
Rubato
Alienato
*/
type Condition struct {
	Id int          `orm:"pk;not null;auto;unique"`
	Description string       `orm:"size(100)"`
	VehicleDG   []*VehicleDG `orm:"reverse(many)"`
}

//Impiego veicolo
/*
Aziendale
Aziendale + Personale
Personale
*/
type Employment struct {
	Id int          `orm:"pk;not null;auto;unique"`
	Description  string       `orm:"size(100)"`
	VehicleDG    []*VehicleDG `orm:"reverse(many)"`
}

/*
type Assegnatari struct{
	IdAssegnatario int `orm:"auto"`
	Nome string
	Cognome string
	CodiceFiscale string
	VehicleDG []*VehicleDG`orm:"reverse(many)"`
}*/

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

//Responsabilit Incidenti
/*
Concorso di colpa
Da accertare
Della controparte
Propria
*/
type Responsabilita struct {
	Id int        `orm:"pk;not null;auto;unique"`
	Descrizione      string     `orm:"size(100)"`
	Incidenti        []*Incidenti `orm:"reverse(many)"`
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

/*
Accesso in senso vietato
Cinture di sicurezza non allacciate
Eccesso di velocità
Guida contromano
Guida pericolosa
Positivo a controlllo alcool
Precedenza non rispettata
Semaforo rosso
Sosta vietata
Utilizzo di telefono cellulare
Violazione di corsia preferenziale
Violazione di ztl
Altro
*/
type TipoInfrazione struct {
	Id int      `orm:"pk;not null;auto;unique"`
	Descrizione      string   `orm:"size(100)"`
	Multe            []*Multe `orm:"reverse(many)"`
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

/*
Alienazione
Assicurazione
Bollo
Contratto Canone
Contratto varie
Lavaggio
Manutenzione ordinaria
Pneumatici
Revisione
Riparazione per sinistro
Riparazione straordinaria
Varie
*/
type TipoSpesa struct {
	Id			int		`orm:"pk;not null;auto;unique"`
	Descrizione string	`orm:"size(100)"`
	Spesa       []*Spesa	`orm:"reverse(many)"`
}

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




//*************FINE DB AUTOMEZZI****************
func init() {
	//Login & App Manager DB
	orm.RegisterModel(new(AuthUser), new(AuthApp))
	//Automezzi DB
	orm.RegisterModel(new(VehicleDG), new(VehicleDT), new(Carburante), new(VehicleType), new(Sector), new(Condition), new(Employment), new(Conducenti), new(Allegati), new(ContrAcquisto), new(ContrLeasing), new(ContrNoleggio), new(Fornitori), new(Incidenti), new(Responsabilita))
	orm.RegisterModel(new(ControparteIncidenti), new(Movimenti), new(Multe), new(TipoInfrazione), new(Rifornimenti), new(TipoSpesa), new(Spesa) )
}
