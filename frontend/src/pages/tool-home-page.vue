<template>
  <main class="home">
    <section class="hero">
      <h1>年瓜工作室工具箱</h1>
      <p>用于统一管理你日常文玩业务中的常用测算工具。</p>
    </section>

    <section class="tool-grid">
      <article class="tool-card">
        <h2>利润率测算</h2>
        <p>按采购、物流、包材和销售数据自动测算单只利润与利润率。</p>
        <button type="button" @click="openProfitMarginTool">进入工具</button>
      </article>
    </section>
  </main>
</template>

<script setup>
const emit = defineEmits(['open-profit-margin-tool'])

const openProfitMarginTool = () => {
  emit('open-profit-margin-tool')
}
</script>

<style scoped>
.home {
  min-height: 100vh;
  padding: 40px 20px;
  box-sizing: border-box;
}

.hero {
  margin: 0 auto 28px;
  max-width: 760px;
}

.hero h1 {
  margin: 0 0 10px;
  font-size: 34px;
}

.hero p {
  margin: 0;
  color: #d0d8e3;
}

.tool-grid {
  display: grid;
  gap: 16px;
  justify-content: center;
}

.tool-card {
  width: min(520px, 92vw);
  padding: 20px;
  border: 1px solid #2e3d52;
  border-radius: 12px;
  background: #172131;
  text-align: left;
}

.tool-card h2 {
  margin: 0 0 10px;
  font-size: 22px;
}

.tool-card p {
  margin: 0 0 16px;
  color: #d0d8e3;
}

.tool-card button {
  width: 108px;
  height: 36px;
  border: 0;
  border-radius: 8px;
  cursor: pointer;
  background: #60a5fa;
  color: #062040;
  font-weight: 700;
}
</style>
