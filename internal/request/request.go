package request

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/mritunjaysharma394/lambda-proxy-http-wrapper/internal/response"
)

var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type APIGatewayProxyRequest struct {
	Resource              string            `json:"resource"` // The resource path defined in API Gateway
	Path                  string            `json:"path"`     // The url path for the caller
	HTTPMethod            string            `json:"httpMethod"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	Body                  string            `json:"body"`
	IsBase64Encoded       bool              `json:"isBase64Encoded,omitempty"`
}

func show(req APIGatewayProxyRequest) (response.APIGatewayProxyResponse, error) {
	// Get the `isbn` query string parameter from the request and
	// validate it.
	isbn := req.QueryStringParameters["isbn"]
	if !isbnRegexp.MatchString(isbn) {
		return clientError(http.StatusBadRequest)
	}

	// Fetch the book record from the database based on the isbn value.
	// bk, err := getItem(isbn)
	// if err != nil {
	// 	return serverError(err)
	// }
	// if bk == nil {
	// 	return clientError(http.StatusNotFound)
	// }

	// // The APIGatewayProxyResponse.Body field needs to be a string, so
	// // we marshal the book record into JSON.
	// js, err := json.Marshal(bk)
	// if err != nil {
	// 	return serverError(err)
	// }

	// Return a response with a 200 OK status and the JSON book record
	// as the body.
	return response.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (response.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return response.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (response.APIGatewayProxyResponse, error) {
	return response.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
