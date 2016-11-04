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

// Observation  | Variable          | Sex   | Employment status
// 123          | Gross weekly pay  | All  | Full Time
// 70           | Gross weekly pay  | All  | Part Time

// segmentation - do attributes need to be attached to the observations or variables/ dimensions?
   // need to be


// should variables be a different concept to dimensions - no


// Darren feedback
// variables do not exist as a concept - they are just dimensions
// attributes / units can differ
//  - should be attached to observations
//  - OR possible attached to the dimension in the context of the other dimension values

type Datasets struct {
	Items []Dataset
	Count int
	Total int
	StartIndex int
	ItemsPerPage int
}

type Dataset struct {
	ID string
	Title string
	Metadata Metadata
	Dimensions []Dimension
	Data Table
}

type Metadata struct {
	Description string
	Taxonomies []string
}

type Dimension struct {
	Name string // Sex

	// possible values?

	// possible units of measure?

	Value string // Male
}

type Row struct {
	Observation interface{} // 123
	Dimensions []Dimension // Sex=Male
}

type Table struct {
	Rows []Row
}

