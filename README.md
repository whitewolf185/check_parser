# Репозиторий с парсингом чека покупок
## Инструкция
Для того, чтобы пользоваться данным парсером, для начала надо написать путь, где расположен JSON файл с данными по чеку.
Данный JSON файл можно получить, если просканировать QR-код чека в приложении `Проверка чеков ФНС России`.

![img.png](doc/imgs/img.png)

Далее программа попросит путь к файлу. Укажите путь полученного JSON. JSON файл должен находиться в папке `doc/JSONs`

Результат программы будет записан в файл `doc/results/<date>.csv`

При загрузке файла в excel, нужно поставить разделитель _табуляция_.

Таблица вывода будет иметь следующий вид:

| Наименование товара | Цена за единицу | Количество | Сумма |
|---------------------|-----------------|------------|-------|
