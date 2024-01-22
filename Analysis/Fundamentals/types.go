package Fundamentals

// Balance Sheet
type GrowthBalanceSheetStatementAsReported struct {
	Date                                                  string      `json:"date,omitempty"`
	Symbol                                                string      `json:"symbol,omitempty"`
	Period                                                string      `json:"period,omitempty"`
	GrowthLiabilitiesandstockholdersequity                interface{} `json:"growthliabilitiesandstockholdersequity,omitempty"`
	GrowthLiabilities                                     interface{} `json:"growthliabilities,omitempty"`
	GrowthLiabilitiescurrent                              interface{} `json:"growthliabilitiescurrent,omitempty"`
	GrowthCommonstocksharesauthorized                     float64     `json:"growthcommonstocksharesauthorized,omitempty"`
	GrowthCashandcashequivalentsatcarryingvalue           interface{} `json:"growthcashandcashequivalentsatcarryingvalue,omitempty"`
	GrowthRetainedearningsaccumulateddeficit              interface{} `json:"growthretainedearningsaccumulateddeficit,omitempty"`
	GrowthLiabilitiesnoncurrent                           interface{} `json:"growthliabilitiesnoncurrent,omitempty"`
	GrowthPropertyplantandequipmentnet                    float64     `json:"growthpropertyplantandequipmentnet,omitempty"`
	GrowthCommonstocksincludingadditionalpaidincapital    interface{} `json:"growthcommonstocksincludingadditionalpaidincapital,omitempty"`
	GrowthCommercialpaper                                 float64     `json:"growthcommercialpaper,omitempty"`
	GrowthLongtermdebtcurrent                             float64     `json:"growthlongtermdebtcurrent,omitempty"`
	GrowthCommonstocksharesoutstanding                    float64     `json:"growthcommonstocksharesoutstanding,omitempty"`
	GrowthOtherliabilitiesnoncurrent                      float64     `json:"growthotherliabilitiesnoncurrent,omitempty"`
	GrowthMarketablesecuritiescurrent                     float64     `json:"growthmarketablesecuritiescurrent,omitempty"`
	GrowthOtherliabilitiescurrent                         float64     `json:"growthotherliabilitiescurrent,omitempty"`
	GrowthAssetscurrent                                   interface{} `json:"growthassetscurrent,omitempty"`
	GrowthLongtermdebtnoncurrent                          float64     `json:"growthlongtermdebtnoncurrent,omitempty"`
	GrowthContractwithcustomerliabilitycurrent            float64     `json:"growthcontractwithcustomerliabilitycurrent,omitempty"`
	GrowthNontradereceivablescurrent                      float64     `json:"growthnontradereceivablescurrent,omitempty"`
	GrowthCommonstocksharesissued                         float64     `json:"growthcommonstocksharesissued,omitempty"`
	GrowthStockholdersequity                              interface{} `json:"growthstockholdersequity,omitempty"`
	GrowthAccountsreceivablenetcurrent                    float64     `json:"growthaccountsreceivablenetcurrent,omitempty"`
	GrowthAccountspayablecurrent                          float64     `json:"growthaccountspayablecurrent,omitempty"`
	GrowthAssets                                          interface{} `json:"growthassets,omitempty"`
	GrowthAssetsnoncurrent                                interface{} `json:"growthassetsnoncurrent,omitempty"`
	GrowthOtherassetscurrent                              float64     `json:"growthotherassetscurrent,omitempty"`
	GrowthOtherassetsnoncurrent                           float64     `json:"growthotherassetsnoncurrent,omitempty"`
	GrowthInventorynet                                    float64     `json:"growthinventorynet,omitempty"`
	GrowthMarketablesecuritiesnoncurrent                  interface{} `json:"growthmarketablesecuritiesnoncurrent,omitempty"`
	GrowthAccumulatedothercomprehensiveincomelossnetoftax float64     `json:"growthaccumulatedothercomprehensiveincomelossnetoftax,omitempty"`
	GrowthOthershorttermborrowings                        float64     `json:"growthothershorttermborrowings,omitempty"`
}

type DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported struct {
	Date                                       string  `json:"dateMatch,omitempty"`
	Symbol                                     string  `json:"symbolMatch,omitempty"`
	Period                                     string  `json:"periodMatch,omitempty"`
	CashAndCashEquivalentsMatch                float64 `json:"cashAndCashEquivalentsMatch,omitempty"`
	PropertyPlantEquipmentNetMatch             float64 `json:"propertyPlantEquipmentNetMatch,omitempty"`
	CommonStockMatch                           float64 `json:"commonStockMatch,omitempty"`
	RetainedEarningsMatch                      float64 `json:"retainedEarningsMatch,omitempty"`
	TotalCurrentAssetsMatch                    float64 `json:"totalCurrentAssetsMatch,omitempty"`
	TotalAssetsMatch                           float64 `json:"totalAssetsMatch,omitempty"`
	TotalCurrentLiabilitiesMatch               float64 `json:"totalCurrentLiabilitiesMatch,omitempty"`
	TotalLiabilitiesMatch                      float64 `json:"totalLiabilitiesMatch,omitempty"`
	TotalStockholdersEquityMatch               float64 `json:"totalStockholdersEquityMatch,omitempty"`
	TotalLiabilitiesAndStockholdersEquityMatch float64 `json:"totalLiabilitiesAndStockholdersEquityMatch,omitempty"`
}

// CashFlow Statement
type CashFlowStatementAsReportedGrowth struct {
	Date                                                                                                                 string      `json:"date,omitempty"`
	Symbol                                                                                                               string      `json:"symbol,omitempty"`
	Period                                                                                                               string      `json:"period,omitempty"`
	PaymentsforrepurchaseofcommonstockGrowth                                                                             float64     `json:"paymentsforrepurchaseofcommonstockgrowth,omitempty"`
	SharebasedcompensationGrowth                                                                                         float64     `json:"sharebasedcompensationgrowth,omitempty"`
	NetincomelossGrowth                                                                                                  float64     `json:"netincomelossgrowth,omitempty"`
	IncreasedecreaseinaccountspayableGrowth                                                                              float64     `json:"increasedecreaseinaccountspayablegrowth,omitempty"`
	ProceedsfrompaymentsforotherfinancingactivitiesGrowth                                                                float64     `json:"proceedsfrompaymentsforotherfinancingactivitiesgrowth,omitempty"`
	PaymentsrelatedtotaxwithholdingforsharebasedcompensationGrowth                                                       float64     `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensationgrowth,omitempty"`
	IncreasedecreaseinotheroperatingliabilitiesGrowth                                                                    float64     `json:"increasedecreaseinotheroperatingliabilitiesgrowth,omitempty"`
	OthernoncashincomeexpenseGrowth                                                                                      float64     `json:"othernoncashincomeexpensegrowth,omitempty"`
	PaymentstoacquirebusinessesnetofcashacquiredGrowth                                                                   float64     `json:"paymentstoacquirebusinessesnetofcashacquiredgrowth,omitempty"`
	DeferredincometaxexpensebenefitGrowth                                                                                float64     `json:"deferredincometaxexpensebenefitgrowth,omitempty"`
	CashcashequivalentsrestrictedcashandrestrictedcashequivalentsGrowth                                                  interface{} `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsgrowth,omitempty"`
	CashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffectGrowth float64     `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffectgrowth,omitempty"`
	NetcashprovidedbyusedinoperatingactivitiesGrowth                                                                     float64     `json:"netcashprovidedbyusedinoperatingactivitiesgrowth,omitempty"`
	ProceedsfromsaleofavailableforsalesecuritiesdebtGrowth                                                               float64     `json:"proceedsfromsaleofavailableforsalesecuritiesdebtgrowth,omitempty"`
	RepaymentsoflongtermdebtGrowth                                                                                       float64     `json:"repaymentsoflongtermdebtgrowth,omitempty"`
	IncometaxespaidnetGrowth                                                                                             float64     `json:"incometaxespaidnetgrowth,omitempty"`
	ProceedsfromissuanceoflongtermdebtGrowth                                                                             float64     `json:"proceedsfromissuanceoflongtermdebtgrowth,omitempty"`
	PaymentstoacquireotherinvestmentsGrowth                                                                              float64     `json:"paymentstoacquireotherinvestmentsgrowth,omitempty"`
	NetcashprovidedbyusedininvestingactivitiesGrowth                                                                     interface{} `json:"netcashprovidedbyusedininvestingactivitiesgrowth,omitempty"`
	IncreasedecreaseincontractwithcustomerliabilityGrowth                                                                float64     `json:"increasedecreaseincontractwithcustomerliabilitygrowth,omitempty"`
	InterestpaidnetGrowth                                                                                                float64     `json:"interestpaidnetgrowth,omitempty"`
	NetcashprovidedbyusedinfinancingactivitiesGrowth                                                                     interface{} `json:"netcashprovidedbyusedinfinancingactivitiesgrowth,omitempty"`
	ProceedsfromrepaymentsofcommercialpaperGrowth                                                                        float64     `json:"proceedsfromrepaymentsofcommercialpapergrowth,omitempty"`
	ProceedsfromsaleandmaturityofotherinvestmentsGrowth                                                                  float64     `json:"proceedsfromsaleandmaturityofotherinvestmentsgrowth,omitempty"`
	PaymentstoacquireavailableforsalesecuritiesdebtGrowth                                                                interface{} `json:"paymentstoacquireavailableforsalesecuritiesdebtgrowth,omitempty"`
	PaymentstoacquirepropertyplantandequipmentGrowth                                                                     float64     `json:"paymentstoacquirepropertyplantandequipmentgrowth,omitempty"`
	PaymentsforproceedsfromotherinvestingactivitiesGrowth                                                                float64     `json:"paymentsforproceedsfromotherinvestingactivitiesgrowth,omitempty"`
	IncreasedecreaseinotherreceivablesGrowth                                                                             float64     `json:"increasedecreaseinotherreceivablesgrowth,omitempty"`
	PaymentsofdividendsGrowth                                                                                            float64     `json:"paymentsofdividendsgrowth,omitempty"`
	IncreasedecreaseininventoriesGrowth                                                                                  float64     `json:"increasedecreaseininventoriesgrowth,omitempty"`
	IncreasedecreaseinaccountsreceivableGrowth                                                                           float64     `json:"increasedecreaseinaccountsreceivablegrowth,omitempty"`
	ProceedsfromissuanceofcommonstockGrowth                                                                              float64     `json:"proceedsfromissuanceofcommonstockgrowth,omitempty"`
	DepreciationdepletionandamortizationGrowth                                                                           float64     `json:"depreciationdepletionandamortizationgrowth,omitempty"`
	ProceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecuritiesGrowth                                          float64     `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecuritiesgrowth,omitempty"`
	IncreasedecreaseinotheroperatingassetsGrowth                                                                         float64     `json:"increasedecreaseinotheroperatingassetsgrowth,omitempty"`
	ProceedsfromothershorttermdebtGrowth                                                                                 float64     `json:"proceedsfromothershorttermdebtgrowth,omitempty"`
}

type DiscrepancyCashFlowStatementAndCashFlowStatementAsReported struct {
	Date                                                          string  `json:"date,omitempty"`
	Symbol                                                        string  `json:"symbol,omitempty"`
	Period                                                        string  `json:"period,omitempty"`
	FillingDate                                                   string  `json:"fillingDate,omitempty"`
	AcceptedDate                                                  string  `json:"acceptedDate,omitempty"`
	NetIncomeDiscrepancyPercentage                                float64 `json:"netIncomeDiscrepancyPercentage,omitempty"`
	DepreciationAndAmortizationDiscrepancyPercentage              float64 `json:"depreciationAndAmortizationDiscrepancyPercentage,omitempty"`
	DeferredIncomeTaxDiscrepancyPercentage                        float64 `json:"deferredIncomeTaxDiscrepancyPercentage,omitempty"`
	StockBasedCompensationDiscrepancyPercentage                   float64 `json:"stockBasedCompensationDiscrepancyPercentage,omitempty"`
	ChangeInWorkingCapitalDiscrepancyPercentage                   float64 `json:"changeInWorkingCapitalDiscrepancyPercentage,omitempty"`
	AccountsReceivablesDiscrepancyPercentage                      float64 `json:"accountsReceivablesDiscrepancyPercentage,omitempty"`
	InventoryDiscrepancyPercentage                                float64 `json:"inventoryDiscrepancyPercentage,omitempty"`
	AccountsPayablesDiscrepancyPercentage                         float64 `json:"accountsPayablesDiscrepancyPercentage,omitempty"`
	OtherNonCashItemsDiscrepancyPercentage                        float64 `json:"otherNonCashItemsDiscrepancyPercentage,omitempty"`
	NetCashProvidedByOperatingActivitiesDiscrepancyPercentage     float64 `json:"netCashProvidedByOperatingActivitiesDiscrepancyPercentage,omitempty"`
	InvestmentsInPropertyPlantAndEquipmentDiscrepancyPercentage   float64 `json:"investmentsInPropertyPlantAndEquipmentDiscrepancyPercentage,omitempty"`
	AcquisitionsNetDiscrepancyPercentage                          float64 `json:"acquisitionsNetDiscrepancyPercentage,omitempty"`
	PurchasesOfInvestmentsDiscrepancyPercentage                   float64 `json:"purchasesOfInvestmentsDiscrepancyPercentage,omitempty"`
	SalesMaturitiesOfInvestmentsDiscrepancyPercentage             float64 `json:"salesMaturitiesOfInvestmentsDiscrepancyPercentage,omitempty"`
	OtherInvestingActivitesDiscrepancyPercentage                  float64 `json:"otherInvestingActivitesDiscrepancyPercentage,omitempty"`
	NetCashUsedForInvestingActivitesDiscrepancyPercentage         float64 `json:"netCashUsedForInvestingActivitesDiscrepancyPercentage,omitempty"`
	DebtRepaymentDiscrepancyPercentage                            float64 `json:"debtRepaymentDiscrepancyPercentage,omitempty"`
	CommonStockIssuedDiscrepancyPercentage                        float64 `json:"commonStockIssuedDiscrepancyPercentage,omitempty"`
	CommonStockRepurchasedDiscrepancyPercentage                   float64 `json:"commonStockRepurchasedDiscrepancyPercentage,omitempty"`
	DividendsPaidDiscrepancyPercentage                            float64 `json:"dividendsPaidDiscrepancyPercentage,omitempty"`
	OtherFinancingActivitesDiscrepancyPercentage                  float64 `json:"otherFinancingActivitesDiscrepancyPercentage,omitempty"`
	NetCashUsedProvidedByFinancingActivitiesDiscrepancyPercentage float64 `json:"netCashUsedProvidedByFinancingActivitiesDiscrepancyPercentage,omitempty"`
	EffectOfForexChangesOnCashDiscrepancyPercentage               float64 `json:"effectOfForexChangesOnCashDiscrepancyPercentage,omitempty"`
	NetChangeInCashDiscrepancyPercentage                          float64 `json:"netChangeInCashDiscrepancyPercentage,omitempty"`
	CashAtEndOfPeriodDiscrepancyPercentage                        float64 `json:"cashAtEndOfPeriodDiscrepancyPercentage,omitempty"`
	CashAtBeginningOfPeriodDiscrepancyPercentage                  float64 `json:"cashAtBeginningOfPeriodDiscrepancyPercentage,omitempty"`
	OperatingCashFlowDiscrepancyPercentage                        float64 `json:"operatingCashFlowDiscrepancyPercentage,omitempty"`
	CapitalExpenditureDiscrepancyPercentage                       float64 `json:"capitalExpenditureDiscrepancyPercentage,omitempty"`
	FreeCashFlowDiscrepancyPercentage                             float64 `json:"freeCashFlowDiscrepancyPercentage,omitempty"`
}

// Income Statement
type DiscrepancyIncomeStatementAndIncomeStatementAsReported struct {
	Date                                      string  `json:"date,omitempty"`
	Symbol                                    string  `json:"symbol,omitempty"`
	Period                                    string  `json:"period,omitempty"`
	NetIncomeDiscrepancy                      float64 `json:"netIncomeDiscrepancy,omitempty"`
	GrossProfitDiscrepancy                    float64 `json:"grossProfitDiscrepancy,omitempty"`
	ResearchAndDevelopmentExpensesDiscrepancy float64 `json:"researchAndDevelopmentExpensesDiscrepancy,omitempty"`
	OperatingIncomeDiscrepancy                float64 `json:"operatingIncomeDiscrepancy,omitempty"`
	EpsDiscrepancy                            float64 `json:"epsDiscrepancy,omitempty"`
	EpsDilutedDiscrepancy                     float64 `json:"epsDilutedDiscrepancy,omitempty"`
	WeightedAverageShsOutDiscrepancy          float64 `json:"weightedAverageShsOutDiscrepancy,omitempty"`
	WeightedAverageShsOutDilDiscrepancy       float64 `json:"weightedAverageShsOutDilDiscrepancy,omitempty"`
	IncomeTaxExpenseDiscrepancy               float64 `json:"incomeTaxExpenseDiscrepancy,omitempty"`
}

type GrowthIncomeStatementAsReported struct {
	Date                                                                                              string      `json:"date,omitempty"`
	Symbol                                                                                            string      `json:"symbol,omitempty"`
	Period                                                                                            string      `json:"period,omitempty"`
	CostofgoodsandservicessoldGrowth                                                                  interface{} `json:"costofgoodsandservicessoldgrowth,omitempty"`
	NetincomelossGrowth                                                                               float64     `json:"netincomelossgrowth,omitempty"`
	ResearchanddevelopmentexpenseGrowth                                                               float64     `json:"researchanddevelopmentexpensegrowth,omitempty"`
	GrossprofitGrowth                                                                                 interface{} `json:"grossprofitgrowth,omitempty"`
	OthercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftaxGrowth   float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftaxgrowth,omitempty"`
	OthercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftaxGrowth                    float64     `json:"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftaxgrowth,omitempty"`
	OthercomprehensiveincomelossderivativesqualifyingashedgesnetoftaxGrowth                           float64     `json:"othercomprehensiveincomelossderivativesqualifyingashedgesnetoftaxgrowth,omitempty"`
	OthercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftaxGrowth      float64     `json:"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftaxgrowth,omitempty"`
	OthercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertaxGrowth      float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertaxgrowth,omitempty"`
	WeightedaveragenumberofdilutedsharesoutstandingGrowth                                             float64     `json:"weightedaveragenumberofdilutedsharesoutstandinggrowth,omitempty"`
	WeightedaveragenumberofsharesoutstandingbasicGrowth                                               float64     `json:"weightedaveragenumberofsharesoutstandingbasicgrowth,omitempty"`
	OthercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftaxGrowth          float64     `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftaxgrowth,omitempty"`
	OperatingincomelossGrowth                                                                         float64     `json:"operatingincomelossgrowth,omitempty"`
	OthercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftaxGrowth         float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftaxgrowth,omitempty"`
	IncomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterestGrowth float64     `json:"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterestgrowth,omitempty"`
	EarningspersharebasicGrowth                                                                       float64     `json:"earningspersharebasicgrowth,omitempty"`
	IncometaxexpensebenefitGrowth                                                                     float64     `json:"incometaxexpensebenefitgrowth,omitempty"`
	RevenuefromcontractwithcustomerexcludingassessedtaxGrowth                                         interface{} `json:"revenuefromcontractwithcustomerexcludingassessedtaxgrowth,omitempty"`
	NonoperatingincomeexpenseGrowth                                                                   float64     `json:"nonoperatingincomeexpensegrowth,omitempty"`
	OperatingexpensesGrowth                                                                           interface{} `json:"operatingexpensesgrowth,omitempty"`
	EarningspersharedilutedGrowth                                                                     float64     `json:"earningspersharedilutedgrowth,omitempty"`
	OthercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftaxGrowth    float64     `json:"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftaxgrowth,omitempty"`
	SellinggeneralandadministrativeexpenseGrowth                                                      float64     `json:"sellinggeneralandadministrativeexpensegrowth,omitempty"`
	OthercomprehensiveincomelossnetoftaxportionattributabletoparentGrowth                             float64     `json:"othercomprehensiveincomelossnetoftaxportionattributabletoparentgrowth,omitempty"`
	ComprehensiveincomenetoftaxGrowth                                                                 float64     `json:"comprehensiveincomenetoftaxgrowth,omitempty"`
	OthercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtaxGrowth         float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtaxgrowth,omitempty"`
	OthercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertaxGrowth                   float64     `json:"othercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertaxgrowth,omitempty"`
}

// Financial Ratios
type FinancialRatiosGrowth struct {
	Symbol                                   string  `json:"symbol" csv:"symbol,omitempty"`
	Date                                     string  `json:"date" csv:"date,omitempty"`
	Period                                   string  `json:"period" csv:"period,omitempty"`
	CurrentRatioGrowth                       float64 `json:"currentRatio" csv:"currentRatioGrowth,omitempty"`
	QuickRatioGrowth                         float64 `json:"quickRatio" csv:"quickRatioGrowth,omitempty"`
	CashRatioGrowth                          float64 `json:"cashRatio" csv:"cashRatioGrowth,omitempty"`
	DaysOfSalesOutstandingGrowth             float64 `json:"daysOfSalesOutstanding" csv:"daysOfSalesOutstandingGrowth,omitempty"`
	DaysOfInventoryOutstandingGrowth         float64 `json:"daysOfInventoryOutstanding" csv:"daysOfInventoryOutstandingGrowth,omitempty"`
	OperatingCycleGrowth                     float64 `json:"operatingCycle" csv:"operatingCycleGrowth,omitempty"`
	DaysOfPayablesOutstandingGrowth          float64 `json:"daysOfPayablesOutstanding" csv:"daysOfPayablesOutstandingGrowth,omitempty"`
	CashConversionCycleGrowth                float64 `json:"cashConversionCycle" csv:"cashConversionCycleGrowth,omitempty"`
	GrossProfitMarginGrowth                  float64 `json:"grossProfitMargin" csv:"grossProfitMarginGrowth,omitempty"`
	OperatingProfitMarginGrowth              float64 `json:"operatingProfitMargin" csv:"operatingProfitMarginGrowth,omitempty"`
	PretaxProfitMarginGrowth                 float64 `json:"pretaxProfitMargin" csv:"pretaxProfitMarginGrowth,omitempty"`
	NetProfitMarginGrowth                    float64 `json:"netProfitMargin" csv:"netProfitMarginGrowth,omitempty"`
	EffectiveTaxRateGrowth                   float64 `json:"effectiveTaxRate" csv:"effectiveTaxRateGrowth,omitempty"`
	ReturnOnAssetsGrowth                     float64 `json:"returnOnAssets" csv:"returnOnAssetsGrowth,omitempty"`
	ReturnOnEquityGrowth                     float64 `json:"returnOnEquity" csv:"returnOnEquityGrowth,omitempty"`
	ReturnOnCapitalEmployedGrowth            float64 `json:"returnOnCapitalEmployed" csv:"returnOnCapitalEmployedGrowth,omitempty"`
	NetIncomePerEBTGrowth                    float64 `json:"netIncomePerEBT" csv:"netIncomePerEBTGrowth,omitempty"`
	EbtPerEbitGrowth                         float64 `json:"ebtPerEbit" csv:"ebtPerEbitGrowth,omitempty"`
	EbitPerRevenueGrowth                     float64 `json:"ebitPerRevenue" csv:"ebitPerRevenueGrowth,omitempty"`
	DebtRatioGrowth                          float64 `json:"debtRatio" csv:"debtRatioGrowth,omitempty"`
	DebtEquityRatioGrowth                    float64 `json:"debtEquityRatio" csv:"debtEquityRatioGrowth,omitempty"`
	LongTermDebtToCapitalizationGrowth       float64 `json:"longTermDebtToCapitalization" csv:"longTermDebtToCapitalizationGrowth,omitempty"`
	TotalDebtToCapitalizationGrowth          float64 `json:"totalDebtToCapitalization" csv:"totalDebtToCapitalizationGrowth,omitempty"`
	InterestCoverageGrowth                   float64 `json:"interestCoverage" csv:"interestCoverageGrowth,omitempty"`
	CashFlowToDebtRatioGrowth                float64 `json:"cashFlowToDebtRatio" csv:"cashFlowToDebtRatioGrowth,omitempty"`
	CompanyEquityMultiplierGrowth            float64 `json:"companyEquityMultiplier" csv:"companyEquityMultiplierGrowth,omitempty"`
	ReceivablesTurnoverGrowth                float64 `json:"receivablesTurnover" csv:"receivablesTurnoverGrowth,omitempty"`
	PayablesTurnoverGrowth                   float64 `json:"payablesTurnover" csv:"payablesTurnoverGrowth,omitempty"`
	InventoryTurnoverGrowth                  float64 `json:"inventoryTurnover" csv:"inventoryTurnoverGrowth,omitempty"`
	FixedAssetTurnoverGrowth                 float64 `json:"fixedAssetTurnover" csv:"fixedAssetTurnoverGrowth,omitempty"`
	AssetTurnoverGrowth                      float64 `json:"assetTurnover" csv:"assetTurnoverGrowth,omitempty"`
	OperatingCashFlowPerShareGrowth          float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShareGrowth,omitempty"`
	FreeCashFlowPerShareGrowth               float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShareGrowth,omitempty"`
	CashPerShareGrowth                       float64 `json:"cashPerShare" csv:"cashPerShareGrowth,omitempty"`
	PayoutRatioGrowth                        float64 `json:"payoutRatio" csv:"payoutRatioGrowth,omitempty"`
	OperatingCashFlowSalesRatioGrowth        float64 `json:"operatingCashFlowSalesRatio" csv:"operatingCashFlowSalesRatioGrowth,omitempty"`
	FreeCashFlowOperatingCashFlowRatioGrowth float64 `json:"freeCashFlowOperatingCashFlowRatio" csv:"freeCashFlowOperatingCashFlowRatioGrowth,omitempty"`
	CashFlowCoverageRatiosGrowth             float64 `json:"cashFlowCoverageRatios" csv:"cashFlowCoverageRatiosGrowth,omitempty"`
	ShortTermCoverageRatiosGrowth            float64 `json:"shortTermCoverageRatios" csv:"shortTermCoverageRatiosGrowth,omitempty"`
	CapitalExpenditureCoverageRatioGrowth    float64 `json:"capitalExpenditureCoverageRatio" csv:"capitalExpenditureCoverageRatioGrowth,omitempty"`
	DividendPaidAndCapexCoverageRatioGrowth  float64 `json:"dividendPaidAndCapexCoverageRatio" csv:"dividendPaidAndCapexCoverageRatioGrowth,omitempty"`
	DividendPayoutRatioGrowth                float64 `json:"dividendPayoutRatio" csv:"dividendPayoutRatioGrowth,omitempty"`
	PriceBookValueRatioGrowth                float64 `json:"priceBookValueRatio" csv:"priceBookValueRatioGrowth,omitempty"`
	PriceToBookRatioGrowth                   float64 `json:"priceToBookRatio" csv:"priceToBookRatioGrowth,omitempty"`
	PriceToSalesRatioGrowth                  float64 `json:"priceToSalesRatio" csv:"priceToSalesRatioGrowth,omitempty"`
	PriceEarningsRatioGrowth                 float64 `json:"priceEarningsRatio" csv:"priceEarningsRatioGrowth,omitempty"`
	PriceToFreeCashFlowsRatioGrowth          float64 `json:"priceToFreeCashFlowsRatio" csv:"priceToFreeCashFlowsRatioGrowth,omitempty"`
	PriceToOperatingCashFlowsRatioGrowth     float64 `json:"priceToOperatingCashFlowsRatio" csv:"priceToOperatingCashFlowsRatioGrowth,omitempty"`
	PriceCashFlowRatioGrowth                 float64 `json:"priceCashFlowRatio" csv:"priceCashFlowRatioGrowth,omitempty"`
	PriceEarningsToGrowthRatioGrowth         float64 `json:"priceEarningsToGrowthRatio" csv:"priceEarningsToGrowthRatioGrowth,omitempty"`
	PriceSalesRatioGrowth                    float64 `json:"priceSalesRatio" csv:"priceSalesRatioGrowth,omitempty"`
	DividendYieldGrowth                      float64 `json:"dividendYield" csv:"dividendYieldGrowth,omitempty"`
	EnterpriseValueMultipleGrowth            float64 `json:"enterpriseValueMultiple" csv:"enterpriseValueMultipleGrowth,omitempty"`
	PriceFairValueGrowth                     float64 `json:"priceFairValue" csv:"priceFairValueGrowth,omitempty"`
}

type FinancialRatiosTTMGrowth struct {
	Symbol                                      string  `csv:"symbol,omitempty"`
	DividendYielTTMGrowth                       float64 `json:"dividendYielTTM" csv:"dividendYielTTM,omitempty"`
	DividendYielPercentageTTMGrowth             float64 `json:"dividendYielPercentageTTM" csv:"dividendYielPercentageTTM,omitempty"`
	PeRatioTTMGrowth                            float64 `json:"peRatioTTM" csv:"peRatioTTM,omitempty"`
	PegRatioTTMGrowth                           float64 `json:"pegRatioTTM" csv:"pegRatioTTM,omitempty"`
	PayoutRatioTTMGrowth                        float64 `json:"payoutRatioTTM" csv:"payoutRatioTTM,omitempty"`
	CurrentRatioTTMGrowth                       float64 `json:"currentRatioTTM" csv:"currentRatioTTM,omitempty"`
	QuickRatioTTMGrowth                         float64 `json:"quickRatioTTM" csv:"quickRatioTTM,omitempty"`
	CashRatioTTMGrowth                          float64 `json:"cashRatioTTM" csv:"cashRatioTTM,omitempty"`
	DaysOfSalesOutstandingTTMGrowth             float64 `json:"daysOfSalesOutstandingTTM" csv:"daysOfSalesOutstandingTTM,omitempty"`
	DaysOfInventoryOutstandingTTMGrowth         float64 `json:"daysOfInventoryOutstandingTTM" csv:"daysOfInventoryOutstandingTTM,omitempty"`
	OperatingCycleTTMGrowth                     float64 `json:"operatingCycleTTM" csv:"operatingCycleTTM,omitempty"`
	DaysOfPayablesOutstandingTTMGrowth          float64 `json:"daysOfPayablesOutstandingTTM" csv:"daysOfPayablesOutstandingTTM,omitempty"`
	CashConversionCycleTTMGrowth                float64 `json:"cashConversionCycleTTM" csv:"cashConversionCycleTTM,omitempty"`
	GrossProfitMarginTTMGrowth                  float64 `json:"grossProfitMarginTTM" csv:"grossProfitMarginTTM,omitempty"`
	OperatingProfitMarginTTMGrowth              float64 `json:"operatingProfitMarginTTM" csv:"operatingProfitMarginTTM,omitempty"`
	PretaxProfitMarginTTMGrowth                 float64 `json:"pretaxProfitMarginTTM" csv:"pretaxProfitMarginTTM,omitempty"`
	NetProfitMarginTTMGrowth                    float64 `json:"netProfitMarginTTM" csv:"netProfitMarginTTM,omitempty"`
	EffectiveTaxRateTTMGrowth                   float64 `json:"effectiveTaxRateTTM" csv:"effectiveTaxRateTTM,omitempty"`
	ReturnOnAssetsTTMGrowth                     float64 `json:"returnOnAssetsTTM" csv:"returnOnAssetsTTM,omitempty"`
	ReturnOnEquityTTMGrowth                     float64 `json:"returnOnEquityTTM" csv:"returnOnEquityTTM,omitempty"`
	ReturnOnCapitalEmployedTTMGrowth            float64 `json:"returnOnCapitalEmployedTTM" csv:"returnOnCapitalEmployedTTM,omitempty"`
	NetIncomePerEBTTTMGrowth                    float64 `json:"netIncomePerEBTTTM" csv:"netIncomePerEBTTTM,omitempty"`
	EbtPerEbitTTMGrowth                         float64 `json:"ebtPerEbitTTM" csv:"ebtPerEbitTTM,omitempty"`
	EbitPerRevenueTTMGrowth                     float64 `json:"ebitPerRevenueTTM" csv:"ebitPerRevenueTTM,omitempty"`
	DebtRatioTTMGrowth                          float64 `json:"debtRatioTTM" csv:"debtRatioTTM,omitempty"`
	DebtEquityRatioTTMGrowth                    float64 `json:"debtEquityRatioTTM" csv:"debtEquityRatioTTM,omitempty"`
	LongTermDebtToCapitalizationTTMGrowth       float64 `json:"longTermDebtToCapitalizationTTM" csv:"longTermDebtToCapitalizationTTM,omitempty"`
	TotalDebtToCapitalizationTTMGrowth          float64 `json:"totalDebtToCapitalizationTTM" csv:"totalDebtToCapitalizationTTM,omitempty"`
	InterestCoverageTTMGrowth                   float64 `json:"interestCoverageTTM" csv:"interestCoverageTTM,omitempty"`
	CashFlowToDebtRatioTTMGrowth                float64 `json:"cashFlowToDebtRatioTTM" csv:"cashFlowToDebtRatioTTM,omitempty"`
	CompanyEquityMultiplierTTMGrowth            float64 `json:"companyEquityMultiplierTTM" csv:"companyEquityMultiplierTTM,omitempty"`
	ReceivablesTurnoverTTMGrowth                float64 `json:"receivablesTurnoverTTM" csv:"receivablesTurnoverTTM,omitempty"`
	PayablesTurnoverTTMGrowth                   float64 `json:"payablesTurnoverTTM" csv:"payablesTurnoverTTM,omitempty"`
	InventoryTurnoverTTMGrowth                  float64 `json:"inventoryTurnoverTTM" csv:"inventoryTurnoverTTM,omitempty"`
	FixedAssetTurnoverTTMGrowth                 float64 `json:"fixedAssetTurnoverTTM" csv:"fixedAssetTurnoverTTM,omitempty"`
	AssetTurnoverTTMGrowth                      float64 `json:"assetTurnoverTTM" csv:"assetTurnoverTTM,omitempty"`
	OperatingCashFlowPerShareTTMGrowth          float64 `json:"operatingCashFlowPerShareTTM" csv:"operatingCashFlowPerShareTTM,omitempty"`
	FreeCashFlowPerShareTTMGrowth               float64 `json:"freeCashFlowPerShareTTM" csv:"freeCashFlowPerShareTTM,omitempty"`
	CashPerShareTTMGrowth                       float64 `json:"cashPerShareTTM" csv:"cashPerShareTTM,omitempty"`
	OperatingCashFlowSalesRatioTTMGrowth        float64 `json:"operatingCashFlowSalesRatioTTM" csv:"operatingCashFlowSalesRatioTTM,omitempty"`
	FreeCashFlowOperatingCashFlowRatioTTMGrowth float64 `json:"freeCashFlowOperatingCashFlowRatioTTM" csv:"freeCashFlowOperatingCashFlowRatioTTM,omitempty"`
	CashFlowCoverageRatiosTTMGrowth             float64 `json:"cashFlowCoverageRatiosTTM" csv:"cashFlowCoverageRatiosTTM,omitempty"`
	ShortTermCoverageRatiosTTMGrowth            float64 `json:"shortTermCoverageRatiosTTM" csv:"shortTermCoverageRatiosTTM,omitempty"`
	CapitalExpenditureCoverageRatioTTMGrowth    float64 `json:"capitalExpenditureCoverageRatioTTM" csv:"capitalExpenditureCoverageRatioTTM,omitempty"`
	DividendPaidAndCapexCoverageRatioTTMGrowth  float64 `json:"dividendPaidAndCapexCoverageRatioTTM" csv:"dividendPaidAndCapexCoverageRatioTTM,omitempty"`
	PriceBookValueRatioTTMGrowth                float64 `json:"priceBookValueRatioTTM" csv:"priceBookValueRatioTTM,omitempty"`
	PriceToBookRatioTTMGrowth                   float64 `json:"priceToBookRatioTTM" csv:"priceToBookRatioTTM,omitempty"`
	PriceToSalesRatioTTMGrowth                  float64 `json:"priceToSalesRatioTTM" csv:"priceToSalesRatioTTM,omitempty"`
	PriceEarningsRatioTTMGrowth                 float64 `json:"priceEarningsRatioTTM" csv:"priceEarningsRatioTTM,omitempty"`
	PriceToFreeCashFlowsRatioTTMGrowth          float64 `json:"priceToFreeCashFlowsRatioTTM" csv:"priceToFreeCashFlowsRatioTTM,omitempty"`
	PriceToOperatingCashFlowsRatioTTMGrowth     float64 `json:"priceToOperatingCashFlowsRatioTTM" csv:"priceToOperatingCashFlowsRatioTTM,omitempty"`
	PriceCashFlowRatioTTMGrowth                 float64 `json:"priceCashFlowRatioTTM" csv:"priceCashFlowRatioTTM,omitempty"`
	PriceEarningsToGrowthRatioTTMGrowth         float64 `json:"priceEarningsToGrowthRatioTTM" csv:"priceEarningsToGrowthRatioTTM,omitempty"`
	PriceSalesRatioTTMGrowth                    float64 `json:"priceSalesRatioTTM" csv:"priceSalesRatioTTM,omitempty"`
	DividendYieldTTMGrowth                      float64 `json:"dividendYieldTTM" csv:"dividendYieldTTM,omitempty"`
	EnterpriseValueMultipleTTMGrowth            float64 `json:"enterpriseValueMultipleTTM" csv:"enterpriseValueMultipleTTM,omitempty"`
	PriceFairValueTTMGrowth                     float64 `json:"priceFairValueTTM" csv:"priceFairValueTTM,omitempty"`
	DividendPerShareTTMGrowth                   float64 `json:"dividendPerShareTTM" csv:"dividendPerShareTTM,omitempty"`
}
