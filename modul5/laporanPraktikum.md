# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[nama] - [NIM]</p>

## Unguided 

### 1. [Soal 2]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(0,n)
}

func CetakDeret (i, n int) {
	if i <= n {
		fmt.Printf("%d ",Fibonacci(i))
		CetakDeret(i + 1, n)
	}
}

func Fibonacci(i int) int {
	if i <= 0 {
		return 0
	}else if i == 1 {
		return 1
	}else {
		return Fibonacci(i-1) + Fibonacci(i-2)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal1.png)
[Program ini menghasilkan barisan bilangan Fibonacci hingga suku ke-$n$ dengan memanfaatkan rekursi ganda, di mana setiap angka diperoleh dari hasil penjumlahan dua angka sebelumnya (Fn-1 + Fn-2). Melalui fungsi CetakDeret yang memanggil fungsi Fibonacci secara berulang, program menampilkan urutan logis pertumbuhan angka yang dimulai dari 0 dan 1 sebagai basis datanya.]

### 2. [Soal 2]
#### soal2.go

```go
package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1,n)
}

func CetakDeret (i, n int) {
	if i <= n {
		CetakBintang(1,i)
		fmt.Println()
		CetakDeret(i + 1, n)
	}
}

func CetakBintang (kolom, jumlah int) {
	if kolom <= jumlah {
		fmt.Print("*")
		CetakBintang(kolom + 1, jumlah)
	}	
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal2.png)
[Program ini mengonversi logika perulangan bersarang (nested loop) menjadi rekursi ganda untuk mencetak pola bintang berbentuk segitiga siku-siku. Fungsi CetakDeret bertanggung jawab menentukan jumlah baris secara vertikal, sementara fungsi CetakBintang bekerja secara horizontal untuk mencetak jumlah karakter bintang yang bertambah seiring bertambahnya indeks baris.]

### 3. [Soal 3]
#### soal3.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakFaktor(1,n)
}

func CetakFaktor (i, n int) {
	if i > n {
		return
	}
	
	if n % i == 0 {
		fmt.Printf("%d ", i)
	}
	CetakFaktor(i + 1, n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal3.png)
[Program ini berfungsi untuk mencari semua pembagi positif dari sebuah bilangan $n$ dengan cara menguji setiap angka $i$ dari 1 hingga $n$ secara berurutan melalui panggilan rekursif. Menggunakan operasi modulus (%), program mendeteksi angka-angka yang tidak memiliki sisa bagi terhadap n, lalu mencetaknya ke layar sebagai faktor resmi dari bilangan tersebut.]

### 4. [Soal 4]
#### soal4.go

```go
package main

import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1, n)
}

func CetakDeret(i, n int){
	if i <= n {
		angkaSekarang := n - i + 1
		fmt.Printf("%d ", angkaSekarang)
		if i < n {
			CetakDeret(i + 1, n)
			fmt.Printf("%d ", angkaSekarang)
		}
	}

}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal4.png)
[Program ini mendemonstrasikan perilaku stack (tumpukan) dalam rekursi dengan mencetak urutan angka yang menurun menuju pusat lalu naik kembali ke nilai asalnya. Dengan menempatkan perintah cetak sebelum dan sesudah panggilan rekursif CetakDeret, program menciptakan efek simetris atau "cermin" yang menampilkan bagaimana memori menyimpan dan mengembalikan nilai variabel saat fungsi selesai dieksekusi.]

### 5. [Soal 5]
#### soal5.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1,n)
}

func CetakDeret (i, n int) {
	if i > n {
		return
	}
	CetakHasil(i)
	CetakDeret(i + 1, n)
}

func CetakHasil (i int) {
	if i % 2 != 0 {
		fmt.Printf("%d ", i)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal5.png)
[Program ini melakukan penelusuran bilangan dari 1 sampai n untuk menyaring dan menampilkan hanya angka yang termasuk kategori ganjil. Melalui fungsi CetakHasil, setiap angka diperiksa apakah memiliki sisa bagi jika dibagi dua; jika hasilnya tidak nol, maka angka tersebut dianggap ganjil dan ditampilkan dalam deret keluaran.]

### 6. [Soal 6]
#### soal6.go

```go
package main
import "fmt"
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasil := HitungPangkat(x, y)
	fmt.Printf("%d ", hasil)
}

func HitungPangkat (x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * HitungPangkat(x, y-1)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul5/output/output-soal6.png)
[Program ini mengimplementasikan konsep matematika eksponen x^y ke dalam bentuk rekursi sederhana dengan basis utama bahwa $x^0$ adalah 1. Proses perhitungan dilakukan dengan mengalikan basis x secara terus-menerus dengan hasil panggilan fungsi itu sendiri sambil mengurangi nilai pangkat y hingga mencapai kondisi berhenti (nol).]
