const decoder = new TextDecoder('utf-8')

/* 
  * @param {string} url
  * @param {object} options
  * @param {string} options.method
  * @param {string } options.contentType
  * @param {string} options.responseType
  * @param {object} options.headers
  * @param {object} options.data
  * @returns {Promise<any>}
*/
// headers['content-type'] = 'application/json'
// headers['content-type'] = 'application/x-www-form-urlencoded'
// headers['content-type'] = 'text/plain'
export async function myFetch (url, options = {}) {
  const {
    method = 'GET',
    contentType = 'json', // json, form
    responseType = 'text', // json, text, stream, original
    streamOptions = {
      type: 'text', // text, original
      onProgress: () => {},
    },
    headers = {},
    data = {},
  } = options
  let reqBody = data
  let reqUrl = url
  switch (method) {
    case 'GET':
      let p = new URLSearchParams(data).toString()
      if (p) {
        reqUrl += '?' + p
      }
      reqBody = null
      break
    case 'POST':
      if (contentType === 'json') {
        reqBody = JSON.stringify(data)
      } else if (contentType === 'form') {
        reqBody = new URLSearchParams(data).toString()
      } else if (contentType === 'text') {
        reqBody = data
      }
      break
  }
  const res = await fetch(reqUrl, {
    method,
    headers,
    body: reqBody,
  })
  switch (responseType) {
    case 'json':
      return await res.json()
    case 'text':
      return await res.text()
    case 'blob':
      return await res.blob()
    case 'stream':
      let data = await readStream(res, streamOptions.type)
      for await (const chunk of data) {
        streamOptions.onProgress(chunk)
      }
      return res
    case 'original':
      return res
  }
}

/* 
  * @param {Response} resp
  * @param {string} format
  * @returns {AsyncGenerator<string>}
 */
async function * readStream (resp, format = 'text') {
  const reader = resp.body.getReader()
  let partialLine = ''
  while (true) {
    const { done, value } = await reader.read();
    if (done) {
      break;
    }
    if (format === 'text') {
      const decodedText = decoder.decode(value, {stream: true});
      const chunk = partialLine + decodedText
      const newLines = chunk.split(/\r?\n/)
      partialLine = newLines.pop() ?? ''
      for (const line of newLines) {
        yield line
      }
    } else {
      yield value;
    }
  }
}