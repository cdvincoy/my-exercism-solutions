// This package provides tools for 
// generating text for the new display in "Tech Palace".
package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customerName := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + customerName
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	displayMsg:= strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
    return displayMsg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanMsg := strings.ReplaceAll(oldMsg,"*","")
    cleanMsg2 := strings.TrimSpace(cleanMsg)
    return cleanMsg2
}
