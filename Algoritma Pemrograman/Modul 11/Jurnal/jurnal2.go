package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}

type tabTopic struct {
	array [100]string
}

type tabTag struct {
	array [100]tag
}

func main() {
	var tabTG tabTag
	var tabTP tabTopic

	fillTabTag(&tabTG, &tabTP)
	fmt.Println(trendingTopic(tabTG))
}

func findTopic(tab tabTag, x string, idx *int, find *bool) {

	*find = false

	i := 0
	for i < len(tab.array) {
		if x == tab.array[i].topik {
			*find = true
			*idx = i
			break
		} else {
			i++
		}
	}
}

func fillTabTag(tabTG *tabTag, tabTP *tabTopic) {
	var topik string
	var ketemu bool
	var index int
	var index_topik int

	i := 0
	for i < len(tabTP.array) {
		fmt.Scan(&topik)
		if topik == "." {
			break
		}
		tabTP.array[i] = topik
		findTopic(*tabTG, topik, &index, &ketemu)
		if ketemu == true {
			tabTG.array[index].banyak++
			i++
		} else {
			tabTG.array[index_topik].topik = topik
			index_topik++
			i++
		}
	}
}

func trendingTopic(tabTG tabTag) string {
	var trending_idx int
	var max int

	max = tabTG.array[0].banyak
	i := 1

	for i < len(tabTG.array) {
		if tabTG.array[i].banyak > max {
			tabTG.array[i].banyak, max = max, tabTG.array[i].banyak
			trending_idx = i
			i++
		} else {
			i++
		}
	}
	return tabTG.array[trending_idx].topik
}

//libur lebaran libur tugas libur lebaran libur mudik libur libur rendang libur puasa uas libur lebaran tugas libur libur 2minggu vaksin
