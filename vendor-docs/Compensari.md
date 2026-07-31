# Structuri import compensări în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Compensari.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=09` |  |
| `TipDocument=COMPENSARE` |  |
| `TotalCompensari=1` |  |
| `[Compensare_1]` |  |
| `NrDoc=1` |  |
| `Zi=17` |  |
| `ObsD=fff ` | Observație cont debit |
| `ObsC=54 ` | Observație cont credit |
| `[Debit_1] ` | Secțiunea debit pentru compensarea 1 |
| `` | cod extern/intern articol – se reglează prin |
| `CodPartener=111` |  |
| `NrLinii=1` |  |
| `` | Tip Document: |
| `Linia_1=200;2000;CH;100` |  1 = Factura de intrare (furnizori);  5 = Factura de intrare din import (DVI); 19 = Factura de ieșire (clienți);  22 = Facturi ieșiri valută;  201 = Plata în avans 200 = Încasare în avans. |
| `` | Număr Document |
| `` | Serie Document |
| `` | Valoare compensată. Poate fi și negativă și se |
| `[Credit_1] ` | Secțiunea credit pentru compensarea 1 |
| `CodPartener=111` |  |
| `NrLinii=1` |  |
| `Linia_1=201;1000;CH;100` |  |
| `[Compensare_2]` |  |
| `NrDoc=2` |  |
| `Zi=17` |  |
| `NrLinii=1` |  |
| `` | Tip document va fi: -1; |
| `` | Contul; |
| `Linia_1=-1;409;4.97;176;EUR` |  |
| `` | Cursul valutar |
| `` | Valoarea |
| `` | Simbol valută |
| `[Credit_2] ` | Secțiunea credit pentru compensarea 2 |
| `CodPartener=111` |  |
| `NrLinii=1` |  |
| `Linia_1=19;13;FF;119` |  |
| `vor fi descrise în fișierul „Partner.txt”.` |  |

Structuri import compensări în WinMENTOR
constanta: „Cod pentru identificare
PARTENER” vezi constante generale > import
date din alte aplicații;
exprimă în moneda în care a fost generată
obligația de plată/încasare
Pentru cazul direct pe cont:
În cazul în care pentru partener sunt coduri noi, neintroduse încă în baza de date, ele
