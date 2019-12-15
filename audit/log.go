package audit

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	log.SetOutput(os.Stdout)
}
