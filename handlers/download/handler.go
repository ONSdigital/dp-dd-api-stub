package download

import (
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	var data, _ = ioutil.ReadFile("./resources/AF001EW.Armed.csv")

	w.Header().Set("Content-Disposition", "attachment; filename=AF001EW.Armed.csv")
	w.Header().Set("Content-Type", "text/csv")
	w.Write(data)
	w.WriteHeader(200)
}
