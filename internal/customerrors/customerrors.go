package customerrors

type CustomError string

func (cErr CustomError) Error() string {
	return string(cErr)
}

var (
	CustomErrorNotFound CustomError = "data not found"
)

func ErrorHTTPStatusCode(err error) (statusCode int) {
	switch err {

	}

	return
}
