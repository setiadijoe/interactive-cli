package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"

	art "go-cli/arithmatic"
	generatoruuid "go-cli/generator-uuid"
	"go-cli/pkg/logger"
)

func main() {

	selectPrompt := promptui.Select{
		Label: "Number Calculation",
		Items: []string{
			"Addition", "Substraction", "Multiplication", "Divition", "Generate UUID", "EXIT",
		},
	}
	renderImages("intro.txt")

	for {
		_, result, err := selectPrompt.Run()
		if err != nil {
			logger.Error(err)
			continue
		}

		fmt.Printf("\nYou choose %+v\n", result)

		switch result {
		case "Addition":
			art.Addition()
			continue
		case "Substraction":
			art.Substraction()
			continue
		case "Multiplication":
			art.Multiplication()
			continue
		case "Divition":
			art.Divition()
			continue
		case "Generate UUID":
			generatoruuid.GenerateUUIDV4()
			continue
		case "EXIT":
			renderImages("exit.txt")
			os.Exit(0)
		}
	}
}

func renderImages(fileName string) {
	path, _ := filepath.Abs(fileName)
	b, _ := ioutil.ReadFile(path)
	fmt.Printf("\033[1;36m%s\033[0m\n", string(b))
}
