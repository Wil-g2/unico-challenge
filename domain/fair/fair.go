package fair

// ID,LONG,LAT,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREF,REGIAO5,REGIAO8,NOME_FEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA,,,,
// 1,-46548146,-23568390,355030885000019,3550308005040,87,VILA FORMOSA,26,ARICANDUVA,Leste,Leste 1,PRA�A LE+O X,7216-8,RUA CODAJ-S,45,VILA FORMOSA,PRA�A  MARECHAL LEIT+O BANDEIRA,,,,
type Fair struct {
	ID           int     `json:"id"`
	Long         float64 `json:"long"`
	Lat          float64 `json:"lat"`
	Setcens      string  `json:"setcens"`
	Areap        string  `json:"areap"`
	CodDist      int     `json:"cod_dist"`
	Distrito     string  `json:"distrito"`
	CodSubPref   int     `json:"cod_sub_pref"`
	Regiao5      string  `json:"regiao5"`
	Regiao8      string  `json:"regiao8"`
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
	Distrito string,
	CodSubPref int,
	Regiao5 string,
	Regiao8 string,
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
		Distrito:     Distrito,
		CodSubPref:   CodSubPref,
		Regiao5:      Regiao5,
		Regiao8:      Regiao8,
		FairName:     FairName,
		Record:       Record,
		Street:       Street,
		Number:       Number,
		Neighborhood: Neighborhood,
		Reference:    Reference,
	}
}
