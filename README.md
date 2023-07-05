# get-nrows-service
Server и client лежат в соответствующих папках.  
Сервер хранит таблицу имен и фамилий в памяти.  
Для получения строчек имен с фамилией, нужно обратится по /rows?n=${сколько имен нужно получить}  
Swagger доступен по адресу localhost:8080/swagger/index.html
# Запус проекта
1. Скопировать в нужную папку
```
git clone https://github.com/rogudator/get-nrows-service.git
```
2. Перейти в папку с проектом
```
cd get-nrows-service
```
3. Забилдить зависимости
```
make build
```
4. Запустить проект
```
make services
```
