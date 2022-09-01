package infrastructure

import "os"

func Getenv(key, def string) string {
	e := os.Getenv(key)

	if e == "" {
		return def
	}

	return e
}
