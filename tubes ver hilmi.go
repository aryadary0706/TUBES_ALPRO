package main

//import eketensi yang diperlukan
import (
	"fmt"
	"os"
)

// Konstanta minimal dari data
const NMAX int = 1000

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
	tempattanggallahir			formattanggal
	gender						string
	beratbadan, tinggibadan		float64
	
	kunjungan					tanggal
}

