## Установка и запуск проекта

### Клонирование репозитория
Для начала работы склонируйте репозиторий в удобную Вам директорию:
```bash
git clone https://github.com/kstsm/wb-level-0
```
### Настройка переменных окружения
Создайте `.env` файл, скопировав в него значения из `.env.example`, и укажите необходимые параметры.
### Запуск проекта
#### Запустите Docker:
```bash
make docker
```
#### Выполните миграции базы данных (при условии, что на вашем ПК установлен `pressly/goose`):
```bash
make goose-up
```
Если goose не установлен, выполните миграции вручную, используя файлы из папки `wb-level-0/consumer/migrations`.

#### Запуск producer и consumer:
```bash
make run-all
```
После запуска сервис будет доступен по адресу: http://localhost:8080/order

### Проведенно нагрузочное тестирование:
<img width="621" height="285" alt="image" src="https://github.com/user-attachments/assets/7a0870a6-5bc4-4e6a-8c72-addfb390a25b" />


