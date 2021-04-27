/*
This package reverses given strings
*/
package reverse

// Reverse function reverses given strings
// Input parameter is a string
func Reverse(word string) (reversed string) {
	buff := make([]string, len(word))

	for i, l := range word {
		buff[len(word)-i-1] = string(l)
	}

	for _, l := range buff {
		reversed += string(l)
	}

	return
}
