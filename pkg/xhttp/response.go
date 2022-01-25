package xhttp

import (
	"encoding/json"
	"net/http"

	"github.com/Codename-Uranium/tunnel/pkg/xerror"
	"go.uber.org/zap"
)

var jsonOkString = []byte(`"OK"`)

// JSONResponse calls the closure and respond with data or error.
// If no error, but the interface is nil it will write "OK" to the response writer.
func JSONResponse(w http.ResponseWriter, closure func() (interface{}, error)) {
	data, err := closure()
	if err != nil {
		WriteJsonError(w, err)
		return
	}

	if data != nil {
		bs, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			// log as warning, not error because it's very likely the API misuse
			zap.S().Warn("failed to marshal %T into JSON: %v", data, err)
			WriteJsonError(w, err)
			return
		}

		WriteData(w, bs)
		return
	}

	WriteData(w, jsonOkString)
}

func WriteData(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		zap.L().Error("can't write response", zap.Error(err))
	}
}

func WriteJsonError(w http.ResponseWriter, err error) {
	if err == nil {
		zap.L().Error("writeError: nil error passed")
		return
	}

	code, message := xerror.ErrorToHttpResponse(err)
	w.WriteHeader(code)
	w.Header().Set("content-type", "application/json")
	if _, err = w.Write(message); err != nil {
		zap.L().Error("can't write response", zap.Error(err))
	}
}
