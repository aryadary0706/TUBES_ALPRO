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

type atributPasien struct {
	id string
	Firstname, lastname string
	umur int
	TTL ttl
	domisili string
	beratBadan, tinggiBadan, tekananDarah int
	kondisi bool
	hasilBMI int
	kunjungan tanggal
}

type Pasien[999]atributPasien

func main(){
	var responden Pasien
	var pilihan, pilihan_2, category int
	var urut int
	urut = 0

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
		fmt.Scan(&pilihan_2)
		if pilihan_2 == 1 {
			ShowAllData(Pasien, n)
		}else if pilihan_2 == 2 {
			// Menampilkan data yang sesuai dengan kategori spesifik yang dipilih di prosedur. Contoh: Umur 19, menampilkan pasien yang hanya umur 19
			fmt.Println("Pilih Kategori Data:")
			fmt.Println("1. Umur")
			fmt.Println("2. Domisili")
			fmt.Println("3. Kondisi")
			fmt.Println("4. Tahun berkunjung")
			fmt.Scan(&category)
			if category == 1 {
				ShowDataByUmur(Pasien, n)
			}else if category == 2 {
				ShowDataByDomisili(Pasien, n)
			}else if category == 3 {
				ShowDataByKondisi(Pasien, n)
			}
		}
	}else if pilihan == 3 {
		fmt.Println("Pilih rekayasa data: 1. Edit, 2. Delete")
	}else{
		//Gua ga tau cara close program :(
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
	fmt.Println("Masukkan umur, TTL, domisili:")
	fmt.Println("format: (umur) (tgl) (bln) (thn) (dom)")
	fmt.Scanln(&P[n].umur, &P[n].TTL.tanggal, &P[n].TTL.bulan, &P[n].TTL.tahun, &P[n].domisili)
	fmt.Println("Masukkan tanggal kedatangan:")
	fmt.Scanln(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)

}

func HitungBMIdanKondisi
