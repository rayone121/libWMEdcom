# Structura fișierului necesar importurilor de

<sub>Source: `22_Structuri import din alte aplicatii__Transfer.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=5` |  |
| `TipDocument=TRANSFER` |  |
| `TotalTransferuri=1` |  |
| `[Transfer_1]` |  |
| `NrDoc=9806834` |  |
| `Data=19.04.2022` |  |
| `` | simbol gestiune de |
| `GestDest=APCCART` |  |
| `TotalArticole=5` |  |
| `Operat= D sau N` |  |
| `Anulat = D sau N` |  |
| `Blocat = D sau N` |  |
| `Observatii=hgagfsfkhfff` |  |
| `` | cod extern/intern articol – se |
| `[Items_1]` | alte aplicații ; |
| `` | UM; |
| `` | cantitate; |
| `` | preț; |
| `` | simbol gestiune livrare |
| `Item_2=1086;BUC;1;36606.00;DEPCENTR;` |  |
| `Item_3=51;BUC;1;3200.00;DEPCENTR;` |  |
| `` | tip contabil implicit pentru |
| `Item_3_TipContabil=9` |  |
| `` | tip contabil implicit pentru |
| `Item_3_TipContabilNIR =9` |  |
| `Item_4=1086;BUC;1;36606.00;DEPCENTR;` |  |
| `Item_5=832;BUC;1;36606.00;DEPCENTR;` |  |
| `transferurilor operate.` |  |
| `„Articole.txt”.` |  |

Structura fișierului necesar importurilor de
transferuri în WinMENTOR
destinație
reglează prin constanta:
„Cod pentru identificare
ARTICOLE”; vezi constante
generale > import date din
(sursă)
livrare
NIR
Obligatoriu în WinMENTOR, gestiunile vor avea asociat tipul contabil în cazul
În cazul unor articole nou apărute în nomenclator, ele vor fi descrise în fișierul
În cazul în care pentru gestiunea de livrare și/sau gestiunea sursă sunt coduri noi,
neintroduse încă în baza de date, ele vor fi descrise în fișierul „Gestiuni.txt”.
În cazul în care transferul se face din tip contabil cu metoda de gestiune valorică, după
simbol gestiune sursă se pot specifica: prețul de înregistrare și prețul de achiziție.
