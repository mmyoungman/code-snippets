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
	fmt.Println("Enter plain text you wish to encrypt:")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.ToLower(text)

	for i, char := range text {
		if (char < 'a' || char > 'z') && char != ' ' {
			fmt.Printf("Plain text contains character '%c' at index %d\n", char, i)
			fmt.Println("Plain text should only contain alphabetic characters and spaces")
			os.Exit(1)
		}
	}

	cipherText := ""
	key1 := -19
	key2 := -21
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

	fmt.Println("And now to decrypt...")

	strBuilder := strings.Builder{}
	for i := 0; i < len(cipherText); i++ {
		if text[i] == ' ' {
			strBuilder.WriteString(string(cipherText[i]))
			continue
		}
		newChar := mod(int(cipherText[i] - 'a') - key2, 26)
		for j := 0; j < 26; j++ {
			if mod(key1 * j, 26) == 1 {
				newChar = mod(newChar * j, 26)
				strBuilder.WriteString(string(newChar + 'a'))
				break
			}
		}
	}
	fmt.Println(strBuilder.String())
}
