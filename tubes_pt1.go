package main

import (
	"fmt"
	"math"
	"os"
)

type ttl struct {
	tanggal, bulan, tahun int
}

type tanggal struct {
	tanggal, bulan, tahun int
}

type nama struct {
	Firstname, lastname string
}

type atributPasien struct {
	id string
	Name nama
	umur int
	gender string
	TTL ttl
	beratBadan, tinggiBadan int
	TDS, TDD int
	Detakjantung int
	kondisi1, kondisi2, kondisi3 string
	hasilBMI int
	kunjungan tanggal
}

type Pasien[999]atributPasien

func main(){
	var pilihan, pilihan_2, category int
	var urut int
	var n int
	n = 0
	urut = 1

	for pilihan == 1 || pilihan == 2 || pilihan == 3 || pilihan == 0{
	fmt.Println("Kamu adalah pasien ke - ", urut)
	fmt.Println("--------------------------------------------")
	fmt.Println("1. Cek Kesehatan Pasien")
	fmt.Println("2. Tammpilkan Data Pasien")
	fmt.Println("3. Edit Data Pasien(MASIH DIBUAT)")
	fmt.Println("4. Close")
	fmt.Println("--------------------------------------------")
	fmt.Println("Pilih nomor di atas")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		cekKesehatan(&Pasien, &n)
	}else if pilihan == 2 {
		fmt.Println("1. Tampilkan semua data")
		fmt.Println("2. Tampilkan berdasarkan kategori yang diinginkan")
		fmt.Println("3. Tampilkan data terurut")
		fmt.Scan(&pilihan_2)
		if pilihan_2 == 1 {
			ShowAllData(Pasien, n)
		}else if pilihan_2 == 2 {
			// Menampilkan data yang sesuai dengan kategori spesifik yang dipilih di prosedur. Contoh: Umur 19, menampilkan pasien yang hanya umur 19
			fmt.Println("Pilih Kategori Data:")
			fmt.Println("1. Umur")
			fmt.Println("2. Kondisi")
			fmt.Println("3. Tahun berkunjung")
			fmt.Scan(&category)
			if category == 1 {
				ShowDataByUmur(Pasien, n)
			}else if category == 2 {
				ShowDataByKondisi(Pasien, n)
			}else if category == 3 {
				ShowDataByYear(Pasien, n)
			}
		}else if pilihan == 3 {
			ShowSortedData(Pasien, n)
		}
	}else if pilihan == 3 {
		fmt.Println("Pilih rekayasa data: 1. Edit, 2. Delete")
	}
}
}

func cekKesehatan(P Pasien, n int){
	/*
	Baca data pasien
	*/
	if n >= 1 && n <= 9 {
		P[n].id = "00" + string(n)
	}else if n >=10 && n <= 99 {
		P[n].id = "0" + string(n)
	}else{
		P[n].id = string(n)
	}
	fmt.Println("id: ", id)
	fmt.Println("Masukkan nama (terdiri dari nama depan dan nama belakang):")
	fmt.Scanln(&P[n].Firstname, &P[n].lastname)
	fmt.Println("Masukkan gender:")
	fmt.Println(&P[n].gender)
	fmt.Println("Masukkan TTL dan domisili:")
	fmt.Println("format:(tgl) (bln) (thn) (dom)")
	fmt.Scanln(&P[n].Tgl, &P[n].bln, &P[n].thn, &P[n].domisili)
	fmt.Println("Masukkan tanggal kedatangan:")
	fmt.Scanln(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)
	fmt.Println("Masukkan berat badan, tinggi badan:")
	fmt.Scanln(&P[n].beratBadan, &P[n].tinggiBadan)
	fmt.Println("Masukkan tekanan darah sistolik dan diastolik:")
	fmt.Scanln(&P[n].TDS, &P[n].TDS)
	fmt.Println("Masukkan Detak Jantung:")
	fmt.Scanln(&P[n].Detakjantung)

	P[n].umur = Hitungumur(P[n].Tgl, P[n].bln, P[n].thn, P[n].kunjungan.tanggal, P[n].kunjungan.bulan, P[n].kunjungan.tahun)

	P[n].kondisi1, P[n].kondisi2, P[n].kondisi3 = PeriksaKesehatan(P[n].beratBadan, P[n].tinggiBadan, P[n].hasilBMI, P[n].tekananDarah, P[n].TDS,P[n].TDD , P[n].umur)
}


func PeriksaKesehatan(BB, TB, TD, TDS, TDD, U int) (string, string, string){
	var BMI float64

	BMI = float64(BB) / math.Pow(float64(TB), 2.0)
	if BMI >= 30.0 {
		return "Obesitas"
	}else if BMI >= 25.0 && BMI < 30.0 {
		return "Overweight"
	}else if BMI >= 18.5 && BMI < 30{
		return "Normal"
	}else{
		return "Underweight"
	}


}

func hitungUmur(tl, bl, yl, tk, bk, yk int)int{
/*
Mengeluarkan nilai umur bertipe integer dari operasi dan percabangan variabel tl(tanggal lahir), bl(bulan lahir), yl(tahun lahir), tk(tahun kunjungan), bk (bulan kunjungan), yk (tahun kunjungan bertipe integer.)
*/ 
	var umur int
	umur = yk - yl
	if bl > bk {
		umur -= 1
	}else if tl < tk {
		umur -= 1
	}else{
		umur = umur
	}
	return umur
	}

func ShowAllData(P Pasien, n int){
	for i := 0; i < n; i++ {
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("Nama: ", P[i].Name.Firstname, P[i].Name.lastname)
		fmt.Println("Umur:", P[i].umur)
		fmt.Println("Kelamin/gender: ", P[i].gender)
		fmt.Println("Tanggal lahir:", P[n].TTL.tanggal, P[n].TTL.bulan, P[n].TTL.tahun)
		fmt.Println("Domisili:", P[i].domisili)
		fmt.Println("")
	}
}
