package datasets

import (
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/stub"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	raw, err := stub.Asset("data/datasets.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	w.Write(raw)
	w.WriteHeader(200)
}
