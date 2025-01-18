package main

/*
Pengenalan testing Package
- Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test
- Berbeda dengan Go, di Go untuk unit test sudah disediakan sebuah package khusus bernama testing
- Selain itu untuk menjalankan unit test, di Go juga sudah disediakan perintahnya
- Hal in imembuat implementasi unit testing di Go sangat mudah dibanding dengan bahasa pemrograman lain
- <https://pkg.go.dev/testing>

testing.T
- Go menyediakan sebuah struct yang bernama testing.T
- Struct ini digunakan untuk unit test di Go

testing.M
- testing.M adalah struct yang disediakan Go untuk mengatur life cycle testing
- Materi ini akan dibahas di chapter Main atau Before and After unit test

testing.B
- testing.B adalah struct yang disediakan Go untuk melakukan benchmarking
- Di Go untuk melakukan benchmark(mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan
*/
