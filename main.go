package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Catatan struct {
	judul     string
	topik     string
	isiMateri string
	waktu     time.Time
	ID        int
}

type Jadwal struct {
	hari         string
	mulaiBelajar string
	akhirBelajar string
	topikBelajar string
}

const catatanMax = 100
const jadwalMax = 50

var (
	catatanData       [catatanMax]Catatan
	jadwalBelajarData [jadwalMax]Jadwal
	catatanAda        int = 0
	jadwalAda         int = 0
	IDcatatanAkhir    int = 0
	IDjadwalAkhir     int = 0
)

type Soal struct {
	soalPertanyaan string
}

func main() {
	var input int
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nSelamat datang di Aplikasi AI Manajemen Belajar,")
		fmt.Println("Ada yang bisa saya bantu?")
		fmt.Println("================================================")
		fmt.Println("1. Kelola catatan materi belajar")
		fmt.Println("2. Kelola jadwal belajar")
		fmt.Println("3. Cari materi")
		fmt.Println("4. Buat soal latihan")
		fmt.Println("0. Keluar")
		fmt.Println("================================================")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&input)
		reader.Scan()

		switch input {
		case 1:
			kelolaCatatan()
		case 2:
			kelolaJadwal()
		case 3:
			cariMateri()
		case 4:
			buatSoal()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini, senang membantu anda!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}

func kelolaCatatan() {
	var input int
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nAda yang bisa saya bantu untuk Catatan anda?")
		fmt.Println("============================================")
		fmt.Println("1. Tambah catatan")
		fmt.Println("2. Lihat catatan")
		fmt.Println("3. Ubah catatan")
		fmt.Println("4. Hapus catatan")
		fmt.Println("0. Kembali")
		fmt.Println("============================================")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&input)
		reader.Scan()

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

func addCatatan(c Catatan) {
	if catatanAda >= catatanMax {
		fmt.Println("Catatan sudah penuh!")
		return
	}
	catatanData[catatanAda] = c
	catatanAda++
}

func tambahCatatan() {
	var catatanBaru Catatan
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var judul string
	var topik string
	var isi string

	fmt.Println("\nTambah Catatan Baru")

	IDcatatanAkhir++
	catatanBaru.ID = IDcatatanAkhir

	fmt.Print("Judul catatan: ")
	reader.Scan()
	judul = reader.Text()
	catatanBaru.judul = judul

	fmt.Print("Topik: ")
	reader.Scan()
	topik = reader.Text()
	catatanBaru.topik = topik

	fmt.Print("Isi materi: ")
	reader.Scan()
	isi = reader.Text()
	catatanBaru.isiMateri = isi

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
		var catatan Catatan = catatanData[i]
		catatanDetail(catatan)
	}
}

func catatanDetail(c Catatan) {
	fmt.Printf("\nID: %d\nJudul: %s\nTopik: %s\nTanggal: %s\nIsi: %s\n",
		c.ID, c.judul, c.topik, c.waktu.Format("2006-01-02 15:04"), c.isiMateri)
}

func editCatatan() {
	var i int
	var index int = -1
	var id int
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var judul string
	var topik string
	var isi string

	fmt.Println("\nEdit Catatan")
	fmt.Print("Masukkan ID catatan: ")
	fmt.Scan(&id)
	reader.Scan()

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

	var catatan *Catatan = &catatanData[index]

	fmt.Printf("Judul (%s): ", catatan.judul)
	reader.Scan()
	judul = reader.Text()
	if judul != "" {
		catatan.judul = judul
	}

	fmt.Printf("Topik (%s): ", catatan.topik)
	reader.Scan()
	topik = reader.Text()
	if topik != "" {
		catatan.topik = topik
	}

	fmt.Printf("Isi (%s): ", catatan.isiMateri)
	reader.Scan()
	isi = reader.Text()
	if isi != "" {
		catatan.isiMateri = isi
	}

	fmt.Println("Catatan berhasil diperbarui")
}

func hapusCatatan() {
	var i int
	var index int = -1
	var id int

	fmt.Println("\nHapus Catatan")
	fmt.Print("Masukkan ID catatan: ")
	fmt.Scan(&id)

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
	var input int
	for {
		fmt.Println("\nAda yang bisa saya bantu untuk Jadwal anda?")
		fmt.Println("===========================================")
		fmt.Println("1. Tambah jadwal")
		fmt.Println("2. Lihat jadwal")
		fmt.Println("3. Hapus jadwal")
		fmt.Println("0. Kembali")
		fmt.Println("===========================================")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&input)

		switch input {
		case 1:
			tambahJadwal()
		case 2:
			lihatJadwal()
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
		var jadwal Jadwal = jadwalBelajarData[i]
		fmt.Printf("\nHari: %s\nWaktu: %s - %s\nTopik: %s\n",
			jadwal.hari, jadwal.mulaiBelajar, jadwal.akhirBelajar, jadwal.topikBelajar)
	}
}

func tambahJadwal() {
	var jadwalBaru Jadwal

	fmt.Println("\nTambah Jadwal Baru")

	if jadwalAda >= jadwalMax {
		fmt.Println("Jadwal sudah penuh!")
		return
	}

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

func addJadwal(j Jadwal) {
	if jadwalAda >= jadwalMax {
		fmt.Println("Jadwal sudah penuh!")
		return
	}
	jadwalBelajarData[jadwalAda] = j
	jadwalAda++
}

func validasiWaktu(waktu string) bool {
	var i int
	var jam int
	var menit int

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

	jam = (int(waktu[0])-'0')*10 + (int(waktu[1]) - '0')
	menit = (int(waktu[3])-'0')*10 + (int(waktu[4]) - '0')

	return jam < 24 && menit < 60
}

func hapusJadwal() {
	var nomor int
	var i int

	fmt.Println("\nHapus Jadwal")
	fmt.Print("Masukkan nomor urut jadwal: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jadwalAda {
		fmt.Println("Nomor tidak valid!")
		return
	}

	for i = nomor - 1; i < jadwalAda-1; i++ {
		jadwalBelajarData[i] = jadwalBelajarData[i+1]
	}
	jadwalAda--

	fmt.Println("Jadwal berhasil dihapus")
}

func cariMateri() {
	var input int
	var keyWord string
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nSilahkan pilih metode pencarian materi berdasarkan Judul/Topiknya")
		fmt.Println("=================================================================")
		fmt.Println("1. Berurutan (sequential)")
		fmt.Println("2. Binary search")
		fmt.Println("0. Kembali")
		fmt.Println("=================================================================")
		fmt.Print("Pilihan: ")

		fmt.Scan(&input)
		reader.Scan()

		switch input {
		case 1:
			fmt.Print("\nKata kunci: ")
			reader.Scan()
			keyWord = strings.ToLower(reader.Text())
			cariSequential(keyWord)
		case 2:
			fmt.Print("\nKata kunci: ")
			reader.Scan()
			keyWord = strings.ToLower(reader.Text())
			cariBinary(keyWord)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func cariSequential(keyWord string) {
	var i int
	var found bool = false

	fmt.Println("\nHasil pencarian:")

	for i = 0; i < catatanAda; i++ {
		var catatan Catatan = catatanData[i]
		if strings.Contains(strings.ToLower(catatan.judul), keyWord) ||
			strings.Contains(strings.ToLower(catatan.topik), keyWord) {
			catatanDetail(catatan)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak menemukan hasil")
	}
}

func cariBinary(keyWord string) { 
	var catatanTerurut []Catatan = make([]Catatan, catatanAda)
	var i, j int
	for i = 0; i < catatanAda; i++ {
		catatanTerurut[i] = catatanData[i]
	}

	var n int = catatanAda
	for i = 0; i < n-1; i++ {
		var idxMin int = i
		for j = i + 1; j < n; j++ {
			if catatanTerurut[j].judul < catatanTerurut[idxMin].judul {
				idxMin = j
			}
		}
		if idxMin != i {
			var temp Catatan = catatanTerurut[i]
			catatanTerurut[i] = catatanTerurut[idxMin]
			catatanTerurut[idxMin] = temp
		}
	}

	var low int = 0
	var high int = catatanAda - 1
	var found bool = false

	fmt.Println("\nHasil pencarian:")

	for low <= high {
		var mid int = low + (high-low)/2
		var currentJudul string = strings.ToLower(catatanTerurut[mid].judul)
		var currentTopik string = strings.ToLower(catatanTerurut[mid].topik)

		if strings.Contains(currentJudul, keyWord) || strings.Contains(currentTopik, keyWord) {
			catatanDetail(catatanTerurut[mid])
			found = true

			var left int = mid - 1
			for left >= 0 {
				currentJudul = strings.ToLower(catatanTerurut[left].judul)
				currentTopik = strings.ToLower(catatanTerurut[left].topik)
				if strings.Contains(currentJudul, keyWord) || strings.Contains(currentTopik, keyWord) {
					catatanDetail(catatanTerurut[left])
					left--
				} else {
					break
				}
			}

			var right int = mid + 1
			for right < catatanAda {
				currentJudul = strings.ToLower(catatanTerurut[right].judul)
				currentTopik = strings.ToLower(catatanTerurut[right].topik)
				if strings.Contains(currentJudul, keyWord) || strings.Contains(currentTopik, keyWord) {
					catatanDetail(catatanTerurut[right])
					right++
				} else {
					break
				}
			}
			return
		} else if catatanTerurut[mid].judul < keyWord {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak menemukan hasil")
	}
}

var daftarSoal = []Soal{ 
	{"Apa ibukota Indonesia?"},
	{"Siapa presiden pertama Indonesia?"},
	{"Apa warna langit pada siang hari yang cerah?"},
	{"Planet terdekat dari matahari?"},
	{"Apa lambang kimia untuk emas?"},
	{"Berapa banyak sisi yang dimiliki segitiga?"},
	{"Apa nama benua terbesar di dunia?"},
	{"Bahasa pemrograman yang dikembangkan Google pada 2009?"},
	{"Apa nama gunung tertinggi di dunia?"},
}

func buatSoal() {
	var i int
	fmt.Println("\nBerikut soal-soal latihannya")

	if len(daftarSoal) < 3 {
		fmt.Println("Maaf, soal yang tersedia tidak cukup!")
		return
	}

	var soalTampil [3]Soal
	var indeksAcak int
	var sudahDipilih = make(map[int]bool)

	for i = 0; i < 3; i++ {
		for {
			indeksAcak = rand.Intn(len(daftarSoal))
			if !sudahDipilih[indeksAcak] {
				sudahDipilih[indeksAcak] = true
				soalTampil[i] = daftarSoal[indeksAcak]
				break
			}
		}
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("\n%d: %s\n", i+1, soalTampil[i].soalPertanyaan)
	}
}
