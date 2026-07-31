# Structura import achiziție în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Structuri import din alte aplicatii - SAF-T__Invoice achizitie.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structura import achiziție în WinMENTOR` |  |
| `[InfoPachet]` |  |
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=6` |  |
| `TipDocument=INVOICE` |  |
| `TotalFacturi=1` |  |
| `[Factura_1]` |  |
| `NrDoc=65` |  |
| ` Factura iniţială:` | 380 sau 0 |
| `ClasificareSAFT=  Factura storno:` | 381 sau 1 |
| ` Autofactura:` | 389 sau 3 |
| `Data=12.06.2022` |  |
| `Lohn=D sau N` |  |
| `aplicaţii.` |  |
| `Moneda=EUR` |  |
| `Scadenta=31.12.2024` |  |
| `Majorari=12,2` |  |
| `Observatii=` |  |
| `TotalArticole=2` |  |
| `TaxareInversa=D sau N` |  |
| `[DVI_1]` |  |
| `NrDoc=6767` |  |
| `Data=12.06.2022` |  |
| ` cod extern/intern/fiscal partener – se` |  |
| `reglează prin constanta: „Cod pentru` |  |
| `identificare ARTICOLE” vezi Constante` |  |
| `CodVama=F000020 generale > Import date din alte` |  |
| `aplicaţii.` |  |
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
| `Total valoare „suplimentare”.` |  |
| `ValoareSuplimentare=10 Se repartizează pe fiecare articol în` |  |
| `parte, la poziția respectivă.` |  |
| `ContSuplimentare=101.01` |  |
| ` TipRepartizare:` |  |
| ` 0 – valoare;` |  |
| `TipRepartizare=0` |  |
| ` 1 – masă;` |  |
| ` 2 – volum.` |  |
| `[Items_1]` |  |
| `aplicaţii;` |  |
| `WinMENTOR;` |  |
| ` Cantitate;` |  |
| `Item_1=E0003621000;MT;1000.00;150.0000;` |  |
| ` Preţ;` |  |
| `C00200;10;S100;TEST` |  |
| ` Discount;` |  |
| `Item_2=D5001110030;MT;1000.00;800.0000;C002` |  |
| `0;;;` |  |
| `„Articole.txt”.` |  |
| `de date, ele vor fi descrise în fişierul „Gestiuni.txt”.` |  |
| `descrise în fişierul „Partner.txt”.` |  |

Parametrul din fișier Explicații
Valori posibile (se poate utiliza codul
din SAF-T sau corespondenţa lui):
 Factura de corecţie: 384 sau 2
 cod extern/intern/fiscal partener – se
reglează prin constanta: „Cod pentru
identificare ARTICOLE” vezi Constante
CodFurnizor=F000020 generale > Import date din alte
 cod extern/intern/fiscal partener – se
reglează prin constanta: „Cod pentru
identificare ARTICOLE” vezi Constante
generale > Import date din alte
 Denumire unitate de măsură din
 Simbol gestiune receptie – numai
pentru articole de tip stoc;
 Simbol cont serviciu;
 Observaţii articol.
În cazul unor articole nou apărute în nomenclator, ele vor fi descrise în fişierul
În cazul în care pentru gestiunea de recepţie sunt coduri noi, neintroduse încă în baza
În cazul în care pentru client sunt coduri noi, neintroduse încă în baza de date, ele vor fi
