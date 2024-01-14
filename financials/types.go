package financials

/*
 * The FinnWorld /balancesheet response object.
 * https://finnworlds.com/documentation/#balancesheets
 */
type BalanceSheetResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			Ticker string `json:"ticker"`
			Period string `json:"period"`
		} `json:"basics"`
		Output struct {
			Balance_Sheet struct {
				Date                                                string `json:"date"`
				Accounts_Payable                                    string `json:"accounts_payable"`
				Accounts_Receivable                                 string `json:"accounts_receivable"`
				Accumulated_Depreciation                            string `json:"accumulated_depreciation"`
				Additional_Paid_In_Capital                          string `json:"additional_paid_in_capital"`
				Capital_Lease_Obligations                           string `json:"capital_lease_obligations"`
				Capital_Stock                                       string `json:"capital_stock"`
				Cash_And_Cash_Equivalents                           string `json:"cash_and_cash_equivalents"`
				Cash_Cash_Equivalents_And_Short_Term_Investments    string `json:"cash_cash_equivalents_and_short_term_investments"`
				Common_Stock                                        string `json:"common_stock"`
				Common_Stock_Equity                                 string `json:"common_stock_equity"`
				Construction_in_Progress                            string `json:"construction_in_progress"`
				Current_Accrued_Expenses                            string `json:"current_accrued_expenses"`
				Current_Assets                                      string `json:"current_assets"`
				Current_Capital_Lease_Obligation                    string `json:"current_capital_lease_obligation"`
				Current_Debt                                        string `json:"current_debt"`
				Current_Debt_and_Capital_Lease_Obligation           string `json:"current_debt_and_capital_lease_obligation"`
				Current_Deferred_Liabilities                        string `json:"current_deferred_liabilities"`
				Current_Deferred_Revenue                            string `json:"current_deferred_revenue"`
				Current_Liabilities                                 string `json:"current_liabilities"`
				Current_Provisions                                  string `json:"current_provisions"`
				Finished_Goods                                      string `json:"finished_goods"`
				Gains_Losses_Not_Affecting_Retained_Earnings        string `json:"gains_losses_not_affecting_retained_earnings"`
				Goodwill                                            string `json:"goodwill"`
				Goodwill_and_Other_Intangible_Assets                string `json:"goodwill_and_other_intangible_assets"`
				Gross_PPE                                           string `json:"gross_ppe"`
				Inventory                                           string `json:"inventory"`
				Invested_Capital                                    string `json:"invested_capital"`
				Land_and_Imporvements                               string `json:"land_and_imporvements"`
				Leases                                              string `json:"leases"`
				Long_Term_Capital_Lease_Obligation                  string `json:"long_term_capital_lease_obligation"`
				Long_Term_Debt                                      string `json:"long_term_debt"`
				Long_Term_Debt_and_Capital_Lease_Obligation         string `json:"long_term_debt_and_capital_lease_obligation"`
				Long_Term_Provisions                                string `json:"long_term_provisions"`
				Machinery_Furniture_Equipment                       string `json:"machinery_furniture_equipment"`
				Minority_Interest                                   string `json:"minority_interest"`
				Net_PPE                                             string `json:"net_ppe"`
				Net_Tangible_Assets                                 string `json:"net_tangible_assets"`
				Non_Current_Accrued_Expenses                        string `json:"non_current_accrued_expenses"`
				Non_Current_Deferred_Liabilities                    string `json:"non_current_deferred_liabilities"`
				Non_Current_Deferred_Revenue                        string `json:"non_current_deferred_revenue"`
				Non_Current_Deferred_Taxes_Liabilities              string `json:"non_current_deferred_taxes_liabilities"`
				Ordinary_Shares_Number                              string `json:"ordinary_shares_number"`
				Other_Current_Assets                                string `json:"other_current_assets"`
				Other_Current_Borrowings                            string `json:"other_current_borrowings"`
				Other_Current_Liabilities                           string `json:"other_current_liabilities"`
				Other_Intangible_Assets                             string `json:"other_intangible_assets"`
				Other_Inventories                                   string `json:"other_inventories"`
				Other_Non_Current_Assets                            string `json:"other_non_current_assets"`
				Other_Non_Current_Liabilities                       string `json:"other_non_current_liabilities"`
				Other_Properties                                    string `json:"other_properties"`
				Other_Short_Term_Investments                        string `json:"other_short_term_investments"`
				Payables                                            string `json:"payables"`
				Payables_And_Accrued_Expenses                       string `json:"payables_and_accrued_expenses"`
				Preferred_Stock                                     string `json:"preferred_stock"`
				Properties                                          string `json:"properties"`
				Raw_Materials                                       string `json:"raw_materials"`
				Receivables                                         string `json:"receivables"`
				Retained_Earnings                                   string `json:"retained_earnings"`
				Share_Issued                                        string `json:"share_issued"`
				StockHolders_Equity                                 string `json:"stockholders_equity"`
				Tangible_Book_Value                                 string `json:"tangible_book_value"`
				Total_Assets                                        string `json:"total_assets"`
				Total_Capitalization                                string `json:"total_capitalization"`
				Total_Debt                                          string `json:"total_debt"`
				Total_Equity_Gross_Minority_Interest                string `json:"total_equity_gross_minority_interest"`
				Total_Liabilities_Net_Minority_Interest             string `json:"total_liabilities_net_minority_interest"`
				Total_Non_Current_Assets                            string `json:"total_non_current_assets"`
				Total_Non_Current_Liabilities_Net_Minority_Interest string `json:"total_non_current_liabilities_net_minority_interest"`
				Total_Tax_Payable                                   string `json:"total_tax_payable"`
				Working_Capital                                     string `json:"working_capital"`
				Work_in_Process                                     string `json:"work_in_process"`
			} `json:"balance_sheet"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /incomestatements response object
* https://finnworlds.com/documentation/#incomestatements
 */
type IncomeStatementResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			Ticker string `json:"ticker"`
			Period string `json:"period"`
		} `json:"basics"`
		Output struct {
			Average_Dilution_Earnings                                 string `json:"average_dilution_earnings"`
			Basic_Average_Shares                                      string `json:"basic_average_shares"`
			Basic_EPS                                                 string `json:"basic_eps"`
			Cost_of_Revenue                                           string `json:"cost_of_revenue"`
			Diluted_Average_Shares                                    string `json:"diluted_average_shares"`
			Diluted_EPS                                               string `json:"diluted_eps"`
			Diluted_Niavailto_Com_Shareholders                        string `json:"diluted_niavailto_com_shareholders"`
			EBIT                                                      string `json:"ebit"`
			Gross_Profit                                              string `json:"gross_profit"`
			Interest_Expense                                          string `json:"interest_expense"`
			Interest_Expense_Non_Operating                            string `json:"interest_expense_non_operating"`
			Interest_Income                                           string `json:"interest_income"`
			Interest_Income_Non_Operating                             string `json:"interest_income_non_operating"`
			Minority_Interests                                        string `json:"minority_interests"`
			Net_Income                                                string `json:"net_income"`
			Net_Income_Common_Stockholders                            string `json:"net_income_common_stockholders"`
			Net_Income_Continuous_Operations                          string `json:"net_income_continuous_operations"`
			Net_Income_From_Continuing_And_Discontinued_Operations    string `json:"net_income_from_continuing_and_discontinued_operations"`
			Net_Income_From_Continuing_Operation_Net_Minority_Interst string `json:"net_income_from_continuing_operation_net_minority_interst"`
			Net_Income_Including_NonControlling_Interests             string `json:"net_income_including_non_controlling_interests"`
			Net_Interest_Income                                       string `json:"net_interest_income"`
			Net_Non_Operating_Interest_Incone_Expense                 string `json:"net_non_operating_interest_incone_expense"`
			Normalized_EBITDA                                         string `json:"normalized_ebitda"`
			Normalized_Income                                         string `json:"normalized_income"`
			Operating_Income                                          string `json:"operating_income"`
			Operating_Revenue                                         string `json:"operating_revenue"`
			Other_Income_Expense                                      string `json:"other_income_expense"`
			Other_Non_Operating_Income_Expenses                       string `json:"other_non_operating_income_expenses"`
			Pretax_Income                                             string `json:"pretax_income"`
			Reconciled_Cost_of_Revenue                                string `json:"reconciled_cost_of_revenue"`
			Reconciled_Depreciation                                   string `json:"reconciled_depreciation"`
			Research_And_Development                                  string `json:"research_and_development"`
			Restrucuring_and_Mergern_Acquisition                      string `json:"restrucuring_and_mergern_acquisition"`
			Selling_General_and_Administration                        string `json:"selling_general_and_administration"`
			Special_Income_Charges                                    string `json:"special_income_charges"`
			Tax_Effect_Of_Unusual_Items                               string `json:"tax_effect_of_unusual_items"`
			Tax_Provision                                             string `json:"tax_provision"`
			Tax_Rate_For_Calcs                                        string `json:"tax_rate_for_calcs"`
			Total_Expenses                                            string `json:"total_expenses"`
			Total_Operating_Income_As_Reported                        string `json:"total_operating_income_as_reported"`
			Total_Revenue                                             string `json:"total_revenue"`
			Total_Unusual_Items                                       string `json:"total_unusual_items"`
			Total_Unusual_Items_Excluding_Goodwill                    string `json:"total_unusual_items_excluding_goodwill"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /cashflow response object
* https://finnworlds.com/documentation/#cashflow
 */
type CashFlowResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			Ticker string `json:"ticker"`
			Period string `json:"period"`
		} `json:"basics"`
		Output struct {
			Cash_Flow []struct {
				Date                                           string `json:"date"`
				Asset_Impairment_Charge                        string `json:"asset_impairment_charge"`
				Beginning_Cash_Position                        string `json:"beginning_cash_position"`
				Capital_Expenditure                            string `json:"capital_expenditure"`
				Capital_Expenditure_Reported                   string `json:"capital_expenditure_reported"`
				Cash_Flow_From_Continuing_Financial_Activities string `json:"cash_flow_from_continuing_financial_activities"`
				Cash_Flow_From_Continuing_Investing_Activities string `json:"cash_flow_from_continuing_investing_activities"`
				Cash_Flow_From_Continuing_Operating_Activities string `json:"cash_flow_from_continuing_operating_activities"`
				Change_In_Inventory                            string `json:"change_in_inventory"`
				Change_In_Other_Current_Assets                 string `json:"change_in_other_current_assets"`
				Change_In_Other_Current_Liabilities            string `json:"change_in_other_current_liabilities"`
				Change_In_Other_Working_Capital                string `json:"change_in_other_working_capital"`
				Change_In_Payables_And_Accrued_Expense         string `json:"change_in_payables_and_accrued_expense"`
				Change_In_Prepaid_Assets                       string `json:"change_in_prepaid_assets"`
				Change_In_Receivables                          string `json:"change_in_receivables"`
				Change_In_Working_Capital                      string `json:"change_in_working_capital"`
				Change_In_Account_Receivables                  string `json:"change_in_account_receivables"`
				Change_In_Cash                                 string `json:"change_in_cash"`
				Depreciation                                   string `json:"depreciation"`
				Depreciation_Amortization_Depletion            string `json:"depreciation_amortization_depletion"`
				Depreciation_And_Amortization                  string `json:"depreciation_and_amortization"`
				Effect_Of_Exchange_Rate_Changes                string `json:"effect_of_exchange_rate_changes"`
				End_Cash_Position                              string `json:"end_cash_position"`
				Financing_Cash_Flow                            string `json:"financing_cash_flow"`
				Free_Cash_Flow                                 string `json:"free_cash_flow"`
				Gain_Loss_On_Scale_Of_PPE                      string `json:"gain_loss_on_scale_of_ppe"`
				Investing_Cash_Flow                            string `json:"investing_cash_flow"`
				Issuance_Of_Debt                               string `json:"issuance_of_debt"`
				Long_Term_Debt_Issuance                        string `json:"long_term_debt_issuance"`
				Long_Term_Debt_Payments                        string `json:"long_term_debt_payments"`
				Net_Foreign_Currency_Exchange_Gain_Loss        string `json:"net_foreign_currency_exchange_gain_loss"`
				Net_Income_From_Continuing_Operations          string `json:"net_income_from_continuing_operations"`
				Net_Intangible_Purchases_And_Sale              string `json:"net_intangible_purchases_and_sale"`
				Net_Investment_Purchase_And_Sale               string `json:"net_investment_purchase_and_sale"`
				Net_Issuance_Payments_Of_Debt                  string `json:"net_issuance_payments_of_debt"`
				Net_Long_Term_Debt_Issuance                    string `json:"net_long_term_debt_issuance"`
				Net_Other_Financing_Charges                    string `json:"net_other_financing_charges"`
				Net_Other_Investing_Changes                    string `json:"net_other_investing_changes"`
				Net_PPEPurchase_And_Sale                       string `json:"net_ppepurchase_and_sale"`
				Operating_Cash_Flow                            string `json:"operating_cash_flow"`
				Operating_Gains_Losses                         string `json:"operating_gains_losses"`
				Other_Non_Cash_Items                           string `json:"other_non_cash_items"`
				Proceeds_From_Stock_Option_Exercised           string `json:"proceeds_from_stock_option_exercised"`
				Purchase_Of_Intangibles                        string `json:"purchase_of_intangibles"`
				Purchase_Of_Investment                         string `json:"purchase_of_investment"`
				Purchase_Of_PPE                                string `json:"purchase_of_ppe"`
				Repayment_Of_Debt                              string `json:"repayment_of_debt"`
				Sale_Of_Intangibles                            string `json:"sale_of_intangibles"`
				Sale_Of_Investment                             string `json:"sale_of_investment"`
				Stock_Based_Compensation                       string `json:"stock_based_compensation"`
			} `json:"cash_flow"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /dividends response object
* https://finnworlds.com/documentation/#dividends
 */
type DividendsResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Basics struct {
			Ticker string `json:"ticker"`
		} `json:"basics"`
		Output struct {
			Dividends struct {
				Dividend_Rate string `json:"dividend_rate"`
				Date          string `json:"date"`
			} `json:"dividends"`
		} `json:"output"`
	} `json:"data"`
}

/*
* The FinnWorld /stocksplits response object
* https://finnworlds.com/documentation/#earnings
 */
type StockSplitsResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Basics struct {
			Ticker string `json:"ticker"`
		} `json:"basics"`
		Output struct {
			StockSplits struct {
				Stock_Split string `json:"stock_split"`
				Date        string `json:"date"`
			} `json:"stocksplits"`
		} `json:"output"`
	} `json:"data"`
}

/*
* The FinnWorld /historicalcandlestick response object
* https://finnworlds.com/documentation/#historicalcandlestick
 */
type HistoricalCandleStickResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Basics struct {
			Ticker     string `json:"ticker"`
			Date_Start string `json:"date_start"`
			Date_End   string `json:"date_end"`
		} `json:"basics"`
		Output struct {
			Daily_Stock_Data struct {
				Date           string `json:"date"`
				OpenTime       string `json:"opentime"`
				Open           string `json:"open"`
				High           string `json:"high"`
				Low            string `json:"low"`
				Close          string `json:"close"`
				CloseTime      string `json:"closetime"`
				Adjusted_Close string `json:"adjusted_close"`
				Trade_Volume   string `json:"trade_volume"`
				Stock_Split    string `json:"stock_split"`
				Dividend_Rate  string `json:"dividend_rate"`
			} `json:"daily_stock_data"`
		} `json:"output"`
	} `json:"data"`
}

/*
* The FinnWorld /macroindicator response object
* https://finnworlds.com/documentation/#macroindicator
 */
type MacroIndiciatorResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Basics struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Details string `json:"details"`
		} `json:"basics"`
		Output []struct {
			Report_Name string `json:"report_name"`
			Previous    string `json:"previous"`
			Actual      string `json:"actual"`
			Unit        string `json:"unit"`
			Report_Date string `json:"report_date"`
		} `json:"output"`
	} `json:"data"`
}

/*
* The FinnWorld /macrocalendar response object
* https://finnworlds.com/documentation/#macrocalendar
 */
type CalendarResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Basics struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Details string `json:"details"`
		} `json:"basics"`
		Output []struct {
			DateTime         string `json:"datetime"`
			ISO_Country_Code string `json:"iso_country_code"`
			Country          string `json:"country"`
			Report_Name      string `json:"report_name"`
			Report_Date      string `json:"report_date"`
			Actual           string `json:"actual"`
			Previous         string `json:"previous"`
			Consensus        string `json:"consensus"`
			Unit             string `json:"unit"`
			Impact           string `json:"impact"`
		} `json:"output"`
	} `json:"data"`
}

/*
* The FinnWorld /financialratios response object
* https://finnworlds.com/documentation/#financialratios
 */
type FinancialRatiosResponse struct {
	Status struct {
		Message string `json:"message"`
	} `json:"status"`
	Results []struct {
		Basics struct {
			Name              string `json:"name"`
			StockTickerSymbol string `json:"stock_ticker_symbol"`
			IsinIdentifier    string `json:"isin_identifier"`
			Exchange          string `json:"exchange"`
			Industry          string `json:"industry"`
			Sector            string `json:"sector"`
		} `json:"basics"`
		Output struct {
			CurrentRatio                 string `json:"currentRatio"`
			QuickRatio                   string `json:"quickRatio"`
			CashRatio                    string `json:"cashRatio"`
			DaysOfSalesOutstanding       string `json:"daysOfSalesOutstanding"`
			DaysOfInventoryOutstanding   string `json:"daysOfInventoryOutstanding"`
			OperatingCycle               string `json:"operatingCycle"`
			DaysOfPayablesOutstanding    string `json:"daysOfPayablesOutstanding"`
			CashConversionCycle          string `json:"cashConversionCycle"`
			GrossProfitMargin            string `json:"grossProfitMargin"`
			OperatingProfitMargin        string `json:"operatingProfitMargin"`
			PretaxProfitMargin           string `json:"pretaxProfitMargin"`
			NetProfitMargin              string `json:"netProfitMargin"`
			EffectiveTaxRate             string `json:"effectiveTaxRate"`
			ReturnOnAssets               string `json:"returnOnAssets"`
			ReturnOnEquity               string `json:"returnOnEquity"`
			ReturnOnCapitalEmployed      string `json:"returnOnCapitalEmployed"`
			NetIncomePerEBT              string `json:"netIncomePerEBT"`
			EbtPerEbit                   string `json:"ebtPerEbit"`
			EbitPerRevenue               string `json:"ebitPerRevenue"`
			DebtRatio                    string `json:"debtRatio"`
			DebtEquityRatio              string `json:"debtEquityRatio"`
			LongTermDebtToCapitalization string `json:"longTermDebtToCapitalization"`
			TotalDebtToCapitalization    string `json:"totalDebtToCapitalization"`
			InterestCoverage             string `json:"interestCoverage"`
			CashFlowToDebtRatio          string `json:"cashFlowToDebtRatio"`
			CompanyEquityMultiplier      string `json:"companyEquityMultiplier"`
			ReceivablesTurnover          string `json:"receivablesTurnover"`
			PayablesTurnover             string `json:"payablesTurnover"`
			InventoryTurnover            string `json:"inventoryTurnover"`
			FixedAssetTurnover           string `json:"fixedAssetTurnover"`
			AssetTurnover                string `json:"assetTurnover"`
			OperatingCashFlowPerShare    string `json:"operatingCashFlowPerShare"`
			FreeCashFlowPerShare         string `json:"freeCashFlowPerShare"`
		} `json:"output"`
	} `json:"results"`
}

/*
* The FinnWorld /technicalindicators response object
* https://finnworlds.com/documentation/#technicalindicators
 */
type TechnicalIndicatorResponse struct {
	Status struct {
		Message string `json:"message"`
	} `json:"status"`
	Results []struct {
		Basics struct {
			Name              string `json:"name"`
			StockTickerSymbol string `json:"stock_ticker_symbol"`
			IsinIdentifier    string `json:"isin_identifier"`
			Exchange          string `json:"exchange"`
		} `json:"basics"`
		Output []struct {
			Value     float64 `json:"value"`
			Backtrack int     `json:"backtrack"`
		} `json:"output"`
	} `json:"results"`
}

/*
* The FinnWorld /companyratings response object
* https://finnworlds.com/documentation/#companyratings
 */
type CompanyRatingsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			CompanyName string `json:"company_name"`
			Ticker      string `json:"ticker"`
		} `json:"basics"`
		Output struct {
			AnalystConsensus struct {
				AnalystConsensus struct {
					ConsensusConclusion string `json:"consensus_conclusion"`
					AnalystAverage      string `json:"analyst_average"`
					AnalystHighest      string `json:"analyst_highest"`
					AnalystLowest       string `json:"analyst_lowest"`
					AnalystsNumber      string `json:"analysts_number"`
					Buy                 string `json:"buy"`
					Hold                string `json:"hold"`
					Sell                string `json:"sell"`
					ConsensusDate       string `json:"consensus_date"`
				} `json:"analyst_consensus"`
				Analysts struct {
					AnalystName string `json:"analyst_name"`
					AnalystFirm string `json:"analyst_firm"`
					AnalystRole string `json:"analyst_role"`
					Rating      struct {
						DateRating  string `json:"date_rating"`
						TargetDate  string `json:"target_date"`
						PriceTarget string `json:"price_target"`
						Rated       string `json:"rated"`
						Conclusion  string `json:"conclusion"`
					} `json:"rating"`
				} `json:"analysts"`
			} `json:"analyst_consensus"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /analystratings response object
* https://finnworlds.com/documentation/#analystratings
 */
type AnalystRatingsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			AnalystName string `json:"analyst_name"`
			AnalystFirm string `json:"analyst_firm"`
			AnalystRole string `json:"analyst_role"`
		} `json:"basics"`
		Output struct {
			CompanyName string `json:"company_name"`
			Ticker      string `json:"ticker"`
			Rating      struct {
				DateRating  string `json:"date_rating"`
				TargetDate  string `json:"target_date"`
				PriceTarget string `json:"price_target"`
				Rated       string `json:"rated"`
				Conclusion  string `json:"conclusion"`
			} `json:"rating"`
			AnalystConsensus struct {
				ConsensusConclusion string `json:"consensus_conclusion"`
				AnalystAverage      string `json:"analyst_average"`
				AnalystHighest      string `json:"analyst_highest"`
				AnalystLowest       string `json:"analyst_lowest"`
				AnalystsNumber      string `json:"analysts_number"`
				Buy                 string `json:"buy"`
				Hold                string `json:"hold"`
				Sell                string `json:"sell"`
				ConsensusDate       string `json:"consensus_date"`
			} `json:"analyst_consensus"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /consensusratings response object
* https://finnworlds.com/documentation/#consensusratings
 */
type ConsensusRatingsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			CompanyName string `json:"company_name"`
			Ticker      string `json:"ticker"`
		} `json:"basics"`
		Output struct {
			AnalystConsensus struct {
				ConsensusConclusion string `json:"consensus_conclusion"`
				StockPrice          string `json:"stock_price"`
				AnalystAverage      string `json:"analyst_average"`
				AnalystHighest      string `json:"analyst_highest"`
				AnalystLowest       string `json:"analyst_lowest"`
				AnalystsNumber      string `json:"analysts_number"`
				Buy                 string `json:"buy"`
				Hold                string `json:"hold"`
				Sell                string `json:"sell"`
				ConsensusDate       string `json:"consensus_date"`
			} `json:"analyst_consensus"`
			EarningsEstimate struct {
				PreviousQuarter struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					LastYearEps     string `json:"last_year_eps"`
					Date            string `json:"date"`
				} `json:"previous_quarter"`
				CurrentQuarter struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					LastYearEps     string `json:"last_year_eps"`
					Date            string `json:"date"`
				} `json:"current_quarter"`
				NextQuarter string `json:"next_quarter"`
				NextYear    struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					LastYearEps     string `json:"last_year_eps"`
					Year            string `json:"year"`
				} `json:"next_year"`
			} `json:"earnings_estimate"`
			EarningsHistory struct {
				ThreeMonthAgo string `json:"3_month_ago"`
				SixMonthAgo   struct {
					EstimatedEps     string `json:"estimated_eps"`
					RealEps          string `json:"real_eps"`
					DifferencePoints string `json:"difference_points"`
					DifferencePerc   string `json:"difference_perc"`
					Date             string `json:"date"`
				} `json:"6_month_ago"`
				NineMonthAgo struct {
					EstimatedEps     string `json:"estimated_eps"`
					RealEps          string `json:"real_eps"`
					DifferencePoints string `json:"difference_points"`
					DifferencePerc   string `json:"difference_perc"`
					Date             string `json:"date"`
				} `json:"9_month_ago"`
				One2MonthAgo struct {
					EstimatedEps     string `json:"estimated_eps"`
					RealEps          string `json:"real_eps"`
					DifferencePoints string `json:"difference_points"`
					DifferencePerc   string `json:"difference_perc"`
					Date             string `json:"date"`
				} `json:"12_month_ago"`
			} `json:"earnings_history"`
			RevenueEstimate struct {
				PreviousQuarter struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					SalesLastYear   string `json:"sales_last_year"`
					GrowthPerc      string `json:"growth_perc"`
					Date            string `json:"date"`
				} `json:"previous_quarter"`
				CurrentQuarter struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					SalesLastYear   string `json:"sales_last_year"`
					GrowthPerc      string `json:"growth_perc"`
					Date            string `json:"date"`
				} `json:"current_quarter"`
				NextQuarter string `json:"next_quarter"`
				NextYear    struct {
					AnalystsNumber  string `json:"analysts_number"`
					AverageEstimate string `json:"average_estimate"`
					LowEstimate     string `json:"low_estimate"`
					HighEstimate    string `json:"high_estimate"`
					SalesLastYear   string `json:"sales_last_year"`
					GrowthPerc      string `json:"growth_perc"`
					Year            string `json:"year"`
				} `json:"next_year"`
			} `json:"revenue_estimate"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /searchtrends response object
* https://finnworlds.com/documentation/#searchtrends
 */
type SearchTrendsResponse struct {
	Status struct {
		Message string `json:"message"`
	} `json:"status"`
	Results []struct {
		Basics struct {
			Name              string `json:"name"`
			StockTickerSymbol string `json:"stock_ticker_symbol"`
			IsinIdentifier    string `json:"isin_identifier"`
			Exchange          string `json:"exchange"`
		} `json:"basics"`
		Terms struct {
			SearchTerm     string `json:"search_term"`
			SearchQuestion string `json:"search_question"`
		} `json:"terms"`
		Output []struct {
			RelatedTerm           string `json:"related_term,omitempty"`
			RelatedTermVolume     string `json:"related_term_volume,omitempty"`
			RelatedQuestion       string `json:"related_question,omitempty"`
			RelatedQuestionVolume string `json:"related_question_volume,omitempty"`
		} `json:"output"`
	} `json:"results"`
}

/*
* The FinnWorld /secfilings response object
* https://finnworlds.com/documentation/#secfilings
 */
type SECFilingsResponse struct {
	Status struct {
		Message string `json:"message"`
	} `json:"status"`
	Results []struct {
		Basics struct {
			Name              string `json:"name"`
			StockTickerSymbol string `json:"stock_ticker_symbol"`
			IsinIdentifier    string `json:"isin_identifier"`
			Exchange          string `json:"exchange"`
		} `json:"basics"`
		Output []struct {
			FormType      string `json:"form_type"`
			Title         string `json:"title"`
			Link          string `json:"link"`
			FilingDate    string `json:"filing_date"`
			ReportingDate string `json:"reporting_date"`
		} `json:"output"`
	} `json:"results"`
}

/*
* The FinnWorld /etfinfo response object
* https://finnworlds.com/documentation/#etfinfo
 */
type ETFInfoResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			FundName   string `json:"fund_name"`
			FileNumber string `json:"file_number"`
			Cik        string `json:"cik"`
			RegLei     string `json:"reg_lei"`
		} `json:"basics"`
		Output struct {
			General struct {
				FundName string `json:"fund_name"`
				Cik      string `json:"cik"`
				RegLei   string `json:"reg_lei"`
				Country  string `json:"country"`
				State    string `json:"state"`
				City     string `json:"city"`
				Zip      string `json:"zip"`
				Street   string `json:"street"`
				Phone    string `json:"phone"`
			} `json:"general"`
			Results []struct {
				Attributes struct {
					SeriesName       string `json:"series_name"`
					SeriesID         string `json:"series_id"`
					SeriesLei        string `json:"series_lei"`
					DateReportPeriod string `json:"date_report_period"`
					EndReportPeriod  string `json:"end_report_period"`
					FinalFiling      int    `json:"final_filing"`
				} `json:"attributes"`
				FundInfo struct {
					TotalAssets           string `json:"total_assets"`
					TotalLiabilities      string `json:"total_liabilities"`
					NetAssets             string `json:"net_assets"`
					AssetsAttrMiscSec     string `json:"assets_attr_misc_sec"`
					InvestedAssets        string `json:"invested_assets"`
					OneYrBanksBorr        string `json:"one_yr_banks_borr"`
					OneYrCtrldComp        string `json:"one_yr_ctrld_comp"`
					OneYrOthAffil         string `json:"one_yr_oth_affil"`
					OneYrOther            string `json:"one_yr_other"`
					AftOneYrBanksBorr     string `json:"aft_one_yr_banks_borr"`
					AftOneYrCtrldComp     string `json:"aft_one_yr_ctrld_comp"`
					AftOneYrOthAffil      string `json:"aft_one_yr_oth_affil"`
					AftOneYrOther         string `json:"aft_one_yr_other"`
					DeliveryDelay         string `json:"delivery_delay"`
					StandbyCommit         string `json:"standby_commit"`
					LiquidPref            string `json:"liquid_pref"`
					CashNotReportedInCorD string `json:"cash_not_reported_in_cor_d"`
					NonCashCollateral     string `json:"non_cash_collateral"`
					Currency              string `json:"currency"`
				} `json:"fund_info"`
			} `json:"results"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /etfreturns response object
* https://finnworlds.com/documentation/#etfreturns
 */
type ETFReturnsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			FundName   string `json:"fund_name"`
			FileNumber string `json:"file_number"`
			Cik        string `json:"cik"`
			RegLei     string `json:"reg_lei"`
		} `json:"basics"`
		Output []struct {
			Attributes struct {
				SeriesName       string `json:"series_name"`
				SeriesID         string `json:"series_id"`
				SeriesLei        string `json:"series_lei"`
				DateReportPeriod string `json:"date_report_period"`
				EndReportPeriod  string `json:"end_report_period"`
				FinalFiling      int    `json:"final_filing"`
			} `json:"attributes"`
			Signature struct {
				DateSigned      string `json:"date_signed"`
				NameOfApplicant string `json:"name_of_applicant"`
				Signature       string `json:"signature"`
				SignerName      string `json:"signer_name"`
				Title           string `json:"title"`
			} `json:"signature"`
			ReturnInfo struct {
				MonthlyTotalReturns []struct {
					Num0 struct {
						Rtn1 string `json:"rtn1"`
						Rtn2 string `json:"rtn2"`
						Rtn3 string `json:"rtn3"`
					} `json:"0"`
					ClassID string `json:"class_id"`
				} `json:"monthly_total_returns"`
				OtherMonthly1 struct {
					NetRealizedGain           string `json:"net_realized_gain"`
					NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
				} `json:"other_monthly1"`
				OtherMonthly2 struct {
					NetRealizedGain           string `json:"net_realized_gain"`
					NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
				} `json:"other_monthly2"`
				OtherMonthly3 struct {
					NetRealizedGain           string `json:"net_realized_gain"`
					NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
				} `json:"other_monthly3"`
				Month1Flow struct {
					Redemption   string `json:"redemption"`
					Reinvestment string `json:"reinvestment"`
					Sales        string `json:"sales"`
				} `json:"month_1_flow"`
				Month2Flow struct {
					Redemption   string `json:"redemption"`
					Reinvestment string `json:"reinvestment"`
					Sales        string `json:"sales"`
				} `json:"month_2_flow"`
				Month3Flow struct {
					Redemption   string `json:"redemption"`
					Reinvestment string `json:"reinvestment"`
					Sales        string `json:"sales"`
				} `json:"month_3_flow"`
				MonthlyReturnCategories struct {
					Commodity struct {
						Month1 struct {
							NetRealizedGain           string `json:"net_realized_gain"`
							NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
						} `json:"month1"`
						Month2 struct {
							NetRealizedGain           string `json:"net_realized_gain"`
							NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
						} `json:"month2"`
						Month3 struct {
							NetRealizedGain           string `json:"net_realized_gain"`
							NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
						} `json:"month3"`
						ForwardCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"forward_category"`
						FutureCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"future_category"`
						OptionCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"option_category"`
						SwaptionCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"swaption_category"`
						SwapCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"swap_category"`
						WarrantCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"warrant_category"`
						OtherCategory struct {
							Instrmon1 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon1"`
							Instrmon2 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon2"`
							Instrmon3 struct {
								NetRealizedGain           string `json:"net_realized_gain"`
								NetUnrealizedAppreciation string `json:"net_unrealized_appreciation"`
							} `json:"instrmon3"`
						} `json:"other_category"`
					} `json:"commodity"`
				} `json:"monthly_return_categories"`
			} `json:"return_info"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /etfholdings response object
* https://finnworlds.com/documentation/#etfholdings
 */
type ETFHoldingsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Result struct {
		Basics struct {
			FundName   string `json:"fund_name"`
			FileNumber string `json:"file_number"`
			Cik        string `json:"cik"`
			RegLei     string `json:"reg_lei"`
		} `json:"basics"`
		Output []struct {
			Attributes struct {
				SeriesName       string `json:"series_name"`
				SeriesID         string `json:"series_id"`
				SeriesLei        string `json:"series_lei"`
				DateReportPeriod string `json:"date_report_period"`
				EndReportPeriod  string `json:"end_report_period"`
				FinalFiling      int    `json:"final_filing"`
			} `json:"attributes"`
			Signature struct {
				DateSigned      string `json:"date_signed"`
				NameOfApplicant string `json:"name_of_applicant"`
				Signature       string `json:"signature"`
				SignerName      string `json:"signer_name"`
				Title           string `json:"title"`
			} `json:"signature"`
			Holdings []struct {
				InvestmentSecurity struct {
					Name              string `json:"name"`
					Ticker            string `json:"ticker"`
					Lei               string `json:"lei"`
					Isin              string `json:"isin"`
					Title             string `json:"title"`
					Cusip             string `json:"cusip"`
					Balance           string `json:"balance"`
					Units             string `json:"units"`
					Currency          string `json:"currency"`
					ValueUsd          string `json:"value_usd"`
					PercentValue      string `json:"percent_value"`
					PayoffProfile     string `json:"payoff_profile"`
					AssetCategory     string `json:"asset_category"`
					IssuerCategory    string `json:"issuer_category"`
					InvestedCountry   string `json:"invested_country"`
					RestrictedSec     string `json:"restricted_sec"`
					FairValueLevel    string `json:"fair_value_level"`
					CashCollateral    string `json:"cash_collateral"`
					NonCashCollateral string `json:"non_cash_collateral"`
					LoanByFund        string `json:"loan_by_fund"`
				} `json:"investment_security"`
			} `json:"holdings"`
		} `json:"output"`
	} `json:"result"`
}

/*
* The FinnWorld /benchmark response object
* https://finnworlds.com/documentation/#benchmark
 */
type BenchmarkResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Region          string `json:"region"`
		Country         string `json:"country"`
		Benchmark       string `json:"benchmark"`
		Price           string `json:"price"`
		PriceChangeDay  string `json:"price_change_day"`
		PercentageDay   string `json:"percentage_day"`
		PercentageWeek  string `json:"percentage_week"`
		PercentageMonth string `json:"percentage_month"`
		PercentageYear  string `json:"percentage_year"`
		Date            string `json:"date"`
	} `json:"data"`
}

/*
* The FinnWorld /bonds response object
* https://finnworlds.com/documentation/#bonds
 */
type BondsResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"status"`
	Data struct {
		Region          string `json:"region"`
		Country         string `json:"country"`
		Type            string `json:"type"`
		Yield           string `json:"yield"`
		PriceChangeDay  string `json:"price_change_day"`
		PercentageWeek  string `json:"percentage_week"`
		PercentageMonth string `json:"percentage_month"`
		PercentageYear  string `json:"percentage_year"`
		Date            string `json:"date"`
	} `json:"data"`
}
