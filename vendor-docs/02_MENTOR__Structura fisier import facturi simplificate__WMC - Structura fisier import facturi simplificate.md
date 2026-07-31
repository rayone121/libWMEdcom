# Structură fișier pentru importul facturilor

<sub>Source: `02_MENTOR__Structura fisier import facturi simplificate__WMC - Structura fisier import facturi simplificate.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `Exemplu:` |  |
| `anul 2013, luna martie.` |  |
| `înregistrări.` |  |
| `Structura fișierului:` |  |
| ` cod fiscal client` | A15 se preiau indiferent dacă au |
| ` denumire client` | sau nu RO înaintea codului fiscal urmând ca acest lucru să se corecteze în declarație A100 dacă nu se completează, |
| ` bază calcul` | programul va atribui automat denumirea „Bon Casa de Marcat” TVA N15 bază impozabilă |
| ` valoare` | TVA N15 TVA |

Structură fișier pentru importul facturilor
simplificate din alte aplicații, în declarația D394
Datele vor fi furnizate în fișiere format .txt.
Denumirea fișierului va avea următoarea structură:
394_cod fiscal firma_an raportare_luna raportare_numar fisier.
Firma care face raportarea are codul fiscal RO123456789, iar datele se referă la
Denumirea fișierului va fi: 394_123456789_2013_03_01.txt
Câmpurile fișierului se pot separa prin caracterul „;”, tab etc.
Dacă fișierul este obținut din excel, verificați să nu aveți și antetul grilei printre
