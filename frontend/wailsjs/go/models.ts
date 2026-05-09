export namespace service {
	
	export class ProfitMarginPayload {
	    purchaseAmount: number;
	    purchaseQuantity: number;
	    expressFee: number;
	    boxUnitPrice: number;
	    bubbleBagUnitPrice: number;
	    miscUnitPrice: number;
	    saleUnitPrice: number;
	    saleQuantity: number;
	
	    static createFrom(source: any = {}) {
	        return new ProfitMarginPayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.purchaseAmount = source["purchaseAmount"];
	        this.purchaseQuantity = source["purchaseQuantity"];
	        this.expressFee = source["expressFee"];
	        this.boxUnitPrice = source["boxUnitPrice"];
	        this.bubbleBagUnitPrice = source["bubbleBagUnitPrice"];
	        this.miscUnitPrice = source["miscUnitPrice"];
	        this.saleUnitPrice = source["saleUnitPrice"];
	        this.saleQuantity = source["saleQuantity"];
	    }
	}
	export class ProfitMarginResult {
	    unitPurchaseCost: number;
	    unitPackageCost: number;
	    totalPackageCost: number;
	    totalExpressCost: number;
	    shipmentTotalCost: number;
	    totalRevenue: number;
	    shipmentProfit: number;
	    shipmentProfitRate: number;
	
	    static createFrom(source: any = {}) {
	        return new ProfitMarginResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.unitPurchaseCost = source["unitPurchaseCost"];
	        this.unitPackageCost = source["unitPackageCost"];
	        this.totalPackageCost = source["totalPackageCost"];
	        this.totalExpressCost = source["totalExpressCost"];
	        this.shipmentTotalCost = source["shipmentTotalCost"];
	        this.totalRevenue = source["totalRevenue"];
	        this.shipmentProfit = source["shipmentProfit"];
	        this.shipmentProfitRate = source["shipmentProfitRate"];
	    }
	}

}

