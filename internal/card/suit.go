package card

type Suit struct {
	LongName  string
	ShortName string
}

var (
	Spades   = Suit{"Spades", "s"}
	Hearts   = Suit{"Hearts", "h"}
	Diamonds = Suit{"Diamonds", "d"}
	Clubs    = Suit{"Clubs", "c"}
)
