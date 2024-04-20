package main

import "fmt"
import "bufio"
import "os"
import "strings"


func main() {
	fmt.Println("Enter caeser ciphertext you wish to brute force:")

	reader := bufio.NewReader(os.Stdin)
	ciphertext, _ := reader.ReadString('\n')

	ciphertext = strings.TrimSuffix(ciphertext, "\n")
	ciphertext = strings.ToLower(ciphertext)

	for i, char := range ciphertext {
		if (char < 'a' || char > 'z') && char != ' ' {
			fmt.Printf("Ciphertext contains character '%c' at index %d\n", char, i)
			fmt.Println("Ciphertext should only contain alphabetic characters and spaces")
			os.Exit(1)
		}
	}

	for key := 0; key < 26; key++ {
		strBuilder := strings.Builder{}
		for _, char := range ciphertext {
			if char == ' ' {
				strBuilder.WriteString(string(char))
				continue
			}
			newChar := int(char) + key
			if newChar > 'z' {
				newChar = newChar - 26
			}
			strBuilder.WriteString(string(newChar))
		}
		fmt.Printf("Key '%c': %s\n", 'a' + key, strBuilder.String())
	}
}
