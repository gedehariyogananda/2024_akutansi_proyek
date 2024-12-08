package Utils

import "strings"

func FormatUsernameClient(input string) string {
	trimmedInput := strings.TrimSpace(input)
	withUnderscore := strings.ReplaceAll(trimmedInput, " ", "_")
	formattedString := strings.ToLower(withUnderscore)
	return formattedString
}
