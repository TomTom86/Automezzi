package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//AuthUser struct for login manager
type AuthUser struct {
	ID            int    `orm:"auto"`
	First         string `orm:"size(20)"`
	Last          string `orm:"size(20)"`
	Email         string `orm:"unique"`
	Password      string `orm:"size(60)"`
	IsApproved    bool
	IDkey         string    `orm:"size(20)"`
	RegDate       time.Time `orm:"auto_now_add;type(datetime)"`
	LastLoginDate time.Time `orm:"auto_now_add;type(datetime)"`
	LastEditDate  time.Time `orm:"auto_now_add;type(datetime)"`
	ResetKey      string    `orm:"size(20)"`
	BlockControll int
	Group         int
	Note          string   `orm:"size(100)"`
	AuthApp       *AuthApp `orm:"rel(one)"`
}

//AuthApp struct for app manager
type AuthApp struct {
	ID        int
	Automezzi bool
	Servizi   bool
	AuthUser  *AuthUser `orm:"reverse(one)"`
}

//***********DB AUTOMEZZI**************

//VeicoliDG contiene dati veicolo
type VeicoliDG struct {
	ID             int             `orm:"auto;unique"`
	Targa          string          `orm:"unique"`
	DataInFlotta   time.Time       `orm:"type(date)"`
	DataFineFlotta time.Time       `orm:"type(date)"`
	Note           string          `orm:"null;size(100)"`
	VeicoliDT      *VeicoliDT      `orm:"rel(one)"`
	TipiVeicolo    *TipiVeicolo    `orm:"rel(fk)"`
	Settori        *Settori        `orm:"rel(fk)"`
	Condizioni     *Condizioni     `orm:"rel(fk)"`
	Impieghi       *Impieghi       `orm:"rel(fk)"`
	Conducenti     *Conducenti     `orm:"rel(fk)"`
	Movimenti      []*Movimenti    `orm:"reverse(many)"`
	Spese          []*Spese        `orm:"reverse(many)"`
	Incidenti      []*Incidenti    `orm:"reverse(many)"`
	Rifornimenti   []*Rifornimenti `orm:"reverse(many)"`
}

//VeicoliDT contiene le specifiche tecniche veicolo
type VeicoliDT struct {
	ID                int             `orm:"pk;not null;auto;unique"`
	MatriculationYear time.Time       `orm:"type(date)"`
	NLibretto         int             `orm:"null"`
	NTelaio           int             `orm:"null"`
	Marca             string          `orm:"size(7)"`
	Modello           string          `orm:"size(40)"`
	NorEuro           int             `orm:"null"`
	Kw                int             `orm:"null"`
	Cilindrata        int             `orm:"null"`
	ConsumoTeorico    int             `orm:"null"`
	KmAnno            int             `orm:"null"`
	CostoKm           int             `orm:"null;digits(12);decimals(4)"`
	Pneumatici        string          `orm:"null;size(20)"`
	VeicoliDG         *VeicoliDG      `orm:"reverse(one)"`
	Allegati          []*Allegati     `orm:"rel(m2m)"`
	Carburante        *Carburante     `orm:"rel(fk)"`
	ContrAcquisti     *ContrAcquisti  `orm:"rel(fk)"`
	ContrLeasing      []*ContrLeasing `orm:"rel(m2m)"`
	ContrNoleggi      []*ContrNoleggi `orm:"rel(m2m)"`
}

//Carburante contiene i tipi di carburante
type Carburante struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(30)"`
	VeicoliDT   []*VeicoliDT `orm:"reverse(many)"`
}

//TipiVeicolo contiene i tipi di veicoli
type TipiVeicolo struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(100)"`
	VeicoliDG   []*VeicoliDG `orm:"reverse(many)"`
}

//Settori contiene i Settori di assegnazione
type Settori struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(100)"`
	VeicoliDG   []*VeicoliDG `orm:"reverse(many)"`
}

//Condizioni contiene i tipi di Condizioni veicoli
type Condizioni struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(100)"`
	VeicoliDG   []*VeicoliDG `orm:"reverse(many)"`
}

//Impieghi contiene i tipi di Impieghi
type Impieghi struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(100)"`
	VeicoliDG   []*VeicoliDG `orm:"reverse(many)"`
}

//Conducenti contiene l'elenco dei conducenti
type Conducenti struct {
	ID            int             `orm:"pk;not null;auto;unique"`
	Nome          string          `orm:"size(20)"`
	Cognome       string          `orm:"size(20)"`
	CodiceFiscale string          `orm:"null;size(16)"`
	VeicoliDG     []*VeicoliDG    `orm:"reverse(many)"`
	Incidenti     []*Incidenti    `orm:"reverse(many)"`
	Movimenti     []*Movimenti    `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti `orm:"reverse(many)"`
	Spese         []*Spese        `orm:"reverse(many)"`
}

//Allegati continee l'elenco degli allegati
type Allegati struct {
	ID            int              `orm:"pk;not null;auto;unique"`
	Percorso      string           `orm:"size(100);not null;"`
	Descrizione   string           `orm:"size(100);not null;"`
	ContrAcquisti []*ContrAcquisti `orm:"reverse(many)"`
	ContrLeasing  []*ContrLeasing  `orm:"reverse(many)"`
	ContrNoleggi  []*ContrNoleggi  `orm:"reverse(many)"`
	Incidenti     []*Incidenti     `orm:"reverse(many)"`
	Movimenti     []*Movimenti     `orm:"reverse(many)"`
	Multe         []*Multe         `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti  `orm:"reverse(many)"`
	Spese         []*Spese         `orm:"reverse(many)"`
	VeicoliDT     []*VeicoliDT     `orm:"reverse(many)"`
}

//ContrAcquisti contiene i contratti di acquisto
type ContrAcquisti struct {
	ID                int          `orm:"pk;not null;auto;unique"`
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
	VeicoliDT         []*VeicoliDT `orm:"reverse(many)"`
}

//ContrLeasing contiene i Contratti di leasing
type ContrLeasing struct {
	ID           int          `orm:"pk;not null;auto;unique"`
	NContratto   string       `orm:"unique;not null;size(20)"`
	DataCont     time.Time    `orm:"auto_now_add;type(date)"`
	PrimaRata    float64      `orm:"null;digits(12);decimals(4)"`
	RataSucc     float64      `orm:"null;digits(12);decimals(4)"`
	NRate        int          `orm:"null"`
	Riscatto     float64      `orm:"null;digits(12);decimals(4)"`
	DataRiscatto time.Time    `orm:"null;type(date)"`
	ImportoTot   float64      `orm:"null;digits(12);decimals(4)"`
	FineCont     time.Time    `orm:"null;type(date)"`
	FineGaranzia time.Time    `orm:"null;type(date)"`
	KmInizioGest int          `orm:"null"`
	KmFineGest   int          `orm:"null"`
	Note         string       `orm:"null;size(100)"`
	Allegati     []*Allegati  `orm:"rel(m2m)"`
	Fornitori    *Fornitori   `orm:"rel(fk)"`
	VeicoliDT    []*VeicoliDT `orm:"reverse(many)"`
}

//ContrNoleggi contiene i Contratti di noleggio
type ContrNoleggi struct {
	ID                   int          `orm:"pk;not null;auto;unique"`
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
	VeicoliDT            []*VeicoliDT `orm:"reverse(many)"`
}

//Fornitori contiene l'elenco fornitori
type Fornitori struct {
	ID            int              `orm:"pk;not null;auto;unique"`
	Descrizione   string           `orm:"size(100)"`
	PI            string           `orm:"null"`
	ContrAcquisti []*ContrAcquisti `orm:"reverse(many)"`
	ContrLeasing  []*ContrLeasing  `orm:"reverse(many)"`
	ContrNoleggi  []*ContrNoleggi  `orm:"reverse(many)"`
	Rifornimenti  []*Rifornimenti  `orm:"reverse(many)"`
}

//Incidenti contiene l'elenco degli incidenti
type Incidenti struct {
	ID                   int                   `orm:"pk;not null;auto;unique"`
	Data                 time.Time             `orm:"type(datetime)"`
	Assicurazione        string                `orm:"size(100)"`
	ImportoDanno         float64               `orm:"null;digits(12);decimals(4)"`
	FranchigiaPagata     float64               `orm:"null;digits(12);decimals(4)"`
	ImportoLiquidato     float64               `orm:"null;digits(12);decimals(4)"`
	DataChiusura         time.Time             `orm:"null;type(datetime)"`
	Feriti               bool                  `orm:"null"`
	AddebitoConducente   bool                  `orm:"null"`
	Note                 string                `orm:"null;size(100)"`
	Descrizione          string                `orm:"null;size(100)"`
	ContropartiIncidenti *ContropartiIncidenti `orm:"rel(one)"`
	Conducenti           *Conducenti           `orm:"rel(fk)"`
	Allegati             []*Allegati           `orm:"rel(m2m)"`
	VeicoliDG            []*VeicoliDG          `orm:"rel(m2m)"`
	Responsabilita       *Responsabilita       `orm:"rel(fk)"`
}

//Responsabilita Incidenti
type Responsabilita struct {
	ID          int          `orm:"pk;not null;auto;unique"`
	Descrizione string       `orm:"size(100)"`
	Incidenti   []*Incidenti `orm:"reverse(many)"`
}

//ContropartiIncidenti contiene i dati delle controparti incidenti
type ContropartiIncidenti struct {
	ID            int        `orm:"pk;not null;auto;unique"`
	Assicurazione string     `orm:"null;size(100)"`
	Targa         string     `orm:"null;size(7)"`
	Marca         string     `orm:"null;size(30)"`
	Modello       string     `orm:"null;size(30)"`
	Proprietario  string     `orm:"null;size(100)"`
	Conducente    string     `orm:"null;size(100)"`
	Riferimento   string     `orm:"null;size(100)"`
	Incidenti     *Incidenti `orm:"reverse(one)"`
}

//Movimenti contiene i movimenti delle auto
type Movimenti struct {
	ID           int          `orm:"pk;not null;auto;unique"`
	DataInizio   time.Time    `orm:"type(datetime)"`
	KmInizio     int          `orm:"unique;not null"`
	Destinazione string       `orm:"not null;size(100)"`
	DataFine     time.Time    `orm:"type(datetime)"`
	KmFine       int          `orm:"not null"`
	Note         string       `orm:"null;size(100)"`
	Conducenti   *Conducenti  `orm:"rel(fk)"`
	Allegati     []*Allegati  `orm:"rel(m2m)"`
	VeicoliDG    []*VeicoliDG `orm:"rel(m2m)"`
}

//Multe contiene i dati delle auto
type Multe struct {
	ID                 int       `orm:"pk;not null;auto;unique"`
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
	TipiInfrazione     []*TipiInfrazione `orm:"rel(m2m)"`
	VeicoliDG          []*VeicoliDG      `orm:"rel(m2m)"`
}

//TipiInfrazione contiene i tipi di infrazioni
type TipiInfrazione struct {
	ID          int      `orm:"pk;not null;auto;unique"`
	Descrizione string   `orm:"size(100)"`
	Multe       []*Multe `orm:"reverse(many)"`
}

//Rifornimenti contiene i dati relativi ai rifornimenti
type Rifornimenti struct {
	ID         int       `orm:"pk;not null;auto;unique"`
	Data       time.Time `orm:"type(datetime)"`
	Km         int
	Importo    float64 `orm:"digits(12);decimals(4)"`
	CostoLitro float64 `orm:"digits(12);decimals(4)"`
	Litri      int
	Note       string       `orm:"null;size(100)"`
	Fornitori  *Fornitori   `orm:"rel(fk)"`
	Conducenti *Conducenti  `orm:"rel(fk)"`
	Allegati   []*Allegati  `orm:"rel(m2m)"`
	VeicoliDG  []*VeicoliDG `orm:"rel(m2m)"`
}

//TipiSpesa contiene i tipi di Spese
type TipiSpesa struct {
	ID          int      `orm:"pk;not null;auto;unique"`
	Descrizione string   `orm:"size(100)"`
	Spese       []*Spese `orm:"reverse(many)"`
}

//Spese contiene i dati relativi alle Spese
type Spese struct {
	ID               int       `orm:"pk;not null;auto;unique"`
	Data             time.Time `orm:"auto_now_add;type(datetime)"`
	Km               int
	Importo          float64      `orm:"digits(12);decimals(4)"`
	Descrizione      string       `orm:"size(100)"`
	NDoc             string       `orm:"null;size(20)"`
	DataDoc          time.Time    `orm:"null;type(datetime)"`
	DataProsScadenza time.Time    `orm:"null;type(datetime)"`
	KmProsScadenza   int          `orm:"null"`
	Note             string       `orm:"null;size(100)"`
	TipiSpesa        *TipiSpesa   `orm:"rel(fk)"`
	Fornitori        *Fornitori   `orm:"rel(fk)"`
	Conducenti       *Conducenti  `orm:"rel(fk)"`
	Allegati         []*Allegati  `orm:"rel(m2m)"`
	VeicoliDG        []*VeicoliDG `orm:"rel(m2m)"`
}

//*************FINE DB AUTOMEZZI****************
func init() {
	//Login & App Manager DB
	orm.RegisterModel(new(AuthUser), new(AuthApp))
	//Automezzi DB
	orm.RegisterModel(new(VeicoliDG), new(VeicoliDT), new(Carburante), new(TipiVeicolo), new(Settori), new(Condizioni), new(Impieghi), new(Conducenti), new(Allegati), new(ContrAcquisti), new(ContrLeasing), new(ContrNoleggi), new(Fornitori), new(Incidenti), new(Responsabilita))
	orm.RegisterModel(new(ContropartiIncidenti), new(Movimenti), new(Multe), new(TipiInfrazione), new(Rifornimenti), new(TipiSpesa), new(Spese))
}
