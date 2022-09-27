package calculation

import (
	"errors"
	"log"
)

func SumNumber(x, y float64) error {
	log.Printf("Result of sum = %.2f", x+y)
	return nil
}

func SubstracNumber(x, y float64) error {
	log.Printf("Result of substract = %.2f", x-y)
	return nil
}

func MultiNumber(x, y float64) error {
	log.Printf("Result of multiplication = %.2f", x*y)
	return nil
}

func DivNumber(x, y float64) error {
	if y == 0.0 {
		return errors.New("second number must not be zero")
	}

	log.Printf("Result of division = %.2f", x/y)
	return nil
}
