package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/manifoldco/promptui"

	"go-cli/calculation"
)

func main() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("The input is not number of float64")
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	selectPrompt := promptui.Select{
		Label: "Number Calculation",
		Items: []string{
			"Addition", "Substraction", "Multipiclation", "Divition", "EXIT",
		},
	}

	firstNumber := promptui.Prompt{
		Label:     "First Number",
		Templates: templates,
		Validate:  validate,
	}

	secondNumber := promptui.Prompt{
		Label:     "Second Number",
		Templates: templates,
		Validate:  validate,
	}

	_, result, err := selectPrompt.Run()
	if err != nil {
		fmt.Println("Prompt failed :", err)
	}

	fmt.Printf("\nYou choose %+v\n", result)

	switch result {
	case "Addition":
		firstAtt, err := firstNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		secondAtt, err := secondNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		first := parsingFloat(firstAtt)
		second := parsingFloat(secondAtt)

		calculation.SumNumber(first, second)
	case "Substraction":
		firstAtt, err := firstNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		secondAtt, err := secondNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		first := parsingFloat(firstAtt)
		second := parsingFloat(secondAtt)

		calculation.SubstracNumber(first, second)
	case "Multipiclation":
		firstAtt, err := firstNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		secondAtt, err := secondNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		first := parsingFloat(firstAtt)
		second := parsingFloat(secondAtt)

		calculation.MultiNumber(first, second)
	case "Divition":
		firstAtt, err := firstNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		secondAtt, err := secondNumber.Run()
		if err != nil {
			fmt.Println(err)
		}

		first := parsingFloat(firstAtt)
		second := parsingFloat(secondAtt)

		err = calculation.DivNumber(first, second)
		log.Fatalln(err)
	case "EXIT":
		renderImages()
		os.Exit(0)
	}
}

func parsingFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func renderImages() {
	path, _ := filepath.Abs("art.txt")
	b, _ := ioutil.ReadFile(path)
	fmt.Printf("\033[1;36m%s\033[0m\n", string(b))
}
