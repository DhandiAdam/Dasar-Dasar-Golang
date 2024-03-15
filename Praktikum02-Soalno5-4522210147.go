package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	usiaStr := scanner.Text()

	// Mengonversi input usia menjadi integer
	usia, err := strconv.Atoi(usiaStr)
	if err != nil {
		fmt.Println("Usia harus berupa angka.")
		return
	}

	// Menampilkan hasil input
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Usia: %d\n", usia)
}
