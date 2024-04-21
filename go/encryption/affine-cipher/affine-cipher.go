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

	key1 := -19
	key2 := -21

	key1_inverse := 0
	for i := 0; i < 26; i++ {
		if mod(key1 * i, 26) == 1 {
			key1_inverse = i
			break
		}
	}

	cipherText := ""

	// encrypt
	{
		strBuilder := strings.Builder{}
		for i := 0; i < len(text); i++ {
			if text[i] == ' ' {
				strBuilder.WriteString(string(text[i]))
				continue
			}
			newChar := mod(((int(text[i]) - 'a') * key1) + key2, 26)
			strBuilder.WriteString(string(newChar + 'a'))
		}
		cipherText = strBuilder.String()
		fmt.Printf("Key '%d/%d': %s\n", key1, key2, cipherText)
	}

	// decrypt
	strBuilder := strings.Builder{}
	for i := 0; i < len(cipherText); i++ {
		if text[i] == ' ' {
			strBuilder.WriteString(string(cipherText[i]))
			continue
		}
		newChar := mod((int(cipherText[i] - 'a') - key2) * key1_inverse, 26)
		strBuilder.WriteString(string(newChar + 'a'))
	}
	fmt.Println(strBuilder.String())
}
