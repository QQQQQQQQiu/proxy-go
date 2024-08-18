import {ref, computed,reactive, nextTick} from 'vue';
import {getElSize} from '@/utils/func';
import {setContainerSize} from '@/store/index';

export const containerEl = ref<HTMLElement | null>(null)


export function init () {
  const size = reactive({
    width: 0,
    height: 0
  })

  const watchEl = async  () => {
    await nextTick()
    const {width, height} = getElSize(containerEl.value as HTMLElement)
    size.width = width
    size.height = height
    setContainerSize({width, height})
  }
  watchEl()
  window.addEventListener('resize', watchEl)

  return {
    size,
    unmounted: () => {
      window.removeEventListener('resize', watchEl)
    }
  }
}