// asciiart/printascii.go:
package asciiart

import (
	"strings"
)

// printAsciiArt prints the given text as ASCII art using the provided map of characters.
func PrintAsciiArt(text string, asciiChars map[byte][]string) string {
	var result strings.Builder

	// Split the text into lines
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		// Print each line of the ASCII art
		for i := 0; i < 8; i++ {
			result.WriteString(printLine(line, asciiChars, i))
			result.WriteString("\n")
		}
	}

	return result.String()
}

// printLine returns a single line of the ASCII art for the given text.
func printLine(text string, asciiChars map[byte][]string, line int) string {
	var lineBuilder strings.Builder

	for _, char := range text {
		if char == '\n' {
			lineBuilder.WriteString("\n")
		} else {
			lineBuilder.WriteString(asciiChars[byte(char)][line])
		}
	}

	return lineBuilder.String()
}
