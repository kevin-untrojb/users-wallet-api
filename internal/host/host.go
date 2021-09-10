package host

import "os"

var (
	env = os.Getenv("Environment")
)
func IsTesting() bool {
	return env == "test"
}