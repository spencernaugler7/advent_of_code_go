package errorhandling

import (
	"errors"
	"fmt"
	"io"
)

type User struct {
	Id int
}

var (
	ErrInvalidID    = errors.New("invalid user id")
	ErrUserNotFound = errors.New("user not found")
)

func getUser(id int) (*User, error) {
	if id < 0 {
		return nil, ErrInvalidID
	}
	// ...
	return nil, ErrUserNotFound
}

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["errorHandling"] = Run
	return registrar
}

func Run(writer io.Writer) {
	user, err := getUser(-1)
	if err != nil {
		if errors.Is(err, ErrInvalidID) {
			// Specific handling
		}
	}

	fmt.Fprintln(writer, user.Id)
}
