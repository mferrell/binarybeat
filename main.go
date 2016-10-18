package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/mferrell/binarybeat/beater"
)

func main() {
	err := beat.Run("binarybeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
