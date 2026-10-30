package errorhandling

import (
	"errors"
	"fmt"
)


type User struct {
	Id int
}
	

var ErrInvalidID = errors.New("invalid user id")
var ErrUserNotFound = errors.New("user not found")

func getUser(id int) (*User, error) {
    if id < 0 {
        return nil, ErrInvalidID
    }
    // ...
    return nil, ErrUserNotFound
}


func Init() {
	
	user, err := getUser(-1)
	if err != nil {
		if errors.Is(err, ErrInvalidID) {
			// Specific handling
		}
	}

	fmt.Println(user.Id)
}
