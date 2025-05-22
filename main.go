package main

import (
	"fmt"
	"time"
)

// struct catatan belajar
type Catatan struct {
	judul     string
	topik     string
	isiMateri string
	waktu     time.Time
	ID        int
}

// struct jadwal belajar
type Jadwal struct {
	hari         string
	mulaiBelajar string
	akhirBelajar string
	topikBelajar string
}

const catatanMax = 100
const jadwalMax = 50

// database yang menampung catatan dan jadwal belajar
var (
	catatanData       [catatanMax]Catatan
	jadwalBelajarData [jadwalMax]Jadwal
	catatanAda        = 0
	jadwalAda         = 0
	IDcatatanAkhir    = 0
	IDjadwalAkhir     = 0
)

func main() {
	for {
		fmt.Println("\nSelamat datang di Aplikasi AI Manajemen Belajar")
		fmt.Println("1. Kelola catatan belajar")
		fmt.Println("2. Kelola jadwal belajar")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			kelolaCatatan()
		case 2:
			kelolaJadwal()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}

func addCatatan(c Catatan) {
	if catatanAda >= catatanMax {
		fmt.Println("Catatan sudah penuh!")
		return
	}
	catatanData[catatanAda] = c
	catatanAda++
}

func addJadwal(j Jadwal) {
	if jadwalAda >= jadwalMax {
		fmt.Println("Jadwal sudah penuh!")
		return
	}
	jadwalBelajarData[jadwalAda] = j
	jadwalAda++
}

func kelolaCatatan() {
	for {
		fmt.Println("\nMenu Catatan Belajar")
		fmt.Println("1. Tambah catatan")
		fmt.Println("2. Lihat catatan")
		fmt.Println("3. Ubah catatan")
		fmt.Println("4. Hapus catatan")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			tambahCatatan()
		case 2:
			lihatSemuaCatatan()
		case 3:
			editCatatan()
		case 4:
			hapusCatatan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahCatatan() {
	fmt.Println("\nTambah Catatan Baru")

	IDcatatanAkhir++
	var catatanBaru Catatan
	catatanBaru.ID = IDcatatanAkhir

	fmt.Print("Judul catatan: ")
	fmt.Scan(&catatanBaru.judul)

	fmt.Print("Topik: ")
	fmt.Scan(&catatanBaru.topik)

	fmt.Print("Isi materi: ")
	fmt.Scan(&catatanBaru.isiMateri)

	catatanBaru.waktu = time.Now()
	addCatatan(catatanBaru)
	fmt.Println("Catatan berhasil ditambahkan!")
}

func lihatSemuaCatatan() {
	var i int
	fmt.Println("\nDaftar Catatan")
	if catatanAda == 0 {
		fmt.Println("Belum ada catatan")
		return
	}

	for i = 0; i < catatanAda; i++ {
		catatan := catatanData[i]
		fmt.Printf("\nID: %d\nJudul: %s\nTopik: %s\nTanggal: %s\nIsi: %s\n",
			catatan.ID, catatan.judul, catatan.topik,
			catatan.waktu.Format("2006-01-02 15:04"), catatan.isiMateri)
	}
}

func editCatatan() {
	var i, index int
	fmt.Println("\nEdit Catatan")
	fmt.Print("Masukkan ID catatan: ")
	var id int
	fmt.Scan(&id)

	index = -1
	for i = 0; i < catatanAda; i++ {
		if catatanData[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Catatan tidak ditemukan")
		return
	}

	catatan := &catatanData[index]

	fmt.Printf("Judul (%s): ", catatan.judul)
	var judul string
	fmt.Scan(&judul)
	if judul != "" {
		catatan.judul = judul
	}

	fmt.Printf("Topik (%s): ", catatan.topik)
	var topik string
	fmt.Scan(&topik)
	if topik != "" {
		catatan.topik = topik
	}

	fmt.Printf("Isi (%s): ", catatan.isiMateri)
	var isi string
	fmt.Scan(&isi)
	if isi != "" {
		catatan.isiMateri = isi
	}

	fmt.Println("Catatan berhasil diperbarui")
}

func hapusCatatan() {
	var i, index int
	fmt.Println("\nHapus Catatan")
	fmt.Print("Masukkan ID catatan: ")
	var id int
	fmt.Scan(&id)

	index = -1
	for i = 0; i < catatanAda; i++ {
		if catatanData[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Catatan tidak ditemukan")
		return
	}

	for i = index; i < catatanAda-1; i++ {
		catatanData[i] = catatanData[i+1]
	}
	catatanAda--

	fmt.Println("Catatan berhasil dihapus")
}

func kelolaJadwal() {
	for {
		fmt.Println("\nMenu Jadwal Belajar")
		fmt.Println("1. Lihat jadwal")
		fmt.Println("2. Tambah jadwal")
		fmt.Println("3. Hapus jadwal")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			lihatJadwal()
		case 2:
			tambahJadwal()
		case 3:
			hapusJadwal()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func lihatJadwal() {
	var i int
	fmt.Println("\nDaftar Jadwal")
	if jadwalAda == 0 {
		fmt.Println("Belum ada jadwal")
		return
	}

	for i = 0; i < jadwalAda; i++ {
		jadwal := jadwalBelajarData[i]
		fmt.Printf("\nHari: %s\nWaktu: %s - %s\nTopik: %s\n",
			jadwal.hari, jadwal.mulaiBelajar, jadwal.akhirBelajar, jadwal.topikBelajar)
	}
}

func tambahJadwal() {
	fmt.Println("\nTambah Jadwal Baru")

	if jadwalAda >= jadwalMax {
		fmt.Println("Jadwal sudah penuh!")
		return
	}

	var jadwalBaru Jadwal

	fmt.Print("Hari (contoh: Senin): ")
	fmt.Scan(&jadwalBaru.hari)

	fmt.Print("Waktu mulai (HH:MM): ")
	fmt.Scan(&jadwalBaru.mulaiBelajar)

	if !validasiWaktu(jadwalBaru.mulaiBelajar) {
		fmt.Println("Format waktu salah! Gunakan HH:MM")
		return
	}

	fmt.Print("Waktu selesai (HH:MM): ")
	fmt.Scan(&jadwalBaru.akhirBelajar)

	if !validasiWaktu(jadwalBaru.akhirBelajar) {
		fmt.Println("Format waktu salah! Gunakan HH:MM")
		return
	}

	fmt.Print("Topik belajar: ")
	fmt.Scan(&jadwalBaru.topikBelajar)

	addJadwal(jadwalBaru)
	fmt.Println("Jadwal berhasil ditambahkan!")
}

func validasiWaktu(waktu string) bool {
	var i int
	if len(waktu) != 5 {
		return false
	}
	if waktu[2] != ':' {
		return false
	}

	for i = 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		if waktu[i] < '0' || waktu[i] > '9' {
			return false
		}
	}

	jam := (waktu[0]-'0')*10 + (waktu[1] - '0')
	menit := (waktu[3]-'0')*10 + (waktu[4] - '0')

	return jam < 24 && menit < 60
}

func hapusJadwal() {
	fmt.Println("\nHapus Jadwal")
	fmt.Print("Masukkan nomor urut jadwal: ")
	var nomor int
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jadwalAda {
		fmt.Println("Nomor tidak valid!")
		return
	}

	for i := nomor - 1; i < jadwalAda-1; i++ {
		jadwalBelajarData[i] = jadwalBelajarData[i+1]
	}
	jadwalAda--

	fmt.Println("Jadwal berhasil dihapus")
}
