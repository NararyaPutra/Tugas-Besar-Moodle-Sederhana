package main

import (
	"fmt"
)

const NMAX = 100

type ArrayGuru struct { // struct ubtuk menyimpan variabel TampG(tampung guru) dan n sebagai batas
	TampG [NMAX]Guru
	n     int
}

type ArrayMahasiswa struct { // struct ubtuk menyimpan variabel TampM(tampung mahasiswa) dan n sebagai batas
	TampM [NMAX]Mahasiswa
	n     int
}

type Guru struct { // struct ubtuk menyimpan variabel Nama (untuk nama guru) dan Course (untuk matkul yang di ajarkan)
	Nama     string
	Password string
}

type Mahasiswa struct { // struct untuk menyimpan variabel Nama (untuk nama mahasiswa) dan variabel Daftarnilai
	Nama        string // yang diambil dari struct ArrayPerCourse untuk nilai mahasiswa
	DaftarNilai ArrayPerCourse
	Password    string
}

// Data mahasiswa

type PerCourse struct {
	NilaiTugas int
	NilaiQuiz  int
}

type ArrayPerCourse struct {
	TampC [NMAX]PerCourse
	n     int
}

// Data soal
type Soal struct {
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
	Penjawab   string
}

type ArrayForum struct {
	TampF [NMAX]Forum
	n     int
}

func main() {
	var PSiswa ArrayMahasiswa
	var PGuru ArrayGuru
	var BTugas ArrayTugas
	var BQuiz ArrayQuiz
	var BForum ArrayForum
	var BCourse ArrayPerCourse
	menu(&PSiswa, &PGuru, &BTugas, &BQuiz, &BForum, &BCourse)
	fmt.Println("===============================================================")
	fmt.Println("			                                                    ")
	fmt.Println("			                                                    ")
	fmt.Println("           Anda Telah Log Out Dari Moodle Sederhana            ")
	fmt.Println("    Sebuah pikiran yang terdidik selalu memiliki harapan       ")
	fmt.Println("                   -Nararya dan Dzakhwan                       ")
	fmt.Println("			                                                    ")
	fmt.Println("			                                                    ")
	fmt.Println("===============================================================")
}

func menu(A *ArrayMahasiswa, B *ArrayGuru, C *ArrayTugas, D *ArrayQuiz, E *ArrayForum, F *ArrayPerCourse) {
	var opsi int
	var input string
	for opsi != 3 {
		fmt.Println("=====================================")
		fmt.Println("||          MOODLE SEDERHANA        ||")
		fmt.Println("=====================================")
		fmt.Println("||        Silahkan Pilih Opsi       ||")
		fmt.Println("||  1. 📝 Login LMS                 ||")
		fmt.Println("||  2. ✏️  Register                  ||")
		fmt.Println("||  3. ❌ Keluar                    ||")
		fmt.Println("=====================================")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			fmt.Println("\n=====================================")
			fmt.Println("||       Salam Bahagia, Silahkan    ||")
			fmt.Println("||  Tuliskan Status Pengguna Dari   ||")
			fmt.Println("||         Pilihan Berikut.         ||")
			fmt.Println("=====================================")
			fmt.Println("||  1. 👨‍🎓 Siswa                   ||")
			fmt.Println("||  2. 👩‍🏫 Guru                    ||")
			fmt.Println("=====================================")
			fmt.Print("Tuliskan Opsi: ")
			fmt.Scan(&input)
			if input == "1" {
				loginsiswa(A, B, input, C, D, E, F)
			} else if input == "2" {
				loginguru(B, input, C, D, E, A, F)
			}
		case 2:
			fmt.Println("\n=====================================")
			fmt.Println("||       Salam Bahagia, Silahkan    ||")
			fmt.Println("||  Tuliskan Status Pengguna Dari   ||")
			fmt.Println("||         Pilihan Berikut.         ||")
			fmt.Println("=====================================")
			fmt.Println("||  1. 👨‍🎓 Siswa                   ||")
			fmt.Println("||  2. 👩‍🏫 Guru                    ||")
			fmt.Println("=====================================")
			fmt.Print("Tuliskan Opsi: ")
			fmt.Scan(&input)
			if input == "1" {
				registersiswa(A)
			} else if input == "2" {
				registerguru(B)
			}
		}
	}
}

func registersiswa(A *ArrayMahasiswa) {
	var user, pass string
	fmt.Println("\n****************************************")
	fmt.Println("*  Salam Bahagia, Silahkan Mohon Untuk *")
	fmt.Println("*   Menuliskan Nama dan Password       *")
	fmt.Println("*              Pengguna.               *")
	fmt.Println("****************************************")
	fmt.Print("📝 Tuliskan Nama: ")
	fmt.Scan(&user)
	fmt.Print("🔒 Tuliskan Password: ")
	fmt.Scan(&pass)
	A.TampM[A.n].Nama = user
	A.TampM[A.n].Password = pass
	A.n++
	InsertSortMhs(A)
	fmt.Println("***************************************")
	fmt.Println("*       Data Berhasil Disimpan!       *")
	fmt.Println("***************************************\n")
}

func loginsiswa(A *ArrayMahasiswa, G *ArrayGuru, input string, W *ArrayTugas, X *ArrayQuiz, Y *ArrayForum, Z *ArrayPerCourse) {
	var pilih int
	var hasil bool
	var nama, pass string
	fmt.Println("\n****************************************")
	fmt.Println("*             Silahkan Login!          *")
	fmt.Println("****************************************")
	fmt.Print("📝 Tuliskan Nama: ")
	fmt.Scan(&nama)
	fmt.Print("🔒 Tuliskan Password: ")
	fmt.Scan(&pass)
	hasil = binSearchLoginMhs(*A, nama, pass)
	if hasil == true {7
		course(W, X, Y, Z, G, A)
	} else {
		fmt.Println("\n****************************************")
		fmt.Println("* Data Tidak Ditemukan, Silahkan        *")
		fmt.Println("* Register atau Login Ulang.            *")
		fmt.Println("****************************************")
		fmt.Println("1. ✏️  Register.")
		fmt.Println("2. 🔄 Login.")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			registersiswa(A)
		case 2:
			loginsiswa(A, G, input, W, X, Y, Z)
		}
	}
}

func registerguru(B *ArrayGuru) {
	var user, pass string
	fmt.Println("🌟 Salam Bahagia, Silahkan Mohon Untuk Menuliskan Nama dan Password Pengguna. 🌟")
	fmt.Print("👤 Tuliskan Nama: ")
	fmt.Scan(&user)
	fmt.Print("🔒 Tuliskan Password: ")
	fmt.Scan(&pass)
	B.TampG[B.n].Nama = user
	B.TampG[B.n].Password = pass
	B.n++
	InserSortGr(B)
	fmt.Println("✅ Data Berhasil Disimpan!")
}

func loginguru(Z *ArrayGuru, input string, A *ArrayTugas, B *ArrayQuiz, C *ArrayForum, D *ArrayMahasiswa, E *ArrayPerCourse) {
	var pilih int
	var hasil bool
	var nama, pass string
	fmt.Println("========================================")
	fmt.Println("| 🔓 Silahkan Login!                  |")
	fmt.Println("========================================")
	fmt.Print("| 👤 Tuliskan Nama: ")
	fmt.Scan(&nama)
	fmt.Println("|")
	fmt.Print("| 🔑 Tuliskan Password: ")
	fmt.Scan(&pass)
	fmt.Println("|")
	hasil = binSearchLoginGr(*Z, nama, pass)
	if hasil == true {
		gurumenu(A, B, C, D, E)
	} else {
		fmt.Println("| ❌ Data Tidak Ditemukan             |")
		fmt.Println("| Silahkan Register atau Login Ulang. |")
		fmt.Println("========================================")
		fmt.Println("| 1️⃣. Register.                      |")
		fmt.Println("| 2️⃣. Login.                         |")
		fmt.Print("| 🔄 Pilih Opsi: ")
		fmt.Scan(&pilih)
		fmt.Println("|")
		switch pilih {
		case 1:
			registerguru(Z)
		case 2:
			loginguru(Z, input, A, B, C, D, E)
		}
	}
	fmt.Println("========================================")
}

func course(W *ArrayTugas, X *ArrayQuiz, Y *ArrayForum, B *ArrayPerCourse, G *ArrayGuru, M *ArrayMahasiswa) {
	var pilihan int
	for pilihan != 4 {
		fmt.Println("🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓🎓")
		fmt.Println("✨✨✨  Menu Course Siswa  ✨✨✨")
		fmt.Println("====================================")
		fmt.Println("| 📘 1️⃣. Tugas 📝                   |")
		fmt.Println("| 🧩 2️⃣. Quiz 🧠                    |")
		fmt.Println("| 💡 3️⃣. Forum 💬                   |")
		fmt.Println("| 🚪 4️⃣. Keluar 🔙                  |")
		fmt.Println("====================================")
		fmt.Print("| 🚀 Pilih opsi: ")
		fmt.Scan(&pilihan)
		fmt.Println("|")
		switch pilihan {
		case 1:
			tugas(W, B, M)
		case 2:
			quiz(X, B, M)
		case 3:
			forum(Y, B)
		}
		fmt.Println("💼💼💼💼💼💼💼💼💼💼💼💼💼💼💼💼")
	}
}

func tugas(A *ArrayTugas, B *ArrayPerCourse, C *ArrayMahasiswa) {
	var jawaban string
	var i, hasil int
	var loop bool = true

	for loop {
		fmt.Println("📚📚📚📚📚📚📚📚📚📚📚📚📚📚")
		fmt.Println("✏️✨ Tugas Time! Jawab Pertanyaan Berikut: ✨✏️")
		fmt.Println("====================================")
		fmt.Println("| 🔍 Pertanyaan:", A.TampT[i].Pertanyaan)
		fmt.Println("| 🖊️ Silahkan Berikan Jawaban: ")

		fmt.Scan(&jawaban)
		A.TampT[i].Jawaban = jawaban

		i++

		if i > A.n-1 {
			loop = false
			fmt.Println("====================================")
			fmt.Println("✅ Tugas Selesai! Good Job! ✅")
			fmt.Println("📚📚📚📚📚📚📚📚📚📚📚📚📚📚")
		}
	}

	C.TampM[i].DaftarNilai.TampC[i].NilaiTugas = hasil
}

func quiz(A *ArrayQuiz, B *ArrayPerCourse, C *ArrayMahasiswa) {
	var jawaban, nama string
	var i, hasil, cari, pilih int
	var loop bool = true

	fmt.Println("📝📝📝📝📝📝📝📝📝📝📝📝📝📝")
	fmt.Println("✨ Quiz Time! Ayo uji kemampuan anda! ✨")
	fmt.Println("========================================")
	fmt.Print("| 🤔 Tuliskan Nama: ")
	fmt.Scan(&nama)
	cari = seqSearchMhs(*C, nama)

	if cari != -1 {
		for loop {
			fmt.Println("| 🔍 Pertanyaan:", A.TampQ[i].Pertanyaan)
			fmt.Print("| 🖊️ Silahkan Berikan Jawaban: ")
			fmt.Scan(&jawaban)

			if jawaban == A.TampQ[i].Jawaban {
				hasil += 10
			}

			i++
			if i > A.n-1 {
				loop = false
				fmt.Println("========================================")
				fmt.Println("| ✅ Seluruh Soal Quiz Telah Dijawab! |")
				fmt.Println("========================================")
			}
		}
	}

	C.TampM[cari].DaftarNilai.TampC[cari].NilaiQuiz = hasil

	fmt.Println("| 🏆 Apakah Anda Ingin Mengetahui Nilai Quiz Anda?")
	fmt.Println("| 1️⃣. Iya")
	fmt.Println("| 2️⃣. Tidak")
	fmt.Print("| 🚀 Tuliskan Opsi: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		if hasil > 80 {
			fmt.Printf("| 🌟 Nilai Anda Adalah %d, Selalu Pertahankan Nilai Anda! 🌟\n", hasil)
		} else if hasil <= 80 && hasil > 70 {
			fmt.Printf("| 👍 Nilai Anda Adalah %d, Kerja Bagus, Tingkatkan Lagi! 👍\n", hasil)
		} else if hasil <= 70 && hasil > 60 {
			fmt.Printf("| 💪 Nilai Anda Adalah %d, Masih Ada Waktu Untuk Meningkatkan Lagi, Semangat! 💪\n", hasil)
		} else {
			fmt.Printf("| 🙏 Nilai Anda Adalah %d, Tingkatkan Lagi Performa Belajarmu, Jangan Patah Semangat! 🙏\n", hasil)
		}
	}

	fmt.Println("📝📝📝📝📝📝📝📝📝📝📝📝📝📝")
}

func forum(A *ArrayForum, B *ArrayPerCourse) {
	var jawaban, user string
	var cari, hasil int

	fmt.Println("📢📢📢📢📢📢📢📢📢📢📢📢📢📢")
	fmt.Println("✨ Forum Diskusi! ✨")
	fmt.Println("========================================")
	fmt.Println("| 🤔 Silahkan Pilih Pertanyaan Yang Ingin Dijawab |")
	fmt.Printf("| %s %10s %30s |\n", "No", "Pertanyaan", "Nama Pengguna")

	for i := 0; i < A.n; i++ {
		fmt.Printf("| %d. %-25s %-25s |\n", i+1, A.TampF[i].Pertanyaan, A.TampF[i].Pengirim)
	}

	fmt.Print("| 🚀 Pilih Nomor Pertanyaan: ")
	fmt.Scan(&cari)
	hasil = seqSearchForum(*A, cari-1)

	if hasil != -1 {
		fmt.Print("| 🖊️ Silahkan Masukkan Username: ")
		fmt.Scan(&user)
		A.TampF[hasil].Penjawab = user

		fmt.Println("| 🔍 Pertanyaan Terpilih:", A.TampF[hasil].Pertanyaan)
		fmt.Print("| 💡 Silahkan Berikan Jawaban: ")
		fmt.Scan(&jawaban)
		A.TampF[hasil].Jawaban = jawaban
	}

	fmt.Println("========================================")
	fmt.Println("✅ Terima Kasih Telah Berpartisipasi di Forum! ✅")
	fmt.Println("📢📢📢📢📢📢📢📢📢📢📢📢📢📢")
}

func gurumenu(A *ArrayTugas, B *ArrayQuiz, C *ArrayForum, D *ArrayMahasiswa, E *ArrayPerCourse) {
	var pilihan int

	fmt.Println("👩‍🏫👨‍🏫👩‍🏫👨‍🏫👩‍🏫👨‍🏫👩‍🏫👨‍🏫")
	fmt.Println("✨ Selamat Datang Di Menu Guru! ✨")
	fmt.Println("========================================")
	for pilihan != 11 {
		fmt.Println("| 1️⃣. Tambah Soal                |")
		fmt.Println("| 2️⃣. Edit Soal                  |")
		fmt.Println("| 3️⃣. Hapus Soal                 |")
		fmt.Println("| 4️⃣. Tambah Forum               |")
		fmt.Println("| 5️⃣. Edit Forum                 |")
		fmt.Println("| 6️⃣. Hapus Forum                |")
		fmt.Println("| 7️⃣. Tampilkan Data Soal        |")
		fmt.Println("| 8️⃣. Tampilkan Data Forum       |")
		fmt.Println("| 9️⃣. Tampilkan Data Mahasiswa   |")
		fmt.Println("| 🔟. Penilaian Mahasiswa        |")
		fmt.Println("| 🔚. Keluar                     |")
		fmt.Println(("Ketik Angka 11 Jika Ingin Keluar"))
		fmt.Println("========================================")
		fmt.Print("| 🚀 Pilih opsi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahSoal(A, B)
		case 2:
			editSoal(A, B)
		case 3:
			hapusSoal(A, B)
		case 4:
			tambahForum(C)
		case 5:
			editForum(C)
		case 6:
			hapusForum(C)
		case 7:
			dataSoal(B, A)
		case 8:
			dataForum(C)
		case 9:
			dataMahasiswa(D, E)
		case 10:
			penilaianMhs(A, D, E)
		case 11:
			fmt.Println("========================================")
			fmt.Println("| 🚪 Anda telah keluar dari menu guru. |")
			fmt.Println("========================================")
			return
		default:
			fmt.Println("| ❗ Pilihan tidak valid. Silakan coba lagi. |")
		}
	}
	fmt.Println("👩‍🏫👨‍🏫👩‍🏫👨‍🏫👩‍🏫👨‍🏫👩‍🏫👨‍🏫")
}

func tambahSoal(A *ArrayTugas, B *ArrayQuiz) {
	var t, p, j string // t adalah tipe soal(quiz atau tugas)
	var Lan int

	fmt.Println("📝📝📝📝📝📝📝📝📝📝📝📝📝📝")
	fmt.Println("✨ Time to Add New Questions! ✨")
	fmt.Println("========================================")
	fmt.Println("| 🧐 Tuliskan Tipe Soal Yang Diinginkan. |")
	fmt.Println("| 1️⃣. Quiz                             |")
	fmt.Println("| 2️⃣. Tugas                            |")
	fmt.Print("| 🚀 Tuliskan Opsi: ")
	fmt.Scan(&t)

	if t == "1" {
		for Lan != 2 {
			fmt.Scanln()
			fmt.Println("| 🖊️ Jika Ingin Menambahkan Pertanyaan dan Jawaban, Mohon Berikan '_' Untuk Setiap Spasi |")
			fmt.Print("| ❔ Silahkan Masukkan Pertanyaan: ")
			fmt.Scan(&p)
			fmt.Scanln()
			fmt.Print("| 💡 Silahkan Masukkan Jawaban: ")
			fmt.Scan(&j)

			B.TampQ[B.n].Pertanyaan = p
			B.TampQ[B.n].Jawaban = j
			B.n++

			fmt.Println("| ✅ Soal Telah Ditambahkan! |")
			fmt.Println("| ➕ Apakah Anda Ingin Melanjutkan Lagi? |")
			fmt.Println("| 1️⃣. Iya                                |")
			fmt.Println("| 2️⃣. Tidak                              |")
			fmt.Print("| 🚀 Pilih Opsi: ")
			fmt.Scan(&Lan)
		}
	} else if t == "2" {
		for Lan != 2 {
			fmt.Scanln()
			fmt.Println("| 🖊️ Jika Ingin Menambahkan Pertanyaan, Mohon Berikan '_' Untuk Setiap Spasi |")
			fmt.Print("| ❔ Silahkan Masukkan Pertanyaan: ")
			fmt.Scan(&p)

			A.TampT[A.n].Pertanyaan = p
			A.n++

			fmt.Println("| ✅ Soal Telah Ditambahkan! |")
			fmt.Println("| ➕ Apakah Anda Ingin Melanjutkan Lagi? |")
			fmt.Println("| 1️⃣. Iya                                |")
			fmt.Println("| 2️⃣. Tidak                              |")
			fmt.Print("| 🚀 Pilih Opsi: ")
			fmt.Scan(&Lan)
		}
	}

	fmt.Println("========================================")
	fmt.Println("✅ Terima Kasih Telah Menambahkan Soal! ✅")
	fmt.Println("📝📝📝📝📝📝📝📝📝📝📝📝📝📝")
}

func tambahForum(A *ArrayForum) {
	var user, p string
	var Lan int

	fmt.Println("🌟👥🌟👥🌟👥🌟👥🌟👥🌟")
	fmt.Println("Selamat Datang di Menu Penambahan Forum!")
	fmt.Println("========================================")
	fmt.Print("| 🙋 Masukkan Username Anda: ")
	fmt.Scan(&user)
	A.TampF[A.n].Pengirim = user

	for Lan != 2 {
		fmt.Print("| ❓ Masukkan Pertanyaan Anda: ")
		fmt.Scan(&p)
		A.TampF[A.n].Pertanyaan = p
		A.n++

		fmt.Println("| ✔️ Pertanyaan Berhasil Ditambahkan! |")
		fmt.Println("| ➡️ Apakah Anda Ingin Melanjutkan?   |")
		fmt.Println("| 1️⃣. Iya                             |")
		fmt.Println("| 2️⃣. Tidak                           |")
		fmt.Print("| 🚀 Pilih Opsi: ")
		fmt.Scan(&Lan)
	}

	for i := 0; i < A.n; i++ {
		A.TampF[i].Pengirim = user
	}

	fmt.Println("========================================")
	fmt.Println("✨ Terima Kasih Telah Berkontribusi ke Forum! ✨")
	fmt.Println("🌟👥🌟👥🌟👥🌟👥🌟👥🌟")
}

func editSoal(A *ArrayTugas, B *ArrayQuiz) {
	var t, p, j string
	var cari, hasil int

	fmt.Println("🌐📝🌐📝🌐📝🌐📝🌐📝🌐")
	fmt.Println("Selamat Datang di Antarmuka Edit Pertanyaan!")
	fmt.Println("========================================")
	fmt.Println("| Silakan Pilih Tipe Soal yang Akan Diubah: |")
	fmt.Println("| 1. Kuis                                   |")
	fmt.Println("| 2. Tugas                                  |")
	fmt.Print("| 🔄 Pilih Opsi: ")
	fmt.Scan(&t)
	fmt.Println("| Silakan Pilih Soal yang Akan Diubah:      |")

	if t == "2" {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < A.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, A.TampT[i].Pertanyaan, A.TampT[i].Jawaban)
		}
	} else if t == "1" {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < B.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, B.TampQ[i].Pertanyaan, B.TampQ[i].Jawaban)
		}
	}
	fmt.Print("| 🔍 Masukkan Nomor Soal: ")
	fmt.Scan(&cari)

	if t == "2" {
		hasil = seqSearchTugas(*A, cari-1)
	} else if t == "1" {
		hasil = seqSearchQuiz(*B, cari-1)
	}

	if hasil != -1 {
		if t == "2" {
			fmt.Println("| Jika Ingin Menambahkan Pertanyaan dan Jawaban, Gunakan '_' untuk Setiap Spasi |")
			fmt.Println("| Silakan Masukkan Pertanyaan Baru.                                             |")
			fmt.Print("| ➕ Masukkan Pertanyaan Baru: ")
			fmt.Scan(&p)
			A.TampT[hasil].Pertanyaan = p
			fmt.Println("| ✔️ Pertanyaan Tugas Berhasil Diperbarui! |")
		} else if t == "1" {
			fmt.Println("| Jika Ingin Menambahkan Pertanyaan dan Jawaban, Gunakan '_' untuk Setiap Spasi |")
			fmt.Println("| Silakan Masukkan Pertanyaan dan Jawaban Baru.                             |")
			fmt.Print("| ➕ Masukkan Pertanyaan Baru: ")
			fmt.Scan(&p)
			fmt.Scanln()
			fmt.Print("| ➕ Masukkan Jawaban Baru: ")
			fmt.Scan(&j)
			B.TampQ[hasil].Pertanyaan = p
			B.TampQ[hasil].Jawaban = j
			fmt.Println("| ✔️ Pertanyaan Kuis Berhasil Diperbarui!   |")
		}
	} else {
		fmt.Println("⚠️ Soal Tidak Ditemukan. Silakan Coba Lagi.")
	}

	fmt.Println("========================================")
	fmt.Println("✨ Soal Berhasil Diperbarui! ✨")
	fmt.Println("🌐📝🌐📝🌐📝🌐📝🌐📝🌐")
}

func editForum(A *ArrayForum) {
	var user, p string
	var cari, hasil int

	fmt.Println("🌟📬🌟📬🌟📬🌟📬🌟📬🌟")
	fmt.Println("Selamat Datang di Antarmuka Edit Forum!")
	fmt.Println("=======================================")
	fmt.Println("| Silakan Masukkan Nama Pengguna Anda.  |")
	fmt.Print("| 🙍 Nama Pengguna: ")
	fmt.Scan(&user)
	fmt.Println("| Silakan Pilih Forum yang Akan Diubah. |")
	fmt.Printf("%s %10s\n", "No", "Pertanyaan")

	for i := 0; i < A.n; i++ {
		fmt.Printf("%d %8s\n", i+1, A.TampF[i].Pertanyaan)
	}

	fmt.Print("| 🔍 Masukkan Nomor Pertanyaan: ")
	fmt.Scan(&cari)

	hasil = seqSearchForum(*A, cari-1)

	if hasil != -1 {
		fmt.Println("| Silakan Masukkan Pertanyaan Baru.      |")
		fmt.Print("| ✏️ Tuliskan Pertanyaan: ")
		fmt.Scan(&p)
		A.TampF[hasil].Pertanyaan = p
		fmt.Println("| ✔️ Pertanyaan Forum Berhasil Diperbarui! |")
	} else {
		fmt.Println("⚠️ Pertanyaan Tidak Ditemukan. Silakan Coba Lagi.")
	}

	fmt.Println("=======================================")
	fmt.Println("✨ Terima Kasih Telah Memperbarui Forum! ✨")
	fmt.Println("🌟📬🌟📬🌟📬🌟📬🌟📬🌟")
}

func hapusSoal(A *ArrayTugas, B *ArrayQuiz) {
	var t string
	var hasil, cari, i int

	fmt.Println("🗑️📝🗑️📝🗑️📝🗑️📝🗑️📝")
	fmt.Println("Selamat Datang di Antarmuka Hapus Soal!")
	fmt.Println("======================================")
	fmt.Println("| Silakan Pilih Tipe Soal yang Akan Dihapus: |")
	fmt.Println("| 1. Kuis                                    |")
	fmt.Println("| 2. Tugas                                   |")
	fmt.Print("| 🔄 Pilih Opsi: ")
	fmt.Scan(&t)
	fmt.Println("| Silakan Pilih Soal yang Akan Dihapus:      |")

	if t == "2" {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < A.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, A.TampT[i].Pertanyaan, A.TampT[i].Jawaban)
		}
		fmt.Print("| 🗑️ Masukkan Nomor Tugas yang Ingin Dihapus: ")
		fmt.Scan(&cari)
	} else if t == "1" {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < B.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, B.TampQ[i].Pertanyaan, B.TampQ[i].Jawaban)
		}
		fmt.Print("| 🗑️ Masukkan Nomor Kuis yang Ingin Dihapus: ")
		fmt.Scan(&cari)
	}

	if t == "2" {
		hasil = seqSearchTugas(*A, cari-1)
	} else if t == "1" {
		hasil = seqSearchQuiz(*B, cari-1)
	}

	if hasil != -1 {
		if t == "2" {
			for i = hasil; i < A.n-1; i++ {
				A.TampT[i].Pertanyaan = A.TampT[i+1].Pertanyaan
				A.TampT[i].Jawaban = A.TampT[i+1].Jawaban
			}
			A.n--
			fmt.Println("| ✔️ Tugas Berhasil Dihapus!                |")
		} else if t == "1" {
			for i = hasil; i < B.n-1; i++ {
				B.TampQ[i].Pertanyaan = B.TampQ[i+1].Pertanyaan
				B.TampQ[i].Jawaban = B.TampQ[i+1].Jawaban
			}
			B.n--
			fmt.Println("| ✔️ Kuis Berhasil Dihapus!                 |")
		}
	} else {
		fmt.Println("⚠️ Soal Tidak Ditemukan. Silakan Coba Lagi.")
	}

	fmt.Println("======================================")
	fmt.Println("✨ Terima Kasih Telah Menghapus Soal! ✨")
	fmt.Println("🗑️📝🗑️📝🗑️📝🗑️📝🗑️📝")
}

func hapusForum(A *ArrayForum) {
	var user string
	var cari, hasil, i int

	fmt.Println("🗑️📜🗑️📜🗑️📜🗑️📜🗑️📜")
	fmt.Println("Selamat Datang di Tampilan Hapus Forum!")
	fmt.Println("=======================================")
	fmt.Println("| Silakan Masukkan Nama Pengguna Anda.  |")
	fmt.Print("| 🙍 Nama Pengguna: ")
	fmt.Scan(&user)
	fmt.Println("| Silakan Pilih Pertanyaan yang Akan Dihapus: |")
	fmt.Printf("%s %10s\n", "No", "Pertanyaan")
	for i := 0; i < A.n; i++ {
		fmt.Printf("%d %8s\n", i+1, A.TampF[i].Pertanyaan)
	}

	fmt.Print("| 🗑️ Masukkan Nomor Pertanyaan: ")
	fmt.Scan(&cari)

	hasil = seqSearchForum(*A, cari-1)

	if hasil != -1 {
		for i = hasil; i < A.n-1; i++ {
			A.TampF[i].Pertanyaan = A.TampF[i+1].Pertanyaan
		}
		A.n--
		fmt.Println("| ✔️ Pertanyaan Forum Berhasil Dihapus! |")
	} else {
		fmt.Println("⚠️ Pertanyaan Tidak Ditemukan. Silakan Coba Lagi.")
	}

	fmt.Println("=======================================")
	fmt.Println("✨ Terima Kasih Telah Menghapus Pertanyaan! ✨")
	fmt.Println("🗑️📜🗑️📜🗑️📜🗑️📜🗑️📜")
}

func dataSoal(A *ArrayQuiz, B *ArrayTugas) {
	var opsi int

	fmt.Println("📊📚📊📚📊📚📊📚📊📚")
	fmt.Println("Selamat Datang di Tampilan Data Soal!")
	fmt.Println("=====================================")
	fmt.Println("| Silakan Pilih Tipe Soal yang Akan Ditampilkan: |")
	fmt.Println("| 1. Kuis                                        |")
	fmt.Println("| 2. Tugas                                       |")
	fmt.Print("| 🔄 Pilih Opsi: ")
	fmt.Scan(&opsi)

	if opsi == 2 {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < B.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, B.TampT[i].Pertanyaan, B.TampT[i].Jawaban)
		}
	} else if opsi == 1 {
		fmt.Printf("%s %10s %20s\n", "No", "Pertanyaan", "Jawaban")
		for i := 0; i < A.n; i++ {
			fmt.Printf("%d %8s %18s\n", i+1, A.TampQ[i].Pertanyaan, A.TampQ[i].Jawaban)
		}
	}

	fmt.Println("=====================================")
	fmt.Println("✨ Data Soal Telah Ditampilkan! ✨")
	fmt.Println("📊📚📊📚📊📚📊📚📊📚")
}

func dataForum(A *ArrayForum) {
	fmt.Printf("%s %10s %30s %20s %30s\n", "No", "Pertanyaan", "Nama Pengirim", "Jawaban", "Nama Penjawab")
	for i := 0; i < A.n; i++ {
		fmt.Printf("%d %9s %24s %25s %29s\n", i+1, A.TampF[i].Pertanyaan, A.TampF[i].Pengirim, A.TampF[i].Jawaban, A.TampF[i].Penjawab)
	}
}

func dataMahasiswa(C *ArrayMahasiswa, B *ArrayPerCourse) {
	var opsi int

	fmt.Println("🎓📈🎓📈🎓📈🎓📈🎓📈")
	fmt.Println("Selamat Datang di Tampilan Data Mahasiswa!")
	fmt.Println("==========================================")
	fmt.Println("| Silakan Pilih Tipe Soal yang Ingin Ditampilkan Data Mahasiswanya: |")
	fmt.Println("| 1. Kuis                                                           |")
	fmt.Println("| 2. Tugas                                                          |")
	fmt.Print("| 🔄 Pilih Opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		fmt.Printf("| %s %10s %20s |\n", "Ranking", "Nama", "Nilai Kuis")
		for i := 0; i < C.n; i++ {
			fmt.Printf("| %d %8s %18d |\n", i+1, C.TampM[i].Nama, C.TampM[i].DaftarNilai.TampC[i].NilaiQuiz)
		}
	} else if opsi == 2 {
		fmt.Printf("| %s %10s %20s |\n", "Ranking", "Nama", "Nilai Tugas")
		for i := 0; i < C.n; i++ {
			fmt.Printf("| %d %8s %18d |\n", i+1, C.TampM[i].Nama, C.TampM[i].DaftarNilai.TampC[i].NilaiTugas)
		}
	}

	fmt.Println("==========================================")
	fmt.Println("✨ Data Mahasiswa Telah Ditampilkan! ✨")
	fmt.Println("🎓📈🎓📈🎓📈🎓📈🎓📈")
}

func penilaianMhs(A *ArrayTugas, B *ArrayMahasiswa, C *ArrayPerCourse) {
	var loop bool = true
	var nama string
	var hasil, i, nilai, cari int

	fmt.Println("‍ Penilaian Tugas Mahasiswa ‍")
	fmt.Println("=================================")
	fmt.Printf("%-10s | %15s | %10s\n", "No", "Pertanyaan", "Nilai")
	fmt.Println("---------------------------------")

	fmt.Println("Silahkan Pilih Nama Mahasiswa Yang Ingin Diberikan Nilai Tugas.")
	for i := 0; i < B.n; i++ {
		fmt.Printf("%d %8s\n", i+1, B.TampM[i].Nama)
	}
	fmt.Println("Masukkan Nama: ")
	fmt.Scan(&nama)

	cari = seqSearchMhs(*B, nama)
	if cari != -1 {
		fmt.Println("Silahkan Beri Nilai Untuk Setiap Jawaban Yang Diberikan.")
		for loop {
			fmt.Printf("%-10d | %15s | ", i+1, A.TampT[i].Pertanyaan)
			if len(A.TampT[i].Jawaban) <= 20 {
				fmt.Printf("%s\n", A.TampT[i].Jawaban)
			} else {
				fmt.Printf("%s...\n", A.TampT[i].Jawaban[:20])
				fmt.Printf("      (Tekan enter untuk melihat jawaban lengkap)\n")
				fmt.Scanln()
				fmt.Printf("%s\n", A.TampT[i].Jawaban)
			}
			fmt.Println("---------------------------------")
			fmt.Print("Masukkan Nilai: ")
			fmt.Scan(&nilai)
			hasil += nilai
			i++
			if i > A.n-1 {
				loop = false
			}
		}
	}

	if loop {
		fmt.Println(" Penilaian selesai! ")
		fmt.Printf("Nilai total untuk %s adalah: %d\n", B.TampM[cari].Nama, hasil)
	} else {
		fmt.Println(" Penilaian dibatalkan. ")
	}

	B.TampM[cari].DaftarNilai.TampC[cari].NilaiTugas = hasil
}

func seqSearchMhs(A ArrayMahasiswa, x string) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < A.n {
		if A.TampM[k].Nama == x {
			found = k
		} else {
			k++
		}
	}
	return found
}

func seqSearchTugas(A ArrayTugas, x int) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < A.n {
		if k == x {
			found = k
		} else {
			k++
		}
	}
	return found
}

func seqSearchQuiz(A ArrayQuiz, x int) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < A.n {
		if k == x {
			found = k
		} else {
			k++
		}
	}
	return found
}

func seqSearchForum(A ArrayForum, x int) int {
	var k int
	k = 0
	found := -1
	for found == -1 && k < A.n {
		if k == x {
			found = k
		} else {
			k++
		}
	}
	return found
}

func binSearchLoginMhs(A ArrayMahasiswa, x string, y string) bool {
	low := 0
	high := A.n - 1
	found := false

	for low <= high && found != true {
		mid := (low + high) / 2
		if A.TampM[mid].Nama < x || (A.TampM[mid].Nama == x && A.TampM[mid].Password < y) {
			low = mid + 1
		} else if A.TampM[mid].Nama > x || (A.TampM[mid].Nama == x && A.TampM[mid].Password > y) {
			high = mid - 1
		} else {
			found = true
		}
	}
	return found
}

func binSearchLoginGr(A ArrayGuru, x, y string) bool {
	low := 0
	high := A.n - 1
	found := false

	for low <= high && found != true {
		mid := (low + high) / 2
		if A.TampG[mid].Nama < x || (A.TampG[mid].Nama == x && A.TampG[mid].Password < y) {
			low = mid + 1
		} else if A.TampG[mid].Nama > x || (A.TampG[mid].Nama == x && A.TampG[mid].Password > y) {
			high = mid - 1
		} else {
			found = true
		}
	}
	return found
}

func SelectSortMhs(A *ArrayMahasiswa) {
	var i, j, temp int
	var t Mahasiswa
	i = 1
	for i <= A.n-1 {
		temp = i - 1
		j = i
		for j < A.n {
			if A.TampM[temp].Nama > A.TampM[j].Nama {
				temp = j
			}
			j++
		}
		t.Nama = A.TampM[temp].Nama
		A.TampM[temp].Nama = A.TampM[i-1].Nama
		A.TampM[i-1].Nama = t.Nama
		i++
	}
}

func InsertSortMhs(A *ArrayMahasiswa) {
	var i, j int
	var t Mahasiswa
	i = 1
	for i <= A.n-1 {
		j = i
		t = A.TampM[j]
		for j > 0 && t.Nama < A.TampM[j-1].Nama {
			A.TampM[j] = A.TampM[j-1]
			j--
		}
		A.TampM[j] = t
		i++
	}
}

func InserSortGr(A *ArrayGuru) {
	var i, j int
	var t Guru
	i = 1
	for i <= A.n-1 {
		j = i
		t = A.TampG[j]
		for j > 0 && t.Nama < A.TampG[j-1].Nama {
			A.TampG[j] = A.TampG[j-1]
			j--
		}
		A.TampG[j] = t
		i++
	}
}
