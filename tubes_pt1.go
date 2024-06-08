package main

//import eketensi yang diperlukan
import (
	"fmt"
	"time"
	// "os"
)

// Konstanta minimal dari data
const NMAX int = 1000

var sesilogin int = -1

var dbc int
var currentTime = time.Now()

// Format untuk data nama lengkap
type namalengkap struct {
	namadepan, namabelakang string
}

// Format untuk data data lengkap
type formattanggal struct {
	tanggal, bulan, tahun int
}

// struct yang isinya hasil2 untuk BMI dan TD
type diagnosa struct {
	hasil    float64
	kelompok string
	kategori string
}

// Atribut yang mencatat data pasien
type datapasien struct {
	idpasien      int
	status        int // 1 = aktif ; 2= tidak aktif
	namapasien    namalengkap
	usia          int
	tempatlahir   string
	tanggallahir  formattanggal
	gender        string
	password      string
	datakesehatan arraykunjungan
}

// Atribut kunjungan untuk per user
type kunjungan struct {
	statusKunjungan         int
	idkunjungan             int
	iduser                  int
	tanggalpengecekan       formattanggal
	tekanandarahdistolik    float64
	tekanandarahsistolik    float64
	beratbadan, tinggibadan float64
	diagnosabmi             diagnosa
	diagnosatekanandarah    diagnosa
}

// Atribut banyak data pasien sesuai banyak limit yang ditentukan di NMAX
type atributpasien [NMAX]datapasien
type arraykunjungan [NMAX]kunjungan

// Interface yang akan di tampilkan
func interface0(P *atributpasien, sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("                SELAMAT DATANG DI               ")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang, "Dengan ID User :", P[*&sesilogin].idpasien)
	}
	fmt.Println("Silahkan layanan yang anda ingin gunakan: ")
	fmt.Println("1. Edit Profil Pasien")
	fmt.Println("2. Menu Data Kunjungan Pasien")
	fmt.Println("3. Cek Kesehatan")
	fmt.Println("4. Keluar dan Logout")
	if sesilogin == 999 {
		fmt.Println("5. Fitur Admin")
	}
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka: ")
}

func interface01(P *atributpasien, sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang, "Dengan ID User :", P[*&sesilogin].idpasien)
	}
	fmt.Println("Pilih layanan akun yang dinginkan?")
	fmt.Println("1. Daftar (Jika Anda belum pernah daftar)")
	fmt.Println("2. Masuk (Jika Anda sudah pernah daftar sebelumnya")
	fmt.Println("3. Ubah Data Diri")
	fmt.Println("4. Ubah Kata Sandi / Password")
	fmt.Println("5. Hapus data diri / akun")
	fmt.Println("6. Keluar / Logout")
	fmt.Println("7. Kembali ke halaman utama")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka: ")
}

func interface02(P *atributpasien, sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang, "Dengan ID User :", P[*&sesilogin].idpasien)
	}
	fmt.Println("Pilih layanan data yang mana kamu inginkan?")
	fmt.Println("1. Lihat Semua Data")
	fmt.Println("2. Rekam/Tambah Data")
	fmt.Println("3. Ubah/Edit Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Kembali ke halaman utama")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka: ")
}

func interface03(P *atributpasien, sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang)
	}
	fmt.Println("Pilih layanan data yang mana kamu inginkan?")
	fmt.Println("1. Cek Kesehatan")
	fmt.Println("2. Kembali ke halaman utama")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka: ")
}


func interface05(sesilogin int) {
	fmt.Println("=====================================================")
	fmt.Println("               DATABASE PASIEN KESEHATAN         ")
	fmt.Println("            RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("=====================================================")
	fmt.Println("<1> Tampilkan Data terurut") // sort descending (selection sort)
	fmt.Println("<2> Check ID Pasien")        //sort ascending (insertion sort)
	fmt.Println("<3> Edit Database Pasien")
	fmt.Println("=====================================================")
	fmt.Print("> ")
}

// Program utama
func main() {
	var datapasien atributpasien
	var pilihan int
	var sesi int
	var dbk int

	datapasien[999].namapasien.namadepan = "admin"
	datapasien[999].idpasien = 999
	datapasien[999].password = "Admin!"

	sesi = -1
	pilihan = 1
	for pilihan >= 1 && pilihan <= 7 {
		interface0(&datapasien, sesi)
		fmt.Println("dbc :", dbc)
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for pilihan >= 1 && pilihan <= 6 {
				interface01(&datapasien, sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					daftarpengguna(&datapasien, &dbc, &sesi)
				} else if pilihan == 2 {
					masukpengguna(&datapasien, &dbc, &sesi)
				} else if pilihan == 3 {
					ubahdatapasien(&datapasien, &dbc, &sesi)
				} else if pilihan == 4 {
					ubahpassword(&datapasien, &dbc, &sesi)
				} else if pilihan == 5 {
					hapusdatapasien(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 6 {
					keluar(&datapasien, &dbc, &sesi)
				}
			}
		} else if pilihan == 2 {
			for pilihan >= 1 && pilihan <= 4 {
				interface02(&datapasien, sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					lihatdata(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 2 {
					tambahdata(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 3 {
					// editdata(&datapasien, &dbc, &sesi)
				} else if pilihan == 4 {
					// hapusdata(&datapasien, &dbc, &sesi)
				}
			}
		} else if pilihan == 3 {
			interface03(&datapasien, sesi)
			for pilihan == 1 {
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					cekkesehatan(&datapasien, &dbc, &dbk, &sesi)
				}
			}
		} else if pilihan == 4 {
			//Keluar()
		} else if pilihan == 5 {
			var uruti int
			pilihan = 1
			for pilihan >= 1 && pilihan <= 4 {
				interface05(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Println()
					fmt.Println("<1>. Tampilkan berdasarkan Data Terkini")
					fmt.Println("<2>. Tampilkan berdasarkan Tahun Lahir Pasien")
					fmt.Println()
					fmt.Println("Masukkan Angka")
					fmt.Scan(&uruti)
					for uruti != 1 && uruti != 2 {
						fmt.Scan(&uruti)
					}
					if uruti == 1 {
						sortSelectionDesc(&datapasien, &dbc)

					} else if uruti == 2 {
						sortInsertionAsx(&datapasien, &dbc)

					}
				} else if pilihan == 2 {
					cariAkunPasien(&datapasien, &dbc, &sesi)
				} else if pilihan == 3 {
					EditDatabasePasien(&datapasien, &dbc)
				} else if pilihan == 4 {
					sesi = -1
				}
			}
		}
	}
}

// Program per-layanan

func daftarpengguna(P *atributpasien, dbc *int, sesi *int) {
	var chunk_gender string
	var chunk_namadepan, chunk_namabelakang, chunk_tempatlahir string
	var chunk_tangallahir, chunk_bulanlahir, chunk_tahunlahir int
	var i int

	idxstatuskosong := pencaristatuskosong(*P, *dbc)
	if idxstatuskosong == -1 {
		fmt.Println("Kamu daftar dengan ID yang baru!")
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tangallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for i < *dbc-1 {
			if chunk_namadepan == P[i].namapasien.namadepan && chunk_namabelakang == P[i].namapasien.namabelakang && chunk_tempatlahir == P[i].tempatlahir && chunk_tangallahir == P[i].tanggallahir.tanggal && chunk_bulanlahir == P[i].tanggallahir.tanggal && chunk_tahunlahir == P[i].tanggallahir.tahun {
				fmt.Println("Maaf pengguna sudah terdaftar silahkan masuk ke menu masuk, dengan ID yang sudah terdaftar")
				break
			}
			i = i + 1
		}
		*dbc = *dbc + 1
		P[*dbc].namapasien.namadepan = chunk_namadepan
		P[*dbc].namapasien.namabelakang = chunk_namabelakang
		P[*dbc].tempatlahir = chunk_tempatlahir
		P[*dbc].tanggallahir.tanggal = chunk_tangallahir
		P[*dbc].tanggallahir.bulan = chunk_bulanlahir
		P[*dbc].tanggallahir.tahun = chunk_tahunlahir
		fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
		fmt.Println("L / l = Laki-laki")
		fmt.Println("P / p = Perempuan")
		fmt.Scan(&chunk_gender)
		if chunk_gender == "L" || chunk_gender == "l" {
			P[*dbc].gender = "Laki-laki"
		} else if chunk_gender == "P" || chunk_gender == "p" {
			P[*dbc].gender = "Perempuan"
		}
		fmt.Println("Masukkan Password untuk user yang di daftarkan")
		fmt.Scan(&P[*dbc].password)

		P[*dbc].status = 1
		P[*dbc].idpasien = *dbc
		*sesi = *dbc
		fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
		fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
	} else {
		fmt.Print(*dbc)
		fmt.Println("Kamu daftar dengan ID yang lama!")
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tangallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for i <= *dbc {
			if chunk_namadepan == P[i].namapasien.namadepan && chunk_namabelakang == P[i].namapasien.namabelakang && chunk_tempatlahir == P[i].tempatlahir && chunk_tangallahir == P[i].tanggallahir.tanggal && chunk_bulanlahir == P[i].tanggallahir.tanggal && chunk_tahunlahir == P[i].tanggallahir.tahun {
				fmt.Println("Maaf pengguna sudah terdaftar silahkan masuk ke menu masuk, dengan ID yang sudah terdaftar")
				break
			}
			i = i + 1
		}
		P[idxstatuskosong].namapasien.namadepan = chunk_namadepan
		P[idxstatuskosong].namapasien.namabelakang = chunk_namabelakang
		P[idxstatuskosong].tempatlahir = chunk_tempatlahir
		P[idxstatuskosong].tanggallahir.tanggal = chunk_tangallahir
		P[idxstatuskosong].tanggallahir.bulan = chunk_bulanlahir
		P[idxstatuskosong].tanggallahir.tahun = chunk_tahunlahir
		fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
		fmt.Println("L / l = Laki-laki")
		fmt.Println("P / p = Perempuan")
		fmt.Scan(&chunk_gender)
		if chunk_gender == "L" || chunk_gender == "l" {
			P[idxstatuskosong].gender = "Laki-laki"
		} else if chunk_gender == "P" || chunk_gender == "p" {
			P[idxstatuskosong].gender = "Perempuan"
		}
		fmt.Println("Masukkan Password untuk user yang di daftarkan")
		fmt.Scan(&P[idxstatuskosong].password)

		P[idxstatuskosong].status = 1
		P[idxstatuskosong].idpasien = idxstatuskosong
		*sesi = idxstatuskosong
		fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
		fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
	}
}

func pencaristatuskosong(P atributpasien, dbc int) int {
	var i int
	i = 0
		for i < dbc {
			if P[i].status == 1{
				return i
			}
			i = i + 1
		}
		return -1
}

func masukpengguna(P *atributpasien, dbc *int, sesi *int) {
	var chunk_id int
	var chunk_password string

	fmt.Println("Silahkan masukkan ID yang sudah terdaftar")
	fmt.Scan(&chunk_id)
	fmt.Println("Silahkan masukkan password sesuai ID yang sudah dimasukkan")
	fmt.Scan(&chunk_password)
	if P[chunk_id].password == chunk_password {
		*sesi = -1
		*sesi = chunk_id
		fmt.Println("Login Berhasil")
		fmt.Println("Selamat datang", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang)
	} else {
		fmt.Println("ID atau Password tidak ditemukan atau tidak sesuai, harap ulangi lagi, atau hubungi petugas jika terddapat kendala")
	}
}

func ubahdatapasien(P *atributpasien, dbc *int, sesi *int) {
	var chunk_pernyataan int
	var chunk_gender string
	var chunk_password, chunk_password2 string

	chunk_pernyataan = 1
	fmt.Println("Apakah data pasien yang akan di rubah adalah ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Println("Masukkan angka 1 / 0 (1 untuk BENAR, 0 untuk SALAH)")
	fmt.Scan(&chunk_pernyataan)
	fmt.Println("Masukkan kata sandi / password :")
	fmt.Scan(&chunk_password)
	fmt.Println("Masukkan kembalikata sandi / password :")
	fmt.Scan(&chunk_password2)
	if *sesi != -1 && chunk_pernyataan == 1 && chunk_password == P[*sesi].password && chunk_password2 == P[*sesi].password && chunk_password == chunk_password2 {
		fmt.Println("Masukkan Nama Depan yang ingin di ubah")
		fmt.Scan(&P[*sesi].namapasien.namadepan)
		fmt.Println("Masukkan Nama Belakang yang ingin di ubah")
		fmt.Scan(&P[*sesi].namapasien.namabelakang)
		fmt.Println("Masukkan Tempat Lahir yang ingin di ubah")
		fmt.Scan(&P[*sesi].tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir yang ingin di ubah (DD/ M/ YYYY)")
		fmt.Scan(&P[*sesi].tanggallahir.tanggal, &P[*sesi].tanggallahir.bulan, &P[*sesi].tanggallahir.tahun)
		fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
		fmt.Println("L / l = Laki-laki")
		fmt.Println("P / p = Perempuan")
		fmt.Scan(&chunk_gender)
		if chunk_gender == "L" || chunk_gender == "l" {
			P[*dbc].gender = "Laki-laki"
		} else if chunk_gender == "P" || chunk_gender == "p" {
			P[*dbc].gender = "Perempuan"
		}
		fmt.Println("Data sudah terubah, dengan keterangan :")
		fmt.Println("Nama Depan           	:", P[*sesi].namapasien.namadepan)
		fmt.Println("Nama Belakang        	:", P[*sesi].namapasien.namabelakang)
		fmt.Println("Tempat Lahir         	:", P[*sesi].tempatlahir)
		fmt.Println("Tanggal/ Tahun Lahir	:", P[*sesi].tanggallahir.tanggal, P[*sesi].tanggallahir.bulan, P[*sesi].tanggallahir.tahun)
		fmt.Println("Gender					:", P[*sesi].gender)
	} else {
		if *sesi == -1 {
			fmt.Println("Pastikan Anda masuk dengan ID atau akun yang benar, jika ada kendala hubungi petugas")
		} else if P[*sesi].password != chunk_password {
			fmt.Println("Password / kata sandi yang anda masukkan salah, silahkan coba lagi")
		}
	}
}

func ubahpassword(P *atributpasien, dbc *int, sesi *int) {
	var chunk_pernyataan int
	var chunk_password, chunk_password2 string
	var chunk_passwordbaru, chunk_passwordbaru2 string

	chunk_pernyataan = 1
	fmt.Println("Apakah data pasien yang akan di rubah adalah ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Println("Masukkan angka 1 / 0 (1 untuk BENAR, 0 untuk SALAH)")
	fmt.Scan(&chunk_pernyataan)
	fmt.Println("Masukkan kata sandi / password :")
	fmt.Scan(&chunk_password)
	fmt.Println("Masukkan kembali kata sandi / password :")
	fmt.Scan(&chunk_password2)
	if *sesi != -1 && chunk_pernyataan == 1 && chunk_password == P[*sesi].password && chunk_password2 == P[*sesi].password && chunk_password == chunk_password2 {
		fmt.Println("Masukkan password baru :")
		fmt.Scan(&chunk_passwordbaru)
		fmt.Println("Masukkan kembali password baru :")
		fmt.Scan(&chunk_passwordbaru2)
		if chunk_passwordbaru == chunk_passwordbaru2 {
			P[*sesi].password = chunk_passwordbaru
			fmt.Println("Password sudah di ubah")
		}
	} else {
		if *sesi == -1 {
			fmt.Println("Pastikan Anda masuk dengan ID atau akun yang benar, jika ada kendala hubungi petugas")
		} else if P[*sesi].password != chunk_password {
			fmt.Println("Password / kata sandi yang anda masukkan salah, silahkan coba lagi")
		}
	}
}

func hapusdatapasien(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var pernyataanpenghapusan int
	var chunk_password, chunk_password2 = "0", "0"

	fmt.Println("================ PERINGATAN ================")
	fmt.Println(" Data Kunjungan dan akun anda akan di hapus ")
	fmt.Println("  dan jika Anda ingin menggunakan aplikasi  ")
	fmt.Println(" Anda harus mendaftar ulang di menu daftar  ")
	fmt.Println("============================================")
	fmt.Println("Akun yang akan dihapus adalah akun dengan ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Println("Apakah anda ingin melanjutkan untuk penghapusan akun ? (Masukkan 1 jika anda SETUJU, masukan 0 jika anda BATAL)")
	fmt.Scan(&pernyataanpenghapusan)
	fmt.Println("Masukkan kata sandi / password :")
	fmt.Scan(&chunk_password)
	fmt.Println("Masukkan kembali kata sandi / password :")
	fmt.Scan(&chunk_password2)
	if pernyataanpenghapusan == 1 && *sesi != -1 && chunk_password == P[*sesi].password && chunk_password2 == P[*sesi].password {
		P[*sesi].status = 2
		fmt.Println("Data akun anda dan kunjungan sudah di hapus")
		*sesi = -1
	} else {
		fmt.Println("|Kamu membatalkan penghapusan akun dengan tidak  |")
		fmt.Println("| menyetujui pernyataan penghapusan ATAU kamu    |")
		fmt.Println("|    Terjadi kesalahan pemasukkan password       |")
	}
}

func keluar(P *atributpasien, dbc *int, sesi *int) {
	*sesi = -1
	fmt.Println("Anda sudah keluar, terimakasih sudah menggunakan aplikasi kami")
}

func lihatdata(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var chunk_tanggal, chunk_bulan, chunk_tahun int
	var chunk_ketemu int

	fmt.Println("Pastikan Anda login dengan ID yang sesuai")
	fmt.Println("Masukkan tanggal pengecekan kesehatan (DD/M/YYYY) : ")
	fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
	chunk_ketemu = pencaritanggal(*P, *dbc, *dbk, *sesi, chunk_tanggal, chunk_bulan, chunk_tahun)
	if chunk_ketemu != 1 {
		fmt.Println("ID Kunjungan : ", P[*sesi].datakesehatan[chunk_ketemu].idkunjungan)
		fmt.Println(P[*sesi].datakesehatan[chunk_ketemu].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[chunk_ketemu].tanggalpengecekan.bulan, P[*sesi].datakesehatan[chunk_ketemu].tanggalpengecekan.tahun)
		fmt.Println("Berat Badan               : ", P[*sesi].datakesehatan[chunk_ketemu].beratbadan)
		fmt.Println("Tinggi Badan              : ", P[*sesi].datakesehatan[chunk_ketemu].tinggibadan)
		fmt.Println("BMI                       : ", P[*sesi].datakesehatan[chunk_ketemu].diagnosabmi)
		fmt.Println("Tekanan Darah Distoltik   : ", P[*sesi].datakesehatan[chunk_ketemu].tekanandarahdistolik)
		fmt.Println("Tekanan Darah Sistolik    : ", P[*sesi].datakesehatan[chunk_ketemu].tekanandarahsistolik)
	} else {
		fmt.Println("Berdasarkan waktu yang anda masukkan, data tidak ditemukan")
	}
}

func pencaritanggal(P atributpasien, dbc, dbk, sesi, chunk_tanggal, chunk_bulan, chunk_tahun int) int {
	var i int = 0
	for i < NMAX {
		if P[sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggal && P[sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulan && P[sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahun {
			return i
		}
		i = i + 1
	}
	return -1
}
func tambahdata(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var chunk_konfirmasitanggal int
	var dbkunjungankosong int

	dbkunjungankosong = pencaridbkunjungankosong(*P, *dbc, *dbk, *sesi)
	if dbkunjungankosong == -1 {
		*dbk = *dbk + 1
	
		fmt.Println("Apakah data pengecekan ini dilakukan pada hari ini?")
		fmt.Println("1 : Ya")
		fmt.Println("2 : Tidak")
		fmt.Print("Masukkan Data :")
		fmt.Scan(&chunk_konfirmasitanggal)
		if chunk_konfirmasitanggal == 1 {
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal = currentTime.Day()
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan = int(currentTime.Month())
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun = currentTime.Year()
		} else {
			fmt.Println("Masukkan Tanggal Pengecekkan (DD/M/YYYY) :")
			fmt.Scan(&P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan, &P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun)
		}
		fmt.Println("Masukkan Berat Badan :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].beratbadan)
		fmt.Println("Masukkan Tinggi Badan :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tinggibadan)
		fmt.Println("Masukkan Tekanan Darah Distolik :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahdistolik)
		fmt.Println("Masukkan Tekanan Darah Sistolik :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahsistolik)
		P[*sesi].datakesehatan[*dbk].idkunjungan = *dbk
		P[*sesi].datakesehatan[*dbk].iduser = *sesi
		fmt.Println("Data sudah di rekam dengan ID Rekam : ", P[*sesi].datakesehatan[*dbk].idkunjungan)
		fmt.Println("Untuk mengecek kesehatan silahkan ke menu cek kesehatan pada halaman utama")
	} else {
		fmt.Println("Apakah data pengecekan ini dilakukan pada hari ini?")
		fmt.Println("1 : Ya")
		fmt.Println("2 : Tidak")
		fmt.Print("Masukkan Data :")
		fmt.Scan(&chunk_konfirmasitanggal)
		if chunk_konfirmasitanggal == 1 {
			P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.tanggal = currentTime.Day()
			P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.bulan = int(currentTime.Month())
			P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.tahun = currentTime.Year()
		} else {
			fmt.Println("Masukkan Tanggal Pengecekkan (DD/M/YYYY) :")
			fmt.Scan(&P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.bulan, &P[*sesi].datakesehatan[dbkunjungankosong].tanggalpengecekan.tahun)
		}
		fmt.Println("Masukkan Berat Badan :")
		fmt.Scan(&P[*sesi].datakesehatan[dbkunjungankosong].beratbadan)
		fmt.Println("Masukkan Tinggi Badan :")
		fmt.Scan(&P[*sesi].datakesehatan[dbkunjungankosong].tinggibadan)
		fmt.Println("Masukkan Tekanan Darah Distolik :")
		fmt.Scan(&P[*sesi].datakesehatan[dbkunjungankosong].tekanandarahdistolik)
		fmt.Println("Masukkan Tekanan Darah Sistolik :")
		fmt.Scan(&P[*sesi].datakesehatan[dbkunjungankosong].tekanandarahsistolik)
		P[*sesi].datakesehatan[dbkunjungankosong].idkunjungan = dbkunjungankosong
		P[*sesi].datakesehatan[dbkunjungankosong].iduser = *sesi
		fmt.Println("Data sudah di rekam dengan ID Rekam : ", P[*sesi].datakesehatan[*dbk].idkunjungan)
		fmt.Println("Untuk mengecek kesehatan silahkan ke menu cek kesehatan pada halaman utama")
	}

}

func pencaridbkunjungankosong(P atributpasien, dbc, dbk, sesi int) int {
	var dbkunjungankosong int
	dbkunjungankosong = 0
	for dbkunjungankosong < dbc {
		if P[sesi].datakesehatan[dbkunjungankosong].statusKunjungan == 2 {
			return dbkunjungankosong
		}
		dbkunjungankosong = dbkunjungankosong + 1
	}
	return -1
}

// func editdata(P *atributpasien, dbc *int, dbk *int, sesi *int) {
// 	var i, xubah int
// 	var d int
// 	var chunk_tanggaldari, chunk_bulandari, chunk_tahundari int
// 	var chunk_tanggalke, chunk_bulanke, chunk_tahunke int
// 	var s1, s2, s3 int

// 	fmt.Println("Masukkan range tanggal kunjungan (Dari - Ke) :")
// 	fmt.Print("Dari :")
// 	fmt.Scan(&chunk_tanggaldari, &chunk_bulandari, &chunk_tahundari)
// 	fmt.Print("Ke :")
// 	fmt.Scan(&chunk_tanggalke, &chunk_bulanke, &chunk_tahunke)
// 	for i <= *dbk && i < NMAX{
// 		if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal >= chunk_tanggaldari && P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal <= chunk_tanggalke {
// 			if P[*sesi].datakesehatan[i].tanggalpengecekan.bulan >= chunk_bulandari && P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal <= chunk_bulanke {
// 				if P[*sesi].datakesehatan[i].tanggalpengecekan.tahun >= chunk_tahundari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun <= chunk_tahunke {
// 					if s1 == 0 {
// 						s1 = i
// 					} else if s1 > 0 {
// 						s2 = i
// 					} else if s2 > 0 {
// 						s3 = i
// 					}
// 				}
// 			}
// 		}
// 		i = i + 1
// 	}
// 	fmt.Println("Pilih data yang akan di edit : ")
// 	fmt.Println("ID     TANGGAL       BULAN     TAHUN")
// 	fmt.Println("Data 1", P[*sesi].datakesehatan[s1].idkunjungan, P[*sesi].datakesehatan[s1].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s1].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s1].tanggalpengecekan.tahun)
// 	fmt.Println("Data 2", P[*sesi].datakesehatan[s2].idkunjungan, P[*sesi].datakesehatan[s2].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s2].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s2].tanggalpengecekan.tahun)
// 	fmt.Println("Data 3", P[*sesi].datakesehatan[s3].idkunjungan, P[*sesi].datakesehatan[s3].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s3].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s3].tanggalpengecekan.tahun)
// 	fmt.Println("Masukkan Nomor :")
// 	fmt.Scan(&d)
// 	if d == 1 {
// 		xubah = s1
// 	} else if d == 2 {
// 		xubah = s2
// 	} else if d == 3 {
// 		xubah = s3
// 	}
// 	fmt.Print("Data yang dirubah adalah data dengan ID kunjungan :", P[*sesi].datakesehatan[xubah].idkunjungan)
// 	fmt.Print("Masukkan berat badan yang ingin diubah :")
// 	fmt.Scan(&P[*sesi].datakesehatan[xubah].beratbadan)
// 	fmt.Print("Masukkan tinggi badan yang ingin diubah :")
// 	fmt.Scan(&P[*sesi].datakesehatan[xubah].tinggibadan)
// 	fmt.Print("Masukkan berat badan yang ingin diubah :")
// 	fmt.Scan(&P[*sesi].datakesehatan[xubah].tekanandarah)
// 	fmt.Print("Masukkan tinggi badan yang ingin diubah :")
// 	fmt.Scan(&P[*sesi].datakesehatan[xubah].tinggibadan)
// }

func hapusdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var chunk_pemilihanmetodecari, chunk_idcari int
	var chunk_tanggal, chunk_bulan, chunk_tahun int
	var wantDeletedID int

	fmt.Println("Berdasarkan apa properti data yang ingin di hapus (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Println("Masukkan ID Kunjungan yang ingin dihapus")
		fmt.Scan(&chunk_idcari)
		wantDeletedID = Sequentialsearch(*P, *dbk, chunk_idcari)

		P[*sesi].datakesehatan[wantDeletedID].statusKunjungan = 2
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan data berdasarkan waktu kunjungan yang ingin dihapus (DD/M/YYYY)")
		fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
		wantDeletedID = pencaritanggal(*P, *dbc, *dbk, *sesi, chunk_tanggal, chunk_bulan, chunk_tahun)

		P[*sesi].datakesehatan[wantDeletedID].statusKunjungan = 2
	}
}

func cekkesehatan(P *atributpasien, dbk *int, dbc *int, sesi *int) {
	var chunk_tanggal, chunk_bulan, chunk_tahun int
	var stopper bool
	var i int

	fmt.Println("Silahkan masukkan tanggal kunjungan")
	fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)

	for i < *dbk && !stopper {
		if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggal && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulan && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahun {
			P[*sesi].datakesehatan[*dbk].diagnosabmi.hasil = P[*sesi].datakesehatan[*dbk].beratbadan / ((P[*sesi].datakesehatan[*dbk].tinggibadan / 100) * (P[*sesi].datakesehatan[*dbk].tinggibadan / 100))
			switch {
			case P[*sesi].datakesehatan[i].diagnosabmi.hasil < 17:
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "Kurus"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "Kekurangan berat badan tingkat rendah"
			case P[*sesi].datakesehatan[i].diagnosabmi.hasil >= 17 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 18.5:
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "Kurus"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "Kekurangan berat badan tingkat sedang"
			case P[*sesi].datakesehatan[i].diagnosabmi.hasil >= 18.5 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 25:
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "Normal"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "Berat badan normal"
			case P[*sesi].datakesehatan[i].diagnosabmi.hasil >= 25 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 30:
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "Gemuk"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "Kelebihan berat badan tingkat rendah"
			case P[*sesi].datakesehatan[i].diagnosabmi.hasil >= 30:
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "Gemuk"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "Kelebihan berat badan tingkat tinggi"
			}

			switch {
			case P[*sesi].datakesehatan[i].tekanandarahsistolik < 90 && P[*sesi].datakesehatan[i].tekanandarahdistolik < 60:
				P[*sesi].datakesehatan[i].diagnosatekanandarah.kelompok = "Rendah"
				P[*sesi].datakesehatan[i].diagnosatekanandarah.kategori = "Tekanan darah rendah (Hipotensi)"
			case P[*sesi].datakesehatan[i].tekanandarahsistolik >= 90 && P[*sesi].datakesehatan[i].tekanandarahsistolik < 120 && P[*sesi].datakesehatan[i].tekanandarahdistolik >= 60 && P[*sesi].datakesehatan[i].tekanandarahdistolik < 80:
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kelompok = "Normal"
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kategori = "Tekanan darah normal"
			case P[*sesi].datakesehatan[i].tekanandarahsistolik >= 120 && P[*sesi].datakesehatan[i].tekanandarahsistolik < 140 && P[*sesi].datakesehatan[i].tekanandarahdistolik >= 80 && P[*sesi].datakesehatan[i].tekanandarahdistolik < 90:
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kelompok = "Pra-Hipertensi"
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Pra-Hipertensi)"
			case P[*sesi].datakesehatan[i].tekanandarahsistolik >= 140 && P[*sesi].datakesehatan[i].tekanandarahdistolik >= 90:
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kelompok = "Tinggi"
				P[*sesi].datakesehatan[*dbk].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Hipertensi)"
				stopper = true
			}
		} else {
			i = i + 1
		}
	}
}

func DisplayAkun(P *atributpasien, dbc *int) {
	var n, i, m int
	n = 0
	m = 0
	fmt.Println("=============================================")
	fmt.Println("      LIST PASIEN RUMAH SAKIT TELKOM         ")
	fmt.Println("=============================================")
	for n <= *dbc {
		if P[n].status != 2 {
			fmt.Println("-----------------------------------------")
			fmt.Println("ID :", P[m].idpasien)
			fmt.Println("Nama: ", P[m].namapasien.namadepan, P[n].namapasien.namabelakang)
			fmt.Println("usia: ", P[m].usia)
			fmt.Println("TTL :", P[m].tempatlahir, ",", P[m].tanggallahir.tanggal, P[m].tanggallahir.bulan, P[m].tanggallahir.tahun)
			fmt.Println("Gender :", P[m].gender)
			i = 0
			for P[n].datakesehatan[i].idkunjungan != 0 {
				fmt.Println("id Kunjungan: ", P[m].datakesehatan[i].idkunjungan)
				fmt.Println("Tanggal Kunjungan: ", P[m].datakesehatan[i].tanggalpengecekan.tanggal, "/", P[m].datakesehatan[i].tanggalpengecekan.bulan, "/", P[m].datakesehatan[i].tanggalpengecekan.tahun)
				fmt.Println("===============[Hasil Diagnosa]===============")
				fmt.Println("Berat Badan: ", P[m].datakesehatan[i].beratbadan)
				fmt.Println("Tinggi Badan: ", P[m].datakesehatan[i].tinggibadan)
				fmt.Println("Hasil BMI: ", P[m].datakesehatan[i].diagnosabmi.hasil)
				fmt.Println("Termasuk dalam kelompok: ", P[m].datakesehatan[i].diagnosabmi.kelompok)
				fmt.Println("detail: ", P[m].datakesehatan[i].diagnosabmi.kategori)
				fmt.Println()
				fmt.Println("Tekanan Darah (Sitolik,Diastolik): ", P[m].datakesehatan[i].tekanandarahsistolik, P[m].datakesehatan[i].tekanandarahdistolik)
				fmt.Println("Hasil Uji tekanan darah: ", P[m].datakesehatan[i].diagnosatekanandarah.hasil)
				fmt.Println("Termasuk dalam kelompok: ", P[m].datakesehatan[i].diagnosatekanandarah.kelompok)
				fmt.Println("detail: ", P[m].datakesehatan[i].diagnosatekanandarah.kategori)
				fmt.Println("=============================================")
				i++
			}
			n++
			m++
		} else {
			m++
		}
	}
}

func EditDatabasePasien(P *atributpasien, dbc *int) {
	var id int
	var Pointedid int
	var chc int
	fmt.Println("Masukkan ID Pasien yang ingin diedit:")
	fmt.Scan(&id)
	Pointedid = Sequentialsearch(*P, *dbc, id)
	if P[Pointedid].idpasien != 0 {
		fmt.Println("-----------------------------------------")
		fmt.Println("ID :", P[Pointedid].idpasien)
		fmt.Println("Nama: ", P[Pointedid].namapasien.namadepan, P[Pointedid].namapasien.namabelakang)
		fmt.Println("usia: ", P[Pointedid].usia)
		fmt.Println("TTL :", P[Pointedid].tempatlahir, ",", P[Pointedid].tanggallahir.tanggal, P[Pointedid].tanggallahir.bulan, P[Pointedid].tanggallahir.tahun)
		fmt.Println("Gender :", P[Pointedid].gender)
		fmt.Println("Pilih Atribut yang akan di-Edit:")
		fmt.Println("-----------------------------------------")
		fmt.Println()
		fmt.Println("<1> Nama")
		fmt.Println("<2> usia")
		fmt.Println("<3> tempat Lahir")
		fmt.Println("<4> tanggal Lahir")
		fmt.Println("<5> gender")
		fmt.Println()
		fmt.Println("Pilih angka sesuai Atribut yang di-Edit:")
		fmt.Scan(&chc)

		if chc == 1 {
			fmt.Println("Masukkan Nama Baru:")
			fmt.Scan(&P[Pointedid].namapasien.namadepan, &P[Pointedid].namapasien.namabelakang)
			fmt.Println("Data sudah di-update!")
		} else if chc == 2 {
			fmt.Println("Masukkan Usia:")
			fmt.Scan(&P[Pointedid].usia)
			fmt.Println("Data sudah di-update!")
		} else if chc == 3 {
			fmt.Println("Masukkan Tempat Lahir:")
			fmt.Scan(&P[Pointedid].tempatlahir)
			fmt.Println("Data sudah di-update!")
		} else if chc == 4 {
			fmt.Println("Masukkan Tanggal Lahir:")
			fmt.Scan(&P[Pointedid].tanggallahir.tanggal, &P[Pointedid].tanggallahir.bulan, &P[Pointedid].tanggallahir.tahun)
			fmt.Println("Data sudah di-update!")
		} else if chc == 5 {
			fmt.Println("Masukkan Gender:")
			fmt.Scan(&P[Pointedid].gender)
			fmt.Println("Data sudah di-update!")
		}
	} else {
		fmt.Print("Data tidak Ditemukan")
	}
}

func sortSelectionDesc(P *atributpasien, dbc *int) {
	var ghostArray atributpasien
	var pass, i, idx, idxghost int
	var temp datapasien
	ghostArray = *P

	for j := 1; j < *dbc && P[j].status == 1; j++ {
		ghostArray[idxghost] = P[j]
		idxghost++
	}

	for pass = 1; pass < idxghost; pass++ {
		idx = pass - 1
		for i = pass; i < idxghost; i++ {
			if ghostArray[idx].idpasien < ghostArray[i].idpasien {
				idx = i
			}
		}
		temp = ghostArray[pass-1]
		ghostArray[pass-1] = ghostArray[idx]
		ghostArray[idx] = temp
	}
	DisplayAkun(&ghostArray, &idxghost)
}

func sortInsertionAsx(P *atributpasien, dbc *int) {
	var pass, i, idxGoib int
	var temp datapasien
	var arrayGoib atributpasien

	for j := 1; j < *dbc && P[j].status == 1; j++ {
		arrayGoib[idxGoib] = P[j]
		idxGoib++
	}
	for pass = 1; pass < idxGoib; pass++ {
		i = pass
		temp = arrayGoib[pass]
		for i > 0 && temp.namapasien.namadepan < arrayGoib[i-1].namapasien.namadepan {
			arrayGoib[i] = arrayGoib[i-1]
			i = i - 1
		}
		P[i] = temp
	}
	DisplayAkun(&arrayGoib, &idxGoib)
}

func Sequentialsearch(P atributpasien, dbc int, x int) int {
	var stopper bool
	var i int
	var indeX int
	i = 1
	stopper = false
	for i <= dbc && !stopper {
		if P[i].idpasien == x {
			indeX = i
			stopper = true
		}
		i++
	}
	return indeX
}

func cariAkunPasien(P *atributpasien, dbc *int, sesi *int) {
	var IDtarget, idx int
	var oldestRecord, BrandRecord, middle int
	var found bool
	var choice string

	idx = 0
	found = false
	oldestRecord = 1
	BrandRecord = *dbc-1
	fmt.Println("------------------------------------------------")
	fmt.Println("Masukkan ID pasien yang ingin kamu cari:")
	fmt.Println("------------------------------------------------")
	fmt.Print("> ")
	fmt.Scan(&IDtarget)

	for oldestRecord <= BrandRecord && found == false {
		middle = (oldestRecord + BrandRecord) / 2
		if P[middle].idpasien == IDtarget {
			idx = middle
			found = true
		} else if P[middle].idpasien > IDtarget {
			BrandRecord = middle - 1
		} else {
			oldestRecord = middle + 1
		}
	}
	if found {
		fmt.Println("ID ditemukan pada record pasien ke-", idx, ". Apa data pasien mau ditampilkan")
		fmt.Println("[Y/y untuk melanjutkan]") // sorry ini gua lagi males
		fmt.Print("> ")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			DisplayAkun(P, dbc)
		} else {
			return
		}
	}
}
