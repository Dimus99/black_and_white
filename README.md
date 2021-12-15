<pre>Задача
Разработать программу, выполняющую обработку изображений.
Описание
Обработчик принимает на вход:
    • имя входной директории;
    • имя директории, в которую необходимо сохранять обработанные изображения.
При старте программа считывает изображения из входной директории, переводит в черно-белый цвет и сохраняет. 
Далее обработчик с некоторой периодичностью проверяет входную директорию на наличие новых изображений.
При появлении новых изображений обработчик также переводит их в черно-белый цвет и сохраняет.
Помимо обработки изображений, программа должна сохранять в базу данных следующую информацию:
    • путь до исходного изображения;
    • путь до обработанного изображения;
    • размер изображения;
    • время, за которое была произведена обработка.
Требования и допущения
Программа должна удовлетворять следующим требованиям:
    • необходимо использовать паттерн producer-consumer;
    • необходимо использовать возможность многопоточной обработки изображений.
При разработке программы допускается:
    • использовать готовую библиотеку для перевода изображения в черно-белый цвет;
    • использовать в качестве базы данных любую реляционную базу данных.</pre>
PostgreSQL
start
<pre>
CREATE DATABASE log_entry; 
CREATE TABLE log_entry (path_from varchar(255), path_to varchar(255), size varchar(255), duration varchar(255));
CREATE USER test WITH password 'test';
GRANT SELECT, INSERT ON log_entry TO test;
</pre>