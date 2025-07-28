package main
// Error is the type of a parse error
// satisfies the error interface
type Error string
func (e Error) Error() string {
	return string(e)
}

// error is a method of regexp that reports parsing errors
// by panicking with an Error
func (regexp *Regexp) error(err string) {
	panic(Error(err))
}
// Compile returns a parsed representation of the regular expression
func Compile(str string) (regexp *Regexp, err error) {
	regexp = new(Regexp)
	defer func() {
		if e := recover(); e != nil {
			regexp = nil
			err = e.(Error)
		}
	}()
	return regexp.doParse(str), nil
}
