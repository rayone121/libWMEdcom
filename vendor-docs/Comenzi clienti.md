# Structura import comenzi clienți în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Comenzi clienti.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=12` |  |
| `Tipdocument=COMANDA` |  |
| `TotalComenzi=1` |  |
| `[Comanda_1]` |  |
| `NrDoc=1` |  |
| `Agent=92` | • marcă agent |
| `SimbolCarnet=` |  |
| `Data=20.12.2022` |  |
| `•` | cod extern/intern/fiscal partener – se |
| `CodClient=3869` | identificare PARTENER” vezi |
| `Locatie=sediul 1` |  |
| `Moneda=LEI` |  |
| `TotalArticole=2` |  |
| `Observatii=Urgent` |  |
| `Discount=5` | aplica (fie ca %AdDim fie ca |
| `[Items_1]` |  |
| `•` | cod extern/intern articol – se reglează |
| `Item_1=22222222;Buc;2;4236;-` |  |
| `5;12.01.2022;obs art;` |  |
| `•` | denumire unitate de măsură din |
| `•` | cantitate; |
| `•` | preț; |
| `•` | adaos/discount: se interpretează în |
| `•` | funcție de valoarea constantei generale de funcționare: „Discountul AUTOMAT la ieșiri evidentiat pe”; termen livrare. |
| `Item_2=44444444;Buc;3;4236;-` |  |
| `3;12.01.2022;` |  |
| `„Articole.txt”.` |  |
| `descrise în fişierul „Partner.txt”.` |  |
| `după termen livrare.` |  |

Structura import comenzi clienți în WinMENTOR
reglează prin constanta: „Cod pentru
constante generale > import date din
alte aplicații.
• procentul discountului.
Pe document acest discount se va
%Discount) conform constantei
TipDiscountClient.
prin constanta: „Cod pentru
identificare ARTICOLE” vezi
constante generale > import date din
alte aplicații;
WinMentor;
În cazul unor articole nou apărute în nomenclator ele vor fi descrise în fişierul
În cazul în care pentru client sunt coduri noi, neintroduse încă în baza de date, ele vor fi
În cazul în care se doreşte introducerea observaţiilor la nivel de articol, acestea se scriu
