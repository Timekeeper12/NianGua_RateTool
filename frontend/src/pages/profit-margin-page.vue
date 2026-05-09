<template>
  <main class="page">
    <header class="header">
      <button class="back-btn" type="button" @click="goBack">返回工具首页</button>
      <h1>利润率测算</h1>
      <p>填写采购、物流、包材与销售数据，自动输出单只利润和总利润情况。</p>
    </header>

    <section class="panel">
      <h2>基础数据</h2>
      <div class="form-grid">
        <label>
          <span>品类采购金额</span>
          <input v-model="form.purchaseAmount" type="number" min="0" step="0.01" placeholder="如：1000" />
        </label>
        <label>
          <span>品类采购数量</span>
          <input v-model="form.purchaseQuantity" type="number" min="0" step="1" placeholder="如：200" />
        </label>
        <label>
          <span>快递费</span>
          <input v-model="form.expressFee" type="number" min="0" step="0.01" placeholder="如：18" />
        </label>
        <label>
          <span>包装盒单价</span>
          <input v-model="form.boxUnitPrice" type="number" min="0" step="0.01" placeholder="如：0.8" />
        </label>
        <label>
          <span>气泡袋单价</span>
          <input v-model="form.bubbleBagUnitPrice" type="number" min="0" step="0.01" placeholder="如：0.2" />
        </label>
        <label>
          <span>标签等杂项单价</span>
          <input v-model="form.miscUnitPrice" type="number" min="0" step="0.01" placeholder="如：0.15" />
        </label>
        <label>
          <span>销售单价</span>
          <input v-model="form.saleUnitPrice" type="number" min="0" step="0.01" placeholder="如：19.9" />
        </label>
        <label>
          <span>销售数量</span>
          <input v-model="form.saleQuantity" type="number" min="0" step="1" placeholder="如：120" />
        </label>
      </div>
    </section>

    <section class="panel">
      <h2>测算结果</h2>
      <div class="result-grid">
        <article>
          <h3>单只拆分成本</h3>
          <p>采购单只成本：¥{{ calculated.unitPurchaseCost }}</p>
          <p>快递分摊单只：¥{{ calculated.unitExpressCost }}</p>
          <p>包材杂项单只：¥{{ calculated.unitPackageCost }}</p>
          <p>单只总成本：¥{{ calculated.unitTotalCost }}</p>
        </article>

        <article>
          <h3>利润与利润率</h3>
          <p>单只利润：¥{{ calculated.unitProfit }}</p>
          <p>单只利润率：{{ calculated.unitProfitRate }}%</p>
          <p>销售总额：¥{{ calculated.totalRevenue }}</p>
          <p>总利润：¥{{ calculated.totalProfit }}</p>
          <p>总利润率：{{ calculated.totalProfitRate }}%</p>
        </article>
      </div>
    </section>
  </main>
</template>

<script setup>
import { reactive, watch } from 'vue'

const emit = defineEmits(['back'])

const form = reactive({
  purchaseAmount: '',
  purchaseQuantity: '',
  expressFee: '',
  boxUnitPrice: '',
  bubbleBagUnitPrice: '',
  miscUnitPrice: '',
  saleUnitPrice: '',
  saleQuantity: ''
})

const createEmptyCalculated = () => ({
  unitPurchaseCost: 0,
  unitExpressCost: 0,
  unitPackageCost: 0,
  unitTotalCost: 0,
  unitProfit: 0,
  unitProfitRate: 0,
  totalRevenue: 0,
  totalProfit: 0,
  totalProfitRate: 0
})

const calculated = reactive(createEmptyCalculated())

const toNumber = (value) => {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : 0
}

const applyCalculatedResult = (result) => {
  Object.assign(calculated, {
    unitPurchaseCost: toNumber(result?.unitPurchaseCost),
    unitExpressCost: toNumber(result?.unitExpressCost),
    unitPackageCost: toNumber(result?.unitPackageCost),
    unitTotalCost: toNumber(result?.unitTotalCost),
    unitProfit: toNumber(result?.unitProfit),
    unitProfitRate: toNumber(result?.unitProfitRate),
    totalRevenue: toNumber(result?.totalRevenue),
    totalProfit: toNumber(result?.totalProfit),
    totalProfitRate: toNumber(result?.totalProfitRate)
  })
}

const calculateByBackend = async () => {
  const payload = {
    purchaseAmount: toNumber(form.purchaseAmount),
    purchaseQuantity: toNumber(form.purchaseQuantity),
    expressFee: toNumber(form.expressFee),
    boxUnitPrice: toNumber(form.boxUnitPrice),
    bubbleBagUnitPrice: toNumber(form.bubbleBagUnitPrice),
    miscUnitPrice: toNumber(form.miscUnitPrice),
    saleUnitPrice: toNumber(form.saleUnitPrice),
    saleQuantity: toNumber(form.saleQuantity)
  }

  try {
    const result = await window.go.main.App.CalculateProfitMargin(payload)
    applyCalculatedResult(result)
  } catch (error) {
    Object.assign(calculated, createEmptyCalculated())
  }
}

watch(form, () => {
  calculateByBackend()
}, {
  deep: true,
  immediate: true
})

const goBack = () => {
  emit('back')
}
</script>

<style scoped>
.page {
  min-height: 100vh;
  padding: 24px 20px 40px;
  box-sizing: border-box;
}

.header {
  width: min(980px, 96vw);
  margin: 0 auto 16px;
  text-align: left;
}

.back-btn {
  width: 118px;
  height: 34px;
  border: 0;
  border-radius: 8px;
  cursor: pointer;
  background: #3b82f6;
  color: #ffffff;
  margin-bottom: 10px;
}

.header h1 {
  margin: 0 0 8px;
}

.header p {
  margin: 0;
  color: #d0d8e3;
}

.panel {
  width: min(980px, 96vw);
  margin: 0 auto 16px;
  border: 1px solid #2e3d52;
  border-radius: 12px;
  padding: 16px;
  background: #172131;
  text-align: left;
}

.panel h2 {
  margin: 0 0 12px;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;
}

.form-grid label {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-grid span {
  color: #d7deea;
  font-size: 14px;
}

.form-grid input {
  height: 36px;
  border: 1px solid #354861;
  border-radius: 8px;
  padding: 0 10px;
  background: #0f1726;
  color: #eff4ff;
}

.result-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 14px;
}

.result-grid article {
  border: 1px solid #2e3d52;
  border-radius: 10px;
  padding: 12px;
  background: #0f1726;
}

.result-grid h3 {
  margin: 0 0 10px;
}

.result-grid p {
  margin: 6px 0;
  color: #dde5f2;
}
</style>
