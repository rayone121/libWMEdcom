# Structură facturi de ieșire importate în WinMENTOR

<sub>Source: `22_Structuri import din alte aplicatii__Facturi clienti storno.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structură facturi de ieșire importate în WinMENTOR` |  |
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=2` |  |
| `TotalFacturi=2` |  |
| `[Factura_1]` |  |
| `NrDoc=17` |  |
| ` se` | completează în cazul |
| `CasaDeMarcat=D` |  |
| `SerieCarnet= AA` |  |
| `MarcaAgent=10` |  |
| `Data=12.02.2022` |  |
| `constanta:` | „Cod pentru |
| `CodClient=C000020` |  |
| `Localitate=Iasi` |  |
| `Scadenta=31.03.2022` |  |
| `TaxareInversa=N` |  |
| `TipTVA=2` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `Locatie=SediuX` |  |
| `Discount=2` |  |
| `TotalArticole=2` |  |
| `Operat=d sau n` |  |
| `[Items_1]` |  |
| `Item_1=A0000013881;LEI;1;21850;` |  |
| `` | Pentru articolele de tip serviciu |
| `se` | completează simbolul |
| `Item_1_Ext=P8202;605` |  |
| `` | cod extern/ intern / fiscal |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` | date din alte aplicații; |
| `` | denumire unitate de măsură din |
| `` | cantitate; |
| `` | pret; |
| `` | simbol gestiune livrare - numai |
| `Item_1_TVA= ` | valoare TVA la nivel de linie |
| `[Factura_2]` |  |
| `NrDoc=18` |  |
| `Data=12.02.2022` |  |
| `CodClient=C000020` |  |
| `Scadenta=12.06.2022` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=d/n` |  |
| `[Items_2]` |  |
| `Item_1=A0000013881;LEI;1;-21850; P8202` |  |
| `` | Stornarea articolului de tip |
| `Item_1_StornoAvans=17;AA` | serviciu de pe factură anterioară |
| `descrierea lor.` |  |
| `Pentru clienți noi se va utiliza fișierul „Partner.txt”.` |  |

Opțiunea de import este în MENTOR> INTERNE> IMPORT DATE DIN ALTE
APLICATIE >Facturi iesire
Tipdocument=FACTURA IESIRE  sau AVIZ EXPEDITIE
facturilor de tip „InfoCM”
 cod extern/ intern / fiscal
partener – se reglează prin
identificare PARTENER” vezi
constante generale > import
date din alte aplicații.
 N = factură obișnuită;
 D = factură cu taxare inversă
 1 …5 = Particularități TVA în
ordinea în care apar în machetă
gestiunii și, eventual, simbolul
contului dacă se lucrează cu
constanta „Modificare ct.
servicii”
partener – se reglează prin
constanta: „Cod pentru
identificare PARTENER” vezi
constante generale > import
WinMENTOR;
pentru articole de tip stoc.
(nr. factură 17, serie carnet AA)
În cazul utilizării de noi articole, se va utiliza fișierul „Articole.txt” pentru descrierea lor.
Pentru utilizarea de gestiuni de livrare noi, se va utiliza fișierul „Gestiuni.txt” pentru
