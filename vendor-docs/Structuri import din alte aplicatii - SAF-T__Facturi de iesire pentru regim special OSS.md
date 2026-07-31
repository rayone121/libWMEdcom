# Structura facturilor de ieşire pentru regim special

<sub>Source: `22_Structuri import din alte aplicatii__Structuri import din alte aplicatii - SAF-T__Facturi de iesire pentru regim special OSS.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=3` |  |
| `Tipdocument=` |  |
| `TotalFacturi=2` |  |
| `[Factura_1]` |  |
| `NrDoc=233233` |  |
| `ClasificareSAFT=` |  |
| `Data=12.03.2022` |  |
| `CodClient=C000020` |  |
| `Moneda=EURO` |  |
| `Curs=4,9842` |  |
| `Scadenta=31.04.2022` |  |
| `TaraTVA=FR` |  |
| `CotaTVA=7` |  |
| `TotalArticole=2` |  |
| `[Items_1]` |  |
| `Item_1=A0000013880;UNIT;1;15;P8201` |  |
| `Item_2=A0000013880;UNIT;1;15;P8201` |  |
| `Item_2_TIPCONTABIL=` | • aici se completeaza cu simbolul tipului |
| `[Factura_2]` |  |
| `MarcaAgent=123` |  |
| `NrDoc=233234` |  |
| `ClasificareSAFT=` |  |
| `Data=12.03.2022` |  |
| `CodClient=C004420` |  |
| `Moneda=EURO` |  |
| `Curs=4,9842` |  |
| `Scadenta=11.04.2022` |  |
| `TaraTVA=IT` |  |
| `CotaTVA=3,5` |  |
| `Majorari=12` |  |
| `TotalArticole=1` |  |
| `[Items_2]` |  |
| `` | Item_x=cod identificare; |
| `` | UM; |
| `Item_1=A0635002080;UNIT;1;5;P8201 ` | cantitate; |
| `` | pret; |
| `` | simbol gestiune (numai daca Kconst |
| `Obligatoriu:` |  |
| `Clienţii sunt persoane fizice din UE.` |  |

Structura facturilor de ieşire pentru regim special
TVA- OSS, importate în WinMENTOR
Factura regim special TVA-OSS
Valori posibile (se poate utiliza codul din
SAF-T sau corespondenţa lui):
• Factura iniţială: 380 sau 0
• Factura storno: 381 sau 1
• Factura de corecţie: 384 sau 2
• Autofactura: 389 sau 3
contabil. Daca se completeaza linia aceasta,
nu mai tine cont de cst. cu tipul predefinit.
VCGestiuni = DA).
Articolele care sunt servicii nu vor avea completată linia cu TIPCONTABIL.
