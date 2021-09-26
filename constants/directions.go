package constants

const (
	Left  string = "L"
	Right string = "R"
	Move  string = "M"
	South string = "S"
	North string = "N"
	West  string = "W"
	East  string = "E"
)

var DirectionRotateNumbers = map[string]int{
	"N": 1,
	"W": 2,
	"S": 3,
	"E": 4,
}
