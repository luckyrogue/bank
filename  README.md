# Bank API

Этот проект предоставляет API для управления банковскими аккаунтами, включая создание аккаунтов, просмотр информации о них, дебет, кредит, перевод средств и заморозку аккаунтов.

## Технологии

- **Go (Golang)** — основной язык для написания серверной части.
- **Mux** — маршрутизация для API.
- **Docker** — для контейнеризации и простоты развертывания.

## Функциональность

- Создание банковских аккаунтов (текущие и сберегательные счета).
- Просмотр информации об аккаунтах.
- Дебет и кредит средств.
- Заморозка (блокировка) счетов.
- Перевод средств между аккаунтами.

## Установка и запуск

### 1. Запуск приложения

```bash
docker build -t bank .
docker run -p 8080:8080 bank
```

## Коллекция Postman

```json
{
	"info": {
		"_postman_id": "ed30f3d9-3091-41e1-b9dd-c917d8835c28",
		"name": "Bank API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "27997264"
	},
	"item": [
		{
			"name": "Accounts",
			"item": [
				{
					"name": "Create",
					"event": [
						{
							"listen": "test",
							"script": {
								"exec": [
									"var response = pm.response.json();",
									"",
									"pm.collectionVariables.set(\"account_id\", response.id);",
									"",
									""
								],
								"type": "text/javascript",
								"packages": {}
							}
						},
						{
							"listen": "prerequest",
							"script": {
								"exec": [
									""
								],
								"type": "text/javascript",
								"packages": {}
							}
						}
					],
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"account_type\": \"Current\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{dev}}/accounts",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts"
							]
						}
					},
					"response": []
				},
				{
					"name": "Read",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{dev}}/accounts/{{account_id}}",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts",
								"{{account_id}}"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"account_type\": \"Savings\",\n  \"balance\": 500,\n  \"is_frozen\": false\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{dev}}/accounts/{{account_id}}",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts",
								"{{account_id}}"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "{{dev}}/accounts/{{account_id}}",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts",
								"{{account_id}}"
							]
						}
					},
					"response": []
				},
				{
					"name": "Debit",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"amount\": 100\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{dev}}/accounts/{{account_id}}/debit",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts",
								"{{account_id}}",
								"debit"
							]
						}
					},
					"response": []
				},
				{
					"name": "Credit",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"amount\": 200\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{dev}}/accounts/{{account_id}}/credit",
							"host": [
								"{{dev}}"
							],
							"path": [
								"accounts",
								"{{account_id}}",
								"credit"
							]
						}
					},
					"response": []
				},
				{
					"name": "Freeze",
					"request": {
						"method": "GET",
						"header": []
					},
					"response": []
				}
			]
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					""
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					""
				]
			}
		}
	],
	"variable": [
		{
			"key": "dev",
			"value": "http://localhost:8080",
			"type": "string"
		},
		{
			"key": "account_id",
			"value": ""
		}
	]
}
```
