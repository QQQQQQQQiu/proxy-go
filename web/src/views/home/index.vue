<template>
  <div class="theme-light bg-[var(--bg-color)] relative w-screen h-screen min-h-screen">
    <div ref="containerEl" class="relative max-w-[1200px] mx-auto">
      <gridContainer>
        <cpu />
        <disk />
      </gridContainer>
    </div>
  </div>
</template>

<script setup lang="ts">
  import gridContainer from '@/components/base/gridContainer/index.vue'
  import disk from '@/components/cards/disk/index.vue'
  import cpu from '@/components/cards/cpu/index.vue'
  import {onUnmounted, ref} from 'vue'
  import {loopQueryConroller} from '@/utils/queryDataBus'
  import {containerEl, init} from '@/views/home/index'
  const {unmounted} = init()
  onUnmounted(unmounted)
  setTimeout(() => {
    loopQueryConroller.start()
  }, 1000)
</script>

<style scoped lang="less"></style>
