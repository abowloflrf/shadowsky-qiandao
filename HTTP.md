## 1. Login Page

```
curl 'https://www.shadowsky.icu/auth/login'
-H 'authority: www.shadowsky.icu'
-H 'pragma: no-cache'
-H 'cache-control: no-cache'
-H 'upgrade-insecure-requests: 1'
-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36'
-H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9'
-H 'sec-fetch-site: none'
-H 'sec-fetch-mode: navigate'
-H 'accept-encoding: gzip, deflate, br'
-H 'accept-language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7'
--compressed
```

## 2. Login Post

```
curl 'https://www.shadowsky.icu/auth/login'
-H 'authority: www.shadowsky.icu'
-H 'accept: application/json, text/javascript, */*; q=0.01'
-H 'x-requested-with: XMLHttpRequest'
-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36'
-H 'content-type: application/x-www-form-urlencoded; charset=UTF-8'
-H 'origin: https://www.shadowsky.icu'
-H 'sec-fetch-site: same-origin'
-H 'sec-fetch-mode: cors'
-H 'referer: https://www.shadowsky.icu/auth/login'
-H 'accept-encoding: gzip, deflate, br'
-H 'accept-language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7'
-H 'cookie: __cfduid=d6218ad105057e1071889f6daf266c3061573298696'
--data 'email=a%40example.com&passwd=123456&remember_me=week'
--compressed
```

**response**

```
Header: set-cookie: sid=b6b1894c612b8cc11bae117cb1e515a47d82353da0f2af7a476a46b87171d447; expires=Sun, 26-Jan-2020 02:10:41 GMT; Max-Age=604800; path=/
```

## 3. Checkin

```
curl 'https://www.shadowsky.icu/user/checkin'
-X POST
-H 'authority: www.shadowsky.icu'
-H 'content-length: 0'
-H 'accept: application/json, text/javascript, */*; q=0.01'
-H 'x-requested-with: XMLHttpRequest'
-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36'
-H 'origin: https://www.shadowsky.icu'
-H 'sec-fetch-site: same-origin'
-H 'sec-fetch-mode: cors'
-H 'referer: https://www.shadowsky.icu/user'
-H 'accept-encoding: gzip, deflate, br'
-H 'accept-language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7'
-H 'cookie: __cfduid=d6218ad105057e1071889f6daf266c3061573298696; sid=b6b1894c612b8cc11bae117cb1e515a47d82353da0f2af7a476a46b87171d447'
--compressed
```

**response**

```json
{ "msg": "获得了 174 MB流量.", "ret": 1 }
```
