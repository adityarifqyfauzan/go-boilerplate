package helper

import "strings"

func ToPlural(word string) string {
	lowerWord := strings.ToLower(word)

	// Rule khusus
	if strings.HasSuffix(lowerWord, "y") && !strings.HasSuffix(lowerWord, "ay") && !strings.HasSuffix(lowerWord, "ey") && !strings.HasSuffix(lowerWord, "iy") && !strings.HasSuffix(lowerWord, "oy") && !strings.HasSuffix(lowerWord, "uy") {
		return word[:len(word)-1] + "ies"
	} else if strings.HasSuffix(lowerWord, "s") || strings.HasSuffix(lowerWord, "x") || strings.HasSuffix(lowerWord, "z") || strings.HasSuffix(lowerWord, "ch") || strings.HasSuffix(lowerWord, "sh") {
		return word + "es"
	} else {
		return word + "s"
	}
}

func ToPascalCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, "")
}
