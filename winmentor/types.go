package winmentor

// Partner represents a business partner (customer/supplier) from WinMENTOR.
type Partner struct {
	ID                string
	Denumire          string
	CodFiscal         string
	Localitate        string
	Adresa            string
	Telefon           string
	PersContact       string
	SimbolClasa       string
	DenClasa          string
	SimbolCatPret     string
	DenCatPret        string
	MarcaAgent        string
	NumeAgent         string
	PrenumeAgent      string
	Scadenta          string
	Discount          string
	DenCritDiscount   string
	SediiPartener     string // "~" separated if multiple
	CodExtern         string
	PartnerBlocat     string // DA/NU
	CreditVanzare     string
	CodFiscal2        string
	ContBanca         string
	LocalitatiSedii   string // "~" separated
	Tara              string
	MarcaAgentiSedii  string // "~" separated
	Observatii        string
	FlagSediuSocial   string // "~" separated D flags
	CodPostalSedii    string // "~" separated
	EmailSedii        string // "~" separated
	TelPersContact    string
	PFsauPJ           string // PF or PJ
	MonedaImplicita   string
	DataAdaugarii     string
	Trasee            string
	PuncteAcumulate   string
	CodFiscalSedii    string // "~" separated
	InfoTipSediu      string // "~" separated
}

// PartnerInput represents the data structure for adding/modifying a partner.
type PartnerInput struct {
	ID                    string
	Denumire              string
	CodFiscal             string
	SediulInLocalitatea   string
	AdresaSediu           string
	TelefonSediu          string
	PersoaneContact       string // "~" separated
	SimbolClasa           string
	SimbolCategoriePret   string
	IDAgentImplicit       string
	NrRegistrulComertului string
	Observatii            string
	SimbolBanca           string // "~" separated
	NumeBanca             string // "~" separated
	LocalitateBanca       string // "~" separated
	ContBanca             string // "~" separated
	ZiImplicitaPlata      string
	NumeSediuSecundar     string // "~" separated
	AdresaSediuSecundar   string // "~" separated
	TelefonSediuSecundar  string // "~" separated
	LocalitateSediuSec    string // "~" separated
	IDAgentSediuSec       string // "~" separated
	CodExtern             string
	SimbolAutoJudetLivr   string
	SimbolAutoJudetSediu  string
	FlagPF                string
	ScadentaImplicita     string
	SimbolTipContabil     string
	FlagProducator        string // P
	EmailSediuSocial      string
	EmailSediiLivrare     string // "~" separated
	TVAIncasare           string // D
	SerAI                 string
	NrAI                  string
	SimbolAutoTaraSediu   string
}

// ToRecord serializes a PartnerInput to a semicolon-separated string.
func (p *PartnerInput) ToRecord() string {
	fields := []string{
		p.ID, p.Denumire, p.CodFiscal, p.SediulInLocalitatea,
		p.AdresaSediu, p.TelefonSediu, p.PersoaneContact,
		p.SimbolClasa, p.SimbolCategoriePret, p.IDAgentImplicit,
		p.NrRegistrulComertului, p.Observatii,
		p.SimbolBanca, p.NumeBanca, p.LocalitateBanca, p.ContBanca,
		p.ZiImplicitaPlata,
		p.NumeSediuSecundar, p.AdresaSediuSecundar, p.TelefonSediuSecundar,
		p.LocalitateSediuSec, p.IDAgentSediuSec,
		p.CodExtern, p.SimbolAutoJudetLivr, p.SimbolAutoJudetSediu,
		p.FlagPF, p.ScadentaImplicita, p.SimbolTipContabil,
		p.FlagProducator, p.EmailSediuSocial, p.EmailSediiLivrare,
		p.TVAIncasare, p.SerAI, p.NrAI, p.SimbolAutoTaraSediu,
	}
	result := ""
	for i, f := range fields {
		if i > 0 {
			result += ";"
		}
		result += f
	}
	return result
}

// StockArticle represents an article in stock.
// PDF fields: CodExtern;Denumire;UM;PretVanzare;Stoc;SimbolClasa;DenClasa;IDProducator;
// Den.Producator;IDFurnizor;DenFurnizor;SimbolGestiune;DenGestiune;CotaTVA;Flag;PretCuTVA;StocRezervat
type StockArticle struct {
	CodExtern       string
	Denumire        string
	UM              string
	PretVanzare     string
	Stoc            string
	SimbolClasa     string
	DenClasa        string
	IDProducator    string
	DenProducator   string
	IDFurnizor      string
	DenFurnizor     string
	SimbolGestiune  string
	DenGestiune     string
	CotaTVA         string
	FlagTVAInclus   string // D/N - whether VAT is included in sale price
	PretCuTVA       string
	StocRezervat    string
}

// ArticleStock represents the stock info for a single article.
type ArticleStock struct {
	CodExtern   string
	Denumire    string
	UM          string
	PretVanzare string
	Stoc        string
}

// SoldPartener represents a partner's balance.
type SoldPartener struct {
	CodExtern string
	Denumire  string
	Sold      string
}

// Employee represents a WinMENTOR employee record.
type Employee struct {
	Nume       string
	Prenume    string
	Marca      string
	CNP        string
	EsteActiv  string // Da/Nu
	EsteAgent  string // Da/Nu
	SerieBuletin string
	NumarBuletin string
	CodPostal  string
}

// DetailedBalance represents a line in the detailed balance (invoice or advance).
type DetailedBalance struct {
	Type         string // "Factura" or "Avans"
	NrDocument   string
	DataDocument string
	Rest         string
}

// Gestiune represents a warehouse/store.
type Gestiune struct {
	Simbol   string
	Denumire string
}

// ClasaParteneri represents a partner class.
type ClasaParteneri struct {
	Simbol   string
	Denumire string
}

// Product represents a product returned by GetProducts.
type Product struct {
	IDArticol             string
	Denumire              string
	DenUM                 string
	IDProducator          string
	DenumireProducator    string
	TipSerie              string
	DataAdaugarii         string
	DataUltimeiModificari string
	TipUM                 string
	CodInternWinMentor    string
	SimbolClasa           string
}

// DeletedProduct represents a deleted product entry.
type DeletedProduct struct {
	CodInternWinMentor string
	DataOraStergerii   string
}

// Bank represents a bank entry.
type Bank struct {
	Simbol   string
	Denumire string
}

// Oferta represents a price offer.
type Oferta struct {
	PartID           string
	ArtID            string
	DataInceput      string
	DataSfarsit      string
	Pret             string
	Cantitate        string
}

// ClasaArticole represents an article class.
type ClasaArticole struct {
	Simbol   string
	Denumire string
}

// NomenclatorArticol represents an article in the full nomenclature.
type NomenclatorArticol struct {
	CodExtern             string
	Denumire              string
	DenUM                 string
	PretVanzare           string
	SimbolClasa           string
	DenClasa              string
	CodExternProducator   string
	DenProducator         string
	GestImplicita         string
	CodExternUnic         string
	CotaTVA               string
	DenUMSecundara        string
	ParitateUMSecundara   string
	Masa                  string
	Serviciu              string
	CodVamal              string
	PretMinim             string
	CantImplicita         string
	PretValuta            string
	DataAdaug             string
	Masa2                 string
	PretVCuTVA            string
	Locatie               string
	PretReferinta         string
}

// VanzareExt represents an extended sale record.
type VanzareExt struct {
	PartID        string
	Zi            string
	PrefixDoc     string
	NrDoc         string
	ArtID         string
	Cant          string
	DenUM         string
	Pret          string
	DenGest       string
	CodInternArt  string
	LocatieClient string
	Adresa        string
	Comision      string
	CodFisca      string
	MarcaAgent    string
	ValAchizitie  string
	CodPostal     string
	ClasaArticol  string
}

// VanzareLuna represents a monthly sale line.
type VanzareLuna struct {
	IDPartener    string
	Zi            string
	NrFactura     string
	IDArticol     string
	NumarComanda  string
	Cant          string
	DenUM         string
	Pret          string
	MarcaAgent    string
	ValoareFactura string
}

// SoldExt represents an extended balance line.
type SoldExt struct {
	IDPartener       string
	Tip              string // "Factura" or "Avans" etc.
	NrFactura        string
	DataFactura      string
	RestDePlata      string
	TermenDePlata    string
	LocatiePartener  string
	MarcaAgent       string
	ValoareFactura   string
	ObservatiiFactura string
}

// ComandaNefacturata represents an uninvoiced order line.
type ComandaNefacturata struct {
	IDArticol    string
	NumarComanda string
	Cant         string
	DenArticol   string
}

// CategoriePret represents a price category.
type CategoriePret struct {
	Simbol   string
	Denumire string
}

// Sold represents an agent balance record.
type Sold struct {
	NrDoc   string
	DataDoc string
	Rest    string
	Termen  string
	Agent   string
	Valoare string
}

// Carnet represents a document book.
type Carnet struct {
	Simbol   string
	Denumire string
}

// ClientInfo represents a client from GetListaClienti.
type ClientInfo struct {
	CodIntern      string
	CodExtern      string
	Denumire       string
	CodFiscal      string
	Localitate     string
	Judet          string
	Adresa         string
	Telefon        string
	MarcaAgent     string
	DataFact       string
	SediiPart      string
	SimbolClasa    string
	DenumireClasa  string
	LocalitSedii   string
}
