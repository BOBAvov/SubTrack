package sub_track

type ErrorPlus interface {
	Error() string
	Loc() string
}

type MyError struct {
	Text     string
	Location string
}

func (e MyError) Error() string {
	return e.Text
}
func (e MyError) Loc() string {
	return e.Location
}

func NewMyError(text string, location string) MyError {
	return MyError{Text: text, Location: location}
}
