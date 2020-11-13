package wahoo

/*
ErrorResponse - Generic Error for all wahoo http errors
*/
type ErrorResponse struct {
	Msg  string
	Code int
}

func (e ErrorResponse) Error() string {
	return e.Msg
}

//ConstructWahooErrorFromResponse - Constructor that will take a response body and construc the error
func constructWahooErrorFromResponse(responseCode int) ErrorResponse {

	errorToReturn := ErrorResponse{
		Code: responseCode,
	}

	switch responseCode {
	case 400:
		errorToReturn.Msg = "Bad Request -- Your request is invalid."
	case 401:
		errorToReturn.Msg = "Unauthorized -- Your API key is wrong."
	case 403:
		errorToReturn.Msg = "Forbidden -- You do not have access to the specified resource."
	case 404:
		errorToReturn.Msg = "Not Found -- The specified resource could not be found."
	case 405:
		errorToReturn.Msg = "Method Not Allowed -- You tried to access a resource with an invalid method."
	case 406:
		errorToReturn.Msg = "Not Acceptable -- You requested a format that isn't json."
	case 410:
		errorToReturn.Msg = "Gone -- The resource requested has been removed from our servers."
	case 422:
		errorToReturn.Msg = "Unprocessable Entity -- One or more parameters supplied are missing or invalid."
	case 429:
		errorToReturn.Msg = "Too Many Requests -- You are sending too many requests in short period of time."
	case 500:
		errorToReturn.Msg = "Internal Server Error -- We had a problem with our server. Try again later."
	case 503:
		errorToReturn.Msg = "Service Unavailable -- We're temporarily offline for maintenance. Please try again later."
	default:
		errorToReturn.Msg = "Unhandeled Response"
	}

	return errorToReturn

}
