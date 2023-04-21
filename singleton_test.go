package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("INFO")
	logrus.Warn("WARN")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("INFO 2")
	logrus.Warn("WARN 2")
}
