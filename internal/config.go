package internal

import (
	"errors"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

func Setup() {
	// check if config dir exists
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	_, err = os.ReadDir(path.Join(home, ".termhub"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(path.Join(home, ".termhub"), 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}
