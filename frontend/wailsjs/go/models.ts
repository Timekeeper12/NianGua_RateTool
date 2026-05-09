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
	    unitExpressCost: number;
	    unitPackageCost: number;
	    unitTotalCost: number;
	    unitProfit: number;
	    unitProfitRate: number;
	    totalRevenue: number;
	    totalProfit: number;
	    totalProfitRate: number;
	
	    static createFrom(source: any = {}) {
	        return new ProfitMarginResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.unitPurchaseCost = source["unitPurchaseCost"];
	        this.unitExpressCost = source["unitExpressCost"];
	        this.unitPackageCost = source["unitPackageCost"];
	        this.unitTotalCost = source["unitTotalCost"];
	        this.unitProfit = source["unitProfit"];
	        this.unitProfitRate = source["unitProfitRate"];
	        this.totalRevenue = source["totalRevenue"];
	        this.totalProfit = source["totalProfit"];
	        this.totalProfitRate = source["totalProfitRate"];
	    }
	}

}

