<template>
  <div ref="boxEl" :class="classComputed" :style="styleComputed" class="rounded-[var(--border-radius)] bg-[var(--bg-color-secondary)] float-left">
      <slot :size />
  </div>
</template>

<script setup lang="ts">
import {onUnmounted} from 'vue';
import { boxEl, init } from '@/components/base/grid/index';
import type {layoutMap, cmdObj} from '@/types/index.d.ts';

const props = defineProps({
  layoutMap: {
    type: Object as () => layoutMap,
    default: () => {
      return {}
    },
  },
  cmdObjArr: {
    type: Object as () => cmdObj[],
    default: () => ([])
  },
})

const {
  size,
  unmounted,
  classComputed,
  styleComputed
} = init(props);
onUnmounted(unmounted)
</script>


<style lang="less">
</style>
