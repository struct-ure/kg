package util

// SupportedLangs defines the 14 languages struct-ure/kg supports.
//
// According to Common Sense Advisory’s latest report, these are the top 14 online languages:
var SupportedLangs = map[string]struct{}{
	"en":           {}, // English
	"zh":           {}, // Simplified Chinese
	"es":           {}, // Spanish
	"ja":           {}, // Japanese
	"de":           {}, // German
	"fr":           {}, // French
	"ru":           {}, // Russian
	"ar":           {}, // Arabic
	"pt":           {}, // Portuguese
	"it":           {}, // Italian
	"ko":           {}, // Korean
	"nl":           {}, // Dutch
	"hi":           {}, // Hindi
	"zh_classical": {}, // Chinese traditional
}

// LangSupported returns true if the passed argument is supported.
func LangSupported(lang string) bool {
	_, ok := SupportedLangs[lang]
	return ok
}
