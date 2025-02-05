package install

import (
	"os"
)

// TODO : Depending on which deployment choose, ***
// the behavior of `install` and `run` will be different,
// so need more considering.

func Run() {
	configPath := os.Getenv("TASKS_CONFIG_PATH")
	if configPath == "" {
		install()
	} else {
		return
	}
}

func install() {

}
