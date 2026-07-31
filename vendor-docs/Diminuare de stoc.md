# Structuri import diminuare de stoc în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Diminuare de stoc.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=9` |  |
| `Tipdocument=DIMINUARE` |  |
| `DE STOC` |  |
| `TotalDiminuari=1` |  |
| `[PV_1]` |  |
| `Operat=D` |  |
| `NrDoc=61` |  |
| `SimbolCarnet=PV` |  |
| `` | Poate lua valorile: |
| `Operatie=A` |  A – adăugare |
| `Data=02.09.2022` |  S – ștergere. |
| `TotalArticole=1` |  |
| `Observatii=` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – se reglează prin |
| `Item_1=22222222;L;1;D1;12;5 ` | denumire unitate de măsură din WinMENTOR; |
| `` | cantitate; |
| `` | simbol gestiune; |
| `` | preț inregistrare; |
| `` | preț achiziție. |

Structuri import diminuare de stoc în WinMENTOR
constanta: „Cod pentru identificare ARTICOLE”
vezi constante generale > import date din alte
aplicații.
Pentru a evita eroarea „Cont debitor necunoscut...”, este important ca în caracterizarea
contabilă a articolului să existe completat contul aferent rubricii Diminuare stoc.
