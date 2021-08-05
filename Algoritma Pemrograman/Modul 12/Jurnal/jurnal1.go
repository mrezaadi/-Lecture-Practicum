package main

import "fmt"

const NMAX = 1000

type statistik_pemain struct {
	nama   string
	gol    int
	assist int
}

type data_pemain struct {
	tabel [NMAX]statistik_pemain
}

func main() {
	var data data_pemain
	var n int
	fmt.Scanln(&n)
	fillData(&data, n)
	descSort(&data, n)

	fmt.Println("\nSetelah Diurutkan:")

	for i := 0; i < n; i++ {
		fmt.Println(data.tabel[i].nama, data.tabel[i].gol, data.tabel[i].assist)
	}

}

func fillData(tab *data_pemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&tab.tabel[i].nama, &tab.tabel[i].gol, &tab.tabel[i].assist)
	}
}

func descSort(tab *data_pemain, n int) {
	var pass, idx, i int
	var temp statistik_pemain

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if tab.tabel[idx].gol == tab.tabel[i].gol {
				if tab.tabel[idx].assist < tab.tabel[i].assist {
					idx = i
				}
			} else if tab.tabel[idx].gol < tab.tabel[i].gol {
				idx = i
			}
			i++
		}
		temp = tab.tabel[pass-1]
		tab.tabel[pass-1] = tab.tabel[idx]
		tab.tabel[idx] = temp
		pass++
	}
}

// 9
// BrunoFernandes 7 3
// JamieVardy 8 1
// Dominic 10 2
// SonHeungMin 9 2
// HarryKane 7 2
// OllieWatkins 6 1
// PattrickBamford 7 1
// CallumWillson 7 0
// MohammedSalah 7 2
