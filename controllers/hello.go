package controllers

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// GetHello controller
func GetHello() (int, map[string]string, error) {
	log.Info("Handling request for /hello")

	status := 200
	body := make(map[string]string)
	body["name"] = "Kristina Vincent"
	body["heading"] = "Portfolio Website"

	return status, body, nil
}
