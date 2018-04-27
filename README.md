# kingsman
bukalapak o2o agent microservice

# Instalasi
1. Download project kingsman: `go get github.com/bukalapak/kingsman`
2. Pergi ke folder kingsman: `cd ~/go/src/github.com/bukalapak/kingsman`
3. Jalankan `make dep`
4. Install `bundle` untuk migrasi basis data: `Gems install bundle`
5. Install dependency Rails untuk migrasi basis data: `bundle install`
6. Gandakan dan sesuaikan berkas environment-nya: `cp env.sample .env` lalu ubah berkas `.env` sesuai dengan konfigurasi lokal
6. Buatlah basis datanya: `bundle exec rake db:create`
7. Migrasikan basis datanya: `bundle exec rake db:migrate`
8. Kingsman siap untuk dijalankan: `make run`

# CP 
Jika membutuhkan bantuan terkait microservice ini, silahkan menghubungi:
- Fadil Muhammad Putra: @fadilmuhput
- Juan Anton: @juan_anton
- Yulistiyan Wardhana: @ywardhana25
