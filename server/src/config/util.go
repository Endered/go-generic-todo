package config

import (
	"os"
	"todo/src/util"
)

func getenv(env string) util.Optional[string] {
	str := os.Getenv(env)
	if str == "" {
		return util.None[string]()
	} else {
		return util.Some(str)
	}
}
