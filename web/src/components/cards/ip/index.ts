import {containerModeStr} from '@/store/index';
import {computed, ref} from 'vue';
import type {layoutMap as _layoutMap, cmdObj, cmdRespObj} from '@/types/index.d.ts';
export const layoutMap = {
  small: {
    w: 3,
    h: 8,
  },
  medium: {
    w: 2,
    h: 8,
  },
  large: {
    w: 2,
    h: 10,
  }
}
export const styleMap = {
  small: {
    box: {
      padding: '4px'
    }
  },
  medium: {
    box: {
      padding: '8px'
    }
  },
  large: {
    box: {
      padding: '14px'
    }
  }
}

export const styleComputed = computed(()=>{
  return styleMap[containerModeStr.value]
})



export const cmdObjArr = ref([
])
export function onCmdResp(arr: cmdRespObj[]) {
}