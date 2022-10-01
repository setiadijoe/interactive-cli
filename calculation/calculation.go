package calculation

import (
	"errors"
	"go-cli/pkg/logger"
)

type Calculate struct {
	log *logger.MyLog
}

func NewCalculate() *Calculate {
	return &Calculate{
		log: logger.NewMyLog(),
	}
}

func (c *Calculate) SumNumber(x, y float64) error {
	c.log.InfoF("%.2f", x+y)
	return nil
}

func (c *Calculate) SubstracNumber(x, y float64) error {
	c.log.InfoF("%.2f", x-y)
	return nil
}

func (c *Calculate) MultiNumber(x, y float64) error {
	c.log.InfoF("%.2f", x*y)
	return nil
}

func (c *Calculate) DivNumber(x, y float64) error {
	if y == 0.0 {
		return errors.New("Second number must not be zero")
	}

	c.log.InfoF("%.2f", x/y)
	return nil
}
