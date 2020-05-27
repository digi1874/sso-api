# Simple Single Sign-On for GO
> 开放登录系统后端
>
> 这是一个开放式登录系统，可供给任何网站做登录用。
>
> 系统安全可靠，数据传输安全可靠，JWT验证安全可靠。使用非对称加密后发给数据，不怕传输数据被拦截，服务器临时生成密钥，任何人都没法查看密钥。JWT的密钥也是服务器运行时随机生成。
>
> * [前端项目](https://github.com/digi1874/sso-web)
> #

## 构建
> 依赖
> 1. go 1.13+
> 2. mysql (4.1+，本项目开发时使用5.7；库需要设置为utf8mb4)
> 3. 项目根目录下创建文件db.json连接mysql数据库
>> ./db.json
>> ```
>> {
>>   "user": "用户名",
>>   "password": "密码",
>>   "localhost": "地址",
>>   "databaseName": "库名"
>> }
>> ```
>> #
> 4. 添加纯真ip信息库qqwry.dat文件，来源[纯真](https://github.com/freshcn/qqwry)
> ```
> # 开发，开启http://localhost:8021/
> $ go run main.go -env=dev
>
> # 生产程序
> $ go build
> ```
> #

## 接口说明
> * response status code:
>> 1. 200: 确定
>> 2. 400: 错误
>> 3. 401: 无权限，token无效
> * response data msg: 回应说明
> #
> 1. 获取公钥
>> * url：/login/rsa
>> * method: GET
>> * response data: {"data": "jwtPayload.jwtSignature"}
>> 1. jwt secret: Response Headers Date 的 Unix 值 （js: new Date( Wed, 27 May 2020 11:22:10 GMT) / 1000）
>> 2. 公钥存放在 jwtPayload pub；[例子](https://github.com/digi1874/sso-web/blob/master/src/api/rsa.js)
>> ```
>> // get: http://localhost:8021/login/rsa
>> // response data:
>> {
>> "data":"eyJwdWIiOiJMUzB0TFMxQ1JVZEpUaUJRVlVKTVNVTWdTMFZaTFMwdExTMEtUVWxIWmsxQk1FZERVM0ZIVTBsaU0wUlJSVUpCVVZWQlFUUkhUa0ZFUTBKcFVVdENaMUZFTTFGck9IbDNVR1ZOTDFsbVdHdElMMWx0YWxseVNtRTRSZ3BNTWtjNWVUUmhlVE4yTW1GU1VIRlJlVUZhWjBzM1IyWjBTMU5pT0hwMFdsQkhUbFJvUTFOSVduTmxZbk51U1RkcE0yTk1RV1o0UmxONkswTjNOWFJtQ201MVowTlRXbVJRVG1aNVZIQlllVU4yTUhaT1RVOUNkamRJZFVWcFIyMXFSRmhYY1dsUlIzTkxaVGh3TVdOdFpVMXZZa2RwVFdaTVRGTnpVRU5LTTFjS2RFcDZUamxLVm5CbFMyWmtSbTk0VFRKUlNVUkJVVUZDQ2kwdExTMHRSVTVFSUZCVlFreEpReUJMUlZrdExTMHRMUW89In0.duHJ2lDNFj15k-Ydq1HvMQMNcmiI7GnR2h8A_3Ez3VY"
>> }
>> ```
>> #

> 2. 检查账号是否已存在
>> * url： /login/number/:number/exist
>> * method: GET
>> * response data: {"data":Boolean(true:存在；false:不存在)}
>> ```
>> // get: http://localhost:8021/login/number/123123/exist
>> // response data:
>> {
>>   "data":true
>> }
>> ```
>> #

> 3. 注册账号
>> * url：/login/register
>> * method: POST
>> * Request Content-Type: application/json;charset=UTF-8
>> * Request Payload: { data: "账号密码加密串", host: "登录网站" }
>> 1. rsa加密[例子](https://github.com/digi1874/sso-web/blob/master/src/utils/crypto.js)
>> 2. rsaEncrypt({ number: 账号, password: 密码 }, 公钥)
>> * response data: {"data":"jwtPayload.jwtSignature"}
>> 1. jwtPayload 使用base64解码后得到 { id: 用户, ip: 用户ip地址, exp: 有效时间unix值 }
>> ```
>> // post: http://localhost:8021/login/register
>>
>> //Request data:
>> {
>>   data: "p4WV+1flEf+r2ko+8g7rNdGctkxgV90yUw9sVCCLQFR7wsWUdOM4oOhdt08lCdajYvPOjS1/sSBa9gp7RSynIjsc8l2zYWhL75WCVvA0A49GRI8nyr9y3944H7yN3wSA6AODHR/sE7Bdis0cMC7FFh/1DmRwmUEO9alrWAPtDY0="
>>   host: "www.ys1994.nl"
>> }
>>
>> // response data:
>> {
>>   "data":"eyJleHAiOjE1ODQ3ODcyNTUsImlkIjoxNCwiaXAiOiIxMjcuMC4wLjEifQ.SbBU8drwIeNuyViPaqnDqGXwipGSkaQq63LOwoQLVOw"
>> }
>> ```
>> #

> 3. 登录
>> * url：/login/
>> 1. 注意别少了最后的"/"
>> * method: POST
>> * Request Content-Type: application/json;charset=UTF-8
>> * Request Payload: { data: "账号密码加密串", host: "登录网站" }
>> 1. rsa加密[例子](https://github.com/digi1874/sso-web/blob/master/src/utils/crypto.js)
>> 2. rsaEncrypt({ number: 账号, password: 密码 }, 公钥)
>> * response data: {"data":"jwtPayload.jwtSignature"}
>> ```
>> // post: http://localhost:8021/login/
>>
>> //Request data:
>> {
>>   data: "p4WV+1flEf+r2ko+8g7rNdGctkxgV90yUw9sVCCLQFR7wsWUdOM4oOhdt08lCdajYvPOjS1/sSBa9gp7RSynIjsc8l2zYWhL75WCVvA0A49GRI8nyr9y3944H7yN3wSA6AODHR/sE7Bdis0cMC7FFh/1DmRwmUEO9alrWAPtDY0="
>>   host: "www.ys1994.nl"
>> }
>>
>> // response data:
>> {
>>   "data":"eyJleHAiOjE1ODQ3ODcyNTUsImlkIjoxNCwiaXAiOiIxMjcuMC4wLjEifQ.SbBU8drwIeNuyViPaqnDqGXwipGSkaQq63LOwoQLVOw"
>> }
>> ```
>> #

> 4. 修改密码
>> * url：/login/password
>> * method: POST
>> * Request Content-Type: application/json;charset=UTF-8
>> * Request Payload: { data: "账号密码加密串" }
>> 1. rsa加密[例子](https://github.com/digi1874/sso-web/blob/master/src/utils/crypto.js)
>> 2. rsaEncrypt({ signature: jwtSignature, password: 当前密码, newPassword: 新密码 }, 公钥)
>> * response data: {"data":"密码修改成功"}
>> ```
>> // post: http://localhost:8021/password
>>
>> //Request data:
>> {
>>   data: "Z37MlUjJAOtL1EdtvWXHjfx69/3g79dAxXgxUnymw5xiDluJ0FHOiGCucWFDg+MTZEHTPdJudzfKvl6liNGqJxar4f0k+Er49Az0yrKDZxmcXGeC8bF+W88F+N0yKSSKfh50QYQTM0D30OA2ZqLQEVEBjKDoBEtL/WU7lt8N2t8="
>> }
>>
>> // response data:
>> {
>>   "data":"密码修改成功"
>> }
>> ```
>> #

> 5. 检查token(jwtPayload.jwtSignature)是否有效
>> * url：/login/verify/:token
>> * method: GET
>> * response data: {"data":Boolean(true:有效；false:无效)}
>> ```
>> // get: http://localhost:8021/login/verify/eyJleHAiOjE1ODQ3NzQyNDMsImlkIjo1LCJpcCI6IjEyNy4wLjAuMSJ9.yaedlsXNNLFxYGknWAPU-ncpUS936V5AigJPWfS3ZxY
>> // response data:
>> {
>>   "data":true
>> }
>> ```
>> #

> 6. 获取登录列表
>> * url：/login/list/:token(jwtPayload.jwtSignature)
>> * method: GET
>> * query:
>> 1. page: 页数
>> 1. size: 每页数量
>> * response data: {"data":Boolean(true:有效；false:无效)}
>> ```
>> // get: http://localhost:8021/login/verify/eyJleHAiOjE1ODQ3NzQyNDMsImlkIjo1LCJpcCI6IjEyNy4wLjAuMSJ9.yaedlsXNNLFxYGknWAPU-ncpUS936V5AigJPWfS3ZxY
>> // response data:
>> {
>>   "data":{
>>     "count": 2,        // 总数
>>     "page": 1,         // 当前页数
>>     "size": 20,        // 每页数量
>>     "data": [          // 列表
>>       {
>>         "ip": "119.129.224.186",       // 用户ip
>>         "country": "广东省广州市",      // 用户地区
>>         "exp": 1584690154,             // 登录过期时间
>>         "message": "",                 // 一些信息
>>         "state": 1,                    // 1: 正常；2: 退出
>>         "createdTime": 1584085354    // 登录时间
>>       },
>>       {
>>         "ip": "119.129.224.186",
>>         "country": "广东省广州市",
>>         "exp": 0,
>>         "message": "登录密码错误",
>>         "state": 2,
>>         "createdTime": 1584085354    // 登录时间
>>       }
>>     ]
>>   }
>> }
>> ```
>> #
