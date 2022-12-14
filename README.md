# tugas bootcamp ke 3

![Alt text](screen11.jpg?raw=true "gambar")

## penjelasan
1. disini saya mengunakan nsq sebagai message broker / queue
2. di runner 1 saya mempublish pesan menggunakan goroutine agar saat mengirim pesan berjalan secara asynchronous atau tidak saling tunggu
3. kemudian runner ke 2 sebagai subscriber dari runner 1(publish)
4. saya handle error menggunakan library logrus, kebetulan saya hanya menampilkan log nya ke terminal; tidak menyimpannya ke dalam file