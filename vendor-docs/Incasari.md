# Structură import încasări în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Incasari.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2021` |  |
| `LunaLucru=5` |  |
| `TotalDocumente=1` |  |
| `[Document1]` |  |
| `Sursa =` | CASA, BANCA sau AVANS DECONTARE |
| `NumeBanca =` | Numele complet al Băncii pentru SURSA=BANCA |
| `Simbol Banca =` | Simbolul băncii dacă SURSA = BANCA |
| `NumeCont =` |  |
| `NrCont =` | Contul bancar pentru SURSA = BANCA |
| `LocalitateaCont =` | Localitatea băncii pentru SURSA = BANCA |
| `FilialaCont =` | Filiala băncii pentru SURSA = BANCA |
| `MonedaBanca =` | Simbol monedă cont bancar sau casă |
| `TipTranzactie =` | • CURS pentru CEC-uri |
| `MarcaAgent =` |  |
| `ZiuaIncasarii =` | Reprezintă ziua registrului de casă/bancă |
| `TotalIncasari =` | Un document conține totalitatea plăților făcute pe |
| `Moneda =` |  |
| `DECONTARE` |  |
| `ContPersonal =` |  |
| `DECONTARE` |  |
| `[Document1-Incasare1]` |  |
| `Reprezinta =` |  |
| `SerieCarnet =` |  |
| `• INTERNA` |  |
| `TipFactura =` |  |
| `contul de client` |  |
| `CursIncasare =` |  |
| `Marca angajatului` | dacă SURSA=AVANS |
| `MarcaAngajat =` |  |
| `DECONTARE` |  |
| `Cont =` | Se completează în cazurile: |
| `[Document1-Incasare2]` |  |
| `DocIncasare =` | Simbolul documentului încasat (C,DP,CD etc.) |
| `NrDocument =` | Numărul documentului |
| `Reprezinta =` |  |
| `NrFactura =` | Numărul facturii încasate |
| `SerieCarnet =` | Se completează seria facturii încasate |
| `CodClient =` | Codul intern/extern de identificare a clientului |
| `LocalitateClient =` | Localitatea clientului |
| `ValIncasata =` | Valoarea încasată |
| `TVAIncasat =` | Valoarea TVA-ului aferent facturii |
| `Cont =` | REPREZINTA=INCASAREPENALITATI |
| `Exemplu:` |  |
| `[InfoPachet] AnLucru = 2022` |  |
| `LunaLucru = 11` |  |
| `TotalDocumente = 1` |  |
| `[Document1]` |  |
| `se face încasarea)` |  |
| `= BANCA)` |  |
| `Filiala Rahova)` |  |
| `LocalitateCont =` |  |
| `FilialaCont =` |  |
| `CEC-uri sau CURENTA}` |  |
| `[Document1-Incasare1] DocIncasare = CD` |  |
| `abonament` |  |
| `NrFactura = 0000000057` |  |
| `SerieCarnet = AAA` |  |
| `la data încasării}` |  |
| `Comision = {valoare comision}` |  |
| `Client = S.C. CAMELIA S.R.L.` |  |
| `CodClient = 000000000002-000017652001` |  |
| `ValIncasata = 100` |  |
| `TVAIncasat = 24` |  |
| `[Document1-Incasare2]` |  |
| `DocIncasare = CD` |  |
| `NrDocument = 9444607` |  |
| `Reprezinta = tip (idem incasare 1)` |  |
| `NrFactura = 0000000057` |  |
| `SerieCarnet = AAA` |  |
| `Client = S.C. CAMELIA S.R.L.` |  |
| `CodClient = 000000000002-000017652001` |  |
| `LocalitateClient = (localitate)` |  |
| `ValIncasata = 0` |  |
| `TVAIncasat = 637937` |  |
| `scris această interfață.` |  |
| `consideră încasat documentul):` |  |
| `[IMPORT_INCASARI] MARJAEROARE=0.5` |  |

Structură import încasări în WinMENTOR
Numele pentru CASSA sau al contului bancar
pentru SURSA = BANCA
Se completează dacă SURSA = BANCA:
• CURENTA
Se completează dacă SURSA = AVANS
DECONTARE
Totalul încasărilor pe acest extras de cont.
un cont (CASA sau BANCA) pentru o anumită zi
(similar unui registru de casă/bancă).
Se completează dacă SURSA=AVANS
Se completează dacă SURSA=AVANS
DocIncasare = Simbolul documentului de încasare(C, DP, CD etc.)
NrDocument = Numărul documentului
Se completează obligatoriu cu una din opțiuni:
PLATA FACTURA, INCASARE PENALITATI,
ALIMENTARE CREDIT – se utilizează pentru
AVANS, DIMINUARE CREDIT, PERSONAL
ANGAJAT, RETUR CLIENT (se folosește pentru
retur client și se va regăsi în Trezorerie în Plăți)
NrFactura = Numărul facturii încasate
Se completează seria facturii încasate
• EXTERNA – facturi în valută
• NC – note contabile ce creează obligație pe
• BA – facturi de la persoane
Comision = Valoarea comisionului
Client = Denumirea clientului
CodClient = Codul intern/extern de identificare a clientului
LocalitateClient = Localitatea clientului
ValIncasata = Valoarea încasată a facturii
TVAIncasat = Valoarea TVA-ului aferent facturii încasate
REPREZINTA=INCASARE PENALITATI
REPREZINTA=ALIMENTARE CREDIT
Poate lua una din valori: PLATA FACTURA, PLATA
PENALITATI, ALIMENTARE CREDIT – se
folosește pentru AVANS, DIMINUARE CREDIT,
PERSONAL ANGAJAT
Se completează în cazurile:
REPREZINTA=ALIMENTARE CREDIT
Sursa = CASA (poate fi CASA, BANCA sau AVANS DECONTARE, în funcție de unde
NumeBanca = (numele complet al băncii la nivel național; se completează dacă SURSA
SimbolBanca = (simbolul băncii; se completează dacă SURSA = BANCA)
NumeCont = Casa 000032 (un nume de CASA sau nume de cont bancar - exemplu
NrCont = (numărul de cont bancar; se completează dacă SURSA = BANCA)
MonedaBanca = {simbol monedă cont bancar sau casă}
TipTranzactie = {valabilă doar pentru sursa=BANCA; poate lua valorile CURS pentru
MarcaAgent =1 (se completează în cazul în care SURSA = AVANS DECONTARE)
ZiuaIncasarii =10 (reprezintă ziua registrului de casă/bancă)
TotalIncasari =2 (total încasări pe acest extras de cont) {după cum vedeți, un document
conține totalitatea încasărilor făcute pe un cont (CASA sau Banca) într-o anumită zi
(similar unui registru de casă într-o anumită zi)}
Moneda = EU (se completează în cazul în care SURSA = AVANS DECONTARE)
ContPersonal = 542 (se completează în cazul în care SURSA = AVANS DECONTARE)
NrDocument = 9444609 (număr chitanță de încasare)
Reprezinta = tip (poate lua una din valorile: INCASARE FACTURA, INCASARE
PENALITATI, ALIMENTARE CREDIT, DIMINUARE CREDIT, PERSONAL ANGAJAT)
FacturaAbonament = D - se completează doar în cazul în care factura este de tip
TipFactura = {tipul facturii stinse; poate lua valorile INTERNA sau EXTERNA (pentru
facturi în valută) sau NC pentu note contabile ce creează obligație pe partener, pe
contul de client) sau LA AVIZ (pentru facturile la aviz)}
FacturaAbonament=N sau D dacă factura provine din contractele de abonament
CursIncasare = {dacă se încasează o factură în valută, CursIncasare reprezintă cursul
MarcaAgent = (marca agentului încasator) LocalitateClient = (localitate)
Cont = simbol cont contabil (pentru INCASARE PENALITATI și ALIMENTARE CREDIT)
Cont = simbol (pentru INCASARE PENALITATI și CREARE CREDIT).
După cum vedeți din exemplu, în cazul când „REPREZINTA” = INCASARE FACTURA
se generează două încasări: una pentru valoarea propriu-zisă, alta pentru TVA. Această
soluție a fost adoptată ca urmare a cererii formulate de Congaz – cei pentru care am
Pentru cazul în care apar diferențe între rest de încasat și valoarea documentului de
încasat, în fișierul Defaults.ini se stabilește o marjă de diferență (până la care se
