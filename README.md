
# Securities Analysis Based AutoTrading

1. Gather up the fundamentals of a ton of companies (at least 1000)
2. Perform and document as many securities analysis functions as possible against the financial statements
3. Give each result a weight from -1.0 to 1.0, the final value will be result*weight
4. Organize each company in order by some aggregate value func(calculations*weights) and pick the top 10 companies to backtest
5. Identify the 'best' weights using a (genetic algorithm, simulated annealing, whatever) and multi-range backtesting
6. Add backtesting purchase/sale rules to make the algorithm more proactive / reactive
7. Add weights to each of the buy/sell indicators and repeat the algorithmic weight optimization process switching between optimizing weights for securities analysis using optimized buy/sell weights and weights for buy/sell signals with optimized security analysis weights
8. Optimization should be based on some function that works to maximize certain portfolio risk/reward values while minimizing others based on their definitions 
9. Research news related to each company in the 10 company portfolio using biztoc and chatgpt 
10. Perform sentiment analysis on news reports and compare it to sentiment provided by chatgpt 
11. give sentiment analysis and chatgpt analysis points weights between -1.0 and 1.0 
12. Perform algorithmic optimization on the weights, this part could require a ton of time

## Progress

- [x] gather company fundamentals and calculations
- [x] create a weight system from -1.0 to 1.0 for each fundamentals item 
    - [ ] TODO: fix the full financial document and employee count retrievers and add the full financial document type to the final results struct type / weight type for weighting
- [x] calculate a normalized weighted value that represents every value in all of the fundamentals / calculations * their field weights
- [ ] set up a genetic algorithm that attempts to maximize the cumulative value based on the weights