# lottery

## Install

`go get -u github.com/syhlion/lottery`

## Usaged

``` go
type Item struct {
    Name string
    Prob int
}
func (i Item) Prob() int{
    return i.Prob
}
func main(){
    lott := New()
    items :=[]lottery.Item{
        Item{Name:"a",Prob:5},
        Item{Name:"b",Prob:20},
        Item{Name:"c",Prob:10},
        Item{Name:"d",Prob:10},
        Item{Name:"e",Prob:50},
        Item{Name:"f",Prob:5},
    }

    //mulit rand
    idx:=lottery.Pick(items ...)
    switch i:= items[idx].(type){
        case Item:
         fmt.Println(i.Name)
    }

    //shuffle rand
    shuffleItems := lott.Shuffle(items ...)

}

```
