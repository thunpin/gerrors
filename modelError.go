package gerrors

type ModelError interface {
	Code() uint
	Obj() interface{}
	error
}
