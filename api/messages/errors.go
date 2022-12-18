package messages

import "net/http"

var ErrorsMessages = map[int]string{
	http.StatusUnprocessableEntity: "unprocessed request",
	http.StatusNotFound:            "entity not found",
}
