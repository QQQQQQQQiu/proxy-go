import {containerModeStr} from '@/store/index';
import {computed, ref} from 'vue';
import type {layoutMap as _layoutMap, cmdObj, cmdRespObj} from '@/types/index.d.ts';
import {fixNumber} from '@/utils/func';

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

export const persent = ref(0)
export const usedComputed = computed(()=>{
  return {
    rotate: `rotate(${(persent.value / 100) * 180 - 90}deg)`
  }
})

export const cmdObjArr = ref([
  {id: 'c1', cmd: 'top -b -n 1 | grep "Cpu(s)"'},
  {id: 'c2', cmd: `ps aux --sort=-%cpu | head -n 6`},
])
export function onCmdResp(arr: cmdRespObj[]) {
  console.log('[cpu] onCmdResp :>> ', arr);
  for (let index = 0; index < arr.length; index++) {
    const {id, output} = arr[index];
    switch (id) {
      case 'c1':
        calcUsedNum(output)
        break;
      case 'c2':
        calcTopCpuProc(output)
        break;
    
      default:
        break;
    }
  }
}

function calcUsedNum(outputStr: string) {
  // 正则表达式匹配 us 和 sy 的值
  const regex = /(\d+\.\d+)\s+us,\s+(\d+\.\d+)\s+sy/;
  const match = outputStr.match(regex);
  if (match) {
      const usValue = match[1]; // us 的值
      const syValue = match[2]; // sy 的值
      persent.value = fixNumber(Number(usValue) + Number(syValue), 1)
  } else {
      console.log("没有找到匹配的值");
  } 
}
function calcTopCpuProc(outputStr: string) {
  console.log('calcTopCpuProc :>> ', outputStr);
}