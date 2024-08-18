<template>
  <div class="w-full grid" :style="style">
      <slot />
  </div>
</template>

<script setup lang="ts">
import { style } from '@/components/base/gridContainer/index.ts';

</script>


<style lang="less">

</style>
