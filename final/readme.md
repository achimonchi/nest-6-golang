# API Spesification Final Project Nest Academy
Ini adalah API Spec untuk final projcet di bootcamp NEST Academy
by: [Reyhan Jovie](https://www.linkedin.com/in/reyhanjovie/)

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
| `gorm.io/drivers/<postgres | mysql>` | Driver untuk ORM `optional`|
| `github.com/lib/pq` | Driver untuk Native Postgres driver `optional` |


## API Spec
### Auth
#### POST /auth/register
Berfungsi untuk melakukan registrasi. Fitur ini akan otomatis membuat user yang registrasi menjadi `customer`.
**Request Body** :
```json
{
    "email" : "string", // must be a valid email
    "password" : "string"
}
```
**Notes** : Pada proses ini, `password` akan di hash menggunakan library `crypto`

**Response Body** :
```json
{
    "status" : 201, // created
    "message" : "REGISTER_SUCCESS",
    "general_info" : "NooBee-Shop"
}
```
Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, // Bad request, or others...
    "message" : "REGISTER_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", // or others 
    }
    "general_info" : "NooBee-Shop"
}
```
> Untuk mengubah menjadi owner / admin, silahkan langsung di update ke databasenya.

#### POST /auth/login
Berfungsi untuk login seluruh user

**Request Body** :
```json
{
    "email" : "string", // must be a valid email
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
        "token" : "string" // JWT 
    },
    "general_info" : "NooBee-Shop"
}
```
**Notes** : Pada proses ini, akan me-generate token menggunakan `JWT`

Jika gagal, maka akan menghasilkan response :
```json
{
    "status" : 400, // Bad request, or others...
    "message" : "LOGIN_FAIL",
    "error" : "BAD_REQUEST",
    "additional_info" : {
        "message" : "invalid request payload", // or others 
    },
    "general_info" : "NooBee-Shop"
}
```






