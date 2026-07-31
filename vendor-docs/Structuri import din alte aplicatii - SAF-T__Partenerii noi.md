# Structură import fișier pentru partenerii noi

<sub>Source: `22_Structuri import din alte aplicatii__Structuri import din alte aplicatii - SAF-T__Partenerii noi.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `[ParteneriNoi_C000020]` |  |
| `Denumire=UP20` |  |
| `ExportSAFT=D` | Valoarea DA indică faptul că partenerul va fi importat în Nomenclator cu bifa Prezent pe declarația SAF-T. |
| `Localitate=BACAU;IASI` | Valoarea NU, determină importul fără bifă. |
| `Tara=Romania` |  |
| `SimbolTara=RO` |  |
| `Judet=BC;IS` |  |
| `CodFiscal=RO256644;889966` |  |
| `RegistruComert=` |  |
| `MarcaAgent=10;11` |  |
| `Sediu=Sediu1;Sediu2` |  Dacă partenerul are mai multe sedii, atunci la Adresă, Localitate, Județ, Telefon, Email, valorile aferente fiecărui sediu, separate prin ; (ordinea este cea în care au fost descrise sediile). |
| `Clasa=10` |  simbolul clasei de încadrare |
| `CodExtern=` |  |
| `CodIntern=` |  |
| `Telefon=0230/215679;0232/222444` |  |
| `PersoanaFizica=NU` |  |
| `InstitPublica=D` | Se acceptă ca parametru și D și Da. |
| `SerieBuletin=` |  |
| `NumarBuletin=` |  |
| `Banci=` |  |
| `Conturi=` |  |
| `` | simbolul caracterizării contabile |
| `TipContabil=Tipic` |  |
| `` | cod extern/intern/fiscal partener |
| `[ParteneriNoi_F000020]` |  |
| `Denumire=Ion Popescu` |  |
| `Localitate=IASI` |  |
| `Tara=Romania` |  |
| `SimbolTara=RO` |  |
| `Judet=IS` |  |
| `CodFiscal=2802101212228` |  |
| `RegistruComert=` |  |
| `MarcaAgent=12` |  |
| `Adresa=Str. Bucium, nr.1` |  |
| `Sediu=Sediu 3` |  |
| `Clasa=10` |  |
| `CodExtern=` |  |
| `CodIntern=` |  |
| `Telefon` |  |
| `Email=sediu3@sediu3.` |  |
| `` | Date de identificare ale |
| `PersoanaFizica=Da` |  |
| `` | Date de identificare ale |
| `SerieBuletin=MX` |  |
| `` | Date de identificare ale |
| `NumarBuletin=808080` |  |
| `Banci=` |  |
| `Conturi=` |  |
| `TipContabil=` |  |
| `ScadentaC=` | Scadența implicită la cumpărare (zile) |
| `ScadentaV=` | Scadența implicită la vânzare (zile) |

Structură import fișier pentru partenerii noi
[PARTNER.txt] în WinMENTOR
Parametrul din fișier Explicații
Adresa=Str.AAA, nr.1;Str.BBB, nr.2 Marcă agent se completează
Email=sediu1@up20.ro;sediu2@up20.ro
aferente partenerilor
– se reglează prin constanta:
„Cod pentru identificare
PARTENERI” vezi constante
generale > import date din alte
aplicații.
partenerilor persoane fizice.
partenerilor persoane fizice.
partenerilor persoane fizice.
Fișierul poate conține toți partenerii din aplicația sursă, consultarea acestui fișier
făcându-se numai în cazul în care codul de partener din fișierul de tranzacții nu este
găsit în nomenclatorul de parteneri din MENTOR.
Se pot adăuga câmpuri noi în aceasta structură, dacă solicitați acest lucru, în măsura în
care există câmpurile respective în WinMentor.
Identificarea localității la partenerii noi se face după cuplu: denumire localitate/simbol
județ, deci nu se va face dacă nu este completat și simbolul județului. Sediul social va fi
primul sediu din cele n sedii ale partenerului (ex: Sediu1).
