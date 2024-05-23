package main

import (
	"fmt"
)

const NMAX = 100000

type ArrayGuru struct {
	TampG [NMAX]Mahasiswa
	n     int
}

type Guru struct {
	Nama   string
	Course string
}

type Mahasiswa struct {
	Nama        string
	DaftarNilai ArrayPerCourse
}

type ArrayMahasiswa struct {
	TampM [NMAX]Mahasiswa
	n     int
}

// Data mahasiswa

type PerCourse struct {
	NamaCourse string
	NilaiTugas int
	NilaiQuiz  int
}

type ArrayPerCourse struct {
	TampC [NMAX]PerCourse
	n     int
}

// Data soal
type Soal struct {
	TipeSoal   string // apakah bertipe "Tugas" atau "Quiz" atau "Forum"
	Pertanyaan string
	Jawaban    string
}

type ArrayTugas struct {
	TampT [NMAX]Soal
	n     int
}

type ArrayQuiz struct {
	TampQ [NMAX]Soal
	n     int
}

type Forum struct {
	Pertanyaan string
	Jawaban    string
	Pengirim   string
}

type ArrayForum struct {
	TampF [NMAX]Forum
	n     int
}

// Data course
type Course struct {
	Nama            string
	Guru            string
	DaftarMahasiswa [NMAX]Mahasiswa
	SoalTugas       ArraySoal
	SoalQuiz        ArraySoal
	Forum           ArrayForum
}

func menu(A *ArrayMahasiswa, B *ArrayGuru) {
	var opsi int
	var input string
	for opsi != 3 {
		fmt.Println("LMS SEDERHANA")
		fmt.Println("Silahkan Pilih Opsi")
		fmt.Println("1. Login LMS")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Println("Pilih opsi: ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			fmt.Scan(&input)
			fmt.Println("Salam Bahagia, Silahkan Pilih Status Pengguna")
			fmt.Println("1. Siswa")
			fmt.Println("2. Guru")
			if input == "Siswa" || input == "siswa" {
				loginsiswa(A, input)
			} else if input == "Guru" || input == "guru" {
				loginguru(B, input)
			}
		case 2:
			fmt.Scan(&input)
			fmt.Println("Salam Bahagia, Silahkan Pilih Status Pengguna")
			fmt.Println("1. Siswa")
			fmt.Println("2. Guru")
			if input == "Siswa" || input == "siswa" {
				registersiswa(A)
			} else if input == "Guru" || input == "guru" {
				registerguru(B)
			}
		}
	}
}

func registersiswa(A *ArrayMahasiswa) {
	var user string
	fmt.Scan(&user)
	A.TampM[A.n].Nama = user
	A.n++
}

func loginsiswa(A *ArrayMahasiswa, input string) {
	var pilih int
	fmt.Println("Silahkan Login!")
	if A.TampM[A.n].Nama == input {
		course()
	} else {
		fmt.Println("Data Tidak Ditemukan, Silahkan Register atau Login Ulang.")
		fmt.Println("1. Regisiter.")
		fmt.Println("2. Login.")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			registersiswa(A)
		case 2:
			loginsiswa(A, input)
		}
	}
}

func registerguru(B *ArrayGuru) {
	var user string
	fmt.Scan(&user)
	B.TampG[B.n].Nama = user
	B.n++
}

func loginguru(B *ArrayGuru, input string) {
	var pilih int
	fmt.Println("Silahkan Login!")
	if B.TampG[B.n].Nama == input {
		gurumenu()
	} else {
		fmt.Println("Data Tidak Ditemukan, Silahkan Register atau Login Ulang.")
		fmt.Println("1. Regisiter.")
		fmt.Println("2. Login.")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			registerguru(B)
		case 2:
			loginguru(B, input)
		}
	}
}

func course() {
	var pilihan int

	for pilihan != 4 {
		fmt.Println("\nMenu Siswa:")
		fmt.Println("1. Soal")
		fmt.Println("2. Quiz ")
		fmt.Println("3. Forum")
		fmt.Println("4. Keluar")
		fmt.Println("Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			soal()
		case 2:
			quiz()
		case 3:
			forum()
		case 4:
			menu(&A, &B)
		}
	}
}

func gurumenu() {
	var pilihan int

	for pilihan != 6 {
		fmt.Println("\nMenu Guru:")
		fmt.Println("1. Tambah Soal")
		fmt.Println("2. Ubah Soal")
		fmt.Println("3. Hapus Soal")
		fmt.Println("4. Tambah Forum")
		fmt.Println("5. Tampilkan Data Mahasiswa")
		fmt.Println("6. Keluar")
		fmt.Println("Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:

		case 2:

		}
	}
}

func tambahSoal(A *ArrayTugas, B *ArrayQuiz) {
	var t, p, j, Lan string // t adalah tipe soal(quiz atau tugas)
	fmt.Scan(&t)
	if t == "Quiz" {
		B.TampQ[A.n].TipeSoal = t
	} else {
		A.TampT[A.n].TipeSoal = t
	}
	for Lan != "Tidak" {
		fmt.Println("Silahkan Masukkan Pertanyaan dan Jawaban.")
		fmt.Scan(&p, &j) // p adalah pertanyaan dan j adalah jawaban
		if t == "Tugas" {
			A.TampT[A.n].Pertanyaan = p
			A.TampT[A.n].Jawaban = j
		} else {
			B.TampQ[A.n].Pertanyaan = p
			B.TampQ[A.n].Jawaban = j
		}
		A.n++
		fmt.Println("Soal Telah Ditambahkan!")
		fmt.Println("Apakah Anda Ingin Melanjutkan Lagi?")
		fmt.Println("1. Iya")
		fmt.Println("2. Tidak")
		fmt.Scan(&Lan)
		// if Lan == "Iya" {
		// 	tambahSoal(A)
		// }
	}
}

func tambahForum(A *ArrayForum) {
	var user, p, Lan string // user adalah pengirim
	fmt.Scan(&user)
	A.TampF[A.n].Pengirim = user
	for Lan != "Tidak" {
		fmt.Println("Silahkan Masukkan Pertanyaan.")
		fmt.Scan(&p) // p adalah pertanyaan
		A.TampF[A.n].Pertanyaan = p
		A.n++
		fmt.Println("Soal Telah Ditambahkan!")
		fmt.Println("Apakah Anda Ingin Melanjutkan Lagi?")
		fmt.Println("1. Iya")
		fmt.Println("2. Tidak")
		fmt.Scan(&Lan)
		// if Lan == "Iya" {
		// 	tambahForum(A)
		// }
	}
}

func editSoal(A *ArrayTugas, B *ArrayQuiz) {
	var t, p, j string // t adalah tipe soal(quiz atau tugas)
	var cari, hasil int
	fmt.Println("Silahkan Pilih Tipe Soal Yang Akan Diubah.")
	fmt.Scan(&t)
	fmt.Println("Silahkan Pilih Soal Yang Akan Diubah.")
	fmt.Scan(&cari)
	if t == "Tugas" {
		hasil = seqSearchTugas(*A, A.n, cari)
	} else {
		hasil = seqSearchQuiz(*B, A.n, cari)
	}
	if hasil != -1 {
		fmt.Println("Silahkan Masukkan Pertanyaan dan Jawaban Baru.")
		fmt.Scan(&p, &j)
		if t == "Tugas" {
			A.TampT[hasil].Pertanyaan = p
			A.TampT[hasil].Jawaban = j
		} else {
			B.TampQ[hasil].Pertanyaan = p
			B.TampQ[hasil].Jawaban = j
		}
	}
}

func editForum(A *ArrayForum) {
	var user, p string // t adalah tipe soal(quiz atau tugas)
	var cari, hasil int
	fmt.Println("Silahkan Masukkan Nama Pengguna.")
	fmt.Scan(&user)
	fmt.Println("Silahkan Pilih Soal Yang Akan Diubah.")
	fmt.Scan(&cari)
	hasil = seqSearchSoal(A, n, cari)
	if hasil != -1 {
		fmt.Println("Silahkan Masukkan Pertanyaan Baru.")
		fmt.Scan(&p)
		A.TampF[hasil].Pertanyaan = p
	}
}

func seqSearchLogin(A ArrayMahasiswa, n int, x string) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < n {
		if A.TampM[k].Nama == x {
			found = k
			k++
		} else {
			found = -1
		}
	}
	return found
}

func binSearch(A ArrayMahasiswa, n int, x string) int {
	low := 0
	high := n - 1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		if A.TampM[mid].Nama < x {
			low = mid + 1
		} else if A.TampM[mid].Nama > x {
			high = mid - 1
		} else {
			found = mid
		}
	}
	return found
}

func seqSearchTugas(A ArrayTugas, n int, x int) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < n {
		if k == x {
			found = k
			k++
		} else {
			found = -1
		}
	}
	return found
}

func seqSearchQuiz(A ArrayQuiz, n int, x int) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < n {
		if k == x {
			found = k
			k++
		} else {
			found = -1
		}
	}
	return found
}

func binSearchLogin(A ArrayMahasiswa, n int, x string) int {
	low := 0
	high := n - 1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		if A.TampM[mid].Nama < x {
			low = mid + 1
		} else if A.TampM[mid].Nama > x {
			high = mid - 1
		} else {
			found = mid
		}
	}
	return found
}

func main() {
	var PSiswa ArrayMahasiswa
	var PGuru ArrayGuru
	var BTugas ArrayTugas
	var BQuiz ArrayQuiz
	var BForum ArrayForum
	menu(&PSiswa, &PGuru)
}
