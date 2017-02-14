package dimension

import (
	"net/http"
	"fmt"
	"errors"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	view := req.URL.Query().Get("view")
	switch view {
	case "hierarchy":
		hierarchyHandler(w, req)
	case "list":
		fallthrough
	case "":
		listHandler(w, req)
		break
	default:
		err := errors.New("Invalid query view string")
		fmt.Println(err.Error())
		w.WriteHeader(404)
	}
}


func LegacyHandler(w http.ResponseWriter, req *http.Request) {
	view := req.URL.Query().Get("view")
	switch view {
	case "hierarchy":
		hierarchyHandler(w, req)
	case "list":
		fallthrough
	case "":
		legacyListHandler(w, req)
		break
	default:
		err := errors.New("Invalid query view string")
		fmt.Println(err.Error())
		w.WriteHeader(404)
	}
}