# DBOTechnicalTest-Golang

1. Setup .env file with your local database configuration
2. Please create new database with name from environtment ```DB_SCHEMA```
3. Run command ```go run main.go``` or ```docker-compose up -d --build``` note : this will be automigrate table to your
   connected database
4. Exampe Request & Response

- API Customer

```
DESCRIPTION     Create Customer
METHOD          POST 
URL             127.0.0.1:8008/v1/customer
BODY
                {
                    "full_name": "Caca Handika",
                    "email": "caca@gmail.com",
                    "password": "123456",
                    "address": "Jakarta Barat"
                }
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "full_name": "Caca Handika",
                        "email": "caca@gmail.com",
                        "password": "$2a$10$5a7l1JHsLaRdAr7wLHLEhuOWGpejuKIFzZODSFTA8SUNXWaw2BgTu",
                        "address": "Jakarta Barat",
                        "created_at": "",
                        "updated_at": "",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
                
                
DESCRIPTION     Get All Customer
METHOD          GET 
URL             127.0.0.1:8008/v1/customer/all?name=budi
HEADER          Bearer {token}
RESPONSE
               {
                   "data": [
                       {
                           "id": 2,
                           "full_name": "Budi Santoso",
                           "email": "budi@gmail.com",
                           "password": "",
                           "address": "Jakarta Barat",
                           "created_at": "2023-04-18 14:19:53",
                           "updated_at": "",
                           "deleted_at": ""
                       }
                   ],
                   "err_response": null,
                   "message": "success",
                   "meta": {
                       "current_page": 1,
                       "last_page": 1,
                       "total": 1,
                       "per_page": 100
                   },
                   "status": true
               }
                
                
DESCRIPTION     Get Customer Details
METHOD          GET
URL             127.0.0.1:8008/v1/customer/{customerID}
HEADER          Bearer {token}
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "full_name": "Caca Handika",
                        "email": "caca@gmail.com",
                        "password": "$2a$10$5a7l1JHsLaRdAr7wLHLEhuOWGpejuKIFzZODSFTA8SUNXWaw2BgTu",
                        "address": "Jakarta Barat",
                        "created_at": "2023-04-18 14:21:02",
                        "updated_at": "",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
                
                
DESCRIPTION     Update Customer
METHOD          PUT
URL             127.0.0.1:8008/v1/customer
HEADER          Bearer {token}
BODY
                {
                    "id": 3,
                    "full_name": "Caca Handika (update)",
                    "email": "cacaupdate@gmail.com",
                    "password": "123456update",
                    "address": "Jakarta Barat Update"
                }
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "full_name": "Caca Handika (update)",
                        "email": "cacaupdate@gmail.com",
                        "password": "$2a$10$5a7l1JHsLaRdAr7wLHLEhuOWGpejuKIFzZODSFTA8SUNXWaw2BgTu",
                        "address": "Jakarta Barat Update",
                        "created_at": "2023-04-18 14:21:02",
                        "updated_at": "2023-04-18 07:35:29",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }               
                
                
DESCRIPTION     Delete Customer
METHOD          DELETE 
URL             127.0.0.1:8008/v1/customer/{customerID}
HEADER          Bearer {token}
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "full_name": "Caca Handika (update)",
                        "email": "cacaupdate@gmail.com",
                        "password": "$2a$10$5a7l1JHsLaRdAr7wLHLEhuOWGpejuKIFzZODSFTA8SUNXWaw2BgTu",
                        "address": "Jakarta Barat Update",
                        "created_at": "2023-04-18 14:21:02",
                        "updated_at": "2023-04-18 07:35:29",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
```

- API Order

```
DESCRIPTION     Create Order
METHOD          POST 
URL             127.0.0.1:8008/v1/order
BODY
                {
                    "customer_id": 2,
                    "item_name": "Gelas",
                    "quantity": 20,
                    "price": 8000
                }
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "customer_id": 2,
                        "item_name": "Gelas",
                        "quantity": 20,
                        "price": 8000,
                        "created_at": "",
                        "updated_at": "",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
                
                
DESCRIPTION     Get All Order
METHOD          GET 
URL             127.0.0.1:8008/v1/order/all?page=1&limit=2
HEADER          Bearer {token}
RESPONSE
                {
                    "data": [
                        {
                            "id": 3,
                            "customer_id": 2,
                            "item_name": "Sendok",
                            "quantity": 50,
                            "price": 5000,
                            "created_at": "2023-04-18 14:54:38",
                            "updated_at": "",
                            "deleted_at": ""
                        },
                        {
                            "id": 2,
                            "customer_id": 2,
                            "item_name": "Piring",
                            "quantity": 25,
                            "price": 8500,
                            "created_at": "2023-04-18 14:54:26",
                            "updated_at": "",
                            "deleted_at": ""
                        }
                    ],
                    "err_response": null,
                    "message": "success",
                    "meta": {
                        "current_page": 1,
                        "last_page": 2,
                        "total": 3,
                        "per_page": 2
                    },
                    "status": true
                }
                
                
DESCRIPTION     Get Order Details
METHOD          GET
URL             127.0.0.1:8008/v1/order/{orderID}
HEADER          Bearer {token}
RESPONSE
                {
                    "data": {
                        "id": 2,
                        "customer_id": 2,
                        "item_name": "Piring",
                        "quantity": 25,
                        "price": 8500,
                        "created_at": "2023-04-18 14:54:26",
                        "updated_at": "",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
                
                
DESCRIPTION     Update Order
METHOD          PUT
URL             127.0.0.1:8008/v1/order
HEADER          Bearer {token}
BODY
                {
                    "id": 3,
                    "customer_id": 2,
                    "item_name": "Sendok Update",
                    "quantity": 5099,
                    "price": 99999
                }
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "customer_id": 2,
                        "item_name": "Sendok Update",
                        "quantity": 5099,
                        "price": 99999,
                        "created_at": "2023-04-18 14:54:38",
                        "updated_at": "2023-04-18 07:58:52",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }          
                
                
DESCRIPTION     Delete Order
METHOD          DELETE 
URL             127.0.0.1:8008/v1/order/{customerID}
HEADER          Bearer {token}
RESPONSE
                {
                    "data": {
                        "id": 3,
                        "full_name": "Caca Handika (update)",
                        "email": "cacaupdate@gmail.com",
                        "password": "$2a$10$5a7l1JHsLaRdAr7wLHLEhuOWGpejuKIFzZODSFTA8SUNXWaw2BgTu",
                        "address": "Jakarta Barat Update",
                        "created_at": "2023-04-18 14:21:02",
                        "updated_at": "2023-04-18 07:35:29",
                        "deleted_at": ""
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
```

- API Authentication

```
DESCRIPTION     Customer Login
METHOD          POST 
URL             127.0.0.1:8008/v1/auth/login
BODY
                {
                    "email": "caca@gmail.com",
                    "password": "123456"
                }
RESPONSE
                {
                    "data": {
                        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE4NDEwMDIsImlkIjozfQ.gMoqIbXIq-dxdUvZr7BA7Dce9rtVtQY9cAU42EgF1NY",
                        "customer": {
                            "id": 3,
                            "full_name": "Caca Handika (update)",
                            "email": "caca@gmail.com",
                            "password": "",
                            "address": "Jakarta Barat Update",
                            "created_at": "2023-04-18 14:21:02",
                            "updated_at": "2023-04-18 07:35:29",
                            "deleted_at": ""
                        }
                    },
                    "err_response": null,
                    "message": "success",
                    "meta": null,
                    "status": true
                }
```

- Another Example

```
DESCRIPTION     Protected API Without Bearer Token
METHOD          GET
URL             127.0.0.1:8008/v1/order/all
RESPONSE
                {
                    "data": null,
                    "err_response": "missing authentication header",
                    "message": "unauthenticate",
                    "meta": null,
                    "status": false
                }
```

5. Also you can check in directory screenshot for the API call image & database diagram image