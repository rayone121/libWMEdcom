# Structură import fișier facturi intrare în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Facturi de intrare.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=3` |  |
| `TipDocument=FACTURA INTRARE` |  |
| `TotalFacturi=3` |  |
| `[Factura_1]` |  |
| `NrDoc=5144257 ` | numărul facturii |
| `ClasificareSAFT=` |  |
| `IDdescarcare` |  ID-ul de descărcare din SPV |
| `Data=02.03.2022 ` | data facturii |
| `SerieCarnet=AA` |  |
| `Locatie= ` | denumirea sediului |
| `` | Se poate specifica nr. NIR dacă |
| `NrNir=999` |  |
| `` | Codul utilizat pentru |
| `CodFurnizor=R 717359` | precizează tipul lui în constante |
| `TVAINCASARE=D` |  |
| `PRORATA=D` |  |
| `AutoFactura=D` |  |
| `FacturaSimplificata=D` |  |
| `` | TipTVA: |
| `TipTVA=0` | 311CF; |
| `Moneda=Euro` | 312CF; instalare și montaj. |
| `Curs=5.01` |  |
| `Scadenta=01.05.2022 ` | scadența facturii |
| `Majorari=0` |  |
| `Observatii=` |  |
| `` | Corespondența dintre numărul |
| `TotalArticole=2` | descrieri de articole este |
| `Discount=11.8300` |  |
| `[Items_1]` |  |
| `` | Cod articol (intern/extern) |
| `` | UM; |
| `` | Cantitate; |
| `Item_1=11126;Buc;1.00;489.47;Falticeni;0;;603.7 ` | Preț achiziție; |
| `7;;;observatii la nivel de articol; ` | Simbol gestiune: pentru |
| `` | Discount linie; |
| `` | Preț înregistrare pentru |
| `` | Observații la nivel de articol. |
| `Item_2=11123;Buc;1.00;37.40;Falticeni;;;51.20;` |  |
| `Item_1_UM1=20 ` | cantitatea în UM alternativă 1 |
| `Item_1_UM2=15 ` | cantitatea în UM alternativă 1 |
| `` | valoarea TVA-ului la nivel de |
| `Item_1_TVA=` |  |
| `` | Se va înscrie serie pentru articol |
| `Item_1_Serii` | SERIE PE LOT, sau pentru |
| `` | Poate lua valorile 0, 1, 3 |
| `Item_1_TVANEDED` | (deductibil 100% , nedeductibil |
| `` | 100%, ½ neded) Tipul contabil pentru recepția |
| `Item_1_TipContabil=` |  |
| `` | articolului Asigură corespondența dintre |
| `[Factura_2]` |  |
| `NrDoc=6` | antetul facturii si conținutul său. |
| `Data=02.03.2022` |  |
| `CodFurnizor=R 717359` |  |
| `Scadenta=01.05.2022` |  |
| `Majorari=0` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Discount=15.0100` |  |
| `[Items_2]` |  |
| `Item_1=11123;Kg;1.00; 49.56;Falticeni;0;;-60.48;` |  |
| `[Factura_3]` |  |
| `NrDoc=7` |  |
| `Data=03.03.2022` |  |
| `CodFurnizor=R 4851409` |  |
| `Scadenta=02.05.2022` |  |
| `Majorari=0` |  |
| `Observatii` |  |
| `TotalArticole=1` |  |
| `Discount=18.0000` |  |
| `ValoareSuplimentare=10` | Se repartizează pe fiecare articol |
| `ContSuplimentare=101.01` |  |
| `` | TipRepartizare: |
| `TipRepartizare=0` |  0 – valoare; |
| `[Items_3]` |  1 – masă;  2 – volum. |
| `Item_1=11124;Lei;1.00;12.00;;;624;;15.12.2008;1` |  |
| `0;articol cu observatii` |  |
| ` Cod intern/extern articol;` |  |
| ` UM;` |  |
| ` Cantitate;` |  |
| ` Preț achiziție;` |  |
| ` Gestiune;` |  |
| ` Discount linie;` |  |
| ` Simbol cont articol serviciu;` |  |
| ` Preț înregistrare;` |  |
| ` Termen garanție;` |  |
| ` Valoare suplimentară;` |  |
| ` Observații la nivel articol.` |  |
| `virgula, așa cum se observă în exemplele de mai sus.` |  |

Structură import fișier facturi intrare în WinMENTOR
Valori posibile (se poate utiliza codul din
SAF-T sau corespondenţa lui):
• Factura iniţială: 380 sau 0
• Factura storno: 381 sau 1
• Factura de corecţie: 384 sau 2
• Autofactura: 389 sau 3
nu se dorește numerotarea
automată de către
WinMENTOR. Funcționează
doar dacă este completat
SerieCarnet.
identificarea furnizorului (se
generale: „Import date din alte
aplicații").
 0 – nedefinit;
 3 – regim special de scutire
 4 – regim special de scutire
 5 – achiziții UE de bunuri cu
de articole și numărul de
validată la preluare.
(Item_1;Item_2)
conform constantei de import
date din alte aplicații: cod
identificare articole;
recepție/repartizare cheltuieli;
articolele ce au tip contabil
implicit cu adaos;
linie
– o singură serie dacă este
fiecare bucată, dacă este
SERIE PE BUCATA .
 Total valoare „suplimentare”.
în parte, la poziția respectivă.
Sinteza articole factură {11 campuri separate prin ";" }: Item_1=;;;;;;;;;;;
Primele 4 câmpuri sunt obligatorii, iar pentru restul se poate lăsa spațiu între punct și
