package customErrors

type FieldErrorBiddingError struct{}

func (err FieldErrorBiddingError) Error() string {
	const message string = "Invalid data type."
	return message
}

func FieldBiddingErrorWrapper() error {
	return FieldErrorBiddingError{}
}

type InvalidParamsErrorBiddingError struct{}

func (err InvalidParamsErrorBiddingError) Error() string {
	const message string = "Invalid parameters"
	return message
}

func InvalidParamsErrorBiddingErrorWrapper() error {
	return InvalidParamsErrorBiddingError{}
}
