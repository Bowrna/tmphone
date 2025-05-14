package tmphone

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tm := New()
	// tmphone_test.go:24: For input மிகவும் => got (MKVM, MKVM, M4KV5M), want (MKVM, MK5VM, MK5VM)
	// tmphone_test.go:24: For input பஞ்சவர்ணம் => got (PNCVRNM, PNCVRN1M, PNCVRN1M), want (, , )
	// tmphone_test.go:24: For input வௌவால் => got (VVL, VVL, V9V3L), want (, , )
	// tmphone_test.go:24: For input அங்காடி => got (ANKD, ANKD, ANK3D4), want (, , )
	// tmphone_test.go:24: For input மோர் => got (MR, MR, M8R), want (, , )
	// tmphone_test.go:24: For input தமிழ் => got (THMZH, THMZH, THM4ZH), want (TML, TM3L, TM3L)
	// tmphone_test.go:24: For input சிப்பாய் => got (SPY, SP2Y, S4P23Y), want (, , )
	// tmphone_test.go:24: For input தண்ணீர்  => got (THNR, THN2R, THN24R), want (, , )
	// tmphone_test.go:24: For input திங்கள் => got (THNKL, THNKL1, TH4NKL1), want (, , )
	tests := map[string][3]string{
		"தமிழ்":      {"TML", "TM3L", "TM3L"},
		"மிகவும்":    {"MKVM", "MKVM", "M4KV5M"},
		"சிப்பாய்":   {"SPY", "SP2Y", "S4P23Y"},
		"தண்ணீர் ":   {"THNR", "THN2R", "THN24R"},
		"பஞ்சவர்ணம்": {"PNCVRNM", "PNCVRN1M", "PNCVRN1M"},
		"திங்கள்":    {"THNKL", "THNKL1", "TH4NKL1"},
		"மோர்":       {"MR", "MR", "M8R"},
		"வௌவால்":     {"VVL", "VVL", "V9V3L"},
		"அங்காடி":    {"ANKD", "ANKD", "ANK3D4"},
	}

	for input, expected := range tests {
		k0, k1, k2 := tm.Encode(input)
		if k0 != expected[0] || k1 != expected[1] || k2 != expected[2] {
			t.Errorf("For input %s => got (%s, %s, %s), want (%s, %s, %s)",
				input, k0, k1, k2, expected[0], expected[1], expected[2])
		}
	}
}
