# Structura import comenzi furnizori în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Comenzi furnizori.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=12` |  |
| `Tipdocument=COMANDA FURNIZOR` |  |
| `TotalComenzi=1` |  |
| `[Comanda_1]` |  |
| `NrDoc=1` |  |
| `SimbolCarnet=` |  |
| `Data=20.12.2022` |  |
| `` | cod extern/intern/fiscal partener – se |
| `CodFurnizor=3869` | identificare PARTENER” vezi |
| `Locatie=sediul 1` |  |
| `Moneda=LEI` |  |
| `TotalArticole=2` |  |
| `Observatii=Urgent` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – se reglează |
| `` | denumire unitate de măsură din |
| `Item_1=22222222;Buc;2;4236;-` |  |
| `` | cantitate; |
| `5;12.01.2022;obs art;` |  |
| `` | preț; |
| `` | adaos/discount: se interpretează în |
| `` | termen livrare; |
| `` | observații articol. |
| `Item_2=44444444;Buc;3;4236;-` |  |
| `3;12.01.2022;` |  |
| `„Articole.txt”.` |  |
| `fi descrise în fişierul „Partner.txt”.` |  |
| `după termen livrare.` |  |

Structura import comenzi furnizori în WinMENTOR
reglează prin constanta: „Cod pentru
constante generale > import date din
alte aplicații
prin constanta: „Cod pentru identificare
ARTICOLE” vezi constante generale >
import date din alte aplicații;
WinMentor;
funcție de valoarea constantei
generale de funcționare: „Discountul
AUTOMAT la ieșiri evidentiat pe”;
În cazul unor articole nou apărute în nomenclator ele vor fi descrise în fişierul
În cazul în care pentru furnizor sunt coduri noi, neintroduse încă în baza de date, ele vor
În cazul în care se doreşte introducerea observaţiilor la nivel de articol, acestea se scriu
