package main

import (
	"fmt"
	"strings"
)

type Expense struct {
	Category string
	Amount   float64
}

var expenses []Expense
var totalBudget float64

func main() {
	totalBudget = 1000 // Contoh anggaran awal
	for {
		showMenu()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addExpense()
		case 2:
			editExpense()
		case 3:
			deleteExpense()
		case 4:
			showTotal(expenses)
		case 5:
			searchExpense()
		case 6:
			sortExpenses()
		case 7:
			showReport()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func showMenu() {
	fmt.Println("\nAplikasi Pengelolaan Budget Traveling")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Ubah Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Tampilkan Total Pengeluaran & Saran")
	fmt.Println("5. Cari Pengeluaran")
	fmt.Println("6. Urutkan Pengeluaran")
	fmt.Println("7. Laporan")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func addExpense() {
	var category string
	var amount float64
	fmt.Print("Kategori: ")
	fmt.Scan(&category)
	fmt.Print("Jumlah: ")
	fmt.Scan(&amount)
	expenses = append(expenses, Expense{Category: strings.ToLower(category), Amount: amount})
	fmt.Println("Pengeluaran ditambahkan.")
}

func editExpense() {
	showExpenses()
	var idx int
	fmt.Print("Pilih nomor yang ingin diubah: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > len(expenses) {
		fmt.Println("Index tidak valid.")
		return
	}
	fmt.Print("Kategori baru: ")
	fmt.Scan(&expenses[idx-1].Category)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&expenses[idx-1].Amount)
	fmt.Println("Pengeluaran diubah.")
}

func deleteExpense() {
	showExpenses()
	var idx int
	fmt.Print("Pilih nomor yang ingin dihapus: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > len(expenses) {
		fmt.Println("Index tidak valid.")
		return
	}
	expenses = append(expenses[:idx-1], expenses[idx:]...)
	fmt.Println("Pengeluaran dihapus.")
}

func showExpenses() {
	for i, e := range expenses {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, e.Category, e.Amount)
	}
}

func showTotal(e expenses) {
	total := 0.0
	for e := range expenses {
		total += e.Amount
	}
	fmt.Printf("Total Pengeluaran: Rp %.2f\n", total)
	if total > totalBudget {
		fmt.Printf("Anggaran terlampaui sebesar Rp %.2f\n", total-totalBudget)
	} else {
		fmt.Printf("Sisa anggaran: Rp %.2f\n", totalBudget-total)
	}
}

func searchExpense() {
	var category string
	fmt.Print("Masukkan kategori (case-insensitive): ")
	fmt.Scan(&category)
	category = strings.ToLower(category)

	// Sequential Search
	fmt.Println("\n[Hasil Pencarian (Sequential Search)]")
	found := false
	for _, e := range expenses {
		if e.Category == category {
			fmt.Printf("- %s: Rp %.2f\n", e.Category, e.Amount)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortExpenses() {
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Jumlah (descending)")
	fmt.Println("2. Kategori (ascending)")
	var opt int
	fmt.Print("Pilihan: ")
	fmt.Scan(&opt)

	if opt == 1 {
		// Selection Sort descending
		n := len(expenses)
		for i := 0; i < n-1; i++ {
			maxIdx := i
			for j := i + 1; j < n; j++ {
				if expenses[j].Amount > expenses[maxIdx].Amount {
					maxIdx = j
				}
			}
			expenses[i], expenses[maxIdx] = expenses[maxIdx], expenses[i]
		}
		fmt.Println("Diurutkan berdasarkan jumlah (desc).")
	} else if opt == 2 {
		// Insertion Sort ascending by category
		for i := 1; i < len(expenses); i++ {
			key := expenses[i]
			j := i - 1
			for j >= 0 && expenses[j].Category > key.Category {
				expenses[j+1] = expenses[j]
				j--
			}
			expenses[j+1] = key
		}
		fmt.Println("Diurutkan berdasarkan kategori (asc).")
	}
}

func showReport() {
	categoryMap := make(map[string]float64)
	total := 0.0
	for _, e := range expenses {
		categoryMap[e.Category] += e.Amount
		total += e.Amount
	}
	fmt.Println("\nLaporan Pengeluaran Berdasarkan Kategori:")
	for cat, amt := range categoryMap {
		fmt.Printf("- %s: Rp %.2f\n", cat, amt)
	}
	fmt.Printf("Total: Rp %.2f | Budget: Rp %.2f\n", total, totalBudget)
	fmt.Printf("Selisih: Rp %.2f\n", totalBudget-total)
}
