package messages

import "net/http"

var SuccessMessages = map[int]string{
	http.StatusOK:      "request success",
	http.StatusCreated: "entity created",
}
