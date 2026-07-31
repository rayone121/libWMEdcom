# Record layouts

Field layouts for every reader in `DocImpServer`, reconciled from three sources and
adversarially verified. **Read this before changing any `splitFields` count.**

## Where these come from, and which source wins

| Source | Authoritative for |
|---|---|
| `Functii pentru interfatarea cu WinMENTOR.pdf` Rev.1.5, 22 May 2026 | **field names**, for the positions it documents |
| `vendor-docs/23_WMGenCmd__Structura_returnata_GetListaParteneri.txt` | 48 numbered fields — beats the manual's 38 |
| `DocImpServer.dll` (a .NET COM interop assembly; `ikdasm` it) | **method signatures** — says nothing about records |
| `wme-raw/*.txt` dumps | **field counts** |

The vendor manual **understates counts** on almost every reader: 24 documented for
`GetNomenclatorArticole` against 40 in the April dump and **42 live**. So take names from the
manual and the count from the data, never the reverse.

Download the current manual from
<http://download.winmentor.ro/WinMentor/Documentatie/24_DocImpServer/>. Per-document *import*
structures (bon de consum, aviz de expeditie, transfer, …) are published separately under
`22_Structuri import din alte aplicatii/`.

## A third hazard: the manual invents columns

`GetListaParteneri` and `GetListaPersonal` are both documented with a `Prenume`
next to `Nume`. The DLL sends neither — the full name arrives in one field, and
mapping a forename beside it shifts everything after it one place left. On
`GetListaPersonal` that gave every employee a CNP in `Marca`, an ID-card series
in `EsteAgent` and an empty `CodPostal`. Both are fixed; the lesson is that the
manual can be wrong about a column's *existence*, not just its count, so a name
taken from it still has to be checked against a value.

## Two hazards that have each already caused a bug

**The separator appears inside the data.** 46 `CodExtern` values contain `;` — `35mmp  ; H01N2-D`,
`TVH/923917 ; 10'`. `SplitN(rec, ";", N)` shifts every field after the code on those rows, so the
unit of measure lands in `PretVanzare`.

**The live width is not the dump width.** A merge conditioned on `len(parts) > N` with `N` taken
from a stale dump fires on *every* record the day the DLL returns a wider row. That is exactly
what happened: `GetNomenclatorArticole` at `N=40` against a live 42 absorbed three fields into the
article code and destroyed the join key for all 64 322 articles (commit `1db37e8`, reverted in
`b00482a`). Any repair must key off a width **measured in the same session**, and must degrade to
"leave it alone" when the width is unexpected.

Measure the live width first: `wmedump -profile` writes a column report beside every dump — width
histogram, fill rate, cardinality, inferred type and samples per position.

## Status

`manual` = positions the vendor names · `true` = positions in the April 2026 dump ·
`go` = current `splitFields` count (`-` = returns raw strings)


| Reader | manual | true | go | status |
|---|--:|--:|--:|---|
| `GetIntrari` | — | 10 | 11 | misaligned |
| `GetListaParteneri` | 38 / **48** | 49 | 49 | **fixed** (ec1b31d) |
| `GetListaPersonal` | 9 | 10 | 10 | **fixed** (6ec93ff) |
| `GetNomenclatorArticole` | 24 | 40 | 40 | misaligned |
| `GetSolduri` | 6 | 11 | 6 | misaligned |
| `GetSolduriExt` | 10 | 13 | 10 | misaligned |
| `GetSolduriFurn` | 10 | 13 | 10 | misaligned |
| `GetStocArticole` | 17 | 21 | 21 | misaligned |
| `GetStocuriPeGestiuni` | — | 10 | 10 | misaligned |
| `GetVanzariExt` | 18 | 23 | 23 | misaligned |
| `GetVanzariLuna` | 10 | 26 | 26 | misaligned |
| `GetComenziNefacturate` | 4 | 24 | 4 | truncating |
| `GetListaBanci` | 2 | 5 | 2 | truncating |
| `GetListaCarnete` | — | 4 | 2 | truncating |
| `GetListaCatPret` | — | 3 | 2 | truncating |
| `GetOferte` | 14 | 11 | 6 | truncating |
| `GetIncasariLuna` | — | 5 | — | unparsed |
| `GetInfoComenzi` | 11 | 9 | — | unparsed |
| `GetInfoIesiri` | 18 | 17 | — | unparsed |
| `GetInfoIesiriExt` | — | 9 | — | unparsed |
| `GetListacarneteExt` | — | 6 | — | unparsed |
| `GetMonede` | — | 2 | — | unparsed |
| `GetReceptii` | — | 21 | — | unparsed |
| `GetStocArticoleExt` | — | 20 | — | unparsed |
| `GetTransferuri` | — | 14 | — | unparsed |
| `GetListaClienti` | 14 | 14 | 14 | correct |
| `GetListaGestiuni` | 2 | 3 | 3 | correct |
| `GetArticoleVandute` | — | — | — | no-data |
| `GetIncasariClienti` | — | — | — | no-data |
| `GetSoldDetaliat` | 4 | — | 4 | no-data |
| `GetSoldPart` | 3 | — | 3 | no-data |
| `GetStocArticol` | 5 | — | 5 | no-data |


## Layouts

### `GetArticoleVandute`

no sample dump, not documented by the vendor.

```
UNKNOWN — no sample and no documented structure. The manual says only that it returns "un sir cu articolele vandute", which does not even establish whether each element is a bare article code or a ;-joined record. Nothing can be named.
```

**Change:** NO CHANGE. Raw []string is correct for a reader with neither sample nor layout; do not invent a struct. Action item: add a GetArticoleVandute call to cmd/wmedump (a known PartID + MarcaAgent 52 + a start month) so the layout can be established from data before anyone parses it. Same for GetUltimeleVanzari (dispid 46), its undocumented sibling at queries.go, which is in the same state.

### `GetComenziNefacturate`

24 positions in the April 2026 dump, 4 of them named by the vendor manual.

```
[0] IDArticol -> string CodExtern. non-empty 1894/2876. MAY CONTAIN ';'
[1] NumarComanda -> string(int). per-document constant; range 0..20258; same numbering space as GetInfoComenzi NrDoc
[2] Cant -> string RO-decimal
[3] DenUM -> string. Lei, L, ORA, Buc, cursa. MANUAL IS WRONG HERE: it names this "Den Articol"
[4] Data -> string dd.mm.yyyy; 30.12.1899 = Delphi null date
[5] IDPartener -> string CodFiscal
[6] MarcaAgent -> string(int). "0" on 2865 rows, "52" on 10
[7] Pret -> string RO-decimal. unit price
[8] Cant (repeat of [2]) -> string; identical to [2] on every sampled row
[9] DenUM (repeat of [3]) -> string; identical to [3] on every sampled row
[10] IDPartener (repeat of [5]) -> string; non-empty 2214/2876
[11] SimbolCarnet -> string. BVAEG (1736), AEG CC (541), COM AEG (444), PRO (152)
[12] NrDocCarnet -> string(int). second per-document number, range 63..19459, parallel slower sequence than [1]. UNRESOLVED which of [1]/[12] WinMENTOR treats as the printed order number
[13] ObservatiiDoc -> string free text. "Conform oferta nr.2183/...", "ANULAT", "La rep."
[14] SediuPartener -> string. SEDIU SOCIAL, Brasov, Prejmer; non-empty 347/2876
[15] <unnamed> -> string; empty on 2875/2875 well-formed rows
[16] TermenLivrare -> string date; 30.12.1899 on 2834 rows
[17] <unnamed> -> string; empty on all well-formed rows
[18] <unnamed> -> string; empty on all well-formed rows
[19] NrLinie -> string(int). CONFIRMED: order 8980 has 32 lines carrying exactly 1..32
[20] NumePartener -> string
[21] <unnamed constant> -> string, "0" on 2875/2875
[22] ObservatieLinie -> string free text, non-empty 58/2876. "Avans 50%", "LINDE H40D", "Brasov-Sibiu", "12V-100Ah"
[23] SimbolMoneda -> string. RON on 2875/2875
[24] <trailing separator artifact> -> always ""; NOT a field
```

**Change:** THIS IS THE WORST OF THE SEVEN AND IT IS LOAD-BEARING. splitFields(rec,4) puts fields 4..24 into DenArticol as one 20-field garbage string: DenArticol = "Lei;01.12.2016;RO4188153;0;-56,09;1;Lei;RO4188153;PRO;5574;Conform oferta...;;;30.12.1899;;;9;EUROPEAN DRINKS S.A.;0;;RON;". There is no article denumire in the record, so the field cannot be salvaged - DELETE DenArticol, do not try to fill it.
Why this matters beyond cosmetics: ComandaNefacturata.NumarComanda ([1]) is the idempotency guard for the whole Dolibarr->WinMENTOR order push (cmd/wmeprobe/main.go:363 and :1171; the syncer refuses to push if this read fails, and skips the import if the number is already present). SplitN(4) happens to leave [0],[1],[2] intact, so the guard works TODAY on well-formed rows - but on a row whose CodExtern contains ';' (1 in 2876 in the sample) NumarComanda receives " Quartz Energy 9000" instead of 4761, so that order number silently drops out of the guard set and a duplicate import becomes possible. Before trusting the guard further, verify on live which of [1]/[12] receives the syncer's NrDoc (500000+id / 600000+id / 800000+id): wmeprobe.log:569-582 shows every candidate as inWME=false, so the guard has never actually matched anything and the assumption is untested.
PARSING RULE FOR THE EMBEDDED ';' (must not use total width): parts := strings.Split(rec, ";"); the article code is the variable-width HEAD, so scan k = 1,2,3,... for the smallest k where parts[k+3] matches ^\d{1,2}\.\d{1,2}\.\d{4}$ (the Data anchor) - for the Quartz row k=1 gives parts[4]="L" (reject), k=2 gives parts[5]="01.12.2016" (accept) - then IDArticol = strings.Join(parts[0:k], ";") and every later field is read at a FRONT-anchored offset k-1+i with a bounds-safe accessor. Extra trailing fields a future WME build appends then fall off the end and are ignored instead of triggering a merge. Never branch on len(parts) > 25: that is exactly the rule that destroyed the nomenclator join key when the live width moved from 40 to 42.

### `GetIncasariClienti`

no sample dump, not documented by the vendor.

```
UNKNOWN — no sample, and the manual gives no "pe structura" list. The two manual entries even contradict each other on the shape of the return:
  #33 says a single total value  -> the return would be 1 element, unsplit
  #49 says a list of collections -> the return would be N ;-joined records
Until a dump exists, no position can be named. Any struct written now would be a guess.
```

**Change:** NO CHANGE. Returning raw []string is the correct conservative behaviour for a reader with no sample and no documented layout; do not add a struct. The doc comment at queries.go:163 states "the total collection value", which follows manual entry #33 while entry #49 says it is a list — soften it to note the contradiction. Action item: add GetIncasariClienti(year,month..,PartID) to cmd/wmedump so a real sample exists before anyone types a struct.

### `GetIncasariLuna`

5 positions in the April 2026 dump, not documented by the vendor.

```
[0] NrDocIncasare -> string/int, the payment document number. Values across the 48 records: '4' x27, '3' x15, '2808', '391336', '391337', '391338', '60', '2'. NOT the day: record '4;01.04.2026;...' has [0]=4 on 01 April, and '3;03.04.2026' / '4;02.04.2026' / '4;04.04.2026' all occur, so [0] is independent of [1]. The small repeated values (2,3,4,60) behave like bank-statement / receipt numbers reused across many lines of one payment; the 391336..391338 run behaves like consecutive payment-order numbers.
[1] DataIncasare -> string date dd.mm.yyyy ('01.04.2026' x14, '02.04.2026' x17, '03.04.2026' x15, '04.04.2026' x2)
[2] DocumentFactura -> string, the settled invoice, always 'F.' + PrefixCarnet + NrFactura ('F.BV AEG35743'); joins 1:1 to GetVanzariLuna[15]
[3] CodExternPartener -> string, customer CIF ('RO43678520'); empty on 3/48 rows (F.BV AEG36043 x2, F.BV AEG36049 x1 — walk-in/cash customers)
[4] SumaIncasata -> RO decimal, amount collected on that invoice ('6670,73','23896,37','42')
No [5]: there is no terminator and no trailing empty column.
```

**Change:** UNPARSED but not broken. If a struct is wanted, it is safe to add one here — this is the only reader in this group with a single stable width, no embedded separators and no free text. Note two things in the struct: (1) there is NO trailing terminator, so a 5-field parse is exact; (2) [0] must be typed as a string, not an int day-of-month, and must not be confused with Zi. Even so, use strings.Split + a bounds-checked accessor rather than SplitN(rec,";",5): if the live DLL ever appends a column, SplitN(5) would silently glue it onto SumaIncasata[4] — the amount, i.e. the one field that must never be corrupted. Keep any surplus in an Extra []string.

### `GetInfoComenzi`

9 positions in the April 2026 dump, 11 of them named by the vendor manual.

```
[0] IDPart -> string. CodFiscal-shaped (RO1093261); non-empty 1349/1351
[1] Data -> string dd.mm.yyyy
[2] NrDoc -> string(int). range 7..20317
[3] IDArticol -> string CodExtern. non-empty 666/1351 ("134660MM", "Azolla ZS46")
[4] CantInit -> string RO-decimal
[5] Pret -> string RO-decimal. 65, 94, 109, 89
[6] AdaosDiminuare -> string. "0" on 697 rows, "" on 654
[7] ProcDiscount -> string. "0" on the same 697 rows (perfectly correlated with [6])
[8] FlagAnulat -> string. "NU" on 1351/1351
DOCUMENTED BUT NOT EMITTED: Marca (10th name), Agent (11th name)
```

**Change:** Add an InfoComanda struct with the 9 documented-and-emitted names. Do NOT reserve Marca/Agent. Note the numbering relationship worth exploiting: [2] NrDoc lives in the SAME numbering space as GetComenziNefacturate field [1] - e.g. GetComenziNefacturate has orders 4217 and 4219 for RO10334547 while GetInfoComenzi has 4218 for the same partner and article - which independently corroborates that GetComenziNefacturate[1] is the manual's "Numar Comanda".

### `GetInfoIesiri`

17 positions in the April 2026 dump, 18 of them named by the vendor manual.

```
[ 0] TipDoc -> int ('19' x7 in the April dump; live wmeprobe shows '19' x873, '24' x395, '22' x3, '' x24)
[ 1] NrDoc -> int ('36050'..'36056')
[ 2] Data -> string date dd.mm.yyyy
[ 3] Numepartener -> string ('MULTICO SRL', '"TRAMAR" SA')
[ 4] NrAuto -> string, empty on 7/7
[ 5] Masa -> numeric, empty on 7/7
[ 6] Observatii -> free text up to 410 chars; MAY CONTAIN ';'
[ 7] Numsediu -> string ('SEDIU SOCIAL' x5, empty x2)
[ 8] Localitate -> string ('ARAD','BUCURESTI SECTORUL 5','GHIMBAV')
[ 9] Judet -> string ('AR','IF','BV','CV')
[10] Tara -> string ('RO' x7)
[11] Adresa -> string
[12] ValDoc -> RO decimal ('350','144215,74','8349','4163,63')
[13] CodExternPart -> string CIF ('RO1678142','RO21542446')
[14] CodInternPart -> int ('2618','2697','4232','358')
[15] CodPostal -> string, empty on 7/7
[16] (terminator) -> always empty
Marca and Agent, positions 16 and 17 in the Rev.1.5 layout, are NOT emitted: the record stops after CodPostal.
```

**Change:** UNPARSED. The Rev.1.5 layout can now be trusted for [0..15] and a struct may be written, but it MUST stop at CodPostal[15] and MUST NOT allocate Marca/Agent — no sample, April or live July, delivers them. Because Observatii[6] is free text that can carry ';', use the same left-anchored rule: [0..5] fixed (TipDoc int, NrDoc int, Data date-shaped — validate), then scan forward from 6 for the first index k where parts[k] is 'SEDIU SOCIAL'-or-empty AND parts[k+3] matches ^[A-Z]{2}$ (Tara) — or, simpler and safer, treat Observatii as the greedy remainder anchored from the RIGHT only after shape-validating Tara/ValDoc/CodExternPart. Never SplitN(rec,";",17).

### `GetInfoIesiriExt`

9 positions in the April 2026 dump, not documented by the vendor.

```
[0] NrLinie -> int, 1-based line number within the document (1,2,3,... resets per document)
[1] DenumireArticol -> string ('Jacheta EMERTON STRECH gri/orange mar.48', 'Transport cu platforma')
[2] Cant -> RO decimal
[3] DenUM -> string ('Buc','Zi','cursa')
[4] (unnamed) -> empty on all 20 sampled rows
[5] Pret -> RO decimal ('131,63','400','2621,9')
[6] CodExtern -> string, article CodExtern. CONFIRMED (see evidence). Empty when the article itself has an empty CodExtern.
[7] IDArticol -> string, the article's unique internal identifier = GetNomenclatorArticole field [9]. CONFIRMED. NOT always an integer.
[8] (terminator) -> always empty
```

**Change:** UNPARSED, and correctly so given there is no dump file — but the layout above is now established well enough to write a 9-field struct if a consumer needs one. Two hard requirements: (1) [7] IDArticol must be typed string (value '429.1' proves it is not an integer); (2) [7] is the correct join key to GetNomenclatorArticole[9] and [6] is the correct join key to GetNomenclatorArticole[0] — but [6] can be empty for service articles, so any Dolibarr matching must key on [7] and fall back to [6], never the reverse. Also add GetInfoIesiriExt to cmd/wmedump so a real width histogram exists; 20 rows scraped from a log is not a sample.

### `GetIntrari`

10 positions in the April 2026 dump, not documented by the vendor.

```
[0] IDPartener -> string CodFiscal. DE204326380, RO11201891, 35135779 (unregistered CUI); empty on 16/6266
[1] Data -> string dd.mm.yyyy zero-padded
[2] NrDocFurnizor -> string. supplier invoice number; VERIFIED == GetReceptii[6]
[3] CodArticol -> string CodExtern; non-empty 5340/6266. MAY CONTAIN ';'
[4] Cant -> string RO-decimal; can be negative (-1 seen)
[5] DenUM -> string. Lei, Buc, L, Set, Kg, Cutii
[6] Pret -> string RO-decimal. VERIFIED == GetReceptii[14] (supplier invoice unit price, NOT the landed cost)
[7] SimbolGestiune -> string. Piese, OI, M aux, Utilaje, Imob - symbols from GetListaGestiuni; empty on 1937/6266
[8] <unnamed numeric> -> string RO-decimal. "0" on 6233/6266; otherwise a per-DOCUMENT constant (2,41 on all three lines of doc 76685; 2,62; 2,6; 2,21; 0,39). Cannot be named from the data
[9] FlagServiciu -> string "DA"/"". CONFIRMED complementary to [7]: cross-tab is 1850 (no gestiune, DA) / 4307 (has gestiune, "") / 87 (no gestiune, "") - never DA with a gestiune
[10] <trailing separator artifact> -> always ""; NOT a field
```

**Change:** WIDTH IS RIGHT, ALIGNMENT IS NOT. splitFields(rec,11) == strings.SplitN(rec,";",11) matches the 11 tokens, so 6244/6266 rows parse correctly - but on the 22 rows whose CodExtern contains ';' every field from [3] on shifts and the struct silently receives text in numeric slots. For `RO6377426;20.05.2025;177856;111846SA ; A1002;1;Buc;501;Piese;0;;` the caller gets CodArticol="111846SA ", Cant=" A1002", DenUM="1", Pret="Buc", DenGest="501", Flag="0". Real examples in the dumps: "111846SA ; A1002", "780040  ; 600ML", "Q8 HAYDN 46 ; 20L", "AD1030 - 5L ; Dreissner", "35mmp  ; H01N2-D", "O5W305L ; 2.5 L".
FIX: strings.Split + bounds-safe accessor; CodArticol is the variable-width field at index 3 with a FIXED 3-token head ([0] CUI, [1] date, [2] doc number - none can contain ';'), so scan k = 1,2,3,... for the smallest k where parts[3+k] parses as an RO-decimal (Cant) AND parts[4+k] is a non-empty non-numeric token (DenUM), then CodArticol = strings.Join(parts[3:3+k], ";") and read the rest at front-anchored offsets. Note the residual ambiguity honestly: a code ending in a numeric chunk followed by a word (e.g. "...; 5;L") could fool the scan; none exists in 6266 rows, and the fallback must be k=1 (no merge), never a width comparison.
Also rename DenGest -> SimbolGestiune, Flag -> FlagServiciu, drop the false "always 0" comment on [8], and keep [8] as a named-unknown.

### `GetListaBanci`

5 positions in the April 2026 dump, 2 of them named by the vendor manual.

```
[0] Simbol -> string (short bank code, e.g. "BCR", "BTRL", "RABO", "Volk-Olden")
[1] Denumire -> string (full bank name; may contain double quotes, e.g. Banca Comerciala "Ion Tiriac")
[2] UNNAMEABLE -> string, empty in 75/75 rows (undocumented; no sample value exists, so it cannot be named from evidence. By analogy with the partner bank block in types.go:72-75 (SimbolBanca/NumeBanca/LocalitateBanca/ContBanca) the plausible candidate is Localitate or Cont, but that is inference, not evidence - keep as named-unknown)
[3] CodSWIFT / BIC -> string (4/75 non-empty, every value a valid SWIFT-BIC: "CITIROBU" (CityBank RO), "RABONL2U" (Rabobank NL), "GENODEF1EDE" (Volksbank DE), "DRESDEFF734" (Commerzbank DE). All four are the foreign/correspondent banks in the list, which is exactly where a BIC is needed)
[4] (trailing separator artifact) -> always empty string
```

**Change:** Split into 5. Keep Bank.Simbol=[0], Bank.Denumire=[1] (now clean), add CodSWIFT=[3] and a named-unknown for [2] to types.go:235-238. At queries.go:71 replace `splitFields(rec, 2)` with the unbounded split + bounds-checked accessor. Note that today every Denumire in this reader carries a ";;;" tail, so any consumer string-matching on bank names is currently failing.

### `GetListaCarnete`

4 positions in the April 2026 dump, not documented by the vendor.

```
[0] SimbolCarnet -> string ("BV AEG", "MAGAEG", "CH AEG", "AEG", "AEG AF", "TEMP" - joins to GetListacarneteExt[0])
[1] TipCarnet / document-class tag -> string (constant "FACT" in all 6 rows. NOT the TipDoc symbol: the same carnets in GetListacarneteExt carry FF / CH / AEG / AF / TEMP at that position. Consistent with the manual's prose that this reader returns only the customer-outgoing/transfer books, i.e. the class is pre-filtered to FACT. Naming is evidence-thin - a single constant value - so treat as provisional)
[2] Denumire -> string ("Facturi AEG Tech", "Facturi AEG Independentei", "Chitante AEG", "Facturi Smart", "AUTOFACTURI", "TEMP")
[3] (trailing separator artifact) -> always empty string
```

**Change:** Split into 4. Add TipCarnet=[1] to types.go:415-418 and move Denumire to [2]. At queries.go:201 replace `splitFields(rec, 2)` with the unbounded split + accessor. Because the manual documents nothing here, add a comment recording that the layout is data-derived from the 6-row April sample and that [1] was constant "FACT" - so a future non-FACT value is expected, not a parse failure.

### `GetListaCatPret`

3 positions in the April 2026 dump, not documented by the vendor.

```
[0] Simbol -> string (empty in the only row)
[1] Denumire -> string ("---- nedefinit ----")
[2] UNNAMEABLE -> int-looking string, value "0" (undocumented; ONE sample of ONE row, so it cannot be named from evidence. Candidates are an ID/index or a boolean 'implicita' flag - the DLL does expose SetCatPretImplicita(IDCatPret), wrapped at queries.go:867 - but with a single all-sentinel row there is nothing to discriminate. Keep as named-unknown)
```

**Change:** Split into 3, keep Simbol=[0] and Denumire=[1] (now clean), add a named-unknown for [2] to types.go:399-402. At articles.go:204 replace `splitFields(rec, 2)`. Given the one-row all-sentinel sample, flag this reader explicitly as LOW-CONFIDENCE in the code comment and re-verify against a company that has price categories before any consumer relies on [2].

### `GetListaClienti`

14 positions in the April 2026 dump, 14 of them named by the vendor manual.

```
[ 0] CodIntern      -> int-as-string (manual 1)
[ 1] CodExtern      -> string (manual 2)
[ 2] Denumire       -> string (manual 3)
[ 3] CodFiscal      -> string (manual 4)
[ 4] Localitate     -> string (manual 5)
[ 5] Judet          -> string, 1-2 letter auto-code (manual 6)
[ 6] Adresa         -> string (manual 7 'Adr')
[ 7] Telefon        -> string, '~'-joined per branch (manual 8 'Tel')
[ 8] MarcaAgent     -> int-as-string (manual 9)
[ 9] DataFact       -> date dd.mm.yyyy (manual 10)
[10] SediiPart      -> string, '~'-joined (manual 11)
[11] SimbolClasa    -> string (manual 12 'Simbol_Clasa'); empty in all 1647 records
[12] DenumireClasa  -> string (manual 13 'Denumire_Clasa'); empty in all 1647 records
[13] LocalitSedii   -> string, '~'-joined, 'TOWN JJ' form (manual 14)
```

**Change:** CORRECT — no field change needed. types.go:421-436 and partners.go:240 already match the manual and the data exactly.

One robustness change only, the same one every reader in this group needs: partners.go:240 uses splitFields, i.e. strings.SplitN(rec, ";", 14), so if the live DLL ever returns a 15th column, LocalitSedii silently swallows it (the exact class of silent corruption that motivated this audit), and if a Denumire or Adresa ever contains a ';' the record is silently shifted with no signal. Replace with strings.Split + explicit width handling: <14 -> loud error; ==14 -> map; >14 -> map [0..13] and keep the remainder in Extra []string, never merge. Shape anchor for detecting a leaked ';' in the free-text head: parts[9] must match dd.mm.yyyy and parts[8] must be numeric-or-empty; both hold for 1647/1647 records today, so a failure is a real alarm.

### `GetListaGestiuni`

3 positions in the April 2026 dump, 2 of them named by the vendor manual.

```
[0] Simbol -> string (warehouse code; EMPTY on the sentinel row. Values: "", F, Imob, OI, EuroDiscount, Pr, "M pr", Serv, "M aux", Piese, "Mf cust", Utilaje, Manop, "Dep Rei", Q8 - note embedded spaces)
[1] Denumire -> string ("---nedefinita---", "Firma", "Imobilizari", "Materiale de natura Ob. invent", ...)
[2] (trailing separator artifact) -> always empty string
```

**Change:** None required. Optionally migrate off splitFields to the shared unbounded split + accessor for uniformity, which would also make it tolerant if a live build ever adds a real field at [2] - today SplitN(rec,3) would silently glue any such new field onto the currently-empty [2] slot rather than dropping it, so the tolerance is worth having even though nothing is broken now.

### `GetListaParteneri`

49 positions in the April 2026 dump, 38 of them named by the vendor manual.

```
[ 0] IDPartener                -> string (manual 1)  e.g. 'RO15129446'; 4840 distinct; equals CodFiscal and CodExtern here
[ 1] Denumire                  -> string (manual 2)
[ 2] CodFiscal                 -> string (manual 3)
[ 3] Localitate                -> string (manual 4)  BRASOV x1270
[ 4] Adresa                    -> string (manual 5)
[ 5] TelefonSedii              -> string, '~'-joined per sediu (manual 6 'Telefon')
[ 6] PersoanaContact           -> string, '~'-joined; MAY CONTAIN A NEWLINE (manual 7)
[ 7] SimbolClasa               -> string (manual 8)  always '...'
[ 8] DenClasa                  -> string (manual 9)  always '---nedefinita---'
[ 9] SimbolCatPret             -> string (manual 10) always empty
[10] DenCatPret                -> string (manual 11) '---- nedefinit ----'
[11] MarcaAgent                -> int-as-string (manual 12)  '0'/'1'/'52'
[12] NumeSiPrenumeAgent        -> string == MANUAL 13 + ' ' + MANUAL 14 COLLAPSED INTO ONE FIELD. This single field is the whole reason every later position is off by one.
[13] Scadenta                  -> int days (manual 15)  15/30/45/0
[14] Discount                  -> decimal RO (manual 16)  empty in all 4866
[15] DenCritDiscount           -> string (manual 17)  empty in all 4866
[16] DenumiriSedii             -> string, '~'-joined (manual 18)  'SEDIU SOCIAL', 'Social'
[17] CodExtern                 -> string (manual 19)
[18] PartenerBlocat            -> enum DA/NU (manual 20)  NU x4865, DA x1
[19] CreditVanzare             -> decimal RO (manual 21)  always '0'
[20] NrRegCom                  -> string (manual 22, WHICH MISNAMES IT 'Cod fiscal')  'J08/10/2003'
[21] ContBanca                 -> string 'SIMBOL-IBAN' (manual 23)  'BTRL-RO80BTRL0080...'
[22] LocalitatiSedii           -> string, '~'-joined (manual 24)  'BRASOV BV'
[23] SimbolAutoJudet           -> string (manual 25, WHICH MISNAMES IT 'Tara')  BV/B/IF/CV, 47 distinct
[24] MarciAgentiSedii          -> string, '~'-joined (manual 26)
[25] Observatii                -> string (manual 27)  17 nonempty
[26] FlagSediuSocial           -> string, '~'-joined D flags (manual 28)
[27] CoduriPostaleSedii        -> string, '~'-joined (manual 29)
[28] EmailuriSedii             -> string, '~'-joined (manual 30)
[29] TelefoanePersContact      -> string, '~'-joined; MAY CONTAIN A NEWLINE (manual 31)
[30] TipPartener               -> enum PJ/PF (manual 32)
[31] MonedaImplicita           -> string (manual 33)  Lei/Eur/LEIM/USD, 6 distinct
[32] DataAdaugarii             -> date dd.mm.yyyy (manual 34)
[33] TraseeSedii               -> string, '~'-joined (manual 35)
[34] PuncteAcumulate           -> decimal RO (manual 36)  '0' in 947, empty in 3919
[35] CoduriFiscaleSedii        -> string, '~'-joined, trailing '~' (manual 37)
[36] TipuriSedii               -> string, '~'-joined, trailing '~' (manual 38)  'Sediul~Punct lucru~'
--- undocumented from here on; the manual stops at position [36] ---
[37] FlagProducator            -> enum DA/NU  INFERRED (DA only for a forklift manufacturer + one garment-tech firm; matches PartnerInput.FlagProducator)
[38] UNNAMED_Flag38            -> enum DA/NU  CANNOT NAME CONFIDENTLY. Either TVAIncasare (PartnerInput ordering) or a VAT-exempt/reverse-charge flag (7 DA rows: 5 foreign + an insurer + an agrotourism firm).
[39] UNNAMED_ListaSediu39      -> string, '~'-joined per sediu, arity identical to [16]; every element empty in 4866 records. CANNOT NAME.
[40] SerieActIdentitate        -> string, 2 letters  INFERRED = PartnerInput.SerAI (set for 35 of 59 PF partners, 1 PJ)
[41] NumarActIdentitate        -> string, 6 digits   INFERRED = PartnerInput.NrAI
[42] Tara                      -> ISO-2 country code  INFERRED = PartnerInput.SimbolAutoTaraSediu (RO/DE/NL/PL/IT/IE/BG, 33 distinct). THIS is the manual's 'Tara', not position [23].
[43] UNNAMED_ListaSediu43      -> string, '~'-joined per sediu, all elements empty. CANNOT NAME.
[44] UNNAMED_Flag44            -> enum DA/NU; 'NU' in all 4866 records. CANNOT NAME.
[45] UNNAMED_Numeric45         -> decimal RO; '0' in the same 947 records as [34]. CANNOT NAME.
[46] UNNAMED_Numeric46         -> decimal RO; '0' in the same 947 records. CANNOT NAME.
[47] UNNAMED_Numeric47         -> decimal RO; '0' in the same 947 records. CANNOT NAME.
[48] UNNAMED_Trailing48        -> string; empty in all 4866 records. CANNOT NAME.
```

**Change:** MISALIGNED at [13]..[16] — the width 49 in partners.go:22 is correct, the names are not. Concretely, today Partner.PrenumeAgent receives the payment term (15/30/45), Partner.Scadenta and Partner.Discount are permanently empty, and Partner.DenCritDiscount receives the branch-name list ('SEDIU SOCIAL'). This is LIVE: internal/sync/orchestrator.go:159 calls GetListaParteneri; the damage is currently confined to cmd/wmecsv/main.go:123-141, which prints all five wrong columns, and to Scadenta being unavailable to any future consumer.

FIX, in /Users/raymond/Documents/Repos/libWMEdcom/winmentor/types.go:6-56 and partners.go:22-73:
1. Replace `NumeAgent f[12]` + `PrenumeAgent f[13]` with a single `NumeSiPrenumeAgent string // [12]`; delete PrenumeAgent (it does not exist in the DLL output). Update cmd/wmecsv/main.go:123-141 accordingly.
2. Re-point: Scadenta f[13], Discount f[14], DenCritDiscount f[15], DenumiriSedii f[16] (a new field — the '~'-joined branch names the module currently throws into DenCritDiscount). CodExtern stays f[17]; delete the comment on types.go:24 ('PDF says SediiPartener here, but actual data = CodExtern') — the PDF is right about the order, it is the collapsed Nume/Prenume field at [12] that shifts everything, and SediiPartener really is present, at [16].
3. Rename Judet -> SimbolAutoJudet at [23] and keep the note that manual 25 misnames it 'Tara'; keep Tara at [42].
4. Name the tail: [37] FlagProducator, [40] SerieActIdentitate, [41] NumarActIdentitate, [42] Tara. Keep [38],[39],[43],[44],[45],[46],[47],[48] as explicitly-unknown named positions, do not delete them.

PARSING RULE (replaces splitFields's SplitN):
Use strings.Split (all parts), never SplitN. Then:
  - len(parts) < 49  -> return a loud error naming the record; do not pad. (splitFields currently pads silently, which would turn a truncated record into a plausible-looking partner.)
  - len(parts) == 49 -> map by index.
  - len(parts) > 49  -> the DLL has widened. Map [0..48] by index and keep the remainder in `Extra []string`. DO NOT MERGE. This is exactly the rule the reverted commit 1db37e8 got wrong: it merged on `len(parts) > count`, which cannot tell 'a field contains a semicolon' from 'the DLL grew a column', and once live width moved it re-glued the wrong fields on every row.
Embedded ';' must be detected by SHAPE, not by width. This reader has hard shape anchors: parts[18] in {DA,NU}, parts[30] in {PJ,PF}, parts[32] matches dd.mm.yyyy, parts[44] in {DA,NU}. If any anchor fails, a semicolon leaked into an earlier free-text field: reject the record with a LOG_ERR naming the partner (parts[0]) rather than guessing. Zero of 4866 records fail these anchors today, so this path is a genuine alarm, not a fallback.
Embedded NEWLINES are real (3 records) and harmless inside Go; guard only the dump/CSV writers (cmd/wmedump, cmd/wmecsv) which currently emit them raw and corrupt their own line structure.

### `GetListaPersonal`

10 positions in the April 2026 dump, 9 of them named by the vendor manual.

```
[0] NumeComplet -> string  (surname + given name in ONE field, e.g. "GURGU ION", "Moarcas Anda-Bianca"; manual's Nume+Prenume are NOT split by the DLL)
[1] Marca -> string (integer-valued employee number, runs 0..88 monotonically)
[2] CNP -> string (13-digit Romanian personal code, e.g. 1661114080041)
[3] EsteActiv -> string enum "Da"/"Nu" (89 Da / 1 Nu)
[4] EsteAgent -> string enum "Da"/"Nu" (1 Da / 89 Nu)
[5] SerieBuletin -> string (2-letter ID-card series: BV=52, ZV=25, KV=3, PH=3, CI/IF/SB/SZ/VS/XR=1 each, 1 empty)
[6] NumarBuletin -> string (6-digit ID-card number, e.g. 623323)
[7] CodPostal -> string (old-style RO postal code: 2200 x58, 2212 x3, 2300 x1, empty x28)
[8] UNNAMEABLE - alternate/alias name -> string (undocumented; 88/90 empty; only "ClaudiaBejan" row 50 and "Anda Hudescu" row 90. Both are person names differing from [0] - looks like an operator login or maiden/previous name. Not enough samples to name confidently; keep as named-unknown, e.g. Extra1)
[9] (trailing separator artifact) -> always empty string
```

**Change:** Split into 10 and re-map. Rename Employee.Nume -> NumeComplet and DROP Employee.Prenume (types.go:181-191) - the DLL has no such field; keeping it guarantees someone re-introduces the off-by-one. Add the [8] named-unknown. Concretely at queries.go:21, replace `splitFields(rec, 9)` with an unbounded `strings.Split(rec, ";")` plus a bounds-checked accessor, and assign NumeComplet=f[0], Marca=f[1], CNP=f[2], EsteActiv=f[3], EsteAgent=f[4], SerieBuletin=f[5], NumarBuletin=f[6], CodPostal=f[7], Extra1=f[8]. Fix the comment at queries.go:12, which currently repeats the manual's shifted layout and is what caused the bug.

### `GetListacarneteExt`

6 positions in the April 2026 dump, not documented by the vendor.

```
[0] SimbolCarnet -> string (may be EMPTY - 11/29 rows; e.g. "NIRMP", "BV AEG", "OF AEG", "")
[1] TipDoc -> string (document-type mnemonic: NIR, F, Cd, AE, NT, CF, M, BC, PV, NP, FF, FE, Dp, CH, AEG, Nir, AF, OFC, CC, TEMP, Temp, INVA, OF, AAA, BBB)
[2] Denumire -> string ("NIR Marfa - Piese", "Factura fiscala", "Bon consum", "Oferte client", ...)
[3] NrStart -> int (first number of the allocated range; STAYS FIXED as the book is consumed - see proof)
[4] NrStop -> int (last number of the allocated range)
[5] NumereRamase -> int (count of still-unused numbers in the range)
```

**Change:** Add a CarnetExt struct (SimbolCarnet, TipDoc, Denumire string; NrStart, NrStop, NumereRamase int) and parse with the unbounded split + accessor, 6 positions. Do NOT emit a trailing-empty field for this reader - unlike GetListaPersonal/Banci/Carnete/Gestiuni it has no trailing ';', which is itself a reason to stop hard-coding counts. Keep the numeric three as strings unless a caller needs arithmetic; if converting, note [0] and [1] are free-form and [0] is empty on 11/29 rows.

### `GetMonede`

2 positions in the April 2026 dump, not documented by the vendor.

```
[0] Denumire -> string (full currency name: "Lei", "Euro", "Forint", "Dolar", "Dolar canadian", "Lire englezesti", "Lei moldovenesti", "Nok", "Slot polonez", plus the sentinel "----nedefinit-----")
[1] Simbol -> string (short code: lei, Eur, HUF, USD, "$ can", Lire, LEIM, NOK, PLN, and "?" for the sentinel). NOTE: values are NOT ISO-4217-clean - "$ can" contains a space and a '$', "Lire" and "lei" are free text - so do not assume a 3-letter code
```

**Change:** Add a Moneda struct (Denumire, Simbol string) and parse 2 positions with the unbounded split + accessor. Put the reversed-order fact in the doc comment, since it contradicts every sibling reader. No trailing-empty position for this reader. Do not normalise or upper-case Simbol - "$ can" and "Lire" would be mangled.

### `GetNomenclatorArticole`

40 positions in the April 2026 dump, 24 of them named by the vendor manual.

```
[0]  -> CodExtern            -> string   (article code per SetIDArtField; CAN CONTAIN ';' — 46 rows)
[1]  -> Denumire             -> string
[2]  -> DenUM                -> string   (Buc:50999, Set:9277, L, Lei, Kg, M …)
[3]  -> PretVanzare          -> decimal RO (comma) — empty on 63433/63441 rows
[4]  -> SimbolClasa          -> string
[5]  -> DenClasa             -> string
[6]  -> CodProducator        -> string   (always empty at AEG)
[7]  -> DenProducator        -> string   (always empty at AEG)
[8]  -> GestImplicita        -> string   (warehouse symbol)
[9]  -> IDIntern             -> string/int — WinMENTOR internal article ID (manual calls it "CodExtern"; it is the identifier NOT used in [0])
[10] -> CotaTVA              -> int      (21/20/19/11/5/0)
[11] -> DenUMSecundara       -> string
[12] -> ParitateUMSecundara  -> decimal
[13] -> Masa                 -> decimal
[14] -> Serviciu             -> enum Da/Nu
[15] -> CodVamal             -> string (8-digit NC)
[16] -> PretMinim            -> decimal
[17] -> CantImplicita        -> decimal
[18] -> PretValuta           -> decimal (always empty)
[19] -> DataAdaug            -> date dd.mm.yyyy
[20] -> Masa2                -> decimal (always empty; manual duplicates "Masa")
[21] -> PretVCuTVA           -> decimal (= [3] x (1+CotaTVA/100), verified on all 8 priced rows)
[22] -> Locatie              -> string (always empty)
[23] -> PretReferinta        -> decimal (always empty)
[24] -> Flag24               -> enum DA/NU  (DA except 1 row) — UNNAMED
[25] -> Unknown25            -> string, always empty — UNNAMED
[26] -> CodSecundar          -> string; the one populated row repeats [0] — UNNAMED (barcode/alt code)
[27] -> Unknown27            -> int, always "0" — UNNAMED
[28] -> Unknown28            -> int-or-empty ("":36663 / "0":26778) — UNNAMED
[29] -> Flag29               -> enum, always "NU" — UNNAMED
[30] -> Unknown30            -> int-or-empty, locked to [32],[34] — UNNAMED
[31] -> Unknown31            -> string, always empty — UNNAMED
[32] -> Unknown32            -> int-or-empty, locked to [30],[34] — UNNAMED
[33] -> Unknown33            -> string, always empty — UNNAMED
[34] -> Unknown34            -> int-or-empty, locked to [30],[32] — UNNAMED
[35] -> Unknown35            -> string, always empty — UNNAMED
[36] -> Flag36               -> enum DA/NU (2 DA) — UNNAMED
[37] -> Unknown37            -> string, always empty — UNNAMED
[38] -> Unknown38            -> string, always empty — UNNAMED
[39] -> (record terminator)  -> always empty in the April dump
--- live only (42 tokens per commit b00482a; never sampled into a file) ---
[39] -> Unknown39            -> UNSAMPLED, added after 6 Apr 2026
[40] -> Unknown40            -> UNSAMPLED
[41] -> (record terminator)
```

**Change:** MISALIGNED + TRUNCATING. Two distinct defects.
(a) TRUNCATION: on live (42 tokens) `SplitN(rec, ";", 40)` makes f[39] = tokens 39+40+41 joined. Fields 0..38 stay correct, so it is bounded — this is why the revert was right.
(b) MISALIGNMENT: on the 46 rows whose CodExtern contains ';', every field shifts left by one. CodExtern becomes "35mmp  ", Denumire becomes " H01N2-D", DenUM becomes the real Denumire, PretVanzare becomes "M"/"Buc". That truncated code is what the sync writes to ref_ext, so those 46 products can never match.

DO NOT re-apply commit 1db37e8 (`splitFieldsMerging(rec, 40, 0)`): it merges whenever len(parts) > count, so with a live width of 42 it fired on EVERY record — CodExtern came back as "704.01;704.01 - Prestari servicii1;Lei", the unit vocabulary collapsed from 107 entries to 17 class names, and ref_ext integrity went from 40424 matches to 0 of 40831.

CORRECT FIX — never hard-code the total width:
1. Replace `strings.SplitN(record, ";", count)` with plain `strings.Split(record, ";")` + right-pad + ignore extras. This alone kills defect (a) and is forward-compatible with any future width; a live 42-token record then parses [0..38] exactly and drops the two unsampled tail tokens instead of jamming them into f[39].
2. Kill defect (b) with a BATCH-DERIVED width, not a constant: over the records returned by this one call, take the modal token count W; for a record with n > W tokens, rejoin tokens [codeIdx .. codeIdx+(n-W)] with ';' into the code field (codeIdx = 0 here). With April data W=40 and 46 rows get k=1; with live data W=42 and the same 46 rows get k=1. No constant is involved, so a width change cannot fire the merge on healthy rows.
3. Guard it with a per-record type assertion so a wrong W can never pass silently: after merging, tokens[codeIdx+10] must match ^[0-9]{1,2}$ (CotaTVA) and tokens[codeIdx+14] must be "Da"/"Nu". If the assertion fails, fail loudly (error, not a shrug) — this is exactly the check that would have caught 1db37e8 before it shipped.
4. Skip the merge entirely when the batch has fewer than 2 records (no mode is derivable).
5. Rename `CodExternUnic` -> `IDIntern` in types.go:268 and fix the comment; it is the internal ID, proved by the 9804/9804 join against GetStocArticole[18] and by GetInfoIesiriExt.
6. Extend the struct to Unknown39/Unknown40 (or better, keep a `Tail []string` for everything past [38]) so a wider live record is preserved rather than silently dropped.
7. Fix the stale comment at articles.go:134 ("the DLL actually returns 40") — it is 42 live.

### `GetOferte`

11 positions in the April 2026 dump, 14 of them named by the vendor manual.

```
[ 0] IDPart -> string, partner CodExtern = CIF (1708 distinct; empty on 45/86066)
[ 1] IDArticol -> string, article CodExtern (13022 distinct; empty on 36605/86066 = offers on a whole machine rather than a part). MAY CONTAIN ';'.
[ 2] DataInreg -> string date dd.mm.yyyy
[ 3] DataExpir -> string date dd.mm.yyyy (always >= [2]; typically +7 days)
[ 4] Pret -> RO decimal
[ 5] CantMin -> RO decimal ('1' x47883, '2' x14414, '4', '5')
[ 6] AdDim -> numeric, adaos/diminuare; '0' x23678 plus exactly three non-zero values in 86066 records ('7','-5','-14'); empty on 62383
[ 7] ProcDiscount -> numeric, '0' on 86066/86066
[ 8] Observatii -> free text up to 50 chars ('LINDE H40D seria H2X394U01652'); empty on 6394. MAY CONTAIN ';' AND embedded newlines.
[ 9] SimbolMoneda -> string, 'Lei' x85137 / 'Eur' x929
[10] CodlaFurn -> string, supplier's own code. Non-empty on only 6 of 86066 records, which is why it looks like a terminator — it is not. PROVEN by the 6 values, each of which is the same code as IDArticol[1] in the supplier's own formatting.
NO position [11] NrDoc, [12] Marca or [13] Agent: the manual names them but the DLL does not emit them (a 12th separator never appears except when Observatii carries one).
```

**Change:** TRUNCATING, and badly — the worst reader in this group. SplitN(rec,";",6) collapses positions 5..10 into one field, so Oferta.Cantitate currently receives, for the very first sample record, the literal string:
  '1;;0;TOYOTA 62-6FDF30 seria 606FDF30-20434 /1999;Lei;'
instead of '1'. Price is fine ([4] is before the cut); quantity and everything after are destroyed on 100% of 86168 records.

Fixes:
(a) Widen struct Oferta (types.go:240-248) from 6 to 11 fields, using the Rev.1.5 names for [0..10]: IDPart, IDArticol, DataInreg, DataExpir, Pret, CantMin, AdDim, ProcDiscount, Observatii, SimbolMoneda, CodlaFurn. Do NOT add NrDoc/Marca/Agent — the manual names them but the DLL never emits them; adding them would invite a future SplitN(14) that re-truncates.
(b) Replace splitFields(rec, 6) at articles.go:228 with:
    parts := strings.Split(rec, ";")            // full split, no N
    // [0..7] fixed head: IDPart, IDArticol, DataInreg, DataExpir, Pret, CantMin, AdDim, ProcDiscount
    //   validate the head by shape: parts[2] and parts[3] must both match ^\d{2}\.\d{2}\.\d{4}$; if not, skip the record and log LOG_ERR rather than storing shifted data
    // Observatii = strings.Join(parts[8:len(parts)-2], ";") when len(parts) >= 11, else parts[8]
    // SimbolMoneda = parts[len(parts)-2];  CodlaFurn = parts[len(parts)-1]
  This is the one reader in the group where a tail-relative merge is defensible, because IDArticol[1] — the join key — sits in the shape-validated fixed head and can never be touched by it; the only field the merge can corrupt is a free-text note. If the live DLL grows a 12th column, SimbolMoneda/CodlaFurn would shift by one, so ALSO shape-guard the tail: only accept parts[len-2] as SimbolMoneda if it matches ^(Lei|Eur|Usd|...)$, otherwise fall back to scanning forward from 8 for the currency token and treat everything after it positionally.
(c) IDArticol may contain ';' (46 known CodExtern values do) — the fixed head at [1] is therefore itself at risk when BOTH IDArticol and Observatii carry separators. The date shape-check on parts[2]/parts[3] detects exactly that case: if parts[2] is not a date, the code must scan forward for the first index k with parts[k] and parts[k+1] both dates, set IDArticol = Join(parts[1:k], ";") and continue from k. That is width-independent.
(d) Records may contain embedded newlines in Observatii (8 records in the sample); never split on \n.

### `GetReceptii`

21 positions in the April 2026 dump, not documented by the vendor.

```
[0] DenGestiune -> string. "Materiale consumab, auxiliare", "Marfuri = piese", "Marfuri = utilaje"
[1] SimbolGestiune -> string. M aux, Piese, Utilaje, OI
[2] NrNir -> string(int). 23935, 33051, 132
[3] DataNir -> string d.m.yyyy NOT zero-padded ("3.4.2026")
[4] NumeFurnizor -> string
[5] IDFurnizor -> string CodFiscal. RO8119423, DE321256424
[6] NrFacturaFurnizor -> string. VERIFIED == GetIntrari[2]
[7] DataFactura -> string dd.mm.yyyy ZERO-PADDED (note the contrast with [3])
[8] CodExtern -> string; non-empty 14/25. MAY CONTAIN ';' in principle (none in this sample)
[9] DenArticol -> string. VERIFIED == GetNomenclatorArticole[1]
[10] ContStoc -> string. 3028.02, 371.01, 371.02, 8039. VERIFIED == GetStocuriPeGestiuni[4]
[11] DenUM -> string. VERIFIED == GetNomenclatorArticole[2]
[12] Cant -> string RO-decimal
[13] Cant (repeat of [12]) -> string; identical to [12] on all 25 rows
[14] PretFactura -> string RO-decimal. VERIFIED == GetIntrari[6]
[15] PretInregistrare -> string RO-decimal, landed/stock cost. VERIFIED == GetStocuriPeGestiuni[8]
[16] ProcTVA -> string(int). 21 (24 rows) / 11 (1 row). VERIFIED == GetNomenclatorArticole[10]
[17] IDDocIntern -> string(int). per-NIR constant: NIR 23935 -> 53306 on all 8 of its lines; 33051 -> 53291; 33050 -> 53310
[18] <unnamed constant> -> string, "1" on 25/25
[19] NrLinie -> string(int). NIR 23935's 9 lines carry 9,1,2,3,4,8,6,7,5 - unique within the document
[20] IDArticolIntern -> string. VERIFIED == GetNomenclatorArticole[9]; rows arrive sorted ascending by it
[21] <trailing separator artifact> -> always ""; NOT a field
```

**Change:** Add a Receptie struct with the 21 content fields. It is worth doing: this is the only reader that carries NIR number, supplier, invoice number, landed cost AND the internal article id on one line, so it is the natural source for a WME->Dolibarr goods-receipt/stock-in feed. Parse with strings.Split + bounds-safe accessor; [8] CodExtern is the variable-width field ([7] before it is a zero-padded date and [10] after it is an account code, giving a clean two-sided anchor). Keep [18] as a named-unknown rather than dropping it. Re-dump with SetFiltruDocNeoperate(0) first, as for GetTransferuri, so the layout is validated on more than 25 rows and more than four warehouses.

### `GetSoldDetaliat`

no sample dump, 4 of them named by the vendor manual.

```
UNVERIFIED — no sample exists. Manual names for the documented prefix; predicted positions marked PREDICTED.

=== parts[0] == "Factura" ===
[ 0] Tip             -> const "Factura"    (manual line 248)
[ 1] NrFactura       -> int-as-string      (manual 248)
[ 2] DataFactura     -> date dd.mm.yyyy    (manual 248)
[ 3] RestDeIncasat   -> decimal RO         (manual 248 'Rest_de_Incasat')
[ 4] TermenDePlata      -> date dd.mm.yyyy  PREDICTED from GetSolduriExt[5]
[ 5] LocatiePartener    -> string           PREDICTED from GetSolduriExt[6]
[ 6] MarcaAgent         -> int-as-string    PREDICTED from GetSolduriExt[7]
[ 7] ValoareFactura     -> decimal RO       PREDICTED from GetSolduriExt[8]
[ 8] ObservatiiFactura  -> free text        PREDICTED from GetSolduriExt[9]
[ 9] SerieDocument      -> string           PREDICTED from GetSolduriExt[10]
[10] TipDocument        -> int-as-string    PREDICTED from GetSolduriExt[11]
[11] PrefixTipCarnet    -> string           PREDICTED from GetSolduriExt[12]

=== parts[0] == "Avans" ===
[ 0] Tip                   -> const "Avans"                       (manual line 250)
[ 1] DocumentAvans         -> string, type+number in ONE field    (manual 250; corroborated by live GetSolduriExt/Furn 'EX 12', 'AVANS 0')
[ 2] DataDocument          -> date dd.mm.yyyy                     (manual 250)
[ 3] SumaAvans             -> decimal RO, ALWAYS NEGATIVE         (manual 250 + 252; sign confirmed on 129 live Avans rows)
[ 4] MarcaAgentIncasator   -> int-as-string   PREDICTED from GetSolduriExt[5]
[ 5] ValoareInitialaAvans  -> decimal RO      PREDICTED from GetSolduriExt[6]
```

**Change:** NO-DATA — cannot be pronounced correct or broken, but the current decoder is wrong on a point the MANUAL ALONE already settles, so one change is safe to make today without a dump.

1. THE TAG IS IGNORED. partners.go:97-114 applies splitFields(rec, 4) to every line regardless of parts[0]. The manual states plainly (lines 251-252) that 'Factura' and 'Avans' lines have different structures, and that on an Avans line position [1] is DocumentAvans and [3] is a negative SumaAvans. DetailedBalance's field names (NrDocument, DataDocument, Rest) happen to survive for Avans, but nothing in the code distinguishes the two, so a caller cannot tell an unpaid invoice from an advance except by string-matching Type. Add the discriminator and, once measured, the two widths.
2. SplitN(4) is the corrupting shape. If the true width is 12 (the prediction), DetailedBalance.Rest becomes 'Rest;Termen;Locatie;Marca;Valoare;Observatii;Serie;19;F.BV AEG' on every invoice line — the same failure now confirmed on 420/420 GetSolduriExt rows. Replace with strings.Split, map [0..3] by index, require len(parts) >= 4 (loud error below that), and keep everything from [4] on in Extra []string. That is correct whether the true width is 4 or 12 and needs no dump to be safe.
3. Shape anchors for a leaked ';': parts[0] must be exactly 'Factura' or 'Avans', parts[2] must match dd.mm.yyyy, parts[3] must parse as an RO decimal (and be negative when parts[0]=='Avans'). Reject the record loudly if any fails; do not merge on width.
4. MEASURE IT. Add GetSoldDetaliat to cmd/wmedump/main.go's sweep with a PartID that is known to have both unpaid invoices and advances — 'RO23315296' has 5 Avans rows in GetSolduri.txt, and 'RO14619270' has 14 Factura rows — so one run yields both variants. Until then, treat the positions above [3] as predicted, not established.

### `GetSoldPart`

no sample dump, 3 of them named by the vendor manual.

```
UNVERIFIED — no sample exists. Documented prefix only; positions beyond [2] are unknown and MUST be tolerated.

[ 0] CodExtern -> string  (manual line 236)
[ 1] Denumire  -> string  (manual line 236)
[ 2] Sold      -> decimal RO  (manual line 236)
[3+] UNKNOWN — not measurable from any artefact in this repo. Predicted by analogy with GetSolduri[9]/[10]: SimbolMoneda then CursValutar. DO NOT hard-code 3 as a closed width.
```

**Change:** NO-DATA — cannot be pronounced correct or broken. splitFields(records[0], 3) at partners.go:88 is consistent with the manual, but the manual has been wrong about width for 5 of the 6 readers in this group that could be checked, so 3 is an unverified assumption. And SplitN(3) is the specific shape that silently corrupts: if the DLL returns 4+ fields, SoldPartener.Sold becomes 'Sold;Lei;1' and every downstream RO-decimal parse of it fails or, worse, succeeds on a truncated prefix.

CHANGES:
1. In partners.go:88 replace splitFields with strings.Split and keep an Extra []string on SoldPartener (types.go:174-178). Map [0],[1],[2] by index; require len(parts) >= 3 and return a loud error below that; put everything from [3] on into Extra rather than into Sold. This is width-tolerant without a fixed total.
2. Add a shape guard on parts[2]: it must parse as an RO decimal (helpers_format.go parsers). That single anchor detects both a widened DLL and a ';' inside Denumire, and it costs nothing.
3. Fix the dropped-records bug at partners.go:85-88: len(records) > 1 is currently ignored. Either return all lines or fail loudly when more than one arrives; do not silently take records[0].
4. MEASURE IT. This is the actionable item: add GetSoldPart to cmd/wmedump/main.go's sweep (it needs a PartID — feed it a CodExtern taken from GetListaParteneri, e.g. 'RO14619270', the partner with the most GetSolduri rows) and re-run on the Windows box. Until a dump exists, no struct here should be treated as settled.

### `GetSolduri`

11 positions in the April 2026 dump, 6 of them named by the vendor manual.

```
TAGGED UNION on position [1]. Both variants share [0..1].

=== parts[1] == "Factura"  (11 fields, 420/467 rows) ===
[ 0] IDPartener        -> string, partner CodExtern, may contain a space  (manual: NOT DOCUMENTED — the manual claims NrDoc here)
[ 1] Tip               -> const "Factura"                                 (manual: NOT DOCUMENTED)
[ 2] NrFactura         -> int-as-string   (manual 1 'NrDoc')
[ 3] DataFactura       -> date dd.mm.yyyy (manual 2 'DataDoc')
[ 4] RestDePlata       -> decimal RO      (manual 3 'Rest')
[ 5] TermenDePlata     -> date dd.mm.yyyy (manual 4 'Termen')
[ 6] LocatiePartener   -> string, branch name (manual: NOT DOCUMENTED for this function; documented for GetSolduriExt at line 485)
[ 7] MarcaAgent        -> int-as-string   (manual 5 'Agent', but at the wrong position)
[ 8] ValoareFactura    -> decimal RO      (manual 6 'Valoare')
[ 9] SimbolMoneda      -> string Lei/Eur  (UNDOCUMENTED, evidence-named)
[10] CursValutar       -> decimal RO, 1 for Lei (UNDOCUMENTED, evidence-named)

=== parts[1] == "Avans"  (9 fields, 47/467 rows) ===
[ 0] IDPartener            -> string, partner CodExtern
[ 1] Tip                   -> const "Avans"           (manual 1 'Doc')
[ 2] DocumentAvans         -> string, TYPE+NUMBER IN ONE FIELD, e.g. 'EX 12' (manual 2 'NrDoc'; GetSolduriExt's manual entry wrongly splits this into 'Tip Document plata avans' + 'Nr Document plata avans')
[ 3] DataDocument          -> date dd.mm.yyyy         (manual 3 'DataDoc')
[ 4] RestAvans             -> decimal RO, ALWAYS NEGATIVE (manual 4 'Suma')
[ 5] MarcaAgentIncasator   -> int-as-string; empty in all 47 rows (manual 5 'Agent')
[ 6] ValoareInitialaAvans  -> decimal RO, positive     (manual 6 'Valoare')
[ 7] SimbolMoneda          -> string Lei               (UNDOCUMENTED, evidence-named)
[ 8] CursValutar           -> decimal RO, '1'          (UNDOCUMENTED, evidence-named)
```

**Change:** MISALIGNED, and worse than a shift — struct Sold has no partner key and no record-type tag at all, so every field lands one-to-two positions early and the last one eats the rest of the line. Today, for BE0445945028;Factura;12;29.06.2012;7000;31.07.2012;;;7000;Eur;4,4508, splitFields(rec, 6) at queries.go:179 produces:
  Sold.NrDoc   = 'BE0445945028'        (the partner code)
  Sold.DataDoc = 'Factura'             (the tag)
  Sold.Rest    = '12'                  (the invoice number)
  Sold.Termen  = '29.06.2012'          (the invoice date)
  Sold.Agent   = '7000'                (the amount still owed)
  Sold.Valoare = '31.07.2012;;;7000;Eur;4,4508'   (SplitN(6) glues the last six fields together)
Not one field is right. This is currently LATENT — grep shows GetSolduri is only reached through cmd/wmedump/main.go:93 (RawQuery, unparsed); internal/sync never calls it — but it must not ship as-is.

FIX, in types.go:405-412 and queries.go:171-190:
1. Replace struct Sold with a tagged type:
   type Sold struct { IDPartener, Tip, NrDocument, DataDocument, Rest, TermenDePlata, LocatiePartener, MarcaAgent, Valoare, SimbolMoneda, CursValutar string }
   For Avans rows, NrDocument holds the merged 'EX 12', TermenDePlata is unset, and LocatiePartener is unset — or, cleaner, define Sold with a Tip discriminator and two named decoders. Do not invent a Termen for advances; the DLL does not send one.
2. Dispatch on parts[1] BEFORE choosing a width. This is the only correct way to read this reader:
   parts := strings.Split(rec, ";")
   switch { case len(parts) >= 2 && parts[1] == "Factura": want = 11
            case len(parts) >= 2 && parts[1] == "Avans":   want = 9
            default: return loud error naming rec }
   len(parts) < want -> loud error (splitFields currently pads, which would fabricate an empty Moneda and Curs on a truncated row). len(parts) == want -> map. len(parts) > want -> map [0..want-1] and keep Extra []string; NEVER merge.
3. Embedded ';' policy: the shape anchors here are strong and cheap — parts[1] is a closed enum, parts[3] is dd.mm.yyyy, and for Factura parts[5] is dd.mm.yyyy and parts[len-1] is a decimal. A ';' inside the only free-text field ([6] LocatiePartener, values like 'Pct. de lucru') would break the parts[len-1] and parts[3] anchors; reject with LOG_ERR naming parts[0] instead of guessing. No record in the 467-row sample fails, so this is an alarm, not a fallback.
4. Add the doc comment that GetSolduri and GetSolduriExt return the SAME records in the SAME order with a common 9-field (Factura) / 7-field (Avans) prefix — verified 467/467 — so the two can be zipped rather than joined.

### `GetSolduriExt`

13 positions in the April 2026 dump, 10 of them named by the vendor manual.

```
TAGGED UNION on position [1]. Widths 13 (Factura) and 7 (Avans) are constant within each tag.

=== parts[1] == "Factura"  (13 fields, 420/467 rows) ===
[ 0] IDPartener         -> string, partner CodExtern, may contain a space (manual line 479 'ID Partener')
[ 1] Tip                -> const "Factura"                                (manual 480)
[ 2] NrFactura          -> int-as-string                                   (manual 481)
[ 3] DataFactura        -> date dd.mm.yyyy                                 (manual 482)
[ 4] RestDePlata        -> decimal RO                                      (manual 483)
[ 5] TermenDePlata      -> date dd.mm.yyyy                                 (manual 484)
[ 6] LocatiePartener    -> string, branch name                             (manual 485)
[ 7] MarcaAgent         -> int-as-string                                   (manual 486)
[ 8] ValoareFactura     -> decimal RO                                      (manual 487)
[ 9] ObservatiiFactura  -> free text; MAY CONTAIN NUL BYTES (2 in the sample) (manual 488)
[10] SerieDocument      -> string, booklet symbol 'BV AEG'/'INVAEG'/'EX AE' (UNDOCUMENTED, evidence-named)
[11] TipDocument        -> int-as-string, 19/22; same code domain as GetInfoIesiri[0] 'TipDoc' (UNDOCUMENTED, evidence-named)
[12] PrefixTipCarnet    -> string, 'F.'+[10] or 'Iv.'+[10]                  (UNDOCUMENTED, evidence-named)

=== parts[1] == "Avans"  (7 fields, 47/467 rows) ===
[ 0] IDPartener            -> string, partner CodExtern            (manual 491)
[ 1] Tip                   -> const "Avans"                        (manual 492)
[ 2] DocumentAvans         -> string, TYPE AND NUMBER MERGED ('EX 12'). THE MANUAL SPLITS THIS INTO TWO FIELDS (lines 493 'Tip Document plata avans' + 494 'Nr Document plata avans'); THE DLL DOES NOT. This is why the manual says 8 and the wire says 7.
[ 3] DataDocument          -> date dd.mm.yyyy                      (manual 495)
[ 4] RestAvans             -> decimal RO, ALWAYS NEGATIVE          (manual 496)
[ 5] MarcaAgentIncasator   -> int-as-string; empty in all 47 rows  (manual 497)
[ 6] ValoareInitialaAvans  -> decimal RO, positive                 (manual 498)
```

**Change:** MISALIGNED and TRUNCATING at the same time. partners.go:270 does splitFields(rec, 10) for both variants:
  - Factura rows (13 fields): SplitN(10) makes SoldExt.ObservatiiFactura = 'Observatii;BV AEG;19;F.BV AEG'. The invoice note is silently glued to three unrelated document fields on all 420 rows — the exact failure mode the CLAUDE.md gotcha warns about ('SplitN with too-small N silently corrupts the last field').
  - Avans rows (7 fields): splitFields pads to 10 and then MISNAMES everything: 'EX 12' lands in NrFactura, the empty MarcaAgentIncasator lands in TermenDePlata, and the advance's original amount lands in LocatiePartener. SoldExt.MarcaAgent, ValoareFactura and ObservatiiFactura are fabricated empties.
Latent today (only cmd/wmedump/main.go:94 reaches it, via RawQuery), but wrong on 467/467 rows the moment anything consumes it.

FIX, in types.go:377-388 and partners.go:262-285:
1. Split SoldExt into the two real shapes, or keep one struct with a Tip discriminator plus the three tail fields: add SerieDocument, TipDocument, PrefixTipCarnet; add DocumentAvans / ValoareInitialaAvans (or document that NrFactura carries the merged 'EX 12' and Valoare carries the initial advance for Avans rows). Do not leave the Avans amount in LocatiePartener.
2. Dispatch on parts[1] before choosing a width: 'Factura' -> 13, 'Avans' -> 7, anything else -> loud error. Use strings.Split, not SplitN. len(parts) < want -> loud error, no padding. len(parts) > want -> map [0..want-1], keep Extra []string, NEVER merge on width.
3. Embedded-';' handling WITHOUT assuming a total width: [9] Observatii is the only free-text field, and it sits between two hard anchors — a decimal at [8] and a fixed 3-field document tail at the end. So for a Factura row, take [0..8] from the LEFT, take the last 3 parts from the RIGHT, and join everything in between with ';' into Observatii, but ONLY after verifying the right-hand anchors actually look like a document tail (parts[len-2] is all digits and parts[len-1] contains parts[len-3]). If those anchors fail, the DLL has widened rather than a semicolon having leaked, and the record must be reported, not merged. That anchor test is precisely what the reverted commit 1db37e8 lacked: it merged on `len(parts) > count` alone, so when live width grew it re-glued every record.
4. NUL bytes are real in this reader's Observatii (2 in 467 rows). Go keeps them; strip or reject them before they reach MySQL/JSON/Dolibarr REST, and note in the code that awk-based inspection of these dumps under-counts fields because of them.

### `GetSolduriFurn`

13 positions in the April 2026 dump, 10 of them named by the vendor manual.

```
TAGGED UNION on position [1]. Positionally identical to GetSolduriExt; the tail semantics differ because these are incoming documents.

=== parts[1] == "Factura"  (13 fields, 404/486 rows) ===
[ 0] IDFurnizor            -> string, supplier CodExtern; CAN BE EMPTY (1 row)      (manual by reference, line 479)
[ 1] Tip                   -> const "Factura"                                        (manual 480)
[ 2] NrFacturaFurnizor     -> string, the supplier's invoice number                   (manual 481)
[ 3] DataFactura           -> date dd.mm.yyyy                                         (manual 482)
[ 4] RestDePlata           -> decimal RO                                              (manual 483)
[ 5] TermenDePlata         -> date dd.mm.yyyy                                         (manual 484)
[ 6] LocatiePartener       -> string; EMPTY in all 404 rows                           (manual 485)
[ 7] MarcaAgent            -> int-as-string; EMPTY in all 404 rows                    (manual 486)
[ 8] ValoareFactura        -> decimal RO                                              (manual 487)
[ 9] ObservatiiFactura     -> free text; EMPTY in all 404 rows                        (manual 488)
[10] SerieDocumentFurnizor -> string, the supplier's own invoice series, free text     (UNDOCUMENTED, evidence-named)
[11] TipDocument           -> int-as-string, 1/3/5; incoming-document type code, same slot as GetSolduriExt[11] but a disjoint value set (UNDOCUMENTED, evidence-named)
[12] PrefixTipCarnet       -> string; EMPTY in all 404 rows (no internal AEG carnet for incoming docs) (UNDOCUMENTED, evidence-named)

=== parts[1] == "Avans"  (7 fields, 82/486 rows) ===
[ 0] IDFurnizor            -> string, supplier CodExtern                (manual 491)
[ 1] Tip                   -> const "Avans"                             (manual 492)
[ 2] DocumentAvans         -> string, TYPE AND NUMBER MERGED: 'EX 10', 'AVANS 0' (manual 493+494 COLLAPSED INTO ONE FIELD)
[ 3] DataDocument          -> date dd.mm.yyyy; '30.12.1899' appears (Delphi zero date) (manual 495)
[ 4] RestAvans             -> decimal RO, ALWAYS NEGATIVE               (manual 496)
[ 5] MarcaAgentIncasator   -> int-as-string; EMPTY in all 82 rows       (manual 497)
[ 6] ValoareInitialaAvans  -> decimal RO, positive                      (manual 498)
```

**Change:** MISALIGNED for the same reason as GetSolduriExt, and it shares the bug because partners.go:288-311 is a copy-paste of partners.go:262-285 down to the splitFields(rec, 10).
  - Factura rows (13 fields): SplitN(10) makes SoldExt.ObservatiiFactura = ';MENT;1;' — i.e. the always-empty Observatii is replaced by a glued string of the supplier series, the document-type code and an empty prefix, on all 404 rows.
  - Avans rows (7 fields): padded to 10 and misnamed exactly as in GetSolduriExt — 'EX 10' into NrFactura, the empty MarcaAgentIncasator into TermenDePlata, and the advance amount into LocatiePartener.
Latent today (cmd/wmedump/main.go:95 only).

FIX, in types.go:377-388 and partners.go:288-311:
1. Apply the same tagged-union decoder as GetSolduriExt: dispatch on parts[1] ('Factura' -> 13, 'Avans' -> 7, else loud error), strings.Split not SplitN, len<want -> loud error, len>want -> Extra []string, never merge on width.
2. Factor the decoder out and call it from BOTH GetSolduriExt and GetSolduriFurn instead of duplicating the body — the manual itself defines Furn as 'aceeasi structura ca si functia GetSolduriExt' (lines 694-695) and the 486-row sample confirms the widths are identical. Keep the naming difference visible in the doc comment: [10] is our booklet symbol for clients and the SUPPLIER'S series for suppliers, and [11]'s value domain is 19/22 for clients vs 1/3/5 for suppliers.
3. Do NOT reuse the GetSolduriExt embedded-';' anchor here unchanged: on this reader [9] Observatii is empty in every row and [10] is the free-text field, so the safe anchor is 'parts[len-1] is empty and parts[len-2] is all digits'. If that fails, report the record; do not merge.
4. Handle the empty IDFurnizor at [0] (1 of 404 rows) explicitly — an empty partner key must be surfaced, not silently used as a join key.

### `GetStocArticol`

no sample dump, 5 of them named by the vendor manual.

```
[0] -> CodExtern    -> string (documented; CAN CONTAIN ';' — the same 46/12/15 codes exist here)
[1] -> Denumire     -> string (documented)
[2] -> UM           -> string (documented)
[3] -> PretVanzare  -> decimal (documented; by analogy with GetStocArticole almost certainly "0")
[4] -> Stoc         -> decimal (documented)
[5..] -> UNKNOWN, UNSAMPLED — the manual's prefix is 5 fields and no dump exists to say how many more the DLL emits. Given GetStocArticole widens 17 -> 21, treat everything past [4] as present-but-unnamed until dumped.
```

**Change:** NO-DATA — cannot be certified. The Go code is not provably wrong, but `splitFields(records[0], 5)` has the exact failure mode that destroyed GetNomenclatorArticole in miniature: f[4] (Stoc) is a SplitN remainder, so if the DLL returns more than 5 fields — which every sibling reader does — Stoc silently becomes "1;Imobillizari;Imobilizari;;;RO11330527;…" and every numeric parse of it fails or, worse, succeeds on the leading digits.

Changes, in order:
1. BLOCKING: dump it. Add GetStocArticol to cmd/wmedump/main.go's reader list (it needs an ArticolID + GestID pair — feed it a code and warehouse taken from GetStocArticole, e.g. "4068400000729" / "EuroDiscount") and re-run on the Windows box. Do not change the struct before that dump exists. This is the same discipline the b00482a revert demanded: the true count comes from live, never from a stale dump and never from the manual.
2. Regardless of the dump, switch to `strings.Split` + pad so f[4] can never absorb a tail. That is a strict improvement available today and cannot regress anything.
3. The single-record shape means the batch-modal-width trick is unavailable: with one record there is no mode. Document explicitly that a ';' inside this article's own CodExtern is unresolvable here, and that callers must therefore look the article up in GetStocArticole/GetNomenclatorArticole (which return batches) when the code is suspect.
4. Leave the manual's 5 names for [0..4] — they align with GetStocArticole[0..4], whose alignment is proven.

### `GetStocArticole`

21 positions in the April 2026 dump, 17 of them named by the vendor manual.

```
[0]  -> CodExtern       -> string   (CAN CONTAIN ';' — 12 rows)
[1]  -> Denumire        -> string
[2]  -> UM              -> string
[3]  -> PretVanzare     -> decimal — "0" on every row; unusable, use GetNomenclatorArticole
[4]  -> Stoc            -> decimal RO (comma)
[5]  -> SimbolClasa     -> string
[6]  -> DenClasa        -> string
[7]  -> IDProducator    -> string (always empty at AEG)
[8]  -> DenProducator   -> string (always empty at AEG)
[9]  -> IDFurnizor      -> string — CUI, per SetIDPartField("CodFiscal")
[10] -> DenFurnizor     -> string
[11] -> SimbolGestiune  -> string
[12] -> DenGestiune     -> string
[13] -> CotaTVA         -> int
[14] -> Flag            -> enum D/N — "NU" on every row (VAT included in sale price?)
[15] -> PretCuTVA       -> decimal — "0" on every row
[16] -> StocRezervat    -> int (reserved by orders; the manual's last documented field)
[17] -> Unknown17       -> int, "0" on every row — UNNAMED (a second reservation counter?)
[18] -> IDIntern        -> string/int — WinMENTOR internal article ID, == GetNomenclatorArticole[9] on 9804/9804 rows
[19] -> Unknown19       -> string, empty on every row — UNNAMED
[20] -> (record terminator) -> always empty
```

**Change:** MISALIGNED on the 12 rows whose CodExtern contains ';' (fields shift left by one: Denumire becomes the code tail, UM becomes the name, PretVanzare becomes the unit, and IDIntern lands in f[19] where the struct expects Unknown19 — so those 12 articles lose BOTH join keys at once). Correct on the other 9804.

The documented prefix 0..16 is exactly right; the extension to 21 in the Go struct is confirmed by the data and the field names for 17/19 are the only genuine gaps.

Changes:
1. Same helper fix as GetNomenclatorArticole — `strings.Split` + pad instead of `SplitN(rec, 21)`, so a wider live record does not collapse into f[20]. Nothing here proves the live width is still 21; assume it can grow.
2. Apply the batch-modal-width merge with codeIdx = 0, guarded by a type assertion: tokens[codeIdx+13] must match ^[0-9]{1,2}$ (CotaTVA) and tokens[codeIdx+14] must be D/N/DA/NU. On the April dump that repairs all 12 rows and leaves all 9804 others byte-identical.
3. Keep Unknown17/Unknown19 as named-unknown; do not invent names for them.
4. articles.go:4 comment ("PDF documents 17 fields but the DLL actually returns 21") is accurate for the April dump — reword to say the 21st token is the record terminator and that the count must be re-verified against live before being relied on.

### `GetStocArticoleExt`

20 positions in the April 2026 dump, not documented by the vendor.

```
[0]  -> CodExtern       -> string  (CAN CONTAIN ';' — 12 rows)
[1]  -> Denumire        -> string
[2]  -> UM              -> string
[3]  -> PretVanzare     -> decimal — "0" on every row
[4]  -> Stoc            -> decimal RO
[5]  -> SimbolClasa     -> string
[6]  -> DenClasa        -> string
[7]  -> IDProducator    -> string (always empty)
[8]  -> DenProducator   -> string (always empty)
[9]  -> IDFurnizor      -> string (CUI)
[10] -> DenFurnizor     -> string
[11] -> SimbolGestiune  -> string
[12] -> DenGestiune     -> string
[13] -> CotaTVA         -> int
[14] -> Flag            -> enum D/N — "NU" on every row
[15] -> PretCuTVA       -> decimal — "0" on every row
   (positions 0..15 verified byte-identical to GetStocArticole[0..15] on all 9816 lines)
[16] -> Unknown16       -> string, empty on every row — UNNAMED (this is where GetStocArticole has StocRezervat; here it is NOT the same field)
[17] -> DataUnknown17   -> date dd.mm.yyyy, "30.12.1899" (Delphi NULL) on every row — UNNAMED (lot/expiry/last-movement date are all plausible; nothing in the data distinguishes them)
[18] -> Unknown18       -> numeric, empty on all 9804 real rows, "0" in the Ext2 "---nedefinit---" placeholder — UNNAMED
[19] -> (record terminator) -> always empty

GetStocArticoleExt2(GestID) -> IDENTICAL layout, same 20 tokens, restricted to one warehouse, plus one "---nedefinit---" placeholder row for warehouses holding no stock.
```

**Change:** UNPARSED — and it should stay that way. There is no struct, no SplitN and therefore no truncation and no misalignment bug: the Go code hands the caller raw records. That is why this reader is the only one of the five with nothing broken.

Do NOT add a typed reader for it. The evidence says GetStocArticoleExt is a downgrade from GetStocArticole for this dataset: identical fields [0..15], but it drops StocRezervat ([16]) and the internal article ID ([18]) — the join key that matches nomenclator[9] on 9804/9804 rows — in exchange for one always-NULL date. Anything the sync needs, GetStocArticole already gives it with more information.

If a typed reader is added later anyway:
1. Reuse the StockArticle field names for [0..15]; they are proven identical byte-for-byte.
2. Do NOT map [16] onto StocRezervat — it is a different field that happens to sit in the same slot.
3. Leave [16]/[17]/[18] as named-unknown. The corpus has literally zero variance in them, so naming them would be invention.
4. Use `strings.Split` + pad + the batch-modal-width merge at codeIdx = 0, with the guard tokens[codeIdx+13] ~ ^[0-9]{1,2}$ and tokens[codeIdx+14] in {D,N,DA,NU} — identical to GetStocArticole, which repairs the same 12 codes.
5. GetStocArticoleExt2 can share the exact same parser; only the placeholder "---nedefinit---" row needs skipping (CodExtern empty and Denumire == "---nedefinit---").

### `GetStocuriPeGestiuni`

10 positions in the April 2026 dump, not documented by the vendor.

```
[0] -> DenGestiune        -> string  (warehouse NAME; GetListaGestiuni stores it as Simbol;Den, i.e. reversed)
[1] -> SimbolGestiune     -> string
[2] -> DenumireArticol    -> string
[3] -> CodExtern          -> string  (CAN CONTAIN ';' — 15 rows; note this is position 3, mid-record)
[4] -> ContContabil       -> string  (371.01 / 371.03 / 371.04 / 2131 / 2133 / 214)
[5] -> UM                 -> string
[6] -> Stoc               -> decimal RO (comma)
[7] -> PretAchizitie      -> decimal, UNIT price, rounded to 2dp  [currently misnamed ValoareStoc]
[8] -> PretInregistrare   -> decimal, UNIT price the warehouse is booked at, full precision — equals the precise cost in cost-valued warehouses, the SALE price in the 371.04 retail warehouse  [currently misnamed ValoareStocPrecisa]
[9] -> CotaTVA            -> int
(no terminator token — this reader emits no trailing ';')
```

**Change:** MISALIGNED on 15 rows, and SEMANTICALLY WRONG on all 11036.

(a) Misalignment: the extra ';' sits at index 3, so on those 15 rows CodExtern truncates to "35mmp  " and ContContabil/UM/Stoc/prices/CotaTVA all shift left by one — CotaTVA ends up merged into f[9] by SplitN. Those 15 rows feed BuildCostPriceMap with a bogus code and a bogus number.

(b) Semantics — the bigger problem: rename [7] ValoareStoc -> PretAchizitie and [8] ValoareStocPrecisa -> PretInregistrare (types.go:168-169), fix the struct comment at types.go:159, and then fix internal/sync/products.go:143-168 to STOP dividing by Stoc. The correct unit cost is [7] (or [8] where full precision is wanted); `val / stoc` corrupts 4875 of 8977 articles' cost prices and zeroes 162. Fix the two comments that repeat the wrong formula at products.go:108 and products.go:144 as well.

(c) Parsing: same helper change (`strings.Split` + pad, not SplitN(10)) plus a batch-modal-width merge with codeIdx = 3. Because the record has no trailing ';' and the tail after the code is strongly typed, the guard here is cheap and exact: after merging, the LAST token must match ^[0-9]{1,2}$ (CotaTVA) and tokens[codeIdx+1] must match ^[0-9]{3}(\.[0-9]+)?$ (ContContabil). On the April dump that repairs all 15 rows and leaves the other 11021 byte-identical. An equivalent width-independent option for this reader specifically: take tokens [0..2] from the left and the last 6 from the right, and join everything in between as CodExtern — that needs no batch statistics at all and is my preferred fix here.

### `GetTransferuri`

14 positions in the April 2026 dump, not documented by the vendor.

```
[0] SimbolGestSursa -> string. "Piese" - a GetListaGestiuni symbol
[1] SimbolGestDest -> string. "F" (Firma) - a GetListaGestiuni symbol
[2] NrDoc -> string(int). 400,401,402,403,404 - one nota transfer per month, Feb..Jun 2026
[3] Data -> string d.m.yyyy NOT zero-padded ("30.6.2026", "2.4.2026") - differs from GetIntrari/GetInfoIesiri which are zero-padded
[4] CodExtern -> string. OC1051, N1323064, "Q8 HAYDN 46 - 20L", "058/233003688071-6069676". MAY CONTAIN ';' (this article family is exactly where the embedded-';' codes live)
[5] DenArticol -> string. "Filtru ulei", "Pompa cu filtru combustibil"
[6] Cant -> string RO-decimal. 1, 3,5, 24, 10
[7] Pret -> string RO-decimal. 13,21 / 371,105883 / 632,9868778
[8] Pret (repeat of [7]) -> string; identical to [7] on all 30 rows
[9] ContSursa -> string. 371.01
[10] ContDest -> string. 371.01
[11] <unnamed> -> string; empty on all 30 rows
[12] ValoareDocument -> string RO-decimal; per-DOCUMENT constant
[13] <unnamed> -> string; empty on all 30 rows
[14] <trailing separator artifact> -> always ""; NOT a field
```

**Change:** Before writing any struct: re-run wmedump for GetTransferuri WITH SetFiltruDocNeoperate(0) and capture a real dump into wme-raw/ - the existing 0-byte file is a dumper artifact, not evidence of an empty reader, and it means 6 months of transfer data (374 rows) has never been in the repo. Only then add a Transfer struct with the 14 fields above. When you do, parse with strings.Split and treat [4] CodExtern as the variable-width field (anchor: [6] Cant must be an RO-decimal and [3] must be a d.m.yyyy date), for the same reason as GetIntrari - the parts catalogue is the one that contains ';' inside CodExtern.

### `GetVanzariExt`

23 positions in the April 2026 dump, 18 of them named by the vendor manual.

```
[ 0] PartID -> string (partner CodExtern = CIF, e.g. 'RO4077201'; empty on 16/10020)
[ 1] Zi -> int, 1..31 (31 distinct)
[ 2] NrDoc -> int, 5-digit invoice number (3249 distinct). The manual puts PrefixDoc at this position; PrefixDoc is NOT emitted here, it is emitted at [19].
[ 3] ArtID -> string, article CodExtern (empty on 5239/10020 service lines). MAY CONTAIN ';'.
[ 4] Cant -> RO decimal, sign-bearing ('1','2','-1','0,1')
[ 5] DenUM -> string ('Buc','ORA','Set','cursa'; 18 distinct)
[ 6] Pret -> RO decimal
[ 7] DenGest -> string ('Piese','Utilaje','Imob','F'); empty exactly on the 4885 rows where [17]=='S'
[ 8] CodInternArt -> int, '0' on 10020/10020
[ 9] LocatieClient -> string ('SEDIU SOCIAL','Prejmer','Brasov','SEDIUL SOCIAL')
[10] MarcaAgent -> int ('0' x9947, '52' x73). PROVEN: the set of NrDoc with [10]=='52' is exactly {33354,33435,33559,33788,33789,33838,34047,34085,34345,35691,35779,35982,35984}, identical to the set of NrDoc with VanzariLuna[8]=='52' (which the manual names MarcaAgent). The manual's ordering (Comision;CodFisca;MarcaAgent) is wrong here.
[11] CodFisca -> string, customer CIF (466 distinct)
[12] Comision | ValAchizitie -> numeric, '0' on 10020/10020
[13] Adresa -> string (manual lists Adresa immediately after Locatie client; it is actually here, after CodFisca)
[14] ValAchizitie | Comision -> numeric, '0' on 10020/10020 — which of [12]/[14] is Comision and which is ValAchizitie CANNOT be resolved: both are constant '0' in all 10026 records
[15] CodPostal -> string, old 4-digit RO codes ('2200' Brasov x3019, '2300', '1900', '3400'; 13 distinct)
[16] ClasaArticol -> string ('4 PIESE','1 MANOPERA','2 DEPLASARE','INCHIRIERI'; 8 distinct)
[17] FlagDescarcareStoc -> string, UNDOCUMENTED. '=' x5135 / 'S' x4885. The 5135 '=' rows are exactly the 5135 rows with DenGest[7] non-empty => '=' is a stock discharge, 'S' is a service line.
[18] CantDescarcata -> RO decimal, UNDOCUMENTED. '0' on all 4885 'S' rows; otherwise mirrors Cant[4] ('1' x3583, '2' x521, '4', '0,1').
[19] PrefixDoc -> string ('BV AEG' x10007, 'INVAEG' x10, 'AEG' x3). PROVEN: equals VanzariLuna[14] (manual PrefixCarnet) on 9711/9711 uniquely-joined line pairs; and VanzariLuna[15] == 'F.' + this + NrDoc.
[20] SimbolMoneda -> string ('Lei' x10010, 'Eur' x10), UNDOCUMENTED
[21] (unnamed) -> always empty, 0/10020 non-empty across all 13 months
[22] (unnamed) -> always empty; this is either the record terminator or a second unused column — indistinguishable from the data
```

**Change:** MISALIGNED on 6/10026 sample records and mis-named on 5 positions. Fixes:

(a) Rename: [8] Unknown8 -> CodInternArt; [10] Unknown10 -> MarcaAgent (proven, see evidence); [12]/[14] Unknown12/Unknown14 -> Comision/ValAchizitie (both constant '0', order unresolved — keep both, document the ambiguity); [17] TipDocument -> FlagDescarcareStoc ('='/'S' is NOT a document type); [18] Unknown18 -> CantDescarcata; [19] PrefixCarnet -> PrefixDoc (this is the manual's PrefixDoc, displaced from documented position 2). [21]/[22] stay named-unknown.

(b) Replace splitFields(rec, 23) at queries.go:90 with a LEFT-SHAPE-ANCHORED parse that never depends on len(parts):
    parts := strings.Split(rec, ";")            // full split, no N
    // [0..2] are fixed: PartID, Zi(int), NrDoc(int)
    // find the smallest j >= 4 with: parts[j] is an RO decimal AND parts[j+1] matches ^[A-Za-z][A-Za-z. ]{0,9}$ (unit) AND parts[j+2] is an RO decimal
    // ArtID = strings.Join(parts[3:j], ";")   // re-joins the embedded separator losslessly
    // everything from j onward is positional: Cant=parts[j], DenUM=parts[j+1], Pret=parts[j+2], ... PrefixDoc=parts[j+15], SimbolMoneda=parts[j+16]
    // index with a bounds-checked accessor that returns "" past the end; collect parts[j+18:] into Extra []string
Why this is safe against a live width larger than 23: the anchor is found by SCANNING FORWARD from a fixed head, never by counting backward from len(parts). Extra live columns land in Extra and cannot touch ArtID/CodFisca/NrDoc. This is the exact inverse of the reverted GetNomenclatorArticole change, which merged on len(parts) > count and therefore fed live width growth straight into the join key.
The triple test cannot false-anchor on the real data: the field after ArtID is always the quantity (numeric), never a unit, so a numeric code fragment such as ' 2003843376002' or '  02420.079.082' fails at the unit test and the scan continues (verified by hand against all 46 semicolon-bearing CodExtern values).

(c) Records may contain embedded newlines (13 dump records show them); do not trim or split on \n.

### `GetVanzariLuna`

26 positions in the April 2026 dump, 10 of them named by the vendor manual.

```
[ 0] IDPartener -> string, partner CodExtern = CIF (empty on 16/9940)
[ 1] Zi -> int, 1..31
[ 2] NrFactura -> int, 5-digit
[ 3] IDArticol -> string, article CodExtern (empty on 5203/9940). MAY CONTAIN ';'.
[ 4] NumarComanda -> string/int, deviz or order number ('19760','20270','19334'); empty on 2858/9940
[ 5] Cant -> RO decimal, sign-bearing
[ 6] DenUM -> string ('Buc','ORA','Set','cursa')
[ 7] Pret -> RO decimal
[ 8] MarcaAgent -> int; only ever '52' (72 rows) or empty
[ 9] ValoareFactura -> RO decimal, document total incl. VAT
[10] DataScadenta -> string date dd.mm.yyyy, UNDOCUMENTED. Always >= [19]; equals [19] + the contractual term (15/30 days) in every inspected row.
[11] FlagTVAInclusaInPret -> string, UNDOCUMENTED. Only value ever emitted: 'NU' (9940/9940).
[12] CotaTVA -> int, UNDOCUMENTED. '21' x6754, '19' x3174, '0' x11, '20' x1.
[13] TipDocument -> string, UNDOCUMENTED. 'F' x9932 (factura), 'Iv' x8.
[14] PrefixCarnet -> string, UNDOCUMENTED. 'BV AEG' x9929, 'INVAEG' x8, 'AEG' x3. Equals VanzariExt[19] on 9711/9711 joined pairs.
[15] SerieDocument -> string, UNDOCUMENTED. Always 'F.' + [14] + [2], e.g. 'F.BV AEG35166'.
[16] DenArticol -> string, UNDOCUMENTED ('Manopera','Materiale necesare reparatiei','Deplasare echipa de interventie').
[17] (unnamed) -> numeric, '0' x9558 / empty x382. UNNAMEABLE: constant where present, so no column in GetVanzariExt discriminates it. The 382 empties are all 'Avans marfa' (advance-invoice) lines.
[18] (unnamed) -> numeric, '0' x7275 / empty x2665. UNNAMEABLE: does NOT correlate with stock-vs-service (cross-tab against VanzariExt DenGest/[17] on 9711 joined pairs shows all four combinations present).
[19] DataFactura -> string date dd.mm.yyyy, UNDOCUMENTED. PROVEN: day-of-month of [19] equals Zi[1] on 9940/9940 records, 0 mismatches.
[20] SediuClient -> string ('SEDIU SOCIAL','Prejmer','Brasov'), UNDOCUMENTED
[21] AdresaClient -> string, UNDOCUMENTED
[22] LocalitateClient -> string ('BRASOV','GHIMBAV','PREJMER'), UNDOCUMENTED
[23] ObservatiiDocument -> free text up to 771 chars, UNDOCUMENTED. MAY CONTAIN ';' AND embedded newlines.
[24] ObservatiiLinie -> free text up to 99 chars ('10%','Manopera - 1 ora','Perioada: 01.05.2025-31.05.2025'), UNDOCUMENTED. MAY CONTAIN ';'.
[25] (unnamed) -> always empty, 0/9940 non-empty; terminator or unused column
```

**Change:** MISALIGNED on 72/10012 sample records. Fixes:

(a) Names: keep [0..9] from the manual. Rename [17] Unknown17 and [18] Unknown18 -> leave as explicitly named-unknown (Unnamed17/Unnamed18) with a comment stating both are constant '0'-or-empty and no evidence names them; do NOT invent ProcDiscount/AdaosDiminuare. Adopt evidence-based names for [10..16] and [19..24] as listed in the layout ([23]/[24] renamed ObservatiiDocument/ObservatiiLinie from the current ObservatiiFactura/Observatii2).

(b) Replace splitFields(rec, 26) at queries.go:130 with the same left-anchored parse, plus a tail rule:
    parts := strings.Split(rec, ";")
    // [0..2] fixed: IDPartener, Zi(int), NrFactura(int)
    // find smallest j >= 5 with parts[j] RO-decimal AND parts[j+1] unit-shaped AND parts[j+2] RO-decimal  -> j is Cant
    //   NumarComanda = parts[j-1];  IDArticol = strings.Join(parts[3:j-1], ";")
    // positional forward from j up to and including LocalitateClient = parts[j+17]
    // ObservatiiDocument = strings.Join(parts[j+18:], ";")   // greedy remainder
  Verified by hand on the misaligning row: parts[3]='111846SA ', [4]=' A1002', [5]='18829', [6]='1', [7]='Buc', [8]='1220'; scanning from j=5 rejects j=5 ('18829','1',...) because '1' is not unit-shaped, accepts j=6 ('1','Buc','1220'); NumarComanda=parts[5]='18829' and IDArticol='111846SA ; A1002'. Correct.

(c) The greedy remainder deliberately puts ALL residual width uncertainty into a free-text note, where it is harmless: extra live columns and extra embedded ';' both end up appended to ObservatiiDocument and can never shift IDArticol, NrFactura or NumarComanda. Cost: ObservatiiLinie[24] is no longer separable. If it must stay separate, split the remainder on the LAST ';' only when the record's separator count equals the modal 25 -- i.e. treat [24] as best-effort, never as a key. Do NOT reconstruct it by counting back from len(parts) unconditionally; that is the reverted-change failure mode.

(d) Records may contain embedded newlines; never split on \n.
