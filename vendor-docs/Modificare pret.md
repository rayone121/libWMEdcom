# Structură import modificare preț în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Modificare pret.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=11` |  |
| `Tipdocument=MODIFICARE PRET` |  |
| `TotalModifPret=1` |  |
| `LogOn=Master` |  |
| `[PV_1]` |  |
| `Operat=D` |  |
| `NrDoc=59` |  |
| `SimbolCarnet=PV •` | serie carnet documente |
| `•` | Poate lua valorile: |
| `Operatie=A` |  |
| `Data=14.11.2022` | ➢ A – adăugare ➢ S – ștergere |
| `TotalArticole=1` |  |
| `•` | Poate lua valorile: |
| `Operat=D` |  |
| `Observatii=` |  |
| `[Items_1]` |  |
| `•` | cod extern/intern articol – se |
| `•` | denumire unitate de măsură |
| `Item_1=11126358;BUC;1;11.00;DC;100;23` | din WinMENTOR; |
| `•` | cantitate; |
| `•` | preț înregistrare; |
| `•` | simbol gestiune. |
| `•` | Preț de înregistrare pentru |
| `•` | Preț de achiziție pentru |
| `Item_1_TipContabil=MAGVAL` |  |
| `Observații:` |  |
| `Rotunjire Preț.` |  |
| `implicit”.` |  |

Structură import modificare preț în WinMENTOR
➢ D – DA
➢ N – NU
reglează prin constanta:
„Cod pentru identificare
ARTICOLE” (vezi constante
generale > import date din
alte aplicații);
livrări valorice;
livrări valorice;
Tip contabil impus de
pachetul de date
1. Pentru a evita eroarea „Cont creditor necunoscut...”, este important ca în
caracterizarea contabilă a articolului să existe completat contul aferent rubricii Dif.
2. Pentru gestiunea pe care se realizează modificarea este important să fie setat „Cont
