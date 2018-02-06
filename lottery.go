package lottery

import (
	"math/rand"
	"sort"
	"time"
)

func New() *Lottery {
	return &Lottery{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type Item interface {
	Prob() int
}
type ItemSort []Item

func (is ItemSort) Len() int {
	return len(is)
}
func (is ItemSort) Less(i, j int) bool {
	return is[i].Prob() < is[j].Prob()
}
func (is ItemSort) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type Lottery struct {
	rand *rand.Rand
}

func (l *Lottery) Rand(items ...Item) int {
	l.rand.Seed(time.Now().UnixNano())
	total := 0
	for _, item := range items {
		total += item.Prob()
	}
	if total <= 0 {
		return -1
	}
	randNum := l.rand.Intn(total) + 1
	tmpNum := 0
	itemIndex := -1
	sort.Sort(ItemSort(items))
	for idx, l := range items {
		tmpNum += l.Prob()
		if tmpNum >= randNum {
			itemIndex = idx
			break
		}

	}
	return itemIndex

}
