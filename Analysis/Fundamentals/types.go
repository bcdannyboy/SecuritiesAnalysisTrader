package Fundamentals

type GrowthBalanceSheetStatementAsReported struct {
	Date                                                  string      `json:"date"`
	Symbol                                                string      `json:"symbol"`
	Period                                                string      `json:"period"`
	GrowthLiabilitiesandstockholdersequity                interface{} `json:"growthliabilitiesandstockholdersequity"`
	GrowthLiabilities                                     interface{} `json:"growthliabilities"`
	GrowthLiabilitiescurrent                              interface{} `json:"growthliabilitiescurrent"`
	GrowthCommonstocksharesauthorized                     float64     `json:"growthcommonstocksharesauthorized"`
	GrowthCashandcashequivalentsatcarryingvalue           interface{} `json:"growthcashandcashequivalentsatcarryingvalue"`
	GrowthRetainedearningsaccumulateddeficit              interface{} `json:"growthretainedearningsaccumulateddeficit"`
	GrowthLiabilitiesnoncurrent                           interface{} `json:"growthliabilitiesnoncurrent"`
	GrowthPropertyplantandequipmentnet                    float64     `json:"growthpropertyplantandequipmentnet"`
	GrowthCommonstocksincludingadditionalpaidincapital    interface{} `json:"growthcommonstocksincludingadditionalpaidincapital"`
	GrowthCommercialpaper                                 float64     `json:"growthcommercialpaper"`
	GrowthLongtermdebtcurrent                             float64     `json:"growthlongtermdebtcurrent"`
	GrowthCommonstocksharesoutstanding                    float64     `json:"growthcommonstocksharesoutstanding"`
	GrowthOtherliabilitiesnoncurrent                      float64     `json:"growthotherliabilitiesnoncurrent"`
	GrowthMarketablesecuritiescurrent                     float64     `json:"growthmarketablesecuritiescurrent"`
	GrowthOtherliabilitiescurrent                         float64     `json:"growthotherliabilitiescurrent"`
	GrowthAssetscurrent                                   interface{} `json:"growthassetscurrent"`
	GrowthLongtermdebtnoncurrent                          float64     `json:"growthlongtermdebtnoncurrent"`
	GrowthContractwithcustomerliabilitycurrent            float64     `json:"growthcontractwithcustomerliabilitycurrent"`
	GrowthNontradereceivablescurrent                      float64     `json:"growthnontradereceivablescurrent"`
	GrowthCommonstocksharesissued                         float64     `json:"growthcommonstocksharesissued"`
	GrowthStockholdersequity                              interface{} `json:"growthstockholdersequity"`
	GrowthAccountsreceivablenetcurrent                    float64     `json:"growthaccountsreceivablenetcurrent"`
	GrowthAccountspayablecurrent                          float64     `json:"growthaccountspayablecurrent"`
	GrowthAssets                                          interface{} `json:"growthassets"`
	GrowthAssetsnoncurrent                                interface{} `json:"growthassetsnoncurrent"`
	GrowthOtherassetscurrent                              float64     `json:"growthotherassetscurrent"`
	GrowthOtherassetsnoncurrent                           float64     `json:"growthotherassetsnoncurrent"`
	GrowthInventorynet                                    float64     `json:"growthinventorynet"`
	GrowthMarketablesecuritiesnoncurrent                  interface{} `json:"growthmarketablesecuritiesnoncurrent"`
	GrowthAccumulatedothercomprehensiveincomelossnetoftax float64     `json:"growthaccumulatedothercomprehensiveincomelossnetoftax"`
	GrowthOthershorttermborrowings                        float64     `json:"growthothershorttermborrowings"`
}
