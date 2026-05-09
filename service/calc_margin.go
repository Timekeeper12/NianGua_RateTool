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
	UnitPurchaseCost   float64 `json:"unitPurchaseCost"`
	UnitPackageCost    float64 `json:"unitPackageCost"`
	TotalPackageCost   float64 `json:"totalPackageCost"`
	TotalExpressCost   float64 `json:"totalExpressCost"`
	ShipmentTotalCost  float64 `json:"shipmentTotalCost"`
	TotalRevenue       float64 `json:"totalRevenue"`
	ShipmentProfit     float64 `json:"shipmentProfit"`
	ShipmentProfitRate float64 `json:"shipmentProfitRate"`
}

func roundTo2(value float64) float64 {
	return math.Round(value*100) / 100
}

func CalculateProfitMargin(payload ProfitMarginPayload) ProfitMarginResult {
	unitPurchaseCost := 0.0
	if payload.PurchaseQuantity > 0 {
		unitPurchaseCost = payload.PurchaseAmount / payload.PurchaseQuantity
	}

	unitPackageCost := payload.BoxUnitPrice + payload.BubbleBagUnitPrice + payload.MiscUnitPrice
	totalPackageCost := unitPackageCost * payload.SaleQuantity

	totalExpressCost := 0.0
	if payload.SaleQuantity > 0 {
		totalExpressCost = payload.ExpressFee
		if payload.SaleQuantity > 3 {
			extraGroups := math.Floor((payload.SaleQuantity - 1) / 3.0)
			totalExpressCost += extraGroups * 3.0
		}
	}

	shipmentTotalCost := (unitPurchaseCost * payload.SaleQuantity) + totalExpressCost + totalPackageCost

	totalRevenue := payload.SaleUnitPrice * payload.SaleQuantity
	shipmentProfit := totalRevenue - shipmentTotalCost

	shipmentProfitRate := 0.0
	if totalRevenue > 0 {
		shipmentProfitRate = shipmentProfit / totalRevenue
	}

	return ProfitMarginResult{
		UnitPurchaseCost:   roundTo2(unitPurchaseCost),
		TotalPackageCost:   roundTo2(totalPackageCost),
		TotalExpressCost:   roundTo2(totalExpressCost),
		ShipmentTotalCost:  roundTo2(shipmentTotalCost),
		TotalRevenue:       roundTo2(totalRevenue),
		ShipmentProfit:     roundTo2(shipmentProfit),
		ShipmentProfitRate: roundTo2(shipmentProfitRate * 100),
	}
}
