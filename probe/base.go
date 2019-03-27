package probe

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Out = os.Stdout
	logger.Level = logrus.InfoLevel
	logger.ReportCaller = true
}
