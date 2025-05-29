package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/genai"
)

// user sudah bisa mengelola catatan & jadwal
// user sudah bisa mencari catatan dengan opsi "Cari materi" (sequential belum di test)
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
	catatanAda        int = 0
	jadwalAda         int = 0
	IDcatatanAkhir    int = 0
	IDjadwalAkhir     int = 0
)

// struct untuk soal pilihan ganda
type Soal struct {
	Pertanyaan string
	Opsi       [4]string
	Jawaban    string // A, B, C, atau D
}

// dummy 15 soal pilihan ganda bahasa Inggris
var soalDummy = []Soal{
	{
		Pertanyaan: "What is the synonym of 'happy'?",
		Opsi:       [4]string{"Sad", "Joyful", "Angry", "Tired"},
		Jawaban:    "B",
	},
	{
		Pertanyaan: "Choose the correct sentence.",
		Opsi:       [4]string{"She go to school.", "She goes to school.", "She going to school.", "She gone to school."},
		Jawaban:    "B",
	},
	{
		Pertanyaan: "What is the past tense of 'run'?",
		Opsi:       [4]string{"Run", "Runned", "Ran", "Running"},
		Jawaban:    "C",
	},
	{
		Pertanyaan: "Fill in the blank: I ___ a book.",
		Opsi:       [4]string{"am reading", "is reading", "are reading", "reading"},
		Jawaban:    "A",
	},
	{
		Pertanyaan: "Which one is a noun?",
		Opsi:       [4]string{"Run", "Beautiful", "Dog", "Quickly"},
		Jawaban:    "C",
	},
	{
		Pertanyaan: "What does 'fast' mean?",
		Opsi:       [4]string{"Slow", "Quick", "Big", "Small"},
		Jawaban:    "B",
	},
	{
		Pertanyaan: "Select the correct plural form of 'child'.",
		Opsi:       [4]string{"Childs", "Childes", "Children", "Child"},
		Jawaban:    "C",
	},
	{
		Pertanyaan: "What is the antonym of 'cold'?",
		Opsi:       [4]string{"Hot", "Warm", "Cool", "Freeze"},
		Jawaban:    "A",
	},
	{
		Pertanyaan: "Choose the correct question: ___ you like coffee?",
		Opsi:       [4]string{"Do", "Does", "Did", "Are"},
		Jawaban:    "A",
	},
	{
		Pertanyaan: "What is the correct form of the verb in past: 'He ___ to the market yesterday.'",
		Opsi:       [4]string{"Go", "Goes", "Went", "Going"},
		Jawaban:    "C",
	},
	{
		Pertanyaan: "Select the correct sentence:",
		Opsi:       [4]string{"They is happy.", "They are happy.", "They am happy.", "They be happy."},
		Jawaban:    "B",
	},
	{
		Pertanyaan: "What does the word 'quickly' describe?",
		Opsi:       [4]string{"Verb", "Adjective", "Adverb", "Noun"},
		Jawaban:    "C",
	},
	{
		Pertanyaan: "Fill in the blank: She has ___ apple.",
		Opsi:       [4]string{"an", "a", "the", "some"},
		Jawaban:    "A",
	},
	{
		Pertanyaan: "What is the correct comparative form of 'good'?",
		Opsi:       [4]string{"Gooder", "Better", "More good", "Best"},
		Jawaban:    "B",
	},
	{
		Pertanyaan: "Choose the correct word: I ___ like swimming.",
		Opsi:       [4]string{"don't", "doesn't", "didn't", "not"},
		Jawaban:    "A",
	},
}

var reader = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Println("\nSelamat datang di Aplikasi AI Manajemen Belajar")
		fmt.Println("1. Kelola catatan belajar")
		fmt.Println("2. Kelola jadwal belajar")
		fmt.Println("3. Cari materi")
		fmt.Println("4. Generate soal berdasarkan materi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		input, err := readInt()
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi!")
			continue
		}

		switch input {
		case 1:
			kelolaCatatan()
		case 2:
			kelolaJadwal()
		case 3:
			cariMateri()
		case 4:
			generateSoal()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
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

		input, err := readInt()
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi!")
			continue
		}

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

	fmt.Println("\nTambah Catatan Baru")

	IDcatatanAkhir++
	catatanBaru.ID = IDcatatanAkhir

	fmt.Print("Judul catatan: ")
	judul, _ := readLine()
	catatanBaru.judul = judul

	fmt.Print("Topik: ")
	topik, _ := readLine()
	catatanBaru.topik = topik

	fmt.Print("Isi materi: ")
	isi, _ := readLine()
	catatanBaru.isiMateri = isi

	catatanBaru.waktu = time.Now()

	addCatatan(catatanBaru)
	fmt.Println("Catatan berhasil ditambahkan!")
}

func lihatSemuaCatatan() {
	if catatanAda == 0 {
		fmt.Println("\nBelum ada catatan")
		return
	}
	fmt.Println("\nDaftar Catatan")
	for i := 0; i < catatanAda; i++ {
		catatanDetail(catatanData[i])
	}
}

func catatanDetail(c Catatan) {
	fmt.Printf("\nID: %d\nJudul: %s\nTopik: %s\nTanggal: %s\nIsi: %s\n",
		c.ID, c.judul, c.topik, c.waktu.Format("2006-01-02 15:04"), c.isiMateri)
}

func editCatatan() {
	fmt.Println("\nEdit Catatan")
	fmt.Print("Masukkan ID catatan: ")
	id, err := readInt()
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	index := -1
	for i := 0; i < catatanAda; i++ {
		if catatanData[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Catatan tidak ditemukan")
		return
	}

	c := &catatanData[index]

	fmt.Printf("Judul (%s): ", c.judul)
	judul, _ := readLine()
	if judul != "" {
		c.judul = judul
	}

	fmt.Printf("Topik (%s): ", c.topik)
	topik, _ := readLine()
	if topik != "" {
		c.topik = topik
	}

	fmt.Printf("Isi (%s): ", c.isiMateri)
	isi, _ := readLine()
	if isi != "" {
		c.isiMateri = isi
	}

	fmt.Println("Catatan berhasil diperbarui")
}

func hapusCatatan() {
	fmt.Println("\nHapus Catatan")
	fmt.Print("Masukkan ID catatan: ")
	id, err := readInt()
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	index := -1
	for i := 0; i < catatanAda; i++ {
		if catatanData[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Catatan tidak ditemukan")
		return
	}

	for i := index; i < catatanAda-1; i++ {
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

		input, err := readInt()
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi!")
			continue
		}

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
		var jadwal Jadwal = jadwalBelajarData[i]
		fmt.Printf("\nHari: %s\nWaktu: %s - %s\nTopik: %s\n",
			jadwal.hari, jadwal.mulaiBelajar, jadwal.akhirBelajar, jadwal.topikBelajar)
	}
}

func tambahJadwal() {
	if jadwalAda >= jadwalMax {
		fmt.Println("Jadwal sudah penuh!")
		return
	}

	fmt.Println("\nTambah Jadwal Baru")

	fmt.Print("Hari (contoh: Senin): ")
	hari, _ := readLine()

	fmt.Print("Waktu mulai (HH:MM): ")
	mulai, _ := readLine()
	if !validasiWaktu(mulai) {
		fmt.Println("Format waktu salah! Gunakan HH:MM")
		return
	}

	fmt.Print("Waktu selesai (HH:MM): ")
	selesai, _ := readLine()
	if !validasiWaktu(selesai) {
		fmt.Println("Format waktu salah! Gunakan HH:MM")
		return
	}

	fmt.Print("Topik belajar: ")
	topik, _ := readLine()

	jadwalBaru := Jadwal{
		hari:         hari,
		mulaiBelajar: mulai,
		akhirBelajar: selesai,
		topikBelajar: topik,
	}

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
	fmt.Println("\nHapus Jadwal")
	fmt.Print("Masukkan nomor urut jadwal: ")
	nomor, err := readInt()
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

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

func cariMateri() {
	for {
		fmt.Println("\nPilih opsi pencarian materi")
		fmt.Println("1. Berurutan (sequential)")
		fmt.Println("2. Binary search")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")

		input, err := readInt()
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi!")
			continue
		}

		switch input {
		case 1:
			fmt.Print("\nKata Kunci: ")
			keyWord, _ := readLine()
			keyWord = strings.ToLower(keyWord)
			searchSequential(keyWord)
		case 2:
			fmt.Print("\nKata Kunci: ")
			keyWord, _ := readLine()
			keyWord = strings.ToLower(keyWord)
			searchBinary(keyWord)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func searchSequential(keyWord string) {
	var i int
	var found bool = false

	fmt.Println("\nHasil pencarian:")

	for i = 0; i < catatanAda; i++ {
		var catatan Catatan = catatanData[i]
		if strings.Contains(strings.ToLower(catatan.judul), keyWord) ||
			strings.Contains(strings.ToLower(catatan.topik), keyWord) ||
			strings.Contains(strings.ToLower(catatan.isiMateri), keyWord) {
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan hasil")
	}
}

func searchBinary(keyWord string) {
	var catatanTerurut []Catatan = make([]Catatan, catatanAda)
	var i int
	for i = 0; i < catatanAda; i++ {
		catatanTerurut[i] = catatanData[i]
	}

	var n int = catatanAda
	for i = 0; i < n-1; i++ {
		var idxMin int = i
		for j := i + 1; j < n; j++ {
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

	fmt.Println("\nHasil Pencarian:")

	for low <= high {
		var mid int = low + (high-low)/2
		var current string = strings.ToLower(catatanTerurut[mid].judul)

		if strings.Contains(current, keyWord) {
			catatanDetail(catatanTerurut[mid])
			found = true

			var left int = mid - 1
			for left >= 0 && strings.Contains(strings.ToLower(catatanTerurut[left].judul), keyWord) {
				catatanDetail(catatanTerurut[left])
				left--
			}

			var right int = mid + 1
			for right < catatanAda && strings.Contains(strings.ToLower(catatanTerurut[right].judul), keyWord) {
				catatanDetail(catatanTerurut[right])
				right++
			}

			return
		} else if catatanTerurut[mid].judul < keyWord {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan hasil")
	}
}

func generateSoal() {
	fmt.Println("\nHalo! Saya AI Manajemen Belajar. ada yang bisa saya bantu?")
	fmt.Println("Ketik apa saja untuk mulai, atau ketik 'keluar' untuk kembali ke menu utama.")

	for {
		fmt.Print("Anda: ")
		input, err := readLine()
		if err != nil {
			fmt.Println("Gagal membaca input, coba lagi.")
			continue
		}

		input = strings.TrimSpace(strings.ToLower(input))
		if input == "keluar" {
			fmt.Println("Kembali ke menu utama...")
			return
		}

		// input selain "keluar", tampilkan soal
		for {
			rand.Seed(time.Now().UnixNano())
			fmt.Println("\nBerikut 5 soal latihan bahasa Inggris:")

			indices := rand.Perm(len(soalDummy))[:5]

			for i, idx := range indices {
				soal := soalDummy[idx]
				fmt.Printf("\nSoal %d: %s\n", i+1, soal.Pertanyaan)
				fmt.Printf("A. %s\n", soal.Opsi[0])
				fmt.Printf("B. %s\n", soal.Opsi[1])
				fmt.Printf("C. %s\n", soal.Opsi[2])
				fmt.Printf("D. %s\n", soal.Opsi[3])
			}
			fmt.Println("\nJawaban benar ditandai dengan huruf (contoh: B).")
			fmt.Println("\nKetik apa saja untuk generate soal baru, atau ketik 'keluar' untuk kembali ke menu utama.")

			fmt.Print("Anda: ")
			input, err = readLine()
			if err != nil {
				fmt.Println("Gagal membaca input, coba lagi.")
				continue
			}

			if strings.TrimSpace(strings.ToLower(input)) == "keluar" {
				fmt.Println("Kembali ke menu utama...")
				return
			}
		}
	}
}

// Fungsi generate 10 soal berdasarkan materi input
func generateSoalByGemini() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "AIzaSyCdtVoP4zjaHOiwfdOxOvbBBsFW9AVsNB0", // Ganti dengan API key Gemini yang valid
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Gagal membuat client: %v", err)
	}

	fmt.Println("\nMasukkan materi untuk generate soal (bahasa Indonesia):")
	reader := bufio.NewScanner(os.Stdin)
	if !reader.Scan() {
		fmt.Println("Gagal membaca input")
		return
	}
	materi := reader.Text()

	fmt.Println("Mohon tunggu, sedang memproses generate soal...")  // <-- info tunggu

	prompt := fmt.Sprintf(`Buatkan 10 soal pilihan ganda dalam bahasa Indonesia berdasarkan materi berikut:
"%s"
Setiap soal terdiri dari satu pertanyaan dan 4 opsi jawaban (A, B, C, D).
Tandai jawaban yang benar dengan jelas.`, materi)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatalf("Gagal generate soal: %v", err)
	}

	fmt.Println("\nSoal yang dihasilkan:")
	fmt.Println(result.Text())
}

// Fungsi helper baca baris string
func readLine() (string, error) {
	if reader.Scan() {
		return reader.Text(), nil
	}
	return "", fmt.Errorf("Gagal baca input")
}

// Fungsi helper baca int dari string input
func readInt() (int, error) {
	line, err := readLine()
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, err
	}
	return i, nil
}
