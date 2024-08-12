# Система управления личными финансами
- Основной функционал: Учёт доходов и расходов, категории транзакций.
- Расширение: Визуализация данных (графики), отчёты по периодам, напоминания о 
платежах, интеграция с банковскими API.

  
# Структура проекта
- personal-finance-management/
    - cmd/
        - server/
            - main.go
    - internal/
        - user/
            - user.go
            - repository.go
            - service.go
        - budget/
            - budget.go
            - repository.go
            - service.go
        - transaction/
            - transaction.go
            - repository.go
            - service.go
        - categories/
                
        - http/
            - handler/
                - account_handler.go
                - transaction_handler.go
                - budget_handler.go
            - middleware/
                - auth_middleware.go
            - router.go
        - config/
            - config.go
        - database/
            - database.go
            - migrations/
        - util/
    - pkg/
        - model/
            - user.go
            - budget.go
            - transaction.go
    - docs/
    - .env
    - .gitignore
    - go.mod
    - go.sum


Структура проекта:
* cmd/ (точка входа в приложение)
* internal/ (внутренняя логика приложения)
* pkg/ (переиспользуемые пакеты)
* configs/ (конфигурационные файлы)
* migrations/ (SQL миграции для базы данных)



