package logger

import (
	"log"
	"os"
)

// TODO implement an actual logger

// Default placeholder for logging
var Default = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
