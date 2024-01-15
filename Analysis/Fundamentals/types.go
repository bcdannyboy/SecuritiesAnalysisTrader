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

type DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported struct {
	Date                                       string  `json:"dateMatch"`
	Symbol                                     string  `json:"symbolMatch"`
	Period                                     string  `json:"periodMatch"`
	CashAndCashEquivalentsMatch                float64 `json:"cashAndCashEquivalentsMatch"`
	PropertyPlantEquipmentNetMatch             float64 `json:"propertyPlantEquipmentNetMatch"`
	CommonStockMatch                           float64 `json:"commonStockMatch"`
	RetainedEarningsMatch                      float64 `json:"retainedEarningsMatch"`
	TotalCurrentAssetsMatch                    float64 `json:"totalCurrentAssetsMatch"`
	TotalAssetsMatch                           float64 `json:"totalAssetsMatch"`
	TotalCurrentLiabilitiesMatch               float64 `json:"totalCurrentLiabilitiesMatch"`
	TotalLiabilitiesMatch                      float64 `json:"totalLiabilitiesMatch"`
	TotalStockholdersEquityMatch               float64 `json:"totalStockholdersEquityMatch"`
	TotalLiabilitiesAndStockholdersEquityMatch float64 `json:"totalLiabilitiesAndStockholdersEquityMatch"`
}

type CashFlowStatementAsReportedGrowth struct {
	Date                                                                                                                 string      `json:"date"`
	Symbol                                                                                                               string      `json:"symbol"`
	Period                                                                                                               string      `json:"period"`
	PaymentsforrepurchaseofcommonstockGrowth                                                                             float64     `json:"paymentsforrepurchaseofcommonstockgrowth"`
	SharebasedcompensationGrowth                                                                                         float64     `json:"sharebasedcompensationgrowth"`
	NetincomelossGrowth                                                                                                  float64     `json:"netincomelossgrowth"`
	IncreasedecreaseinaccountspayableGrowth                                                                              float64     `json:"increasedecreaseinaccountspayablegrowth"`
	ProceedsfrompaymentsforotherfinancingactivitiesGrowth                                                                float64     `json:"proceedsfrompaymentsforotherfinancingactivitiesgrowth"`
	PaymentsrelatedtotaxwithholdingforsharebasedcompensationGrowth                                                       float64     `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensationgrowth"`
	IncreasedecreaseinotheroperatingliabilitiesGrowth                                                                    float64     `json:"increasedecreaseinotheroperatingliabilitiesgrowth"`
	OthernoncashincomeexpenseGrowth                                                                                      float64     `json:"othernoncashincomeexpensegrowth"`
	PaymentstoacquirebusinessesnetofcashacquiredGrowth                                                                   float64     `json:"paymentstoacquirebusinessesnetofcashacquiredgrowth"`
	DeferredincometaxexpensebenefitGrowth                                                                                float64     `json:"deferredincometaxexpensebenefitgrowth"`
	CashcashequivalentsrestrictedcashandrestrictedcashequivalentsGrowth                                                  interface{} `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsgrowth"`
	CashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffectGrowth float64     `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffectgrowth"`
	NetcashprovidedbyusedinoperatingactivitiesGrowth                                                                     float64     `json:"netcashprovidedbyusedinoperatingactivitiesgrowth"`
	ProceedsfromsaleofavailableforsalesecuritiesdebtGrowth                                                               float64     `json:"proceedsfromsaleofavailableforsalesecuritiesdebtgrowth"`
	RepaymentsoflongtermdebtGrowth                                                                                       float64     `json:"repaymentsoflongtermdebtgrowth"`
	IncometaxespaidnetGrowth                                                                                             float64     `json:"incometaxespaidnetgrowth"`
	ProceedsfromissuanceoflongtermdebtGrowth                                                                             float64     `json:"proceedsfromissuanceoflongtermdebtgrowth"`
	PaymentstoacquireotherinvestmentsGrowth                                                                              float64     `json:"paymentstoacquireotherinvestmentsgrowth"`
	NetcashprovidedbyusedininvestingactivitiesGrowth                                                                     interface{} `json:"netcashprovidedbyusedininvestingactivitiesgrowth"`
	IncreasedecreaseincontractwithcustomerliabilityGrowth                                                                float64     `json:"increasedecreaseincontractwithcustomerliabilitygrowth"`
	InterestpaidnetGrowth                                                                                                float64     `json:"interestpaidnetgrowth"`
	NetcashprovidedbyusedinfinancingactivitiesGrowth                                                                     interface{} `json:"netcashprovidedbyusedinfinancingactivitiesgrowth"`
	ProceedsfromrepaymentsofcommercialpaperGrowth                                                                        float64     `json:"proceedsfromrepaymentsofcommercialpapergrowth"`
	ProceedsfromsaleandmaturityofotherinvestmentsGrowth                                                                  float64     `json:"proceedsfromsaleandmaturityofotherinvestmentsgrowth"`
	PaymentstoacquireavailableforsalesecuritiesdebtGrowth                                                                interface{} `json:"paymentstoacquireavailableforsalesecuritiesdebtgrowth"`
	PaymentstoacquirepropertyplantandequipmentGrowth                                                                     float64     `json:"paymentstoacquirepropertyplantandequipmentgrowth"`
	PaymentsforproceedsfromotherinvestingactivitiesGrowth                                                                float64     `json:"paymentsforproceedsfromotherinvestingactivitiesgrowth"`
	IncreasedecreaseinotherreceivablesGrowth                                                                             float64     `json:"increasedecreaseinotherreceivablesgrowth"`
	PaymentsofdividendsGrowth                                                                                            float64     `json:"paymentsofdividendsgrowth"`
	IncreasedecreaseininventoriesGrowth                                                                                  float64     `json:"increasedecreaseininventoriesgrowth"`
	IncreasedecreaseinaccountsreceivableGrowth                                                                           float64     `json:"increasedecreaseinaccountsreceivablegrowth"`
	ProceedsfromissuanceofcommonstockGrowth                                                                              float64     `json:"proceedsfromissuanceofcommonstockgrowth"`
	DepreciationdepletionandamortizationGrowth                                                                           float64     `json:"depreciationdepletionandamortizationgrowth"`
	ProceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecuritiesGrowth                                          float64     `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecuritiesgrowth"`
	IncreasedecreaseinotheroperatingassetsGrowth                                                                         float64     `json:"increasedecreaseinotheroperatingassetsgrowth"`
	ProceedsfromothershorttermdebtGrowth                                                                                 float64     `json:"proceedsfromothershorttermdebtgrowth"`
}

type DiscrepancyCashFlowStatementAndCashFlowStatementAsReported struct {
	Date                                                          string  `json:"date"`
	Symbol                                                        string  `json:"symbol"`
	Period                                                        string  `json:"period"`
	FillingDate                                                   string  `json:"fillingDate"`
	AcceptedDate                                                  string  `json:"acceptedDate"`
	NetIncomeDiscrepancyPercentage                                float64 `json:"netIncomeDiscrepancyPercentage"`
	DepreciationAndAmortizationDiscrepancyPercentage              float64 `json:"depreciationAndAmortizationDiscrepancyPercentage"`
	DeferredIncomeTaxDiscrepancyPercentage                        float64 `json:"deferredIncomeTaxDiscrepancyPercentage"`
	StockBasedCompensationDiscrepancyPercentage                   float64 `json:"stockBasedCompensationDiscrepancyPercentage"`
	ChangeInWorkingCapitalDiscrepancyPercentage                   float64 `json:"changeInWorkingCapitalDiscrepancyPercentage"`
	AccountsReceivablesDiscrepancyPercentage                      float64 `json:"accountsReceivablesDiscrepancyPercentage"`
	InventoryDiscrepancyPercentage                                float64 `json:"inventoryDiscrepancyPercentage"`
	AccountsPayablesDiscrepancyPercentage                         float64 `json:"accountsPayablesDiscrepancyPercentage"`
	OtherNonCashItemsDiscrepancyPercentage                        float64 `json:"otherNonCashItemsDiscrepancyPercentage"`
	NetCashProvidedByOperatingActivitiesDiscrepancyPercentage     float64 `json:"netCashProvidedByOperatingActivitiesDiscrepancyPercentage"`
	InvestmentsInPropertyPlantAndEquipmentDiscrepancyPercentage   float64 `json:"investmentsInPropertyPlantAndEquipmentDiscrepancyPercentage"`
	AcquisitionsNetDiscrepancyPercentage                          float64 `json:"acquisitionsNetDiscrepancyPercentage"`
	PurchasesOfInvestmentsDiscrepancyPercentage                   float64 `json:"purchasesOfInvestmentsDiscrepancyPercentage"`
	SalesMaturitiesOfInvestmentsDiscrepancyPercentage             float64 `json:"salesMaturitiesOfInvestmentsDiscrepancyPercentage"`
	OtherInvestingActivitesDiscrepancyPercentage                  float64 `json:"otherInvestingActivitesDiscrepancyPercentage"`
	NetCashUsedForInvestingActivitesDiscrepancyPercentage         float64 `json:"netCashUsedForInvestingActivitesDiscrepancyPercentage"`
	DebtRepaymentDiscrepancyPercentage                            float64 `json:"debtRepaymentDiscrepancyPercentage"`
	CommonStockIssuedDiscrepancyPercentage                        float64 `json:"commonStockIssuedDiscrepancyPercentage"`
	CommonStockRepurchasedDiscrepancyPercentage                   float64 `json:"commonStockRepurchasedDiscrepancyPercentage"`
	DividendsPaidDiscrepancyPercentage                            float64 `json:"dividendsPaidDiscrepancyPercentage"`
	OtherFinancingActivitesDiscrepancyPercentage                  float64 `json:"otherFinancingActivitesDiscrepancyPercentage"`
	NetCashUsedProvidedByFinancingActivitiesDiscrepancyPercentage float64 `json:"netCashUsedProvidedByFinancingActivitiesDiscrepancyPercentage"`
	EffectOfForexChangesOnCashDiscrepancyPercentage               float64 `json:"effectOfForexChangesOnCashDiscrepancyPercentage"`
	NetChangeInCashDiscrepancyPercentage                          float64 `json:"netChangeInCashDiscrepancyPercentage"`
	CashAtEndOfPeriodDiscrepancyPercentage                        float64 `json:"cashAtEndOfPeriodDiscrepancyPercentage"`
	CashAtBeginningOfPeriodDiscrepancyPercentage                  float64 `json:"cashAtBeginningOfPeriodDiscrepancyPercentage"`
	OperatingCashFlowDiscrepancyPercentage                        float64 `json:"operatingCashFlowDiscrepancyPercentage"`
	CapitalExpenditureDiscrepancyPercentage                       float64 `json:"capitalExpenditureDiscrepancyPercentage"`
	FreeCashFlowDiscrepancyPercentage                             float64 `json:"freeCashFlowDiscrepancyPercentage"`
}
