package service

import "math"

type ProfitMarginPayload struct {
	PurchaseAmount     float64 `json:"purchaseAmount"`
	PurchaseQuantity   float64 `json:"purchaseQuantity"`
	ExpressFee         float64 `json:"expressFee"`
	BoxUnitPrice       float64 `json:"boxUnitPrice"`
	BubbleBagUnitPrice float64 `json:"bubbleBagUnitPrice"`
	MiscUnitPrice      float64 `json:"miscUnitPrice"`
	SaleUnitPrice      float64 `json:"saleUnitPrice"`
	SaleQuantity       float64 `json:"saleQuantity"`
}

type ProfitMarginResult struct {
	UnitPurchaseCost float64 `json:"unitPurchaseCost"`
	UnitExpressCost  float64 `json:"unitExpressCost"`
	UnitPackageCost  float64 `json:"unitPackageCost"`
	UnitTotalCost    float64 `json:"unitTotalCost"`
	UnitProfit       float64 `json:"unitProfit"`
	UnitProfitRate   float64 `json:"unitProfitRate"`
	TotalRevenue     float64 `json:"totalRevenue"`
	TotalProfit      float64 `json:"totalProfit"`
	TotalProfitRate  float64 `json:"totalProfitRate"`
}

func roundTo2(value float64) float64 {
	return math.Round(value*100) / 100
}

func CalculateProfitMargin(payload ProfitMarginPayload) ProfitMarginResult {
	validPurchaseQuantity := 0.0
	if payload.PurchaseQuantity > 0 {
		validPurchaseQuantity = payload.PurchaseQuantity
	}

	unitPurchaseCost := 0.0
	if validPurchaseQuantity > 0 {
		unitPurchaseCost = payload.PurchaseAmount / validPurchaseQuantity
	}

	unitExpressCost := 0.0
	if validPurchaseQuantity > 0 {
		unitExpressCost = payload.ExpressFee / validPurchaseQuantity
	}

	unitPackageCost := payload.BoxUnitPrice + payload.BubbleBagUnitPrice + payload.MiscUnitPrice
	unitTotalCost := unitPurchaseCost + unitExpressCost + unitPackageCost
	unitProfit := payload.SaleUnitPrice - unitTotalCost

	unitProfitRate := 0.0
	if payload.SaleUnitPrice > 0 {
		unitProfitRate = unitProfit / payload.SaleUnitPrice
	}

	totalRevenue := payload.SaleUnitPrice * payload.SaleQuantity
	totalProfit := unitProfit * payload.SaleQuantity

	totalProfitRate := 0.0
	if totalRevenue > 0 {
		totalProfitRate = totalProfit / totalRevenue
	}

	return ProfitMarginResult{
		UnitPurchaseCost: roundTo2(unitPurchaseCost),
		UnitExpressCost:  roundTo2(unitExpressCost),
		UnitPackageCost:  roundTo2(unitPackageCost),
		UnitTotalCost:    roundTo2(unitTotalCost),
		UnitProfit:       roundTo2(unitProfit),
		UnitProfitRate:   roundTo2(unitProfitRate * 100),
		TotalRevenue:     roundTo2(totalRevenue),
		TotalProfit:      roundTo2(totalProfit),
		TotalProfitRate:  roundTo2(totalProfitRate * 100),
	}
}
