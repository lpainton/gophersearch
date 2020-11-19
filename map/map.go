package map

type dangerLevel = int

const (
	unknown dangerLevel = iota
	zero    dangerLevel = iota
	one     dangerLevel = iota
	two     dangerLevel = iota
	three   dangerLevel = iota
	four    dangerLevel = iota
	five    dangerLevel = iota
	six     dangerLevel = iota
	seven   dangerLevel = iota
	eight   dangerLevel = iota
	gopher    dangerLevel = iota
)

type distance = int
type seed = string

// Minefield describes a complete game state
type Game struct {
	Width   distance
	Height  distance
	Squares []dangerLevel
	Seed    seed
}

func create(width distance, height distance) Game {

}