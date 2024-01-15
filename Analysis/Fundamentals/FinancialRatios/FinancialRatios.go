package Fundamentals

import (
	"fmt"
	fundamentals "github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"reflect"
)

func AnalyzeFinancialRatios(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.FinancialRatios, []objects.FinancialRatiosTTM, []*fundamentals.FinancialRatiosGrowth, []*fundamentals.FinancialRatiosTTMGrowth, error) {
	var FIN_RATIOS []objects.FinancialRatios
	var FIN_RATIOS_TTM []objects.FinancialRatiosTTM
	var FR_GROWTH []*fundamentals.FinancialRatiosGrowth
	var FR_TTM_GROWTH []*fundamentals.FinancialRatiosTTMGrowth

	FIN_RATIOS, err := APIClient.CompanyValuation.FinancialRatios(
		objects.RequestFinancialRatios{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return FIN_RATIOS, FIN_RATIOS_TTM, FR_GROWTH, FR_TTM_GROWTH, fmt.Errorf("failed to get financial ratios: %s", err.Error())
	}

	FIN_RATIOS_TTM, err = APIClient.CompanyValuation.FinancialRatiosTTM(Symbol)
	if err != nil {
		return FIN_RATIOS, FIN_RATIOS_TTM, FR_GROWTH, FR_TTM_GROWTH, fmt.Errorf("failed to get financial ratios TTM: %s", err.Error())
	}

	FR_GROWTH = GetGrowthOfFinancialRatios(FIN_RATIOS)
	FR_TTM_GROWTH = GetGrowthOfFinancialRatiosTTM(FIN_RATIOS_TTM)

	return FIN_RATIOS, FIN_RATIOS_TTM, FR_GROWTH, FR_TTM_GROWTH, nil
}

func GetGrowthOfFinancialRatios(ratios []objects.FinancialRatios) []*fundamentals.FinancialRatiosGrowth {
	growth := []*fundamentals.FinancialRatiosGrowth{}
	var lastRatios objects.FinancialRatios

	for i, ratio := range ratios {
		growthObj := &fundamentals.FinancialRatiosGrowth{
			Symbol: ratio.Symbol,
			Date:   ratio.Date,
			Period: ratio.Period,
		}

		if i > 0 {
			calculateGrowth(reflect.ValueOf(ratio), reflect.ValueOf(lastRatios), reflect.ValueOf(growthObj).Elem())
		}

		growth = append(growth, growthObj)
		lastRatios = ratio
	}

	return growth
}

func GetGrowthOfFinancialRatiosTTM(ratiosTTM []objects.FinancialRatiosTTM) []*fundamentals.FinancialRatiosTTMGrowth {
	growthTTM := []*fundamentals.FinancialRatiosTTMGrowth{}
	var lastRatiosTTM objects.FinancialRatiosTTM

	for i, ratioTTM := range ratiosTTM {
		growthObjTTM := &fundamentals.FinancialRatiosTTMGrowth{
			Symbol: ratioTTM.Symbol,
		}

		if i > 0 {
			calculateGrowth(reflect.ValueOf(ratioTTM), reflect.ValueOf(lastRatiosTTM), reflect.ValueOf(growthObjTTM).Elem())
		}

		growthTTM = append(growthTTM, growthObjTTM)
		lastRatiosTTM = ratioTTM
	}

	return growthTTM
}

func calculateGrowth(valCurrent, valLast, valGrowth reflect.Value) {
	for j := 0; j < valCurrent.NumField(); j++ {
		fieldCurrent := valCurrent.Field(j)
		fieldLast := valLast.Field(j)
		fieldGrowth := valGrowth.Field(j)

		if fieldCurrent.Kind() == reflect.Float64 {
			growthValue := fieldCurrent.Float() - fieldLast.Float()
			fieldGrowth.SetFloat(growthValue)
		}

		// Handle the case for Interface type fields
		if fieldCurrent.Kind() == reflect.Interface && !fieldCurrent.IsNil() {
			curVal, okCur := fieldCurrent.Interface().(float64)
			lastVal, okLast := fieldLast.Interface().(float64)
			if okCur && okLast {
				fieldGrowth.SetFloat(curVal - lastVal)
			}
		}
	}
}
