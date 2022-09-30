package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"

	"go-cli/calculation"
)

var buf bytes.Buffer
var logError = log.New(&buf, "ERROR : ", log.Lmsgprefix)

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

	for {
		_, result, err := selectPrompt.Run()
		if err != nil {
			logError.Print(err)
			logrus.Error(&buf)
			continue
		}

		fmt.Printf("\nYou choose %+v\n", result)

		switch result {
		case "Addition":
			firstAtt, err := firstNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			secondAtt, err := secondNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			first := parsingFloat(firstAtt)
			second := parsingFloat(secondAtt)

			calculation.SumNumber(first, second)
			continue
		case "Substraction":
			firstAtt, err := firstNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			secondAtt, err := secondNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			first := parsingFloat(firstAtt)
			second := parsingFloat(secondAtt)

			calculation.SubstracNumber(first, second)
			continue
		case "Multipiclation":
			firstAtt, err := firstNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			secondAtt, err := secondNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			first := parsingFloat(firstAtt)
			second := parsingFloat(secondAtt)

			calculation.MultiNumber(first, second)
			continue
		case "Divition":
			firstAtt, err := firstNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			secondAtt, err := secondNumber.Run()
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}

			first := parsingFloat(firstAtt)
			second := parsingFloat(secondAtt)

			err = calculation.DivNumber(first, second)
			if err != nil {
				logError.Print(err)
				logrus.Error(&buf)
			}
			continue
		case "EXIT":
			renderImages()
			os.Exit(0)
		}
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
