# Structura import articole noi [ARTICOLE.txt] în

<sub>Source: `22_Structuri import din alte aplicatii__Articole noi.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Structura import articole noi [ARTICOLE.txt] în
                                              WinMENTOR

          Parametrul din fișier                               Explicații

               [InfoPachet]

[ArticoleNoi_P1200069195018]
Denumire=TELA SDRAIO S.S.
Serviciu=D
ContServiciu=707

[ArticoleNoi_A0000013880]
Denumire=BORDO IN TESSUTO MM.
Serviciu=N
ContServiciu=
                                           Se va completa cu cod extern, cod fiscal
IDProducator=                              sau cod intern în funcție de constanta de
                                           căutare parteneri.
GestiuneImplicita=
Clasa=
PretVanzare=

TVAInclus=D
                                           0 - normal

TipEFACT=                                  1 – risc fiscal ridicat
                                           2 – construcție nouă

                                            cod extern/intern articol – se reglează
                                             prin constanta: „Cod pentru identificare
[ArticoleNoi_A0635002080]                    ARTICOLE” vezi Constante generale >
                                             Import date din alte aplicaţii.

Denumire=NASTROPOLIPROP.SPIG
Serviciu=D

ContServiciu=711

GestiuneImplicita= simbol gestiune

Clasa=

PretVanzare=

TVAInclus=

ProcTVA=
ZeroCuDeducere=                   N – fără deducere; D – cu deducere.
                                  Poate lua una din cele trei valori:
                                  1 – nu are serie;
TipSerie=
                                  2 – serie pe bucată ;.
                                  3 – serie pe lot.
TipContabil=simbol tip contabil
                                        Este echivalentul SAF-T pentru cod
CodVamal=
                                         NC

Greutate=
Volum=
Suprafata=
CantImplicita=
CodCPV=
CodCNAS=
ZilePlata=
StocMinim=
PretReferinta=
Descriere=

[ArticoleNoi_A1092999899]
Denumire=NYLONH.150 BLUNOTT
Serviciu=D
ContServiciu=768

GestiuneImplicita=
Clasa=
PretVanzare=
TVAInclus=

[ArticoleNoi_D5001110030]

Denumire=ETICH.TRANSFERT CHIC

Serviciu= N

ContServiciu=
GestiuneImplicita=

Clasa=

PretVanzare=

TVAInclus=

Fişierul poate conţine toate articolele din aplicaţia sursă, consultarea acestui fişier
făcându-se numai în cazul în care codul de articol din fişierul de tranzacţii nu este găsit
în nomenclatorul de articole din WinMENTOR.

Fişierul poate fi importat doar la importul unui document (factură intrare, factură ieşire,
transfer, comandă etc.) în care se face referire la articolele prezente în document.
Unitatea de măsură implicită va fi preluată din tranzacţia importată (intrări, ieşiri
etc.).
Se pot adăuga câmpuri noi în această structură, dacă solicitaţi acest lucru, în
măsura în care există câmpurile respective în WinMENTOR.
