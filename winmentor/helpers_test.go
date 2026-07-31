package winmentor

import "testing"

// A handful of WinMENTOR CodExtern values contain the record separator itself.
// Before splitFieldsMerging these shifted every later field left by one, so the
// unit of measure was stored as the sale price. Rows are verbatim from
// wme-raw/GetNomenclatorArticole.txt and wme-raw/GetStocuriPeGestiuni.txt.
func TestSplitFieldsMerging(t *testing.T) {
	nomenclator := "35mmp  ; H01N2-D;Cablu sudura;M;;4 Piese;Marfuri = piese;;;Piese;377;21;;;;Nu;;;;;26.05.2009;;0;;;DA;;;0;;NU;;;;;;;NU;;;"
	f := splitFieldsMerging(nomenclator, 40, 0)
	if len(f) != 40 {
		t.Fatalf("got %d fields, want 40", len(f))
	}
	if f[0] != "35mmp  ; H01N2-D" {
		t.Errorf("CodExtern = %q, want %q", f[0], "35mmp  ; H01N2-D")
	}
	if f[1] != "Cablu sudura" {
		t.Errorf("Denumire = %q, want %q", f[1], "Cablu sudura")
	}
	if f[2] != "M" {
		t.Errorf("DenUM = %q, want %q", f[2], "M")
	}
	if f[3] != "" {
		t.Errorf("PretVanzare = %q, want empty (was the unit before the fix)", f[3])
	}

	stoc := "Marfuri = piese;Piese;Cablu sudura;35mmp  ; H01N2-D;371.01;M;50;25,78;25,78;21"
	g := splitFieldsMerging(stoc, 10, 3)
	if len(g) != 10 {
		t.Fatalf("got %d fields, want 10", len(g))
	}
	if g[3] != "35mmp  ; H01N2-D" {
		t.Errorf("CodExtern = %q, want %q", g[3], "35mmp  ; H01N2-D")
	}
	if g[6] != "50" {
		t.Errorf("Stoc = %q, want %q", g[6], "50")
	}
	if g[8] != "25,78" || g[9] != "21" {
		t.Errorf("ValoareStocPrecisa/CotaTVA = %q/%q, want 25,78/21", g[8], g[9])
	}

	// Records without an embedded separator must be untouched.
	clean := splitFieldsMerging("A;B;C", 3, 0)
	if clean[0] != "A" || clean[1] != "B" || clean[2] != "C" {
		t.Errorf("clean record altered: %q", clean)
	}
	// Short records still pad.
	if p := splitFieldsMerging("A;B", 4, 0); len(p) != 4 || p[0] != "A" || p[3] != "" {
		t.Errorf("short record = %q", p)
	}
}
