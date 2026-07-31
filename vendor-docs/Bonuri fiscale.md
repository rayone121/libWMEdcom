# Bonuri fiscale intrare

<sub>Source: `22_Structuri import din alte aplicatii__Bonuri fiscale.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Bonuri fiscale intrare` |  |
| `[InfoPachet]` |  |
| `AnLucru=2006` |  |
| `LunaLucru=3` |  |
| `TipDocument=BON FISCAL INTRARE` |  |
| `TotalBonuri=3` |  |
| `Numarul bonului` |  |
| `[Bon_1]` | Data bonului |
| `NrDoc=5144257` |  |
| `Data=02.03.2006` |  |
| `Discount=11.8300` |  |
| `[Items_1]` |  |
| `Item_1=11126;Buc;1.00;489.47;Falticeni;0;;603.77;` | Corespondenta dintre numarul de articole |
| `Item_2=11123;Buc;1.00;37.40;Falticeni;;;51.20;` | si numarul de descrieri de articole este |
| `Item_1_TVA=------valoarea TVA-ului la nivel de linie` | validat la preluare. |
| `[Bon_2]` |  |
| `Data=02.03.2006` |  |
| `CodFurnizor=R 717359` |  |
| `TotalArticole=1 Simbol gestiune` |  |
| `Discount=15.0100` | Discount linie |
| `[Items_2]` |  |
| `Item_1=11123;Kg;1.00; 49.56;Falticeni;0;;-60.48;` | Pret inregistrare pt. articolele ce au |
| `[Bon_3]` | tip contabil implicit cu adaos |
| `CodFurnizor=R 4851409` |  |
| `TotalArticole=1` | Termen de garantie |
| `Discount=18.0000` |  |
| `[Items_3]` |  |
| `Item_1=11124;Lei;1.00;12.00;;;624;;15.12.2008;` |  |

CodFurnizor=R 717359 Codul utilizat pt. identificarea furnizorului (se precizeaza
TotalArticole=2 tipul lui in constante generale: "Import date din alte aplicatii".
NrDoc=6 Asigura corespondenta dintre antetul bonului si continutul sau.
NrDoc=7 Cod articol (intern/extern) conform constantei de import date din
Data=03.03.2006 alte aplicatii: cod identificare articole
{ 9 campuri separate prin ";" } Simbol cont pt.articole de tip serviciu.
