package Calculations

func TangibleNetWorth(TotalAssets float64, TotalLiabilities float64, IntangibleAssets float64) float64 {
	// Tangible net worth is most commonly a calculation of the net worth of a company that excludes any value derived from intangible assets such as copyrights.

	return TotalAssets - TotalLiabilities - IntangibleAssets
}
