# Financial-Market-Microservice
___________________________________

## Financial-Market-Microservice - это микросервис реализованный на Go позволяющий пользователю получить информацию о курсе валют по отношению к рублю и информацию об отношении рубля к валютам.
___________________________________

## APIs:
Для сервиса были разработаны API для получения up-to-date информации о курсе и отношениях валют к рублю. API реализованы как при помощи **RESTful API** архитектуры, так и при помощи **gRPC** архитектуры.
___________________________________

## Документация

Документация реализована при помощи **go-swagger API**, расположенная по ссылке: https://github.com/go-swagger/go-swagger.
В проект добавлен Makefile, выполнение которого создаст файл, содержащий данные для отображения документации. Документация отображается на локальном хосте **(localhost:3000/docs)**. Для того, чтобы выполнить Makefile, в корневой директории проекта пропишите:

     make swagger
     
В качестве результата будет создан файл swagger.yaml необходимый для отображения документации на локальном хосте.
Если **swagger-API** не установлена, то Makefile загрузит его самостоятельно. Результат выполнения Makefile:

     which swagger || GO111MODULE=off go get github.com/go-swagger/go-swagger/cmd/swagger
     /home/ellofae/go/bin/swagger
     GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

Помимо Makefile для **go-swagger API**, добавлен Makefile для генерации **proto** файла для реализации **gRPC** архитектуры.
Чтобы обновить **.proto** файл и внести изменения в проект, в директории ClientServing, которая находится в корневой директории проекта, пропишите:
  make proto

В качестве результата будет создан **.proto** файл необходимый для поддержки **gRPC** архитекруры API.
___________________________________

### Отобржанение главной страницы:
![result1](https://github.com/ellofae/Financial-Market-Microservice/blob/main/imgs/pr1.PNG?raw=true)

### Отображение сервисной страницы Exchange Rates
![result2](https://github.com/ellofae/Financial-Market-Microservice/blob/main/imgs/pr2.PNG?raw=true)

### Отображение сервисной страницы Currency Rates
![result3](https://github.com/ellofae/Financial-Market-Microservice/blob/main/imgs/pr3.PNG?raw=true)

### Отображение страницы Team
![result4](https://github.com/ellofae/Financial-Market-Microservice/blob/main/imgs/pr4.PNG?raw=true)

### Отображение страницы Contact
![result5](https://github.com/ellofae/Financial-Market-Microservice/blob/main/imgs/pr5.PNG?raw=true)
___________________________________
