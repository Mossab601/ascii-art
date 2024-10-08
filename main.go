package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	ascii_art "ascii_art/Functions"
)

func main() {
	//==================================================== asci-art1 v1 ==================================
	/*arg:=os.Args[1:]
	// check if numbers args not equal 1
	if len(arg) != 1 {
		fmt.Println("the numbers of args not correct")
		return
	}

	// open file
	file, err := os.Open("./Symbols/standard.txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}*/
	//==================================================== asci-art1-fs  ==================================================
	banner := "standard"
	// Check if the user passed exactly one argument.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	if len(os.Args) == 3 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		} else {
			banner = os.Args[2]
		}
	}

	// Open the file containing the symbol definitions.
	file, err := os.Open("./Symbols/" + banner + ".txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	defer file.Close() // Ensures the file is closed after use.

	sentence := os.Args[1]

	// Checks that the word is not empty
	if len(sentence) == 0 {
		return
	}

	scanner := bufio.NewScanner(file) // Initialize a scanner to read the file line by line

	count := 0
	symbole := []string{}
	symboles := [][]string{}

	// Loop through each line of the text file
	for scanner.Scan() {
		symbole = append(symbole, scanner.Text())
		count++

		// Each ASCII symbol in this format consists of 9 lines
		if count == ascii_art.Hight_symbole_with_ligne {
			symboles = append(symboles, symbole) // Adds the entire symbol to the symbol list
			symbole = []string{}
			count = 0
		}

	}
	// Checks if the file contains at least 95 symbols.
	if len(symboles) < ascii_art.Nbr_char_printble {
		log.Fatal("Please make sure all characters are present in the file.")
	}
	result := ascii_art.PrintWords(ascii_art.Split(sentence), symboles)

	// show result
	fmt.Println(result)
}
