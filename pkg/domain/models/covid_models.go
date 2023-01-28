package models

type CovidModel struct {
	Year           int    `json:"year,omitempty" bson:"year,omitempty"`
	Weeknum        int    `json:"weeknum,omitempty" bson:"weeknum,omitempty"`
	Gender         string `json:"gender,omitempty" bson:"gender, omitempty"`
	AgeNumber      string `json:"age_number,omitempty" bson:"age_number, omitempty"`
	AgeRange       string `json:"age_range,omitempty" bson:"age_range, omitempty"`
	Job            string `json:"job,omitempty" bson:"job, omitempty"`
	Risk           string `json:"risk,omitempty" bson:"risk, omitempty"`
	PatientType    string `json:"patient_type,omitempty" bson:"patient_type, omitempty"`
	Province       string `json:"province,omitempty" bson:"province, omitempty"`
	ReportingGroup string `json:"reporting_group,omitempty" bson:"reporting_group, omitempty"`
	RegionOdpc     string `json:"region_odpc,omitempty" bson:"region_odpc, omitempty"`
	Region         string `json:"region,omitempty" bson:"region, omitempty"`
	UpdateDate     string `json:"update_date,omitempty" bson:"update_date, omitempty"`
}

type CovidDataResponse struct {
	Data  []CovidModel `json:"data,omitempty" bson:"data, omitempty"`
	Links struct {
		First string      `json:"first,omitempty"`
		Last  string      `json:"last,omitempty"`
		Prev  interface{} `json:"prev,omitempty"`
		Next  string      `json:"next,omitempty"`
	} `json:"links,omitempty"`
	Meta struct {
		CurrentPage int `json:"current_page,omitempty"`
		From        int `json:"from,omitempty"`
		LastPage    int `json:"last_page,omitempty"`
		Links       []struct {
			URL    interface{} `json:"url,omitempty"`
			Label  string      `json:"label,omitempty"`
			Active bool        `json:"active,omitempty"`
		} `json:"links,omitempty"`
		Path    string `json:"path,omitempty"`
		PerPage int    `json:"per_page,omitempty"`
		To      int    `json:"to,omitempty"`
		Total   int    `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}

type CovidByYear struct {
	Year  int `json:"year,omitempty" bson:"year,omitempty"`
	Count int `json:"count,omitempty" bson:"count,omitempty"`
}

type CovidByJob struct {
	Job   string `json:"job,omitempty" bson:"job,omitempty"`
	Count int    `json:"count,omitempty" bson:"count,omitempty"`
}
