# Structuri import încasare direct pe cont în

<sub>Source: `22_Structuri import din alte aplicatii__Import incasare direct pe cont.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` | WinMENTOR |
| `AnLucru=2022` |  |
| `LunaLucru=7` |  |
| `TotalDocumente=1` |  |
| `[Document1]` |  |
| `` | Poate fi CASA sau BANCA, în |
| `Sursa=BANCA` | funcție de unde se face |
| `` | numele complet al băncii la nivel |
| `NumeBanca=BANCA COMERCIALA ROMANA` | național; se completează dacă |
| `` | simbolul băncii; se completează |
| `SimbolBanca=BCR` |  |
| `` | un nume de CASĂ sau nume de |
| `NumeCont=CONT BCR` |  |
| `` | cod extern/intern/fiscal partener |
| `NrCont=22222` | „Cod pentru identificare |
| `LocalitateCont=2 MAI` |  |
| `FilialaCont=2 MAI` |  |
| `ZiuaIncasarii=03 ` | ziua registrului de casă/bancă |
| `` | total încasări pe acest extras de |
| `TotalIncasari=1` |  |
| `` | valabilă doar pentru sursa = |
| `TipTranzactie=CURENTA` |  |
| `[Document1-Incasare1]` |  |
| `DocIncasare=OP` |  |
| `NrDocument=5555` |  |
| `Reprezinta=DIRECT PE VENITURI` |  |
| `SIMBOLCONT=101.06` |  |
| `ValIncasata=200` |  |
| `TVAIncasat=0` |  |

Structuri import încasare direct pe cont în
încasarea
SURSA = BANCA
dacă SURSA = BANCA
cont bancar
– se reglează prin constanta:
PARTENER” vezi constante
generale >
cont
BANCA; poate lua valorile
CURS pentru cec-uri sau
CURENTĂ
