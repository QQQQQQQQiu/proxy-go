
/* 
    示例：
    export const EventsMap = {
      FetchCmdData: 'fetchCmdData'
    }
    export const {
      emitEventListener,
      addEventListener,
      removeEventListener,
    } = createEventInstance(EventsMap)
*/

export function createEventInstance(eventsMap = {}) {
  const _mirrorObj = Object.keys(eventsMap).reduce((acc, key) => {
    acc[key] = []
    return acc
  }, {})
  const _listenerFnObj = () => JSON.parse(JSON.stringify(_mirrorObj))
  let listenerFnObj = _listenerFnObj()

  async function emitEventListener(type = '', ...args) {
    let arr = listenerFnObj[type]
    return Promise.all(arr.map(fn => fn(...args)))
  }
  function addEventListener(type = '', fn = () => { }) {
    listenerFnObj[type].push(fn)
    console.log('addEventListener :', type, listenerFnObj[type])
    return ()=>{
      removeEventListener(type, fn)
    }
  }
  function removeEventListener(type = '', fn) {
    if (!type) {
      listenerFnObj = _listenerFnObj()
      console.log('removeEventListener listenerFnObj :', listenerFnObj)
      return
    }
    if (!fn) {
      listenerFnObj[type] = []
      return
    }
    let index = listenerFnObj[type].findIndex(targetFn => targetFn === fn)
    index >= 0 && listenerFnObj[type].splice(index, 1)
  }

  return {
    emitEventListener,
    addEventListener,
    removeEventListener,
  }
}
