# Payment Processor

## Описание
Этот проект представляет собой систему обработки платежей на языке Go с поддержкой нескольких платёжных методов:
- **PayPal**
- **Stripe**
- **Банковский перевод**

Каждый платёжный процессор реализует общий интерфейс `PaymentProcessor`, который содержит метод `ProcessPayment`. Система логирует все операции в файл `payment.log` и одновременно выводит в консоль.

## Установка и запуск
### 1. Склонируйте репозиторий
```sh
git clone https://github.com/yourusername/payment-processor.git
cd payment-processor
```

### 2. Запустите приложение
```sh
go run main.go
```

## Структура кода
- `PaymentProcessor` – интерфейс для всех платёжных систем.
- `PayPalProcessor`, `StripeProcessor`, `BankTransferProcessor` – реализации интерфейса для разных платёжных систем.
- `Log()` – функция создания логгера, записывающего в файл `payment.log` и консоль.
- `MakePayment()` – функция, принимающая процессор платежей и выполняющая транзакцию.
- `main()` – точка входа в приложение, тестирующая оплату через три метода.

## Пример использования
```go
logger, logFile := Log()
defer logFile.Close()

paypal := NewPayPalProcessor(logger)
stripe := NewStripeProcessor(logger)
bank := NewBankTransferProcessor(logger)

paypalDetails := map[string]string{"email": "test.99@mail.ru"}
stripeDetails := map[string]string{"cardNumber": "1232131231231643"}
bankDetails := map[string]string{"accountNumber": "923542523", "routingNumber": "1672734534"}

MakePayment(logger, paypal, 100.0, "USD", paypalDetails)
MakePayment(logger, stripe, 200.0, "EUR", stripeDetails)
MakePayment(logger, bank, 300.0, "RUB", bankDetails)
```

## Логирование
Все операции логируются в файл `payment.log`. Пример содержимого лога:
```
2025/02/08 14:25:30.123456 Processing payment via PayPal... Amount: 100.00 USD, Email: test.99@mail.ru
2025/02/08 14:25:32.654321 PayPal payment successful: 100.00 USD to test.99@mail.ru
```

## Возможные ошибки
- Отсутствие обязательных данных (email для PayPal, номер карты для Stripe, реквизиты для банка).
- Случайные сбои при обработке платежей (моделируются с вероятностью ошибки).

## TODO
- [ ] Добавить поддержку новых платёжных систем (например, криптовалюты).
- [ ] Реализовать API для взаимодействия с клиентскими приложениями.
- [ ] Улучшить обработку ошибок и добавить механизмы повторных попыток оплаты.

## Лицензия
MIT License

