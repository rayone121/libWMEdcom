# Structura avizelor de ieșire importate în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Avize expeditie iesire.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structura avizelor de ieșire importate în WinMENTOR` |  |
| `Opțiunea de import este în MENTOR> INTERNE> IMPORT DATE DIN ALTE` |  |
| `APLICATII>Avize expeditie iesire` |  |
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=2` |  |
| `Tipdocument=AVIZ EXPEDITIE` |  |
| `TotalAvize=2` |  |
| `[Aviz_1]` |  |
| `NrDoc=17` |  |
| `SAF-T sau corespondenţa lui):` |  |
| `• Factura iniţială:` | 380 sau 0 |
| `• Factura storno:` | 381 sau 1 |
| `• Factura de corecţie:` | 384 sau 2 |
| `• Autofactura:` | 389 sau 3 |
| `• Cu factura la bon:` | 751 sau 4 |
| `Data=12.02.2022` |  |
| `CodClient=C000020` |  |
| `identificare PARTENER”` | vezi |
| `alte aplicații.` |  |
| `Scadenta=31.03.2022` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `TotalArticole=2` |  |
| `Operat=d sau n` |  |
| `[Items_1]` |  |
| `Item_1=A0000013880;BUC;1.2;21850;P8201 prin constanta: „Cod` | pentru |
| `identificare ARTICOLE”` | vezi |
| `alte aplicații;` |  |
| `WinMENTOR;` |  |
| `• cantitate;` |  |
| `• preț;` |  |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` |  |
| `[Aviz _2]` |  |
| `NrDoc=18` |  |
| `Data=12.02.2022` |  |
| `CodClient=C000020` |  |
| `Scadenta=12.06.2022` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=d/n` |  |
| `[Items_2]` |  |
| `Item_1=A0000013880;BUC;1;33600;P8201` |  |
| `descrierea lor.` |  |
| `Pentru clienți noi se va utiliza fișierul „Partner.txt”.` |  |
| `– discount-ul utilizat la vânzare,` |  |
| `– preţ înregistrare (pentru articole „valorice”)` |  |
| `– observaţii articol.` |  |

ClasificareSAFT= Valori posibile (se poate utiliza codul din
• cod extern/intern/fiscal partener – se
reglează prin constanta: „Cod pentru
constante generale > import date din
• cod extern/intern articol – se reglează
constante generale > import date din
• denumire unitate de măsură din
• simbol gestiune livrare – numai
pentru articole de tip stoc.
În cazul utilizării de noi articole se va utiliza fişierul „Articole.txt” pentru descrierea lor.
Pentru utilizarea de gestiuni de livrare noi se va utiliza fişierul „Gestiuni.txt” pentru
La nivel de articol se mai pot introduce şi următoarele informaţii, în ordinea enumerării
lor şi separate prin „;” (dacă nu există unul dintre acestea trebuie pus „;”):
