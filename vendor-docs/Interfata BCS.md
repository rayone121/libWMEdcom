# Interfaţa BCS

<sub>Source: `22_Structuri import din alte aplicatii__Interfata BCS.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Interfaţa BCS
Interfaţa BCS este dedicată importului articolelor, din fişiere în format .txt, în vederea
simplificării procedurii de introducere a datelor.

Fişierele utilizate pentru import pot fi generate de către alte aplicaţii, de MENTOR sau
de către utilizatori prin prelucrarea documentelor primite de la diverşi parteneri, în
formate diverse (.doc, .xls, .pdf etc.).

În constante generale, există o categorie specială de constante: „Interfaţa BCS” prin
care se configurează modul de lucru.

Metode de identificare a articolelor importate

Cu ajutorul constantei „Articolele se identifică prin” se configurează modul de
identificare a articolelor în baza de date. Există următoarele posibilităţi:
    Denumire;
    Cod extern;
    Cod intern;
    Cod catalog partener.

Ultima metodă de identificare foloseşte Cod catalog furnizor pentru comenzile către
furnizori şi tranzacţiile de intrare în stoc, şi Cod catalog clienţi pentru comenzile de la
clienţi şi tranzacţiile de ieşire din stoc.

Structuri de import/export

1. La importurile de tranzacţii de intrare/ieşire din stoc fişierul trebuie să aibă forma:

Identificator,0,22.3,1209.54[,articol,um]
unde :
    Identificator – identificator articol ce poate avea una din următoarele valori:
        denumire articol, cod extern, cod intern, cod catalog partener;
    0 – unitatea de măsură principală sau secundară din num1.db (0 – principală, 1–
        secundară);
    22.3 – cantitate;
    1209.54 – preţ.

În cazul adăugării în nomenclatorul de articole a unor articole inexistente în baza de
date, prin intermediul importului articolelor tranzacţiei, formatul va cuprinde obligatoriu şi
câmpurile:
     articol – denumirea articolului inexistent în baza de date ce urmează a fi introdus
       în nomenclatorul de articole;
     um – denumirea unităţii de măsură corespunzătoare articolului (există în baza de
       date).
În cazul exportării articolelor tranzacţiilor, câmpurile se completează automat.

2. La importul comenzilor către furnizori/de la clienţi structura este asemănătoare
cu cea de mai sus, diferenţa constând în tipul separatorului:

Identificator;0;22.3;1209.54[,articol,um]

unde :
Identificator – identificator articol ce poate avea una din următoarele valori: denumire
articol, cod extern, cod intern, cod catalog partener.

3. importul inventarului:

Identificator, Cant,[serie]

unde :
    Identificator – identificator articol ce poate avea una din următoarele valori:
       denumire articol, cod extern, cod intern.
    Cant – cantitate faptic exprimată în unitatea de măsură principală (cea din
       nomenclatorul de articole). Pentru articolele cu serie pe bucata cantitatea va fi
       „1”.
    Serie – element opţional, este folosit doar pentru articolele cu serie (pe bucată
       sau lot). Opţiunea este utilizabilă începând cu versiunea 841.16.

4. importul stocurilor iniţiale:

Identificator, 22.3, 1209.54, 1300

unde :
    Identificator – denumirea articolului, cod extern, cod intern;
    22,3 – cantitatea exprimată în unitatea de măsură principală;
    1029.54 – preţul de achiziţie – pentru marfă cu adaos la preţ de înregistrare
       (poate lipsi);
    1300 – preţul de înregistrare;

Structura minimală rămâne : Identificator, Cantitate.

În cazul iniţializării, articolele noi pot actualiza nomenclatorul de articole. Structura în
acest caz trebuie să aibă următoarea formă:

Identificator, cantitate, pretAchizitie, pretInregistrare, codextern, codintern, denumireArt,
denumireUM, simbolClasa, simbol gestiune implicită, pretVanzare, DataIN, TG.

Unde:
Identificator – identificator articol (denumire, cod intern sau cod extern articol)
DataIN – reprezinta data intrare in gestiune
TG – termenul de garantie

Cu click dreapta pe View Initializare stocuri se va putea afisa optiunea Import stocuri de
la interfata BCS.
Câmpurile: codextern, codintern, denumireArt, denumireUM, simbolClasa, simbol
gestiune implicită, preţ vânzare – reprezintă secţiunea de informaţii pentru adăugarea
articolului nou.

Din toate aceste câmpuri din secţiunea de informaţii pentru adăugare articol nou, care
caracterizează articolul, obligatorii sunt cele subliniate.

Ex: Trebuie să adăugăm un articol cu Denumirea: Zahăr,
UM: Kg,
Cod Extern: 123456,
Cod Intern: 9ABC200,
Gestiunea implicită: DC (Depozit central). Constanta de identificare este prin : Cod
intern.

9ABC200,100,5.00,5.50,123456,9ABC200,Zahăr,Kg,SimbolClasa,DC,34

unde:
    100 – cantitatea exprimată în unitatea de măsură principală;
    5.00 – preţul de achiziţie;
    5.50 – preţul de înregistrare;
    DC – simbolul gestiunii implicite;
    34 – preţul de vânzare;
    Separatorul zecimal pentru cantitate şi preţuri este „.”.
