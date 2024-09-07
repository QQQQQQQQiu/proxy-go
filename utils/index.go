package utils

import (
	"flag"
)

func ParseFlag(name string, defaultValue string) string {
	value := flag.String(name, defaultValue, "desc")
	flag.Parse()
	return *value
}