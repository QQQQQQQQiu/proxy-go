import {ref, computed, nextTick} from 'vue';
export const containerSize = ref({
  width: 0,
  height: 0
})
export const containerModeMap = {
  Small: 'Small',
  Medium: 'Medium',
  Large: 'Large'
}
export const containerModeStr = computed(() => {
  if (containerSize.value.width < 768) {
    return 'small'
  } else if (containerSize.value.width < 1024) {
    return 'medium'
  } else {
    return 'large'
  }
})

export const isSmall = computed(() => containerModeStr.value === 'small')
export const isMedium = computed(() => containerModeStr.value === 'medium')
export const isLarge = computed(() => containerModeStr.value === 'large')

export function setContainerSize({width, height}){
  containerSize.value.width = width
  containerSize.value.height = height
}