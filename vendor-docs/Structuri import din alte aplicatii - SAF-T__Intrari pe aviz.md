# Structura avizelor de intrare importate în

<sub>Source: `22_Structuri import din alte aplicatii__Structuri import din alte aplicatii - SAF-T__Intrari pe aviz.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `intrare.` |  |
| `[InfoPachet]` |  |
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=2` |  |
| `Tipdocument= AVIZ INTRARE` |  |
| `TotalAvize=2` |  |
| `[Aviz_1]` |  |
| `NrDoc=17` |  |
| `ClasificareSAFT=` |  |
| `Data=12.02.2022` |  |
| `CodFurnizor=C000020` | identificare PARTENER” vezi |
| `Moneda=Eur` |  |
| `Curs=4.01` |  |
| `Scadenta=31.03.2022` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `TotalArticole=2` |  |
| `Operat=d sau n` |  |
| `[Items_1]` |  |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` |  |
| `` | cod extern/intern articol – se |
| `Item_1=A0000013880;BUC;1.2;21850;P8201` | reglează prin constanta: „Cod pentru identificare ARTICOLE” vezi Constante generale > Import date din alte aplicaţii;  denumire unitate de măsură din |
| `` | WinMENTOR; cantitate; |
| `` | preţ; |
| `` | simbol gestiune. |
| `Item_1_UM1=20` |  cantitatea în UM alternativă 1 |
| `Item_1_UM2=15` |  cantitatea în UM alternativă 1 |
| `Item_1_TVA=` |  valoarea TVA-ului la nivel de linie |
| `[Aviz_2]` |  |
| `NrDoc=18` |  |
| `ClasificareSAFT=` |  |
| `Data=12.02.2022` |  |
| `CodFurnizor=C000020` |  |
| `Moneda=Eur` |  |
| `Curs=4.97` |  |
| `Scadenta=12.06.2022` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=d/n` |  |
| `[Items_2]` |  |
| `Item_1=A0000013880;BUC;1;33600;P8201` |  |

Structura avizelor de intrare importate în
WinMENTOR
Opţiunea de import este în MENTOR > Interne > Import date din alte aplicaţii > Avize
 cod extern/intern/fiscal partener – se
reglează prin constanta: „Cod pentru
Constante generale > Import date
din alte aplicaţii.
Valori posibile (se poate utiliza codul din
SAF-T sau corespondenţa lui):
 Factura initiala: 380 sau 0
 Factura storno: 381 sau 1
 Factura de corectie: 384 sau 2
 Autofactura: 389 sau 3
În cazul utilizării de noi articole, se va utiliza fişierul „Articole.txt” pentru descrierea lor.
Pentru utilizarea de gestiuni de livrare noi, se va utiliza fişierul „Gestiuni.txt” pentru
descrierea lor. Pentru clienţi noi se va utiliza fişierul „Partner.txt”.
