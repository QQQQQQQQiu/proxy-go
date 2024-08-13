# 简介

***代理请求、命令行***

### 进程参数

| 参数名 | 类型 | 描述 | 是否必须 |
| ------ | ---- | ---- | ---- |
| -s | string | 接口请求令牌 | 是 |

### 服务端口
```:801 ```

### 接口列表

| 路径  |描述 | 
| --- | ---|
| /api |  代理接口 |
| /doc |  文档说明 |

### 代理接口请求方式

1. POST

2. GET（参数保存在链接上的q字段,如/q=${...body}）

### 接口令牌存放(二选一)
1. 请求头
	```
    Authorization:${secret}
    ```
2. 请求参数
	```
    {
		s:${secret}
        ...
    }
    ```

    
### 请求参数

  | 参数名 | 类型 | 描述 | 是否必须 |
  | ------ | ---- | ---- | ---- |
  | s | string | 密钥 | 是 |
  | type | string | 请求类型 | 是 |
  | data | any | 请求数据 | 是 |

### 参数示例

* 执行命令

	```
        {
            "s": "${secret}",
            "type": "command",
            "data": [
                {
                    "id": "c1",
                    "sh": "pwd"
                },
                {
                    "id": "c2",
                    "sh": "ls"
                }
            ]
        }
    ```
    响应：[{id: 'c1', output: ''},{id: 'c2', output: ''}]
        

* 接口代理 

	```

        {
            "s": "${secret}",
            "type": "xhr",
            "data": {
                "throwHeaders": false, // 是否把响应头放进body
                "method": "GET",
                "url": "https://cn.bing.com/hp/api/model",
                "headers": {
                    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
                },
                body: ""
            }
        }
	```