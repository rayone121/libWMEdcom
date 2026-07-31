# Structura facturilor de ieșire pentru regim special

<sub>Source: `22_Structuri import din alte aplicatii__Facturi de iesire pentru regim special OSS.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2021` |  |
| `LunaLucru=7` |  |
| `Tipdocument=FACTURA REGIM` |  |
| `SPECIAL TVA-OSS` |  |
| `TotalFacturi=2` |  |
| `[Factura_1]` |  |
| `NrDoc=233233` |  |
| `Data=12.07.2021` |  |
| `CodClient=C000020` |  |
| `Moneda=EURO` |  |
| `Curs=4,5842` |  |
| `Scadenta=31.07.2021` |  |
| `TaraTVA=FR` |  |
| `CotaTVA=7` |  |
| `TotalArticole=2` |  |
| `[Items_1]` |  |
| `Item_1=A0000013880;UNIT;1;15;P8201` |  |
| `Item_2=A0000013880;UNIT;1;15;P8201` |  |
| `` | aici se completează cu simbolul |
| `Item_2_TIPCONTABIL=` | tipului contabil; dacă se completează |
| `[Factura_2]` |  |
| `MarcaAgent=123` |  |
| `NrDoc=233234` |  |
| `Data=12.03.2021` |  |
| `CodClient=C004420` |  |
| `Moneda=EURO` |  |
| `Curs=4,5842` |  |
| `Scadenta=11.04.2021` |  |
| `TaraTVA=IT` |  |
| `CotaTVA=3,5` |  |
| `Majorari=12` |  |
| `TotalArticole=1` |  |
| `[Items_2]` |  |
| `` | cod identificare; |
| `` | UM; |
| `` | cantitate; |
| `Item_1=A0635002080;UNIT;1;5;P8201` |  |
| `` | preț; |
| `` | simbol gestiune (numai dacă Kconst |
| `Obligatoriu:` |  |
| `Clienții sunt persoane fizice din UE.` |  |

Structura facturilor de ieșire pentru regim special
TVA – OSS, importate în WinMENTOR
linia aceasta, nu mai ține cont de
constanta cu tipul predefinit
VCGestiuni = DA)
Articolele care sunt servicii electronice nu vor avea linie cu TIPCONTABIL.
