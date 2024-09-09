import {ref, reactive ,computed, nextTick} from 'vue';
import {containerModeStr} from '@/store/index';
import {getElSize} from '@/utils/func';
import type {layoutMap as _layoutMap, cmdObj, cmdRespObj} from '@/types/index.d.ts';
import {addEventListener as addEventListener_queryDataBus, EventsMap as EventsMap_queryDataBus} from '@/utils/queryDataBus';

const GridRowHeight = 12

export const boxEl = ref<HTMLElement | null>(null)

export const boxSize = reactive({
  width: 0,
  height: 0
})

interface Props {
  layoutMap: _layoutMap;
  cmdObjArr: cmdObj[];
  onCmdResp: (resp: cmdRespObj[]) => void
}
export function init (props: Props) {
  const {
    layoutMap,
    cmdObjArr = [],
    onCmdResp = () => {}
  } = props
  const size = reactive({
    width: 0,
    height: 0
  })
  const watchEl = async () => {
    await nextTick()
    const {width, height} = getElSize(boxEl.value as HTMLElement)
    size.width = width
    size.height = height
  }
  watchEl()
  window.addEventListener('resize', watchEl)
  const classComputed = computed(() => {
    return {
    }
  })
  const styleComputed = computed(() => {
    let obj = layoutMap[containerModeStr.value]
    return {
      [`grid-column`]: `span ${obj.w}`,
      height: `${obj.h * GridRowHeight}px`,
      [`grid-row`]: `span ${obj.h}`,
    }
  })
  const stop = addEventListener_queryDataBus(EventsMap_queryDataBus.FetchCmdData, () => {
    return {
      cmdObjArr: cmdObjArr,
      callback: (cmdRespObjArr: cmdRespObj[]) => {
        // console.log('addEventListener_queryDataBus callback cmdRespObjArr :>> ', cmdRespObjArr);
        onCmdResp(cmdRespObjArr);
      }
    }
  })

  return {
    size,
    styleComputed,
    classComputed,
    unmounted: () => {
      window.removeEventListener('resize', watchEl)
      stop()
    }
  }
}