# crowdfunding
Aplikasi Crowdfunding adalah aplikasi web penggalangan dana yang melibatkan kontribusi dari banyak individu untuk mendanai proyek, usaha, atau tujuan tertentu.

    Aplikasi crowdfunding ini dibangun menggunakan stack 
    1.  Golang pada sisi backend
        1.  Gin Framework Restful API
        2.  ORM GORM 
    2.  Vue NuxtJS untuk sisi frontend 
    3.  PostgreSQL untuk sisi Database


### Restore DB pada PostgreSQL v16

1. create database crowdfunding
2. jalankan commandline restore :

        $ pg_restore -h localhost -p 5432 -U postgres -d crowdfunding "db_crowfunding_20250416.sql"