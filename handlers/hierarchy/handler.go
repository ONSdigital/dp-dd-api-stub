package hierarchy

import (
	"net/http"
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	hierarchyId := req.URL.Query().Get(":hierarchyId")
	raw, err := stub.Asset("data/hierarchies/" + hierarchyId + ".json")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(404)
		return
	}

	w.Write(raw)
	w.WriteHeader(200)
}
