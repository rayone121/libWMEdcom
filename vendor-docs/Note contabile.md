# Structură import fișier note contabile în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Note contabile.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=9` |  |
| `Tipdocument=NOTA CONTABILA` |  |
| `TotalNote=2` |  |
| `[Nota_1]` |  |
| `NrDoc=151520` |  |
| `Data=29.11.2022` |  |
| `Moneda=EU` |  |
| `Curs=4.95` |  |
| `Observatii=Observatie` |  |
| `` | cod extern/intern/fiscal partener |
| `Partener=RO 1216789` |  |
| `TotalLinii=2` |  |
| `` | Pentru ca TVA să apară în |
| `Jurnal=Vanzari` | „Vânzări”, respectiv „Cumpărări” |
| `Baza=1 ` | Index linie bază impozitare |
| `TVA=2 ` | Index linie TVA |
| `[Items_1]` |  |
| `` | Setare bifa partener pentru |
| `Item_1=P;401;X;593;1000;DC; Obs linii ` |  P – este pusă bifa  X – nu este pusă bifa Contul de debit din planul de |
| `` | Setare bifă partener pentru |
| `Item_2=X;401;X;593;300;DC; Obs linii` |  P – este pusă bifa  X – nu este pusă bifa.  Contul de credit din planul de conturi;  Valoarea liniei din nota contabilă;  Simbol gestiune;  Observații linii. |
| `[Nota_2]` |  |
| `NrDoc=151521` |  |
| `Data=30.11.2022` |  |
| `Moneda=Lei` |  |
| `Observatii=` |  |
| `Partener=RO 1216789` |  |
| `TotalLinii=1` |  |
| `[Items_2]` |  |
| `Item_1=P;409.01;X;401;700` |  |
| `fi descrise în fișierul „Partner.txt”.` |  |

Structură import fișier note contabile în WinMENTOR
- se reglează prin constanta:
„Cod pentru identificare
PARTENER” vezi constante
generale > import date din alte
aplicații.
jurnalul de vânzari se trece
ca să apară în jurnalul de
cumpărări.
contul de debit:
conturi;
contul de credit:
În cazul în care pentru partener sunt coduri noi, neintroduse încă în baza de ate, ele vor
