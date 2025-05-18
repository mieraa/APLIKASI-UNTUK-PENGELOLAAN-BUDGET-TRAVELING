package main

import (
	"fmt"
	"time"
)

const NMAX int = 20

type tabInt [NMAX]float64
type tabString [NMAX]string

var peng tabInt
var ket tabString
var npeng int

var masukan tabInt
var sumber tabString
var nmasuk int

var tglTujuan time.Time

func main() {
	hitungMundur()
	menu()
}

func hitungMundur() {
	var hari, bulan, tahun int
	fmt.Println("Masukkan tanggal keberangkatan:")
	fmt.Println("Masukkan Hari, Bulan, Tahun dengan angka ya..")
	fmt.Print("Hari: ")
	fmt.Scan(&hari)
	fmt.Print("Bulan: ")
	fmt.Scan(&bulan)
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	tglTujuan = time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.Local)
	now := time.Now()
	sel := tglTujuan.Sub(now)

	if sel < 0 {
		fmt.Println("Tanggal keberangkatan sudah lewat.")
	} else {
		hariTersisa := int(sel.Hours() / 24)
		fmt.Printf("Menuju keberangkatan: %d hari lagi\n\n", hariTersisa)
	}
}

func menu() {
	var pilih int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Rencana Tabungan Perjalanan")
		fmt.Println("2. Edit Pengeluaran")
		fmt.Println("3. Tambah Pemasukan")
		fmt.Println("4. Edit Pemasukan")
		fmt.Println("5. Laporan")
		fmt.Println("6. Rekomendasi Tabungan Harian")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahPengeluaran(&peng, &ket, &npeng)
		case 2:
			editPengeluaran(&peng, &ket, npeng)
		case 3:
			tambahPemasukan(&masukan, &sumber, &nmasuk)
		case 4:
			editPemasukan(&masukan, &sumber, nmasuk)
		case 5:
			laporan(peng, ket, npeng, masukan, sumber, nmasuk)
		case 6:
			rekomendasiTabungan(peng, npeng)
		case 0:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahPengeluaran(p *tabInt, ket *tabString, n *int) {
	fmt.Println("Masukkan jumlah pengeluaran untuk masing-masing kategori:")

	fmt.Print("Transportasi = ")
	fmt.Scan(&(*p)[*n])
	(*ket)[*n] = "Transportasi"
	*n = *n + 1

	fmt.Print("Makan = ")
	fmt.Scan(&(*p)[*n])
	(*ket)[*n] = "Makan"
	*n = *n + 1

	fmt.Print("Hotel = ")
	fmt.Scan(&(*p)[*n])
	(*ket)[*n] = "Hotel"
	*n = *n + 1

	fmt.Print("Tiket Acara = ")
	fmt.Scan(&(*p)[*n])
	(*ket)[*n] = "Tiket Acara"
	*n = *n + 1

	fmt.Print("Deskripsi Lain = ")
	fmt.Scan(&(*p)[*n])
	(*ket)[*n] = "Lain-lain"
	*n = *n + 1

	total := 0.0
	for i := 0; i < *n; i++ {
		total += (*p)[i]
	}
	fmt.Printf("Total Pengeluaran: Rp %.2f\n", total)
}

func editPengeluaran(A *tabInt, K *tabString, n int) {
	cetakElemen(*A, *K, n, "Pengeluaran")
	fmt.Print("Ubah nomor (0 untuk batal): ")
	var idx int
	fmt.Scan(&idx)
	if idx < 1 || idx > n {
		fmt.Println("Tidak ada yang diubah")
		return
	}
	idx--
	fmt.Print("Deskripsi baru: ")
	fmt.Scan(&(*K)[idx])
	fmt.Print("Nominal baru: Rp ")
	fmt.Scan(&(*A)[idx])
}

func tambahPemasukan(A *tabInt, K *tabString, n *int) {
	if *n >= NMAX {
		fmt.Println("Data penuh.")
		return
	}
	fmt.Print("Sumber pemasukan: ")
	fmt.Scan(&(*K)[*n])
	fmt.Print("Nominal: Rp ")
	fmt.Scan(&(*A)[*n])
	*n++
}

func editPemasukan(A *tabInt, K *tabString, n int) {
	cetakElemen(*A, *K, n, "Pemasukan")
	fmt.Print("Ubah nomor: ")
	var idx int
	fmt.Scan(&idx)
	if idx < 1 || idx > n {
		fmt.Println("Nomor tidak valid.")
		return
	}
	idx--
	fmt.Print("Sumber baru: ")
	fmt.Scan(&(*K)[idx])
	fmt.Print("Nominal baru: Rp ")
	fmt.Scan(&(*A)[idx])
}

func cetakElemen(A tabInt, K tabString, n int, judul string) {
	fmt.Println("\n", judul, ":")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s = Rp %.2f\n", i+1, K[i], A[i])
	}
}

func total(A tabInt, n int) float64 {
	t := 0.0
	for i := 0; i < n; i++ {
		t += A[i]
	}
	return t
}

func laporan(P tabInt, K tabString, nP int, M tabInt, S tabString, nM int) {
	fmt.Println("\n==== Laporan Budget ====")
	cetakElemen(P, K, nP, "Pengeluaran")
	cetakElemen(M, S, nM, "Pemasukan")
	totP := total(P, nP)
	totM := total(M, nM)
	fmt.Printf("Total Pengeluaran: Rp %.2f\n", totP)
	fmt.Printf("Total Pemasukan: Rp %.2f\n", totM)
	fmt.Printf("Sisa: Rp %.2f\n", totM-totP)

	if totP < totM {
		fmt.Println("Biaya perjalanan ada sudah tercukupi!")
	} else {
		fmt.Println("Ayo semangat nabung lagi!!!")
	}
}

func rekomendasiTabungan(A tabInt, n int) {
	selisih := tglTujuan.Sub(time.Now())
	hari := int(selisih.Hours() / 24)
	if hari <= 0 {
		fmt.Println("Sudah lewat/tiba hari-H")
		return
	}
	totalBiaya := total(A, n)
	fmt.Printf("Tabungan per hari (%.2f / %d): Rp %.2f\n", totalBiaya, hari, totalBiaya/float64(hari))
}
