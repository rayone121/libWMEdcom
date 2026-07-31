# Mod funcționare WMGenCmd

<sub>Source: `23_WMGenCmd__ModFunctionare_WMGenCmd.pdf` — WinMENTOR official documentation, converted from PDF.</sub>
Mod funcționare WMGenCmd

    În fișierul “WMGenCmd.ini” se setează calea către directorul în care aplicația web va scrie
    fișierele de documente ce se doresc importate în WinMENTOR, sau interogări ale stocurilor
    și articolelor din WinMENTOR (directorul prin care cele două aplicații comunică).

    “FirmaLucru” reprezintă numele prescurtat al firmei pe care se lucrează, așa cum apare ea
    în baza de date WinMENTOR.

    “IDArticol” reprezintă coloana după care se identifică unic articolele din baza de date
    WinMENTOR. Coloana implicită este CODINTERN.

    “IDPartener” reprezintă coloana după care se identifică unic partenerii din baza de date
    WinMENTOR. Coloana implicită este CODINTERN. Poate lua valorile CODEXTERN,
    CODFISCAL, CODINTERN.

    “Gestiune” se poate introduce în momentul când se dorește interogarea stocurilor doar
    dintr-un număr limitat de gestiuni. Se vor introduce simbolurile acestor gestiuni din baza de
    date WinMENTOR separate cu ';'.

    “User” reprezintă un nume valid de utilizator la baza de date Oracle, care are drepturi de
    logare la firma de lucru.

    “Parola“ reprezintă parola pentru autentificare.

   Pentru INTEROGAREA STOCURILOR aplicația web scrie în directorul de lucru fișierul
    “GetStoc.txt”. Răspunsul este primit în fișierul “'Stocuri.txt' pe structura:
        IDArticol;Denumire Articol;DenUM;Stoc;Stoc rezervat;Pret vanzare;Simbol clasa;Den
        Clasa
     unde stoc rezervat reprezintă stocul rezervat de alte comenzi.

   Pentru INTEROGAREA INFORMAȚIILOR privind articolele din baza de date
    WinMENTOR aplicația WEB scrie în directorul de lucru fișierul ”Getinfoart.txt”.

        Răspunsul este primit în fișierul “Infoart.txt” pe structura:
             Cod Extern;Cod Intern;Denumire Articol;Producator;Unitatea de          Masura;Pret
             Vanzare;Pret Valuta ;TVA;Data Adaugarii;SombolClasa;DenClasa;

   Pentru ÎNREGISTRAREA COMENZILOR DE LA CLIENȚI aplicația web va scrie în
    directorul de lucru fișiere de comenzi pe structura care se poate vedea ca exemplu în

    fișierul “Comanda_1.txt”.
    Denumirea acestor fișiere trebuie să respecte formatul Comanda_N.txt
    unde N este un număr unic gestionat de aplicația web.

 Pentru ÎNREGISTRAREA COMENZILOR DE LA FURNIZORI aplicația web va scrie în
  directorul de lucru fișiere de comenzi pe structura care se poate vedea ca exemplu în
  fisierul “Comanda_furn.txt”.

    WMGenCmd.exe procesează aceste date, și dacă formatul este valid comanda va fi
    importată în baza de date WinMENTOR (prin DocImpServer) după care va șterge din
    director fișierul respectiv.
    Dacă importul nu a reușit, fișierul va fi redenumit de către WMGenCmd.exe sub forma
    Comanda_N.err iar în fișierul ”Erori.txt” vor apare explicații privind erorile apărute.

   Pentru ÎNREGISTRAREA FACTURILOR aplicația web va scrie în directorul de lucru fișiere
    de facturi pe structura care se poate vedea ca exemplu în fișierul “Factura_1.txt”.

    Denumirea acestor fișiere trebuie să respecte formatul “Factura_N.txt” unde “N” este
    un număr unic gestionat de aplicația web.

    “WMGenCmd.exe” procesează aceste date și dacă formatul este valid, pachetul va fi
    importat în baza de date WinMENTOR (prin DocImpServer), după care va șterge din
    director fișierul respectiv.

    Dacă importul nu a reușit, fișierul va fi redenumit de către “WMGenCmd.exe” sub forma

    “Factura_N.err” iar în fișierul “Erori.txt” vor apare explicații privind erorile apărute.

   Pentru ÎNREGISTRAREA PARTENERILOR NOI, aplicația web va scrie în directorul de
    lucru fișiere cu denumire Partener_N.txt
    unde N este un număr unic gestionat de aplicația web.

    Structura unei linii cu informații privind noul partener este următoarea:
        ID Partener;
        Nume partener;
        CodFiscal;
        Localitate sediu social;

    2                                        Rev 1.0 – 24.05.2024                         Clasificare: Public
     Strada sediu social;
     Telefon Sediusocial;
     Pers contact;
     Clasa;
     Simbol categ pret;
     Marca agent implicit;
     RegCom;
     Observatii;
     Simbol Banca;
     Sucursala;
     Localitate Sucursala banca;
     Cont Bancar;
     scadenta implicita;
     nume sedii de livrare; (daca sunt mai multe se vor separa cu caracterul ~)
     adrese sedii; (daca sunt mai multe ordinea va fi cea de la nume sedii)
     telefoane sedii (eventual separate cu ~)
     localitati sedii ( eventual separate cu ~)
     marci agenti sedii; ( eventual separate cu ~)
     denumire sediu social;
     flag persoana fizica; (DA/NU);

    Tipul sediilor în ordinea de la nume sedii, separate cu '~' //1 = livrare; 2 = Facturare; 3 =
    Livrare + Facturare
    ID-uri sediilor în ordinea de la nume sedii, separate cu '~'
    Simbol gest listare sedii în ordinea de la nume sedii, separate cu '~'
    Partener blocat (DA/NU)

     codintern;numepartener;ro-ul        firmei;localitate;adresa;nr       de    telefon;persoana   de
     contact;;;1;j-ul firmei;;banca;;;contulbancar;
    Exemplu:
    2121;SC MAREX SRL;RO 544344;Iasi;str                   culturii   nr   71   bis;0232712844;Ionescu
    Adrian;;;1;J-21-22-94;;BRD;;;111;

   Pentru interogarea informațiilor privind prețurile multiple ale articolelor din
    WinMENTOR, aplicația WEB scrie în directorul de lucru fișierul “GetPretArt.txt”.

    Răspunsul este primit în fișierul “ListaArtCatPret.txt” pe structura:
        IDARticol;SimboCategPret;Pret fara TVA;Implicit(Da).
    Pentru un articol care are mai multe prețuri, vor apare în acest fișier text mai multe linii.

   Pentru interogarea NOMENCLATORULUI DE GESTIUNI de gestiuni aplicația WEB scrie în
    directorul de lucru fișierul ”GetListaGestiuni.txt”.

    Răspunsul este primit în fișierul “ListaGestiuni.txt” pe structura :
        SImbolGestiune;Denumire;Simbol categ.pret;

   Pentru interogarea INCASARILOR DE LA CLIENTI, aplicația web scrie în directorul de
    lucru fișierul “GetIncasariClienti.txt”.

    Pe prima linie a fișierului trebuie puși parametrii:

        An Inceput analiza, Luna inceput analiza, An sfarsit analiza, Luna sfarsit analiza,
        IDPartener (separati cu virgula)

    Exemplu:
    2018,1,2018,7,RO24903480

    Dacă IDPartener este șir vid, interogarea returnează toate încasările din intervalul
    selectat.

    Răspunsul este salvat în fișierul “IncasariClienti.txt” pe structura:

    IDPartener;Document incasare;Data incasarii;Suma Incasata;Documente incasate

    Secțiunea "Documente Incasate" include detalierea facturilor care fac obiectul încasării, pe
    structura :

         Serie si numar Factura/DataFactura=Suma incasata concatenate cu separatorul ~

    În cazul avansurilor, informația este sub forma “Avans=Suma incasata”

    Exemplu de linie returnată:

         RO24903480;Ch 122;01.03.2018;100;~F.MOLD 8/01.03.2018=17,85~F.MOLD

   Pentru ÎNREGISTRAREA FACTURILOR DE LA FURNIZORI aplicația web va scrie în
    directorul de lucru fișiere de FACTURI ale căror denumiri trebuie să respecte formatul
    FacturaFurn_N.txt;

    4                                        Rev 1.0 – 24.05.2024                      Clasificare: Public
    WMGenCmd.exe procesează aceste date, și dacă formatul este valid factura va fi importată
    în baza de date WinMENTOR (prin DocImpServer), după care va șterge din director fișierul
    respectiv.
    Dacă importul nu a reușit, fișierul va fi redenumit de către WMGenCmd.exe sub forma
    “FacturaFurn_N.err” iar în fișierul “Erori.txt” vor apare explicații privind erorile date.

   Pentru setartea unei liste de comenzi facturabile web va scrie în directorul de lucru
    fisiere text care respectă denumirea SetCmdFacturabile_N.txt, unde N reprezintă un ID
    unic.

    Structura unei linii din acest fișier trebuie să respecte structura :
        NumarComanda

   Pentru înregistrarea ARTICOLELOR NOI, aplicația web va scrie în directorul de lucru
    fișiere cu denumire Articol_N.txt

    unde N este un număr unic gestionat de aplicația web.
    Structura unei linii cu informații privind noul articol este următoarea:
        IDArticol;
        Denumire;
        Denumire UM;
        Simbol cont (pentru articole de tip serviciu);
        Denumire UM ambalare;
        Cod extern UM ambalare;
        Cod intern UM ambalare;
        Masa UM ambalare;
        Greutate specifica UM ambalare;

   Pentru funcția de modificare date parteneri prin care să se actualizeze coloana “Blocat”
    trebuie să existe pe disc un fișier update Partener_N.txt în care să se precizeze IDPART; D
    sau IDPART; N
        D = blocare
        N = deblocare
   Materiale rețetă: se copie fișierul “GetMaterialeReteta.txt” care pe prima linie conține
    IDProdus.
    Se returnează datele în “MaterialeReteta.txt”.

   Pentru consultarea facturilor emise în baza unor comenzi Online, aplicația client scrie

    în folderul de lucru fișierul : “GetFacturiCmdOnline.txt”

    Răspunsul este dat în fișierul “FacturiCmdOnline.txt” pe structura:

        IDClient;DataDocument;NumarDocument;IDarticol;Numar
        comanda;Cantitate;UM;Pret;MarcaAgent;Valoare factura;TVA aferent liniei de pe
        factura;Flag factura anulata (DA/NU);Rest de plata factura;Termen plata factura;Numar
        inregistrare factura;

    Observații:

    1. pentru a seta faptul ca o comandă este postată din aplicatii Online, în pachetul comenzii
       trebuie adaugată sub sectiunea [COMANDA_...] linia :

        DinOnline=D

    2. Dacă se dorește ca interogarea “FacturiCmdOnline.txt” să aducă inclusiv facturi care
       sting comenzi care nu provin din aplicații Online, în fișierul “WMGenCmd.ini” se adaugă
       linia :

        InclusivFacturiOffline=D

   Pentru INTEROGAREA NOMENCLATORULUI DE PARTENERI aplicația web scrie în
    directorul de lucru fișierul GETLISTAPARTENERI.TXT

    Răspunsul este primit în fisierul Parteneri.txt. Structura o găsiți detaliată în fișierul
    Structura_returnata_GetListaParteneri.txt.

   Pentru INREGISTRAREA FACTURILOR CÂTRE FURNIZORI aplicația web va scrie în
    directorul de lucru fișiere cu denumirea ”FacturaFurn_x.txt” unde x este un ID unic al
    pachetului.

   Interogare GetCriteriiDiscPeArticole.Txt
    Returnează           în        fișierul          “CriteriiDiscPeArticole.txt”     informațiile:
    Den.Criteriu;IDArticol;ProcentDiscount;

   Interogare GetCriteriiDiscPeClase.txt
    Returnează          în        fișierul     “CriteriiDiscPeClase.Txt”              informațiile:
    Den.Criteriu;Den.Clasa;ProcentDiscount;SimbolClasa;

   Interogare GetCriteriiDiscPart.txt
    Returnează în “CriteriiDiscPart.txt” informațiile:
    IDPartener;Den.Criteriu ;

    6                                       Rev 1.0 – 24.05.2024                    Clasificare: Public
