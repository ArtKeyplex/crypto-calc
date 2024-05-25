# Проект конвертации валют

## Описание

Этот проект предназначен для конвертации валют с использованием API FastForex. Проект состоит из трех контейнеров Docker:
- `db`: Контейнер с базой данных PostgreSQL, где хранятся данные о валютах.
- `app`: Основное приложение, предоставляющее REST API для конвертации валют.
- `updater`: Сервис для периодического обновления курсов валют с использованием команды `update-rates`.

## Установка и запуск

1. Склонируйте репозиторий:

    ```sh
    git clone git@github.com:ArtKeyplex/crypto-calc.git
    cd crypto-calc
    ```

2. Скопируйте файл `.env` в корневой директории проекта из примера:

    ```
    cp .env.example .env
    ```

3. Запустите контейнеры Docker с помощью Makefile:

    ```sh
    make up
    ```

4. Для остановки контейнеров:

    ```sh
    make down
    ```
   
5. Тестовые данные создаются в бд при первом запуске 