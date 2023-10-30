Сервис предоставляет через REST API доступ хранилищу кодов

# REST API

## Добавить продукт
Добавляет продукт если его нет в базе. Проверяется формат gtin и наличие описания
### Request
`POST /v1/goods`
``````
curl -D - -X POST http://localhost/v1/goods -d '{ "gtin": "04000000000025", "desc": "Молоко 3,5%"}'
``````
### Response
``````
HTTP/1.1 201 Created
Date: Mon, 30 Oct 2023 07:28:25 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OK
``````
