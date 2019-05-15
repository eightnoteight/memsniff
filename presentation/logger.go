package presentation

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func (u *uiContext) runLogger(exitChan chan bool) error {
	logFile, err := os.OpenFile("/tmp/memsniff.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	logger := logrus.New()
	logger.SetOutput(logFile)
	ticker := time.Tick(10 * time.Second)

	for {
		select {
		case <-ticker:
			logger.Print(u.analysis.Report(rand.Int31n(10) == 0))
		case <-exitChan:
			return nil
		}
	}
}
