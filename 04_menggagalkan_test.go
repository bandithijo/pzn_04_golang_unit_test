package main

/*
Menggagalkan Unit Test
- Menggagalkan unit test menggunakna panic bukanlah hal yang bagus
- Go sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
- Terdapat function Fail(), FailNow(), Error(), dan Fatal() jika kita ingin menggagalkan unit test

t.Fail() dan t.FailNow()
- Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Bedanya?
- Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
- FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

t.Error(args...) dan t.Fatal(args...)
- Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
- Error() function lebih seperti melakukan log(print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
- Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
- Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
*/
