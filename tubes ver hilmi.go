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
	tekanandarah				float64
	beratbadan, tinggibadan		float64
	bmi							float64
}

// Atribut banyak data pasien sesuai banyak limit yang ditentukan di NMAX
type atributpasien [NMAX]datapasien
type arraykunjungan [NMAX]kunjungan 

// Interface yang akan di tampilkan
func interface0(sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("                SELAMAT DATANG DI               ")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	fmt.Println("Saat ini, Terdaftar sebagai :", sesilogin)
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
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

func interface01(sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	fmt.Println("Saat ini, terdaftar sebagai :", sesilogin)
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
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

func interface02(sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	fmt.Println("Saat ini, Terdaftar sebagai :", sesilogin)
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
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

func interface03(sesilogin int) {
	fmt.Println("================================================")
	fmt.Println("            APLIKASI PENGECEK KESEHATAN         ")
	fmt.Println("          RUMAH SAKIT TELKOM UNIVERTSITY        ")
	fmt.Println("================================================")
	fmt.Println("Saat ini, Terdaftar sebagai :", sesilogin)
	if sesilogin == -1 {
		fmt.Println("Anda belum masuk (login), silahkan login terlebih dahulu di menu Isi atau Pilih data Pasien")
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
	for pilihan >= 1 && pilihan <= 5 {
		interface0(sesi)
		fmt.Scan(&pilihan)
		
		if pilihan == 1 {	
			for pilihan >= 1 && pilihan <= 6 {
				interface01(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					daftarpengguna(&datapasien, &dbc, &sesi)
					fmt.Print(sesi)
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
				interface02(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					// lihatdata(&datapasien, &dbc, &sesi)
				} else if pilihan == 2 {
					tambahdata(&datapasien, &dbc, &dbk, &sesi)
				} else if pilihan == 3 {
					// editdata(&datapasien, &dbc, &sesi)
				} else if pilihan == 4 {
					// hapusdata(&datapasien, &dbc, &sesi)
				}
			}
			//data(&datapasien, &dbc ,&pilihan, &sesilogin)
		} else if pilihan == 3 {
			for pilihan == 1 {
				interface03(sesi)
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					// cekkesehatan(&datapasien, &dbc, &sesi)
				}
			}
			//kesehatan(&datapasien, &dbc ,&pilihan, &sesilogin)
		} 
	}
}

// Program per-layanan

func daftarpengguna(P *atributpasien, dbc *int, sesi *int) {
	var chunk_gender string
	var chunk_namadepan, chunk_namabelakang, chunk_tempatlahir string
	var chunk_tangallahir, chunk_bulanlahir, chunk_tahunlahir int
	var i int
	
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
			fmt.Print("Maaf pengguna sudah terdaftar silahkan masuk ke menu masuk, dengan ID yang sudah terdaftar")
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
	fmt.Print("Apakah data pasien yang akan di rubah adalah ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Print("Masukkan angka 1 / 0 (1 untuk BENAR, 0 untuk SALAH)")
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
		fmt.Println("Nama Depan           :", P[*sesi].namapasien.namadepan)
		fmt.Println("Nama Belakang        :", P[*sesi].namapasien.namabelakang)
		fmt.Println("Tempat Lahir         :", P[*sesi].tempatlahir)
		fmt.Println("Tanggal/ Tahun Lahir :", P[*sesi].tanggallahir.tanggal, P[*sesi].tanggallahir.bulan, P[*sesi].tanggallahir.tahun)
		fmt.Println("Gender				  :", P[*sesi].gender)	
	} else {
		if sesi == -1 {
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
	fmt.Print("Apakah data pasien yang akan di rubah adalah ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
	fmt.Print("Masukkan angka 1 / 0 (1 untuk BENAR, 0 untuk SALAH)")
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
	fmt.Print("Akun yang akan dihapus adalah akun dengan ID Pasien :", P[*sesi].idpasien, "dengan atas nama", P[*sesi].namapasien.namadepan, P[*sesi].namapasien.namabelakang, "?")
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
				P[*sesi].datakesehatan[i].tekanandarah = -1
				P[*sesi].datakesehatan[i].beratbadan = -1
				P[*sesi].datakesehatan[i].tinggibadan = -1
				P[*sesi].datakesehatan[i].bmi = -1
			}
			i = i + 1
		}
		*sesi = -1
		fmt.Println("Data akun anda dan kunjungan sudah di hapus")
	}

}

func keluar(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	*sesi = -1
	fmt.Println("Anda sudah keluar, terimakasih sudah menggunakan aplikasi kami")
}

func lihatdata(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var chunk_tanggal, chunk_bulan, chunk_tahun, chunk_search int
	var i int
	var s1, s2, s3, s int
	
	fmt.Println("Pastikan Anda login dengan ID yang sesuai")
	fmt.Println("Masukkan tanggal pengecekan kesehatan (DD/M/YYYY) : ")
	fmt.Scan(&chunk_tanggal, &chunk_bulan, &chunk_tahun)
	for i <= *dbk && i < NMAX{
		if P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal == chunk_tanggal && P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan == chunk_bulan && P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun == chunk_tahun {
			if s1 == 0 {
				s1 = i 
			} else if s1 > 0 {
				s2 = i
			} else if s2 > 0 {
				s3 = i
			}
			s = s + 1
		}
		i = i + 1
	}
	fmt.Println("Data yang di temukan dengan ID kunjungan :")
	fmt.Println("ID    TANGGAL      BULAN      TAHUN")
	for x := i ; x < s ; i++ {
		if x == 1 {
			fmt.Printf("%d    %d      %d      %d\n", s1, P[*sesi].datakesehatan[s1].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s1].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s1].tanggalpengecekan.tahun)
		} else if x == 2 && s2 > 0{
			fmt.Printf("%d    %d      %d      %d\n", s2, P[*sesi].datakesehatan[s2].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s2].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s2].tanggalpengecekan.tahun)
		} else if x == 3 && s3 > 0 {
			fmt.Printf("%d    %d      %d      %d\n", s3, P[*sesi].datakesehatan[s3].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s3].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s3].tanggalpengecekan.tahun)
		}
	}
	fmt.Println("Pilih data yang tersedia dengan menuliskan ID Kunjungan yang sesuai pada tanggal :")
	fmt.Scan(&chunk_search)
	if chunk_search == s1 {
		fmt.Println(P[*sesi].datakesehatan[s1].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s1].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s1].tanggalpengecekan.tahun)
		fmt.Printf("Berat Badan : %d", P[*sesi].datakesehatan[s1].beratbadan)
		fmt.Printf("Tinggi Badan : %d", P[*sesi].datakesehatan[s1].tinggibadan)
		fmt.Printf("BMI : %d", P[*sesi].datakesehatan[s1].bmi)
		fmt.Printf("Tekanan Darah : %f", P[*sesi].datakesehatan[s1].tekanandarah)
	} else if chunk_search == s2{
		fmt.Println(P[*sesi].datakesehatan[s2].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s2].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s2].tanggalpengecekan.tahun)
		fmt.Printf("Berat Badan : %d", P[*sesi].datakesehatan[s2].beratbadan)
		fmt.Printf("Tinggi Badan : %d", P[*sesi].datakesehatan[s2].tinggibadan)
		fmt.Printf("BMI : %d", P[*sesi].datakesehatan[s2].bmi)
		fmt.Printf("Tekanan Darah : %d", P[*sesi].datakesehatan[s2].tekanandarah)
	} else if chunk_search == s3 {
		fmt.Println(P[*sesi].datakesehatan[s3].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[s3].tanggalpengecekan.bulan, P[*sesi].datakesehatan[s3].tanggalpengecekan.tahun)
		fmt.Printf("Berat Badan : %d", P[*sesi].datakesehatan[s3].beratbadan)
		fmt.Printf("Tinggi Badan : %d", P[*sesi].datakesehatan[s3].tinggibadan)
		fmt.Printf("BMI : %d", P[*sesi].datakesehatan[s3].bmi)
		fmt.Printf("Tekanan Darah : %d", P[*sesi].datakesehatan[s3].tekanandarah)
	}
}
		
func tambahdata(P *atributpasien, dbc *int, dbk *int, sesi *int) {
	var c int
	*dbk = *dbk + 1
	
	fmt.Print("Apakah data pengecekan ini dilakukan pada hari ini?")
	fmt.Print("1 : Ya")
	fmt.Print("2 : Tidak")
	fmt.Print("Masukkan Data :")
	fmt.Scan(&c)
	if c == 1 {
		P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal = currentTime.Day()
		P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan = int(currentTime.Month())
		P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun = currentTime.Year()
	} else {
		fmt.Print("Masukkan Tanggal Pengecekkan (DD/M/YYYY) :")
		fmt.Scan(&P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tanggal, P[*sesi].datakesehatan[*dbk].tanggalpengecekan.bulan, &P[*sesi].datakesehatan[*dbk].tanggalpengecekan.tahun)
	}
	fmt.Print("Masukkan Berat Badan :")
	fmt.Scan(P[*sesi].datakesehatan[*dbk].beratbadan)
	fmt.Print("Masukkan Tinggi Badan :")
	fmt.Scan(P[*sesi].datakesehatan[*dbk].tinggibadan)
	fmt.Print("Masukkan Tekanan Darah :")
	fmt.Scan(P[*sesi].datakesehatan[*dbk].tekanandarah)
	P[*sesi].datakesehatan[*dbk].bmi = P[*sesi].datakesehatan[*dbk].beratbadan / P[*sesi].datakesehatan[*dbk].tinggibadan * P[*sesi].datakesehatan[*dbk].tinggibadan
	P[*sesi].datakesehatan[*dbk].idkunjungan = *dbk
	P[*sesi].datakesehatan[*dbk].iduser = *sesi 
	fmt.Print("Data sudah di rekam dengan ID Rekam : ", P[*sesi].datakesehatan[*dbk].idkunjungan)
	fmt.Print("Untuk mengecek kesehatan silahkan ke menu cek kesehatan pada halaman utama")
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

	// func hapusdata(P *atributPasien, dbc *int, sesi *int) {

	// }

	// func cekkesehatan {
		
	// }