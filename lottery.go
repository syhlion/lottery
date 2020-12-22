package lottery

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	rand "math/rand"
	"sort"
)

var src cryptoSource

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func New() *Lottery {
	return &Lottery{
		rand: rand.New(src),
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

func (l *Lottery) Shuffle(items ...Item) (dest []Item) {
	dest = make([]Item, len(items))
	perm := l.rand.Perm(len(items))
	for i, v := range perm {
		dest[v] = items[i]
	}
	return
}

func (l *Lottery) Pick(items ...Item) int {
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
