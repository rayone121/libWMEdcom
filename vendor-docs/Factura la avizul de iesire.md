# Structură import factură la avizul

<sub>Source: `22_Structuri import din alte aplicatii__Factura la avizul de iesire.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `[InfoPachet]` |  |
| `AnLucru=2023` |  |
| `LunaLucru=5` |  |
| `Tipdocument= FACTURA LA AVIZE IESIRE` |  |
| `TotalFacturi=2` |  |
| `[Factura_1]` |  |
| `NrDoc=17` | • numărul facturii |
| `ClasificareSAFT=` |  |
| `•` | se completează în cazul facturilor |
| `CasaDeMarcat=D` |  |
| `SerieCarnet=AA` | de tip „InfoCM” |
| `MarcaAgent=10` |  |
| `Data=12.05.2023` |  |
| `•` | cod extern/intern/fiscal partener |
| `CodClient=C000020` | se reglează prin constanta: „Cod pentru identificare PARTENER” |
| `Localitate=IASI` | vezi constante generale > import date din alte aplicații. |
| `•` | se completează numai dacă este |
| `TVAINCASARE=D` |  |
| `AutoFactura=N` | cazul • D= Da; N= Nu |
| `FacturaSimplificata=N` | • D= Da; N= Nu |
| `TipPlata=X` |  |
| `2. Ordin Plata` |  |
| `3. Cec` |  |
| `4. Bilet la Ordin` |  |
| `5. Compensare` |  |
| `6. Majorari` |  |
| `7. Online` |  |
| `8. BO avalizat` |  |
| `9. Definit de comun acord` |  |
| `10. Instrument nedefinit` |  |
| `11. Voucher CARD` |  |
| `12. Voucher CEC` |  |
| `EmisClient=D • D= Da; N= Nu` |  |
| `EmisTert=N • D= Da; N= Nu` |  |
| `Stornare=D • D= Da; N= Nu` |  |
| `• TipTVA:` |  |
| `[0;1] – nedefinit;` |  |
| `României;` |  |
| `TipTVA=0` |  |
| `294.2 (a/d);` |  |
| `4 – idem(b/c);` |  |
| `5 – regim special de scutire` |  |
| `art.311CF;` |  |
| `6 – idem 312CF.` |  |
| `ANULAT=N* • D= Da; N= Nu` |  |
| `• N = factură obișnuită,` |  |
| `TaxareInversa=N` |  |
| `Majorari=12.45` |  |
| `Observatii=hgdhgfhgfhgf` |  |
| `NrVoucher=` |  |
| `primit de la client.` |  |
| `EFactObs=VOUCHER Se va completa cu textul VOUCHER;` |  |
| `Locatie=SediuX` |  |
| `Discount=2` |  |
| `TotalArticole=2` |  |
| `Operat=D sau N` |  |
| `ComandaEF=100` |  |
| `ContractEF=100` |  |
| `ProiectEF=100` |  |
| `[Scadente_1]` |  |
| `31.05.2023=10.9;10;2` | • 10.9 este valoarea |
| `31.06.2023=` | • valoare la a doua scadență; |
| `[Items_1]` |  |
| `Item_1=A0000013881;LEI;1;21850` |  |
| `•` | Pentru articolele de tip serviciu, |
| `Item_1_Ext=P8202;605` |  |
| `•` | cod extern/intern articol – se |
| `Item_2=A0000013880;BUC;1.5;23467;P8201` |  |
| `•` | denumire unitate de măsură din |
| `•` | Cantitate; |
| `•` | Preț; |
| `•` | simbol gestiune livrare – numai |
| `Item_2_TVA=___________________` |  |
| `•` | _valoare TVA la nivel de linie |
| `Item_2_NrAviz = 9 •` | Numărul avizului stins |
| `[Factura_2]` |  |
| `NrDoc=18` |  |
| `Data=12.05.2023` |  |
| `CodClient=C000020` |  |
| `Scadenta=12.06.2023` |  |
| `Majorari=` |  |
| `Observatii=` |  |
| `TotalArticole=1` |  |
| `Operat=D/N •` | D= Da; N= Nu |
| `[Items_2]` |  |
| `Item_1=A0000013881;LEI;1;-21850; P8202` |  |
| `Item_1_StornoAvans=17;AA •` | Stornarea articolului de tip |
| `Observaţii:` | (nr. factura 17, serie carnet AA) |
| `APLICAŢIE > Facturi la avize.` |  |
| `descrierea lor.` |  |
| `Pentru clienţi noi se va utiliza fişierul „Partner.txt”.` |  |
| `• Simbol gestiune;` |  |
| `• Discount-ul utilizat la vânzare;` |  |
| `• Preţ înregistrare (pentru articole „valorice”);` |  |
| `• Observaţii articol;` |  |
| `• Preţ achiziţie (pentru articole „valorice”).` |  |
| `terminaţia „_Serii”.` |  |
| `serie pentru fiecare linie „item”.` |  |
| `EX: Item_3=11127;Buc;2.00;49.47;DEP22;;;100` |  |
| `Item_3_Serii=33333;44444` |  |
| `achiziție, atunci linia cu articolul va arăta astfel:` |  |
| `Item_1=45545;Buc;1.00;584;TZL;;;Obs linie;12;` |  |
| `în depozit, adică fără adaos.` |  |

Structură import factură la avizul
de ieșire WinMENTOR
Valori posibile (se poate utiliza codul din
SAF-T sau corespondenţa lui):
• Factura iniţială: 380 sau 0
• Factura storno: 381 sau 1
• Factura de corecţie: 384 sau 2
• Autofactura: 389 sau 3
unde x poate lua valorile :
1. Numerar
2 – locul livrarii/prestării în afara
3 – intracomunitar scutit conform art.
• D = factură cu taxare inversă
Se va completa numărului voucherului
se pot defini mai multe scadențe
pentru o factură data_scadenta =
valoare;majorari;tipplata
unde:
• 10 lei majorări
• TipPlata=2 (ordin plată)
majorare;
se completează simbolul gestiunii
și eventual simbolul contului,
dacă se lucrează cu constanta
„Modificare ct. servicii”.
reglează prin: „Cod pentru
identificare ARTICOLE” vezi
constante generale > import date
din alte aplicații;
WinMENTOR;
pentru articole de tip stoc.
serviciu de pe factura anterioară
Opţiunea de import este în MENTOR> INTERNE> IMPORT DATE DIN ALTE
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
