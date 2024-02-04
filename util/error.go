package util

type PError struct {
  Msg string
}

func (e *PError) Error() string {
	return e.Msg
}