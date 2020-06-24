package utils

import (
	"fmt"
	"log"
)

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Warn(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}
