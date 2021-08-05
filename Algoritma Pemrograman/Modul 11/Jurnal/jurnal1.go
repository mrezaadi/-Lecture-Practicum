package main

import "fmt"

const NMAX = 1000

type waktu struct {
	menit, detik int
}

type lagu struct {
	judul, penyanyi string
	durasi          waktu
}

type playlist struct {
	list [NMAX]lagu
}

func main() {
	var data playlist
	var n int
	var long int

	makePlaylist(&data, &n)
	long = longestSong(data, n)
	showPlaylist(data, n, long)

}

func findSong(x, y string, tab playlist) bool {
	var find bool

	i := 0
	for i < len(tab.list) {
		if x == tab.list[i].judul && y == tab.list[i].penyanyi {
			find = true
			i++
			break
		} else {
			find = false
			i++
		}
	}
	return find
}

func makePlaylist(tab *playlist, n *int) {
	var judul, penyanyi string
	var menit, detik int
	var find bool

	i := 0
	for i < NMAX {
		if judul != "#" && penyanyi != "#" {
			fmt.Scanln(&judul, &penyanyi, &menit, &detik)
			if judul == "#" && penyanyi == "#" {
				break
			}
			find = findSong(judul, penyanyi, *tab)
			if find == false {
				tab.list[i].judul = judul
				tab.list[i].penyanyi = penyanyi
				tab.list[i].durasi.menit = menit
				tab.list[i].durasi.detik = detik
				*n++
				i++
			}
		} else {
			break
		}
	}
}

func longestSong(tab playlist, n int) int {
	var longestminute, longestsecond int
	var longest int

	longestminute = tab.list[0].durasi.menit
	longestsecond = tab.list[0].durasi.detik
	longest = 0

	i := 1
	for i < n {
		if tab.list[i].durasi.menit == longestminute {
			if tab.list[i].durasi.detik > longestsecond {
				longestsecond, longestminute = tab.list[i].durasi.detik, tab.list[i].durasi.menit
				longest = i
				i++
			} else {
				i++
			}
		} else if tab.list[i].durasi.menit > longestminute {
			longestsecond, longestminute = tab.list[i].durasi.detik, tab.list[i].durasi.menit
			longest = i
			i++
		} else {
			i++
		}
	}
	return longest
}

func showPlaylist(tab playlist, n int, longest int) {
	i := 0

	for i < n {
		if i == longest {
			fmt.Println("*", tab.list[i].judul, tab.list[i].durasi.menit, "menit", tab.list[i].durasi.detik, "detik")
			i++
		} else {
			fmt.Println(tab.list[i].judul)
			i++
		}
	}
}

/*
Snowman Sia 2 53
Bertaut NadinAmizah 5 21
Anyone JustinBieber 3 16
Cuek RizkyFebian 4 20
Anyone JustiBieber 4 24
Sofia Clairo 3 8
Snowman Sia 2 42
# #

SomebodyToldMe TheKillers 3 21
SomebodyToldMe TheKillers 3 22
SomebodyToldMe TheKillers 3 23
SomebodyToldMe TheKillers 3 24
# #
*/
