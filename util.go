package data

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes a JSON response with the given status code and data.
func WriteJSON(w http.ResponseWriter, status int, data ...interface{}) error {
	if status == http.StatusNoContent {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		return nil
	}

	js, err := json.Marshal(data[0])
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// ErrorJSON writes a JSON response with the given error message and status code.
func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, statusCode, theError)
}

// SuccessResponse writes a JSON response with the given data and status code.
func SuccessResponse(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"data":    &data,
	})
}

// NotFound writes a JSON response with the given data and status code.
func BadRequest(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "data validation failed",
		"data":    &data,
	})
}


// NotFound writes a JSON response with the given data and status code.
func InternalServerError(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "internal server error",
		"data":    &data,
	})
}