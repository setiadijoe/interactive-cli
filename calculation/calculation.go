package calculation

import (
	"bytes"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
)

var buf bytes.Buffer
var logInfo = log.New(&buf, "RESULT : ", log.Lmsgprefix)

func SumNumber(x, y float64) error {
	logInfo.Printf("%.2f", x+y)
	logrus.Info(&buf)
	return nil
}

func SubstracNumber(x, y float64) error {
	logInfo.Printf("%.2f", x-y)
	logrus.Info(&buf)
	return nil
}

func MultiNumber(x, y float64) error {
	logInfo.Printf("%.2f", x*y)
	logrus.Info(&buf)
	return nil
}

func DivNumber(x, y float64) error {
	if y == 0.0 {
		return errors.New("Second number must not be zero")
	}

	logInfo.Printf("%.2f", x/y)
	logrus.Info(&buf)
	return nil
}
