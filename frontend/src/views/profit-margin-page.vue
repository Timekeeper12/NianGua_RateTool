<template>
  <main class="page">
    <section class="title-wrap">
      <el-card shadow="never" class="title-card">
        <div class="title-row">
          <el-button type="warning" plain @click="goBack">返回工具首页</el-button>
          <el-tag type="warning" effect="light" round>文玩利润模型</el-tag>
        </div>
        <h1>利润率测算</h1>
        <p>录入采购与销售数据，自动生成单只成本拆解、利润与利润率结果。</p>
      </el-card>
    </section>

    <section class="content-wrap">
      <el-row :gutter="16">
        <el-col :xs="24" :lg="14">
          <el-card shadow="hover" class="panel-card">
            <template #header>
              <span>基础数据</span>
            </template>

            <el-form label-position="top" class="form-grid">
              <el-form-item label="品类采购金额">
                <el-input-number v-model="form.purchaseAmount" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="品类采购数量">
                <el-input-number v-model="form.purchaseQuantity" :min="0" :step="1" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="快递费">
                <el-input-number v-model="form.expressFee" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="包装盒单价">
                <el-input-number v-model="form.boxUnitPrice" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="气泡袋单价">
                <el-input-number v-model="form.bubbleBagUnitPrice" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="标签等杂项单价">
                <el-input-number v-model="form.miscUnitPrice" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="销售单价">
                <el-input-number v-model="form.saleUnitPrice" :min="0" :step="0.01" :precision="2" controls-position="right" class="full-input" />
              </el-form-item>
              <el-form-item label="销售数量">
                <el-input-number v-model="form.saleQuantity" :min="0" :step="1" controls-position="right" class="full-input" />
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>

        <el-col :xs="24" :lg="10">
          <el-card shadow="hover" class="panel-card">
            <template #header>
              <span>测算结果</span>
            </template>

            <div class="stats-group">
              <h3>单只拆分成本</h3>
              <el-statistic title="采购单只成本" :value="calculated.unitPurchaseCost" prefix="¥" />
              <el-statistic title="快递分摊单只" :value="calculated.unitExpressCost" prefix="¥" />
              <el-statistic title="包材杂项单只" :value="calculated.unitPackageCost" prefix="¥" />
              <el-statistic title="单只总成本" :value="calculated.unitTotalCost" prefix="¥" />
            </div>

            <el-divider />

            <div class="stats-group">
              <h3>利润与利润率</h3>
              <el-statistic title="单只利润" :value="calculated.unitProfit" prefix="¥" />
              <el-statistic title="单只利润率" :value="calculated.unitProfitRate" suffix="%" />
              <el-statistic title="销售总额" :value="calculated.totalRevenue" prefix="¥" />
              <el-statistic title="总利润" :value="calculated.totalProfit" prefix="¥" />
              <el-statistic title="总利润率" :value="calculated.totalProfitRate" suffix="%" />
            </div>
          </el-card>
        </el-col>
      </el-row>
    </section>
  </main>
</template>

<script setup>
import { reactive, watch } from 'vue'
import {
  ElButton,
  ElCard,
  ElCol,
  ElDivider,
  ElForm,
  ElFormItem,
  ElInputNumber,
  ElRow,
  ElStatistic,
  ElTag
} from 'element-plus'

const emit = defineEmits(['back'])

const form = reactive({
  purchaseAmount: 0,
  purchaseQuantity: 0,
  expressFee: 0,
  boxUnitPrice: 0,
  bubbleBagUnitPrice: 0,
  miscUnitPrice: 0,
  saleUnitPrice: 0,
  saleQuantity: 0
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
  padding: 22px 20px 38px;
  box-sizing: border-box;
  background:
    radial-gradient(circle at 12% 0, rgba(215, 173, 109, 0.3), transparent 25%),
    radial-gradient(circle at 88% 100%, rgba(233, 205, 166, 0.35), transparent 30%),
    linear-gradient(180deg, #f8f3ea, #efe6d8);
}

.title-wrap,
.content-wrap {
  width: min(1140px, 96vw);
  margin: 0 auto 16px;
}

.title-card {
  border: 1px solid #d4bc97;
  border-radius: 16px;
  background: linear-gradient(140deg, #fff8ef, #f7eddc);
  color: #4a3520;
}

.title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
}

.title-card h1 {
  margin: 12px 0 8px;
  color: #382616;
  font-size: clamp(26px, 3.5vw, 36px);
}

.title-card p {
  margin: 0;
  color: #6d5437;
}

.panel-card {
  min-height: 100%;
  border: 1px solid #deceb5;
  border-radius: 14px;
  background: linear-gradient(180deg, #fffdfa, #f9f1e6);
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0 10px;
}

.full-input {
  width: 100%;
}

.stats-group {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 12px;
}

.stats-group h3 {
  margin: 0 0 6px;
  grid-column: 1 / -1;
  color: #6f5538;
  font-size: 16px;
}

:deep(.el-card__header) {
  border-bottom-color: #eadfcc;
  color: #4f3821;
  font-size: 16px;
  font-weight: 600;
}

:deep(.el-form-item__label) {
  color: #70573a;
}

:deep(.el-input-number) {
  width: 100%;
}

:deep(.el-statistic__head) {
  color: #7e6750;
  font-size: 13px;
}

:deep(.el-statistic__content) {
  color: #3f2b1a;
  font-size: 22px;
}

@media (max-width: 860px) {
  .form-grid,
  .stats-group {
    grid-template-columns: 1fr;
  }
}
</style>
