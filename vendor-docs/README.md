# Vendor documentation

WinMENTOR's own documentation, converted to markdown. Everything here is
published openly at <http://download.winmentor.ro/WinMentor/Documentatie/> —
no login, no licence gate. The tree has 40 folders and 682 files; this is the
subset that bears on integration.

Regenerate with `python3 pdf2md.py <pdf-dir> <md-dir>`.

## What is worth reading

| File | Why |
|---|---|
| `24_DocImpServer__Functii pentru interfatarea cu WinMENTOR.md` | the COM function reference, **Rev.1.5 of 22 May 2026** |
| `24_DocImpServer__Functions for interfacing…md` | the same, in English |
| `23_WMGenCmd__Structura_returnata_GetListaParteneri.txt` | **48 numbered fields** for `GetListaParteneri` — far better than the 38 in the main manual |
| `23_WMGenCmd__*.txt` | real packet examples: partner, customer order, supplier order, invoice |
| `Bonuri de consum.md`, `Avize expeditie iesire.md`, `Transfer.md`, … | one import structure per document type |

## Two things this settled

**The aviz packet.** `Avize expeditie iesire.md` gives `TotalAvize` / `[Aviz_n]`
/ `Scadenta=` as a plain key. We were sending `TotalFacturi` / `[Factura_n]` /
a `[Scadente_1]` section, which is why WinMENTOR answered "nu pare a fi un
fisier care sa contina FACTURI IESIRE" — invoice sections under a delivery-note
type. Avize can be imported.

**The bon de consum packet is correct.** `Bonuri de consum.md` confirms
`TipDocument=BON DE CONSUM`, `TotalBonuri`, `[Bon_n]`, `GestConsum` and the
8-field item line, all matching what `internal/sync/bonuri_consum.go` emits.

## Caveats

The manual **understates field counts** on nearly every reader — 24 documented
for `GetNomenclatorArticole` against 40 in an April 2026 dump and 42 live. Take
names from here, counts from the data. See `../LAYOUTS.md`.

Rev.1.5 added `GetInfoIesiri` and `GetInfoComenzi` layouts and corrected
`GetOferte` to 14 fields; Rev.1.4 documented none of them.

Conversion is mechanical and imperfect: the source PDFs are two-column tables
whose right column wraps unpredictably, so long explanations can arrive split
across rows. Parameter names and packet structures — the load-bearing part —
come through clean. When precision matters, read the PDF.
