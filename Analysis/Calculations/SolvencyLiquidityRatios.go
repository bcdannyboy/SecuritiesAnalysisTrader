package Calculations

import "fmt"

func CurrentRatio(CurrentAssets float64, CurrentLiabilities float64) float64 {
	// The current ratio is a liquidity ratio that measures a company’s ability to pay short-term obligations or those due within one year. It tells investors and analysts how a company can maximize the current assets on its balance sheet to satisfy its current debt and other payables.
	if CurrentLiabilities == 0 {
		fmt.Printf("got a zero value for CurrentLiabilities, with CurrentAssets: %f and CurrentLiabilities: %f\n", CurrentAssets, CurrentLiabilities)
		return 0
	}
	return CurrentAssets / CurrentLiabilities
}

func AcidTestRatio(CurrentAssets float64, Inventory float64, CurrentLiabilities float64) float64 {
	// The acid-test ratio is a strong indicator of whether a firm has sufficient short-term assets to cover its immediate liabilities. It is calculated as follows:
	// Acid-test ratio = (Cash and Cash Equivalents + Marketable Securities + Accounts Receivable) / Current Liabilities
	if CurrentLiabilities == 0 {
		fmt.Printf("got a zero value for CurrentLiabilities, with CurrentAssets: %f and CurrentLiabilities: %f\n", CurrentAssets, CurrentLiabilities)
		return 0
	}
	return (CurrentAssets - Inventory) / CurrentLiabilities
}

func CashRatio(CashAndCashEquivalents float64, CurrentLiabilities float64) float64 {
	// The cash ratio is a measurement of a company's liquidity, specifically the ratio of a company's total cash and cash equivalents to its current liabilities.
	if CurrentLiabilities == 0 {
		fmt.Printf("got a zero value for CurrentLiabilities, with CashAndCashEquivalents: %f and CurrentLiabilities: %f\n", CashAndCashEquivalents, CurrentLiabilities)
		return 0
	}
	return CashAndCashEquivalents / CurrentLiabilities
}

func DefensiveIntervalRatio(CashAndCashEquivalents, AccountsReceivable, MarketableSecurities, OperatingExpenses, NonCashCharges, PeriodInDays float64) float64 {
	// The defensive interval ratio (DIR) is a financial liquidity ratio that indicates how many days a company can operate without needing to tap into capital sources other than its current assets. It is also known as the basic defense interval ratio (BDIR) or the defensive interval period ratio (DIPR).
	if PeriodInDays == 0 {
		fmt.Printf("got a zero value for PeriodInDays, with CashAndCashEquivalents: %f, AccountsReceivable: %f, MarketableSecurities: %f, OperatingExpenses: %f, NonCashCharges: %f, PeriodInDays: %f\n", CashAndCashEquivalents, AccountsReceivable, MarketableSecurities, OperatingExpenses, NonCashCharges, PeriodInDays)
		return 0
	}
	if ((OperatingExpenses - NonCashCharges) / PeriodInDays) == 0 {
		fmt.Printf("got a zero value for ((OperatingExpenses - NonCashCharges) / PeriodInDays), with CashAndCashEquivalents: %f, AccountsReceivable: %f, MarketableSecurities: %f, OperatingExpenses: %f, NonCashCharges: %f, PeriodInDays: %f\n", CashAndCashEquivalents, AccountsReceivable, MarketableSecurities, OperatingExpenses, NonCashCharges, PeriodInDays)
		return 0
	}
	return (CashAndCashEquivalents + AccountsReceivable + MarketableSecurities) / ((OperatingExpenses - NonCashCharges) / PeriodInDays)
}

func DrySalesRatio(AccountsReceivable, NetRevenue float64) float64 {
	// The dry sales ratio is a liquidity ratio that measures a company’s ability to cover its short-term obligations with its liquid assets. It is calculated by dividing a company’s liquid assets by its net credit sales.
	if NetRevenue == 0 {
		fmt.Printf("got a zero value for NetRevenue, with AccountsReceivable: %f and NetRevenue: %f\n", AccountsReceivable, NetRevenue)
		return 0
	}
	return AccountsReceivable / (NetRevenue / 365)
}
