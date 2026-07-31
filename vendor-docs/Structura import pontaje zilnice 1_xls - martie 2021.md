# WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Structura import pontaje zilnice 1_xls - martie 2021.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `WinMENTOR` |  |
| `Structura import pontaje zilnice 1` |  |
| `- Import file type: .xls (Excel)` |  |
| `- Import file name: no restriction` |  |
| `No. Continut` | [Field Field Type Max. |
| `1 An pontaj` | An Numeric 4 |
| `2 Luna pontaj` | Luna Numeric 2 |
| `3 Tip pontaj (avans-1/lichidare-2)` | TipP Numeric 1 |
| `4 Simbol echipa (formatie) - unic` | SimbolE Alfanumeric 10 |
| `5 Marca angajat - unic` | Marca Numeric 10 |
| `6 Total zile lucrate` | TZL Numeric 2 |
| `7 Zi1 - numar de ore lucrate in ziua precizata sau` | Zi1 Numeric |
| `codul intreruperii conform legendei de mai jos` | 2 |
| `8 Zi2 idem` | Zi2 Numeric 2 |
| `9 Zi3 idem` | Zi3 Numeric 2 |
| `10 Zi4 idem` | Zi4 Numeric 2 |
| `11 Zi5 idem` | Zi5 Numeric 2 |
| `12 Zi6 idem` | Zi6 Numeric 2 |
| `13 Zi7 idem` | Zi7 Numeric 2 |
| `14 Zi8 idem` | Zi8 Numeric 2 |
| `15 Zi9 idem` | Zi9 Numeric 2 |
| `16 Zi10 idem` | Zi10 Numeric 2 |
| `17 Zi11 idem` | Zi11 Numeric 2 |
| `18 Zi12 idem` | Zi12 Numeric 2 |
| `19 Zi13 idem` | Zi13 Numeric 2 |
| `20 Zi14 idem` | Zi14 Numeric 2 |
| `21 Zi15 idem` | Zi15 Numeric 2 |
| `22 Zi16 nu se poate completa pentru TipP = 1` | Zi16 Numeric 2 |
| `23 Ore suplimentare I` | TS1 Numeric 2 |
| `24 Ore suplimentare II` | TS2 Numeric 2 |
| `25 Ore noapte` | TN Numeric 2 |
| `26 Identificator echipa (formatie) - unic` | IdE Alfanumeric 13 |
| `27 Zile Suspendate nu se poate completa pentru TipP=1` | Suspendate Numeric 2 |
| `28 Zile Active nu se poate completa pentru TipP=1` | ZileCActiv Numeric 2 |
| `Legenda codificare intreruperi:` |  |
| `(Cod/Simbol/Denumire/Zile Active (A) sau Zile Suspendate (S))` |  |
| ` 25 - Bo - boala obisnuita - S` |  35 - Sc - scolarizare |
| ` 26 - Bp - boala profesionala - S` |  36 - N - nemotivat- S |
| ` 27 - Am - accident de munca - S` |  37 - Prm - program redus maternitate- S |
| ` 28 - M - maternitate- S` |  38 - Prb - program redus boala- S |
| ` 29 - I - invoiri/concedii fara plata- S` |  39 - Ip - invoire platita- A |
| ` 30 - It - intreruperi tehnologice- S` |  40 - Iu - invoire in regim de urgenta - A |
| ` 31 - Oc - obligatii cetatenesti` |  41 - It4 - somaj urgenta 994 - S |
| ` 32 - Om - obligatii militare` |  42 - It5 - somaj urgenta 995 - S |
| ` 33 - Co - concediu odihna` |  43 - It8 - somaj urgenta 998 - S |
| ` 34 - D - delegatie` |  45 - Zs1- Zi redusa cu o ora - A |
| `WinMENTOR` |  |
| ` 46 - Zs2 - Zi redusa cu doua ore - A ` | 51 - Tm2 - Telemunca 2 ore pe zi- A |
| ` 47 - Zs3 - Zi redusa cu trei ore - A ` | 52 - Tm3 - Telemunca 3 ore pe zi- A |
| ` 48 - Zs4 - Zi redusa cu patru ore- A ` | 57 - Tm - Telemunca 8 ore pe zi- A |
| ` 49 - Izs – Toata ziua este redusa- A` |  |
| ` 50 - Tm1 - Telemunca 1 ora pe zi- A` |  |
| `celor existente. Procedura de actualizare este urmatoarea:` |  |
| `Observatii:` |  |

Name] Digits
Incepand cu versiunea de WinMENTOR 892.02 se pot actualiza pontajele fara fi necesara stergerea
1. Se va mofica simbolul intreruperii pentru ziua ce se doreste modificata
2. Se va completa 0 daca intreruperea existenta nu trebuie sa mai apara.
Campurile verzi sunt obligatoriu de completat, cu precizarea ca, pentru identificarea formatiei, daca
„Simbol” formatie nu este unic (din diverse motive), se va utiliza campul „Identificator” de pe
macheta, in acest caz coloana „Simbol” din tabelul .xls fiind necompletata.
