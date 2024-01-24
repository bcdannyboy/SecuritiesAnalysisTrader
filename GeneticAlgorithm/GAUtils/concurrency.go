package GAUtils

import "runtime"

func DetermineOptimalBatchSize(totalItems int, totalCompanies int) int {
	// Example: Adjust this logic based on your specific requirements and testing
	numCPU := runtime.NumCPU()
	optimalSize := (totalItems * totalCompanies) / (numCPU)
	if optimalSize < numCPU {
		return numCPU
	}
	return optimalSize
}
