package world

import "fmt"

const (
	//North Represents north.
	North int = iota
	//NorthWest Represents north-west.
	NorthWest
	//West Represents west.
	West
	//SouthWest Represents south-west.
	SouthWest
	//South represents south.
	South
	//SouthEast represents south-east
	SouthEast
	//East Represents east.
	East
	//NorthEast Represents north-east.
	NorthEast
	//MaxX Width of the game
	MaxX = 944
	//MaxY Height of the game
	MaxY = 3776
)

//Location A tile in the game world.
type Location struct {
	X, Y int
}

//DeathSpot The spot where mobs go to die.
var DeathSpot = NewLocation(0, 0)

//NewLocation Returns a reference to a new instance of the Location data structure.
func NewLocation(x, y int) *Location {
	return &Location{x, y}
}

//String Returns a string representation of the location
func (l *Location) String() string {
	return fmt.Sprintf("[%d,%d]", l.X, l.Y)
}

//Equals Returns true if this location points to the same location as o
func (l *Location) Equals(o interface{}) bool {
	if o, ok := o.(*Location); ok {
		return l.X == o.X && l.Y == o.Y
	}
	if o, ok := o.(Location); ok {
		return l.X == o.X && l.Y == o.Y
	}
	return false
}

//DeltaX Returns the difference between this locations X coord and the other locations X coord
func (l *Location) DeltaX(other Location) (deltaX int) {
	if l.X > other.X {
		deltaX = l.X - other.X
	} else if other.X > l.X {
		deltaX = other.X - l.X
	}
	return
}

//DeltaY Returns the difference between this locations Y coord and the other locations Y coord
func (l *Location) DeltaY(other Location) (deltaY int) {
	if l.Y > other.Y {
		deltaY = l.Y - other.Y
	} else if other.Y > l.Y {
		deltaY = other.Y - l.Y
	}
	return
}

//LongestDelta Returns the largest difference in coordinates between receiver and other
func (l *Location) LongestDelta(other Location) int {
	deltaX, deltaY := l.DeltaX(other), l.DeltaY(other)
	if deltaX > deltaY {
		return deltaX
	}
	return deltaY
}

//WithinRange Returns true if the other location is within radius tiles of the receiver location, otherwise false.
func (l *Location) WithinRange(other Location, radius int) bool {
	return l.LongestDelta(other) <= radius
}

//ParseDirection Tries to parse the direction indicated in s.  If it can not match any direction, returns the zero-value for direction: north.
func ParseDirection(s string) int {
	switch s {
	case "northeast":
		return NorthEast
	case "ne":
		return NorthEast
	case "northwest":
		return NorthWest
	case "nw":
		return NorthWest
	case "east":
		return East
	case "e":
		return East
	case "west":
		return West
	case "w":
		return West
	case "south":
		return South
	case "s":
		return South
	case "southeast":
		return SouthEast
	case "se":
		return SouthEast
	case "southwest":
		return SouthWest
	case "sw":
		return SouthWest
	case "n":
		return North
	case "north":
		return North
	}

	return North
}