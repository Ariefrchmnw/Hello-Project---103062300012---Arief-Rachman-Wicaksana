// Nama  : Arief Rachman Wicaksana
// NIM   : 103062300012
// Kelas : S1IT-KJ-23-001
package main

import (
	"fmt"
	"os"
	"sort"
)

// Struktur untuk menyimpan informasi tenant (penyewa)
type Tenant struct {
	Nama            string  // Nama tenant
	TotalTransaksi  float64 // Total uang transaksi tenant
	JumlahTransaksi int     // Jumlah transaksi yang dilakukan tenant
}

// Struktur untuk menyimpan informasi transaksi
type Transaksi struct {
	NamaTenant string  // Nama tenant yang melakukan transaksi
	Jumlah     float64 // Jumlah transaksi yang dilakukan
}

// Slice untuk menyimpan data tenant dan transaksi
var tenants []Tenant
var transaksi []Transaksi

// Prosedur untuk menambahkan tenant baru ke daftar tenants
func tambahTenant(nama string) {
	tenant := Tenant{Nama: nama}      // Membuat tenant baru
	tenants = append(tenants, tenant) // Menambahkan tenant ke slice tenants
}

// Prosedur untuk mengubah nama tenant yang sudah ada
func ubahTenant(namaLama string, namaBaru string) {
	for i := range tenants { // Looping untuk mencari tenant berdasarkan nama lama
		if tenants[i].Nama == namaLama { // Jika ditemukan tenant dengan nama yang sesuai
			tenants[i].Nama = namaBaru // Mengubah nama tenant
			break
		}
	}
}

// Prosedur untuk menghapus tenant dari daftar tenants
func hapusTenant(nama string) {
	for i := range tenants { // Looping untuk mencari tenant yang akan dihapus
		if tenants[i].Nama == nama { // Jika ditemukan tenant dengan nama yang sesuai
			tenants = append(tenants[:i], tenants[i+1:]...) // Menghapus tenant dari slice
			break
		}
	}
}

// Prosedur untuk mencatat transaksi baru ke daftar transaksi
func tambahTransaksi(namaTenant string, jumlah float64) {
	transaksiBaru := Transaksi{NamaTenant: namaTenant, Jumlah: jumlah} // Membuat transaksi baru
	transaksi = append(transaksi, transaksiBaru)                       // Menambahkan transaksi ke slice transaksi

	for i := range tenants { // Looping untuk mencari tenant yang sesuai
		if tenants[i].Nama == namaTenant { // Jika tenant ditemukan
			tenants[i].TotalTransaksi += jumlah // Menambah total transaksi tenant
			tenants[i].JumlahTransaksi++        // Menambah jumlah transaksi tenant
			break
		}
	}
}

// Fungsi untuk menghitung pendapatan tenant dan admin kantin
func hitungPendapatan() ([]float64, float64) {
	pendapatanTenant := make([]float64, len(tenants)) // Membuat slice untuk menyimpan pendapatan tiap tenant
	var pendapatanAdmin float64                       // Variabel untuk menyimpan pendapatan admin

	for _, t := range transaksi { // Looping melalui setiap transaksi
		bagianTenant := t.Jumlah * 0.75  // 75% pendapatan untuk tenant
		bagianAdmin := t.Jumlah * 0.25   // 25% pendapatan untuk admin
		for i, tenant := range tenants { // Looping untuk mencocokkan tenant dengan transaksi
			if tenant.Nama == t.NamaTenant { // Jika tenant cocok dengan nama transaksi
				pendapatanTenant[i] += bagianTenant // Menambahkan bagian tenant
				break
			}
		}
		pendapatanAdmin += bagianAdmin // Menambahkan bagian admin
	}
	return pendapatanTenant, pendapatanAdmin // Mengembalikan pendapatan tenant dan admin
}

// Prosedur untuk menampilkan daftar tenant berdasarkan banyaknya transaksi
func daftarTenantBerdasarkanTransaksi() {
	sort.SliceStable(tenants, func(i, j int) bool { // Mengurutkan tenants berdasarkan jumlah transaksi (descending)
		return tenants[i].JumlahTransaksi > tenants[j].JumlahTransaksi
	})

	file, _ := os.Create("daftar_tenant.txt") // Membuat file daftar_tenant.txt
	defer file.Close()                        // Menutup file setelah selesai

	fmt.Println("Daftar Tenant berdasarkan banyak transaksi:")
	for _, tenant := range tenants { // Looping untuk menampilkan dan menulis daftar tenant
		output := fmt.Sprintf("Nama: %s, Jumlah Transaksi: %d, Total Uang: %.2f\n",
			tenant.Nama, tenant.JumlahTransaksi, tenant.TotalTransaksi)
		file.WriteString(output) // Menulis daftar tenant ke file
		fmt.Print(output)        // Menampilkan daftar tenant di terminal
	}
	fmt.Println("Daftar tenant berhasil ditulis ke daftar_tenant.txt")
}

// Prosedur untuk menampilkan pendapatan tenant dan admin ke dalam file
func tampilkanPendapatanKeFile() {
	pendapatanTenant, pendapatanAdmin := hitungPendapatan() // Memanggil fungsi untuk menghitung pendapatan
	file, _ := os.Create("pendapatan.txt")                  // Membuat file pendapatan.txt
	defer file.Close()                                      // Menutup file setelah selesai

	fmt.Println("Pendapatan Tenant:")
	for i, tenant := range tenants { // Looping untuk menampilkan dan menulis pendapatan tenant
		output := fmt.Sprintf("Tenant Nama: %s, Pendapatan: %.2f\n", tenant.Nama, pendapatanTenant[i])
		file.WriteString(output) // Menulis pendapatan tenant ke file
		fmt.Print(output)        // Menampilkan pendapatan tenant di terminal
	}
	outputAdmin := fmt.Sprintf("Pendapatan Admin: %.2f\n", pendapatanAdmin) // Format pendapatan admin
	file.WriteString(outputAdmin)                                           // Menulis pendapatan admin ke file
	fmt.Print(outputAdmin)                                                  // Menampilkan pendapatan admin di terminal

	fmt.Println("Pendapatan berhasil ditulis ke pendapatan.txt")
}

// Fungsi utama yang menyediakan menu interaktif
func main() {
	var pilihan int
	for {
		// Menampilkan menu pilihan
		fmt.Println(`
===================================
|         Menu                    |
| 1. Tambah Tenant                |
| 2. Ubah Tenant                  |
| 3. Hapus Tenant                 |
| 4. Tambah Transaksi             |
| 5. Tampilkan Pendapatan         |
| 6. Tampilkan Daftar Tenant      |
|    Berdasarkan Banyak Transaksi |
| 7. Keluar                       |
===================================
		`)
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan) // Menerima input pilihan dari pengguna

		// Mengecek pilihan pengguna dan menjalankan fungsionalitas yang sesuai
		if pilihan == 1 {
			var nama string
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&nama)
			tambahTenant(nama) // Menambahkan tenant baru
			fmt.Println("Tenant berhasil ditambahkan.")
		} else if pilihan == 2 {
			var namaLama, namaBaru string
			fmt.Print("Masukkan Nama Tenant yang ingin diubah: ")
			fmt.Scan(&namaLama)
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&namaBaru)
			ubahTenant(namaLama, namaBaru) // Mengubah nama tenant
			fmt.Println("Tenant berhasil diubah.")
		} else if pilihan == 3 {
			var nama string
			fmt.Print("Masukkan Nama Tenant yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusTenant(nama) // Menghapus tenant
			fmt.Println("Tenant berhasil dihapus.")
		} else if pilihan == 4 {
			var namaTenant string
			var jumlah float64
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&namaTenant)
			fmt.Print("Masukkan Jumlah Transaksi: ")
			fmt.Scan(&jumlah)
			tambahTransaksi(namaTenant, jumlah) // Mencatat transaksi baru
			fmt.Println("Transaksi berhasil dicatat.")
		} else if pilihan == 5 {
			tampilkanPendapatanKeFile() // Menampilkan pendapatan tenant dan admin
		} else if pilihan == 6 {
			daftarTenantBerdasarkanTransaksi() // Menampilkan daftar tenant berdasarkan transaksi
		} else if pilihan == 7 {
			fmt.Println("Keluar dari program.") // Keluar dari program
			return
		} else {
			fmt.Println("Opsi tidak valid.") // Menampilkan pesan jika opsi tidak valid
		}
	}
}
