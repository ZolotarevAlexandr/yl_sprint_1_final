# Calculation Server

## Description

GitHub repository contains 2 Go modules: `calculator`, which contains all logic
related to calculation and `calculator_server`, which handles requests

Server is running on localhost on `8080` port.  
Server has one endpoint: `/api/v1/calculate`, which receives `POST` requests with JSON
of following schema:

```json
{
  "type": "object",
  "properties": {
    "expression": {
      "type": "string"
    }
  },
  "required": [
    "expression"
  ]
}
```

example:

```json
{
  "expression": "2 + 2"
}
```

Endpoint returns either `result` or `error` with JSONs of following schemas:

On **success**:

```json
{
  "type": "object",
  "properties": {
    "result": {
      "type": "number"
    }
  }
}
```

example:

```json
{
  "result": 4
}
```

On **error**:

```json
{
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    }
  }
}
```

example:

```json
{
  "error": "Expression is not valid"
}
```

## Run server

To start server just run:

```bash
go run ./calculator_server/cmd/server/main.go
```

from root of project

## Examples:

1. Request:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"expression":"3+5*2"}' http://localhost:8080/api/v1/calculate
    ```
   Returns:
    ```json
    {"result":13}
    ```
   With code `200`


2. Request:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"expression":"(4+7)/2"}' http://localhost:8080/api/v1/calculate
    ```
   Returns:
    ```json
    {"result":5.5}
    ```
   With code `200`


3. Request:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"expression":"(((4+7/2))"}' http://localhost:8080/api/v1/calculate
    ```
   Returns:
    ```json
    {"error": "Expression is not valid"}
    ```
   With code `422`, because of unbalanced parentheses


4. Request:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"expression":"1 / 0"}' http://localhost:8080/api/v1/calculate
    ```
   Returns:
    ```json
    {"error": "Expression is not valid"}
    ```
   With code `422`, because of division by zero

On errors not related to calculation (i.e. error while reading request body) server returns
```json
{
    "error": "Internal server error"
}
```
With code `500`
