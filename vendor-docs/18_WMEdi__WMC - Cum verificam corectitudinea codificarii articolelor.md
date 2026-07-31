# Cum verificăm corectitudinea codificării articolelor

<sub>Source: `18_WMEdi__WMC - Cum verificam corectitudinea codificarii articolelor.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Cum verificăm corectitudinea codificării articolelor
                  tranzacționate prin intermediul importului de
                 comenzi/exportului de facturi în format .XML?

Verificarea corectitudinii codificării articolelor, tranzacționate prin intermediul fișierelor în
format .xml, se poate face prin configurarea unei noi versiuni de listă generată pe baza
datelor utilizate pentru generarea situației „Informații către client” din WinMENTOR,
modulul COMERCIAL(vezi figura1).

                                               Fig.1

                                               Fig.2
Lista permite selectarea unui anumit partener (vezi figura 2).
Structura noii liste se poate observa în figura 3, inclusiv modul de grupare.
Pe site, în directorul versiunii curente WinMENTOR, subdirectorul „ _CONFIGURARI NOI
LISTE WinMENTOR”, este configurarea acestei liste (CODCLI.ml, CODCLI.ml1),

configurare ce poate fi importatăprin intermediul opțiunii din COMERCIAL > LISTE >
UNELTE > „Manager liste”: Clienți > Informații către client.

                                             Fig.3
ATENȚIE!
Câmpul „Cod extern” din nomenclatorul de articole va conține întotdeauna codul EAN al
articolului, indiferent de tipul contului și indiferent de tipul importului sau exportului de
tranzacții.

Review-uri document
