# Structură import facturi clienți

<sub>Source: `22_Structuri import din alte aplicatii__Facturi clienti.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Structură import facturi clienți` |  |
| `în WinMENTOR` |  |
| `[InfoPachet]` |  |
| `AnLucru=2022` |  |
| `LunaLucru=2` |  |
| `` | sau AVIZ EXPEDITIE |
| `Tipdocument= FACTURA IESIRE` |  |
| `TotalFacturi=2` |  |
| `[Factura_1]` |  |
| `NrDoc=17 ` | numărul facturii |
| `ClasificareSAFT=` | • Factura storno: 381 sau 1 |
| `` | se completează în cazul facturilor |
| `CasaDeMarcat=D` |  |
| `NumarBonuri= ` | Se completează cu numărul de |
| `SerieCarnet=AA` |  |
| `MarcaAgent=10` |  |
| `Data=12.02.2022` |  |
| `` | cod extern/intern/fiscal partener |
| `CodClient=C000020` | pentru identificare PARTENER” |
| `Localitate=IASI` |  |
| `` | se completează numai dacă este |
| `TVAINCASARE=D` |  |
| `IDDescarcare=2565252` |  |
| `AutoFactura=N  D= Da; N= Nu` |  |
| `FacturaSimplificata=N  D= Da; N= Nu` |  |
| `1. Numerar` |  |
| `2. Ordin Plata` |  |
| `3. Cec` |  |
| `4. Bilet la Ordin` |  |
| `5. Compensare` |  |
| `TipPlata=X 6. Majorari` |  |
| `7. Online` |  |
| `8. BO avalizat` |  |
| `9. Definit de comun acord` |  |
| `10. Instrument nedefinit` |  |
| `11. Voucher CARD` |  |
| `12. Voucher CEC` |  |
| `EmisClient=D  D= Da; N= Nu` |  |
| `EmisTert=N  D= Da; N= Nu` |  |
| `Stornare=D  D= Da; N= Nu` |  |
| ` TipTVA:` |  |
| `[0;1] – nedefinit;` |  |
| `României;` |  |
| `TipTVA=0` |  |
| `294.2 (a/d);` |  |
| `4 – idem(b/c);` |  |
| `art.311CF;` |  |
| `6 – idem 312CF.` |  |
| `ANULAT=N*  D= Da; N= Nu` |  |
| `TaxareInversa=N` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `NrVoucher=` |  |
| `primit de la client.` |  |
| `Locatie=SediuX` |  |
| `Discount=2` |  |
| `TotalArticole=2` |  |
| `Operat=D sau N` |  |
| `ComandaEF=100` |  |
| `ContractEF=100` |  |
| `ProiectEF=100` |  |
| `unde:` |  |
| `majorare;` |  |
| `[Items_1]` |  |
| `Item_1=A0000013881;LEI;1;21850` |  |
| `` | Pentru articolele de tip serviciu, |
| `Item_1_Ext=P8202;605` |  |
| `` | cod extern/intern articol – se |
| `identificare` | ARTICOLE” vezi |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` |  |
| `` | denumire unitate de măsură din |
| `` | Cantitate; |
| `` | Preț; |
| `` | simbol gestiune livrare – numai |
| `Item_2_TVA=___________________` |  |
| `` | _valoare TVA la nivel de linie |
| `[Factura_2]` |  |
| `NrDoc=18` |  |
| `Data=12.02.2022` |  |
| `CodClient=C000020` |  |
| `Scadenta=12.06.2022` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=D/N ` | D= Da; N= Nu |
| `[Items_2]` |  |
| `Item_1=A0000013881;LEI;1;-21850; P8202` |  |
| `Item_1_StornoAvans=17;AA` |  |
| `` | Stornarea articolului de tip |
| `Observaţii:` |  |
| `Opţiunea de import este în MENTOR> INTERNE> IMPORT DATE DIN ALTE` |  |
| `APLICAŢIE > Facturi ieşire.` |  |
| `descrierea lor.` |  |
| `Pentru clienţi noi se va utiliza fişierul „Partner.txt”.` |  |
| ` Simbol gestiune;` |  |
| ` Discount-ul utilizat la vânzare;` |  |
| ` Preţ înregistrare (pentru articole „valorice”);` |  |
| ` Observaţii articol;` |  |
| ` Preţ achiziţie (pentru articole „valorice”).` |  |
| `terminaţia „_Serii”.` |  |
| `serie pentru fiecare linie „item”.` |  |
| `EX: Item_3=11127;Buc;2.00;49.47;DEP22;;;100` |  |
| `Item_3_Serii=33333;44444` |  |
| `achiziție, atunci linia cu articolul va arăta astfel:` |  |
| `Item_1=45545;Buc;1.00;584;TZL;;;Obs linie;12;` |  |
| `în depozit, adică fără adaos.` |  |

Valori posibile (se poate utiliza codul din
SAF-T sau corespondenţa lui):
• Factura iniţială: 380 sau 0
• Factura de corecţie: 384 sau 2
• Autofactura: 389 sau 3
• Cu factura la bon: 751 sau 4
de tip „InfoCM”
bonuri pentru care s-a emis
factura.
se reglează prin constanta: „Cod
vezi constante generale > import
date din alte aplicații.
cazul
unde x poate lua valorile :
2 – locul livrarii/prestării în afara
3 – intracomunitar scutit conform art.
5 – regim special de scutire
 N = factură obișnuită,
 D = factură cu taxare inversă
Se va completa numărului voucherului
EFactObs=VOUCHER Se va completa cu textul VOUCHER;
CNPDelegat= Daca se trece delegatul pe fisa partener,
se poate pune valoare si la import si se
aduce automat delegatul
se pot defini mai multe scadențe
[Scadente_1] pentru o factură data_scadenta =
valoare;majorari;tipplata
31.03.2022=10.9;10;2  10.9 este valoarea
 10 lei majorări
 TipPlata=2 (ordin plată)
31.07.2022=  valoare la a doua scadență;
se completează simbolul gestiunii
și eventual simbolul contului,
dacă se lucrează cu constanta
„Modificare ct. servicii”.
reglează prin: „Cod pentru
constante generale > import date
din alte aplicații;
WinMENTOR;
pentru articole de tip stoc.
serviciu de pe factura anterioară
(nr. factura 17, serie carnet AA)
În cazul utilizării de noi articole, se va utiliza fişierul „Articole.txt” pentru descrierea lor.
Pentru utilizarea de gestiuni de livrare noi, se va utiliza fişierul „Gestiuni.txt” pentru
La nivel de articol, se mai pot introduce şi următoarele informaţii, în ordinea enumerării
lor şi separate prin „;” (dacă nu există unul dintre acestea trebuie pus „;;”):
Pentru articole cu serii, se poate specifica şi seria pe o linie nouă cu acelaşi item şi
La articolele cu serie pe bucată se validează ca numărul seriilor să corespundă cantităţii
precizate (vezi exemplul de mai jos), iar pentru cele cu serie pe lot să fie specificată o
Dacă se dorește ca livrarea să fie făcută de pe o linie de stoc cu un anume preț de
Explicațiile rămân aceleași ca mai sus, singura mențiune va fi că pe poziția 9 va fi prețul
de achiziție de pe care se dorește livrarea. Modificarea este valabilă doar pentru marfa
