package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Message returns the messages from model to the controller
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond returns messages from the controller to the user
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//FilterArrayParams reads a filter array parameter from an HTTP request
func FilterArrayParams(r *http.Request, b []byte) ([]byte, error) {
	vals := r.URL.Query()
	filter := vals.Get("filter")
	if len(filter) == 0 {
		return b, nil
	}
	words := strings.Split(filter, ",")
	prevJSON := []map[string]interface{}{}
	json.Unmarshal(b, &prevJSON)
	newJSON := []map[string]interface{}{}
	for _, pjs := range prevJSON {
		njs := map[string]interface{}{}
		for _, v := range words {
			tmp, ok := pjs[v]
			if ok {
				njs[v] = tmp
			}
		}
		newJSON = append(newJSON, njs)
	}
	bs, err := json.Marshal(newJSON)
	return bs, err
}

//GetUintParam reads a uint parameter from an HTTP request
func GetUintParam(r *http.Request, param string) (uint, error) {
	sid, ok := mux.Vars(r)[param]
	if !ok {
		return 0, errors.New(ParameterNotExistsErr(param))
	}
	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

//GetUUIDParam reads a UUID parameter from an HTTP request
func GetUUIDParam(r *http.Request, param string) (string, error) {
	sid, ok := mux.Vars(r)[param]
	if !ok {
		return "", errors.New(ParameterNotExistsErr(param))
	}
	if len(strings.Split(sid, "-")) != 5 {
		return "", errors.New(ParameterNotUUIDErr(param))
	}
	return sid, nil
}

//GetStringParam reads a string parameter from an HTTP request
func GetStringParam(r *http.Request, param string) (string, error) {
	sid, ok := mux.Vars(r)[param]
	if !ok {
		return "", errors.New(ParameterNotExistsErr(param))
	}
	return sid, nil
}

//GetIntParam reads an integer parameter from an HTTP request
func GetIntParam(r *http.Request, param string) (int, error) {
	sid, ok := mux.Vars(r)[param]
	if !ok {
		return 0, errors.New(ParameterNotExistsErr(param))
	}
	id, err := strconv.Atoi(sid)
	if err != nil {
		return 0, err
	}
	return id, nil
}
