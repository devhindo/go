package main

type chessPiece struct {
	name string
	value int
	place int
}

func (c *chessPiece) move() {
	c.place += 5
}