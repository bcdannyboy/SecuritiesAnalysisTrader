package GAUtils

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
)

func PrintBacktestResults(results map[string]Backtest.PortfolioResults) {
	for strategy, backtestResult := range results {
		fmt.Printf("Result for strategy: %s\n", strategy)
		fmt.Printf("Total Return: %f\n", backtestResult.Total.TotalProfitLoss)
		fmt.Printf("Annualized Return: %f\n", backtestResult.Total.AnnualizedReturn)
		fmt.Printf("Volatility: %f\n", backtestResult.Total.Volatility)
		fmt.Printf("Sharpe Ratio: %f\n", backtestResult.Total.SharpeRatio)
		fmt.Printf("Sortino Ratio: %f\n", backtestResult.Total.SortinoRatio)
		fmt.Printf("Max Drawdown: %f\n", backtestResult.Total.MaxDrawdown)
		fmt.Printf("YoY Profit/Loss:\n")
		for year, profitLoss := range backtestResult.Total.YoYProfitLoss {
			fmt.Printf("\t%s: %f\n", year, profitLoss)
		}

		fmt.Printf("Tickers in portfolio:\n")
		for ticker, _ := range backtestResult.IndividualStocks {
			fmt.Printf("\t%s\n", ticker)
		}

	}
}
