package fair

type CreateDTO struct {
	Long         float64
	Lat          float64
	Setcens      string
	Areap        string
	CodDist      int
	Distrito     string
	CodSubPref   int
	Regiao5      string
	Regiao8      string
	FairName     string
	Record       string
	Street       string
	Number       string
	Neighborhood string
	Reference    string
}

type UpdateDTO struct {
	Long         float64
	Lat          float64
	Setcens      string
	Areap        string
	CodDist      int
	Distrito     string
	CodSubPref   int
	Regiao5      string
	Regiao8      string
	FairName     string
	Record       string
	Street       string
	Number       string
	Neighborhood string
	Reference    string
}
