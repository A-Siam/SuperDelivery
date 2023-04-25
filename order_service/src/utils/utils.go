package utils

import (
	"errors"
	"os"
)

func GetEnvWithCheck(s string) (string, error) {
	v := os.Getenv(s)
	if v == "" {
		return "", errors.New("Env variable not set")
	}
	return v, nil
}
