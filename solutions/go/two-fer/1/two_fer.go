
// This package contains a function for sharing an extra cookie.
package twofer

// This function returns a message based on whether the person's name is provided.
func ShareWith(name string) string {
    // Start with an empty string
	str := ""

    // If no name is provided, "you" is used in the message.
    if name == "" {
        str = "One for you, one for me."
    } else {
    	// If a name is provided, the name is used in the message.
		str = "One for " + name + ", one for me."
    }
	return str
}
