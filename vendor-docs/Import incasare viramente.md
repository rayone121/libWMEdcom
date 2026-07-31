# Structură import încasare viramente în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Import incasare viramente.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=3` |  |
| `TotalDocumente=1` |  |
| `[Document1]` |  |
| `` | Poate fi CASA sau BANCA, |
| `Sursa=BANCA` | în funcție de unde se face |
| `` | numele complet al băncii la |
| `NumeBanca=BANCA COMERCIALA ROMANA` | nivel național; se completează |
| `` | simbolul băncii; se |
| `SimbolBanca=BCR` | completează dacă SURSA = |
| `` | un nume de CASĂ sau nume |
| `NumeCont=CONT BCR` |  |
| `` | cod extern/intern/fiscal |
| `NrCont=22222` | constanta: „Cod pentru |
| `LocalitateCont=2 MAI` |  |
| `FilialaCont=2 MAI` |  |
| `ZiuaIncasarii=03 ` | ziua registrului de casă/bancă |
| `` | total încasări pe acest extras |
| `TotalIncasari=1` |  |
| `` | valabilă doar pentru sursa = |
| `TipTranzactie=CURENTA` |  |
| `[Document1-Incasare1]` |  |
| `DocIncasare=OP` |  |
| `NrDocument=5555` |  |
| `Reprezinta=INCASARE VIRAMENT` |  |
| `NrContDest=1111` |  |
| `ValIncasata=200` |  |
| `DocViramentStins=OP` |  |
| `NrViramentStins=1333` |  |

Structură import încasare viramente în WinMENTOR
încasarea
dacă SURSA = BANCA
BANCA
de cont bancar
partener – se reglează prin
identificare PARTENER” vezi
constante generale
de cont
BANCA; poate lua valorile
CURS pentru cec-uri sau
CURENTĂ
