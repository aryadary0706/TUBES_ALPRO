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

//struct yang isinya hasil2 untuk BMI dan TD
type diagnosa struct {
	hasil float64
	kelompok string
	kategori string
}

// Atribut yang mencatat data pasien
type datapasien struct {
	idpasien					int
	namapasien					namalengkap
	usia						int
	tempatlahir					string
	tanggallahir				formattanggal
	gender						string
	password					string
	datakesehatan				arraykunjungan
}

// Atribut kunjungan untuk per user
type kunjungan struct {
	idkunjungan					int
	iduser						int
	tanggalpengecekan			formattanggal
	tekanandarahdistolik		float64
	tekanandarahsistolik		float64
	beratbadan, tinggibadan		float64
	diagnosabmi					diagnosa
	diagnosatekanandarah		diagnosa
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

// Program utama
func main() {
	var datapasien atributpasien
	var pilihan int
	var sesi int
	var dbk int

	sesi = -1
	pilihan = 1
	for pilihan >= 1 && pilihan <= 7 {
		interface0(&datapasien, sesi)
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
			for pilihan == 1 {
				interface03(&datapasien, sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					// cekkesehatan(&datapasien, &dbc, &sesi)
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
	
	var dbkosong int

	dbkosong = pencaridbkosong(*P, *dbc)
	if dbkosong == -1 {
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tangallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for i < *dbc {
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
		
		P[*dbc].idpasien = *dbc
		*sesi = *dbc
		fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
		fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
	} else {
		fmt.Println("Masukkan Nama Depan")
		fmt.Scan(&chunk_namadepan)
		fmt.Println("Masukkan Nama Belakang")
		fmt.Scan(&chunk_namabelakang)
		fmt.Println("Masukkan Tempat Lahir")
		fmt.Scan(&chunk_tempatlahir)
		fmt.Println("Masukkan Tanggal Lahir (DD/ M/ YYYY)")
		fmt.Scan(&chunk_tangallahir, &chunk_bulanlahir, &chunk_tahunlahir)
		for i < *dbc {
			if chunk_namadepan == P[i].namapasien.namadepan && chunk_namabelakang == P[i].namapasien.namabelakang && chunk_tempatlahir == P[i].tempatlahir && chunk_tangallahir == P[i].tanggallahir.tanggal && chunk_bulanlahir == P[i].tanggallahir.tanggal && chunk_tahunlahir == P[i].tanggallahir.tahun {
				fmt.Println("Maaf pengguna sudah terdaftar silahkan masuk ke menu masuk, dengan ID yang sudah terdaftar")
				break
			}
			i = i + 1
		}
		P[dbkosong].namapasien.namadepan = chunk_namadepan
		P[dbkosong].namapasien.namabelakang = chunk_namabelakang
		P[dbkosong].tempatlahir = chunk_tempatlahir
		P[dbkosong].tanggallahir.tanggal = chunk_tangallahir
		P[dbkosong].tanggallahir.bulan = chunk_bulanlahir
		P[dbkosong].tanggallahir.tahun = chunk_tahunlahir
		fmt.Println("Pilih Gender (Laki-Laki / Perempuan)")
		fmt.Println("L / l = Laki-laki")
		fmt.Println("P / p = Perempuan")
		fmt.Scan(&chunk_gender)
		if chunk_gender == "L" || chunk_gender == "l" {
			P[dbkosong].gender = "Laki-laki"
		} else if chunk_gender == "P" || chunk_gender == "p" {
			P[dbkosong].gender = "Perempuan"
		}
		fmt.Println("Masukkan Password untuk user yang di daftarkan")
		fmt.Scan(&P[dbkosong].password)
		
		P[dbkosong].idpasien = dbkosong
		*sesi = dbkosong
		fmt.Println("Pengguna sudah di daftarkan dengan ID :", *sesi)
		fmt.Println("Harap Ingat ID dan Password yang sudah terdaftar")
	}
}

func pencaridbkosong (P atributpasien, dbc int) int {
	var dbkosong int
	
	for dbkosong < dbc {
		if P[dbkosong].namapasien.namadepan == "-1" && P[dbkosong].namapasien.namabelakang == "-1" {
			return dbkosong
		}
		dbkosong = dbkosong + 1
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
	var i int = 0


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
	if pernyataanpenghapusan == 1 && *sesi != -1 && pernyataanpenghapusan == 1 && chunk_password == P[*sesi].password && chunk_password2 == P[*sesi].password && chunk_password == chunk_password2 {
		P[*sesi].namapasien.namadepan = "-1"
		P[*sesi].namapasien.namabelakang = "-1"
		P[*sesi].usia = -1
		P[*sesi].tempatlahir = "-1"
		P[*sesi].tanggallahir.tanggal = -1
		P[*sesi].tanggallahir.bulan = -1
		P[*sesi].tanggallahir.tahun = -1
		P[*sesi].gender = "-1"
		P[*sesi].password = "-1"
		for i < *dbk {
			if P[*sesi].datakesehatan[i].iduser == *sesi {
				P[*sesi].datakesehatan[i].tanggalpengecekan.tanggal = -1
				P[*sesi].datakesehatan[i].tanggalpengecekan.bulan = -1
				P[*sesi].datakesehatan[i].tanggalpengecekan.tahun = -1
				P[*sesi].datakesehatan[i].tekanandarahdistolik = -1
				P[*sesi].datakesehatan[i].tekanandarahsistolik = -1
				P[*sesi].datakesehatan[i].beratbadan = -1
				P[*sesi].datakesehatan[i].tinggibadan = -1
				P[*sesi].datakesehatan[i].diagnosabmi.hasil = -1
				P[*sesi].datakesehatan[i].diagnosabmi.kelompok = "-1"
				P[*sesi].datakesehatan[i].diagnosabmi.kategori = "-1"
				P[*sesi].datakesehatan[i].diagnosatekanandarah.hasil = -1
				P[*sesi].datakesehatan[i].diagnosatekanandarah.kelompok = "-1"
				P[*sesi].datakesehatan[i].diagnosatekanandarah.kategori = "-1"
			}
			i = i + 1
		}
		*sesi = -1
		fmt.Println("Data akun anda dan kunjungan sudah di hapus")
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

func pencaritanggal (P atributpasien, dbc, dbk, sesi, chunk_tanggal, chunk_bulan, chunk_tahun int) int {
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

	dbkunjungankosong = pencaridbkunjungankosong(*P, *dbc, *dbk)
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

func pencaridbkunjungankosong (P atributpasien, dbc, dbk int) int {
	var dbkunjungankosong int
	
	for dbkunjungankosong < dbc {
		if P[dbkunjungankosong].namapasien.namadepan == "-1" && P[dbkunjungankosong].namapasien.namabelakang == "-1" {
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

func hapusdatakunjungan(P *atributPasien, dbc *int, dbk *int, sesi *int) {
	var chunk_pemilihanmetodecari, chunk_idcari int
	var chunk_tanggal, chunk_bulan, chunk_tahun int = 0, 0, 0
	var i int = 0

	fmt.Println("Berdasarkan apa properti data yang ingin di hapus (ID Kunjungan/ Tanggal Kunjungan) ?")
	fmt.Println("1 : ID Kunjungan")
	fmt.Println("2 : Tanggal Kunjungan")
	fmt.Scan(&chunk_pemilihanmetodecari)
	if chunk_pemilihanmetodecari == 1 {
		fmt.Println("Masukkan ID Kunjungan yang ingin dihapus")
		fmt.Scan(&chunk_idcari)
		
	} else if chunk_pemilihanmetodecari == 2 {
		fmt.Println("Masukkan data berdasarkan waktu kunjungan yang ingin dihapus (DD/M/YYYY)")
		fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
		pencaritanggal(*P, *dbc, *dbk, *sesi, chunk_tanggal, chunk_bulan, chunk_tahun)
		
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
			case P[*sesi].datakesehatan[i].tekanandarahsistolik >= 120 && P[*sesi].datakesehatan[i].tekanandarahsistolik < 140 && P[*sesi].datakesehatan[i].tekanandarahdistolik >= 80 && P[*sesi].datakesehatan[i].tekanandarahdistolik  < 90:
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