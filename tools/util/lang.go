package util

// SupportedLangs defines the 14 languages struct-ure/kg supports.
//
// According to Common Sense Advisory’s latest report, these are the top 14 online languages:
// English
// Simplified Chinese
// Spanish
// Japanese
// German
// French
// Russian
// Arabic
// Portuguese
// Italian
// Korean
// Dutch
// Hindi
// Chinese traditional
var SupportedLangs = map[string]struct{}{
	"en":           {},
	"zh":           {},
	"es":           {},
	"ja":           {},
	"de":           {},
	"fr":           {},
	"ru":           {},
	"ar":           {},
	"pt":           {},
	"it":           {},
	"ko":           {},
	"nl":           {},
	"hi":           {},
	"zh_classical": {},
}

// LangSupported returns true if the passed argument is supported.
func LangSupported(lang string) bool {
	_, ok := SupportedLangs[lang]
	return ok
}
