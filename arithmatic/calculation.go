package arithmatic

import (
	"errors"

	"go-cli/pkg/logger"
)

type Arithmatic struct {
}

func NewArithmatic() *Arithmatic {
	return &Arithmatic{}
}

func (c *Arithmatic) SumNumber(x, y float64) error {
	logger.InfoF("%.2f", x+y)
	return nil
}

func (c *Arithmatic) SubstracNumber(x, y float64) error {
	logger.InfoF("%.2f", x-y)
	return nil
}

func (c *Arithmatic) MultiNumber(x, y float64) error {
	logger.InfoF("%.2f", x*y)
	return nil
}

func (c *Arithmatic) DivNumber(x, y float64) error {
	if y == 0.0 {
		return errors.New("Second number must not be zero")
	}

	logger.InfoF("%.2f", x/y)
	return nil
}
