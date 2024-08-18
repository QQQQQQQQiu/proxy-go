import {ref, computed, nextTick} from 'vue';
import {containerModeStr} from '@/store/index';

export const style = computed<{ gap: number }>(() => {
  const mode = containerModeStr.value
  let obj =  {
    small: {
      gap: 4,
    },
    medium: {
      gap: 8
    },
    large: {
      gap: 14
    }
  }[mode]
  return {
    [`gap`]: `${obj.gap}px`,
    [`grid-template-columns`]: `repeat(12, 1fr)`,
  }
})