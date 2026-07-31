# Structură import mărire de stoc în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Mariri de stoc.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=9` |  |
| `Tipdocument=MARIRE DE STOC` |  |
| `TotalMariri=1` |  |
| `[PV_1]` |  |
| `Operat=D` |  |
| `NrDoc=61` |  |
| `SimbolCarnet=PV` |  Serie carnet document |
| `Operatie=A` | A - adăugare |
| `Data=02.09.2022` |  |
| `TotalArticole=1` |  |
| `Observatii=` |  |
| `[Items_1]` |  |
| `` | Cod extern/intern articol - se |
| `Item_1=22222222 ;L ;1 ;D1 ;12 ;5` | reglează prin constanta: „Cod pentru identificare ARTICOLE”, vezi constante generale > import date din alte aplicații;  Denumire unitate de măsură din |
| `Item_1_Serii=ABCD4556` |  Seria articolului |
| `Observații:` |  |
| `stoc”.` |  |
| `documentul „Mărire de stoc” va fi neoperat.` |  |
| `implicit”.` |  |

Structură import mărire de stoc în WinMENTOR
 Poate lua valorile:
S - ștergere
WinMENTOR;
 Cantitate;
 Simbol gestiune;
 Preț înregistrare;
 Preț achiziție.
1. Pentru a evita eroarea „Cont creditor necunoscut...”, este important ca în
caracterizarea contabilă a articolului să existe completat contul aferent rubricii „Mărire
2. În cazul în care, pentru articolul care se face mărire nu există nicio intrare,
3. Pentru gestiunea pe care se realizează mărirea este important să fie setat „Cont
