package originstamp

import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}
