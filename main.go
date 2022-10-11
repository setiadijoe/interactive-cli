package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"

	"go-cli/calculation"
	"go-cli/pkg/logger"
)

func main() {

	selectPrompt := promptui.Select{
		Label: "Number Calculation",
		Items: []string{
			"Addition", "Substraction", "Multipiclation", "Divition", "EXIT",
		},
	}

	for {
		_, result, err := selectPrompt.Run()
		if err != nil {
			logger.Error(err)
			continue
		}

		fmt.Printf("\nYou choose %+v\n", result)

		switch result {
		case "Addition":
			calculation.Addition()
			continue
		case "Substraction":
			calculation.Substraction()
			continue
		case "Multipiclation":
			calculation.Multipiclation()
			continue
		case "Divition":
			calculation.Divition()
			continue
		case "EXIT":
			renderImages()
			os.Exit(0)
		}
	}
}

func renderImages() {
	path, _ := filepath.Abs("art.txt")
	b, _ := ioutil.ReadFile(path)
	fmt.Printf("\033[1;36m%s\033[0m\n", string(b))
}
