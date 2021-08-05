package main

import "fmt"

const NMAX = 100

type partai struct {
	nama  string
	suara int
}

type tab_partai struct {
	tabel [NMAX]partai
}

func main() {
	var data tab_partai
	var suara string
	var index int
	var n int

	k := 0
	for k < NMAX {
		fmt.Scan(&suara)
		if suara != "-1" {
			index = posisi(data, suara)
			if index == -1 {
				data.tabel[k].nama = suara
				data.tabel[k].suara++
				k++
				n++
			} else {
				data.tabel[index].suara++
			}
		} else {
			break
		}
	}

	var pass, idx, i int
	var temp partai

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if data.tabel[idx].suara < data.tabel[i].suara {
				idx = i
			}
			i++
		}
		temp = data.tabel[pass-1]
		data.tabel[pass-1] = data.tabel[idx]
		data.tabel[idx] = temp
		pass++
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%v(%v) ", data.tabel[j].nama, data.tabel[j].suara)
	}
}

func posisi(data tab_partai, x string) int {
	var idx int

	idx = -1
	for i := 0; i < len(data.tabel); i++ {
		if data.tabel[i].nama == x {
			idx = i
			break
		}
	}
	return idx
}

// 14 10 13 13 14 10 11 13 13 12 15 11 10 -1
