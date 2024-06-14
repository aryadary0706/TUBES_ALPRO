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
	status        int // 1 = aktif ; 2 = tidak aktif
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
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang)
		fmt.Println("Dengan ID User :", P[*&sesilogin].idpasien)
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
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang)
		fmt.Println("Dengan ID User :", P[*&sesilogin].idpasien)
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
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang)
		fmt.Println("Dengan ID User :", P[*&sesilogin].idpasien)
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
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau pilih data Pasien")
	} else {
		fmt.Println("Saat ini, Terdaftar sebagai :", P[*&sesilogin].namapasien.namadepan, P[*&sesilogin].namapasien.namabelakang)
		fmt.Println("Dengan ID User :", P[*&sesilogin].idpasien)
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
		// fmt.Println("dbc :", dbc)
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
					lihatdatakunjungan(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 2 {
					tambahdatakunjungan(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 3 {
					ubahdatakunjungan(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 4 {
					hapusdatakunjungan(&datapasien, &dbc, &dbk, &sesi)
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
			keluar(&datapasien, &dbc, &sesi)
		} else if pilihan == 5 {
			var uruti int
			pilihan = 1
			for pilihan >= 1 && pilihan <= 4 {
				interface05(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Println()
					fmt.Println("<1>. Tampilkan berdasarkan Data Terkini")
					fmt.Println("<2>. Tampilkan berdasarkan Nama Pasien")
					fmt.Println()
					fmt.Println("Masukkan Angka")
					fmt.Scan(&uruti)
					for uruti != 1 && uruti != 2 {
						fmt.Scan(&uruti)
					}
					if uruti == 1 {
						sortSelectionDesc(datapasien, dbc)

					} else if uruti == 2 {
						sortInsertionAsx(datapasien, dbc)

					}
				} else if pilihan == 2 {
					cariAkunPasien(datapasien, dbc)
				} else if pilihan == 3 {
					EditDatabasePasien(&datapasien, dbc)
				} else if pilihan == 4 {
					sesi = -1
				}
			}
		}
	}
}

// SubProgram per-layanan

func daftarpengguna(P *atributpasien, dbc *int, sesi *int) {
	var chunk_gender, chunk_password, chunk_password2 string
	var chunk_namadepan, chunk_namabelakang, chunk_tempatlahir string
	var chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir int
	var dbsama int
	var dbkosong int
	var datefillcounter int
	var loopstatuscode int

	// idxstatuskosong := pencaristatuskosong(*P, *dbc)
	dbkosong = pencaristatuskosong(*P, *dbc)
	if dbkosong == -1 { //tidak ada DB yang kosong
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for pengecektanggal(chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir) == false && datefillcounter < 3 {
			fmt.Println("Tanggal yang anda masukkan salah atau format tidak sesuai, silahkan coba lagi")
			fmt.Print("Masukkan tanggal yang benar (DD/ M/ YYYY): ")
			fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
			datefillcounter = datefillcounter + 1
		}
		dbsama = pencaridbsama(*P, *dbc, chunk_namadepan, chunk_namabelakang, chunk_tempatlahir, chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir)
		if dbsama == -1 {
			*dbc = *dbc + 1
			P[*dbc].namapasien.namadepan = chunk_namadepan
			P[*dbc].namapasien.namabelakang = chunk_namabelakang
			P[*dbc].tempatlahir = chunk_tempatlahir
			P[*dbc].tanggallahir.tanggal = chunk_tanggallahir
			P[*dbc].tanggallahir.bulan = chunk_bulanlahir
			P[*dbc].tanggallahir.tahun = chunk_tahunlahir
			fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
			fmt.Println("L / l = Laki-laki")
			fmt.Println("P / p = Perempuan")
			for loopstatuscode == 0 {
				fmt.Scan(&chunk_gender)
				if chunk_gender == "L" || chunk_gender == "l" {
					P[*dbc].gender = "Laki-laki"
					loopstatuscode = 1
				} else if chunk_gender == "P" || chunk_gender == "p" {
					P[*dbc].gender = "Perempuan"
					loopstatuscode = 1
				} else {
					fmt.Print("Masukkan Gender yang sesuai")
				}
			}
			fmt.Println("Masukkan Password untuk user yang di daftarkan")
			fmt.Scan(&chunk_password)
			fmt.Println("Masukkan lagi password untuk user yang di daftarkan")
			fmt.Scan(&chunk_password2)
			if chunk_password == chunk_password2 {
				P[*dbc].password = chunk_password
			} else {
				for chunk_password != chunk_password2 {
					fmt.Println("Password tidak sama, silahkan coba lagi")
					fmt.Println("Masukkan Password untuk user yang di daftarkan")
					fmt.Scan(&chunk_password)
					fmt.Println("Masukkan lagi password untuk user yang di daftarkan")
					fmt.Scan(&chunk_password2)
				}
				P[*dbc].password = chunk_password
			}

			P[*dbc].idpasien = *dbc
			*sesi = *dbc
			fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
			fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
		} else {
			fmt.Println("Identitas yang anda masukkan sudah terdaftar di sistem, silahkan hubungi petugas jika ada kendala")
		}
	} else {
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for pengecektanggal(chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir) == false && datefillcounter < 3 {
			fmt.Println("Tanggal yang anda masukkan salah atau format tidak sesuai, silahkan coba lagi")
			fmt.Print("Masukkan tanggal lahir yang benar (DD/ M/ YYYY) : ")
			fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
			datefillcounter = datefillcounter + 1
		}
		dbsama = pencaridbsama(*P, *dbc, chunk_namadepan, chunk_namabelakang, chunk_tempatlahir, chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir)
		fmt.Println(dbsama)
		if dbsama == -1 {
			P[dbkosong].namapasien.namadepan = chunk_namadepan
			P[dbkosong].namapasien.namabelakang = chunk_namabelakang
			P[dbkosong].tempatlahir = chunk_tempatlahir
			P[dbkosong].tanggallahir.tanggal = chunk_tanggallahir
			P[dbkosong].tanggallahir.bulan = chunk_bulanlahir
			P[dbkosong].tanggallahir.tahun = chunk_tahunlahir
			fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
			fmt.Println("L / l = Laki-laki")
			fmt.Println("P / p = Perempuan")
			for loopstatuscode == 0 {
				fmt.Scan(&chunk_gender)
				if chunk_gender == "L" || chunk_gender == "l" {
					P[dbkosong].gender = "Laki-laki"
					loopstatuscode = 1
				} else if chunk_gender == "P" || chunk_gender == "p" {
					P[dbkosong].gender = "Perempuan"
					loopstatuscode = 1
				} else {
					fmt.Print("Masukkan Gender yang sesuai")
				}
			}
			fmt.Println("Masukkan Password untuk user yang di daftarkan")
			fmt.Scan(&chunk_password)
			fmt.Println("Masukkan lagi password untuk user yang di daftarkan")
			fmt.Scan(&chunk_password2)
			if chunk_password == chunk_password2 {
				P[*dbc].password = chunk_password
			} else {
				for chunk_password != chunk_password2 {
					fmt.Println("Password tidak sama, silahkan coba lagi")
					fmt.Println("Masukkan Password untuk user yang di daftarkan")
					fmt.Scan(&chunk_password)
					fmt.Println("Masukkan lagi password untuk user yang di daftarkan")
					fmt.Scan(&chunk_password2)
				}
				P[dbkosong].password = chunk_password
			}
			
			P[dbkosong].status = 1
			P[dbkosong].idpasien = dbkosong
			*sesi = dbkosong
			fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
			fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
		} else {
			fmt.Println("Identitas yang anda masukkan sudah terdaftar di sistem, silahkan hubungi petugas jika ada kendala")
		}
	}
}

func pengecektanggal(check_tanggal, check_bulan, check_tahun int) bool {
	var validasi bool

	if check_tahun > 0 {
		if (check_tahun % 4 == 0 && check_tahun % 100 != 0) || (check_tahun % 400 == 0) {
			if check_bulan == 1 || check_bulan == 3 || check_bulan == 5 || check_bulan == 7 || check_bulan == 8 || check_bulan == 10 || check_bulan == 12 {
				if check_tanggal > 0 && check_tanggal <= 31 {
					validasi = true
				} else {
					validasi = false
				}
			} else if check_bulan == 4 || check_bulan == 6 || check_bulan == 9 || check_bulan == 11 {
				if check_tanggal > 0 && check_tanggal <= 30 {
					validasi = true
				} else {
					validasi = false
				}
			} else if check_bulan == 2 {
				if check_tanggal > 0 && check_tanggal <= 29 {
					validasi = true
				} else {
					validasi = false
				}
			} 
		} else {
			if check_bulan == 1 || check_bulan == 3 || check_bulan == 5 || check_bulan == 7 || check_bulan == 8 || check_bulan == 10 || check_bulan == 12 {
				if check_tanggal > 0 && check_tanggal <= 31 {
					validasi = true
				} else {
					validasi = false
				}
			} else if check_bulan == 4 || check_bulan == 6 || check_bulan == 9 || check_bulan == 11 {
				if check_tanggal > 0 && check_tanggal <= 30 {
					validasi = true
				} else {
					validasi = false
				}
			} else if check_bulan == 2 {
				if check_tanggal > 0 && check_tanggal <= 28 {
					validasi = true
				} else {
					validasi = false
				}
			}
		}
	}
	return validasi
}
func pencaridbsama(P atributpasien, dbc int, namadepan, namabelakang, tempatlahir string, tanggallahir, bulanlahir, tahunlahir int) int {
	var i int = 0

	for i < dbc {
		if namadepan == P[i].namapasien.namadepan && namabelakang == P[i].namapasien.namabelakang && tempatlahir == P[i].tempatlahir && tanggallahir == P[i].tanggallahir.tanggal && bulanlahir == P[i].tanggallahir.bulan && tahunlahir == P[i].tanggallahir.tahun {
			return i
		}
		i = i + 1
	}
	return -1
}

func pencaristatuskosong(P atributpasien, dbc int) int {
	var i int
	i = 0
	for i < dbc {
		if P[i].status == 2 {
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
	var chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir int
	var datefillcounter int
	var loopstatuscode int

	chunk_pernyataan = 1
	fmt.Println("Apakah data pasien yang akan di rubah adalah ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Println("Masukkan angka 1 / 0 (1 untuk BENAR, 0 untuk SALAH)")
	fmt.Scan(&chunk_pernyataan)
	fmt.Println("Masukkan kata sandi / password :")
	fmt.Scan(&chunk_password)
	fmt.Println("Masukkan kembali kata sandi / password :")
	fmt.Scan(&chunk_password2)
	if *sesi != -1 && chunk_pernyataan == 1 && chunk_password == P[*sesi].password && chunk_password2 == P[*sesi].password && chunk_password == chunk_password2 {
		fmt.Println("Masukkan Nama Depan yang ingin di ubah")
		fmt.Scan(&P[*sesi].namapasien.namadepan)
		fmt.Println("Masukkan Nama Belakang yang ingin di ubah")
		fmt.Scan(&P[*sesi].namapasien.namabelakang)
		fmt.Println("Masukkan Tempat Lahir yang ingin di ubah")
		fmt.Scan(&P[*sesi].tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir yang ingin di ubah (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for pengecektanggal(chunk_tanggallahir, chunk_bulanlahir, chunk_tahunlahir) == false && datefillcounter < 3 {
			fmt.Println("Tanggal yang anda masukkan salah atau format tidak sesuai, silahkan coba lagi")
			fmt.Print("Masukkan tanggal lahir yang benar (DD/ M/ YYYY) : ")
			fmt.Scan(&chunk_tanggallahir, &chunk_bulanlahir, &chunk_tahunlahir)
			datefillcounter = datefillcounter + 1
		}
		P[*sesi].tanggallahir.tanggal = chunk_tanggallahir
		P[*sesi].tanggallahir.bulan = chunk_bulanlahir
		P[*sesi].tanggallahir.tahun = chunk_tahunlahir
		fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
		fmt.Println("L / l = Laki-laki")
		fmt.Println("P / p = Perempuan")
		for loopstatuscode == 0 {
			fmt.Scan(&chunk_gender)
			if chunk_gender == "L" || chunk_gender == "l" {
				P[*sesi].gender = "Laki-laki"
				loopstatuscode = 1
			} else if chunk_gender == "P" || chunk_gender == "p" {
				P[*sesi].gender = "Perempuan"
				loopstatuscode = 1
			} else {
				fmt.Print("Masukkan Gender yang sesuai")
			}
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

	fmt.Println("================= PERINGATAN =================")
	fmt.Println("  Data Kunjungan dan akun anda akan di hapus  ")
	fmt.Println("   dan jika Anda ingin menggunakan aplikasi   ")
	fmt.Println("  Anda harus mendaftar ulang di menu daftar   ")
	fmt.Println("==============================================")
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
		fmt.Println("| Kamu membatalkan penghapusan akun dengan tidak |")
		fmt.Println("| menyetujui pernyataan penghapusan atau kamu    |")
		fmt.Println("|    Terjadi kesalahan pemasukkan password       |")
	}
}

func keluar(P *atributpasien, dbc *int, sesi *int) {
	*sesi = -1
	fmt.Println("Anda sudah keluar, terimakasih sudah menggunakan aplikasi kami")
}

func lihatdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
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
func tambahdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var chunk_konfirmasitanggal int
	var dbkunjungankosong int
	var chunk_tanggalkunjungan, chunk_bulankunjungan, chunk_tahunkunjungan int
	var datefillcounter int

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
			fmt.Scan(&chunk_tanggalkunjungan, &chunk_bulankunjungan, &chunk_tahunkunjungan)
				for pengecektanggal(chunk_tanggalkunjungan, chunk_bulankunjungan, chunk_tahunkunjungan) == false && datefillcounter < 3 {
				fmt.Println("Tanggal yang anda masukkan salah atau format tidak sesuai, silahkan coba lagi")
				fmt.Print("Masukkan tanggal kunjungan yang benar (DD/ M/ YYYY) : ")
				fmt.Scan(chunk_tanggalkunjungan, &chunk_bulankunjungan, &chunk_tahunkunjungan)
				datefillcounter = datefillcounter + 1
			}
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal = chunk_tanggalkunjungan
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan = chunk_bulankunjungan
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun =chunk_tahunkunjungan
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
			fmt.Println("Masukkan Tanggal Pengecekkan (DD/M/YYYY) :")
			fmt.Scan(&chunk_tanggalkunjungan, &chunk_bulankunjungan, &chunk_tahunkunjungan)
				for pengecektanggal(chunk_tanggalkunjungan, chunk_bulankunjungan, chunk_tahunkunjungan) == false && datefillcounter < 3 {
				fmt.Println("Tanggal yang anda masukkan salah atau format tidak sesuai, silahkan coba lagi")
				fmt.Print("Masukkan tanggal kunjungan yang benar (DD/ M/ YYYY) : ")
				fmt.Scan(chunk_tanggalkunjungan, &chunk_bulankunjungan, &chunk_tahunkunjungan)
				datefillcounter = datefillcounter + 1
			}
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal = chunk_tanggalkunjungan
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan = chunk_bulankunjungan
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun = chunk_tahunkunjungan
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
	for dbkunjungankosong <= dbc {
		if P[sesi].datakesehatan[dbkunjungankosong].statusKunjungan == 2 {
			return dbkunjungankosong
		}
		dbkunjungankosong = dbkunjungankosong + 1
	}
	return -1
}

func ubahdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var i, n, x int
	var chunk_tanggal, chunk_bulan, chunk_tahun int
	var chunk_tanggalcari, chunk_bulancari, chunk_tahuncari int
	var chunk_id int
	var chunk_beratbadan, chunk_tinggibadan, chunk_tekanandarahdistolik, chunk_tekanandarahsistolik float64
	var chunk_pemilihanmetodecari, chunk_konfirmasiperubahan int
	var arrpencaritanggal [100]int

	fmt.Println("Berdasarkan apa properti data yang ingin di ubah (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Print("Masukkan ID Kunjungan yang ingin diubah datanya :")
		fmt.Scan(&chunk_id)
		if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].idkunjungan != -1 {
			fmt.Println("Data ditemukan dengan keterangan berikut : ")
			fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
			fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
			fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
			fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
			fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
			fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
			fmt.Println("Apakah data yang ingin di ubah benar?")
			fmt.Scan(&chunk_konfirmasiperubahan)
			if chunk_konfirmasiperubahan == 1 {
				fmt.Println("Masukkan tanggal kunjungan yang baru (isi 0 0 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tanggal kunjungan data lama : ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
				fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
				if chunk_tanggal != 0 && chunk_bulan != 0 && chunk_tahun != 0 {
					P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal = chunk_tanggal
					P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan = chunk_bulan
					P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun = chunk_tahun
				}
				fmt.Println("Masukkan berat badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Berat badan data lama : ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
				fmt.Scan(&chunk_beratbadan)
				if chunk_beratbadan != 0 {
					P[*sesi].datakesehatan[chunk_id].beratbadan = chunk_beratbadan
				}
				fmt.Println("Masukkan tinggi badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tinggi badan data lama : ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
				fmt.Scan(&chunk_tinggibadan)
				if chunk_tinggibadan != 0 {
					P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tinggibadan
				}
				fmt.Println("Masukkan tekanan darah distolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tekanan darah distolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Scan(&chunk_tekanandarahdistolik)
				if chunk_tekanandarahdistolik != 0 {
					P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tekanandarahdistolik
				}
				fmt.Println("Masukkan tekanan darah sistolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tekanan darah sistolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				fmt.Scan(&chunk_tekanandarahsistolik)
				if chunk_tekanandarahsistolik != 0 {
					P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tekanandarahsistolik
				}
				fmt.Println("Data kunjungan sudah di rubah dengan isi :")
				fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
				fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
				fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
				fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
				fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
			} else {
				fmt.Println("Anda membatalkan aksi mengubah data")
			}
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	} else {
		fmt.Println("Masukkan tanggal kunjungan :")
		fmt.Scan(&chunk_tanggalcari, &chunk_bulancari, &chunk_tahuncari)
		for i < *dbk {
			if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggalcari && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulancari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahuncari {
				arrpencaritanggal[n] = i
				n = n + 1
			}
			i = i + 1
		}
		if n != 0 {
			fmt.Println("Data yang di temukan dengan prefrensi tanggal, adalah :")
			fmt.Printf("TANGGAL KUNJUNGAN       ID KUNJUNGAN       BERAT BADAN       TINGGI BADAN       TEKANAN DARAH DISTOLIK       TEKANAN DARAH SISTOLIK")
			for x <= n {
				fmt.Printf("%d/%d/%d %7d %7d kg %7d m %7f mmHg %7f mmHg", P[*sesi].datakesehatan[arrpencaritanggal[x]].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[arrpencaritanggal[x]].tanggalpengecekan.bulan, P[*sesi].datakesehatan[arrpencaritanggal[x]].tanggalpengecekan.tahun, P[*sesi].datakesehatan[arrpencaritanggal[x]].idkunjungan, P[*sesi].datakesehatan[arrpencaritanggal[x]].beratbadan, P[*sesi].datakesehatan[arrpencaritanggal[x]].tinggibadan, P[*sesi].datakesehatan[arrpencaritanggal[x]].tekanandarahdistolik, P[*sesi].datakesehatan[arrpencaritanggal[x]].tekanandarahsistolik)
				x = x + 1
			}
			fmt.Println("Masukkan data sesuai ID Kunjungan :")
			fmt.Scan(&chunk_id)
			if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].idkunjungan != -1 {
				fmt.Println("Data ditemukan dengan keterangan berikut : ")
				fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
				fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
				fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
				fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
				fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				fmt.Println("Apakah data yang ingin di ubah benar?")
				fmt.Scan(&chunk_konfirmasiperubahan)
				if chunk_konfirmasiperubahan == 1 {
					fmt.Println("Masukkan tanggal kunjungan yang baru (isi 0 0 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tanggal kunjungan data lama : ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
					fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
					if chunk_tanggal != 0 && chunk_bulan != 0 && chunk_tahun != 0 {
						P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal = chunk_tanggal
						P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan = chunk_bulan
						P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun = chunk_tahun
					}
					fmt.Println("Masukkan berat badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Berat badan data lama : ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
					fmt.Scan(&chunk_beratbadan)
					if chunk_beratbadan != 0 {
						P[*sesi].datakesehatan[chunk_id].beratbadan = chunk_beratbadan
					}
					fmt.Println("Masukkan tinggi badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tinggi badan data lama : ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
					fmt.Scan(&chunk_tinggibadan)
					if chunk_tinggibadan != 0 {
						P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tinggibadan
					}
					fmt.Println("Masukkan tekanan darah distolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tekanan darah distolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
					fmt.Scan(&chunk_tekanandarahdistolik)
					if chunk_tekanandarahdistolik != 0 {
						P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tekanandarahdistolik
					}
					fmt.Println("Masukkan tekanan darah sistolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tekanan darah sistolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
					fmt.Scan(&chunk_tekanandarahsistolik)
					if chunk_tekanandarahsistolik != 0 {
						P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tekanandarahsistolik
					}
					fmt.Println("Data kunjungan sudah di rubah dengan isi :")
					fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
					fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
					fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
					fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
					fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
					fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				} else {
					fmt.Println("Anda membatalkan aksi mengubah data")
				}
			} else {
				fmt.Println("Data tidak ditemukan, silahkan coba lagi")
			}
		}
	}
}

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

func DisplayAkun(P atributpasien, dbc int) {
	var n, i, m int
	n = 0
	m = 0
	fmt.Println("=============================================")
	fmt.Println("      LIST PASIEN RUMAH SAKIT TELKOM         ")
	fmt.Println("=============================================")
	for n <= dbc-1 {
		if P[m].status != 2 {
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

func EditDatabasePasien(P *atributpasien, dbc int) {
	var id int
	var Pointedid int
	var chc int
	fmt.Println("Masukkan ID Pasien yang ingin diedit:")
	fmt.Scan(&id)
	Pointedid = Sequentialsearch(*P, dbc, id)
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

func sortSelectionDesc(P atributpasien, dbc int) {
	var ghostArray atributpasien
	var pass, i, idx, idxghost int
	var temp datapasien
	ghostArray = P

	for j := 1; j <= dbc && P[j].status == 1; j++ {
		ghostArray[idxghost] = P[j]
		fmt.Println("Nama: ", ghostArray[idxghost].namapasien.namadepan, ghostArray[idxghost].namapasien.namabelakang)
		idxghost++
	}
	fmt.Println("index array baru: ", idxghost)

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
	DisplayAkun(ghostArray, idxghost)
}

func sortInsertionAsx(P atributpasien, dbc int) {
	var pass, i, idxGoib int
	var temp datapasien
	var arrayGoib atributpasien

	for j := 1; j <= dbc && P[j].status == 1; j++ {
		arrayGoib[idxGoib] = P[j]
		fmt.Println("Nama: ", arrayGoib[idxGoib].namapasien.namadepan, arrayGoib[idxGoib].namapasien.namabelakang)
		idxGoib++
	}
	fmt.Println("index array baru: ", idxGoib)
	for pass = 1; pass < idxGoib; pass++ {
		i = pass
		temp = arrayGoib[pass]
		for i > 0 && temp.namapasien.namadepan < arrayGoib[i-1].namapasien.namadepan {
			arrayGoib[i] = arrayGoib[i-1]
			i = i - 1
		}
		P[i] = temp
	}
	DisplayAkun(arrayGoib, idxGoib)
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

func cariAkunPasien(P atributpasien, dbc int) {
	var IDtarget, idx int
	var oldestRecord, BrandRecord, middle int
	var found bool
	var choice string

	idx = 0
	found = false
	oldestRecord = 0
	BrandRecord = dbc - 1
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
		fmt.Println("[Y/y untuk melanjutkan] dan [T/t untuk kembali ke Menu]") // sorry ini gua lagi males
		fmt.Print("> ")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {

		} else {
			return
		}
	}
}
