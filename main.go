package main

import "fmt"

const NMAX = 999

type perangkat struct {
	id int
	nama string
	watt int
	durasi int
	ruangan string
}

type daftar [NMAX]perangkat

func main() {
	var pilih, a int
	var data daftar

	for {
		fmt.Println()
		fmt.Println("SELAMAT DATANG DI APLIKASI POWERLOG")
		fmt.Println()
		
		fmt.Println("List Menu:")
		fmt.Println("1. Cek data perangkat rumah")
		fmt.Println("2. Keluar")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)

		switch pilih{
		case 1:
			menuCekData(&data, &a)
		case 2:
			fmt.Println("Terima kasih telah menggunakan aplikasi POWERLOG :)")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuCekData(data *daftar, a *int){
	var pilih int
	
	for{
		fmt.Println()
		fmt.Print("MENU CEK DATA PERANGKAT RUMAH")
		fmt.Println()
		fmt.Println("List Menu:")
		fmt.Println("0. Isi terlebih dahulu datanya")
		fmt.Println("1. Tampilkan data")
		fmt.Println("2. Menambah data")
		fmt.Println("3. Menghapus data")
		fmt.Println("4. Mengubah data")
		fmt.Println("5. Mencari data menggunakan (Sequential Search)")
		fmt.Println("6. Mencari data menggunakan (Binary Search)")
		fmt.Println("7. Kembali")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
		case 0:
			awalData(data, a)
		case 1:
			tampilData(data, a)
		case 2:
			if *a == 0 {
				fmt.Println("Data masih kosong! Memasuki isi terlebih dahulu datanya...")
				awalData(data, a)
			} else {
				tambahData(data, a)
			}
		case 3:
			apusData(data, a)
		case 4:
			ubahData(data, a)
		case 5:
			cariSequential(data, a)
		case 6:
			cariBinary(data, a)
		case 7:
			return
		}
	}
}

func tampilData(data *daftar, a *int){
	var i int

	if *a == 0 {
		fmt.Println()
		fmt.Println("Data masih kosong!")
		return
	}
	fmt.Println()
	fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
	fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
	fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")

	for i = 0; i < *a; i++ {
		fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
	}

	fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
}

func awalData(data *daftar, a *int){
	var i int
	i = 0
	
	fmt.Println()
	fmt.Print("Masukan ID: ")
	fmt.Scan(&data[i].id)

	fmt.Print("Masukan nama perangkat: ")
	fmt.Scan(&data[i].nama)

	fmt.Print("Masukan konsumsi Watt: ")
	fmt.Scan(&data[i].watt)

	fmt.Print("Masukan durasi Pemakaian: ")
	fmt.Scan(&data[i].durasi)
	
	fmt.Print("Masukan lokasi ruangan perangkat: ")
	fmt.Scan(&data[i].ruangan)

	i = i + 1
	*a = i
}

func tambahData(data *daftar, a *int){
	var i, n, j int
	fmt.Println()
	fmt.Print("Mau berapa banyak data: ")
	fmt.Scan(&n)

	if *a == NMAX {
		fmt.Println("Data penuh!")
		return
	}

	for i = *a; i < *a + n; i++ {

		fmt.Print("Masukan ID: ")
		fmt.Scan(&data[i].id)
		
		for j = 0; j < i; j++ {
			if data[i].id == data[j].id {
				fmt.Println("Data duplikat, mohon inputkan data kembali!")
				fmt.Print("Masukan ID: ")
				fmt.Scan(&data[i].id)
				j = -1 
			}
		}
		fmt.Print("Masukan nama perangkat: ")
		fmt.Scan(&data[i].nama)

		fmt.Print("Masukan konsumsi Watt: ")
		fmt.Scan(&data[i].watt)

		fmt.Print("Masukan durasi Pemakaian: ")
		fmt.Scan(&data[i].durasi)
		
		fmt.Print("Masukan lokasi ruangan perangkat: ")
		fmt.Scan(&data[i].ruangan)
	}

	*a = i
	fmt.Println("Data perangkat baru sudah ditambahkan!")
}

func apusData(data *daftar, a *int){
	var i, j, target int
	var found bool
	
	fmt.Println()
	fmt.Print("Masukan ID yang dihapus: ")
	fmt.Scan(&target)
	
	for i = 0; i < *a; i++{
		if data[i].id == target{
			found =  true
				
			*a = *a - 1
			
			for j = i; j < *a; j++ {
				data[j] = data[j+1]
				data[i].id = data[i].id - 1
			}
				
			fmt.Printf("ID %d berhasil dihapus!\n", target)
			return
		}
	}
	if found == false{
		fmt.Print("ID tidak ditemukan!")
	}
}

func ubahData(data *daftar, a *int){
	var i, j,  target int
	var found bool
	
	fmt.Println()
	fmt.Print("Masukan ID yang ingin diubah: ")
	fmt.Scan(&target)
	
	for i = 0; i < *a; i++{
		if data[i].id == target{
			found =  true
			for j = i; j < *a; j++ {
				data[j] = data[j+1]
				data[i].id = data[i].id + 1
			}
			
			fmt.Print("Masukan nama perangkat: ")
			fmt.Scan(&data[i].nama)

			fmt.Print("Masukan konsumsi Watt: ")
			fmt.Scan(&data[i].watt)

			fmt.Print("Masukan durasi Pemakaian: ")
			fmt.Scan(&data[i].durasi)
		
			fmt.Print("Masukan lokasi ruangan perangkat: ")
			fmt.Scan(&data[i].ruangan)
				
			fmt.Printf("ID %d berhasil diubah!\n", target)
			return
		}
	}
	if found == false{
		fmt.Print("ID tidak ditemukan!")
	}
}

func cariSequential(data *daftar, a *int){
	var i, pilih int
	var target string
	var found bool
	
	fmt.Println()
	fmt.Println("List data yang ingin dicari: ")
	fmt.Println("1. Nama perangkat")
	fmt.Println("2. Lokasi ruang perangkat ")
	fmt.Println()
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&pilih)
	
	switch pilih{
		case 1:
			fmt.Println()
			fmt.Print("Nama perangkat: ")
			fmt.Scan(&target)
			
			found = false
			
			for i = 0; i < *a; i++ {
				if data[i].nama == target {
					if found == false {
						fmt.Println()
						fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
						fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
						fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
						found = true
					}
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",
					data[i].id, data[i].nama, (data[i].watt*data[i].durasi), data[i].durasi, data[i].ruangan)
				}
			}

			if found == true {
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			} else {
				fmt.Println("Nama perangkat tidak ditemukan!")
			}
		case 2:
			fmt.Println()
			fmt.Print("Lokasi perangkat: ")
			fmt.Scan(&target)
			
			found = false
			
			for i = 0; i < *a; i++ {
				if data[i].ruangan == target {
					if found == false {
						fmt.Println()
						fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
						fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
						fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
						found = true
					}
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",
					data[i].id, data[i].nama, (data[i].watt*data[i].durasi), data[i].durasi, data[i].ruangan)
				}
			}

			if found == true {
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			} else {
				fmt.Println("Lokasi perangkat tidak ditemukan!")
			}
	}
}
func cariBinary(data *daftar, a *int){
	var idx, pilih int
	var target string
	var kiri, kanan, tengah int
	
	fmt.Println()
	fmt.Println("List data yang ingin dicari: ")
	fmt.Println("1. Nama perangkat")
	fmt.Println("2. Lokasi ruang perangkat ")
	fmt.Println()
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&pilih)
	
	switch pilih{
		case 1:
			SortNamaAscend(data, *a)
			fmt.Println()
			fmt.Print("Nama perangkat: ")
			fmt.Scan(&target)
			
			kiri = 0
			kanan = *a - 1
			idx = -1 
			
			for kiri <= kanan {
				tengah = (kiri + kanan) / 2
				if data[tengah].nama == target {
					idx = tengah 
					break
				}else if data[tengah].nama < target {
					kiri = tengah + 1
				}else {
					kanan = tengah - 1
				}
			}	
			if idx != -1 {
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				for idx > -1 && data[idx].nama == target {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n", data[idx].id, data[idx].nama, (data[idx].watt * data[idx].durasi), data[idx].durasi, data[idx].ruangan)
					idx = idx - 1
				}
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			}else {
				fmt.Println("Nama perangkat tidak ditemukan!")
			}
			
		case 2:
			SortRuanganAscend(data, *a)
			fmt.Println()
			fmt.Print("Lokasi perangkat: ")
			fmt.Scan(&target)
			
			kiri = 0
			kanan = *a - 1
			idx = -1 
			
			for kiri <= kanan {
				tengah = (kiri + kanan) / 2
				if data[tengah].ruangan == target {
					idx = tengah 
					break
				}else if data[tengah].ruangan < target {
					kiri = tengah + 1
				}else {
					kanan = tengah - 1
				}
			}	
			if idx != -1 {
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				for idx > -1 && data[idx].ruangan == target {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n", data[idx].id, data[idx].nama, (data[idx].watt * data[idx].durasi), data[idx].durasi, data[idx].ruangan)
					idx = idx - 1
				}
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			}else {
				fmt.Println("Lokasi perangkat tidak ditemukan!")
			}
	}
}

func SortNamaAscend(data *daftar, a int){
	var i, j, idx int
	var temp perangkat
	for i = 0; i < a - 1; i++ {
		idx = i
		for j = i + 1; j < a; j++ {
			if data[j].nama < data[idx].nama {
				idx = j
			}
		}
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func SortRuanganAscend(data *daftar, a int){
	var i, j, idx int
	var temp perangkat
	for i = 0; i < a - 1; i++ {
		idx = i
		for j = i + 1; j < a; j++ {
			if data[j].ruangan < data[idx].ruangan {
				idx = j
			}
		}
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}
