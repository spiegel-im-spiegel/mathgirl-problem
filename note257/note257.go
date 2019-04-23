package note257

import (
	"github.com/spiegel-im-spiegel/mathgirl-problem/prob"
)

type People struct {
	infect <-chan bool
}

type Person bool

func NewPeople() *People {
	return &People{infect: prob.New(0.01)}
}

func (ppl *People) SelectPersion() Person {
	return Person(<-ppl.infect)
}

func (psn Person) Infection() bool {
	return bool(psn)
}

type TestKit struct {
	probability <-chan bool
}

func NewTestKit() *TestKit {
	return &TestKit{probability: prob.New(0.9)}
}

func (tk *TestKit) Inspect(psn Person) bool {
	if psn.Infection() {
		return <-tk.probability
	}
	return !<-tk.probability
}
