

### Requirements
* Docker and Go
### Usage
Clone the repository with:
```bash
git clone github.com/SurkovIlya/statistics-app
```
Copy the `env.example` file to a `.env` file.
```bash
cp .env.example .env
```
Update the postgres variables declared in the new `.env` to match your preference. 

Build and start the services with:
```bash
docker-compose up --build
```
### Statistic app API
<details>
<summary> <h4>{statistic-app-api-host}/orderbook/get - Get orders from DB</h4></summary>
  
#### Method: POST
#### Request: 
```json
{
	"exchange_name": "bybit",
	"pair": "USD/RUB"
}
```
#### Response:
```json
[
	{
		"price": 331.4,
		"base_qty": 3.66
	},
	{
		"price": 222.02,
		"base_qty": 5.66
	}
]
```
</details>
<details>
<summary> <h4>{statistic-app-api-host}/orderbook/save - Save order in DB</h4></summary>
  
#### Method: POST
#### Request: 
```json
{
  "exchange_name": "bybit",
  "pair": "USD/RUB",
  "order_book": 
    {
      "asks": 
        {
          "price": 1.40,
          "base_qty": 3.66
        }
      ,
      "bids": 
        {
          "price": 3.02,
          "base_qty": 5.66
        }
    }
}
```
#### Response:
```json
"OK"
```
</details>
<details>
<summary> <h4> {statistic-app-api-host}/orderhistory/get - Get orders history from DB </h4></summary>
  
#### Method: POST
#### Request: 
```json
{
  "client_name": "John Doe",
  "exchange_name": "Example Exchange",
  "label": "Order123",
  "pair": "BTC/USDT"
}
```
#### Response:
```json
[
	{
		"client_name": "John Doe",
		"exchange_name": "Example Exchange",
		"label": "Order123",
		"pair": "BTC/USDT",
		"side": "Buy",
		"type": "Limit",
		"base_qty": 1.5,
		"price": 40000.25,
		"algorithm_name_placed": "AlgorithmXYZ",
		"lowest_sell_prc": 40200.75,
		"highest_buy_prc": 39950.5,
		"commission_quote_qty": 2,
		"time_placed": "2022-01-15T10:30:00Z"
	}
]
```
</details>
<details>
<summary> <h4> {statistic-app-api-host}/orderhistory/save - Save order history in DB</h4> </summary>

#### Method: POST
#### Request: 
```json
{
  "client_name": "John Doe",
  "exchange_name": "Example Exchange",
  "label": "Order123",
  "pair": "BTC/USDT",
  "side": "Buy",
  "type": "Limit",
  "base_qty": 1.5,
  "price": 40000.25,
  "algorithm_name_placed": "AlgorithmXYZ",
  "lowest_sell_prc": 40200.75,
  "highest_buy_prc": 39950.50,
  "commission_quote_qty": 2.0,
  "time_placed": "2022-01-15T10:30:00Z"
}
```
#### Response:
```json
"OK"
```
</details>

#### Load testing
Testing was carried out on a computer with an "AMD Ryzen 5 5600H with Radeon Graphics" processor, processor frequency - 3.30 GHz, 1 core was used (runtime.GOMAXPROCS(1))
![Test result](https://github.com/SurkovIlya/statistics-app/blob/main/loadTest.jpg)
