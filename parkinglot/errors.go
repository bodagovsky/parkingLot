package parkinglot

const (
	SlotIsEmptyError = `The slot you are trying to leave is empty`
)

type SlotIsEmpty struct {
	msg string
}

func SlotIsEmptyErr(message string) *SlotIsEmpty {
	return &SlotIsEmpty{
		msg: message,
	}
}

func (e *SlotIsEmpty) Error() string {
	return e.msg
}

