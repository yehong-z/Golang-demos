package trace

import (
	"errors"
	"fmt"
	"runtime/debug"

	log "github.com/jeanphorn/log4go"
)

func A() error {
	err := B()
	if err != nil {
		log.LOGGER("Test").Error(err.Error(), fmt.Sprintf("%s", debug.Stack()))
	}
	return err
}

func B() error {
	return C()
}

func C() error {
	return D()
}

func D() error {
	return E()
}

func E() error {
	return errors.New("an error")
}
