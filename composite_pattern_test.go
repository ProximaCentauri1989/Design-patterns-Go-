package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompositePattern(t *testing.T) {
	//Box with wires (box with 3 objects)
	usbCable := NewObject("Xiaomi2A", 3.0)
	typeCCable := NewObject("Xiaomi2.4A", 5.0)
	powerCable := NewObject("PowerPlugCable", 2.0)

	boxWithWires := NewBox("My tangled wires")
	boxWithWires.Add(usbCable)
	boxWithWires.Add(typeCCable)
	boxWithWires.Add(powerCable)
	assert.Equal(t, float32(10.0), boxWithWires.Price())
	assert.Equal(t, 3, boxWithWires.Amount())

	//Box with subjects for writing (bix with 4 objects)
	boxWithWritibleObjects := NewBox("My writible objects")
	pen := NewObject("Pen", 0.25)
	blackMarker := NewObject("Black marker", 0.75)
	redMarker := NewObject("Red marker", 0.50)
	pencil := NewObject("Gray pencil", 0.25)
	boxWithWritibleObjects.Add(pen)
	boxWithWritibleObjects.Add(blackMarker)
	boxWithWritibleObjects.Add(redMarker)
	boxWithWritibleObjects.Add(pencil)

	//Box with one object and two previous boxes
	bigBox := NewBox("My stuff")
	bigBox.Add(boxWithWires)
	assert.Equal(t, float32(10.0), bigBox.Price())
	assert.Equal(t, 3, bigBox.Amount())

	bigBox.Add(boxWithWritibleObjects)
	assert.Equal(t, float32(11.75), bigBox.Price())
	assert.Equal(t, 7, bigBox.Amount())

	//Add one more element to the big box
	bigBox.Add(NewObject("Santa Claus's mustache", 0.25))
	assert.Equal(t, float32(12.0), bigBox.Price())
	assert.Equal(t, 8, bigBox.Amount())

	bigBox.ListObjects(0)

	bigBox.Clear()
	assert.Equal(t, 0, bigBox.Amount())
}
