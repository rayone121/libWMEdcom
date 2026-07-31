# Structură import fișier facturi în așteptare în

<sub>Source: `22_Structuri import din alte aplicatii__Factura in asteptare.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=3` |  |
| `TipDocument=FACTURA INTRARE` |  |
| `TotalFacturi=1` |  |
| `[Factura_1]` |  |
| `FactInAsteptare=D` |  |
| `Data=02.03.2022  data facturii` |  |
| `SerieCarnet=AA` |  |
| ` Codul` | utilizat pentru |
| `aplicații").` |  |
| `TVAINCASARE=D` |  |
| `PRORATA=D` |  |
| `AutoFactura=D` |  |
| `FacturaSimplificata=D` |  |
| ` TipTVA:` |  |
| `311CF;` |  |
| `TipTVA=0` |  |
| `312CF;` |  |
| `Moneda=Euro` |  |
| `Curs=5.01` |  |
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
| `Item_1_SimbolCont = 473` | Cont reprezentativ |
| `Item_1_UM1=20 ` | cantitatea în UM alternativă 1 |
| `Item_1_UM2=15 ` | cantitatea în UM alternativă 1 |
| `` | valoarea TVA-ului la nivel de |
| `Item_1_TVA=` |  |
| `` | Se va înscrie serie pentru articol |
| `Item_1_Serii` | SERIE PE LOT, sau pentru |
| `` | Poate lua valorile 0, 1, 3 |
| `Item_1_TVANEDED` | (deductibil 100% , nedeductibil |
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

Structură import fișier facturi în așteptare în
WinMENTOR
NrDoc=5144257  numărul facturii
identificarea furnizorului (se
CodFurnizor=R 717359 precizează tipul lui în constante
generale: „Import date din alte
 0 – nedefinit;
 3 – regim special de scutire
 4 – regim special de scutire
 5 – achiziții UE de bunuri cu
instalare și montaj.
Scadenta=01.05.2022  scadența facturii
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
100%, ½ neded)
Sinteza articole factură {11 campuri separate prin ";" }: Item_1=;;;;;;;;;;;
Primele 4 câmpuri sunt obligatorii, iar pentru restul se poate lăsa spațiu între punct și
