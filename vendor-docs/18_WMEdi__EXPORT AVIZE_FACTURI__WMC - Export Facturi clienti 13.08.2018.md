# EXPORT FACTURI CLIENȚI FORMAT .xml

<sub>Source: `18_WMEdi__EXPORT AVIZE_FACTURI__WMC - Export Facturi clienti 13.08.2018.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
EXPORT FACTURI CLIENȚI FORMAT .xml

IDENTIFICAREA PARTENERILOR TRANZACȚIEI DE IEȘIRE:
Se face selectând din WMEdi opțiunea Identificare >> Configurare import/export pe
sedii clienți (vezi figura 1).

În fereastra afișată se vor configura elementele ce intervin în importul /exportul datelor
dintre firme la nivel de format selectat (codificare sediu livrare la client, codificare firmă
furnizoare în evidența clientului, metoda identificare articole la client).
Selecția sediilor clientului se face de pe plusul verde al grilei (vezi figura 2).

                                             Fig. 2
Completarea GLN-ului (Global Location Number) sau Id-ului sediului de livrare se poate
face direct în grilă, în cazul în care nu au fost deja specificate în adresa sediului de livrare
din interfața nomenclatorului de parteneri. La salvare, GLN-ul/Id-ul se va actualiza în
nomenclatorul partenerilor.
Elementele de identificare ce țin de furnizor (firma) se pot completa automat pe toate
rândurile din grilă, specificând acest lucru în coloana de marcare, folosind butonul „+”.
Liniile grilei se pot deselecta folosind butonul „-“ (vezi figura 3).

În final, grila arată ca în figura 4.
Nu uitați să salvați configurările, în caz contrar, o veți lua de la capăt!

                                              Fig. 4

                                              Fig. 5
După configurarea tipului de export, facturile înregistrate în mod obișnuit vor putea fi
exportate în formatul specificat, selectând opțiunea „Export WMEdi”, după salvarea lor
(vezi figura 5).

ELEMENTE SPECIFICE FORMATELOR UTILIZATE
1. SMART INVOICE
(beneficiari: CARREFOUR...)
Parametri configurare WMEdi
       Se va specifica GLN-ul sediilor de livrare ale clientului și a sediului social dacă nu au
        fost deja încărcate.
       Identificarea firmei furnizorului: Tip Id.Firma - va fi specific, iar valoarea
        identificatorului va fi cea primită de la client.
       Codul articolului utilizat pentru identificare este codul extern al articolului (dacă este
        EAN) în caz contrar, codul de catalog al clientului.
Mod de lucru/Date obligatorii
       Factura nu poate conține articole ce provin de pe mai multe comenzi.
       În CONSTANTE GENERALE > FUNCȚIONARE > SECȚIUNEA: IEȘIRI DIN STOC
        > Identificare Taxa Verde se va selecta taxa corespunzătoare, în cazul în care
        articolele pe care le facturați conțin această taxă.
       Codul GLN este completat pe baza datelor primite de la client atât pentru sediul
        social, cât și pentru sediile de livrare.
       Cantitatea articolelor pe factură se va specifica cu maximum 3 zecimale.
       Toate mărimile referitoare la preț, taxa verde, TVA, discount corespunzătoare
        articolelor se vor înscrie pe factură cu maximum 4 zecimale.
       La facturare se va folosi discount-ul pe coloană.
       Înregistrarea delegatului este obligatorie.

2. DOCXCHANGE
(beneficiari: CARREFOUR, AUCHAN,...)
Parametri configurare WMEdi
       Se va specifica GLN-ul sediilor de livrare ale clientului și a sediului social dacă nu au
        fost deja încărcate.
       Identificarea firmei furnizorului: Tip Id.Firma - va fi specific, iar valoarea
        identificatorului va fi cea primită de la client.
       Codul articolului utilizat pentru identificare este codul de catalog al clientului.
Atenție!!!
Codul de catalog CARREFOUR începe cu caracterul „C”.
Mod de lucru/Date obligatorii
       Obligatoriu, factura este generată pe baza unei comenzi client.
       Factura nu poate conține articole ce provin de pe mai multe comenzi.

        Codul fiscal al firmei trebuie să fie fără spații între indicatorul fiscal și codul numeric
         (RO88888).
        Să fie completat contul bancar implicit al firmei furnizoare în CONSTANTE
         UTILIZATOR > CONSTANTE DE OPERARE > SECȚIUNEA: TRANZACȚII
         TREZORERIE > Bancă implicită. Fiind o constantă utilizator trebuie să fie
         completată pentru toți operatorii care facturează.
        În CONSTANTE GENERALE > FUNCȚIONARE > SECȚIUNEA: IEȘIRI DIN STOC
         > Identificare Taxa Verde se va selecta taxa corespunzătoare, în cazul în care
         articolele pe care le facturați conțin această taxă.
        În cazul în care factura conține articole cu taxă va trebui să dați constantei: „Taxe
         facturabile: auto taxare pentru facturi de ieșire” valoarea DA.
        Codul fiscal al clientului respectă aceeași regulă ca mai sus (fără spațiu).
        Clientul are specificat contul bancar implicit în nomenclator.
        Denumirea articolului exportat este de maximum 50 caractere.
        Este obligatoriu ca articolele să aibă completat codul extern (codul EAN) și codul de
         catalog la client (în rubrica „Informații de la client”).
        Cantitatea articolelor pe factură se va specifica cu maximum 3 zecimale.
        Toate mărimile referitoare la preț, taxa verde, valoare, TVA, discount
         corespunzătoare articolelor se vor înscrie pe factura cu maximum 4 zecimale.
        La facturare se va folosi discount-ul pe coloană.
        Nu se va folosi scadența multiplă pentru facturare.
OBSERVAȚII:
               Este implementat și formatul numit format extins, care în aplicație este
                identificat prin denumirea DOCXCHANGE 2 (inclusiv exportul facturilor cu
                articole taxabile și/ sau discount pe coloană).
                      – Spre deosebire de formatul DOCXCHANGE, taxa verde se va
                        declara în: WMEdi > Identificare >> Identificare taxe facturate
                        clienților, în caz contrar, vor fi tratate ca articole obișnuite.
                        Beneficiari CARREFOUR, AUCHAN...

3. EXPERT SOLUTIONS
(beneficiari: REAL)
Parametri configurare WMEdi
        Se va specifica GLN-ul sediilor de livrare ale clientului și a sediului social dacă nu au
         fost deja încărcate.
        Identificarea firmei furnizorului: Tip Id.Firma - va fi specific, iar valoarea
         identificatorului va fi cea primită de la client.
        Codul articolului utilizat pentru identificare este codul de catalog al clientului.

Mod de lucru/Date obligatorii
       Obligatoriu, factura este generată pe baza unei comenzi a clientului.
       Factura nu poate conține articole ce provin de pe mai multe comenzi.
       Să fie completat contul bancar implicit al firmei furnizoare în CONSTANTE
        UTILIZATOR > CONSTANTE DE OPERARE > SECȚIUNEA: TRANZACȚII
        TREZORERIE > Bancă implicită. Fiind o constantă utilizator trebuie să fie
        completată pentru toți operatorii care facturează.
       În CONSTANTE GENERALE > FUNCȚIONARE > SECȚIUNEA: IEȘIRI DIN STOC
        > Identificare Taxa Verde se va selecta taxa corespunzătoare, în cazul în care
        articolele pe care le facturați conțin această taxă.
       Codul GLN este completat pe baza datelor primite de la client, atât pentru sediul
        social, cât și pentru sediile de livrare.
       Clientul are specificat contul bancar implicit în nomenclator.
       Cantitatea articolelor pe factură se va specifica cu maximum 3 zecimale.
       Toate mărimile referitoare la preț, valoare, TVA, discount corespunzătoare
        articolelor se vor înscrie pe factură cu maximum 4 zecimale.
       Taxa verde se va exprima cu maximum 2 zecimale.
       La facturare, se va folosi discount-ul pe coloană.
       Este obligatoriu ca articolele să aibă completat codul extern și codul de catalog la
        client.
       Unitățile de măsură utilizate pentru exportul articolelor vor fi cele convenite cu
        clientul prin intermediul ofertelor prealabile.

                                                  Fig.6
4. EDINET
(beneficiari: AUCHAN, DEDEMAN, CARREFOUR, CORA, OBI, LIDL, PRAKTIKER,
METRO, PENNY MARKET, ... (http://www.edinet.ro/lanturi-magazine-solutie-edi/))
Parametri configurare WMEdi
       Se va specifica GLN-ul sediilor de livrare ale clientului și a sediului social dacă nu au
        fost deja încărcate.
       Identificarea firmei furnizorului (Tip Id.Firma) se va face prin GLN-ul firmei.
       Codul articolului utilizat pentru identificare este codul de catalog al clientului.

Mod de lucru/Date obligatorii
        Factura este generată obligatoriu pe baza unei comenzi a clientului.
        Pe comanda dată livrării articolelor este completată obligatoriu.
        Factura nu poate conține articole ce provin de pe mai multe comenzi sau cu mai
         multe termene de livrare.
        Să fie completat contul bancar implicit al firmei furnizoare în CONSTANTE
         UTILIZATOR > CONSTANTE DE OPERARE > SECȚIUNEA: TRANZACȚII
         TREZORERIE > Banca implicită. Fiind o constantă utilizator trebuie să fie
         completată pentru toți operatorii care facturează.
        Se va completa GLN-ul furnizorului în FIRME.
        Codul GLN pentru client este completat pe baza datelor primite de la el, atât pentru
         sediul social, cât și pentru sediile de livrare.
        Cantitatea articolelor pe factură se va specifica cu maximum 3 zecimale.
        Toate mărimile referitoare la preț, valoare, TVA, discount corespunzătoare
         articolelor se vor înscrie pe factură cu maximum 2 zecimale.
        La facturare se va folosi discount-ul pe coloană.
        Este obligatoriu ca articolele să aibă completat codul de catalog la client și codul
         EAN preluat din codul extern al articolelor.
OBSERVAȚII:
               Pentru acest format este implementată și factura de corecție pentru care
                există documentație separată de utilizare, precum și exportul avizului de
                expediție.

5. NOVENSYS
(beneficiari: HORNBACH...)
Parametri configurare WMEdi
        Se va specifica GLN-ul sediilor de livrare ale clientului și a sediului social dacă nu au
         fost deja încărcate.
        Identificarea firmei vânzătorului (Tip Id.Firma) se va face prin identificatorul stabilit în
         prelabil cu clientul și stabilit în fereastra de identificare a clienților (identificator
         specific sau GLN).
        Codul articolului utilizat pentru identificare este codul EAN (codul extern al
         articolului). Opțional, se transmite și codul articolului la client dacă este completat în
         „Informații de la client” din nomenclatorul de articole.
Mod de lucru/Date obligatorii
        Factura este generată obligatoriu pe baza unei comenzi a clientului și a unui aviz de
         expediție pe baza căruia se face livrarea.
        Să fie completat contul bancar implicit al firmei furnizoare în CONSTANTE
         UTILIZATOR > CONSTANTE DE OPERARE > SECȚIUNEA: TRANZACȚII

        TREZORERIE > Bancă implicită. Fiind o constantă utilizator trebuie să fie
        completată pentru toți operatorii care facturează.
       Cantitatea articolelor pe factură se va specifica cu maximum 3 zecimale.
       Toate mărimile referitoare la preț, valoare, TVA, discount corespunzătoare
        articolelor se vor înscrie pe factură cu maximum 2 zecimale.

6. TECNET
(beneficiari: AUCHAN, REAL HYPER MAGAZINE SRL...)
Parametri configurare WMEdi
       Se vor specifica GLN-urile sediilor de livrare ale clientului și a sediului social dacă nu
        au fost deja încărcate în nomenclatorul de parteneri.
       Identificarea firmei vânzătorului (Tip Id.Firma) se va face prin identificatorul primit în
        prealabil de la client și stabilit în fereastra de identificare a clienților (identificator
        specific).

ATENȚIE!
        Soluția de implementare pentru AUCHAN, BILLA și PENNY presupune, în cazul în
        care articolele livrate clienților aparțin mai multor grupe de articole (clase de articole)
        din clasificarea clienților, să fie facturate separat pe fiecare clasa în parte și să aibă
        cod furnizor corespunzător grupei de articole.
        În mod special, implementarea pentru BILLA presupune asocierea unui cod GLN
        unic al vânzătorului pentru fiecare cod furnizor asociat grupelor de articole.
        În aceste condiții s-au introdus două metode noi de identificare a firmei față de client
        (vezi figura 7):
                             Id specific gr.art. - Id specific (cod furnizor) asociat grupelor de
                       articole client. Valoarea Id-ului este furnizată de către client.
                            GLN+ Id specific gr.art. - GLN asociat Id-ului specific furnizat
                       de client. În acest caz GLN-ul trebuie obținut de firmă (de la GS1
                       România) și comunicat clientului.
        Soluția s-a completat prin introducerea unei noi opțiuni de identificare în WMEdi:
        „Configurare grupare articole clienți”, care presupune asocierea Id-ului (cod
        furnizor), eventual a GLN-ului, articolelor vândute clienților, identificați prin una din
        cele două metode (vezi figura 8).
       Codul articolului utilizat pentru identificare este codul articolului la client. Obligatoriu,
        se transmite și codul extern al articolului (EAN) completat în nomenclatorul de
        articole în rubrica „Cod extern (EAN)”.
       Pentru identificarea taxelor facturabile asociate articolelor, se vor specifica în
        prealabil articolele de tip serviciu, utilizate la facturarea lor, în fereastra opțiunii
        Identificare >> Identificare taxe facturate clienților, în caz contrar vor fi tratate ca
        articole obișnuite. Se pot defini două tipuri de taxe: Eco Tax și Green Tax.

                             Fig.7

                             Fig.8

Mod de lucru/Date obligatorii
        Factura este generată obligatoriu pe baza unei comenzi a clientului și a unui aviz de
         expediție pe baza căruia se face livrarea.
        Cantitatea articolelor pe factură se va specifica cu maximum 2 zecimale.
        Toate mărimile referitoare la preț, valoare, discount corespunzătoare articolelor se
         vor înscrie pe factură cu maximum 4 zecimale.
        Exportul permite înregistrarea facturilor obișnuite, de retur și de corecție. Facturile
         de retur și corecție nu pot face referire la mai multe facturi obișnuite. Dacă nu s-a
         identificat automat tipul facturii, la momentul exportului, identificarea se poate face
         manual în interfața de generare. Datele solicitate sunt obligatorii.

Review-uri document
