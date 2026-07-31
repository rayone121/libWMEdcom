# Import comandă client format.xml (structura

<sub>Source: `18_WMEdi__IMPORT COMENZI__WMC - Import comanda client DOCPROCESS 04.05.2015.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Import comandă client format.xml (structura
                                DOCXCHANGE / DOCPROCESS)

Identificare parteneri tranzacție
A. Vânzător
Identificarea firmei se face prin cod atribuit de cumpărător vânzătorului, suplimentar
identificării prin GLN.

                                          Fig.1

                                         Fig. 2

Înregistrarea codului GLN se face în tab-ul „Adresa” al machetei de introducere a datelor
firmei (figura 1), iar codul atribuit de client va fi asociat prin intermediul programului WMEdi,
interfața opțiunii „Identificare (vezi figura 2).

B. Cumpărător
Înregistrarea codului GLN pentru identificarea locațiilor partenerilor (clienți/furnizori), se
face în fereastra sediului locației partenerului (vezi figura 3), indiferent de natura codului.

                                               Fig. 3

Identificare articole tranzacție:
Pentru identificarea articolelor comandate se folosesc:
       implicit: codul extern din nomenclatorul de articole al vânzătorului, în format EAN
        (8,12,13) - tag <Item><StandardItemIdentification>;
       suplimentar: codul articolului din catalogul clientului -
        tag<Item><BuyersItemIdentification>;
       opțional: codul intern al vânzătorului - tag <Item><SellersItemIdentification>.
Indiferent de codul utilizat pentru identificarea articolelor, acesta trebuie să fie unic. Codul
suplimentar utilizat se menționează în interfața opțiunii „Identificare”( vezi figura 2).
O vedere de ansamblu asupra articolelor tranzacționate cu un anumit client, asupra
denumirii articolelor și codurilor utilizate se poate obține din COMERCIAL > LISTE >
CLIENȚI > „Informații de la clienți”, configurată ca în figura 4. În felul acesta se pot opera
corecții când este cazul, în nomenclatorul de articole.

Unitatea de măsură a înregistrării articolelor pe comanda clienților trebuie să coincidă cu
unitatea de măsură principală sau secundară a articolelor comandate, în nomenclatorul
vânzătorului, în caz contrar nu se importă comanda.

Observații:
Deoarece structura de import a comenzilor clienți presupune preluarea numărului comenzii
furnizat de client, acest lucru se va opera atribuind constantei generale de funcționare
Comenzi clienți:pe carnete de documente valoarea NU, numărul comenzii client fiind
înregistrat automat în celula „Serie” a comenzii client.
Identificarea în continuare a comenzii se va face fără alte intervenții ale operatorului, atât
pe interfețele tranzacțiilor de stingere a comenzii cât și în listele care fac referire la
comenzi, prin intermediul acestui număr înregistrat în câmpul „Serie”.

                                               Fig. 4

Review-uri document
