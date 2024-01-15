package Fundamentals

// Balance Sheet
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

// CashFlow Statement
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

// Income Statement
type DiscrepancyIncomeStatementAndIncomeStatementAsReported struct {
	Date                                      string  `json:"date"`
	Symbol                                    string  `json:"symbol"`
	Period                                    string  `json:"period"`
	NetIncomeDiscrepancy                      float64 `json:"netIncomeDiscrepancy"`
	GrossProfitDiscrepancy                    float64 `json:"grossProfitDiscrepancy"`
	ResearchAndDevelopmentExpensesDiscrepancy float64 `json:"researchAndDevelopmentExpensesDiscrepancy"`
	OperatingIncomeDiscrepancy                float64 `json:"operatingIncomeDiscrepancy"`
	EpsDiscrepancy                            float64 `json:"epsDiscrepancy"`
	EpsDilutedDiscrepancy                     float64 `json:"epsDilutedDiscrepancy"`
	WeightedAverageShsOutDiscrepancy          float64 `json:"weightedAverageShsOutDiscrepancy"`
	WeightedAverageShsOutDilDiscrepancy       float64 `json:"weightedAverageShsOutDilDiscrepancy"`
	IncomeTaxExpenseDiscrepancy               float64 `json:"incomeTaxExpenseDiscrepancy"`
}

type GrowthIncomeStatementAsReported struct {
	Date                                                                                              string      `json:"date"`
	Symbol                                                                                            string      `json:"symbol"`
	Period                                                                                            string      `json:"period"`
	CostofgoodsandservicessoldGrowth                                                                  interface{} `json:"costofgoodsandservicessoldgrowth"`
	NetincomelossGrowth                                                                               float64     `json:"netincomelossgrowth"`
	ResearchanddevelopmentexpenseGrowth                                                               float64     `json:"researchanddevelopmentexpensegrowth"`
	GrossprofitGrowth                                                                                 interface{} `json:"grossprofitgrowth"`
	OthercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftaxGrowth   float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftaxgrowth"`
	OthercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftaxGrowth                    float64     `json:"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftaxgrowth"`
	OthercomprehensiveincomelossderivativesqualifyingashedgesnetoftaxGrowth                           float64     `json:"othercomprehensiveincomelossderivativesqualifyingashedgesnetoftaxgrowth"`
	OthercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftaxGrowth      float64     `json:"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftaxgrowth"`
	OthercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertaxGrowth      float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertaxgrowth"`
	WeightedaveragenumberofdilutedsharesoutstandingGrowth                                             float64     `json:"weightedaveragenumberofdilutedsharesoutstandinggrowth"`
	WeightedaveragenumberofsharesoutstandingbasicGrowth                                               float64     `json:"weightedaveragenumberofsharesoutstandingbasicgrowth"`
	OthercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftaxGrowth          float64     `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftaxgrowth"`
	OperatingincomelossGrowth                                                                         float64     `json:"operatingincomelossgrowth"`
	OthercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftaxGrowth         float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftaxgrowth"`
	IncomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterestGrowth float64     `json:"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterestgrowth"`
	EarningspersharebasicGrowth                                                                       float64     `json:"earningspersharebasicgrowth"`
	IncometaxexpensebenefitGrowth                                                                     float64     `json:"incometaxexpensebenefitgrowth"`
	RevenuefromcontractwithcustomerexcludingassessedtaxGrowth                                         interface{} `json:"revenuefromcontractwithcustomerexcludingassessedtaxgrowth"`
	NonoperatingincomeexpenseGrowth                                                                   float64     `json:"nonoperatingincomeexpensegrowth"`
	OperatingexpensesGrowth                                                                           interface{} `json:"operatingexpensesgrowth"`
	EarningspersharedilutedGrowth                                                                     float64     `json:"earningspersharedilutedgrowth"`
	OthercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftaxGrowth    float64     `json:"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftaxgrowth"`
	SellinggeneralandadministrativeexpenseGrowth                                                      float64     `json:"sellinggeneralandadministrativeexpensegrowth"`
	OthercomprehensiveincomelossnetoftaxportionattributabletoparentGrowth                             float64     `json:"othercomprehensiveincomelossnetoftaxportionattributabletoparentgrowth"`
	ComprehensiveincomenetoftaxGrowth                                                                 float64     `json:"comprehensiveincomenetoftaxgrowth"`
	OthercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtaxGrowth         float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtaxgrowth"`
	OthercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertaxGrowth                   float64     `json:"othercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertaxgrowth"`
}

// Financial Ratios
type FinancialRatiosGrowth struct {
	Symbol                                   string  `json:"symbol" csv:"symbol"`
	Date                                     string  `json:"date" csv:"date"`
	Period                                   string  `json:"period" csv:"period"`
	CurrentRatioGrowth                       float64 `json:"currentRatio" csv:"currentRatioGrowth"`
	QuickRatioGrowth                         float64 `json:"quickRatio" csv:"quickRatioGrowth"`
	CashRatioGrowth                          float64 `json:"cashRatio" csv:"cashRatioGrowth"`
	DaysOfSalesOutstandingGrowth             float64 `json:"daysOfSalesOutstanding" csv:"daysOfSalesOutstandingGrowth"`
	DaysOfInventoryOutstandingGrowth         float64 `json:"daysOfInventoryOutstanding" csv:"daysOfInventoryOutstandingGrowth"`
	OperatingCycleGrowth                     float64 `json:"operatingCycle" csv:"operatingCycleGrowth"`
	DaysOfPayablesOutstandingGrowth          float64 `json:"daysOfPayablesOutstanding" csv:"daysOfPayablesOutstandingGrowth"`
	CashConversionCycleGrowth                float64 `json:"cashConversionCycle" csv:"cashConversionCycleGrowth"`
	GrossProfitMarginGrowth                  float64 `json:"grossProfitMargin" csv:"grossProfitMarginGrowth"`
	OperatingProfitMarginGrowth              float64 `json:"operatingProfitMargin" csv:"operatingProfitMarginGrowth"`
	PretaxProfitMarginGrowth                 float64 `json:"pretaxProfitMargin" csv:"pretaxProfitMarginGrowth"`
	NetProfitMarginGrowth                    float64 `json:"netProfitMargin" csv:"netProfitMarginGrowth"`
	EffectiveTaxRateGrowth                   float64 `json:"effectiveTaxRate" csv:"effectiveTaxRateGrowth"`
	ReturnOnAssetsGrowth                     float64 `json:"returnOnAssets" csv:"returnOnAssetsGrowth"`
	ReturnOnEquityGrowth                     float64 `json:"returnOnEquity" csv:"returnOnEquityGrowth"`
	ReturnOnCapitalEmployedGrowth            float64 `json:"returnOnCapitalEmployed" csv:"returnOnCapitalEmployedGrowth"`
	NetIncomePerEBTGrowth                    float64 `json:"netIncomePerEBT" csv:"netIncomePerEBTGrowth"`
	EbtPerEbitGrowth                         float64 `json:"ebtPerEbit" csv:"ebtPerEbitGrowth"`
	EbitPerRevenueGrowth                     float64 `json:"ebitPerRevenue" csv:"ebitPerRevenueGrowth"`
	DebtRatioGrowth                          float64 `json:"debtRatio" csv:"debtRatioGrowth"`
	DebtEquityRatioGrowth                    float64 `json:"debtEquityRatio" csv:"debtEquityRatioGrowth"`
	LongTermDebtToCapitalizationGrowth       float64 `json:"longTermDebtToCapitalization" csv:"longTermDebtToCapitalizationGrowth"`
	TotalDebtToCapitalizationGrowth          float64 `json:"totalDebtToCapitalization" csv:"totalDebtToCapitalizationGrowth"`
	InterestCoverageGrowth                   float64 `json:"interestCoverage" csv:"interestCoverageGrowth"`
	CashFlowToDebtRatioGrowth                float64 `json:"cashFlowToDebtRatio" csv:"cashFlowToDebtRatioGrowth"`
	CompanyEquityMultiplierGrowth            float64 `json:"companyEquityMultiplier" csv:"companyEquityMultiplierGrowth"`
	ReceivablesTurnoverGrowth                float64 `json:"receivablesTurnover" csv:"receivablesTurnoverGrowth"`
	PayablesTurnoverGrowth                   float64 `json:"payablesTurnover" csv:"payablesTurnoverGrowth"`
	InventoryTurnoverGrowth                  float64 `json:"inventoryTurnover" csv:"inventoryTurnoverGrowth"`
	FixedAssetTurnoverGrowth                 float64 `json:"fixedAssetTurnover" csv:"fixedAssetTurnoverGrowth"`
	AssetTurnoverGrowth                      float64 `json:"assetTurnover" csv:"assetTurnoverGrowth"`
	OperatingCashFlowPerShareGrowth          float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShareGrowth"`
	FreeCashFlowPerShareGrowth               float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShareGrowth"`
	CashPerShareGrowth                       float64 `json:"cashPerShare" csv:"cashPerShareGrowth"`
	PayoutRatioGrowth                        float64 `json:"payoutRatio" csv:"payoutRatioGrowth"`
	OperatingCashFlowSalesRatioGrowth        float64 `json:"operatingCashFlowSalesRatio" csv:"operatingCashFlowSalesRatioGrowth"`
	FreeCashFlowOperatingCashFlowRatioGrowth float64 `json:"freeCashFlowOperatingCashFlowRatio" csv:"freeCashFlowOperatingCashFlowRatioGrowth"`
	CashFlowCoverageRatiosGrowth             float64 `json:"cashFlowCoverageRatios" csv:"cashFlowCoverageRatiosGrowth"`
	ShortTermCoverageRatiosGrowth            float64 `json:"shortTermCoverageRatios" csv:"shortTermCoverageRatiosGrowth"`
	CapitalExpenditureCoverageRatioGrowth    float64 `json:"capitalExpenditureCoverageRatio" csv:"capitalExpenditureCoverageRatioGrowth"`
	DividendPaidAndCapexCoverageRatioGrowth  float64 `json:"dividendPaidAndCapexCoverageRatio" csv:"dividendPaidAndCapexCoverageRatioGrowth"`
	DividendPayoutRatioGrowth                float64 `json:"dividendPayoutRatio" csv:"dividendPayoutRatioGrowth"`
	PriceBookValueRatioGrowth                float64 `json:"priceBookValueRatio" csv:"priceBookValueRatioGrowth"`
	PriceToBookRatioGrowth                   float64 `json:"priceToBookRatio" csv:"priceToBookRatioGrowth"`
	PriceToSalesRatioGrowth                  float64 `json:"priceToSalesRatio" csv:"priceToSalesRatioGrowth"`
	PriceEarningsRatioGrowth                 float64 `json:"priceEarningsRatio" csv:"priceEarningsRatioGrowth"`
	PriceToFreeCashFlowsRatioGrowth          float64 `json:"priceToFreeCashFlowsRatio" csv:"priceToFreeCashFlowsRatioGrowth"`
	PriceToOperatingCashFlowsRatioGrowth     float64 `json:"priceToOperatingCashFlowsRatio" csv:"priceToOperatingCashFlowsRatioGrowth"`
	PriceCashFlowRatioGrowth                 float64 `json:"priceCashFlowRatio" csv:"priceCashFlowRatioGrowth"`
	PriceEarningsToGrowthRatioGrowth         float64 `json:"priceEarningsToGrowthRatio" csv:"priceEarningsToGrowthRatioGrowth"`
	PriceSalesRatioGrowth                    float64 `json:"priceSalesRatio" csv:"priceSalesRatioGrowth"`
	DividendYieldGrowth                      float64 `json:"dividendYield" csv:"dividendYieldGrowth"`
	EnterpriseValueMultipleGrowth            float64 `json:"enterpriseValueMultiple" csv:"enterpriseValueMultipleGrowth"`
	PriceFairValueGrowth                     float64 `json:"priceFairValue" csv:"priceFairValueGrowth"`
}

type FinancialRatiosTTMGrowth struct {
	Symbol                                      string  `csv:"symbol"`
	DividendYielTTMGrowth                       float64 `json:"dividendYielTTM" csv:"dividendYielTTM"`
	DividendYielPercentageTTMGrowth             float64 `json:"dividendYielPercentageTTM" csv:"dividendYielPercentageTTM"`
	PeRatioTTMGrowth                            float64 `json:"peRatioTTM" csv:"peRatioTTM"`
	PegRatioTTMGrowth                           float64 `json:"pegRatioTTM" csv:"pegRatioTTM"`
	PayoutRatioTTMGrowth                        float64 `json:"payoutRatioTTM" csv:"payoutRatioTTM"`
	CurrentRatioTTMGrowth                       float64 `json:"currentRatioTTM" csv:"currentRatioTTM"`
	QuickRatioTTMGrowth                         float64 `json:"quickRatioTTM" csv:"quickRatioTTM"`
	CashRatioTTMGrowth                          float64 `json:"cashRatioTTM" csv:"cashRatioTTM"`
	DaysOfSalesOutstandingTTMGrowth             float64 `json:"daysOfSalesOutstandingTTM" csv:"daysOfSalesOutstandingTTM"`
	DaysOfInventoryOutstandingTTMGrowth         float64 `json:"daysOfInventoryOutstandingTTM" csv:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTMGrowth                     float64 `json:"operatingCycleTTM" csv:"operatingCycleTTM"`
	DaysOfPayablesOutstandingTTMGrowth          float64 `json:"daysOfPayablesOutstandingTTM" csv:"daysOfPayablesOutstandingTTM"`
	CashConversionCycleTTMGrowth                float64 `json:"cashConversionCycleTTM" csv:"cashConversionCycleTTM"`
	GrossProfitMarginTTMGrowth                  float64 `json:"grossProfitMarginTTM" csv:"grossProfitMarginTTM"`
	OperatingProfitMarginTTMGrowth              float64 `json:"operatingProfitMarginTTM" csv:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTMGrowth                 float64 `json:"pretaxProfitMarginTTM" csv:"pretaxProfitMarginTTM"`
	NetProfitMarginTTMGrowth                    float64 `json:"netProfitMarginTTM" csv:"netProfitMarginTTM"`
	EffectiveTaxRateTTMGrowth                   float64 `json:"effectiveTaxRateTTM" csv:"effectiveTaxRateTTM"`
	ReturnOnAssetsTTMGrowth                     float64 `json:"returnOnAssetsTTM" csv:"returnOnAssetsTTM"`
	ReturnOnEquityTTMGrowth                     float64 `json:"returnOnEquityTTM" csv:"returnOnEquityTTM"`
	ReturnOnCapitalEmployedTTMGrowth            float64 `json:"returnOnCapitalEmployedTTM" csv:"returnOnCapitalEmployedTTM"`
	NetIncomePerEBTTTMGrowth                    float64 `json:"netIncomePerEBTTTM" csv:"netIncomePerEBTTTM"`
	EbtPerEbitTTMGrowth                         float64 `json:"ebtPerEbitTTM" csv:"ebtPerEbitTTM"`
	EbitPerRevenueTTMGrowth                     float64 `json:"ebitPerRevenueTTM" csv:"ebitPerRevenueTTM"`
	DebtRatioTTMGrowth                          float64 `json:"debtRatioTTM" csv:"debtRatioTTM"`
	DebtEquityRatioTTMGrowth                    float64 `json:"debtEquityRatioTTM" csv:"debtEquityRatioTTM"`
	LongTermDebtToCapitalizationTTMGrowth       float64 `json:"longTermDebtToCapitalizationTTM" csv:"longTermDebtToCapitalizationTTM"`
	TotalDebtToCapitalizationTTMGrowth          float64 `json:"totalDebtToCapitalizationTTM" csv:"totalDebtToCapitalizationTTM"`
	InterestCoverageTTMGrowth                   float64 `json:"interestCoverageTTM" csv:"interestCoverageTTM"`
	CashFlowToDebtRatioTTMGrowth                float64 `json:"cashFlowToDebtRatioTTM" csv:"cashFlowToDebtRatioTTM"`
	CompanyEquityMultiplierTTMGrowth            float64 `json:"companyEquityMultiplierTTM" csv:"companyEquityMultiplierTTM"`
	ReceivablesTurnoverTTMGrowth                float64 `json:"receivablesTurnoverTTM" csv:"receivablesTurnoverTTM"`
	PayablesTurnoverTTMGrowth                   float64 `json:"payablesTurnoverTTM" csv:"payablesTurnoverTTM"`
	InventoryTurnoverTTMGrowth                  float64 `json:"inventoryTurnoverTTM" csv:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTMGrowth                 float64 `json:"fixedAssetTurnoverTTM" csv:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTMGrowth                      float64 `json:"assetTurnoverTTM" csv:"assetTurnoverTTM"`
	OperatingCashFlowPerShareTTMGrowth          float64 `json:"operatingCashFlowPerShareTTM" csv:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTMGrowth               float64 `json:"freeCashFlowPerShareTTM" csv:"freeCashFlowPerShareTTM"`
	CashPerShareTTMGrowth                       float64 `json:"cashPerShareTTM" csv:"cashPerShareTTM"`
	OperatingCashFlowSalesRatioTTMGrowth        float64 `json:"operatingCashFlowSalesRatioTTM" csv:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTMGrowth float64 `json:"freeCashFlowOperatingCashFlowRatioTTM" csv:"freeCashFlowOperatingCashFlowRatioTTM"`
	CashFlowCoverageRatiosTTMGrowth             float64 `json:"cashFlowCoverageRatiosTTM" csv:"cashFlowCoverageRatiosTTM"`
	ShortTermCoverageRatiosTTMGrowth            float64 `json:"shortTermCoverageRatiosTTM" csv:"shortTermCoverageRatiosTTM"`
	CapitalExpenditureCoverageRatioTTMGrowth    float64 `json:"capitalExpenditureCoverageRatioTTM" csv:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTMGrowth  float64 `json:"dividendPaidAndCapexCoverageRatioTTM" csv:"dividendPaidAndCapexCoverageRatioTTM"`
	PriceBookValueRatioTTMGrowth                float64 `json:"priceBookValueRatioTTM" csv:"priceBookValueRatioTTM"`
	PriceToBookRatioTTMGrowth                   float64 `json:"priceToBookRatioTTM" csv:"priceToBookRatioTTM"`
	PriceToSalesRatioTTMGrowth                  float64 `json:"priceToSalesRatioTTM" csv:"priceToSalesRatioTTM"`
	PriceEarningsRatioTTMGrowth                 float64 `json:"priceEarningsRatioTTM" csv:"priceEarningsRatioTTM"`
	PriceToFreeCashFlowsRatioTTMGrowth          float64 `json:"priceToFreeCashFlowsRatioTTM" csv:"priceToFreeCashFlowsRatioTTM"`
	PriceToOperatingCashFlowsRatioTTMGrowth     float64 `json:"priceToOperatingCashFlowsRatioTTM" csv:"priceToOperatingCashFlowsRatioTTM"`
	PriceCashFlowRatioTTMGrowth                 float64 `json:"priceCashFlowRatioTTM" csv:"priceCashFlowRatioTTM"`
	PriceEarningsToGrowthRatioTTMGrowth         float64 `json:"priceEarningsToGrowthRatioTTM" csv:"priceEarningsToGrowthRatioTTM"`
	PriceSalesRatioTTMGrowth                    float64 `json:"priceSalesRatioTTM" csv:"priceSalesRatioTTM"`
	DividendYieldTTMGrowth                      float64 `json:"dividendYieldTTM" csv:"dividendYieldTTM"`
	EnterpriseValueMultipleTTMGrowth            float64 `json:"enterpriseValueMultipleTTM" csv:"enterpriseValueMultipleTTM"`
	PriceFairValueTTMGrowth                     float64 `json:"priceFairValueTTM" csv:"priceFairValueTTM"`
	DividendPerShareTTMGrowth                   float64 `json:"dividendPerShareTTM" csv:"dividendPerShareTTM"`
}
