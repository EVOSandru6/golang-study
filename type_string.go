package main

import "fmt"

func main() {
	first := "First string\a"
	second := `Second string \n`

	fmt.Println(first, second)
	fmt.Println(len(second))

	stringOperations()
}

func stringOperations() {
	first := "first--"
	second := "second"
	space := " "

	ukranian := "Украиньська"
	runeUkranian := []rune(ukranian)

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println(first + space + second)

	fmt.Println(first[1])
	fmt.Println(string(first[1]))
	fmt.Println(first[3:])
	fmt.Println(`________________________`,first[:7])
	fmt.Println(first[2:7], "\n")

	fmt.Println(ukranian[1])
	fmt.Println(ukranian[3:])
	fmt.Println(ukranian[:7])
	fmt.Println(ukranian[2:7], "\n")

	fmt.Println(runeUkranian[1])
	fmt.Println(runeUkranian[3:])
	fmt.Println(runeUkranian[:7])
	fmt.Println(runeUkranian[2:7], "\n")

	fmt.Println(string(runeUkranian[1]))
	fmt.Println(string(runeUkranian[3:]))
	fmt.Println(string(runeUkranian[:7]))
	fmt.Println(string(runeUkranian[2:7]), "\n")


}
