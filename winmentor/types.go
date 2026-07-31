package winmentor

// Partner represents a business partner (customer/supplier) from WinMENTOR.
// PDF documents 38 fields but the DLL actually returns 49.
// The field layout diverges from the PDF starting at index [17].
type Partner struct {
	ID                    string // [0]
	Denumire              string // [1]
	CodFiscal             string // [2]
	Localitate            string // [3]
	Adresa                string // [4]
	Telefon               string // [5]
	PersContact           string // [6]
	SimbolClasa           string // [7]
	DenClasa              string // [8]
	SimbolCatPret         string // [9]
	DenCatPret            string // [10]
	MarcaAgent            string // [11]
	NumeAgent             string // [12]
	Scadenta              string // [13] payment terms in days
	Discount              string // [14] always empty in observed data
	DenCritDiscount       string // [15] always empty in observed data
	DenumiriSediiPartener string // [16] "~" separated branch names
	CodExtern             string // [17] PDF says SediiPartener here, but actual data = CodExtern
	PartnerBlocat         string // [18] DA/NU
	CreditVanzare         string // [19]
	NrRegCom              string // [20] trade register number (e.g. "J08/10/2003"), undocumented in PDF
	ContBanca             string // [21]
	LocalitatiSedii       string // [22] "~" separated
	Judet                 string // [23] county code (e.g. "BV"), undocumented in PDF
	MarcaAgentiSedii      string // [24] "~" separated
	Observatii            string // [25]
	FlagSediuSocial       string // [26] "~" separated D flags
	CodPostalSedii        string // [27] "~" separated
	EmailSedii            string // [28] "~" separated
	TelPersContact        string // [29]
	PFsauPJ               string // [30] PF or PJ
	MonedaImplicita       string // [31]
	DataAdaugarii         string // [32]
	Trasee                string // [33]
	PuncteAcumulate       string // [34]
	CodFiscalSedii        string // [35] "~" separated
	InfoTipSediu          string // [36] "~" separated
	FlagClient            string // [37] "DA" when the partner is a customer
	FlagFurnizor          string // [38] "NU" when the partner is a supplier
	GLNSedii              string // [39] "~" separated
	SerieBuletin          string // [40] ID-card series, natural persons only
	NumarBuletin          string // [41] ID-card number, natural persons only
	Tara                  string // [42] country code (e.g. "RO", "BE")
	NrOrdineTraseu        string // [43] position within the assigned route
	BlocatIesiri          string // [44] blocked for outgoing documents
	ZileNeplata           string // [45]
	IDPrimaZiNeplata      string // [46]
	SoldLimita            string // [47]
	Unknown48             string // [48] always empty in observed data
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
// PDF documents 17 fields but the DLL actually returns 21.
type StockArticle struct {
	CodExtern      string // [0]
	Denumire       string // [1]
	UM             string // [2]
	PretVanzare    string // [3]
	Stoc           string // [4]
	SimbolClasa    string // [5]
	DenClasa       string // [6]
	IDProducator   string // [7]
	DenProducator  string // [8]
	IDFurnizor     string // [9]
	DenFurnizor    string // [10]
	SimbolGestiune string // [11]
	DenGestiune    string // [12]
	CotaTVA        string // [13]
	FlagTVAInclus  string // [14] D/NU - whether VAT is included in sale price
	PretCuTVA      string // [15]
	StocRezervat   string // [16]
	Unknown17      string // [17] undocumented, always "0" in observed data
	IDIntern       string // [18] internal WinMENTOR article ID
	Unknown19      string // [19] undocumented, always empty in observed data
	Unknown20      string // [20] undocumented, always empty in observed data
}

// ArticleStock represents the stock info for a single article.
type ArticleStock struct {
	CodExtern   string
	Denumire    string
	UM          string
	PretVanzare string
	Stoc        string
}

// StocGestiune represents a stock entry per warehouse from GetStocuriPeGestiuni.
// Format: "DenGestiune;SimbolGestiune;Denumire;CodExtern;ContContabil;UM;Stoc;ValoareStoc;ValoareStocPrecisa;CotaTVA"
type StocGestiune struct {
	DenGestiune        string // [0]
	SimbolGestiune     string // [1]
	Denumire           string // [2]
	CodExtern          string // [3]
	ContContabil       string // [4]
	UM                 string // [5]
	Stoc               string // [6]
	ValoareStoc        string // [7] total stock value (rounded)
	ValoareStocPrecisa string // [8] total stock value (precise)
	CotaTVA            string // [9]
}

// SoldPartener represents a partner's balance.
type SoldPartener struct {
	CodExtern string
	Denumire  string
	Sold      string
}

// Employee represents a WinMENTOR employee record.
//
// The manual documents "Nume;Prenume;Marca;..." but the DLL sends the full name
// in a single field: all 90 sampled rows carry 2 to 4 words in [0] ("GURGU ION",
// "OLARIU DIANA ELENA"), never a bare surname. Mapping a Prenume here shifted
// every later field one place left.
type Employee struct {
	Nume           string // [0] full name, surname first
	Marca          string // [1] payroll number — the Marca that MarcaAgent and GetIstoricVanzari refer to
	CNP            string // [2]
	EsteActiv      string // [3] Da/Nu
	EsteAgent      string // [4] Da/Nu
	SerieBuletin   string // [5]
	NumarBuletin   string // [6]
	CodPostal      string // [7]
	NumeUtilizator string // [8] WinMENTOR user name, i.e. a Logon value; set on 2 of 90 rows
	Unknown9       string // [9] empty in every sampled row
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
	PartID      string
	ArtID       string
	DataInceput string
	DataSfarsit string
	Pret        string
	Cantitate   string
}

// ClasaArticole represents an article class.
type ClasaArticole struct {
	Simbol   string
	Denumire string
}

// NomenclatorArticol represents an article in the full nomenclature.
// PDF documents 24 fields but the DLL actually returns 40.
type NomenclatorArticol struct {
	CodExtern           string // [0]
	Denumire            string // [1]
	DenUM               string // [2]
	PretVanzare         string // [3]
	SimbolClasa         string // [4]
	DenClasa            string // [5]
	CodExternProducator string // [6]
	DenProducator       string // [7]
	GestImplicita       string // [8]
	CodExternUnic       string // [9]
	CotaTVA             string // [10]
	DenUMSecundara      string // [11]
	ParitateUMSecundara string // [12]
	Masa                string // [13]
	Serviciu            string // [14] Da/Nu
	CodVamal            string // [15]
	PretMinim           string // [16]
	CantImplicita       string // [17]
	PretValuta          string // [18]
	DataAdaug           string // [19]
	Masa2               string // [20]
	PretVCuTVA          string // [21]
	Locatie             string // [22]
	PretReferinta       string // [23]
	Unknown24           string // [24] DA/NU flag
	Unknown25           string // [25]
	Unknown26           string // [26]
	Unknown27           string // [27] always "0"
	Unknown28           string // [28]
	Unknown29           string // [29] NU flag
	Unknown30           string // [30]
	Unknown31           string // [31]
	Unknown32           string // [32]
	Unknown33           string // [33]
	Unknown34           string // [34]
	Unknown35           string // [35]
	Unknown36           string // [36] NU flag
	Unknown37           string // [37]
	Unknown38           string // [38]
	Unknown39           string // [39]
}

// VanzareExt represents an extended sale record.
// PDF documents 18 fields but the DLL actually returns 23.
type VanzareExt struct {
	IDPartener    string // [0] always empty
	Zi            string // [1] day of month
	NrFactura     string // [2] invoice number
	CodArticol    string // [3] article code
	Cant          string // [4] quantity
	DenUM         string // [5] unit of measure
	Pret          string // [6] price
	DenGest       string // [7] warehouse name
	Unknown8      string // [8] always "0"
	LocatieClient string // [9] client location/branch
	Unknown10     string // [10] always "0"
	CodFiscal     string // [11] customer tax ID
	Unknown12     string // [12] always "0"
	Adresa        string // [13] customer address
	Unknown14     string // [14] always "0"
	CodPostal     string // [15] postal code
	ClasaArticol  string // [16] article class
	TipDocument   string // [17] document type ("=", "S")
	Unknown18     string // [18]
	PrefixCarnet  string // [19] booklet prefix
	Moneda        string // [20] currency
	Unknown21     string // [21]
	Unknown22     string // [22]
}

// VanzareLuna represents a monthly sale line.
// PDF documents 10 fields but the DLL actually returns 26.
type VanzareLuna struct {
	IDPartener        string // [0] always empty
	Zi                string // [1] day of month
	NrFactura         string // [2] invoice number
	CodArticol        string // [3] article code
	NumarComanda      string // [4] order number
	Cant              string // [5] quantity
	DenUM             string // [6] unit of measure
	Pret              string // [7] price
	MarcaAgent        string // [8] agent code
	ValoareFactura    string // [9] invoice value with VAT
	DataScadenta      string // [10] due date
	TVAInclus         string // [11] VAT included flag ("NU")
	CotaTVA           string // [12] VAT rate
	TipDocument       string // [13] document type ("F")
	PrefixCarnet      string // [14] booklet prefix
	SerieDocument     string // [15] full document series
	DenArticol        string // [16] article description
	Unknown17         string // [17] always "0"
	Unknown18         string // [18]
	DataEmitere       string // [19] issue date
	SediuClient       string // [20] client branch
	AdresaClient      string // [21] client address
	LocalitateClient  string // [22] client town
	ObservatiiFactura string // [23] invoice notes
	Observatii2       string // [24] additional notes
	Unknown25         string // [25]
}

// Intrare represents an incoming entry (purchase) line.
// DLL returns 11 fields.
type Intrare struct {
	IDPartener string // [0] always empty
	Data       string // [1] date
	NrDoc      string // [2] document number
	CodArticol string // [3] article code
	Cant       string // [4] quantity
	DenUM      string // [5] unit of measure
	Pret       string // [6] price/value
	DenGest    string // [7] warehouse name
	Unknown8   string // [8] always "0"
	Flag       string // [9] "DA" or empty
	Unknown10  string // [10]
}

// SoldExt represents an extended balance line.
type SoldExt struct {
	IDPartener        string
	Tip               string // "Factura" or "Avans" etc.
	NrFactura         string
	DataFactura       string
	RestDePlata       string
	TermenDePlata     string
	LocatiePartener   string
	MarcaAgent        string
	ValoareFactura    string
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
	CodIntern     string
	CodExtern     string
	Denumire      string
	CodFiscal     string
	Localitate    string
	Judet         string
	Adresa        string
	Telefon       string
	MarcaAgent    string
	DataFact      string
	SediiPart     string
	SimbolClasa   string
	DenumireClasa string
	LocalitSedii  string
}
