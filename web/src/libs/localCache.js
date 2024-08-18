import localforage from "localforage";

let storageData = {}
let databaseName = 'app'

export async function init(name, defaultValue = {}) {
  databaseName = name
  storageData = await get(name, defaultValue)
  console.log('storageData :>> ', storageData, name, defaultValue);
}


/**
 *  设置storage
 * @param {string} path
 * @param {*} value 
 */
export async function setStorage(path = '', value) {
  const keys = path.split('.')
  let obj = storageData
  for (let i = 0; i < keys.length - 1; i++) {
    const key = keys[i]
    if (obj[key] === undefined) {
      obj[key] = {}
    }
    obj = obj[key]
  }
  obj[keys[keys.length - 1]] = value
  await set(databaseName, JSON.parse(JSON.stringify(storageData)))
}

/* 
  根据节点路径从storageData获取缓存
 */
export function getStorage(path = '', defaultValue) {
  const keys = path.split('.')
  let obj = storageData
  for (let i = 0; i < keys.length; i++) {
    obj = obj[keys[i]]
    if (!obj) {
      return defaultValue
    }
  }
  return obj
}

async function get(key = '', defaultValue) {
  const value = await localforage.getItem(key)
  // console.log('get value :>> ',key, value);
  return value !== null ? value : defaultValue
}

async function set(key = '', value = null) {
  await localforage.setItem(key, value)
  return value
}