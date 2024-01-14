package financials

type FinancialsAPI struct {
	APIURL string
	APIKey string
}

func New(APIKey string, APIURL string) *FinancialsAPI {
	return &FinancialsAPI{
		APIURL: APIURL,
		APIKey: APIKey,
	}
}
