package dimension

import (
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	view := req.URL.Query().Get("view")
	switch view {
	case "hierarchy":
		hierarchyHandler(w, req)
		break
	case "list":
		fallthrough
	default:
		listHandler(w, req)
		break
	}
}
