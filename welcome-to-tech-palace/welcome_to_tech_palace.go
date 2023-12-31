package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var name string = strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return name
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	var newStr string = strings.ReplaceAll(oldMsg, "*", "")
	var newStr2 = strings.TrimSpace(newStr)

	return newStr2

}
