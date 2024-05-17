package main

import (
	"fmt"
)

type tanggal struct {
	tanggal, bulan, tahun int
}

type nama struct {
	Firstname, lastname string
}

type atributPasien struct {
	id                      int
	Name                    nama
	umur                    int
	gender                  string
	TTL                     tanggal
	beratBadan, tinggiBadan float64
	TDS, TDD                int
	diagnosaBMI, diagnosaTD string
	hasilBMI                float64
	kunjungan               tanggal
}

const NMAX int = 1000

type Pasien [NMAX]atributPasien

func main() {
	//UI Interface
	var P Pasien
	var pilihan, pilihan_2, pilihan_3, category int
	var n int
	n = 0

	for pilihan == 1 || pilihan == 2 || pilihan == 0 || pilihan == 3 {
		fmt.Println("================================================")
		fmt.Println("         _       _ _         __  _       _ _        _")
		fmt.Println("  /\\ /\\ (_)_ __ (_) | __    /  \\(_) __ _(_) |_ __ _| |")
		fmt.Println(" / //_/ | | '_ \\| | |/ /   / /\\ / |/ _` | | __/ _` | |")
		fmt.Println("/ __ \\| | | | | | |   <   / /_//| | (_| | | || (_| | |")
		fmt.Println("\\/  \\/|_|_|_| |_|_|_|\\_\\ /___,/ |_|\\__, |_|\\__\\__,_|_|")
		fmt.Println("                                   |___/              ")
		fmt.Println("================================================")
		fmt.Println("No. Urut Pasien", n+1)
		fmt.Println("Silahkan Pilih menu berikut: ")
		fmt.Println("1. Cek Kesehatan")
		fmt.Println("2. Tampilkan Data Pasien")
		fmt.Println("3. Edit Data Pasien")
		fmt.Println()
		fmt.Println()
		fmt.Println("Program akan berhenti jika memilih angka selain di menu")
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			if n < NMAX {
				cekKesehatan(&P, n)
				n = n + 1
			} else {
				fmt.Print("Kapasitas Pasien Penuh")
			}
		} else if pilihan == 2 {
			fmt.Println("Pilih data berikut:")
			fmt.Println("1. Tampilkan semua data")
			fmt.Println("2. Tampilkan berdasarkan kategori")
			fmt.Println("3. Tampilkan data dari yang terkini")
			fmt.Scan(&pilihan_2)
			if pilihan_2 == 1 {
				ShowData(P, n)
			} else if pilihan_2 == 2 {
				//hanya menampilkan data yang sesuai dengan kategori spesifik dari input
				fmt.Println()
				fmt.Println("Pilih Kategori Data:")
				fmt.Println("1. Umur")
				fmt.Println("2. gender")
				fmt.Println("3. Tahun berkunjung")
				fmt.Print("Masukkan angka: ")
				fmt.Scan(&category)
				for category != 1 && category != 2 && category != 3 {
					fmt.Println("Masukan tidak valid")
					fmt.Scan(&category)
				}
				if category == 1 {
					MakeDataBasedAge(P, n)
				} else if category == 2 {
					MakeDataBasedGender(P, n)
				} else if category == 3 {
					MakeDataBasedYear(P, n)
				}
			} else if pilihan_2 == 3 {
				MakeSortedData(P, n)
			}
		}else if pilihan == 3 {
			fmt.Println("Pilih jenis rekayasa data: ")
			fmt.Println("1. Edit Data Pasien")
			fmt.Println("2. Hapus data Pasien")
			fmt.Print("Masukkan angka: ")
			fmt.Scan(&pilihan_3)
			for pilihan_3 != 1 && pilihan != 2 && pilihan != 3 {
				fmt.Println("Masukan tidak valid")
				fmt.Scan(&pilihan_3)
			}
			if pilihan_3 == 1 {
				EditData(&P, n)
			}else if pilihan_3 == 2 {
				DeleteData(&P, &n)
			}
		}
	}
	fmt.Println("Program sudah selesai")
}

func cekKesehatan(P *Pasien, n int) {
	/*
		Baca data pasien
	*/
	fmt.Println("Masukkan nama (terdiri dari nama depan dan nama belakang):")
	fmt.Print("Nama Depan: ")
	fmt.Scan(&P[n].Name.Firstname)
	fmt.Print("Nama Belakang: ")
	fmt.Scan(&P[n].Name.lastname)
	
	fmt.Println("Masukkan gender(Male/Female):")
	fmt.Scan(&P[n].gender)

	for P[n].gender != "Male" && P[n].gender != "Female" {
		fmt.Println("Input tidak terbaca. Masukkan kembali:")
		fmt.Scan(&P[n].gender)
	}

	fmt.Println("Masukkan Tanggal Lahir:")
	fmt.Println("(format: DD MM YYYY)")
	fmt.Scan(&P[n].TTL.tanggal, &P[n].TTL.bulan, &P[n].TTL.tahun)

	if !(VerifDate(P[n].TTL)) || P[n].TTL.tanggal < P[n-1].TTL.tanggal && P[n].TTL.bulan < P[n-1].TTL.bulan && P[n].TTL.tahun < P[n-1].TTL.tahun{
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&P[n].TTL.tanggal, &P[n].TTL.bulan, &P[n].TTL.tahun)
	}

	fmt.Println("Masukkan tanggal kedatangan:")
	fmt.Println("format: DD MM YYYY")
	fmt.Scan(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)

	if !(VerifDate(P[n].kunjungan)) || P[n].kunjungan.tanggal < P[n-1].kunjungan.tanggal && P[n].kunjungan.bulan < P[n-1].kunjungan.bulan && P[n].kunjungan.tahun < P[n-1].kunjungan.tahun{
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)
	}

	fmt.Println("Masukkan berat badan, tinggi badan:")
	fmt.Print("Berat Badan (kg): ")
	fmt.Scan(&P[n].beratBadan)
	fmt.Print("Tinggi badan (m): ")
	fmt.Scan(&P[n].tinggiBadan)

	fmt.Println("Masukkan tekanan darah sistolik dan diastolik:")
	fmt.Print("Tekanan darah Sistolik (mmHg): ")
	fmt.Scan(&P[n].TDS)
	fmt.Print("Tekanan darah Sistolik (mmHg): ")
	fmt.Scan(&P[n].TDD)

	BuatId(P, n)

	P[n].umur = hitungUmur(P[n].TTL.tanggal, P[n].TTL.bulan, P[n].TTL.tahun, P[n].kunjungan.tanggal, P[n].kunjungan.bulan, P[n].kunjungan.tahun)

	Diagnosa(P, n)
}

func BuatId(P *Pasien, n int) {
	var date int
	date = n
	if n > 0 {
		if P[n].kunjungan.tahun != P[n-1].kunjungan.tahun || P[n].kunjungan.bulan != P[n-1].kunjungan.bulan || P[n].kunjungan.tanggal != P[n-1].kunjungan.tanggal {
			date = 0
		}
	}
	P[n].id = (P[n].kunjungan.tahun%100) * 100000000 + P[n].kunjungan.bulan * 1000000 + P[n].kunjungan.tanggal * 10000 + (date+1)
}

func VerifDate(P Pasien, n int) bool{
	if (P[n].TTL.tanggal >= 1 && P[n].TTL.tanggal <= 31) && (P[n].TTL.bulan >= 1 && P[n].TTL.bulan <= 12) {
		return true
	}else{
		return false
	}
}

func Diagnosa(P *Pasien, n int) {
	var BMI float64
	BMI = (P[n].beratBadan / (P[n].tinggiBadan * P[n].tinggiBadan)) * 10000
	if BMI >= 30 {
		P[n].diagnosaBMI = "Obesitas"
	} else if BMI < 30 && BMI >= 23 {
		P[n].diagnosaBMI = "Berat badan berlebih"
	} else if BMI < 23 && BMI >= 18.5 {
		P[n].diagnosaBMI = "Berat badan Ideal"
	} else if BMI < 18.5 {
		P[n].diagnosaBMI = "Berat badan kurang"
	}
	P[n].hasilBMI = BMI

	if P[n].TDS > 180 && P[n].TDD > 120 {
		P[n].diagnosaTD = "Krisis Hipertensi"
	} else if P[n].TDS > 140 && P[n].TDD > 90 {
		P[n].diagnosaTD = "Hipertensi Derajat 2"
	} else if (P[n].TDS >= 130 && P[n].TDS < 140) && (P[n].TDD >= 80 && P[n].TDD < 90) {
		P[n].diagnosaTD = "Hipertensi Derajat 1"
	} else if (P[n].TDS >= 120 && P[n].TDS < 130) && (P[n].TDD < 80) {
		P[n].diagnosaTD = "Prahipertensi"
	} else if (P[n].TDS < 120 && P[n].TDS >= 90) && (P[n].TDD < 80 && P[n].TDD >= 60) {
		P[n].diagnosaTD = "Normal"
	} else if P[n].TDS < 90 && P[n].TDD < 60 {
		P[n].diagnosaTD = "Hipotensi"
	} else {
		P[n].diagnosaTD = "Data tidak valid"
	}
}

func hitungUmur(tl, bl, yl, tk, bk, yk int) int {
	/*
		Mengeluarkan nilai umur bertipe integer dari operasi dan percabangan variabel tl(tanggal lahir), bl(bulan lahir), yl(tahun lahir), tk(tahun kunjungan), bk (bulan kunjungan), yk (tahun kunjungan bertipe integer.)
	*/
	var umur int
	umur = yk - yl
	if bl > bk {
		umur -= 1
	} else if tl < tk {
		umur -= 1
	} else {
		umur = umur
	}
	return umur
}

func ShowData(P Pasien, n int) {
	if n > 0 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		for i := 0; i < n; i++ {
			
			fmt.Println("-----------------------------------------------------")
			fmt.Printf("No. ID:  %d\n", P[i].id)
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Nama:", P[i].Name.Firstname, P[i].Name.lastname)
			fmt.Println("Umur:", P[i].umur)
			fmt.Println("Gender:", P[i].gender)
			fmt.Println("Tanggal Lahir:", P[i].TTL.tanggal, "/", P[i].TTL.bulan, "/", P[i].TTL.tahun)
			fmt.Println("Tanggal Kunjungan:", P[i].kunjungan.tanggal, "/", P[i].kunjungan.bulan, "/", P[i].kunjungan.tahun)
			fmt.Println("Berat Badan:", P[i].beratBadan)
			fmt.Println("Tinggi Badan:", P[i].tinggiBadan)
			fmt.Println("Diagnosa BMI:", P[i].diagnosaBMI)
			fmt.Println("Diagnosa Tekanan Darah:", P[i].diagnosaTD)
			fmt.Printf("Hasil BMI: %.2f \n", P[i].hasilBMI)
			fmt.Println("----------------------------------------------------")
		}
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func MakeDataBasedAge(P Pasien, n int) {
	var B Pasien
	var m int
	m = 0
	var age int
	fmt.Println("Masukkan Umur:")
	fmt.Scan(&age)
	for i := 0; i < n; i++ {
		if P[i].umur == age {
			B[m] = P[i]
			m++
		}
	}
	ShowData(B, m)
}

func MakeDataBasedGender(P Pasien, n int) {
	var C Pasien
	var m int
	m = 0
	var gender string
	fmt.Println("Male/Female?")
	fmt.Scan(&gender)
	for i := 0; i < n; i++ {
		if P[i].gender == gender {
			C[m] = P[i]
			m++
		}
	}
	ShowData(C, m)
}

func MakeDataBasedYear(P Pasien, n int) {
	var D Pasien
	var m int
	m = 0
	var year int	
	fmt.Println("Masukkan Tahun:")
	fmt.Scan(&year)
	for i := 0; i < n; i++ {
		if P[i].kunjungan.tanggal == year {
			D[m] = P[i]
			m++
		}
	}
	ShowData(D, m)
}

func MakeSortedData(P Pasien, n int) {
	var E Pasien
	var m int
	m = 0
	for i := 0; i < n; i++ {
		E[m] = P[n-1-i]
		m++
	}
	ShowData(E, m)
}


//fungsi ini masih BELUM DICOBA (BELUM DIRUN)


func EditData(P *Pasien, n int) {
	var id int
	fmt.Println("Masukkan ID pasien yang ingin diedit:")
	fmt.Scan(&id)

	found := false // Variabel flag untuk menandai apakah ID ditemukan atau tidak

	for i := 0; i < n; i++ {
		if P[i].id == id {
			// Menampilkan dan memperbarui data pasien
			found = true
			fmt.Println("Data pasien yang akan diedit:")
			fmt.Println("ID:", P[i].id)
			fmt.Println("Nama:", P[i].Name.Firstname, P[i].Name.lastname)
			fmt.Println("Umur:", P[i].umur)
			fmt.Println("Gender:", P[i].gender)
			fmt.Println("Tanggal Lahir:", P[i].TTL.tanggal, "/", P[i].TTL.bulan, "/", P[i].TTL.tahun)
			fmt.Println("Tanggal Kunjungan:", P[i].kunjungan.tanggal, "/", P[i].kunjungan.bulan, "/", P[i].kunjungan.tahun)
			fmt.Println("Berat Badan:", P[i].beratBadan)
			fmt.Println("Tinggi Badan:", P[i].tinggiBadan)
			fmt.Println("Diagnosa BMI:", P[i].diagnosaBMI)
			fmt.Println("Diagnosa Tekanan Darah:", P[i].diagnosaTD)
			fmt.Printf("Hasil BMI: %.2f \n", P[i].hasilBMI)
			fmt.Println("---------------------")

			// Memperbarui data pasien
			var choice int
			fmt.Println("Apa yang ingin Anda edit?")
			fmt.Println("1. Nama")
			fmt.Println("2. Gender")
			fmt.Println("3. Tanggal Lahir")
			fmt.Println("4. Tanggal Kunjungan")
			fmt.Println("5. Berat Badan")
			fmt.Println("6. Tinggi Badan")
	
			fmt.Scan(&choice)

			if choice == 1 {
				fmt.Println("Masukkan nama baru (terdiri dari nama depan dan nama belakang):")
				fmt.Scan(&P[i].Name.Firstname, &P[i].Name.lastname)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 2 {
				fmt.Println("Masukkan gender baru:")
				fmt.Scan(&P[i].gender)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 3 {
				fmt.Println("Masukkan Tanggal Lahir baru (format: DD MM YYYY):")
				fmt.Scan(&P[i].TTL.tanggal, &P[i].TTL.bulan, &P[i].TTL.tahun)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 4 {
				fmt.Println("Masukkan tanggal kunjungan baru (format: DD MM YYYY):")
				fmt.Scan(&P[i].kunjungan.tanggal, &P[i].kunjungan.bulan, &P[i].kunjungan.tahun)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 5 {
				fmt.Println("Masukkan berat badan baru:")
				fmt.Scan(&P[i].beratBadan)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 6 {
				fmt.Println("Masukkan tinggi badan baru:")
				fmt.Scan(&P[i].tinggiBadan)
				fmt.Println("Data berhasil diupdate!")
			}else {
				fmt.Println("Pilihan tidak valid")
			}
		}
	}

	if !found {
		fmt.Println("ID pasien tidak ditemukan.")
	}
}

func DeleteData(P *Pasien, n *int) {
	var id int
	fmt.Println("Masukkan ID pasien yang ingin dihapus:")
	fmt.Scan(&id)

	found := false // Variabel flag untuk menandai apakah ID ditemukan atau tidak

	for i := 0; i < *n; i++ {
		if P[i].id == id {
			found = true
			// Geser data ke kiri untuk menutup celah
			for j := i; j < *n-1; j++ {
				P[j] = P[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus.")
			break // Keluar dari loop setelah data ditemukan dan dihapus
		}
	}

	if !found {
		fmt.Println("ID pasien tidak ditemukan.")
	}
}
