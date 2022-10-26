# API Spesification Final Project Nest Academy
Ini adalah API Spec untuk final projcet di bootcamp NEST Academy
by: [Reyhan Jovie](https:#www.linkedin.com/in/reyhanjovie/)

## Brief
Kamu diminta untuk membuat sebuah aplikasi oleh seorang client :
> Halo, saya NooBee ingin meminta jasa kamu untuk membuat aplikasi untuk toko saya. Saya menjual berbagai pakaian, seperti Baju, Celana, Jilbab, Jacket, dan lainnya. Sebelumnya saya berjualan di toko hijau. Namun karena toko saya semakin besar, jadinya biaya admin dan operasional di toko hijau pun semakin besar.

Lalu, client tersebut ingin membuat aplikasi online shop-nya sendiri, yang mana mereka mempunyai beberapa jenis karyawan seperti berikut :
> Saya ada beberapa karyawan, bisa dibilang ada 2 kategori. Ada admin yang bertugas untuk memantau karyawan, dan kasir yang bertugas untuk melakukan transaksi kepada konsumen

Adapun, client ingin karyawan dan konsumennya nanti bisa melakukan hal hal berikut :
- Owner / admin
  - Manage data karyawan (Create, Read, Update, Delete)
  - Melihat seluruh data member (Read, in general data)
  - Melihat total transaksi per bulan (Read)
  - Manage seluruh produk (Create, Read, Update, Delete)
  - Login
- Kasir
  - Approve transaksi dari customer
  - Update status transaksi (Approve -> Sending)
  - Login
- Member / Customer
  - Registrasi
  - Login
  - Melihat produk
  - Melihat detail produk
  - Membeli produk

## Architecture
![Architecture](https://res.cloudinary.com/noobeeid/image/upload/v1666584475/bootcamp/others/Screen_Shot_2022-10-24_at_11.07.40_dwcp1k.png)
- System dibuat menggunakan arsitektur monolith
- Database yang digunakan adalah `MySQL` atau `PostgreSQL`
- Terdapat 4 services, yaitu :
  - Auth
  - User
  - Product
  - Transaction
- Terintegrasi dengan Third Party, yaitu **Raja Ongkir**
- Package yg dibutuhkan :

| Nama Package | Fungsi | 
| --- | --- |
| `golang.org/x/crypto` | Hash dan Verify Password |
| `github.com/golang-jwt/jwt` | Generate dan Verify Token |
| `gorm.io/gorm` | ORM di golang `optional`|
| `gorm.io/drivers/<postgres, mysql>` | Driver untuk ORM `optional`|
| `github.com/lib/pq` | Driver untuk Native Postgres driver `optional` |


## API Spec
### Auth
#### POST /auth/register
Berfungsi untuk melakukan registrasi. Fitur ini akan otomatis membuat user yang registrasi menjadi `customer`.
**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string"
}
```
**Notes** : Pada proses ini, `password` akan di hash menggunakan library `crypto`

**Response Body** :
```json
{
    "status" : 201, # created
    "message" : "REGISTER_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "REGISTER_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```
> Untuk mengubah menjadi owner / admin, silahkan langsung di update ke databasenya.

#### POST /auth/login
Berfungsi untuk login seluruh user

**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string"
}
```
**Notes** : Pada proses ini, `password` akan di verify menggunakan library `crypto`

**Response Body** :
```json
{
    "status" : 200, 
    "message" : "LOGIN_SUCCESS",
    "payload" : {
        "token" : "string" # JWT 
    },
    "general_info" : "NooBee-Shop"
}
```
**Notes** : Pada proses ini, akan me-generate token menggunakan `JWT`

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "LOGIN_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


### User
#### POST /users
Berfungsi untuk create user

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "fullname" : "string",
    "gender" : "string",
    "contact" : "string",
    "street" : "string",
    "city_id" : "string", 
    "province_id" : "string"
}
```
**Notes** : `province_id` dan `city_id` di dapat dari API Raja ongkir yang sudah di wrap ke aplikasi kamu.

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATED_USER_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "CREATED_USER_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```

#### GET /users
Akan menampilkan seluruh data users. Endpoint ini hanya berlaku untuk `admin` atau `owner`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Query String**
- limit : `int` with default is 25  | optional
- page  : `int` with default is 1   | optional

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_ALL_USERS_SUCCESS",
    "payload" : [
        {
            "id" : "string",
            "fullname" : "string",
            "address" : {
                "city" : {
                    "id" : "string",
                    "name" : "string",
                },
                "province" : {
                    "id" : "string",
                    "name" : "string"
                },
                "street" : "string"
            },
            "auth" : {
                "email" : "string"
            }
        }
    ],
    "query" : {
        "limit" : 25, # default is 25
        "page" : 1, # default is 1
        "total" : 3
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_ALL_USERS_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


#### GET /users/profile
Akan menampilkan detail profile dari user yang sedang login.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_USER_PROFILE_SUCCESS",
    "payload" : {
        "id" : "string",
        "fullname" : "string",
        "gender" : "string",
        "contact" : "string",
        "address" : {
            "city" : {
                "id" : "string",
                "name" : "string",
            },
            "province" : {
                "id" : "string",
                "name" : "string"
            },
            "street" : "string"
        },
        "auth" : {
            "email" : "string"
        }
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 401, # Unauthorized, Not Found, or others...
    "message" : "GET_USER_PROFILE_FAIL",
    "error" : "UNAUTHORIZED",
    "additional_info" : {
        "message" : "you need to login for access this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


#### GET /users/email/:email
Akan menampilkan detail profile dari user berdasarkan email. Endpoint ini hanya berlaku untuk `admin` atau `owner`.

**Params**
- email : `string` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_USER_BY_EMAIL_SUCCESS",
    "payload" : {
        "id" : "string",
        "fullname" : "string",
        "gender" : "string",
        "contact" : "string",
        "address" : {
            "city" : {
                "id" : "string",
                "name" : "string",
            },
            "province" : {
                "id" : "string",
                "name" : "string"
            },
            "street" : "string"
        },
        "auth" : {
            "email" : "string"
        }
    }
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Forbidden, Not Found, or others...
    "message" : "GET_USER_BY_EMAIL_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```

#### PUT /users/profile 
Berfungsi untuk mengubah data user yang sedang login

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "fullname" : "string",
    "gender" : "string",
    "contact" : "string",
    "street" : "string",
    "city_id" : "string", 
    "province_id" : "string"
}
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_USER_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, # Bad request, or others...
    "message" : "UPDATE_USER_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


### Product
#### POST /products
Berfungsi untuk menambah product baru. Endpoint ini hanya bisa di akses oleh `admin`

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name" : "string",
    "category" : "string",
    "description" : "string",
    "price" : 0, 
    "stock" : 0,
    "weight" : 0, #gram
    "img_url" : "string"
}
```
**Notes** : untuk `img_url`, cukup lampirkan saja url dari image yang ada di google. 

**Response Body**
```json
{
    "status" : 201,
    "message" : "CREATE_PRODUCT_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "CREATE_PRODUCT_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


#### GET /products
Berfungsi untuk melihat daftar product. 

**Query String**
- limit : `int` with default is 25  | optional
- page  : `int` with default is 1   | optional

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_ALL_PRODUCTS_SUCCESS",
    "payload" : [
        {
            "id" : "string",
            "name" : "string",
            "category" : "string",
            "description" : "string",
            "price" : 0, 
            "stock" : 0,
            "img_url" : "string"
        }
    ],
    "query" : {
        "limit" : 25, # default is 25
        "page" : 1, # default is 1
        "total" : 3
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 404, # Bad request, or others...
    "message" : "GET_ALL_PRODUCTS_FAIL",
    "error" : "NOT_FOUND",
    "additional_info" : {
        "message" : "data not found in this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```

#### GET /products/id/:id
Berfungsi untuk melihat detail product berdasarkan ID product. Endpoint ini hanya bisa di akses jika telah login

**Params**
- id : `string` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_DETAIL_PRODUCT_SUCCESS",
    "payload" : {
        "id" : "string",
        "name" : "string",
        "category" : "string",
        "description" : "string",
        "price" : 0, 
        "stock" : 0,
        "img_url" : "string",
        "created_at" : "timestamp",
        "updated_at" : "timestamp"
    }
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 404, # Bad request, or others...
    "message" : "GET_DETAIL_PRODUCT_FAIL",
    "error" : "NOT_FOUND",
    "additional_info" : {
        "message" : "data not found in this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


#### PUT /products/id/:id
Berfungsi untuk mengubah data dari product. Endpoint ini hanya bisa di akses oleh `admin`

**Params**
- id : `string` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "name" : "string",
    "category" : "string",
    "description" : "string",
    "price" : 0, 
    "stock" : 0,
    "img_url" : "string"
}
```
**Notes** : untuk `img_url`, cukup lampirkan saja url dari image yang ada di google. 

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_PRODUCT_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "UPDATE_PRODUCT_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```

#### DELETE /products/id/:id
Berfungsi untuk menghapus product. Endpoint ini hanya bisa di akses oleh `admin`

**Params**
- id : `string` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "DELETE_PRODUCT_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
**Notes** : Saat delete product, hanya mengubah status product dari `1` menjadi `0` atau biasa di sebut dengan `soft delete`.

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "DELETE_PRODUCT_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```



### Transaction
#### POST /transactions/inquire
Ini berfungsi untuk inquiry transaksi. Pada proses ini, akan dilakukan pengecekan untuk estimasi ongkos kirim. Proses ini hanya bisa dilakukan oleh `customer`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "product_id" : "string",
    "product_name" : "string",
    "quantity" : "int",
    "destination" : "int", #id dari kota customer
    "weight" : "int", #calculation dari quantity * product_weight
    "total_price" : "int", #calculation dari price * quantity
    "courier" : "string", #silahkan cek di Raja Ongkir
}
```
**Notes** : Pada proses ini, akan melakukan pengecekan ke stok product, apakah stok masih tersedia atau tidak. Dan juga akan melakukan pengecekan untuk estimasi ongkos kirim ke ***API Raja Ongkir***

**Response Body**
```json
{
    "status" : 200,
    "message" : "INQUIRY_TRANSACTION_SUCCESS",
    "payload" : {
        "product" : {
            "id" : "string",
            "name" : "string", 
            "img_url" : "string",
            "price" : "int"
        },
        "quantity" : "int",
        "destination" : "int", 
        "weight" : "int", 
        "total_price" : "int", 
        "servcies_courier" : [ #this result is from Raja Ongkir
            {
                "code" : "string", #code courier
                "name" : "string",
                "costs" : [
                    {
                        "services" : "string",
                        "description" : "string",
                        "cost" : [
                            {
                                "value" : "int",
                                "estimation" : "string", #in days
                                "note" : "string"  
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "general_info" : "NooBee-Shop"
}
```

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 422, # Unprocessable entity, or others...
    "message" : "INQUIRY_TRANSACTION_FAIL",
    "error" : "UNPROCESSABLE_ENTITY",
    "additional_info" : {
        "message" : "stock prodcut not enough", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```

#### POST /transactions/confirm
Ini berfungsi untuk confirmation transaction. Pada proses ini, akan dilakukan proses update stok pada produk dan insert ke table transactions. Proses ini hanya bisa dilakukan oleh `customer`.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "product_id" : "string",
    "product_name" : "string",
    "quantity" : "int",
    "destination" : "int", #id dari kota customer
    "weight" : "int", #calculation dari quantity * product_weight
    "total_price" : "int", #calculation dari price * quantity
    "courier" : {
        "code" : "string", #silahkan cek di Raja Ongkir
        "service" : "string", #OKE, REG, atau yang lainnya
        "cost" : "int",
        "estimation" : "string"
    }
}
```
**Notes** : Pada proses ini, akan melakukan pengecekan ke stok product, apakah stok masih tersedia atau tidak. Dan juga akan melakukan pengecekan untuk estimasi ongkos kirim ke ***API Raja Ongkir***. Secara default, status yang di set adalah `WAITING`

**Response Body**
```json
{
    "status" : 200,
    "message" : "CONFIRM_TRANSACTION_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```


Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 403, # Bad request, or others...
    "message" : "CONFIRM_TRANSACTION_FAIL",
    "error" : "FORBIDDEN_ACCESS",
    "additional_info" : {
        "message" : "you dont have access for this resources", # or others 
    },
    "general_info" : "NooBee-Shop"
}
```


#### GET /transactions/histories/me
Ini berfungsi untuk melihat riwayat transaction si customer yang sedang login.

**Query String**
- limit : `int` with default is 25  | optional
- page  : `int` with default is 1   | optional

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_TRANSACTION_HISTORIES_SUCCESS",
    "payload" : [
        {
            "id" : "string",
            "product_id" : "string",
            "product_name" : "string",
            "quantity": "string",
            "destination": {
                "city" : "string",
                "province" : "string",
            },
            "weight" : "int",
            "total_price" : "int",
            "courier" : {
                "code" : "string", 
                "service" : "string",
                "cost" : "int",
                "estimation" : "string"
            },
            "status" : "string", #waiting, pickup, on the way, arrived. 
            "estimation_arrived" : "string",
            "created_at" : "timestamp",
            "updated_at" : "timestamp",
        }
    ],
     "query" : {
        "limit" : 25, # default is 25
        "page" : 1, # default is 1
        "total" : 3
    }
    "general_info" : "NooBee-Shop"
}
```

#### GET /transactions/histories/list
Ini berfungsi untuk melihat riwayat transaction seluruh user. Endpoint ini hanya bisa di akses oleh `admin` dan `kasir`.

**Query String**
- limit : `int` with default is 25  | optional
- page  : `int` with default is 1   | optional

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body**
```json
{
    "status" : 200,
    "message" : "GET_TRANSACTION_HISTORIES_SUCCESS",
    "payload" : [
        {
            "id" : "string",
            "product_id" : "string",
            "product_name" : "string",
            "quantity": "string",
            "destination": {
                "city" : "string",
                "province" : "string",
            },
            "weight" : "int",
            "total_price" : "int",
            "status" : "string", #waiting, pickup, on the way, arrived
            "created_at" : "timestamp",
            "updated_at" : "timestamp",
        }
    ],
     "query" : {
        "limit" : 25, # default is 25
        "page" : 1, # default is 1
        "total" : 3
    }
    "general_info" : "NooBee-Shop"
}
```

#### PUT /transactions/id/:id/status
Ini berfungsi untuk mengubah status transaksi customer. endpoint ini hanya bisa di akses oleh `kasir`

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body**
```json
{
    "status" : "string"
}
```
List status : 
- WAITING 
- PICKUP
- ON THE WAY
- ARRIVED

**Response Body**
```json
{
    "status" : 200,
    "message" : "UPDATE_STATUS_TRANSACTION_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
