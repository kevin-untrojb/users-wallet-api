# users-wallet-api


#Ejecución de la api

Para buildear la API localmente  se debe ejecutar desde la raíz del proyecto:

```bash
make build-api
```

Para borrar las dependencias y cerrar ls containers de docker :

```bash
make clean
```

Para ejecutar todos los test :

```bash
make test-all
```
## Endpoints

### Obtener usuario
Se obtiene un usuario y todas sus billeteras
```http
GET/users/:user_id
curl -X GET 'localhost:8080/users/1'
```
ejemplo de response 
```json
{
  "id": 1,
  "first_name": "roberto",
  "last_name": "robertino",
  "alias": "tito",
  "email": "tito12@htmail.com",
  "wallets": [
    {
      "id": 1,
      "currency_name": "ARS",
      "current_balance": "369369370369.00"
    },
    {
      "id": 2,
      "currency_name": "BTC",
      "current_balance": "0.00023124"
    },
    {
      "id": 3,
      "currency_name": "USDT",
      "current_balance": "34.24"
    }
  ]
}
```

### Crear usuario usuario
Ejemplo de request
```http
POST/users
curl -X POST 'localhost:8080/users'
```
body de ejemplo
```json
{
   "first_name":"name",
   "last_name":"last name",
   "alias":"alias",
   "email":"email@email.com"
}
```
body de response
```json
{
  "id": 6,
  "first_name":"name",
  "last_name":"last name",
  "alias":"alias",
  "email":"email@email.com",
  "wallets": null
}
```
### Search de transacciones
Parametros opcionales para el search:
    currency : puede ser "ARS","BTC","USDT"
    tramsaction_type : puede ser "deposit" o "extraction"
    limit : maximo de elementos (10 por default)
    offset: offset de elementos (0 por default)
Ejemplo de request
```http
GET/users/:user_id/wallet
curl -X GET 'localhost:8080/users/1/wallet?currency=ARS&transaction_type=&limit=2&offset=1'
```

ejemplo deresponse
```json
{
  "paging": {
    "total": 6,
    "limit": 2,
    "offset": 1
  },
  "results": [
    {
      "id": 5,
      "transaction_type": "deposit",
      "date_create": "2021-09-20T22:50:25Z",
      "amount": "123123123123.00",
      "currency": "ARS"
    },
    {
      "id": 4,
      "transaction_type": "deposit",
      "date_create": "2021-09-20T22:50:24Z",
      "amount": "123123123123.00",
      "currency": "ARS"
    }
  ]
}
```

### Nueva transacción
Se definió que esta nueva transaccion es para cada billetera.
No se puede quedar en negativo ninguna cuenta
```http
POST/users/:user_id/wallet/:wallet_id/new_transaction
curl -X POST 'localhost:8080/users/1/wallet/1/new_transaction'
```

body 
```json
{
  "transaction_type":"deposit",
  "amount":"100.00"
}
```
ejemplo deresponse
```json
{
  "id": 7,
  "wallet_id": 1,
  "transaction_type": "deposit",
  "user_id": 1,
  "date_create": "0001-01-01T00:00:00Z",
  "amount": "100.00",
  "currency": ""
}
```