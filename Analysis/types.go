package Analysis

import (
	fundamentals "github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type CompanyFundamentals struct {
	Symbol string

	BalanceSheetStatements                                             []objects.BalanceSheetStatement                                                   `json:"balancesheetstatements,omitempty"`
	BalanceSheetStatementGrowth                                        []objects.BalanceSheetStatementGrowth                                             `json:"balancesheetstatementgrowth,omitempty"`
	BalanceSheetStatementAsReported                                    []objects.BalanceSheetStatementAsReported                                         `json:"balancesheetstatementasreported,omitempty"`
	GrowthBalanceSheetStatementAsReported                              []fundamentals.GrowthBalanceSheetStatementAsReported                              `json:"growthbalancesheetstatementasreported,omitempty"`
	DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported []fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported `json:"discrepancybalancesheetstatementandbalancesheetstatementasreported,omitempty"`

	IncomeStatement                                        []objects.IncomeStatement                                             `json:"incomestatement,omitempty"`
	IncomeStatementGrowth                                  []objects.IncomeStatementGrowth                                       `json:"incomestatementgrowth,omitempty"`
	IncomeStatementAsReported                              []objects.IncomeStatementAsReported                                   `json:"incomestatementasreported,omitempty"`
	GrowthIncomeStatementAsReported                        []fundamentals.GrowthIncomeStatementAsReported                        `json:"growthincomestatementasreported,omitempty"`
	DiscrepancyIncomeStatementAndIncomeStatementAsReported []fundamentals.DiscrepancyIncomeStatementAndIncomeStatementAsReported `json:"discrepancyincomestatementandincomestatementasreported,omitempty"`

	CashFlowStatement                                          []objects.CashFlowStatement                                               `json:"cashflowstatement,omitempty"`
	CashFlowStatementGrowth                                    []objects.CashFlowStatementGrowth                                         `json:"cashflowstatementgrowth,omitempty"`
	CashFlowStatementAsReported                                []objects.CashFlowStatementAsReported                                     `json:"cashflowstatementasreported,omitempty"`
	CashFlowStatementAsReportedGrowth                          []fundamentals.CashFlowStatementAsReportedGrowth                          `json:"cashflowstatementasreportedgrowth,omitempty"`
	DiscrepancyCashFlowStatementAndCashFlowStatementAsReported []fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported `json:"discrepancycashflowstatementandcashflowstatementasreported,omitempty"`

	FinancialRatios          []objects.FinancialRatios               `json:"financialratios,omitempty"`
	FinancialRatiosTTM       []objects.FinancialRatiosTTM            `json:"financialratiosttm,omitempty"`
	FinancialRatiosGrowth    []fundamentals.FinancialRatiosGrowth    `json:"financialratiosgrowth,omitempty"`
	FinancialRatiosTTMGrowth []fundamentals.FinancialRatiosTTMGrowth `json:"financialratiosttmgrowth,omitempty"`

	FullFinancialStatement       []objects.FullFinancialStatementAsReported `json:"fullfinancialstatementasreported,omitempty"`
	FullFinancialStatementGrowth []objects.FinancialStatementsGrowth        `json:"fullfinancialstatementgrowth,omitempty"`
}

type FundamentalsCalculationsResults struct {
	Symbol           string                         `json:"symbol,omitempty"`
	Fundamentals     CompanyFundamentals            `json:"fundamentals,omitempty"`
	PeriodLength     objects.CompanyValuationPeriod `json:"periodlength,omitempty"`
	Outlook          CompanyOutlook                 `json:"outlook,omitempty"`
	NumEmployees     float64                        `json:"numemployees,omitempty"`
	CostOfEquity     float64                        `json:"costofequity,omitempty"`
	Beta             float64                        `json:"beta,omitempty"`
	EffectiveTaxRate float64                        `json:"effectivetaxrate,omitempty"`

	BalanceSheet struct {
		DifferenceInLengthBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported int                  `json:"differenceinlengthbetweenbalancesheetstatementandbalancesheetstatementasreported,omitempty"`
		StatementAndReportDiscrepancyGrowth                                              map[string][]float64 `json:"statementandreportdiscrepancygrowth,omitempty"`

		TotalGapsInBalanceSheetStatementPeriods                  int `json:"totalgapsinbalancesheetstatementperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatement int `json:"totalconsecutiveperiodswithnogapsinbalancesheetstatement,omitempty"`
		TotalConsecutiveMissingPeriodsInBalanceSheetStatement    int `json:"totalconsecutivemissingperiodsinbalancesheetstatement,omitempty"`

		TotalGapsInBalanceSheetStatementAsReportedPeriods                  int `json:"totalgapsinbalancesheetstatementasreportedperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatementAsReported int `json:"totalconsecutiveperiodswithnogapsinbalancesheetstatementasreported,omitempty"`
		TotalConsecutiveMissingPeriodsInBalanceSheetStatementAsReported    int `json:"totalconsecutivemissingperiodsinbalancesheetstatementasreported,omitempty"`

		MeanSTDBalanceSheetStatement                          map[string][]interface{} `json:"meanstdbalancesheetstatement,omitempty"`
		MeanSTDBalanceSheetStatementAsReported                map[string][]interface{} `json:"meanstdbalancesheetstatementasreported,omitempty"`
		MeanSTDBalanceSheetStatementGrowth                    map[string][]interface{} `json:"meanstdbalancesheetstatementgrowth,omitempty"`
		MeanSTDBalanceSheetStatementAsReportedGrowth          map[string][]interface{} `json:"meanstdbalancesheetstatementasreportedgrowth,omitempty"`
		MeanSTDBalanceSheetDiscrepancies                      map[string][]interface{} `json:"meanstdbalancesheetdiscrepancies,omitempty"`
		MeanZippedSTDBalanceSheetStatementAndAsReported       map[string][]interface{} `json:"meanstdzippedbalancesheetstatementandasreported,omitempty"`
		MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth map[string][]interface{} `json:"meanstdzippedbalancesheetstatementandasreportedgrowth,omitempty"`
	} `json:"balancesheet,omitempty"`

	IncomeStatement struct {
		DifferenceInLengthBetweenIncomeStatementAndIncomeStatementAsReported int                  `json:"differenceinlengthbetweenincomestatementandincomestatementasreported,omitempty"`
		StatementAndReportDiscrepancyGrowth                                  map[string][]float64 `json:"statementandreportdiscrepancygrowth,omitempty"`

		TotalGapsInIncomeStatementPeriods                  int `json:"totalgapsinincomestatementperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInIncomeStatement int `json:"totalconsecutiveperiodswithnogapsinincomestatement,omitempty"`
		TotalConsecutiveMissingPeriodsInIncomeStatement    int `json:"totalconsecutivemissingperiodsinincomestatement,omitempty"`

		TotalGapsInIncomeStatementAsReportedPeriods                  int `json:"totalgapsinincomestatementasreportedperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInIncomeStatementAsReported int `json:"totalconsecutiveperiodswithnogapsinincomestatementasreported,omitempty"`
		TotalConsecutiveMissingPeriodsInIncomeStatementAsReported    int `json:"totalconsecutivemissingperiodsinincomestatementasreported,omitempty"`

		MeanSTDIncomeStatement                          map[string][]interface{} `json:"meanstdincomestatement,omitempty"`
		MeanSTDIncomeStatementAsReported                map[string][]interface{} `json:"meanstdincomestatementasreported,omitempty"`
		MeanSTDIncomeStatementGrowth                    map[string][]interface{} `json:"meanstdincomestatementgrowth,omitempty"`
		MeanSTDIncomeStatementAsReportedGrowth          map[string][]interface{} `json:"meanstdincomestatementasreportedgrowth,omitempty"`
		MeanSTDIncomeStatementDiscrepancies             map[string][]interface{} `json:"meanstdincomestatementdiscrepancies,omitempty"`
		MeanZippedSTDIncomeStatementAndAsReported       map[string][]interface{} `json:"meanstdincomestatementandasreported,omitempty"`
		MeanZippedSTDIncomeStatementAndAsReportedGrowth map[string][]interface{} `json:"meanstdincomestatementandasreportedgrowth,omitempty"`
	} `json:"incomestatement,omitempty"`

	CashFlowStatement struct {
		DifferenceInLengthBetweenCashFlowStatementAndCashFlowStatementAsReported int                  `json:"differenceinlengthbetweencashflowstatementandcashflowstatementasreported,omitempty"`
		StatementAndReportDiscrepancyGrowth                                      map[string][]float64 `json:"statementandreportdiscrepancygrowth,omitempty"`

		TotalGapsInCashFlowStatementPeriods                  int `json:"totalgapsincashflowstatementperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInCashFlowStatement int `json:"totalconsecutiveperiodswithnogapsincashflowstatement,omitempty"`
		TotalConsecutiveMissingPeriodsInCashFlowStatement    int `json:"totalconsecutivemissingperiodsincashflowstatement,omitempty"`

		TotalGapsInCashFlowStatementAsReportedPeriods                  int `json:"totalgapsincashflowstatementasreportedperiods,omitempty"`
		TotalConsecutivePeriodsWithNoGapsInCashFlowStatementAsReported int `json:"totalconsecutiveperiodswithnogapsincashflowstatementasreported,omitempty"`
		TotalConsecutiveMissingPeriodsInCashFlowStatementAsReported    int `json:"totalconsecutivemissingperiodsincashflowstatementasreported,omitempty"`

		MeanSTDCashFlowStatement                          map[string][]interface{} `json:"meanstdcashflowstatement,omitempty"`
		MeanSTDCashFlowStatementAsReported                map[string][]interface{} `json:"meanstdcashflowstatementasreported,omitempty"`
		MeanSTDCashFlowStatementGrowth                    map[string][]interface{} `json:"meanstdcashflowstatementgrowth,omitempty"`
		MeanSTDCashFlowStatementAsReportedGrowth          map[string][]interface{} `json:"meanstdcashflowstatementasreportedgrowth,omitempty"`
		MeanSTDCashFlowStatementDiscrepancies             map[string][]interface{} `json:"meanstdcashflowstatementdiscrepancies,omitempty"`
		MeanZippedSTDCashFlowStatementAndAsReported       map[string][]interface{} `json:"meanstdcashflowstatementandasreported,omitempty"`
		MeanZippedSTDCashFlowStatementAndAsReportedGrowth map[string][]interface{} `json:"meanstdcashflowstatementandasreportedgrowth,omitempty"`
	}

	FinancialRatios struct {
		FPMRatios          []objects.FinancialRatios               `json:"financialratios,omitempty"`
		FPMRatiosTTM       []objects.FinancialRatiosTTM            `json:"financialratiosttm,omitempty"`
		FPMRatiosGrowth    []fundamentals.FinancialRatiosGrowth    `json:"financialratiosgrowth,omitempty"`
		FPMRatiosTTMGrowth []fundamentals.FinancialRatiosTTMGrowth `json:"financialratiosttmgrowth,omitempty"`

		AverageSTDFPMRatios                     map[string][]interface{} `json:"averagestdfinancialratios,omitempty"`
		AverageSTDFPMRatiosTTM                  map[string][]interface{} `json:"averagestdfinancialratiosttm,omitempty"`
		AverageSTDFPMRatiosGrowth               map[string][]interface{} `json:"averagestdfinancialratiosgrowth,omitempty"`
		AverageSTDFPMRatiosTTMGrowth            map[string][]interface{} `json:"averagestdfinancialratiosttmgrowth,omitempty"`
		AverageSTDFZippedFPMRationsAndTTMRatios map[string][]interface{} `json:"averagestdzippedfinancialratiosandfinancialratiosttm,omitempty"`
	} `json:"financialratios,omitempty"`

	CustomCalculations              []map[string]*float64    `json:"customcalculations,omitempty"`
	CustomCalculationsGrowth        map[string][]float64     `json:"customcalculationsgrowth,omitempty"`
	MeanSTDCustomCalculations       map[string][]interface{} `json:"meanstdcustomcalculations,omitempty"`
	MeanSTDCustomCalculationsGrowth map[string][]interface{} `json:"meanstdcustomcalculationsgrowth,omitempty"`

	CustomCalculationsAsReported              []map[string]*float64    `json:"customcalculationsasreported,omitempty"`
	CustomCalculationsAsReportedGrowth        map[string][]float64     `json:"customcalculationsasreportedgrowth,omitempty"`
	MeanSTDCustomCalculationsAsReported       map[string][]interface{} `json:"meanstdcustomcalculationsasreported,omitempty"`
	MeanSTDCustomCalculationsAsReportedGrowth map[string][]interface{} `json:"meanstdcustomcalculationsasreportedgrowth,omitempty"`

	MeanZippedSTDCustomCalculationsAndAsReported       map[string][]interface{} `json:"meanstdzippedcustomcalculationsandasreported,omitempty"`
	MeanZippedSTDCustomCalculationsAndAsReportedGrowth map[string][]interface{} `json:"meanstdzippedcustomcalculationsandasreportedgrowth,omitempty"`
}

type CompanyOutlook struct {
	Beta                       float64                  `json:"beta,omitempty"`
	AvgVolume                  float64                  `json:"avgvolume,omitempty"`
	MarketCap                  float64                  `json:"marketcap,omitempty"`
	LastDividend               float64                  `json:"lastdividend,omitempty"`
	Changes                    float64                  `json:"changes,omitempty"`
	FMPDcfDiff                 float64                  `json:"fmpdcfdiff,omitempty"`
	FMPDcf                     float64                  `json:"fmpdcf,omitempty"`
	StockPrice                 float64                  `json:"stockprice,omitempty"`
	TotalExecutivePay          float64                  `json:"totalexecutivepay,omitempty"`
	InsiderTrades              []map[string]*float64    `json:"insidertrades,omitempty"`
	InsiderTradesGrowth        map[string][]float64     `json:"insidertradesgrowth,omitempty"`
	InsiderTradesMeanSTD       map[string][]interface{} `json:"insidertradesmeanstd,omitempty"`
	InsiderTradesGrowthMeanSTD map[string][]interface{} `json:"insidertradesgrowthmeanstd,omitempty"`
	Splits                     []float64                `json:"splits,omitempty"`
	SplitMeanSTD               map[string][]interface{} `json:"splitmeanstd,omitempty"`
	DividendHist               []map[string]*float64    `json:"dividendhist,omitempty"`
	DividendHistGrowth         map[string][]float64     `json:"dividendhistgrowth,omitempty"`
	DividendHistMeanSTD        map[string][]interface{} `json:"dividendhistmeanstd,omitempty"`
	DividendHistGrowthMeanSTD  map[string][]interface{} `json:"dividendhistgrowthmeanstd,omitempty"`
	DividendYieldTTM           float64                  `json:"dividendyieldttm,omitempty"`
	CurrentVolume              float64                  `json:"currentvolume,omitempty"`
	YearHigh                   float64                  `json:"yearhigh,omitempty"`
	YearLow                    float64                  `json:"yearlow,omitempty"`
	Ratios                     []map[string]*float64    `json:"ratios,omitempty"`
	RatiosGrowth               map[string][]float64     `json:"ratiosgrowth,omitempty"`
	RatiosMeanSTD              map[string][]interface{} `json:"ratiosmeanstd,omitempty"`
	RatiosGrowthMeanSTD        map[string][]interface{} `json:"ratiosgrowthmeanstd,omitempty"`
}

type FinalNumbers struct {
	CalculationsOutlookFundamentals FundamentalsCalculationsResults `json:"calculationsoutlookfundamentals,omitempty"`
	FMPDCF                          []objects.DiscountedCashFlow    `json:"fmpdcf,omitempty"`
	FMPDCFMeanSTD                   map[string][]interface{}        `json:"fmpdcfmeanstd,omitempty"`
	FMPMeanSTDDCF                   map[string][]interface{}        `json:"fmpmeanstddcf,omitempty"`
	EmployeeCount                   float64                         `json:"employeecount,omitempty"`
	FMPRatings                      []map[string]*float64           `json:"fmpratings,omitempty"`
	FMPRatingsGrowth                map[string][]float64            `json:"fmpratingsgrowth,omitempty"`
	FMPRatingsMeanSTD               map[string][]interface{}        `json:"fmpratingsmeanstd,omitempty"`
	FMPRatingsGrowthMeanSTD         map[string][]interface{}        `json:"fmpratingsgrowthmeanstd,omitempty"`
}

type CompanyData struct {
	Ticker       string
	CandleSticks []objects.StockCandle
	Data         FinalNumbers
}
