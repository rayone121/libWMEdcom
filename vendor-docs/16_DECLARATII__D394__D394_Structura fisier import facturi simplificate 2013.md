# Structură fișier

<sub>Source: `16_DECLARATII__D394__D394_Structura fisier import facturi simplificate 2013.pdf` — WinMENTOR official documentation, converted from PDF.</sub>

| Parametru | Explicație |
|---|---|
| `raportare_luna raportare_numar fisier.` | Structură fișier în declarația D394 |
| `Exemplu:` |  |
| `anul 2013, luna martie.` |  |
| `înregistrări.` |  |
| `Structura fișierului:` |  |
| ` cod fiscal client` | A15 se preiau codurile, indiferent dacă au sau nu RO |
| ` denumire client` | înaintea codului fiscal, urmând ca acest lucru să se corecteze în declarație A100 dacă nu se completează programul, va atribui |
| ` bază calcul TVA` | automat denumirea „Bon Casa de Marcat” N15 bază impozabilă |
| ` valoare TVA` | N15 TVA |
| `Exemplu de conținut fișier:` |  |

pentru importul facturilor simplificate din alte aplicații,
Datele vor fi furnizate în fișiere format .txt.
Denumirea fișierului va avea următoarea structură: 394_cod fiscal firma_an
Firma care face raportarea are codul fiscal RO123456789, iar datele se referă la
Denumirea fișierului va fi: 394_123456789_2013_03_01.txt
Câmpurile fișierului se pot separa prin caracterul „;”, tab etc.
Dacă fișierul este obținut din excel, verificați să nu aveti și antetul grilei printre
