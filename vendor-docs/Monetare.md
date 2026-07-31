# Structura fișierului necesar importurilor de

<sub>Source: `22_Structuri import din alte aplicatii__Monetare.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=5` |  |
| `TipDocument=MONETAR` |  |
| `TotalMonetare=1` |  |
| `[Monetar_1]` |  |
| `Operat=D` |  |
| `NrDoc=59` |  |
| ` Seria` | carnetului de |
| `SimbolCarnet=M` |  |
| `documente` |  |
| `Operatie=A` |  A – adăugare, |
| `marcat:` |  |
| `CasaDeMarcat=D` |  |
| `NumarBonuri=2` |  |
| `Data=20.05.2022` |  |
| `Casa=CasaLei` |  |
| `TotalArticole=5` |  |
| `CEC=0` |  |
| `CARD=0` |  |
| `BONVALORIC=0` |  |
| `Observatii=hgagfsfkhfff` |  |
| `MarcaAgent` |  |
| `Discount=3.1` |  |
| `TVADiscount=1.1` |  |
| `[Items_1]` |  |
| `` | cod extern/intern articol – |
| `Item_1=51;BUC;1;3200.00;DEPCENTR` | alte aplicații; |
| `` | unitate de măsură; |
| `` | cantitate; |
| `` | preț; |
| `` | simbol gestiune livrare. |
| `Item_2=1086;BUC;1;36606.00;DEPCENTR;` |  |
| `Item_3=51;BUC;1;3200.00;DEPCENTR;` |  |
| `Item_4=1086;BUC;1;36606.00;DEPCENTR;` |  |
| `Item_5=832;BUC;1;36606.00;DEPCENTR;` |  |
| `[Monetar_1_Facturi]` |  |
| `` | cod extern/intern partener – |
| `Factura_1=222;1;F;` |  |
| `` | numărul facturii; |
| `` | seria facturii. |
| `[MONETAR_1_DetaliiDiscount]` |  |
| `TVA_19=11.9;1.9` |  |
| `TVA_9=24;1.98` |  |
| `` | valoare (cu TVA); |
| `TVA_5=26.25;1.25` |  |
| `` | din care TVA. |
| `descrise în fișierul „Articole.txt”.` |  |

Structura fișierului necesar importurilor de
transferuri în WinMENTOR
 Poate lua valorile:
 S – ștergere.
 Preluat de la casa de
 D – Da,
 N – Nu.
 Numărul de bonuri de la CM
preluat pe monetarul curent
se reglează prin constanta:
„Cod pentru identificare
ARTICOLE” vezi constante
generale > import date din
se reglează prin constanta:
„Cod de identificare
Parteneri” vezi constante
generale > import date din
alte aplicații;
Obligatoriu în WinMENTOR gestiunile vor avea asociat tipul contabil în cazul
monetarelor operate. În cazul unor articole nou apărute în nomenclator, ele vor fi
În cazul articolelor gesionate prin metoda de gestiune valorică, după simbolul gestiunii
de livrare se pot specifica: prețul de înregistrare, preț achiziție.
