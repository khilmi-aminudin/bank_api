# Bank API

Merupakan aplikasi RestAPI bank sederhana dengan fitur sebagai berikkut :
- Register User/Customers Bank
- Register Bank Account
- Transaksi dasar Bank seperti (Topup, Withdrawal, Transfer dan Pembyaran ke merchant)

### Alat-alat
- Git
- Golang versi 19 atau lebih
- Docker

## Cara Menggunakan

### Dengan docker
- Clone repository
  ```shell
  $ git clone https://github.com/khilmi-aminudin/bank_api.git
  $ cd bank_api
  ```
- Ubah app.env.example menjadi app.env dan sesuaikan konfigurasi di dalamnya dengan env milik anda
- Pastikan di komputer sudah terinstall Docker
  ``` shell
  $ docker version
    # Client: Docker Engine - Community
    # Version:           24.0.2
  ```
- Menjalankan aplikasi  
  *UNIX/MacOS*
  ``` shell
    # build aplikasi
    $ make build

    # membuat database dan menjalankan migrasi database
    $ make postgresql && make createdb && make migrateup

    # menjalankan aplikasi
    $ make runapp

  ```
  *Windows*
  ``` shell
    # build aplikasi
    $ docker build -t bank-api .
    
    # membuat postgresql container
    $ docker run --name postgresql -p 5432:5432 -e TZ=Asia/Jakarta -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
    
    # membuat database di dalam postgresql container 
    $ docker exec -it postgresql createdb --username=root --owner=root bank_db

    # menjalankan migrasi database
    $ migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up

    # menjalankan aplikasi
    $ docker run --rm --name bankapi --network="host" -p 8080:8080 bank-api

  ```


### Tanpa Docker
- - Clone repository
  ```shell
  $ git clone https://github.com/khilmi-aminudin/bank_api.git
  $ cd bank_api
  ```
- Pastikan anda memiliki postgresql instance
- Buat database di postgresql 
- Ubah app.env.example menjadi app.env dan sesuaikan konfigurasi di dalamnya dengan env milik anda, serta sesuaikan juga konfigurasi DB_SOURCE didalam Makefile 
- Menjalankan aplikasi  
  *UNIX/MacOS*
  ``` shell
    # melakukan test aplikasi
    $ make test

    # menjalankan migrasi database
    $ && make migrateup

    # menjalankan aplikasi
    $ make runserver

  ```
  *Windows*
  ``` shell
    # test aplikasi
    $ go test -v -cover ./...
    
    # menjalankan migrasi database
    # sesuaikan konfigrasi database
    $ migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up

    # menjalankan aplikasi
    $ go run cmd/main.go

  ```
- Sebagai catatan, aplikasi ini sudah menerapkan CI untuk testing aplikasi saat di push ke github pada branch main


## Dokumentasi API
- Untuk dokumentasi API telah disediakan postman collection yang dapat di import di Postman atau API Client lainnya.
- Data di database masih kosong, hanya tersedia  1 user admin dengan username dan password sebagai berikut:
  ```bash
    "username" : "admin",
    "password" : "admin123Dev"
  ```
  </br>
  </br>
  </br>
  



NB : Aplikasi ini masih banyak yang perlu diperbaiki, dimana saya selaku developer menyadari bahwa banyak code yang penempatannya kurang sesuai, dimana banyak Bussines Logic yang masih saya lakukan di Handlers yang seharusnya dilakukan di Service