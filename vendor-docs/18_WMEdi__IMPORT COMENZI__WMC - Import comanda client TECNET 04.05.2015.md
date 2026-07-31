# IMPORT COMANDĂ CLIENT FORMAT .xml

<sub>Source: `18_WMEdi__IMPORT COMENZI__WMC - Import comanda client TECNET 04.05.2015.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
IMPORT COMANDĂ CLIENT FORMAT .xml
                                   (STRUCTURA TECNET)

IDENTIFICARE PARTENERI TRANZACȚIE:
A. VÂNZĂTOR
Identificarea firmei se face prin cod atribuit de cumpărător vânzătorului, suplimentar
identificării prin GLN.

                                         Fig. 1

                                         Fig. 2

Înregistrarea codului GLN se face în tab-ul „Adresa” a machetei de introducere a datelor
firmei (figura 1), iar codul atribuit de client va fi asociat prin intermediul programului WMEdi,
interfața opțiunii „Identificare” (vezi figura 2).

B. CUMPĂRĂTOR
Înregistrarea codului GLN pentru identificarea locațiilor partenerilor (clienți/furnizori), se
face în fereastra sediului locației partenerului (vezi figura 3), indiferent de natura codului.

                                              Fig. 3

IDENTIFICARE ARTICOLE TRANZACȚIE:
Pentru identificarea articolelor comandate se folosește atât codul extern (EAN) al
articolului, cât și codul din catalogul clientului:
       codul extern din nomenclatorul de articole a vânzătorului, în format EAN (8, 12,13) -
        tag<Line-Item><EAN>;
       codul articolului în catalogul clientului - tag <Line-Item><BuyerItemCode>
Indiferent de codul utilizat pentru identificarea articolelor, acesta trebuie să fie unic.
Codul suplimentar utilizat se menționează în interfața opțiunii „Identificare”( vezi figura 2).
O vedere de ansamblu asupra articolelor tranzacționate cu un anumit client, asupra
denumirii articolelor și codurilor utilizate se poate obține din COMERCIAL > LISTE >
CLIENȚI > „Informații de la clienți”, configurată ca în figura 4. În felul acesta, se pot
opera corecții când este cazul, în nomenclatorul de articole.
Unitatea de măsură a înregistrării articolelor pe comanda clienților trebuie să coincidă cu
unitatea de măsură principală sau secundară a articolelor comandate, în nomenclatorul
vânzătorului, în caz contrar, nu se importă comanda.

OBSERVAȚII:
Deoarece structura de import a comenzilor clientți presupune preluarea numărului comenzii
furnizat de client, acest lucru se va opera atribuind constantei generale de funcționare
„Comenzi clienți: pe carnete de documente” valoarea NU, numărul comenzii client fiind
înregistrat automat în celula „Serie” a comenzii client.
Identificarea în continuare a comenzii se va face fără alte intervenții ale operatorului, atât
pe interfețele tranzacțiilor de stingere a comenzii, cât și în listele care fac referire la
comenzi, prin intermediul acestui număr înregistrat în câmpul serie.

                                               Fig. 4

Review-uri document
