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
// Format: "DenGestiune;SimbolGestiune;Denumire;CodExtern;ContContabil;UM;Stoc;PretUnitar;PretStoc;CotaTVA"
//
// [7] and [8] are UNIT prices, not stock values, whatever their old names said.
// Two measurements settle it: the same article held at 28 and at 29 units
// carries the identical [7], which no total can do; and all 25 rows with Stoc=0
// carry a non-zero [8], which no total can do either. Reading them as totals
// understated every cost by the quantity on hand.
type StocGestiune struct {
	DenGestiune    string // [0]
	SimbolGestiune string // [1]
	Denumire       string // [2]
	CodExtern      string // [3]
	ContContabil   string // [4]
	UM             string // [5]
	Stoc           string // [6]
	PretUnitar     string // [7] unit acquisition price, 2 decimals
	PretStoc       string // [8] unit price this warehouse records the article at, full precision:
	//     acquisition cost in cost-valued warehouses, retail price where the
	//     gestiune is held at selling price (account 371.04)
	CotaTVA string // [9]
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
// DetailedBalance is one line of GetSoldDetaliat. The manual documents 4
// fields; the live DLL sends 11, so splitting at 4 made Rest swallow the rest of
// the record — the scadenta, the currency and the carnet among them.
type DetailedBalance struct {
	Type          string // [0] "Factura" or "Avans"
	NrDocument    string // [1]
	DataDocument  string // [2]
	Rest          string // [3]
	TermenDePlata string // [4]
	Unknown5      string // [5] empty in every sampled row
	Moneda        string // [6]
	Curs          string // [7]
	Unknown8      string // [8] empty in every sampled row
	PrefixCarnet  string // [9]
	Unknown10     string // [10] trailing, empty
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
// Bank is one row of GetListaBanci. The manual documents 2 fields; the DLL
// returns 5, and position [3] carries the SWIFT/BIC code — which splitting at 2
// glued onto the end of Denumire.
type Bank struct {
	Simbol   string // [0]
	Denumire string // [1]
	Unknown2 string // [2] empty on every sampled row
	Swift    string // [3] BIC, set on 4 of 75 banks
	Unknown4 string // [4] trailing, empty
}

// Oferta represents a price offer.
// Oferta is one row of GetOferte, 13 fields live. Splitting at 6 made Cantitate
// absorb seven columns, so it never parsed as a number.
//
// Names follow the manual's Rev.1.5 list, which gives 14 and ends
// "...NrDoc;Marca;Agent". Only 13 arrive, and the values line up through Marca,
// so Agent is the column the manual has that the DLL does not send — the same
// thing it does with Prenume on Partner and Employee.
type Oferta struct {
	PartID       string // [0] partner code
	ArtID        string // [1] article code
	DataInreg    string // [2]
	DataExpir    string // [3]
	Pret         string // [4]
	CantMin      string // [5]
	AdDim        string // [6] adaos/diminuare
	ProcDiscount string // [7]
	Observatii   string // [8] equipment description in practice
	SimbolMoneda string // [9]
	CodlaFurn    string // [10]
	NrDoc        string // [11]
	Marca        string // [12] agent payroll number
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
	Unknown39           string // [39] "0" on 42% of rows
	Unknown40           string // [40] "0" on 42% of rows — absent from the April 2026 dump
	Unknown41           string // [41] empty — absent from the April 2026 dump
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
// SoldExt is one row of GetSolduriExt or GetSolduriFurn, 13 fields for a
// Factura. Splitting at 10 made ObservatiiFactura swallow the last three.
//
// Both readers share this shape but populate it differently: GetSolduriExt
// fills Observatii and a "BV AEG"-style carnet prefix, GetSolduriFurn leaves
// Observatii empty and puts a supplier series such as "UNIGOM" in the same slot.
//
// Note this is NOT a superset of Sold. From [9] on the two diverge — Sold has
// Moneda and Curs there, SoldExt has the observation and document series, and
// neither carries the other's columns.
//
// 48 of 464 rows (80 of 468 for Furn) are 7-field "Avans" rows that stop after
// [6]; check Tip before reading anything past it.
type SoldExt struct {
	IDPartener        string // [0]
	Tip               string // [1] "Factura" or "Avans"
	NrFactura         string // [2]
	DataFactura       string // [3]
	RestDePlata       string // [4]
	TermenDePlata     string // [5] Factura only
	LocatiePartener   string // [6]
	Unknown7          string // [7] empty on every sampled row
	ValoareFactura    string // [8] Factura only
	ObservatiiFactura string // [9] Factura only; empty for GetSolduriFurn
	PrefixCarnet      string // [10] "BV AEG" for Ext, a supplier series for Furn
	MarcaAgent        string // [11]
	SerieDocument     string // [12] e.g. "F.BV AEG"; empty for GetSolduriFurn
}

// ComandaNefacturata represents an uninvoiced order line.
// ComandaNefacturata is one row of GetComenziNefacturate, 25 fields live.
// Splitting at 4 left DenArticol holding the other 21 joined together — the
// field was never an article name at all, since the DLL puts the unit there.
//
// Only NumarComanda is read in production, as the idempotency guard before an
// order is pushed, and it parsed correctly even at width 4.
type ComandaNefacturata struct {
	CodArticol       string // [0]
	NumarComanda     string // [1]
	Cant             string // [2]
	DenUM            string // [3]
	DataComanda      string // [4]
	CodFiscalClient  string // [5]
	Unknown6         string // [6] "0" throughout
	Pret             string // [7] negative on the sampled rows
	CantLivrata      string // [8]
	DenUMLivrare     string // [9]
	CodFiscalLivrare string // [10]
	Unknown11        string // [11] 4 values, "PRO" dominant; not a warehouse symbol
	NrDocIntern      string // [12]
	Observatii       string // [13]
	TipSediu         string // [14] e.g. "SEDIU SOCIAL"
	Unknown15        string // [15] empty
	DataLivrare      string // [16] "30.12.1899" is the Delphi null date
	Unknown17        string // [17] empty
	Unknown18        string // [18] empty
	MarcaAgent       string // [19]
	DenPartener      string // [20]
	Unknown21        string // [21] "0" throughout
	ObservatiiAvans  string // [22] e.g. "Avans 50%"
	Moneda           string // [23]
	Unknown24        string // [24] trailing, empty
}

// CategoriePret represents a price category.
// CategoriePret is one row of GetListaCatPret. Undocumented, and this install
// has only the "nedefinit" placeholder, so [2] is named from its shape alone.
type CategoriePret struct {
	Simbol   string // [0]
	Denumire string // [1]
	Unknown2 string // [2] "0" on the single sampled row
}

// Sold represents an agent balance record.
// Sold is one row of GetSolduri, 11 fields for a Factura. Splitting at 6 shifted
// every field: NrDoc received the partner CUI and Valoare swallowed the tail.
//
// The reader also emits 9-field "Avans" rows — 48 of 464 here. Tip says which,
// and the advance layout stops after [8]; the currency and rate are absent, so
// read those only when Tip is "Factura".
type Sold struct {
	IDPartener      string // [0] partner code, per SetIDPartField
	Tip             string // [1] "Factura" or "Avans"
	NrDocument      string // [2]
	DataDocument    string // [3]
	RestDePlata     string // [4]
	TermenDePlata   string // [5] Factura only
	LocatiePartener string // [6]
	Unknown7        string // [7] empty on every sampled row
	ValoareDocument string // [8]
	Moneda          string // [9] Factura only
	Curs            string // [10] Factura only
}

// Carnet represents a document book.
// Carnet is one row of GetListaCarnete. Undocumented by the manual; the DLL
// returns 4 fields. Splitting at 2 made Denumire swallow the document-type
// discriminator, which is the only thing that makes the list useful.
type Carnet struct {
	Simbol   string // [0]
	TipDoc   string // [1] document type, e.g. "FACT"
	Denumire string // [2]
	Unknown3 string // [3] trailing, empty
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
