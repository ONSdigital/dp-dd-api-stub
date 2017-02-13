package models

// Dataset: Earnings by place of work
// Dimensions: Sex | Employment status
// Variables: gross weekly pay, weekly pay excluding overtime

// Table
// Observation  | Variable (Earnings)| Sex   | Employment status
// 123          | Gross weekly pay  | Male  | Full Time
// 70           | Gross weekly pay  | Male  | Part Time
// 119          | Gross weekly pay  | Female| Full Time
// 70           | Gross weekly pay  | Female| Part Time
// 123          | weekly pay excluding overtime  | Male  | Full Time
// 70           | weekly pay excluding overtime  | Male  | Part Time
// 119          | weekly pay excluding overtime  | Female| Full Time
// 70           | weekly pay excluding overtime  | Female| Part Time

// Table -  Observation | DimensionID | DimensionOptionID | DimensionOptionLabel
// Observation  | Geography ID  | Geography Label   | Sex ID | Sex Label    | Age ID    | Age Label | Residence Type ID | Residence Type Label
// 105          | K04000001     | England and Wales | 1      | Females      | 1         | 50+       | 1                 | Lives in a communal establishment

// human readable flag - include labels / replace them?

// The csv output should not rely on fixed column locations. columns should be identified by their header value/ label.
// by doing this we can remove all of the blank

//Denormalisation
// ---------------
// The ID / details / label for the dimension can be held at the dataset level - it should not be repeated throughout the table.

// Fixed dimension options - if the dimension only ever has one value for the dataset, we can define it as a fixed dimension.
// This prevents it being repeated for each observation.

// Create ID's for dimensions and their options
// when multiple datasets have the same dimensions they can be the same ID instead of having various string representations of the same thing.
// e.g. year can mean different things in different datasets - financial year or calendar year

// can we translate dimensions to ID's as part of the data baker step??

type Datasets struct {
	Items        []*Dataset `json:"items,omitempty"`
	Count        int        `json:"count"`
	Total        int        `json:"total"`
	StartIndex   int        `json:"startIndex"`
	ItemsPerPage int        `json:"itemsPerPage"`
	First        string     `json:"first"`
	Next         string     `json:"next,omitempty"`
	Prev         string     `json:"prev,omitempty"`
	Last         string     `json:"last"`
	Page         int        `json:"page"`
	TotalPages   int        `json:"totalPages"`
}

type DataResource struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Metadata *Metadata `json:"metadata"`
	Datasets []string  `json:"datasets"`
}

type Dataset struct {
	ID               string       `json:"id"`
	CustomerFacingID string       `json:"customerFacingId"`
	Title            string       `json:"title"`
	URL              string       `json:"url,omitempty"`
	Metadata         *Metadata    `json:"metadata,omitempty"`
	Dimensions       []*Dimension `json:"dimensions,omitempty"`
	Data             *Table       `json:"data,omitempty"`
}

type GeoArea struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	ExtCode     string        `json:"extCode"`
	Metadata    *Metadata     `json:"metadata,omitempty"`
	RelGeoArea  string        `json:"relid,omitempty"`
	AreaType    string        `json:"areaType"`
	AreaLevel   string        `json:"areaLevel"`
	Populations []*Population `json:"populations,omitempty"`
}

type GeoHierarchy struct {
	Name     string     `json:"name"`
	GeoAreas []*GeoArea `json:"geoAreas,omitempty"`
}

type Population struct {
	GeoAreaID    string `json:"geoAreaID"`
	ExtCode      string `json:"extCode"`
	TimePeriodID string `json:"timePeriodID"`
}

type TimePeriod struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	TimeTypeID string `json:"timeTypeID"`
}

type TimeType struct {
	ID string `json:"id"`
}

type Metadata struct {
	Description        string         `json:"description,omitempty"`
	Taxonomies         []string       `json:"taxonomies,omitempty"`
	Contact            *Contact       `json:"contact"`
	ReleaseDate        string         `json:"releaseDate"`
	NextRelease        string         `json:"nextReleaseDate"`
	NationalStatistics bool           `json:"nationalStatistics"` // whether these are official National Statistics
	Publications       []string       `json:"associatedPublications"`
	Methodology        []*Methodology `json:"methodology"`
	TermsAndConditions string         `json:"termsAndConditions,omitempty"`
}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type Methodology struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Dimension struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"` // eg.: Sex
	Type           string             `json:"type"` // standard, classification, time, geography
	URL            string             `json:"url"`
	Hierarchical   bool               `json:"hierarchical"`
	Options        []*DimensionOption `json:"options,omitempty"`
	SelectedOption *DimensionOption   `json:"selectedOption,omitempty"`
}

type DimensionOption struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"` // Male
	Options []*DimensionOption `json:"options,omitempty"`
}

type Hierarchy struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	Options []*HierarchyEntry `json:"options,omitempty"`
}

type HierarchyEntry struct {
	Code      string              `json:"code"`
	Name      string              `json:"name"`
	LevelType *HierarchyLevelType `json:"levelType,omitempty"`
	HasData   bool                `json:"hasData,omitempty"`	// used only for sparsely populated hierarchy
	Options   []*HierarchyEntry   `json:"options,omitempty"`
}

type HierarchyLevelType struct {
	ID  string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Row struct {
	Observation interface{}        // 123
	Dimensions  []*DimensionOption // Sex=Male
}

type Table struct {
	Rows []*Row
}
