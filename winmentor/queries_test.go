package winmentor

import "testing"

// The manual lists "Nume;Prenume;Marca;..." but the DLL sends the full name in
// one field, so mapping a Prenume shifted everything after it one place left:
// Marca received the CNP, EsteAgent received the ID-card series, and CodPostal
// came back empty. Record is verbatim from wme-raw/GetListaPersonal.txt.
func TestEmployeeFieldsAreNotShiftedByAPhantomPrenume(t *testing.T) {
	rec := "GURGU ION;1;1661114080041;Da;Nu;BV;623323;2200;;"

	f := splitFields(rec, 10)
	if len(f) != 10 {
		t.Fatalf("got %d fields, want 10", len(f))
	}
	e := Employee{
		Nume: f[0], Marca: f[1], CNP: f[2], EsteActiv: f[3], EsteAgent: f[4],
		SerieBuletin: f[5], NumarBuletin: f[6], CodPostal: f[7],
		NumeUtilizator: f[8], Unknown9: f[9],
	}

	for _, c := range []struct{ name, got, want string }{
		{"Nume", e.Nume, "GURGU ION"},
		{"Marca", e.Marca, "1"},
		{"CNP", e.CNP, "1661114080041"},
		{"EsteActiv", e.EsteActiv, "Da"},
		{"EsteAgent", e.EsteAgent, "Nu"},
		{"SerieBuletin", e.SerieBuletin, "BV"},
		{"NumarBuletin", e.NumarBuletin, "623323"},
		{"CodPostal", e.CodPostal, "2200"},
	} {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}

	// The name is one field: every sampled row carries 2-4 words here, so a
	// split that produced a bare surname would mean the layout changed.
	if e.Nume == "GURGU" {
		t.Error("Nume lost its second half — the Prenume shift is back")
	}
}
