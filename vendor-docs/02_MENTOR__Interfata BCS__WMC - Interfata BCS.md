# Interfața BCS

<sub>Source: `02_MENTOR__Interfata BCS__WMC - Interfata BCS.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Interfața BCS
Interfața BCS este dedicată importului articolelor, din fișiere în format .txt, în vederea
simplificării procedurii de introducere a datelor.
Fișierele utilizate pentru import pot fi generate de către alte aplicații, de MENTOR sau de
către utilizatori prin prelucrarea documentelor primite de la diverși parteneri, în formate
diverse (.doc, .xls, .pdf etc.).
În constante generale există o categorie specială de constante: „Interfața BCS” prin care se
configurează modul de lucru.

Metode de identificare a articolelor importate
Cu ajutorul constantei „Articolele se identifică prin” se configurează modul de identificare a
articolelor în baza de date. Există următoarele posibilități:
              Denumire;
              Cod extern;
              Cod intern;
              Cod catalog partener.
Ultima metodă de identificare folosește Cod catalog furnizor pentru comenzile către
furnizori și tranzacțiile de intrare în stoc, și Cod catalog clienți pentru comenzile de la clienți
și tranzacțiile de ieșire din stoc.

Structuri de import/export
    1. importurile de tranzacții de intrare/ieșire din stoc – fișierul trebuie să aibă
       forma:
Identificator,0,22.3,1209.54[,articol,um]
unde :
               Identificator – identificator articol ce poate avea una din următoarele valori:
                denumire articol, cod extern, cod intern, cod catalog partener;
               0 – unitatea de măsură principală sau secundară din num1.db (0 – principală,
                1 – secundară);
               22.3 – cantitatea;
               1209.54 – preț.
În cazul adăugării în nomenclatorul de articole a unor articole inexistente în baza de date,
prin intermediul importului articolelor tranzacției, formatul va cuprinde obligatoriu și
câmpurile:
     articol – denumirea articolului inexistent în baza de date ce urmează a fi introdus în
      nomenclatorul de articole;
     um – denumirea unității de măsură corespunzătoare articolului (există în baza de
      date).
În cazul exportării articolelor tranzacțiilor, câmpurile se completează automat.

    2. importul comenzilor către furnizori/de la clienți – structura este asemănătoare
       cu cea de mai sus, diferența constând în tipul separatorului:
Identificator;0;22.3;1209.54[,articol,um]
unde :
            Identificator – identificator articol ce poate avea una din următoarele valori:
             denumire articol, cod extern, cod intern, cod catalog partener.
    3. importul inventarului:
Identificator, Cant,[serie]
unde :
            Identificator – identificator articol ce poate avea una din următoarele valori:
             denumire articol, cod extern, cod intern.
            Cant – cantitate faptic exprimată în unitatea de măsură principală (cea din
             nomenclatorul de articole). Pentru articolele cu serie pe bucată cantitatea va fi
             „1”.
            Serie – element opțional, este folosit doar pentru articolele cu serie (pe
             bucată sau lot). Opțiunea este utilizabilă începând cu versiunea 841.16.
    4. importul stocurilor inițiale:
Identificator, 22.3, 1209.54, 1300
unde :
            Identificator – denumirea articolului, cod extern, cod intern;
            22,3 – cantitatea exprimată în unitatea de măsură principală;
            1029.54 – prețul de achiziție – pentru marfă cu adaos la preț de înregistrare
             (poate lipsi) ;
            1300 – prețul de înregistrare.
Structura minimală rămâne : Identificator, Cantitate.
În cazul inițializării, articolele noi pot actualiza nomenclatorul de articole. Structura în acest
caz trebuie să aibă următoarea formă:
Identificator, cantitate, pretAchizitie, pretInregistrare, codextern, codintern, denumireArt,
denumireUM, simbolClasa, simbol gestiune implicită, pretVanzare.
Unde:
     Identificator – identificator articol (denumire, cod intern sau cod extern articol).
     Câmpurile: codextern, codintern, denumireArt, denumireUM, simbolClasa, simbol
      gestiune implicită, preț vânzare – reprezintă secțiunea de informații pentru
      adăugarea articolului nou.
Din toate aceste câmpuri din secțiunea de informații pentru adăugare articol nou, care
caracterizează articolul, obligatorii sunt cele subliniate.
Ex: Trebuie să adăugăm un articol cu:
            Denumirea: Zahăr,

               UM: Kg,
               Cod Extern: 123456,
               Cod Intern: 9ABC200,
               Gestiunea implicită: DC (Depozit central).
               Constanta de identificare este prin:
                Cod intern.
9ABC200,100,5.00,5.50,123456,9ABC200,Zahar,Kg,DC,34
unde:
               100 – cantitatea exprimată în unitatea de măsură principală,
               5.00 – prețul de achiziție,
               5.50 – prețul de înregistrare,
               DC – simbolul gestiunii implicite ,
               34 - prețul de vânzare.
Separatorul zecimal pentru cantitate și prețuri este „.”.
