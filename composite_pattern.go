package composite

import "fmt"

//Interface for tree structure
type Component interface {
	Add(c Component)
	Clear()
	Price() float32
	IsBox() bool
	ListObjects(int)
	Amount() int
}

//Box implements box with items (can be box or object)
type Box struct {
	compnts []Component
	title   string
}

func NewBox(title string) Component {
	return &Box{title: title}
}

func (b *Box) Add(c Component) {
	b.compnts = append(b.compnts, c)
}

func (b *Box) Price() float32 {
	var price float32
	for _, c := range b.compnts {
		price += c.Price()
	}
	return price
}

func (b *Box) Clear() {
	for pos, c := range b.compnts {
		if !c.IsBox() {
			b.compnts = append(b.compnts[:pos], b.compnts[pos+1:]...)
		}
	}
}

func (b *Box) IsBox() bool {
	return true
}

func (b *Box) Name() string {
	return b.title
}

func (b *Box) ListObjects(indent int) {
	PrintIndent(indent)

	fmt.Println(fmt.Sprintf("Box with '%s' contains: ", b.title))
	for _, c := range b.compnts {
		c.ListObjects(indent + 1)
	}
}

func (b *Box) Amount() int {
	var amount int
	for _, c := range b.compnts {
		amount += c.Amount()
	}
	return amount
}

type Object struct {
	price float32
	name  string
}

func NewObject(name string, price float32) Component {
	return &Object{
		name:  name,
		price: price,
	}
}

func (o *Object) Name() string {
	return o.name
}

func (o *Object) Add(c Component) {
	//Empty
}

func (o *Object) Clear() {
	//Empty
}

func (o *Object) ListObjects(indent int) {
	PrintIndent(indent)
	fmt.Println(fmt.Sprintf("Item with name '%s'", o.Name()))
}

func (o *Object) Price() float32 {
	return o.price
}

func (o *Object) IsBox() bool {
	return false
}

func (o *Object) Amount() int {
	return 1
}

func PrintIndent(indent int) {
	for count := indent; count > 0; count-- {
		fmt.Print("\t")
	}
}
