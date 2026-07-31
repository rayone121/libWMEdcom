# Structura avizelor de ieșire către subunități importate

<sub>Source: `22_Structuri import din alte aplicatii__Avize catre subunitati-iesiri.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=2` |  |
| `Tipdocument=AVIZ IESIRE SUBUNITATI` |  |
| `TotalAvize=2` |  |
| `[Aviz_1]` |  |
| `NrDoc=17` |  |
| `Data=12.02.2022` |  |
| `Subunitatea=Denumirea subunitatii` |  |
| `Scadenta=31.03.2022` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `TotalArticole=2` |  |
| `Operat=d sau n` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – se |
| `` | denumire unitate de măsură din |
| `Item_1=A0000013880;BUC;1.2;21850;P8201` |  |
| `` | cantitate; |
| `` | preț; |
| `` | simbol gestiune livrare – numai |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` |  |
| `[Aviz _2]` |  |
| `NrDoc=18` |  |
| `Data=12.02.2022` |  |
| `Subunitatea=Denumirea subunitatii` |  |
| `Scadenta=12.06.2022` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=d/n` |  |
| `[Items_2]` |  |
| `Item_1=A0000013880;BUC;1;33600;P8201` |  |
| `descrierea lor.` |  |
| `– discount-ul utilizat la vânzare,` |  |
| `– observaţii articol.` |  |

Structura avizelor de ieșire către subunități importate
în WinMENTOR
Opțiunea de import este în MENTOR> INTERNE> IMPORT DATE DIN ALTE
APLICATII>Avize iesire subunitati
reglează prin constanta: „Cod
pentru identificare ARTICOLE”
vezi constante generale > import
date din alte aplicații;
WinMENTOR;
pentru articole de tip stoc.
În cazul utilizării de noi articole se va utiliza fişierul „Articole.txt” pentru descrierea lor.
Pentru utilizarea de gestiuni de livrare noi se va utiliza fişierul „Gestiuni.txt” pentru
La nivel de articol se mai pot introduce şi următoarele informaţii, în ordinea enumerării
lor şi separate prin „;” (dacă nu există unul dintre acestea trebuie pus „;”):
– preţ înregistrare (pentru articole „valorice”)
