<template>
  <main class="home-page">
    <section class="hero-wrap">
      <el-card shadow="never" class="hero-card">
        <div class="hero-top">
          <el-tag type="warning" effect="light" round>文玩经营工具台</el-tag>
          <span class="hero-year">NianGua Studio</span>
        </div>
        <h1>年瓜工作室 · 运营测算台</h1>
        <p>围绕进货、物流、包材与销售维度，让每一串珠、每一件器物都有清晰利润账。</p>
        <div class="hero-actions">
          <el-button type="warning" size="large" @click="openProfitMarginTool">进入利润率测算</el-button>
          <el-button plain size="large">更多工具筹备中</el-button>
        </div>
      </el-card>
    </section>

    <section class="board-wrap">
      <el-card shadow="hover" class="board-card">
        <template #header>
          <div class="card-header">
            <span>常用工具</span>
            <el-tag type="info" effect="plain">持续扩展</el-tag>
          </div>
        </template>

        <div class="tool-grid">
          <article class="tool-item active">
            <div class="tool-item-title">
              <h2>利润率测算</h2>
              <el-tag type="success" effect="light">可用</el-tag>
            </div>
            <p>按采购、物流、包材和销售数据自动测算单只利润与总利润率。</p>
            <el-divider />
            <div class="tool-item-foot">
              <el-button type="warning" @click="openProfitMarginTool">立即使用</el-button>
            </div>
          </article>

          <article class="tool-item">
            <div class="tool-item-title">
              <h2>库存周转分析</h2>
              <el-tag type="info" effect="light">规划中</el-tag>
            </div>
            <p>跟踪批次进出库与周转天数，辅助优化压货与补货节奏。</p>
            <el-divider />
            <div class="tool-item-foot">
              <el-button disabled>敬请期待</el-button>
            </div>
          </article>
        </div>
      </el-card>
    </section>
  </main>
</template>

<script setup>
import { ElButton, ElCard, ElDivider, ElTag } from 'element-plus'

const emit = defineEmits(['open-profit-margin-tool'])

const openProfitMarginTool = () => {
  emit('open-profit-margin-tool')
}
</script>

<style scoped>
.home-page {
  min-height: 100vh;
  padding: 28px 20px 40px;
  box-sizing: border-box;
  background:
    radial-gradient(circle at 0 0, rgba(215, 173, 109, 0.35), transparent 36%),
    radial-gradient(circle at 100% 100%, rgba(225, 196, 154, 0.35), transparent 30%),
    linear-gradient(180deg, #f7f2e8, #efe6d8);
}

.hero-wrap,
.board-wrap {
  width: min(1100px, 96vw);
  margin: 0 auto 20px;
}

.hero-card {
  border: 1px solid #cfb182;
  border-radius: 18px;
  background: linear-gradient(140deg, #fff8ee, #f7ebd8);
  color: #4b3420;
}

.hero-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-bottom: 12px;
}

.hero-year {
  color: #8d6e45;
  font-size: 13px;
  letter-spacing: 1px;
}

.hero-card h1 {
  margin: 4px 0 10px;
  font-size: clamp(24px, 3vw, 36px);
  line-height: 1.2;
  color: #3b2817;
}

.hero-card p {
  margin: 0 0 20px;
  color: #6f5231;
  max-width: 740px;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.board-card {
  border: 1px solid #d8c7ab;
  border-radius: 16px;
  background: rgba(255, 250, 242, 0.9);
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.tool-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 14px;
}

.tool-item {
  border: 1px solid #ddcfb8;
  border-radius: 14px;
  padding: 16px;
  background: linear-gradient(180deg, #fffdfa, #f8f0e4);
  text-align: left;
}

.tool-item.active {
  border-color: #b68443;
  box-shadow: 0 10px 24px rgba(182, 132, 67, 0.16);
}

.tool-item-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
}

.tool-item h2 {
  margin: 0;
  color: #3d2a1a;
  font-size: 20px;
}

.tool-item p {
  margin: 10px 0 0;
  color: #6a5340;
  min-height: 42px;
}

.tool-item-foot {
  display: flex;
  justify-content: flex-end;
}
</style>
