package host

import "os"

var (
	env = os.Getenv("Environment")
)

func IsProduction() bool {
	return env == "production"
}
