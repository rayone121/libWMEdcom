# IMPORT PLATI DIN ALTE APLICATII

<sub>Source: `22_Structuri import din alte aplicatii__VENITURI PROFESIONALE__Import plati venituri profesionale - utilizare1.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
IMPORT PLATI DIN ALTE APLICATII
           IN PLATA VENITURILOR PROFESIONALE
DESCRIERE:

        Permite importul platilor din alte aplicatii, simplificand procedura de inregistrare a platilor, mai
ales in cazul contractelor de arenda inregistrate in alte aplicatii.

FUNCTIONARE:

        Se selecteaza optiunea "Import plati din alte aplicatii" din meniul afisat la selectarea butonului "+"
albastru (vezi figura 1). Se solicita selectarea fisierului .xls ce contine platile generate din alte aplicatii
pentru luna si anul de lucru curent (vezi structura mai jos).
        In interfata afisata (figura 2) se va specifica linia si coloana din care incep datele de importat
ale fisierului Excel. Ordinea coloanelor trebuie sa coincida cu structura fisierului de import specificat in
randurile de mai jos.
        Inregistrarile importate se adauga la cele existente numai in cazul in care nu se dubleaza si
structura inregistrarilor este completa (toate campurile sunt completate). Pot exista mai multe plati pentru
o persoana.

                                                   Fig. 1

Creare: 27.01.2015; Modificare:12/13/2017

                                                 Fig. 2

STRUCTURA FISIER IMPORT PLATI (format .xls):
 • Data plata      Se valideaza ca anul si luna din data sa coincida cu anul si luna de lucru
                   (zz.ll.aaaa)
 • CNP             CNP titular contract
 • Tip contract VP Poate lua valorile: 6, 7, 8, 9, 10, 11, 26
 • Valoare plata
 • Numar contract

Observatii:
Tip Contract VP - decodificarea codurilor utilizate:
   • 6 - Drepturi de proprietate intelectuala
   • 7 - Contracte / conventii civile
   • 8 - Expertiza contabila, tehnica, etc.
   • 10 - Asociere potrivit titlului III din CF
   • 11 - Asociere potrivit titlului II din CF
   • 26 - Venituri din arenda bunurilor agricole

Importul completeaza grila 1 din foaia de calcul si lanseaza executia procedurii "Calcul stopaj".

Creare: 27.01.2015; Modificare:12/13/2017
