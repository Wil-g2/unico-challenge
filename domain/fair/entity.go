package fair

type Fair struct {
	ID           int     `json:"id"`
	Long         float64 `json:"long"`
	Lat          float64 `json:"lat"`
	Setcens      string  `json:"setcens"`
	Areap        string  `json:"areap"`
	CodDist      int     `json:"cod_dist"`
	District     string  `json:"district"`
	CodSubPref   int     `json:"cod_sub_pref"`
	Region5      string  `json:"region_5"`
	Region8      string  `json:"region_8"`
	FairName     string  `json:"fair_name"`
	Record       string  `json:"record"`
	Street       string  `json:"street"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	Reference    string  `json:"reference"`
}

func NewFair(Long float64,
	Lat float64,
	Setcens string,
	Areap string,
	CodDist int,
	District string,
	CodSubPref int,
	Region5 string,
	Region8 string,
	FairName string,
	Record string,
	Street string,
	Number string,
	Neighborhood string,
	Reference string) *Fair {
	return &Fair{
		Long:         Long,
		Lat:          Lat,
		Setcens:      Setcens,
		Areap:        Areap,
		CodDist:      CodDist,
		District:     District,
		CodSubPref:   CodSubPref,
		Region5:      Region5,
		Region8:      Region8,
		FairName:     FairName,
		Record:       Record,
		Street:       Street,
		Number:       Number,
		Neighborhood: Neighborhood,
		Reference:    Reference,
	}
}
