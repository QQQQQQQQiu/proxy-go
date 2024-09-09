export function generateUUID () : string {
  var d = new Date().getTime();
  var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = (d + Math.random()*16)%16 | 0;
    d = Math.floor(d/16);
    return (c=='x' ? r : (r&0x3|0x8)).toString(16);
  });
  return uuid;
}

export function debounce(fn: Function, delay: number) {
  let timerId: any = null;
  return function(...args: any[]) {
    const context = this
    clearTimeout(timerId);
    timerId = setTimeout(() => {
      fn.apply(context, args);
    }, delay);
  };
}


export function downloadTxtFile(text: string, fileName: string = 'file') {
  const blob = new Blob([text], { type: "text/plain" });
  // 创建一个下载链接
  const a = document.createElement("a");
  a.href = URL.createObjectURL(blob);
  a.download = `${fileName}.txt`;

  // 将链接添加到页面并触发点击下载
  document.body.appendChild(a);
  a.click();
  // 清理
  window.URL.revokeObjectURL(a.href);
  document.body.removeChild(a);
}

export function chooseFile(cb: (file: File) => void) {
  const input = document.createElement('input');
  input.type = 'file';
  input.accept = '.txt'; // 指定仅接受TXT文件

  input.addEventListener('change', function() {
    const file = this.files[0];
    if (file) {
      cb(file)
    }
  });
  input.click();
}

export async function wait(
  n: number = 0, 
  exportReject: (reject: (reason?: any) => void) => void = () => {}
): Promise<void> {
  let stop: ((reason?: any) => void) | null = null;
  let pass: (() => void) | null = null;

  const control = new Promise<void>((resolve, reject) => {
    pass = resolve;
    stop = reject
  });
  exportReject(stop);
  
  const timer = new Promise<void>((resolve, reject) => {
    setTimeout(() => {
      stop = null;
      resolve();
      pass && pass(); // 确保 pass 存在后再调用
    }, n);
  });

  
  await Promise.all([control, timer]);
}

export function getElSize (el: HTMLElement) {
  return {
    width: el.clientWidth || 0,
    height: el.clientHeight || 0
  }
}

export function fixNumber(n: number, precision: number) {
  return Number(n.toFixed(precision))
}