package calculation

import (
	"errors"

	"go-cli/pkg/logger"
)

type Calculate struct {
}

func NewCalculate() *Calculate {
	return &Calculate{}
}

func (c *Calculate) SumNumber(x, y float64) error {
	logger.InfoF("%.2f", x+y)
	return nil
}

func (c *Calculate) SubstracNumber(x, y float64) error {
	logger.InfoF("%.2f", x-y)
	return nil
}

func (c *Calculate) MultiNumber(x, y float64) error {
	logger.InfoF("%.2f", x*y)
	return nil
}

func (c *Calculate) DivNumber(x, y float64) error {
	if y == 0.0 {
		return errors.New("Second number must not be zero")
	}

	logger.InfoF("%.2f", x/y)
	return nil
}
