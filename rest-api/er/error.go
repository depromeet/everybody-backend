package er

// WrongTimeFormatErr 는 잘못된 시간을 나타내는
type WrongTimeFormatErr struct {
	Err error
}

func (e *WrongTimeFormatErr) Error() string {
	return "잘못된 시간 포맷입니다: " + e.Err.Error()
}
