package hierarchy

import (
	"net/http"
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	hierarchyId := req.URL.Query().Get(":hierarchyId")
	raw, err := stub.Asset("data/hierarchies/" + hierarchyId + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	w.Write(raw)
	w.WriteHeader(200)
}
