# RedSpear

RedSpear is simple Golang server for parsing markdown.

## demo
### Form data
request
```bash
curl -X POST http://153.126.139.150:8080/api/md -d 'md=hoge'
```
response
```html
<p>hoge</p>
```

### JSON data
request
```bash
curl -X POST http://153.126.139.150:8080/api/md -H 'Content-Type: application/json' -d '{"md":"hoge"}'
```
response
```html
<p>hoge</p>
```

## License
Copyright (c) 2019 Mstn
Released under the [MIT License](https://opensource.org/licenses/mit-license.php)
