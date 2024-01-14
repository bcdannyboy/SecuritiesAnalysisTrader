package financials

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
