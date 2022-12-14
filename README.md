# Куриная ферма
Решение задачи про ферму


# Quick start #

Скопируйте репозиторий в указанную вами директорию и выполните в ней команды:

` docker-compose build `

` docker-compose up -d `

Конфигурационный файл /config/app.env:

+ CHICKS_COUNT - количество куриц
+ EGGS_MIN_SPAWN_COUNT - минимальное количество снесенных яиц за раз
+ EGGS_MAX_SPAWN_COUNT - максимальное количество снесенных яиц за раз
+ EGGS_SPAWN_MIN_DELAY - минимальная задержка перед снесением яйца
+ EGGS_SPAWN_MAX_DELAY - максимальная задержка перед снесением яйца
+ MIN_CHECK_DELAY - минимальное время интервала в которое приходит фермер
+ MAX_CHECK_DELAY - максимальное время интервала через которое приходит фермер
+ MIN_NEEDED_QUANTITY - мин количество яиц, которые нужны фермеру
+ MAX_NEEDED_QUANTITY - макс количество яиц, которые нужны фермеру

Информация по количеству яиц в холодильнике
http://127.0.0.1:8081/

Что происходит в программе:
Курицы - это n отдельных горутин, которые изменяют один общий ресурс - количество яиц в холодильнике.
В каждой горутине в бесконечном цикле осуществляется процесс отсидки курицей своих яиц. С определённой задержкой i-ая курица откладывает яйцо в в задаваемом целочисленном интервале (об этом сообщается в логах контейнера). 

В отдельной горутине запущен процесс проверки фермером отложенных его курицами яиц. Об этом также можно получить сообщение в логах 

Когда количество яиц в холодильнике доходит до максимального значения Int64, куриные горутины завершают свою работу. А одинокий фермер продолжает собирать это огромное количество яиц


# Описание задачи #
На ферме живут курицы. Известно, что курица несет от 2 до 5 яиц каждые 2-10 секунд (курицы несут яйца независимо друг от друга). Все снесенные яйца складываются в холодильник. И, через 10-15 секунд приходит фермер за определенным количеством яиц (случайное число от 10-20). Если в холодильнике достаточное количество яиц, он забирает необходимое, иначе забирает все яйца из холодильника, если их было меньше.

Входные параметры
Диапазоны (количества яиц и временных промежутков) задаются при старте;
Число из диапазонов выбирается случайно во время выполнения программы.
Задание
Необходимо реализовать программу эмулирующую алгоритм взаимодействия, описанный в условии. Дополнительно сделать http-handler для получения текущего количества яиц в холодильнике.

Ожидается, что вместе с проектом будет приложен README с описанием как запустить и проверить работоспособность приложения;
Расположить код в виде оформленного github-репозитория;
Примерное время выполнения задания: 2 часа.
Преимуществом будет
Разумное использование сторонних библиотек;
Логирование;
Unit-тесты;
Все константы, настраиваемые через конфиг;
Оформленный Makefile;
Сборка Docker-образа с приложением;
Логическое разделение на коммиты.
