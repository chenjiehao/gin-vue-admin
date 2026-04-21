<template>
  <div class="h-screen w-full flex items-center justify-center bg-slate-50 dark:bg-slate-900">
    <div class="text-center">
      <h1 class="text-3xl font-semibold text-slate-900 dark:text-slate-100">
        欢迎回来，开始今天的Coding节奏
      </h1>
      <p class="mt-4 text-base text-slate-500 dark:text-slate-400">
        {{ today }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const today = computed(() => {
  try {
    const d = new Date()
    return d.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    })
  } catch (e) {
    return new Date().toISOString().slice(0, 10)
  }
})

defineOptions({
  name: 'Dashboard'
})
</script>

<style lang="scss" scoped></style>
