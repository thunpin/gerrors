package gerrors

type Error struct {
	Code    int
	Message string
	Obj     interface{}
}

func (e Error) Error() string {
	return e.Message
}
