запуск

Для запуска проекта выполните следующие шаги:

Склонируйте репозиторий:
git clone https://github.com/dimakirio/calc-sprint-1.git
cd calc_go
Убедитесь, что Go установлен и находится в $PATH (проверить версию можно командой go version).

Запустите API-сервер командой:

go run .\cmd\main.go
Сервер запустится на порту 8080. Если необходимо изменить порт, установите флаг --port:

go run .\cmd\main.go --port 9090
Использование API
Эндпоинт
POST /api/v1/calculate
Заголовки
Content-Type: application/json
Тело запроса

Пример:

{
  "expression": "2+2*2"
}
Ответы
Успешный запрос

Статус-код: 200 OK

Пример ответа:

{
  "result": "6"
}
Ошибка обработки выражения

Статус-код: 422 Unprocessable Entity

Пример ответа:

{
  "error": "Error calculation"
}
Неподдерживаемый метод

Статус-код: 405 Method Not Allowed

Пример ответа:

{
  "error": "Wrong Method"
}
Некорректное тело запроса

Статус-код: 400 Bad Request

Пример ответа:

{
  "error": "Invalid Body"
}

Примеры использования

Успешный запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:

{
  "result": "6"
}
