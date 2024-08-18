export type stroageCache = {
  settings: {
    loopTime: number
  }
}
export type cmdObj = {
  id: string
  cmd: string
}
export type cmdRespObj = {
  id: string
  output: string
}

/* 
  卡片尺寸
  w: 宽度等级： 1-12
  h: 高度等级：(n * 12) px
 */
export type layoutItem = {
  w: number
  h: number
}
export type layoutMap = {
  small: layoutItem
  medium: layoutItem
  large: layoutItem
}