package section2

import (
	"fmt"
	"testing"
)

func TestViolinist_Greetings(t *testing.T) {
	musicians := []MusicalPlayer{
		&Violinist{Musician{"Raul"}},
		&Trumpeter{Musician{"David"}},
	}
	for _, musician := range musicians {
		fmt.Println(musician.Greetings())
	}
}
