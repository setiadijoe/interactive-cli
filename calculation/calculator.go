package calculation

import (
	"errors"
	"strconv"

	"github.com/manifoldco/promptui"

	"go-cli/pkg/logger"
)

var (
	validate = func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("The input is not number of float64")
		}

		return nil
	}

	templates = &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	firstNumber = promptui.Prompt{
		Label:     "First Number",
		Templates: templates,
		Validate:  validate,
	}

	secondNumber = promptui.Prompt{
		Label:     "Second Number",
		Templates: templates,
		Validate:  validate,
	}

	calc = NewCalculate()
)

func Addition() {
	firstAtt, err := firstNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	secondAtt, err := secondNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	first := parsingFloat(firstAtt)
	second := parsingFloat(secondAtt)

	calc.SumNumber(first, second)
}

func Substraction() {
	firstAtt, err := firstNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	secondAtt, err := secondNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	first := parsingFloat(firstAtt)
	second := parsingFloat(secondAtt)

	calc.SubstracNumber(first, second)
}

func Multipiclation() {
	firstAtt, err := firstNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	secondAtt, err := secondNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	first := parsingFloat(firstAtt)
	second := parsingFloat(secondAtt)

	calc.MultiNumber(first, second)
}

func Divition() {
	firstAtt, err := firstNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	secondAtt, err := secondNumber.Run()
	if err != nil {
		logger.Error(err)
	}

	first := parsingFloat(firstAtt)
	second := parsingFloat(secondAtt)

	err = calc.DivNumber(first, second)
	if err != nil {
		logger.Error(err)
	}
}

func parsingFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
