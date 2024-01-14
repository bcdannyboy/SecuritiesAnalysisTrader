package financials

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
