package text

import (
	"fmt"
	"strings"
)

// Insensitive shows how to compare strings in a insensitive way

func testInsensitive() {

	// comparing two strings can become dificult when just moving to lowercase or uppercase is not enough
	// EqualFold is the way we can test it

	word1 := "Something"
	word2 := "somEtHiNg"

	w1w2 := strings.EqualFold(word1, word2)

	word3 := "σ" // sigma lowercase
	word4 := "Σ" // sigma uppercase

	w3w4 := strings.EqualFold(word3, word4)

	fmt.Printf("We can test differnt strings to compare if similar for different languages, in this example words %s and %s are similar %t, and %s - %s are also similar %t\n", word1, word2, w1w2, word3, word4, w3w4)
}
