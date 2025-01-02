package middleware

import (
	"encoding/json"
	"net/http"
	"reflect"
	"trivium/internal/presentation/format"
)

type HandlerFunction func(input interface{}) (interface{}, error)

func JsonMiddleware(handler HandlerFunction, inputType interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := reflect.New(reflect.TypeOf(inputType).Elem()).Interface()
		if r.Body != nil {
			defer r.Body.Close()
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
				return
			}
		}

		response, err := handler(input)
		if err != nil {
			format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		format.WriteSuccessResponse(w, http.StatusOK, response)
	}
}
