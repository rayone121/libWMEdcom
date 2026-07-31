# Structură import fișier plați în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Plati.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=11` |  |
| `TotalDocumente=1` |  |
| `[Document1]` |  |
| `` | poate fi CASA, BANCA sau |
| `Sursa=CASA` | AVANS DECONTARE, în funcție |
| `` | numele complet al băncii la nivel |
| `NumeBanca=` | național; se completează dacă |
| `` | simbolul băncii; se completează |
| `SimbolBanca=` |  |
| `` | un nume de CASĂ sau nume de |
| `NumeCont=Casa 000032` | cont bancar - exemplu Filiala |
| `` | numărul de cont bancar; se |
| `NrCont=` | completează dacă SURSA = |
| `LocalitateCont=` |  |
| `FilialaCont=` |  |
| `` | simbol moneda cont bancar sau |
| `MonedaBanca=` |  |
| `` | valabilă doar pentru SURSA = |
| `TipTranzactie=` |  |
| `` | se completează în cazul în care |
| `MarcaAgent=1` |  |
| `` | reprezintă ziua registrului de |
| `ZiuaPlatii=10` |  |
| `` | total plăți pe acest extras de |
| `TotalPlati=2` |  |
| `Moneda=EU` |  |
| `ContPersonal=542` |  |
| `[Document1-Plata1]` |  |
| `DocPlata=CD` |  |
| ` PLATA FACTURĂ` |  |
| `NrFactura=0000000057` |  |
| `SerieCarnet=AAA` |  |
| `valorile` | INTERNĂ sau |
| `Furnizor=S.C. CAMELIA S.R.L.` |  |
| `CodFurnizor=1111111` |  |
| `LocalitateFurnizor=  localitate` |  |
| `ValPlatita=100` |  |
| `TVAPlatit=24` |  |
| `Cont= PLATĂ` | PENALITĂȚI și |
| `[Document1-Plata2]` |  |
| `DocIPlata=CD` |  |
| `NrDocument=9444607` |  |
| `Reprezinta=tip ` | idem încasare 1 |
| `NrFactura=0000000057` |  |
| `SerieCarnet=AAA` |  |
| `Client=S.C. CAMELIA S.R.L.` |  |
| `CodClient=222222222` |  |
| `LocalitateFurnizor= ` | localitate |
| `ValPlatita=0` |  |
| `TVAPlatit=637937` |  |
| `` | (pentru PLATĂ PENALITĂȚI și |
| `Cont=simbol` | când „reprezintă” = PLATĂ |

Structură import fișier plați în WinMENTOR
de unde se face plata
SURSA = BANCA
dacă SURSA = BANCA
Rahova
BANCA
casă
BANCA; poate lua valorile
CURS pentru cec - uri sau
CURENTĂ
SURSA = AVANS DECONTARE
casă/bancă
cont (după cum vedeți, un
document conține totalitatea
plăților făcute pe un cont, CASA
sau BANCA, într - o anumită zi,
similar unui registru de casă
într-o anumită zi)
 se completează în cazul în care
SURSA = AVANS DECONTARE
 se completează în cazul în care
SURSA = AVANS DECONTARE
NrDocument=9444609  număr chitanța de plată
poate lua una din valorile:
 DIRECT PE CHELTUIELI
Reprezinta=tip  PLATA PENALITĂȚI
 ALIMENTARE CREDIT
 DIMINUARE CREDIT
 PERSONAL ANGAJAT
 tipul facturii stinse; poate lua
EXTERNĂ (pentru facturi în
valută) sau NC pentu note
TipFactura = contabile ce creează obligație
pe partener pe contul de client),
BA (pentru intrările de la
persoane fizice) sau BF (pentru
bonuri fiscale)
 dacă se plătește o factură în
CursPlata= valută, CursPlata reprezintă
cursul la data plății
Comision=  valoare comision
MarcaAgent=  marca agentului plătitor
 simbol cont contabil (pentru
ALIMENTARE CREDIT)
CREARE CREDIT) După cum
vedeți din exemplu, în cazul
FACTURĂ se generează două
plăți : una pentru valoarea
propriu - zisă, alta pentru TVA.
