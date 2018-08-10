package gerrors

type ModelError interface {
	Code() int
	Obj() interface{}
	error
}
