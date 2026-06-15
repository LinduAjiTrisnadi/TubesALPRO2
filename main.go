package main

import "fmt"

const NMAX = 999

type perangkat struct {
	id, watt, durasi int
	nama, ruangan string
}

type daftar [NMAX]perangkat

func main() {
	var pilih, a int
	var data daftar
	var keluar bool
	
	keluar = false
	for keluar == false{
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
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuCekData(data *daftar, a *int){
	var pilih int
	var keluar bool
	
	keluar = false
	for keluar == false{
		fmt.Println("\nMENU CEK DATA PERANGKAT RUMAH")
		fmt.Println()
		fmt.Println("List Menu:")
		if *a <= 0{
			fmt.Println("0. Isi terlebih dahulu datanya")
		}
		fmt.Println("1. Tampilkan data")
		fmt.Println("2. Menambah/Menghapus/Mengubah data")
		fmt.Println("3. Mencari data menggunakan (Sequential/Binary Search)")
		fmt.Println("4. Mencari data tertinggi menggunakan (Selection/Insertion Sort)")
		fmt.Println("5. Tampilkan Statistik data perangkat")
		fmt.Println("6. Kembali")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
		case 0:
			awalData(data, a)
		case 1:
			tampilData(data, a)
		case 2:
			menuCrud(data, a)
		case 3:
			menuSequBinary(data, a)
		case 4:
			menuSelectInsert(data, a)
		case 5:
			menuStatis(data, a)
		case 6:
			keluar = true
		}
	}
}

func menuCrud(data *daftar, a *int){
	var pilih int
	var keluar bool
	
	keluar = false
	for keluar == false{
		fmt.Println("\nMENU CRUD")
		fmt.Println()
		fmt.Println("List Menu:")
		fmt.Println("1. Menambah data")
		fmt.Println("2. Menghapus data")
		fmt.Println("3. Mengubah data")
		fmt.Println("4. Kembali")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
		case 1:
			if *a == 0 {
				fmt.Println("Data masih kosong! Memasuki isi terlebih dahulu datanya...")
				awalData(data, a)
			} else {
				tambahData(data, a)
			}
		case 2:
			apusData(data, a)
		case 3:
			ubahData(data, a)
		case 4:
			keluar = true
		}
	}
}

func menuSequBinary(data *daftar, a *int){
	var pilih int
	var keluar bool
	
	keluar = false
	for keluar == false{
		fmt.Println("\nMENU Sequ/Binary")
		fmt.Println()
		fmt.Println("List Menu:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("3. Kembali")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
		case 1:
			cariSequential(data, a)
		case 2:
			cariBinary(data, a)
		case 3:
			keluar = true
		}
	}
}

func menuSelectInsert(data *daftar, a *int){
	var pilih int
	var keluar bool
	
	keluar = false
	for keluar == false{
		fmt.Println("\nMENU Select/Insert")
		fmt.Println()
		fmt.Println("List Menu:")
		fmt.Println("1. Selection")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Kembali")
		
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
		case 1:
			menuSelection(data, a)
		case 2:
			menuInsertion(data, a)
		case 3:
			keluar = true
		}
	}
}

func tampilData(data *daftar, a *int){
    var i int
    
    if *a == 0 {
        fmt.Println()
        fmt.Println("Data masih kosong!")
    } else {
        fmt.Println()
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
        fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
        for i = 0; i < *a; i++ {
            fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
        }
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
    }
}

func awalData(data *daftar, a *int){
    
    data[0].id = 1
    data[0].nama = "Kipas"
    data[0].watt = 4
    data[0].durasi = 12
    data[0].ruangan = "Ruang_tamu"
    
    data[1].id = 2
    data[1].nama = "AC"
    data[1].watt = 100
    data[1].durasi = 12
    data[1].ruangan = "Kamar_tidur"
    
    data[2].id = 3
    data[2].nama = "TV"
    data[2].watt = 20
    data[2].durasi = 5
    data[2].ruangan = "Ruang_tamu"
    
    *a = 3
    
    fmt.Println()
    fmt.Println("Data sudah terisi")
}

func tambahData(data *daftar, a *int){
    var i, n int
    fmt.Println()
    fmt.Print("Mau berapa banyak data: ")
    fmt.Scan(&n)
    fmt.Println()

    if *a == NMAX {
        fmt.Println("Data penuh!")
    } else {
        for i = *a; i < *a + n; i++ {
            data[i].id = i + 1  

            fmt.Print("Masukan nama perangkat: ")
            fmt.Scan(&data[i].nama)
            fmt.Print("Masukan konsumsi Watt: ")
            fmt.Scan(&data[i].watt)
            fmt.Print("Masukan durasi Pemakaian (perjam): ")
            fmt.Scan(&data[i].durasi)
            fmt.Print("Masukan lokasi ruangan perangkat: ")
            fmt.Scan(&data[i].ruangan)
            fmt.Println()
        }
        *a = i
        fmt.Println("Data perangkat baru sudah ditambahkan!")
    }
}

func apusData(data *daftar, a *int) {
    var i, j, target int
    var found bool

    fmt.Println()
    fmt.Print("Masukan ID yang dihapus: ")
    fmt.Scan(&target)

    i = 0
	found = false
	
    for i < *a && found == false {
        if data[i].id == target {
            found = true
            *a = *a - 1

            for j = i; j < *a; j++ {
                data[j] = data[j+1]
            }

            for j = 0; j < *a; j++ {
                if data[j].id > target {
                    data[j].id = data[j].id - 1
                }
            }

            fmt.Printf("ID %d berhasil dihapus!\n", target)
        }
        i = i + 1
    }

    if found == false {
        fmt.Println("ID tidak ditemukan!")
    }
}

func ubahData(data *daftar, a *int){
    var i, target int
    var found bool

    fmt.Println()
    fmt.Print("Masukan ID yang ingin diubah: ")
    fmt.Scan(&target)

    i = 0
	found = false
    for i < *a && found == false {
        if data[i].id == target {
            found = true

            fmt.Print("Masukan nama perangkat: ")
            fmt.Scan(&data[i].nama)
            fmt.Print("Masukan konsumsi Watt: ")
            fmt.Scan(&data[i].watt)
            fmt.Print("Masukan durasi Pemakaian: ")
            fmt.Scan(&data[i].durasi)

            fmt.Print("Masukan lokasi ruangan perangkat: ")
            fmt.Scan(&data[i].ruangan)

            fmt.Printf("ID %d berhasil diubah!\n", target)
        }
        i = i + 1
    }

    if found == false {
        fmt.Println("ID tidak ditemukan!")
    }
}

func cariSequential(data *daftar, a *int){
	var i, pilih int
	var target string
	var found, keluar bool
	
	keluar = false
	
	for keluar == false{
		fmt.Println()
		fmt.Println("List data yang ingin dicari: ")
		fmt.Println("1. Nama perangkat")
		fmt.Println("2. Lokasi ruang perangkat ")
		fmt.Println("3. Kembali ")
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
			case 3:
				keluar = true
		}
	}
}

func cariBinary(data *daftar, a *int){
	var idx, pilih, i int
	var target string
	var kiri, kanan, tengah, batasKanan, batasKiri int
	var temp daftar
	var keluar bool
	
	keluar = false
	
	for i = 0; i < *a; i++ {
		temp[i] = data[i]
	}
	
	for keluar == false{
		fmt.Println()
		fmt.Println("List data yang ingin dicari: ")
		fmt.Println("1. Nama perangkat")
		fmt.Println("2. Lokasi ruang perangkat ")
		fmt.Println("3. Kembali ")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
			case 1:
				SortNamaAscend(&temp, *a)
				fmt.Println()
				fmt.Print("Nama perangkat: ")
				fmt.Scan(&target)
				
				kiri = 0
				kanan = *a - 1
				idx = -1
				for kiri <= kanan {
					tengah = (kiri + kanan) / 2
					if temp[tengah].nama == target {
						idx = tengah
						kanan = tengah - 1 
					} else if temp[tengah].nama < target {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}
				batasKiri = idx

				kiri = 0
				kanan = *a - 1
				idx = -1
				for kiri <= kanan {
					tengah = (kiri + kanan) / 2
					if temp[tengah].nama == target {
						idx = tengah
						kiri = tengah + 1 
					} else if temp[tengah].nama < target {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}
				batasKanan = idx

				if batasKiri != -1 {
					fmt.Println()
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
					fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
					for i = batasKiri; i <= batasKanan; i++ {
						fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",
						temp[i].id, temp[i].nama, (temp[i].watt*temp[i].durasi), temp[i].durasi, temp[i].ruangan)
					}
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				} else {
					fmt.Println("Nama perangkat tidak ditemukan!")
				}
				
			case 2:
				SortRuanganAscend(&temp, *a)
				fmt.Println()
				fmt.Print("Lokasi perangkat: ")
				fmt.Scan(&target)
				
				kiri = 0
				kanan = *a - 1
				idx = -1
				
				for kiri <= kanan {
					tengah = (kiri + kanan) / 2
					if temp[tengah].ruangan == target {
						idx = tengah
						kanan = tengah - 1 
					} else if temp[tengah].ruangan < target {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}
				batasKiri = idx

				kiri = 0
				kanan = *a - 1
				idx = -1
				for kiri <= kanan {
					tengah = (kiri + kanan) / 2
					if temp[tengah].ruangan == target {
						idx = tengah
						kiri = tengah + 1 
					} else if temp[tengah].ruangan < target {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}
				batasKanan = idx

				if batasKiri != -1 {
					fmt.Println()
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
					fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
					for i = batasKiri; i <= batasKanan; i++ {
						fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",
							temp[i].id, temp[i].nama, (temp[i].watt*temp[i].durasi), temp[i].durasi, temp[i].ruangan)
					}
					fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				} else {
					fmt.Println("Lokasi perangkat tidak ditemukan!")
				}
			case 3:
				keluar = true
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

func SortKonsumsiDescend(data *daftar, a int){
	var i, j, idx int
	var temp perangkat
	
	for i = 0; i < a - 1; i++ {
		idx = i
		for j = i + 1; j < a; j++ {
			if data[j].watt > data[idx].watt {
				idx = j
			}
		}
		
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func SortNamaDescend(data *daftar, a int){
	var i, j, idx int
	var temp perangkat
	for i = 0; i < a - 1; i++ {
		idx = i
		for j = i + 1; j < a; j++ {
			if data[j].nama > data[idx].nama {
				idx = j
			}
		}
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func insertSortKonsumsi(data *daftar, a int) {
	var i, j int
	var temp perangkat
	for i = 1; i < a; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].watt < temp.watt {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = temp
	}
}

func insertSortNama(data *daftar, a int) {
	var i, j int
	var temp perangkat
	for i = 1; i < a; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].nama < temp.nama {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = temp
	}
}

func menuSelection(data *daftar, a *int){
	var pilih, i int
	var keluar bool
	
	keluar = false
	
	for keluar == false{
			
		fmt.Println()
		fmt.Println("List data untuk diurutkan: ")
		fmt.Println("1. Konsumsi Watt")
		fmt.Println("2. Nama Perangkat ")
		fmt.Println("3. Kembali ")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
			case 1:
				SortKonsumsiDescend(data, *a)
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")

				for i = 0; i < *a; i++ {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
				}

				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			case 2:
				SortNamaDescend(data, *a)
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")

				for i = 0; i < *a; i++ {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
				}

				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			case 3:
				keluar = true
		}
	}
}

func menuInsertion(data *daftar, a *int){
	var pilih, i int
	var keluar bool
	
	keluar = false
	
	for keluar == false{
		
		fmt.Println()
		fmt.Println("List data untuk diurutkan: ")
		fmt.Println("1. Konsumsi Watt")
		fmt.Println("2. Nama Perangkat ")
		fmt.Println("3. Kembali ")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		
		switch pilih{
			case 1:
				insertSortKonsumsi(data, *a)
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")

				for i = 0; i < *a; i++ {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
				}

				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			case 2:
				insertSortNama(data, *a)
				fmt.Println()
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
				fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")

				for i = 0; i < *a; i++ {
					fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[i].id,data[i].nama,(data[i].watt*data[i].durasi),data[i].durasi,data[i].ruangan)
				}

				fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
			case 3:
				keluar = true
		}
	}
	
}

func menuStatis(data *daftar, a *int){
    var i, total, max int

    if *a == 0 {
        fmt.Println("Data masih kosong!")
    } else {
        total = 0
        max = 0

        for i = 0; i < *a; i++ {
            total = total + (data[i].watt * data[i].durasi)
        }
        for i = 0; i < *a; i++ {
            if data[i].watt * data[i].durasi > data[max].watt * data[max].durasi {
                max = i
            }
        }
        fmt.Println()
        fmt.Printf("Total penggunaan Daya harian: %d W\n", total)
        fmt.Println()
        fmt.Println("Perangkat yang paling boros Energi:")
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
        fmt.Println("| ID    | Nama Perangkat         | Konsumsi Watt   | Durasi Pemakaian   | Lokasi Ruangan Perangkat |")
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
        fmt.Printf("| %-5d | %-22s | %-14dW | %-15dJam | %-24s |\n",data[max].id,data[max].nama,(data[max].watt*data[max].durasi),data[max].durasi,data[max].ruangan)
        fmt.Println("+-------+------------------------+-----------------+--------------------+--------------------------+")
    }
}
