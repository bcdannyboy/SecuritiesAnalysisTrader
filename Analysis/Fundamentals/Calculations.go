package Fundamentals

import "math"

func CurrentRatio(CurrentAssets float64, CurrentLiabilities float64) float64 {
	// The current ratio is a liquidity ratio that measures a company’s ability to pay short-term obligations or those due within one year. It tells investors and analysts how a company can maximize the current assets on its balance sheet to satisfy its current debt and other payables.

	return CurrentAssets / CurrentLiabilities
}

func AcidTestRatio(CurrentAssets float64, Inventory float64, CurrentLiabilities float64) float64 {
	// The acid-test ratio is a strong indicator of whether a firm has sufficient short-term assets to cover its immediate liabilities. It is calculated as follows:
	// Acid-test ratio = (Cash and Cash Equivalents + Marketable Securities + Accounts Receivable) / Current Liabilities

	return (CurrentAssets - Inventory) / CurrentLiabilities
}

func DebtToEquityRatio(TotalLiabilities float64, TotalShareholderEquity float64) float64 {
	// The debt-to-equity (D/E) ratio is calculated by dividing a company’s total liabilities by its shareholder equity. These numbers are available on the balance sheet of a company’s financial statements. The ratio is used to evaluate a company’s financial leverage.

	return TotalLiabilities / TotalShareholderEquity
}

func ReturnOnEquity(NetIncome float64, TotalShareholderEquity float64) float64 {
	// Return on equity (ROE) is a measure of financial performance calculated by dividing net income by shareholders' equity. Because shareholders' equity is equal to a company’s assets minus its debt, ROE is considered the return on net assets.

	return NetIncome / TotalShareholderEquity
}

func AssetTurnoverRatio(NetSales float64, TotalAssets float64) float64 {
	// Asset turnover ratio is the ratio between the value of a company’s sales or revenues and the value of its assets. The asset turnover ratio measures the ability of a company to use its assets to efficiently generate sales.

	return NetSales / TotalAssets
}

func InventoryTurnoverRatio(CostOfGoodsSold float64, AverageInventory float64) float64 {
	// The inventory turnover ratio is a key measure for evaluating how effective a company's management is at managing inventory levels and generating sales from it. The formula for calculating inventory turnover equals the cost of goods sold divided by the average inventory.

	return CostOfGoodsSold / AverageInventory
}

func AccountsReceivableTurnoverRatio(NetCreditSales float64, AverageAccountsReceivable float64) float64 {
	// The accounts receivable turnover ratio measures a companies effectiveness in terms of qualifying their credit borrowers and collecting monies owed from them. The A/R turnover ratio is an indication to how many times the accounts receivables are "turned over" throughout the year

	return NetCreditSales / AverageAccountsReceivable
}

func FixedAssetTurnoverRatio(NetSales float64, NetFixedAssets float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	return NetSales / NetFixedAssets
}

func DebtToCapitalRatio(ShortTermDebt float64, LongTermDebt float64, TotalShareholderEquity float64) float64 {
	// The debt-to-capital ratio is a measurement of a company's financial leverage. The debt-to-capital ratio is calculated by taking the company's debt, including both short- and long-term liabilities and dividing it by the total capital.

	return (ShortTermDebt + LongTermDebt) / TotalShareholderEquity
}

func InterestCoverageRatio(EBIT float64, InterestExpense float64) float64 {
	// The interest coverage ratio is used to determine how easily a company can pay their interest expenses on outstanding debt. The ratio is calculated by dividing a company's earnings before interest and taxes (EBIT) by the company's interest expenses for the same period.

	return EBIT / InterestExpense
}

func CashConversionCycle(DaysInventoryOutstanding float64, DaysSalesOutstanding float64, DaysPayablesOutstanding float64) float64 {
	// The cash conversion cycle (CCC) is a metric that expresses the length of time, in days, that it takes for a company to convert resource inputs into cash flows. The cash conversion cycle attempts to measure the amount of time each net input dollar is tied up in the production and sales process before it is converted into cash through sales to customers.

	return DaysInventoryOutstanding + DaysSalesOutstanding - DaysPayablesOutstanding
}

func NetWorkingCapital(AccountsReceivable float64, Inventory float64, AccountsPayable float64) float64 {
	// Net working capital is calculated as current assets minus current liabilities. It is a derivation of working capital, that is commonly used in valuation techniques such as discounted cash flows (DCF) that look at cash flows over the next five to ten years.

	return AccountsReceivable + Inventory - AccountsPayable
}

func PPETurnoverRatio(NetSales float64, NetFixedAssets float64, AccumulatedDepreciation float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	return NetSales / (NetFixedAssets - AccumulatedDepreciation)
}

func DebtServiceCoverageRatio(NetOperatingIncome float64, TotalDebtService float64) float64 {
	// The debt service coverage ratio (DSCR) measures the ability of a company to use its operating income to repay all its debt obligations, including repayment of principal and interest on both short-term and long-term debt.

	return NetOperatingIncome / TotalDebtService
}

func TangibleNetWorth(TotalAssets float64, TotalLiabilities float64, IntangibleAssets float64) float64 {
	// Tangible net worth is most commonly a calculation of the net worth of a company that excludes any value derived from intangible assets such as copyrights.

	return TotalAssets - TotalLiabilities - IntangibleAssets
}

func TangibleNetWorthRatio(TNW float64, TotalAssets float64) float64 {
	// Tangible net worth is most commonly a calculation of the net worth of a company that excludes any value derived from intangible assets such as copyrights.

	return TNW / TotalAssets
}

func TimeInterestEarnedRatio(EBIT float64, InterestExpense float64) float64 {
	// The interest coverage ratio is used to determine how easily a company can pay their interest expenses on outstanding debt. The ratio is calculated by dividing a company's earnings before interest and taxes (EBIT) by the company's interest expenses for the same period.

	return EBIT / InterestExpense
}

func ReturnOnAssetRatio(NetIncome float64, TotalAssets float64) float64 {
	// The return on assets ratio, often called the return on total assets, is a profitability ratio that measures the net income produced by total assets during a period by comparing net income to the average total assets.

	return NetIncome / TotalAssets
}

func LeverageRatio(TotalDebt float64, EBITDA float64) float64 {
	// The leverage ratio is a measure of the financial leverage of a company. It is calculated by dividing total debt (the sum of current liabilities and long-term liabilities) by total assets (the sum of current assets, fixed assets, and other assets such as 'goodwill').

	return TotalDebt / EBITDA
}

func CapitalizationRatio(TotalDebt float64, ShareHolderEquity float64) float64 {
	// The capitalization ratio, often called the Cap ratio, is a financial metric that measures a company's solvency by calculating the total debt component of the company's capital structure of debt and equity.

	return TotalDebt / ShareHolderEquity
}

func EquityMultiplierRatio(TotalAssets float64, ShareHolderEquity float64) float64 {
	// The equity multiplier is a financial leverage ratio that measures the portion of company’s assets that are financed by stockholder's equity. It is calculated by dividing total assets by total equity.

	return TotalAssets / ShareHolderEquity
}

func NonPerformingAssetRatio(NonPerformingAssets float64, TotalAssets float64) float64 {
	// The nonperforming asset ratio is a measurement of the percentage of nonperforming assets to the total assets of a bank or company. A nonperforming asset refers to loans or advances that are in jeopardy of default.

	return NonPerformingAssets / TotalAssets
}

func DeferredTaxLiabilityToEquityRatio(DeferredTaxLiabilities float64, ShareHolderEquity float64) float64 {
	// Deferred tax liability is a tax that is assessed or is due for the current period but has not yet been paid. The deferral arises because of timing differences between the accrual of the tax and payment of the tax.

	return DeferredTaxLiabilities / ShareHolderEquity
}

func CashRatio(CashAndCashEquivalents float64, CurrentLiabilities float64) float64 {
	// The cash ratio is a measurement of a company's liquidity, specifically the ratio of a company's total cash and cash equivalents to its current liabilities.

	return CashAndCashEquivalents / CurrentLiabilities
}

func GrossProfitMargin(GrossProfit float64, NetSales float64) float64 {
	// Gross profit margin is a ratio that indicates the performance of a company's sales and production. It measures how much out of every dollar of sales a company actually keeps in earnings.

	return GrossProfit / NetSales
}

func OperatingProfitMargin(OperatingIncome float64, NetSales float64) float64 {
	// Operating margin is a profitability ratio measuring revenue after covering operating and non-operating expenses of a business. It is calculated by dividing operating profit by revenue.

	return OperatingIncome / NetSales
}

func NetProfitMargin(NetIncome float64, NetSales float64) float64 {
	// Net profit margin is equal to how much net income or profit is generated as a percentage of revenue. Net profit margin is the ratio of net profits to revenues for a company or business segment.

	return NetIncome / NetSales
}

func AverageCollectionPeriod(AccountsReceivableTurnoverRatio float64) float64 {
	// The average collection period is the average number of days in a year that a company takes to collect payments owed, after a sale of its goods or services has been made.

	return 365 / AccountsReceivableTurnoverRatio
}

func AccountsPayableTurnoverRatio(SupplierPurchases float64, AverageAccountsPayable float64) float64 {
	// The accounts payable turnover ratio is a short-term liquidity measure used to quantify the rate at which a company pays off its suppliers. Accounts payable turnover ratio is calculated by taking the total purchases made from suppliers and dividing it by the average accounts payable amount during the same period.

	return SupplierPurchases / AverageAccountsPayable
}

func AverageAccountsPayablePaymentPeriod(AccountsPayableTurnoverRatio float64) float64 {
	// The average payable period is the average number of days it takes a company to pay off credit accounts payable.

	return 365 / AccountsPayableTurnoverRatio
}

func InventoryToWorkingCapitalRatio(Inventory float64, WorkingCapital float64) float64 {
	// The inventory to working capital ratio is a liquidity ratio that measures a company's ability to pay off its current liabilities with inventory.

	return Inventory / WorkingCapital
}

func BookValuePerShare(ShareHolderEquity float64, SharesOutstanding float64) float64 {
	// Book value per share (BVPS) takes the ratio of a firm's common equity divided by its number of shares outstanding. Book value of equity per share effectively indicates a firm's net asset value (total assets - total liabilities) on a per-share basis.

	return ShareHolderEquity / SharesOutstanding
}

func NetTangibleAssetsPerShare(TangibleNetWorth float64, SharesOutstanding float64) float64 {
	// Net tangible assets per share is a company's total tangible assets minus total liabilities, divided by its number of shares of common stock outstanding.

	return TangibleNetWorth / SharesOutstanding
}

func DefensiveIntervalRatio(CashAndCashEquivalents, AccountsReceivable, MarketableSecurities, OperatingExpenses, NonCashCharges, PeriodInDays float64) float64 {
	// The defensive interval ratio (DIR) is a financial liquidity ratio that indicates how many days a company can operate without needing to tap into capital sources other than its current assets. It is also known as the basic defense interval ratio (BDIR) or the defensive interval period ratio (DIPR).

	return (CashAndCashEquivalents + AccountsReceivable + MarketableSecurities) / ((OperatingExpenses - NonCashCharges) / PeriodInDays)
}

func TotalDebtToEBITDA(TotalDebt float64, EBITDA float64) float64 {
	// The leverage ratio is a measure of the financial leverage of a company. It is calculated by dividing total debt (the sum of current liabilities and long-term liabilities) by total assets (the sum of current assets, fixed assets, and other assets such as 'goodwill').

	return TotalDebt / EBITDA
}

func TangibleEquityRatio(CommonShareHolderEquity, IntangibleAssets, TotalAssets float64) float64 {
	// The tangible common equity (TCE) ratio measures the percentage of a company’s common stock that is tangible common equity. The ratio is used to calculate a bank's ability to deal with potential losses. The higher the ratio, the more likely it is that the bank will be able to absorb the losses it incurs.

	return (CommonShareHolderEquity - IntangibleAssets) / (TotalAssets - IntangibleAssets)
}

func IntangiblesRatio(IntangibleAssets, TotalAssets float64) float64 {
	// The intangible ratio is a financial ratio that measures the percentage of intangible assets in comparison to a company's total assets. The intangible ratio is calculated by dividing intangible assets by total assets.

	return IntangibleAssets / TotalAssets
}

func InvestmentTurnoverRatio(NetSales, TotalInvestments float64) float64 {
	// The investment turnover ratio measures a company's efficiency in using its assets to generate revenue or sales. The ratio formula is to divide net sales by the average total assets.

	return NetSales / TotalInvestments
}

func ReturnOnCapitalEmployed(EBIT, TotalAssets, TotalLiabilities float64) float64 {
	// Return on capital employed (ROCE) is a financial ratio that measures a company's profitability and the efficiency with which its capital is employed. ROCE is calculated as: ROCE = EBIT / Capital Employed.

	return EBIT / (TotalAssets - TotalLiabilities)
}

func WorkingCapitalTurnoverRatio(NetSales, WorkingCapital float64) float64 {
	// The working capital turnover ratio measures how efficiently a company is using its working capital to support a given level of sales. It is calculated by dividing net sales by the average working capital.

	return NetSales / WorkingCapital
}

func NetGearingRatio(LongTermDebt, ShortTermDebt, AdditionalLiabilities, ShareHoldersEquity float64) float64 {
	// The net gearing ratio is a liquidity ratio that measures a company's ability to repay its debts if they were all due today. It is calculated by dividing a company's net debt by its net assets.

	return (LongTermDebt + ShortTermDebt + AdditionalLiabilities) / ShareHoldersEquity
}

func FixedChargeCoverageRatio(EBIT, LeasePayments, InterestExpense float64) float64 {
	// The fixed-charge coverage ratio (FCCR) is a measure of a firm's ability to cover its fixed charges, such as debt payments, interest expense, and equipment lease expense. The ratio is calculated by dividing a company's earnings before interest and taxes (EBIT) by its total fixed charges.

	return (EBIT + LeasePayments) / (LeasePayments + (LeasePayments * InterestExpense))
}

func EBITDAMarginRatio(EBITDA, NetSales float64) float64 {
	// EBITDA margin is a profitability ratio that measures how much earnings a company is generating before interest, taxes, depreciation, and amortization, as a percentage of revenue.

	return EBITDA / NetSales
}

func DividendPayoutRatio(Dividends, NetIncome float64) float64 {
	// The dividend payout ratio is the ratio of the total amount of dividends paid out to shareholders relative to the net income of the company. It is the percentage of earnings paid to shareholders in dividends.

	return Dividends / NetIncome
}

func PriceToBookValueRatio(MarketCapitalization float64, BookValueOfEquity float64) float64 {
	// The price-to-book ratio compares a company's market value to its book value. The market value of a company is its share price multiplied by the number of outstanding shares. The book value is the net assets of a company.

	return MarketCapitalization / BookValueOfEquity
}

func WeightedAverageCostOfCapital(MarketValueOfEquity, MarketValueOfDebt, CostOfEquity, CostOfDebt, CorporateTaxRate float64) float64 {
	// Weighted average cost of capital (WACC) represents a company's cost of capital, with each category of capital (debt and equity) proportionately weighted.
	E := MarketValueOfEquity
	D := MarketValueOfDebt
	V := E + D

	return ((E / V) * CostOfEquity) + ((D / V) * CostOfDebt * (1 - CorporateTaxRate))
}

func CostOfEquity(RiskFreeRate, MarketReturn, Beta float64) float64 {
	// The cost of equity is the return a company requires to decide if an investment meets capital return requirements. Firms often use it as a capital budgeting threshold for the required rate of return.

	return RiskFreeRate + (Beta * (MarketReturn - RiskFreeRate))
}

func CostOfDebt(InterestExpense, MarketValueOfDebt float64) float64 {
	// The cost of debt is the effective interest rate a company pays on its debt obligations, including bonds, mortgages, and any other forms of debt the company may have.

	return InterestExpense / MarketValueOfDebt
}

func MarketValueOfEquity(SharePrice, SharesOutstanding float64) float64 {
	// Market value of equity is the total dollar value of a company's equity calculated by multiplying the current stock price by total outstanding shares.

	return SharePrice * SharesOutstanding
}

func MarketValueOfDebt(SharePrice, SharesOutstanding, BookValueOfDebt float64) float64 {
	// Market value of debt is the total dollar value of a company's debt, calculated by multiplying the current market price of a company's debt by its total outstanding debt.

	return (SharePrice * SharesOutstanding) + BookValueOfDebt
}

func EconomicValueAdded(NOPAT, WACC, TotalCapital float64) float64 {
	// Economic value added (EVA) is a measure of a company's financial performance based on the residual wealth calculated by deducting its cost of capital from its operating profit, adjusted for taxes on a cash basis.

	return NOPAT - (WACC * TotalCapital)
}

func MarketToBookRatio(MarketCapitalization, BookValueOfEquity float64) float64 {
	// The market-to-book (M/B) ratio is calculated by dividing the market price per share by book value per share.

	return MarketCapitalization / BookValueOfEquity
}

func RetentionRate(Dividends, NetIncome float64) float64 {
	// The retention ratio is the proportion of earnings kept back in the business as retained earnings. Retention ratio refers to the percentage of net income that is retained to grow the business, rather than being paid out as dividends.

	return (NetIncome - Dividends) / NetIncome
}

func SustainableGrowthRate(RetentionRatio, ReturnOnEquity float64) float64 {
	// The sustainable growth rate (SGR) is the maximum rate of growth that a company can sustain without raising additional equity or taking on new debt. The sustainable growth rate can be calculated by multiplying a company's earnings retention ratio by its return on equity (ROE).

	return RetentionRatio * ReturnOnEquity
}

func FreeCashFlowToEquity(EBITDA, DepreciationAmortization, InterestPayments, Taxes, ChangeInWorkingCapital, CapitalExpenditures, NetDebt float64) float64 {
	// Free cash flow to equity (FCFE) is a measure of how much cash can be paid to the equity shareholders of a company after all expenses, reinvestment, and debt are paid. FCFE is a measure of equity capital usage.

	return EBITDA - (CapitalExpenditures - DepreciationAmortization) - ChangeInWorkingCapital + (NetDebt - InterestPayments) - Taxes
}

func FreeCashFlowToFirm(NetIncome, NonCashCharges, InterestPayments, TaxRate, LongTermInvestments, InvestmentsInWorkingCapital float64) float64 {
	// Free cash flow to the firm (FCFF) represents the amount of cash flow from operations available for distribution after certain expenses are paid. FCF is a measure of a company's profitability after all expenses and reinvestments.

	return NetIncome + NonCashCharges - (InterestPayments * (1 - TaxRate)) - LongTermInvestments - InvestmentsInWorkingCapital
}

func CapitalExpenditureRatio(CapitalExpenditures, OperatingCashFlow float64) float64 {
	// The capital expenditure ratio is a ratio that measures the amount of a company's capital expenditures to the amount of its operating cash flow.

	return CapitalExpenditures / OperatingCashFlow
}

func OperatingCashFlowRatio(OperatingCashFlow, NetSales float64) float64 {
	// The operating cash flow ratio is a measure of how well current liabilities are covered by the cash flows generated from a company's operations.

	return OperatingCashFlow / NetSales
}

func EarningsPerShare(NetIncome, PreferredDividends, SharesOutstanding float64) float64 {
	// Earnings per share (EPS) is calculated as a company's profit divided by the outstanding shares of its common stock. The resulting number serves as an indicator of a company's profitability.

	return (NetIncome - PreferredDividends) / SharesOutstanding
}

func EBITDAPerShare(EBITDA, SharesOutstanding float64) float64 {
	// EBITDA per share is a ratio used to measure a company's return on investment. It is calculated by dividing EBITDA by the number of outstanding shares.

	return EBITDA / SharesOutstanding
}

func GrossMarginOnInventory(GrossProfit, Inventory float64) float64 {
	// Gross margin on inventory is a ratio that measures how profitable a company is relative to its inventory. It is calculated by dividing gross profit by inventory.

	return GrossProfit / Inventory
}

func EconomicOrderQuantity(CarryingCostPerUnit, OrderingCostPerOrder, AnnualDemand float64) float64 {
	// Economic order quantity (EOQ) is an equation for inventory that determines the ideal order quantity a company should purchase for its inventory given a set cost of production, demand rate, and other variables.

	return math.Sqrt((2 * AnnualDemand * OrderingCostPerOrder) / CarryingCostPerUnit)
}

func CashFlowReturnOnEquity(OperatingCashFlow, ShareHolderEquity float64) float64 {
	// Cash flow return on equity (CFROE) is a measure of a company's ability to generate cash flow from equity. It is calculated by dividing cash flow from operations by shareholders' equity.

	return OperatingCashFlow / ShareHolderEquity
}

func EfficiencyRatio(NonInterestExpense, NetRevenue float64) float64 {
	// The efficiency ratio is typically used to analyze how well a company uses its assets and liabilities internally. An efficiency ratio can calculate the turnover of receivables, the repayment of liabilities, the quantity and usage of equity, and the general use of inventory and machinery.

	return NonInterestExpense / NetRevenue
}

func LoanToDepositRatio(Loans, Deposits float64) float64 {
	// The loan-to-deposit ratio (LDR) is used to assess a bank's liquidity by comparing a bank's total loans to its total deposits for the same period. The LDR is expressed as a percentage.

	return Loans / Deposits
}

func RevenuePerEmployee(NetRevenue, NumberOfEmployees float64) float64 {
	// Revenue per employee is an important ratio that looks at a company's revenue in relation to the number of employees it has. The ratio is generally used as a way of measuring productivity.

	return NetRevenue / NumberOfEmployees
}

func Economicprofit(TotalRevenue, ExplicitCost, ImplicitCost float64) float64 {
	// Economic profit is the difference between the revenue received from the sale of an output and the costs of all inputs, including opportunity costs.

	return TotalRevenue - (ExplicitCost + ImplicitCost)
}
