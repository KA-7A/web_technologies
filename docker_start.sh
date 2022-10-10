docker pull nginx   # Скачиваем образ

# docker images     # Показать образы
# docker run -it --rm -d -p 8080:80 -v ./src/:/usr/share/nginx/html --name web nginx    # Первая итерация запуска контейнера
# Генерирует ssl-сертификат для работы http/2.0
#
#       sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./ssl/localhost.key -out ./ssl/localhost.crt
#
# Вводим информацию (хотя можно ничего не писать, один фиг он самоподписанный и браузер будет ругаться)
#
#       sudo openssl dhparam -out ./ssl/dhparam.pem 2048 2>/dev/null
#
# Добавим Диффи-Хелмана (доооолго ждем)
# Строки сверху выполнять необязательно, потому что все ключи уже сгенерированы

docker-compose up -d # запуск контейнера с nginx в detached режиме

docker ps           # Посмотрим тут ID запущенного контейнера
docker logs webserver -f   # Будем получить логи в постоянном режиме. 
# После пары тестовых (GET/POST/DELETE) запросов через PostMan получил такие логи:
# 172.18.0.1 - - [29/Sep/2022:07:41:14 +0000] "GET /first_page.html HTTP/1.1" 200 3126 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.114 YaBrowser/22.9.1.1079 Yowser/2.5 Safari/537.36" "-"
# 172.18.0.1 - - [29/Sep/2022:07:42:39 +0000] "POST / HTTP/1.1" 403 153 "-" "PostmanRuntime/7.29.2" "-"
# 172.18.0.1 - - [29/Sep/2022:07:42:52 +0000] "DELETE / HTTP/1.1" 405 157 "-" "PostmanRuntime/7.29.2" "-"

# А тут мы с помощью curl посмотрим, получается ли сделать HTTP/2.0 или нет (из другого терминала)
curl -ks -I https://localhost/index.html -o/dev/null -w '%{http_version}\n'
# Выхлоп тут: 2
# Лог на сервере:
# 172.18.0.1 - - [29/Sep/2022:10:38:14 +0000] "HEAD /index.html HTTP/2.0" 200 0 "-" "curl/7.79.1"
# Ура, HTTP/2.0 работает

# С помощью утилиты nghttp проверяем, что у нас получилось с пушами:
nghttp -ans https://127.0.0.1
# Получаем такой результат
# id  responseEnd requestStart  process code size request path
#  13     +7.37ms       +709us   6.66ms  200   1K /
#   2     +9.86ms *    +5.23ms   4.63ms  200  541 /styles/style.css
#   4    +12.10ms *    +5.24ms   6.86ms  200  652 /script.js
#  15    +14.59ms      +7.48ms   7.10ms  200   1K /favicon.ico
# Где звёздочка означает, что файл был отправлен через HTTP/2

docker stop webserver      # Остановка работающего контейнера 

docker container prune # Удалить все завершенные контейнеры, чтобы глаза не мозолили