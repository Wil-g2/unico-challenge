package fair

type CreateDTO struct {
	Long         float64 `json:"long" validate:"required"`
	Lat          float64 `json:"lat" validate:"required"`
	Setcens      string  `json:"setcens" validate:"required"`
	Areap        string  `json:"areap" validate:"required"`
	CodDist      int     `json:"cod_dist" validate:"required"`
	District     string  `json:"district" validate:"required"`
	CodSubPref   int     `json:"cod_sub_pref" validate:"required"`
	Region5      string  `json:"region_5" validate:"required"`
	Region8      string  `json:"region_8" validate:"required"`
	FairName     string  `json:"fair_name" validate:"required"`
	Record       string  `json:"record" validate:"required"`
	Street       string  `json:"street" validate:"required"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	Reference    string  `json:"reference"`
}

type UpdateDTO struct {
	Long         float64 `json:"long" validate:"required"`
	Lat          float64 `json:"lat" validate:"required"`
	Setcens      string  `json:"setcens" validate:"required"`
	Areap        string  `json:"areap" validate:"required"`
	CodDist      int     `json:"cod_dist" validate:"required"`
	District     string  `json:"district" validate:"required"`
	CodSubPref   int     `json:"cod_sub_pref" validate:"required"`
	Region5      string  `json:"region_5" validate:"required"`
	Region8      string  `json:"region_8" validate:"required"`
	FairName     string  `json:"fair_name" validate:"required"`
	Record       string  `json:"record" validate:"required"`
	Street       string  `json:"street" validate:"required"`
	Number       string  `json:"number" validate:"required"`
	Neighborhood string  `json:"neighborhood" validate:"required"`
	Reference    string  `json:"reference" validate:"required"`
}
