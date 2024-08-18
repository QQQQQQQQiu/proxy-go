import {containerModeStr} from '@/store/index';
import {computed} from 'vue';

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

export const usedComputed = computed(()=>{
  let persent = 0
  return {
    rotate: `rotate(${(persent / 100) * 180 - 90}deg)`
  }
})