package hierarchies

import (
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	raw, err := stub.Asset("data/hierarchies.json")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(404)
		return
	}

	w.Write(raw)
	w.WriteHeader(200)
}
