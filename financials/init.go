package financials

type FinancialsAPI struct {
	APIURL string
	APIKey string
}

func New(APIURL string, APIKey string) *FinancialsAPI {
	return &FinancialsAPI{
		APIURL: APIURL,
		APIKey: APIKey,
	}
}
