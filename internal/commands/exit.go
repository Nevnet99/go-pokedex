package commands

import (
	"errors"
	"fmt"
)

func Exit() error {
	fmt.Println("Goodbye!")
	return errors.New("exit")
}
