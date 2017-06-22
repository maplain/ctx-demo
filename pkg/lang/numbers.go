package lang

type Lang string

const (
	Swedish    Lang = "Swedish"
	Portuguese Lang = "Portuguese"
	Latin      Lang = "Latin"
	Finnish    Lang = "Finnish"
	Catalan    Lang = "Catalan"
	German     Lang = "German"
	Greek      Lang = "Greek"
	Spanish    Lang = "Spanish"
	Welsh      Lang = "Welsh"
	French     Lang = "French"
	Norwegian  Lang = "Norwegian"
	Sanskrit   Lang = "Sanskrit"
	English    Lang = "English"
	Russian    Lang = "Russian"
	Polish     Lang = "Polish"
	Arabic     Lang = "Arabic"
	Basque     Lang = "Basque"
	Swahili    Lang = "Swahili"
	Italian    Lang = "Italian"
)

var numbers = map[Lang][]string{
	Swedish:    {"ett", "två", "tre", "fyra", "fem", "sex", "sju", "åtta", "nio", "tio"},
	Portuguese: {"um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez"},
	Latin:      {"unus", "duo", "tres", "quattor", "quinque", "sex", "septem", "octo", "novem", "decem"},
	Finnish:    {"yksi", "kaksi", "kolme", "neljä", "viisi", "kuusi", "seitsemän", "kahdeksan", "yhdeksan", "kymmenen"},
	Catalan:    {"un", "dos", "tres", "quatre", "cinc", "sis", "set", "vuit", "nou", "deu"},
	German:     {"eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn"},
	Greek:      {"eîs", "dúo", "treiîs", "téssares", "pénte", "hex", "heptá", "októ", "ennéa", "déka"},
	Spanish:    {"uno", "dos", "tres", "cuatro", "cinco", "seis", "seite", "ocho", "nueve", "diez"},
	Welsh:      {"un", "dau", "tri", "pedwar", "pump", "chwech", "saith", "wyth", "naw", "deg"},
	French:     {"un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix"},
	Norwegian:  {"en", "to", "tre", "fire", "fem", "seks", "syv", "åtte", "ni", "ti"},
	Sanskrit:   {"ekab", "dvi", "trayah", "chatvarah", "pancha", "shash", "sapta", "ashta", "nava", "dasha"},
	English:    {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
	Russian:    {"odin", "dva", "tri", "chetyre", "pyat'", "shest'", "sem'", "vosem'", "devyat'", "desyat'"},
	Polish:     {"jeden", "dwa", "trzy", "cztery", "pięć", "sześć", "siedem", "osiem", "dziewięć", "dziesięć"},
	Arabic:     {"wahed", "ithnain", "thelatha", "arba'a", "hamza", "sitta", "seba'a", "themania", "tisa'a", "ashara"},
	Basque:     {"bat", "bi", "hiru", "lau", "bost", "sei", "zazpi", "zortzi", "beheratzi", "hamar"},
	Swahili:    {"moja", "bili", "tatu", "'nne", "tano", "sita", "sabbah", "nanne", "tissa", "kumi"},
	Italian:    {"uno", "due", "tre", "quattro", "cinque", "sei", "sette", "otto", "nove", "dieci"},
}
