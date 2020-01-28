package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jenusek/resourcepack/internal/session"
)

func RenderResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		RenderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RenderError(w http.ResponseWriter, err error) {
	sessionError, ok := err.(session.Error)
	if !ok {
		sessionError = session.ServerError(err)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(sessionError.StatusCode)
	w.Write([]byte(fmt.Sprintf("%s", sessionError.Description)))
}
