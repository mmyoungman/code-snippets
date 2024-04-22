package main

import "fmt"
import "bufio"
import "os"
import "strings"

func mod(value int, mod int) int {
	value = value % mod
	if value < 0 {
		value += mod
	}
	return value
}

func main() {
	fmt.Println("Enter text you wish to encrypt:")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.ToLower(text)

	for i, char := range text {
		if (char < 'a' || char > 'z') && char != ' ' {
			fmt.Printf("Text contains character '%c' at index %d\n", char, i)
			fmt.Println("Text should only contain alphabetic characters and spaces")
			os.Exit(1)
		}
	}

	//key := "abcdefghijklmnopqrstuvwxyz"
	key := "zyxwvutsrqponmlkjihgfedcba"

	cipherText := ""

	// encrypt
	{
		strBuilder := strings.Builder{}
		for i := 0; i < len(text); i++ {
			if text[i] == ' ' {
				strBuilder.WriteString(string(text[i]))
				continue
			}
			newChar := key[text[i] - 'a']
			strBuilder.WriteString(string(newChar))
		}
		cipherText = strBuilder.String()
		fmt.Printf("Key '%s': %s\n", key, cipherText)
	}

	// decrypt
	strBuilder := strings.Builder{}
	for i := 0; i < len(cipherText); i++ {
		if cipherText[i] == ' ' {
			strBuilder.WriteString(string(cipherText[i]))
			continue
		}
		for j := 0; j < len(key); j++ {
			if cipherText[i] == key[j] {
				strBuilder.WriteString(string(j + 'a'))
				break
			}
		}
	}
	fmt.Println(strBuilder.String())
}
