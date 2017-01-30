package hierarchy

import (
	"encoding/json"
	"github.com/ONSdigital/dp-dd-api-stub/models"
	"net/http"
)

func createTimeMonthEntries(name string) *models.HierarchyEntry {
	var levelType *models.HierarchyLevelType = &models.HierarchyLevelType{
		Code:  "month",
		Name:  "month",
		Level: 1,
	}

	return &models.HierarchyEntry{
		ID:   name,
		Name: name,
		LevelType: &models.HierarchyLevelType{
			Code:  "year",
			Name:  "year",
			Level: 0,
		},
		Options: []*models.HierarchyEntry{
			{ID: name + ".01", Name: "January", LevelType: levelType},
			{ID: name + ".02", Name: "February", LevelType: levelType},
			{ID: name + ".03", Name: "March", LevelType: levelType},
			{ID: name + ".04", Name: "April", LevelType: levelType},
			{ID: name + ".05", Name: "May", LevelType: levelType},
			{ID: name + ".06", Name: "June", LevelType: levelType},
			{ID: name + ".07", Name: "July", LevelType: levelType},
			{ID: name + ".08", Name: "August", LevelType: levelType},
			{ID: name + ".09", Name: "September", LevelType: levelType},
			{ID: name + ".10", Name: "October", LevelType: levelType},
			{ID: name + ".11", Name: "November", LevelType: levelType},
			{ID: name + ".12", Name: "December", LevelType: levelType},
		},
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {

	var hierarchiesStub map[string]*models.Hierarchy = make(map[string]*models.Hierarchy)
	hierarchiesStub["Time"] = &models.Hierarchy{
		ID:   "Time",
		Name: "Time",
		Type: "time",
		Options: []*models.HierarchyEntry{
			createTimeMonthEntries("2015"),
			createTimeMonthEntries("2016"),
			createTimeMonthEntries("2017"),
		},
	}

	hierarchiesStub["CI_000641"] = &models.Hierarchy{
		ID:   "CI_000641",
		Name: "Special Aggregate",
		Type: "classification",
		Options: []*models.HierarchyEntry{
			{ID: "CI_0004216", Name: "CPI (overall index)"},
			{ID: "CI_0004331", Name: "01 Food and non-alcoholic beverages", Options: []*models.HierarchyEntry{
				{ID: "CI_0004217", Name: "01.1 Food", Options: []*models.HierarchyEntry{
					{ID: "CI_0004252", Name: "01.1.1 Bread and cereals"},
					{ID: "CI_0004253", Name: "01.1.2 Meat"},
					{ID: "CI_0004254", Name: "01.1.3 Fish"},
					{ID: "CI_0004255", Name: "01.1.4 Milk"},
					{ID: "CI_0004256", Name: "01.1.5 Oils and fats"},
					{ID: "CI_0004257", Name: "01.1.6 Fruit"},
					{ID: "CI_0004258", Name: "01.1.7 Vegetables including potatoes and tubers"},
					{ID: "CI_0004259", Name: "01.1.8 Sugar, jam, syrups, chocolate and confectionery"},
					{ID: "CI_0004260", Name: "01.1.9 Food products (nec)"},
				}},
				{ID: "CI_0004218", Name: "01.2 Non-alcoholic beverages", Options: []*models.HierarchyEntry{
					{ID: "CI_0004252", Name: "01.2.1 Coffee, tea and cocoa"},
					{ID: "CI_0004253", Name: "01.2.2 Mineral waters, soft drinks and juices"},
				}},
			}},
			{ID: "CI_0004332", Name: "02 Alcoholic beverages and tobacco", Options: []*models.HierarchyEntry{
				{ID: "CI_0004219", Name: "02.1 Alcoholic beverages", Options: []*models.HierarchyEntry{
					{ID: "CI_0004263", Name: "02.1.1 Spirits"},
					{ID: "CI_0004264", Name: "02.1.2 Wine "},
					{ID: "CI_0004265", Name: "02.1.3 Beer"},
				}},
				{ID: "CI_0004220", Name: "02.2 Tobacco", Options: []*models.HierarchyEntry{
					{ID: "CI_0004266", Name: "02.2.0 Tobacco"},
				}},
			}},
		},
	}

	hierarchyId := req.URL.Query().Get(":hierarchyId")
	if _, ok := hierarchiesStub[hierarchyId]; !ok {
		w.WriteHeader(404)
		return
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(hierarchiesStub[hierarchyId])

	w.WriteHeader(200)
}
