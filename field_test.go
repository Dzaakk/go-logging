package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "My_Username").Info("Hello World")

	logger.WithField("username", "user1").
		WithField("name", "user1_name").
		Info("Hello World")
}
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "username2",
		"name":     "user2_name",
	}).Info("Hello World")
}
