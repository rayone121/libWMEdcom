# Structuri import plata viramente

<sub>Source: `22_Structuri import din alte aplicatii__Import plata virament.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structuri import plata viramente` |  |
| `[InfoPachet]` | functie de unde se face incasarea |
| `AnLucru=2010` |  |
| `LunaLucru=3` |  |
| `TotalDocumente=1` | numele complet al bancii la nivel national; |
| `[Document1]` | simbolul bancii; se completeaza daca |
| `Sursa=BANCA` |  |
| `NumeBanca=BANCA AGRICOLA` |  |
| `SimbolBanca=BA` |  |
| `NumeCont=CONT BA` | un nume de CASA sau nume de cont bancar |
| `NrCont=11111` |  |
| `LocalitateCont=2 MAI` |  |
| `FilialaCont=2 MAI` |  |
| `ZiuaIncasarii=03` | cod extern/intern/fiscal partener - se regleaza |
| `TotalIncasari=1` | prin constanta: “Cod pentru identificare |
| `TipTranzactie=CURENTA` | PARTENER” vezi constante generale > |
| `[Document1-Incasare1]` |  |
| `DocIncasare=OP` |  |
| `NrDocument=5555` | Total incasari pe acest extras de cont |
| `Reprezinta=PLATA VIRAMENT` |  |

Poate fi CASA sau BANCA, in
se completeaza daca SURSA = BANCA
SURSA = BANCA
import date din alte aplicatii.
Reprezinta ziua registrului de casa/banca
NrContDest=22222 Valabila doar pentru sursa=BANCA; poate lua valorile
ValIncasata=200 CURS pentru cec-uri sau CURENTA
