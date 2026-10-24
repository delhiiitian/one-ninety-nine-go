// Package greeting builds friendly messages.
//
// A package is a folder of .go files that share one package name.
// Names that start with a capital letter are exported (visible to other packages).
package greeting

import "strings"

// DefaultName is used when no name is given. Exported: other packages can read it.
const DefaultName = "friend"

// Hello returns a greeting in the chosen language. Exported.
func Hello(name, lang string) string {
	return word(lang) + ", " + tidy(name) + "!"
}

// word is unexported (lowercase): only code inside package greeting can call it.
func word(lang string) string {
	switch lang {
	case "hi":
		return "Namaste"
	case "es":
		return "Hola"
	case "fr":
		return "Bonjour"
	default:
		return "Hello"
	}
}

func tidy(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return DefaultName
	}
	return strings.ToUpper(name[:1]) + name[1:]
}
