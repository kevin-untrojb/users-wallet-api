package host

import "os"

var (
	env = os.Getenv("ENVIRONMENT")
)

func IsProduction() bool {
	return env == "prod"
}
