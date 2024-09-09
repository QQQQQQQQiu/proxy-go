import type { cmdObj, cmdRespObj } from "@/types/index.d.ts";
import { getStorage } from "@/libs/localCache.js"
import {wait} from '@/utils/func';
import {myFetch} from '@/libs/fetch.js';
import {createEventInstance} from '@/libs/event.js';

type fetchCmdDataArgs = {
  cmdObjArr: cmdObj[]
  callback: (res: cmdRespObj[]) => void
}

export const EventsMap = {
  FetchCmdData: 'FetchCmdData',
  LoopQueryData: 'LoopQueryData'
}

export const {
  emitEventListener,
  addEventListener,
  removeEventListener,
} = createEventInstance(EventsMap)

export const loopQueryConroller = {
  once: () => {
    loopQueryConroller.stop()
    loopQuery(true)
  },
  start: () => {
    loopQueryConroller.stop()
    loopQuery()
  },
  stop: () => StopLoopHandle()
}

export async function proxyFetch (options = {}) {
  return await myFetch(`http://10.0.0.116:801/api/xhr/xx`, {
    contentType: 'json',
    responseType: 'text',
    method: 'POST',
    data: options
  })
}

async function fetchCmdData() {
  let objArr: fetchCmdDataArgs[] = await emitEventListener(EventsMap.FetchCmdData)
  // console.log('objArr :>> ', objArr);
  let cmdQueueArr: cmdObj[] = []
  cmdQueueArr = objArr.map(item => item.cmdObjArr).flat()
  // console.log('cmdQueueArr :>> ', cmdQueueArr);
  const res = await myFetch(`http://10.0.0.116:801/api/cmd/xx`,{
    method: 'POST',
    contentType: 'json',
    responseType: 'json',
    headers: {},
    data: cmdQueueArr,
  })
  // console.log('res :>> ', res);
  for (let index = 0; index < objArr.length; index++) {
    const element = objArr[index];
    // element.callback([{id: 'cmd1', output: 'lssss'}, {id: 'cmd2', output: 'pwdddd'}, {id: 'cmd3', output: 'pwdddd'}])
    element.callback(res)
  }
}
 let StopLoopHandle = ()=>{}
 async function loopQuery(isOnce = false) {
   await fetchCmdData()
   emitEventListener(EventsMap.LoopQueryData)
   if (isOnce) return
   const waitTime = getStorage('settings.loopTime')
   console.log('waitTime :>> ', waitTime);
  await wait(waitTime, (reject) => StopLoopHandle = reject)
  await loopQuery()
}

