package datasets

import (
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"io/ioutil"
	"net/http"
	"os"
	//"strconv"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	raw, err := ioutil.ReadFile("./data/datasets.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var stubData models.Datasets = models.Datasets{}
	json.Unmarshal(raw, &stubData)

	//baseUrl := "http://" + req.Host + req.URL.EscapedPath()
	//stubData.First = baseUrl + "?page=1"
	//stubData.Last = baseUrl + "?page=" + strconv.Itoa(stubData.TotalPages)

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(stubData)

	w.WriteHeader(200)
}
