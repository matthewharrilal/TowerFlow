package main

type errorString struct {
	str string
}

func (err *errorString) Error() string {
	// The official error struct has an interface looks for anyting with an error method so can extend functionality such as we are doing here
	return err.str
}
