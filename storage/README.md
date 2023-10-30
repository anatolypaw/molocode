Сервис предоставляет через REST API доступ хранилищу кодов

# REST API

## Добавить продукт
Добавляет продукт если его нет в базе. Проверяется формат gtin и наличие описания
### Request
`POST /v1/goods`

curl -X POST http://localhost/v1/goods
-d '{"gtin": "04000000000005", "desc": "Молоко 3,5%"}'

### Response
