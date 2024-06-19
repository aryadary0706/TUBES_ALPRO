package main

//import eketensi yang diperlukan
import (
	"fmt"
	"time"
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
	guladarah               float64
	beratbadan, tinggibadan float64
	diagnosabmi             diagnosa
	diagnosatekanandarah    diagnosa
	diagnosaguladarah       diagnosa
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
	fmt.Println("1. Profil Pasien")
	fmt.Println("2. Data Kunjungan Pasien")
	fmt.Println("3. Cek Kesehatan")
	fmt.Println("4. Keluar dan Logout")
	if sesilogin == 999 {
		fmt.Println("5. Fitur Admin")
	}
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka : ")
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
	fmt.Print("Masukkan angka : ")
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
	fmt.Println("1. Lihat Data")
	fmt.Println("2. Rekam/Tambah Data")
	fmt.Println("3. Ubah/Edit Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Kembali ke halaman utama")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka : ")
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
	fmt.Println("1. Cek Kesehatan (BMI)")
	fmt.Println("2. Cek Gula Darah (Diabet)")
	fmt.Println("5. Kembali ke halaman utama")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan angka : ")
}

func interface05(sesilogin int) {
	fmt.Println("=====================================================")
	fmt.Println("               DATABASE PASIEN KESEHATAN         ")
	fmt.Println("            RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("=====================================================")
	fmt.Println("<1> Tampilkan Data terurut") // sort descending (selection sort)
	fmt.Println("<2> Check ID Pasien")        //sort ascending (insertion sort)
	fmt.Println("<3> Edit Database Pasien")
	fmt.Println("<4> Keluar dari Menu")
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
			for pilihan >= 1 && pilihan <= 3 {
				interface03(&datapasien, sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					cekkesehatan(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 2 {
					cekguladarah(&datapasien, &dbc, &dbk, &sesi)
				}
			}
		} else if pilihan == 4 {
			keluar(&datapasien, &dbc, &sesi)
		} else if pilihan == 5 {
			var uruti int
			pilihan = 1
			for pilihan >= 1 && pilihan <= 3 {
				interface05(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Println()
					fmt.Println("<1>. Tampilkan berdasarkan Data Terkini")
					fmt.Println("<2>. Tampilkan berdasarkan Umur Pasien (Secara ascended)")
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
				}
			}
		}
	}
}

// SubProgram per-layanan

func daftarpengguna(P *atributpasien, dbc *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc, sesi adalah indikator id akun yang dipakai
	// proses : prosedur mencari database dengan status tidak aktif. jika tidak ditemukan, maka akun akan disimpan di dbc baru
	//			Pengguna memasukkan nama depan, nama belakang, domisili, TTL, gender, dan password.
	//F.S. : Jika syarat-syarat pendaftaran terpenuhi, maka pendaftaran akun berhasil pada dbc baru atau dbc yang statusnya tidak aktif
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
			fmt.Println("Masukkan Usia untuk user yang di daftarkan")
			fmt.Scan(&P[*dbc].usia)
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
			P[*dbc].status = 1
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
			fmt.Println("Masukkan Usia untuk user yang di daftarkan")
			fmt.Scan(&P[*dbc].usia)
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

func masukpengguna(P *atributpasien, dbc *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc, sesi adalah indikator id akun yang dipakai
	//proses : pengguna memasukkan id dan passwod akun
	//F.S : Jika id dan password ada di database, maka sesi mengganti nilai id pasien yang baru masuk. Jika gagal, maka akan keluar dari prosedur
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
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc, sesi adalah indikator id akun yang dipakai
	//proses : pengguna akan menyatakan validitas akun dan memasukkan atribut datapasien yang baru
	//F.S : datapasien pada id akan berubah sesuai datapasien yang dimasukkan
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
		fmt.Println("Masukkan Usia untuk user yang ingin di ubah")
		fmt.Scan(&P[*dbc].usia)
		fmt.Println("Data sudah terubah, dengan keterangan :")
		fmt.Println("Nama Depan           	:", P[*sesi].namapasien.namadepan)
		fmt.Println("Nama Belakang        	:", P[*sesi].namapasien.namabelakang)
		fmt.Println("Tempat Lahir         	:", P[*sesi].tempatlahir)
		fmt.Println("Tanggal/ Tahun Lahir	:", P[*sesi].tanggallahir.tanggal, P[*sesi].tanggallahir.bulan, P[*sesi].tanggallahir.tahun)
		fmt.Println("Gender					:", P[*sesi].gender)
		fmt.Println("Usia					:", P[*sesi].usia)
	} else {
		if *sesi == -1 {
			fmt.Println("Pastikan Anda masuk dengan ID atau akun yang benar, jika ada kendala hubungi petugas")
		} else if P[*sesi].password != chunk_password {
			fmt.Println("Password / kata sandi yang anda masukkan salah, silahkan coba lagi")
		}
	}
}

func ubahpassword(P *atributpasien, dbc *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc, sesi adalah indikator id akun yang dipakai
	//proses : pengguna akan memverifikasi password dan memasukkan password yang baru
	//F.S : password akun pengguna pada sesi terkini akan berubah jika valid. Jika gagal, maka akan keluar dari prosdedur
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
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses : pengguna memasukkan pernyataan penghapusan dan password akun sebelum menghapus akun
	//F.S : Jika masukan data valid, maka atribut status pada akun akan diganti 2 (status tidak aktif). Jika gagal, maka akan keluar dari prosedur
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
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc, sesi adalah indikator id akun yang dipakai
	//F.S : nilai pada atribut sesi akan berganti ke -1, menandakan tidak ada akun yang tertaut saat ini
	*sesi = -1
	fmt.Println("Anda sudah keluar, terimakasih sudah menggunakan aplikasi kami")
}

func lihatdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses : pengguna memasukkan tanggal, bulan, tahun dari arraykunjungan yang ingin dimasukkan. lalu program memproses fungsi pencaritanggal untuk mencari index akun yang tanggal kunjungan  sesuai pada arraykunjungan pada P tersebut
	//F.S : jika index data ditemukan, maka akan men-output data kunjungan. Jika tidak, maka akan keluar dari prosedur
	var chunk_tanggalcari, chunk_bulancari, chunk_tahuncari int
	var chunk_pemilihanmetodecari int
	var chunk_id int
	var arrpencaritanggal [100]int
	var i, n, x int

	fmt.Println("Pastikan Anda login dengan ID yang sesuai")
	fmt.Println("Berdasarkan apa properti data yang ingin di pakai (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Print("Masukkan ID Kunjungan yang ingin digunakan :")
		fmt.Scan(&chunk_id)
		if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].statusKunjungan != 2 {
			fmt.Println("Data ditemukan dengan keterangan berikut : ")
			fmt.Println("ID Kunjungan 				: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
			fmt.Println("Tanggal Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
			fmt.Println("Berat Badan				: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
			fmt.Println("Tinggi Badan				: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
			fmt.Println("Tekanan Darah Distolik 	: ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
			fmt.Println("Tekanan Darah Sistolik 	: ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
			fmt.Println("Gula Darah					: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
		} else {
			fmt.Println("Berdasarkan waktu yang anda masukkan, data tidak ditemukan")
		}
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan tanggal kunjungan :")
		fmt.Scan(&chunk_tanggalcari, &chunk_bulancari, &chunk_tahuncari)
		for i <= *dbk {
			if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggalcari && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulancari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahuncari {
				arrpencaritanggal[n] = i
				n++
			}
			i++
		}
		if n != 0 {
			fmt.Println("Data yang ditemukan dengan preferensi tanggal, adalah:")
			fmt.Printf("%-20s %-15s %-15s %-15s %-25s %-25s\n", "TANGGAL KUNJUNGAN", "ID KUNJUNGAN", "BERAT BADAN (kg)", "TINGGI BADAN (m)", "TEKANAN DARAH DISTOLIK (mmHg)", "TEKANAN DARAH SISTOLIK (mmHg)")
			fmt.Println("==============================================================================================================")
			for x < n {
				datakesehatan := P[*sesi].datakesehatan[arrpencaritanggal[x]]
				fmt.Printf("%02d/%02d/%04d %-15d %-15.2f %-15.2f %-25.2f %-25.2f\n",
					datakesehatan.tanggalpengecekan.tanggal,
					datakesehatan.tanggalpengecekan.bulan,
					datakesehatan.tanggalpengecekan.tahun,
					datakesehatan.idkunjungan,
					datakesehatan.beratbadan,
					datakesehatan.tinggibadan,
					datakesehatan.tekanandarahdistolik,
					datakesehatan.tekanandarahsistolik)
				x++
			}
			fmt.Println("Masukkan data sesuai ID Kunjungan :")
			fmt.Scan(&chunk_id)
			if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].idkunjungan != -1 {
				fmt.Println("Data ditemukan dengan keterangan berikut : ")
				fmt.Println("ID Kunjungan 				: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
				fmt.Println("Tanggal Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
				fmt.Println("Berat Badan				: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
				fmt.Println("Tinggi Badan				: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
				fmt.Println("Tekanan Darah Distolik 	: ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Println("Tekanan Darah Sistolik 	: ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				fmt.Println("Gula Darah					: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
			} else {
				fmt.Println("Berdasarkan waktu yang anda masukkan, data tidak ditemukan")
			}
		} else {
			fmt.Println("Berdasarkan waktu yang anda masukkan, data tidak ditemukan")
		}
	} else {
		fmt.Println("Anda membatalkan pengecekkan")
	}
}

func tambahdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses: dbk akun pasien bertambah, memasukkan pilihan pengecekan tanggal dilakukan hari ini atau tanggal lainnya. Lalu pengguna memasukkan data kunjungan
	// F.S : data kunjungan terekan dalam dbk akun pasien
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
			P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun = chunk_tahunkunjungan
		}
		fmt.Println("Masukkan Berat Badan (kg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].beratbadan)
		fmt.Println("Masukkan Tinggi Badan (cm) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tinggibadan)
		fmt.Println("Masukkan Tekanan Darah Distolik (mmHg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahdistolik)
		fmt.Println("Masukkan Tekanan Darah Sistolik (mmHg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahsistolik)
		fmt.Println("Masukkan Tingkat Gula Darah (mg/dL): ")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].guladarah)
		P[*sesi].datakesehatan[*dbk].idkunjungan = *dbk
		P[*sesi].datakesehatan[*dbk].iduser = *sesi
		P[*sesi].datakesehatan[*dbk].statusKunjungan = 1
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
		fmt.Println("Masukkan Berat Badan (kg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].beratbadan)
		fmt.Println("Masukkan Tinggi Badan (cm) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tinggibadan)
		fmt.Println("Masukkan Tekanan Darah Distolik (mmHg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahdistolik)
		fmt.Println("Masukkan Tekanan Darah Sistolik (mmHg) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tekanandarahsistolik)
		fmt.Println("Masukkan Tingkat Gula Darah (mg/dL) : ")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].guladarah)
		P[*sesi].datakesehatan[dbkunjungankosong].idkunjungan = dbkunjungankosong
		P[*sesi].datakesehatan[dbkunjungankosong].iduser = *sesi
		fmt.Println("Data sudah di rekam dengan ID Rekam : ", P[*sesi].datakesehatan[*dbk].idkunjungan)
		P[*sesi].datakesehatan[*dbk].statusKunjungan = 1
		fmt.Println("Untuk mengecek kesehatan silahkan ke menu cek kesehatan pada halaman utama")
	}

}

func ubahdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses: pengguna memilih metode cari id kunjungan (ID atau tanggal kunjungan)
	//			1. ID kunjungan
	//				pengguna memasukkan id kunjungan
	//F.S :
	var i, n, x int = 1, 0, 0
	var chunk_tanggal, chunk_bulan, chunk_tahun int
	var chunk_tanggalcari, chunk_bulancari, chunk_tahuncari int
	var chunk_id int
	var chunk_beratbadan, chunk_tinggibadan, chunk_guladarah, chunk_tekanandarahdistolik, chunk_tekanandarahsistolik float64
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
			fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
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
				if chunk_beratbadan > 0 {
					P[*sesi].datakesehatan[chunk_id].beratbadan = chunk_beratbadan
				}
				fmt.Println("Masukkan tinggi badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tinggi badan data lama : ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
				fmt.Scan(&chunk_tinggibadan)
				if chunk_tinggibadan > 0 {
					P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tinggibadan
				}
				fmt.Println("Masukkan tekanan darah distolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tekanan darah distolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Scan(&chunk_tekanandarahdistolik)
				if chunk_tekanandarahdistolik > 0 {
					P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik = chunk_tekanandarahdistolik
				}
				fmt.Println("Masukkan tekanan darah sistolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Tekanan darah sistolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				fmt.Scan(&chunk_tekanandarahsistolik)
				if chunk_tekanandarahsistolik > 0 {
					P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik = chunk_tekanandarahsistolik
				}
				fmt.Println("Masukkan gula darah yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
				fmt.Println("Gula darah data lama : ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
				fmt.Scan(&chunk_guladarah)
				if chunk_guladarah > 0 {
					P[*sesi].datakesehatan[chunk_id].guladarah = chunk_guladarah
				}
				fmt.Println("Data kunjungan sudah di rubah dengan isi :")
				fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
				fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
				fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
				fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
				fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
				fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
				fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
			} else {
				fmt.Println("Anda membatalkan aksi mengubah data")
			}
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	} else {
		fmt.Println("Masukkan tanggal kunjungan :")
		fmt.Scan(&chunk_tanggalcari, &chunk_bulancari, &chunk_tahuncari)
		for i <= *dbk {
			if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggalcari && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulancari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahuncari {
				arrpencaritanggal[n] = i
				n++
			}
			i++
		}
		if n != 0 {
			fmt.Println("Data yang ditemukan dengan preferensi tanggal, adalah:")
			fmt.Printf("%-20s %-15s %-15s %-15s %-25s %-25s\n", "TANGGAL KUNJUNGAN", "ID KUNJUNGAN", "BERAT BADAN (kg)", "TINGGI BADAN (m)", "TEKANAN DARAH DISTOLIK (mmHg)", "TEKANAN DARAH SISTOLIK (mmHg)")
			fmt.Println("==============================================================================================================")
			for x < n {
				datakesehatan := P[*sesi].datakesehatan[arrpencaritanggal[x]]
				fmt.Printf("%02d/%02d/%04d %-15d %-15.2f %-15.2f %-25.2f %-25.2f\n",
					datakesehatan.tanggalpengecekan.tanggal,
					datakesehatan.tanggalpengecekan.bulan,
					datakesehatan.tanggalpengecekan.tahun,
					datakesehatan.idkunjungan,
					datakesehatan.beratbadan,
					datakesehatan.tinggibadan,
					datakesehatan.tekanandarahdistolik,
					datakesehatan.tekanandarahsistolik)
				x++
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
				fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
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
					if chunk_beratbadan > 0 {
						P[*sesi].datakesehatan[chunk_id].beratbadan = chunk_beratbadan
					}
					fmt.Println("Masukkan tinggi badan yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tinggi badan data lama : ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
					fmt.Scan(&chunk_tinggibadan)
					if chunk_tinggibadan > 0 {
						P[*sesi].datakesehatan[chunk_id].tinggibadan = chunk_tinggibadan
					}
					fmt.Println("Masukkan tekanan darah distolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tekanan darah distolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
					fmt.Scan(&chunk_tekanandarahdistolik)
					if chunk_tekanandarahdistolik > 0 {
						P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik = chunk_tekanandarahdistolik
					}
					fmt.Println("Masukkan tekanan darah sistolik yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Tekanan darah sistolik data lama : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
					fmt.Scan(&chunk_tekanandarahsistolik)
					if chunk_tekanandarahsistolik > 0 {
						P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik = chunk_tekanandarahsistolik
					}
					fmt.Println("Masukkan gula darah yang ingin di ubah (isi 0 jika bagian bulan tak ingin di ubah)")
					fmt.Println("Gula darah data lama : ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
					fmt.Scan(&chunk_guladarah)
					if chunk_guladarah > 0 {
						P[*sesi].datakesehatan[chunk_id].guladarah = chunk_guladarah
					}
					fmt.Println("Data kunjungan sudah di rubah dengan isi :")
					fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
					fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
					fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
					fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].tinggibadan, "m")
					fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
					fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
					fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
				} else {
					fmt.Println("Anda membatalkan aksi mengubah data")
				}
			} else {
				fmt.Println("Data tidak ditemukan, silahkan coba lagi")
			}
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

func hapusdatakunjungan(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	//I.S. : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses : pengguna memilih metode pencarian indeks arraykunjungan, yaitu berdasarkan ID dan tanggal kunjungan. Jika data ditemukan, patribut status pada kunjungan akan diubah nilai menjadi 2 (tidak aktif). Jika tidak, pengguna akan keluar dari prosedur
	//F.S : datapasien masih ada, namun dengan status tidak aktif
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
		fmt.Println("Data sudah terhapus, untuk menambah atau mengubah gunakan fitur yang disediakan")
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan data berdasarkan waktu kunjungan yang ingin dihapus (DD/M/YYYY)")
		fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
		wantDeletedID = pencaritanggal(*P, *dbc, *dbk, *sesi, chunk_tanggal, chunk_bulan, chunk_tahun)

		P[*sesi].datakesehatan[wantDeletedID].statusKunjungan = 2
		fmt.Println("Data sudah terhapus, untuk menambah atau mengubah gunakan fitur yang disediakan")
	}
}

func cekkesehatan(P *atributpasien, dbk *int, dbc *int, sesi *int) {
	//I.S. : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	//proses : pengguna memasukkan tanggal, bulan, tahun kunjungan untuk mencari indeks array yang menyimpan data tsb. Jika data ditemukan, program akan menghitung BMI dan mengelompokkan kategori dan kelompok dari BMI dan tekanan darah dan menyimpannya dalam atribut diagnosaBMI dan diagnosatekanandarah.
	//F.S : atribut diagnosaBMI dan diagnosatekanandarah diperbarui
	var chunk_tanggalcari, chunk_bulancari, chunk_tahuncari int
	var chunk_konfirmasipenggunaan int
	var chunk_pemilihanmetodecari int
	var arrpencaritanggal [100]int
	var chunk_id int
	var n, i, x int

	fmt.Println("Berdasarkan apa properti data yang ingin di pakai (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Print("Masukkan ID Kunjungan yang ingin digunakan :")
		fmt.Scan(&chunk_id)
		if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].statusKunjungan != 2 {
			fmt.Println("Data ditemukan dengan keterangan berikut : ")
			fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
			fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
			fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
			fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
			fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
			fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
			fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
			fmt.Println("Apakah data yang ingin di gunakan benar (1 jika benar/ 0 jika salah) ?")
			fmt.Scan(&chunk_konfirmasipenggunaan)
			if chunk_konfirmasipenggunaan == 1 {
				P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil = P[*sesi].datakesehatan[chunk_id].beratbadan / ((P[*sesi].datakesehatan[chunk_id].tinggibadan / 100) * (P[*sesi].datakesehatan[chunk_id].tinggibadan / 100))
				if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil < 17 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Kurus"
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kekurangan berat badan tingkat rendah"
				} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 17 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 18.5 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Kurus"
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kekurangan berat badan tingkat sedang"
				} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 18.5 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 25 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Normal"
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Berat badan normal"
				} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 25 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 30 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Gemuk"
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kelebihan berat badan tingkat rendah"
				} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 30 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Gemuk"
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kelebihan berat badan tingkat tinggi"
				} else {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tidak Jelas"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tidak Jelas"
				}

				// Diagnosa Tekanan Darah
				if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 90 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 60 {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Rendah"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah rendah (Hipotensi)"
				} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 90 && P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 120 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 60 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 80 {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Normal"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah normal"
				} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 120 && P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 140 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 80 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 90 {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Pra-Hipertensi"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Pra-Hipertensi)"
				} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 140 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 90 {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tinggi"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Hipertensi)"
				} else {
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tidak Jelas"
					P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tidak Jelas"
				}
				fmt.Println("Hasil cek kesehatan anda dengan keterangan : ")
				fmt.Println("Kelompok : ", P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok)
				fmt.Println("Kategori : ", P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori)
			} else {
				fmt.Println("Anda membatalkan pengecekan kesehatan")
			}
		} else {
			fmt.Println("Anda memasukkan ID Kunjungan yang salah")
		}
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan tanggal kunjungan :")
		fmt.Scan(&chunk_tanggalcari, &chunk_bulancari, &chunk_tahuncari)
		for i <= *dbk {
			if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggalcari && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulancari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahuncari {
				arrpencaritanggal[n] = i
				n++
			}
			i++
		}
		if n != 0 {
			fmt.Println("Data yang ditemukan dengan preferensi tanggal, adalah:")
			fmt.Printf("%-20s %-15s %-15s %-15s %-25s %-25s\n", "TANGGAL KUNJUNGAN", "ID KUNJUNGAN", "BERAT BADAN (kg)", "TINGGI BADAN (m)", "TEKANAN DARAH DISTOLIK (mmHg)", "TEKANAN DARAH SISTOLIK (mmHg)")
			fmt.Println("==============================================================================================================")
			for x < n {
				datakesehatan := P[*sesi].datakesehatan[arrpencaritanggal[x]]
				fmt.Printf("%02d/%02d/%04d %-15d %-15.2f %-15.2f %-25.2f %-25.2f\n",
					datakesehatan.tanggalpengecekan.tanggal,
					datakesehatan.tanggalpengecekan.bulan,
					datakesehatan.tanggalpengecekan.tahun,
					datakesehatan.idkunjungan,
					datakesehatan.beratbadan,
					datakesehatan.tinggibadan,
					datakesehatan.tekanandarahdistolik,
					datakesehatan.tekanandarahsistolik)
				x++
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
				fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
				fmt.Println("Apakah data yang ingin di gunakan benar?")
				fmt.Scan(&chunk_konfirmasipenggunaan)
				if chunk_konfirmasipenggunaan == 1 {
					P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil = P[*sesi].datakesehatan[chunk_id].beratbadan / ((P[*sesi].datakesehatan[chunk_id].tinggibadan / 100) * (P[*sesi].datakesehatan[chunk_id].tinggibadan / 100))
					if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil < 17 {
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Kurus"
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kekurangan berat badan tingkat rendah"
					} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 17 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 18.5 {
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Kurus"
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kekurangan berat badan tingkat sedang"
					} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 18.5 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 25 {
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Normal"
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Berat badan normal"
					} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 25 && P[*sesi].datakesehatan[i].diagnosabmi.hasil < 30 {
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Gemuk"
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kelebihan berat badan tingkat rendah"
					} else if P[*sesi].datakesehatan[chunk_id].diagnosabmi.hasil >= 30 {
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kelompok = "Gemuk"
						P[*sesi].datakesehatan[chunk_id].diagnosabmi.kategori = "Kelebihan berat badan tingkat tinggi"
					} else {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tidak Jelas"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tidak Jelas"
					}

					// Diagnosa Tekanan Darah
					if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 90 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 60 {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Rendah"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah rendah (Hipotensi)"
					} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 90 && P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 120 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 60 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 80 {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Normal"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah normal"
					} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 120 && P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik < 140 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 80 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik < 90 {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Pra-Hipertensi"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Pra-Hipertensi)"
					} else if P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik >= 140 && P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik >= 90 {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tinggi"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tekanan darah tinggi (Hipertensi)"
					} else {
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok = "Tidak Jelas"
						P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori = "Tidak Jelas"
					}
					fmt.Println("Hasil cek kesehatan anda dengan keterangan : ")
					fmt.Println("Kelompok : ", P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kelompok)
					fmt.Println("Kategori : ", P[*sesi].datakesehatan[chunk_id].diagnosatekanandarah.kategori)
				} else {
					fmt.Println("Anda membatalkan pengecekan kesehatan")
				}
			} else {
				fmt.Println("Anda memasukkan ID Kunjungan yang salah")
			}
		} else {
			fmt.Println("Data tidak di temukan pada tanggal kunjungan yang anda")
		}
	} else {
		fmt.Println("Anda membatalkan pengecekan kesehatan")
	}
}

func cekguladarah(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang
	// proses : pengguna memilih metode pencarian indeks arraykunjungan, yaitu berdasarkan ID dan tanggal kunjungan. Jika data ditemukan, program akan mengoutput arraykunjungan untuk memastikan pengguna apa benar data kunjungan yang dimaksud akan diolah. Jika iya, program akan mengelompokkan dan mengkategorikan hasil tekanan darah
	//F.S : Jika berhasil, data akan disimpan di atribut diagnosaguladarah. Jika dibatalkan, maka program keluar dari prosedur
	// var chunk_tanggal, chunk_bulan, chunk_tahun int
	var chunk_tanggalcari, chunk_bulancari, chunk_tahuncari int
	var chunk_konfirmasipenggunaan int
	var chunk_pemilihanmetodecari int
	var arrpencaritanggal [100]int
	var chunk_id int
	var n, i, x int

	fmt.Println("Berdasarkan apa properti data yang ingin di pakai (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Print("Masukkan ID Kunjungan yang ingin digunakan :")
		fmt.Scan(&chunk_id)
		if P[*sesi].datakesehatan[chunk_id].idkunjungan != 0 && P[*sesi].datakesehatan[chunk_id].idkunjungan != -1 {
			fmt.Println("Data ditemukan dengan keterangan berikut : ")
			fmt.Println("ID Kunjungan 			: ", P[*sesi].datakesehatan[chunk_id].idkunjungan)
			fmt.Println("Tanggal Kunjungan 		: ", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tanggal, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.bulan, "/", P[*sesi].datakesehatan[chunk_id].tanggalpengecekan.tahun)
			fmt.Println("Berat Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "kg")
			fmt.Println("Tinggi Badan			: ", P[*sesi].datakesehatan[chunk_id].beratbadan, "m")
			fmt.Println("Tekanan Darah Distolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahdistolik, "mmHg")
			fmt.Println("Tekanan Darah Sistolik : ", P[*sesi].datakesehatan[chunk_id].tekanandarahsistolik, "mmHg")
			fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
			fmt.Println("Apakah data yang ingin di gunakan benar (1 jika benar/ 0 jika salah) ?")
			fmt.Scan(&chunk_konfirmasipenggunaan)
			if chunk_konfirmasipenggunaan == 1 {
				if P[*sesi].datakesehatan[chunk_id].guladarah < 35 {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Sangat Rendah"
				} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 35 && P[*sesi].datakesehatan[chunk_id].guladarah < 70 {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Rendah"
				} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 70 && P[*sesi].datakesehatan[chunk_id].guladarah < 100 {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Normal"
				} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 100 && P[*sesi].datakesehatan[chunk_id].guladarah < 125 {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Pra-diabetes"
				} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 125 && P[*sesi].datakesehatan[chunk_id].guladarah < 200 {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Diabetes"
				} else {
					P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Diabetes Tinggi"
				}
				fmt.Println("Hasil cek gula darah anda dengan keterangan : ")
				fmt.Println("Gula Darah : ", P[*sesi].datakesehatan[chunk_id].guladarah)
				fmt.Println("Kelompok   : ", P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok)
			} else {
				fmt.Println("Anda membatalkan pengecekan gula darah")
			}
		} else {
			fmt.Println("Anda memasukkan ID Kunjungan yang salah")
		}
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan tanggal kunjungan :")
		fmt.Scan(&chunk_tanggalcari, &chunk_bulancari, &chunk_tahuncari)
		for i <= *dbk {
			if P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggalcari && P[*sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulancari && P[*sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahuncari {
				arrpencaritanggal[n] = i
				n++
			}
			i++
		}
		if n != 0 {
			fmt.Println("Data yang ditemukan dengan preferensi tanggal, adalah:")
			fmt.Printf("%-20s %-15s %-15s %-15s %-25s %-25s\n", "TANGGAL KUNJUNGAN", "ID KUNJUNGAN", "BERAT BADAN (kg)", "TINGGI BADAN (m)", "TEKANAN DARAH DISTOLIK (mmHg)", "TEKANAN DARAH SISTOLIK (mmHg)")
			fmt.Println("==============================================================================================================")
			for x < n {
				datakesehatan := P[*sesi].datakesehatan[arrpencaritanggal[x]]
				fmt.Printf("%02d/%02d/%04d %-15d %-15.2f %-15.2f %-25.2f %-25.2f\n",
					datakesehatan.tanggalpengecekan.tanggal,
					datakesehatan.tanggalpengecekan.bulan,
					datakesehatan.tanggalpengecekan.tahun,
					datakesehatan.idkunjungan,
					datakesehatan.beratbadan,
					datakesehatan.tinggibadan,
					datakesehatan.tekanandarahdistolik,
					datakesehatan.tekanandarahsistolik)
				x++
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
				fmt.Println("Gula Darah				: ", P[*sesi].datakesehatan[chunk_id].guladarah, "mg/dL")
				fmt.Println("Apakah data yang ingin di gunakan benar?")
				fmt.Scan(&chunk_konfirmasipenggunaan)
				if chunk_konfirmasipenggunaan == 1 {
					if P[*sesi].datakesehatan[chunk_id].guladarah < 35 {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Sangat Rendah"
					} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 35 && P[*sesi].datakesehatan[chunk_id].guladarah < 70 {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Rendah"
					} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 70 && P[*sesi].datakesehatan[chunk_id].guladarah < 100 {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Normal"
					} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 100 && P[*sesi].datakesehatan[chunk_id].guladarah < 125 {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Pra-diabetes"
					} else if P[*sesi].datakesehatan[chunk_id].guladarah >= 125 && P[*sesi].datakesehatan[chunk_id].guladarah < 200 {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Diabetes"
					} else {
						P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok = "Diabetes Tinggi"
					}
					fmt.Println("Hasil cek gula darah anda dengan keterangan : ")
					fmt.Println("Gula Darah : ", P[*sesi].datakesehatan[chunk_id].guladarah)
					fmt.Println("Kelompok   : ", P[*sesi].datakesehatan[chunk_id].diagnosaguladarah.kelompok)
				}
			}
		} else {
			fmt.Println("Data tidak di temukan pada tanggal kunjungan yang anda")
		}
	} else {
		fmt.Println("Anda membatalkan pengecekan gula darah")
	}
}

func DisplayAkun(P atributpasien, dbc int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc dan arraykunjungan sebagai atribut datapasien, sesi sebagai indikator id pasien yang dipakai sekarang. Program hanya bisa dijalankan oleh pengguna berstatus admin
	// F.S : program akan mendisplay seluruh datapasien dan arraykunjungan pengguna
	var n, i, m int
	n = 0
	m = 0
	i = 1
	fmt.Println("=============================================")
	fmt.Println(i, "      LIST PASIEN RUMAH SAKIT TELKOM         ")
	fmt.Println("=============================================")
	for n < dbc {
		fmt.Println("-----------------------------------------")
		fmt.Println("ID :", P[m].idpasien)
		fmt.Println("Nama: ", P[m].namapasien.namadepan, P[n].namapasien.namabelakang)
		fmt.Println("usia: ", P[m].usia)
		fmt.Println("TTL :", P[m].tempatlahir, ",", P[m].tanggallahir.tanggal, P[m].tanggallahir.bulan, P[m].tanggallahir.tahun)
		fmt.Println("Gender :", P[m].gender)
		fmt.Println("Gender :", P[m].usia)
		for P[m].datakesehatan[i].statusKunjungan != 2 && P[m].datakesehatan[i].statusKunjungan != 0 {
			fmt.Println("id Kunjungan: ", P[m].datakesehatan[i].idkunjungan)
			fmt.Println("Tanggal Kunjungan: ", P[m].datakesehatan[i].tanggalpengecekan.tanggal, "/", P[m].datakesehatan[i].tanggalpengecekan.bulan, "/", P[m].datakesehatan[i].tanggalpengecekan.tahun)
			fmt.Println("===============[Hasil Diagnosa]===============")
			fmt.Println("Berat Badan: ", P[m].datakesehatan[i].beratbadan)
			fmt.Println("Tinggi Badan: ", P[m].datakesehatan[i].tinggibadan)
			fmt.Println("Hasil BMI: ", P[m].datakesehatan[i].diagnosabmi.hasil)
			fmt.Println("Termasuk dalam kelompok: ", P[m].datakesehatan[i].diagnosabmi.kelompok)
			fmt.Println("detail: ", P[m].datakesehatan[i].diagnosabmi.kategori)
			fmt.Println()
			fmt.Println("Tekanan Darah (Sitolik, Diastolik): ", P[m].datakesehatan[i].tekanandarahsistolik, P[m].datakesehatan[i].tekanandarahdistolik)
			fmt.Println("Hasil Uji tekanan darah: ", P[m].datakesehatan[i].diagnosatekanandarah.hasil)
			fmt.Println("Termasuk dalam kelompok: ", P[m].datakesehatan[i].diagnosatekanandarah.kelompok)
			fmt.Println("Detail: ", P[m].datakesehatan[i].diagnosatekanandarah.kategori)
			fmt.Println("Kadar Gula Darah: ", P[m].datakesehatan[i].guladarah)
			fmt.Println("Hasil Uji Gula darah: ", P[m].datakesehatan[i].diagnosaguladarah.hasil)
			fmt.Println("Termasuk dalam kelompok: ", P[m].datakesehatan[i].diagnosaguladarah.kelompok)
			fmt.Println("Detail: ", P[m].datakesehatan[1].diagnosaguladarah.kategori)
			fmt.Println("=============================================")
			i++
		}
		n++
		m++
	}
}

func EditDatabasePasien(P *atributpasien, dbc int) {
	//I.S. : terdefinisi array P bertipe atributpasien sebanyak dbc bilangan bulat
	//proses : Pengguna akan memasukkan ID pasien yang akan diedit. ID dicari dengan sequential search. Jika ID ditemukan, maka program mengoutput akun pengguna. Lalu, pengguna memilih atribuut akun yang akan diedit. Lalu pengguna memasukkan input yang baru untuk atribut akun tersebut.
	//F.S : Jika telah memenuhi syarat, makan akun pengguna berhasil di-update
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

func cariAkunPasien(P atributpasien, dbc int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc bilangan bulat
	//proses : Masukkan ID yang mau diccari. ID dicari menggunakan binary search.
	//F.S : Jika ID ditemukan, maka pengguna bisa memilih untuk menampilkan id dan masuk prosedur DisplayAkun_personal atau tidak menampilkan dan kembali ke prosedur
	var IDtarget, idx int
	var oldestRecord, BrandRecord, middle int
	var found bool

	idx = 0
	found = false
	oldestRecord = 1
	BrandRecord = dbc
	fmt.Println("------------------------------------------------")
	fmt.Println("Masukkan ID pasien yang ingin kamu cari:")
	fmt.Println("------------------------------------------------")
	fmt.Print("> ")
	fmt.Scan(&IDtarget)

	for oldestRecord <= BrandRecord && found == false {
		middle = (oldestRecord + BrandRecord) / 2
		if P[middle].idpasien == IDtarget && P[middle].status != 2 {
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
	}
}

// Subfungsi dan Tools

func pengecektanggal(check_tanggal, check_bulan, check_tahun int) bool {
	//Fungsi mengembalikan boolean yang berisi bila tanggal, bulan, dan tahun yang dimasukkan adalah valid
	var validasi bool

	if check_tahun > 0 {
		if (check_tahun%4 == 0 && check_tahun%100 != 0) || (check_tahun%400 == 0) {
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
	//fungsi mengembalikan integer yang bernilai indeks dari akun pasien yang memiliki nama depan, nama belankang, tanggal, bulan, dan tahun lahir yang sama. Fungsi mengembalikan -1 jika tidak ditemukan
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
	//fungsi mengembalikan integer bernilai indeks dari akun yang memiliki atribut status == 2 {tidak aktif}
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

func pencaritanggal(P atributpasien, dbc, dbk, sesi, chunk_tanggal, chunk_bulan, chunk_tahun int) int {
	//fungsi mengembalikan indeks arraykunjungan yang memiliki tanggal, bulan dan tahun kunjungan yang sama dari variabel chunk
	var i int = 1
	for i <= dbk {
		if P[sesi].datakesehatan[i].tanggalpengecekan.tanggal == chunk_tanggal && P[sesi].datakesehatan[i].tanggalpengecekan.bulan == chunk_bulan && P[sesi].datakesehatan[i].tanggalpengecekan.tahun == chunk_tahun {
			return i
		}
		i = i + 1
	}
	return -1
}

func pencaridbkunjungankosong(P atributpasien, dbc, dbk, sesi int) int {
	//fungsi mengembalikan integer bernilai indeks dari id kunjungan  yang memiliki atribut status == 2 {tidak aktif}
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

func sortSelectionDesc(P atributpasien, dbc int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc anggota
	//F.S : dibuat array baru bernama ghostArray sebanyak idxghost yang merupakan hasil sorting selection descended dari P dan akan diteruskan ke prosedur DisplayAkun
	var ghostArray atributpasien
	var pass, i, idx, idxghost int
	var temp datapasien
	ghostArray = P

	for j := 1; j <= dbc && P[j].status == 1; j++ {
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
	DisplayAkun(ghostArray, idxghost)
}

func sortInsertionAsx(P atributpasien, dbc int) {
	//I.S : terdefinisi array P bertipe atributpasien sebanyak dbc anggota
	//F.S : dibuat array baru bernama ghostArray sebanyak idxghost yang merupakan hasil sorting Insertion AScended dari P dan akan diteruskan ke prosedur DisplayAkun
	var pass, i, idxGoib int
	var temp datapasien
	var arrayGoib atributpasien

	for j := 1; j <= dbc && P[j].status == 1; j++ {
		arrayGoib[idxGoib] = P[j]
		idxGoib++
	}

	for pass = 1; pass < idxGoib; pass++ {
		i = pass
		temp = arrayGoib[pass]
		for i > 0 && temp.usia <= arrayGoib[i-1].usia {
			arrayGoib[i] = arrayGoib[i-1]
			i = i - 1
		}
		arrayGoib[i] = temp
	}
	DisplayAkun(arrayGoib, idxGoib)
}

func Sequentialsearch(P atributpasien, dbc int, x int) int {
	// I.S : terdefinisi array P bertipe atributpasien sebanyak dbc anggota dan x yang menandakan id yang akan dicari
	//F.S : prosedur mengembalikan indeks dari id akun pasien jika ditemukan. Jika tidak ditemukan, akan mengembalikan -1
	var stopper bool
	var i int
	var indeX int
	i = 1
	indeX = -1
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
