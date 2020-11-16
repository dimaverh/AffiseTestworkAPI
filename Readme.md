# Pricegen API

## Общее описание

API предназначен для парсинга данных с внешних апи.

## API

* getData

### Получить данные

Метод: __GET__
URI: `/`
Пример запроса:

```bash
  curl --header "Content-Type: application/json" --request POST --data '{"data":["https://reqres.in/api/products/3",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/",
      "http://date.jsontest.com/"
      ]}' \
    http://localhost:8080/.
```

Перед запуском проекта нужно сделать make configure и make compile.