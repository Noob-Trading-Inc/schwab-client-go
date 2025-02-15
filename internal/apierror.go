package internal

import (
	"strings"

	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

type ApiError struct {
	Errors []Error
	JSon   string
}

type Error struct {
	ID     string
	Status int
	Title  string
	Detail string
}

func NewApiError(errResponse string) (apierror ApiError) {
	if strings.Trim(errResponse, "\n ") != "" {
		util.Deserialize(errResponse, &apierror)
		apierror.JSon = errResponse
	}
	return
}

func (c ApiError) Error() string {
	return c.JSon
}
