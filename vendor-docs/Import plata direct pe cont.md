# Structuri import plată direct pe cont în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Import plata direct pe cont.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structuri import plată direct pe cont în WinMENTOR` |  |
| `[InfoPachet]` |  |
| `AnLucru=2024` |  |
| `LunaLucru=02` |  |
| `TotalDocumente=1` |  |
| `[Document1]` |  |
| `•` | Poate fi CASA sau BANCA, în |
| `Sursa=BANCA` |  |
| `•` | numele complet al băncii la nivel |
| `NumeBanca=Banca Comerciala Romana` | național; se completează dacă |
| `•` | simbolul băncii; se completează |
| `SimbolBanca=BCR` |  |
| `LocalitateCont=Iasi •` | Localitatea filialei |
| `•` | un nume de CASĂ sau nume de |
| `NumeCont=CONT BCR` |  |
| `•` | cod extern/intern/fiscal partener |
| `NrCont=RO93RNCB0260003080950009` | „Cod pentru identificare |
| `•` | valabilă doar pentru sursa = |
| `TipTranzactie=CURENTA` |  |
| `ZiuaPlatii=10 •` | Ziua in care se face plata |
| `•` | total plăți pe acest extras de |
| `TotalPlati=1` |  |
| `[Document1-Plata1]` |  |
| `DocPlata=Ext` |  |
| `NrDocument=99` |  |
| `• DATORII TAXE` |  |
| `Reprezinta=tip` |  |
| `• DIRECT PE CHELTUIELI` |  |
| `ValPlatita=119` |  |
| `SIMBOLCONT=627` |  |
| `TvaPlatit=0` |  |

funcție de unde se face plata
SURSA = BANCA
dacă SURSA = BANCA
cont bancar
– se reglează prin constanta:
PARTENER” vezi constante
generale.
BANCA; poate lua valorile
CURS pentru cec-uri sau
CURENTĂ
cont
