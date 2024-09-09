<template>
  <grid :layoutMap="layoutMap" :cmdObjArr="cmdObjArr" :onCmdResp="onCmdResp">
    <template v-slot:default="{size}">
      <div class="w-full h-full" :style="styleComputed.box">
        <p>公网ipv4：{{ipv4_pub}}%</p>
      </div>
    </template>
  </grid>
</template>

<script setup lang="ts">
import grid from '@/components/base/grid/index.vue';
import {containerModeStr} from '@/store/index';
import {layoutMap, styleComputed, usedComputed, persent, onCmdResp, cmdObjArr} from './index.ts';
import {ref} from 'vue';

</script>


<style lang="less">
.loader {
  width: 96px;
  height: 48px;
  display: inline-block;
  position: relative;
  background: #5e5e5e;
  border-radius: 48px 48px 0 0;
  box-sizing: border-box;
  overflow: hidden;
  &::after {
    content: '';
    box-sizing: border-box;
    position: absolute;
    width: 24px;
    height: 12px;
    border-radius: 24px 24px 0 0;
    background: yellow;
    left: 50%;
    transform: translateX(-50%);
    bottom: 0;
  }
  &::before {
    content: '';
    position: absolute;
    width: 4px;
    height: 32px;
    left: 0;
    right: 0;
    margin: auto;
    bottom: 0;
    background: yellow;
    transform-origin: 50% 100%;
    box-sizing: border-box;
    transform: var(--rotate);
    transition: transform 0.6s ease;
  }
}

</style>
