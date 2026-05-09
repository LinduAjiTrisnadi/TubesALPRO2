package main

import "fmt"

const NMAX = 999

type perangkat struct {
	nama   string
	watt   int
	durasi float64
}

type cacatan [NMAX]perangkat

func main() {
	var n int

	fmt.Print("\n")
	fmt.Println("SELAMAT DATANG DI APLIKASI PowerLog")
	fmt.Print("\n")

	fmt.Println("Pilihan Menu: ")
	fmt.Println("Menambahkan Data (1)")
	fmt.Println("Mengubah Data (2)")
	fmt.Println("Menghapus Data (3)")
	fmt.Println("Mau pilih ke menu berapa??")
	fmt.Scan(&n)
	inputData(n)
	fmt.Print("\n")
}

func inputData(n int) {
	var i, x int
	var data cacatan
	var kondisi string

	fmt.Print("\n")
	if n == 1 {
		fmt.Println("Mau Menambahkan berapa data?")
		fmt.Scan(&x)
		fmt.Print("\n")
		fmt.Println("Inputkan nama,watt, dan durasinya:")
		for i = 0; i < x; i++ {
			fmt.Scan(&data[i].nama, &data[i].watt, &data[i].durasi)
		}

		fmt.Println("+------------------------+-----------------+--------------------+")
		fmt.Println("| Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   |")
		fmt.Println("+------------------------+-----------------+--------------------+")

		for i = 0; i < x; i++ {
			fmt.Printf("| %-22s | %-15d | %-15.2f    | \n", data[i].nama, data[i].watt, data[i].durasi)
		}
		fmt.Println("+------------------------+-----------------+--------------------+")

		fmt.Print("Kembali ke menu (Y/T)")
		fmt.Scan(&kondisi)

		if kondisi == "T" {
			fmt.Println("Menambahkan Data (1)")
			fmt.Println("Mengubah Data (2)")
			fmt.Println("Menghapus Data (3)")
			fmt.Println("Mau pilih ke menu berapa??")
		}
	}

}