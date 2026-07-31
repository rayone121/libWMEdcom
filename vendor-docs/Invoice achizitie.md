# Structură import achiziție în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Invoice achizitie.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=6` |  |
| `TipDocument=INVOICE` |  |
| `TotalFacturi=1` |  |
| `[Factura_1]` |  |
| `NrDoc=65` |  |
| `Data=12.06.2022` |  |
| `Lohn=D sau N` |  |
| `` | cod extern/intern/fiscal partener |
| `„Cod` | pentru identificare |
| `CodFurnizor=F000020` |  |
| `Moneda=EUR` |  |
| `Scadenta=31.12.2022` |  |
| `Majorari=12,2` |  |
| `Observatii=` |  |
| `TotalArticole=2` |  |
| `TaxareInversa=D sau N` |  |
| `[DVI_1]` |  |
| `NrDoc=6767` |  |
| `Data=12.06.2022` |  |
| `` | cod extern/intern/fiscal partener |
| `„Cod` | pentru identificare |
| `CodVama=F000020` |  |
| `Curs=11` |  |
| `[Transport_1]` |  |
| `Nrdoc=` |  |
| `CodTransportator=` |  |
| `Moneda=` |  |
| `Curs=` |  |
| `Valoare=` |  |
| `Scadenta=` |  |
| `Majorari=` |  |
| `[Asigurare_1]` |  |
| `NrDoc=` |  |
| `CodAsigurator=` |  |
| `Moneda=` |  |
| `Curs=` |  |
| `Valoare=` |  |
| `Scadenta=` |  |
| `Majorari=` |  |
| `ValoareSuplimentare=10` | Se repartizează pe fiecare articol în |
| `ContSuplimentare=101.01` |  |
| `` | TipRepartizare: |
| `TipRepartizare=0` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – se |
| `` | denumire unitate de măsură din |
| `Item_1=E0003621000;MT;1000.00;150.0000;C00` |  |
| `` | cantitate; |
| `200;-10;S100;TEST` |  |
| `` | preț; |
| `` | simbol gestiune recepție – |
| `` | discount; |
| `` | simbol cont serviciu; |
| `` | observații articol. |
| `Item_2=D5001110030;MT;1000.00;800.0000;C00` |  |
| `20;;;` |  |
| `„Articole.txt”.` |  |
| `descrise în fișierul „Partner.txt”.` |  |

Structură import achiziție în WinMENTOR
– se reglează prin constanta:
ARTICOLE” vezi constante
generale > import date din alte
aplicații.
– se reglează prin constanta:
ARTICOLE” vezi constante
generale > import date din alte
aplicații.
Total valoare „suplimentare”.
parte, la poziția respectivă.
 0 – valoare;
 1 – masă;
 2 – volum.
reglează prin constanta: „Cod
pentru identificare ARTICOLE”
vezi constante generale >import
date din alte aplicații;
WinMENTOR;
numai pentru articole de tip
stoc;
În cazul unor articole nou apărute în nomenclator ele vor fi descrise în fișierul
În cazul în care pentru gestiunea de recepție sunt coduri noi, neintroduse încă în baza
de date, ele vor fi descrise în fișierul „Gestiuni.txt”.
În cazul în care pentru client sunt coduri noi, neintroduse încă în baza de date, ele vor fi
