# AQTechincalTest-Golang
1. Setup .env file with your local database configuration
2. Please create new database with name from environtment ```DB_SCHEMA```
3. Run command ```go run main.go``` note : this will be automigrate table to your connected database
4. Request & Response

```
METHOD  POST 
URL     127.0.0.1:8008/v1/auth/login
BODY
        {
            "username": "angga123",
            "password": "123456"
        }
RESPONSE
        {
            "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkzMjQ0ODcsInVzZXJJRCI6Ijk2NDY1ZjNmLWM2YmMtMTFlZC05MTM4LTAwZmZjMmVmMjY0NSJ9.sjkeTvWzMyTuUC6JfvCU7lVB3CvUN7yXjKAe0K_qXFs",
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```

```
METHOD  POST
URL     127.0.0.1:8008/v1/book
HEADER  Authorization   Bearer {token}
BODY
        {
            "judul": "judul 1",
            "deskripsi": "deskripsi 1",
            "harga": 999.55,
            "stok": 50,
            "penerbit": "penerbit 1",
            "book_category": [
                {
                    "category": "kategori A"
                },
                {
                    "category": "kategori B"
                },
                {
                    "category": "kategori C"
                }
            ],
            "book_keyword": [
                {
                    "keyword": "keyword A"
                },
                {
                    "keyword": "keyword B"
                },
                {
                    "keyword": "keyword C"
                }
            ]
        }
RESPONSE
        {
            "data": null,
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```


```
METHOD  PUT
URL     127.0.0.1:8008/v1/book
HEADER  Authorization   Bearer {token}
BODY
        {
            "id": 8,
            "judul": "judul 1 (updated)",
            "deskripsi": "deskripsi 1 (updated)",
            "harga": 11999.55,
            "stok": 1150,
            "penerbit": "penerbit 1 (updated)"
        }
RESPONSE
        {
            "data": null,
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```

```
METHOD  GET
URL     127.0.0.1:8008/v1/book/all?limit=1&page=1
HEADER  Authorization   Bearer {token}
RESPONSE
        {
            "data": [
                {
                    "id": 2,
                    "judul": "judul 1 (updated)",
                    "deskripsi": "deskripsi 1 (updated)",
                    "harga": 11999.55,
                    "stok": 1150,
                    "penerbit": "penerbit 1 (updated)",
                    "created_at": "2023-04-11 23:24:17",
                    "updated_at": "2023-04-11 16:24:37",
                    "deleted_at": "",
                    "book_category": [
                        {
                            "id": 2,
                            "book_id": 2,
                            "category": "kategori A",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        },
                        {
                            "id": 3,
                            "book_id": 2,
                            "category": "kategori B",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        },
                        {
                            "id": 4,
                            "book_id": 2,
                            "category": "kategori C",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        }
                    ],
                    "book_keyword": [
                        {
                            "id": 2,
                            "book_id": 2,
                            "keyword": "keyword A",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        },
                        {
                            "id": 3,
                            "book_id": 2,
                            "keyword": "keyword B",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        },
                        {
                            "id": 4,
                            "book_id": 2,
                            "keyword": "keyword C",
                            "created_at": "2023-04-11 23:24:17",
                            "updated_at": "",
                            "deleted_at": ""
                        }
                    ]
                }
            ],
            "err_response": null,
            "message": "success",
            "meta": {
                "current_page": 1,
                "last_page": 2,
                "total": 2,
                "per_page": 1
            },
            "status": true
        }
```

```
METHOD  GET
URL     127.0.0.1:8008/v1/book/1
HEADER  Authorization   Bearer {token}
RESPONSE
        {
            "data": {
                "id": 1,
                "judul": "judul 2",
                "deskripsi": "deskripsi 2",
                "harga": 999.55,
                "stok": 50,
                "penerbit": "penerbit 2",
                "created_at": "2023-04-11 23:24:07",
                "updated_at": "",
                "deleted_at": "",
                "book_category": [
                    {
                        "id": 1,
                        "book_id": 1,
                        "category": "kategori A",
                        "created_at": "2023-04-11 23:24:07",
                        "updated_at": "",
                        "deleted_at": ""
                    }
                ],
                "book_keyword": [
                    {
                        "id": 1,
                        "book_id": 1,
                        "keyword": "keyword A",
                        "created_at": "2023-04-11 23:24:07",
                        "updated_at": "",
                        "deleted_at": ""
                    }
                ]
            },
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```

```
METHOD  DELETE
URL     127.0.0.1:8008/v1/book
HEADER  Authorization   Bearer {token}
BODY
        {
            "id": [
                1,
                2
            ]
        }
RESPONSE
        {
            "data": null,
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```