# Structura import bonuri de consum în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Bonuri de consum.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structura import bonuri de consum în WinMENTOR` |  |
| `[InfoPachet]` |  |
| `AnLucru=2023` |  |
| `LunaLucru=4` |  |
| `TipDocument=BON DE CONSUM` |  |
| `TotalBonuri=3` |  |
| `[Bon_1]` |  |
| `NrDoc=14` |  |
| `Data=15.04.2023` |  |
| `GestConsum=C83388` |  Simbol gestiune consumatoare |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – se |
| `` | UM; |
| `` | cantitate; |
| `` | preț; |
| `` | simbol gestiune livrare |
| `Item_1=A0000013880;MT;3.00;0;DEPCENTR;aaa;10000` |  |
| `` | (trebuie completat tipul contabil la nivel de gestiune pentru a se putea efectua livrarea); Tip contabil la nivel de gestiune |
| `` | Pret achizitie (pentru tip |
| `` | Observatii la nivel de articol. |
| `[Bon_2]` |  |
| `NrDoc=15` |  |
| `Data=16.04.2023` |  |
| `GestConsum=C83389` |  |
| `Observatii=` |  |
| `TotalArticole=3` |  |
| `[Items_2]` |  |
| `Item_1=A0000112345;BUC;10.00;25000;DEPCENTR;` |  |
| `` | Seria articolului, dacă articolul |
| `Item_1_Serii=111` |  |
| `Item_2=B45000;BUC;1;15000;DEPCENTR;` | este cu serii |
| `Item_3=B66665800;M;150;3500;DEPCENTR;` |  |
| `[Bon_3]` |  |
| `NrDoc=16` |  |
| `Data=10.02.2023` |  |
| `` | Număr comandă internă; se |
| `ComandaInterna=111` | poate specifica legătura cu o anumită comandă internă, la |
| `GestConsum= C83390` | nivel de document și nu la nivel de linii / articole. |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=d` |  |
| `[Items_3] Item_1=7777;BUC;1;160; C83391;` |  |
| `Item_1_Serii=122` |  |

reglează prin constanta: cod
pentru identificare ARTICOLE
vezi Constante generale >
Import date din alte aplicații.
contabil valoric);
În cazul unor articole nou apărute în nomenclator ele vor fi descrise în fișierul „Articole.txt”.
În cazul în care gestiunea consumatoare sau gestiunea de livrare sunt coduri noi,
neintroduse încă în baza de date, ele vor fi descrise în fișierul „Gestiuni.txt”.
