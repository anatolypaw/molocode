Сервис предоставляет через REST API доступ хранилищу кодов

## Добавить продукт в базу
Принимает:
- GTIN продукта
- Описание продукта

Возвращает ошибки:
- некорректный gtin
- продукт уже есть в базе
- отсутсвует описание
`POST /v1/goods`
``````
{ "gtin": "04000000000025", "description": "Молоко 3,5%"}
``````

## Получить список продуктов
Выводит список всех продуктов
### Request
`GET /v1/goods`

### Respone
``````
--
``````
