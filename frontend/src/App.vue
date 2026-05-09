
<template>
  <ToolHomePage
    v-if="currentPage === 'home'"
    @open-profit-margin-tool="openProfitMarginTool"
  />
  <ProfitMarginPage
    v-else-if="currentPage === 'profit-margin'"
    @back="backToHome"
  />
</template>

<script setup>
import { ref } from 'vue'
import ToolHomePage from './pages/tool-home-page.vue'
import ProfitMarginPage from './pages/profit-margin-page.vue'

const currentPage = ref('home')

const openProfitMarginTool = () => {
  currentPage.value = 'profit-margin'
}

const backToHome = () => {
  currentPage.value = 'home'
}
</script>
