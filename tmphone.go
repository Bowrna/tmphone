package tmphone

import (
	"regexp"
	"strings"
)

var vowels = map[string]string{
	"அ": "A",
	"ஆ": "A",
	"இ": "I",
	"ஈ": "I",
	"உ": "U",
	"ஊ": "U",
	"எ": "E",
	"ஏ": "E",
	"ஐ": "AI",
	"ஒ": "O",
	"ஓ": "O",
	"ஔ": "AU",
}

var consonants = map[string]string{
	"க": "K",
	"ச": "C",
	"ட": "T",
	"த": "D",
	"ப": "P",
	"ற": "TR",
	"ங": "NG",
	"ஞ": "NJ",
	"ண": "N1",
	"ந": "N",
	"ம": "M",
	"ன": "N1",
	"ய": "Y",
	"ர": "R",
	"ல": "L",
	"வ": "V",
	"ழ": "ZH",
	"ள": "L1",
	"ஃ": "",
	"ஷ": "S1", "ஸ": "S", "ஹ": "H", "ஜ": "J",
}
var modifiers = map[string]string{
	"ா": "3", // Long vowel, no change needed in encoding
	"ி": "4", // Vowel modifier for "i"
	"ீ": "4", // Vowel modifier for "ii"
	"ு": "5", // Vowel modifier for "u"
	"ூ": "5", // Vowel modifier for "uu"
	"ெ": "6", // Vowel modifier for "e"
	"ே": "6", // Vowel modifier for "ee"
	"ை": "7", // Vowel modifier for "ai"
	"ொ": "8", // Vowel modifier for "o"
	"ோ": "8", // Vowel modifier for "oo"
	"ௌ": "9", // Vowel modifier for "au"
	"்": "",  // Tamil Virama or halant (no sound, used for consonant clusters)
}
var compounds = map[string]string{
	"க்க": "K2",
	"ச்ச": "C2",
	"த்த": "D2",
	"ப்ப": "P2",
	"ல்ல": "L2",
	"வ்வ": "V2",
	"ண்ண": "N2",
	"ம்ம": "M2",
	"ற்ற": "TR2",
	"ட்ட": "T2",
	"ஞ்ஞ": "NJ2",

	"ன்ற":  "NR",
	"ண்ட":  "NT",
	"ங்க":  "NK",
	"ஞ்ச":  "NC",
	"ந்த":  "ND",
	"ம்ப":  "MP",
	"ந்ன":  "NN",
	"ற்க":  "RK",
	"ர்ப்": "RP",
	"க்த":  "KT",
}

var (
	regexKey0, _        = regexp.Compile(`[1,2,4-9]`)
	regexKey1, _        = regexp.Compile(`[2,4-9]`)
	regexNonTamil, _    = regexp.Compile(`[\P{Tamil}]`)
	regexAlphaNum, _    = regexp.Compile(`[^0-9A-Z]`)
	regexSpecialCase, _ = regexp.Compile(`^(A|V|T|S|U|M|O)L(K|S)`)
)

// TMphone is the Tamizh-phone tokenizer.
type TMphone struct {
	modCompounds  *regexp.Regexp
	modConsonants *regexp.Regexp
	modVowels     *regexp.Regexp
}

// New returns a new instance of the TMPhone tokenizer.
func New() *TMphone {
	var (
		glyphs []string
		mods   []string
		tm     = &TMphone{}
	)

	// modifiers.
	for t := range modifiers {
		mods = append(mods, t)
	}

	// compounds.
	for t := range compounds {
		glyphs = append(glyphs, t)
	}
	tm.modCompounds, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// consonants.
	glyphs = []string{}
	for k := range consonants {
		glyphs = append(glyphs, k)
	}
	tm.modConsonants, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// vowels.
	glyphs = []string{}
	for k := range vowels {
		glyphs = append(glyphs, k)
	}
	tm.modVowels, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	return tm
}

// Encode encodes a unicode Tamizh string to its Roman TMPhone hash.
// Ideally, words should be encoded one at a time, and not as phrases
// or sentences.
func (t *TMphone) Encode(input string) (string, string, string) {
	// key2 accounts for hard and modified sounds.
	key2 := t.process(input)

	// key1 loses numeric modifiers that denote phonetic modifiers.
	key1 := regexKey1.ReplaceAllString(key2, "")

	// key0 loses numeric modifiers that denote hard sounds, doubled sounds,
	// and phonetic modifiers.
	key0 := regexKey0.ReplaceAllString(key2, "")

	return key0, key1, key2
}

func (t *TMphone) process(input string) string {
	// Remove all non-malayalam characters.
	input = regexNonTamil.ReplaceAllString(strings.Trim(input, ""), "")

	// All character replacements are grouped between { and } to maintain
	// separatability till the final step.

	// Replace and group modified compounds.
	input = t.replaceModifiedGlyphs(input, compounds, t.modCompounds)

	// Replace and group unmodified compounds.
	for k, v := range compounds {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group modified consonants and vowels.
	input = t.replaceModifiedGlyphs(input, consonants, t.modConsonants)
	input = t.replaceModifiedGlyphs(input, vowels, t.modVowels)

	// Replace and group unmodified consonants.
	for k, v := range consonants {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group unmodified vowels.
	for k, v := range vowels {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace all modifiers.
	for k, v := range modifiers {
		input = strings.ReplaceAll(input, k, v)
	}

	// Remove non alpha numeric characters (losing the bracket grouping).
	return regexAlphaNum.ReplaceAllString(input, "")
}

func (k *TMphone) replaceModifiedGlyphs(input string, glyphs map[string]string, r *regexp.Regexp) string {
	for _, matches := range r.FindAllStringSubmatch(input, -1) {
		for _, m := range matches {
			if rep, ok := glyphs[m]; ok {
				input = strings.ReplaceAll(input, m, rep)
			}
		}
	}
	return input
}
