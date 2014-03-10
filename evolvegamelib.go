package evolvegamelib

import (
	"fmt"
	"math/rand"
)

type Pos struct {
	x int
	y int
}

type Animal struct {
	x                  int
	y                  int
	dir                int
	energy             int
	genes              [8]int
	reproductionEnergy int
}

type Arena struct {
	width        int
	height       int
	jungleX      int
	jungleY      int
	jungleWidth  int
	jungleHeight int
	plantEnergy  int
	plants       map[Pos]bool
	animals      []Animal
	moveDecr     int
}

func (a *Arena) AddPlants() {
	a.plants[Pos{rand.Intn(a.width), rand.Intn(a.height)}] = true
	a.plants[Pos{a.jungleX + rand.Intn(a.jungleWidth),
		a.jungleY + rand.Intn(a.jungleHeight)}] = true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (animal *Animal) Move(arena *Arena) {
	dir := animal.dir

	// get delta in x axis
	deltaX := 0
	if dir == 2 || dir == 3 || dir == 4 {
		deltaX = 1
	}
	if dir == 0 || dir == 7 || dir == 6 {
		deltaX = -1
	}

	// get delta in y axis
	deltaY := 0
	if dir == 0 || dir == 1 || dir == 2 {
		deltaY = 1
	}
	if dir == 6 || dir == 6 || dir == 4 {
		deltaY = -1
	}

	animal.x = abs((animal.x+deltaX)+arena.width) % arena.width
	animal.y = abs((animal.y+deltaY)+arena.height) % arena.height

	animal.energy = animal.energy - 5
}

func angle(genes [8]int, n int) int {
	count := 0
	for i, val := range genes {
		count += val
		if n <= count {
			return i
		}
	}
	return 7
}

func (animal *Animal) Turn() {
	somaGenes := 0
	for _, v := range animal.genes {
		somaGenes += v
	}
	if somaGenes == 0 {
		somaGenes = 10
	}
	x := rand.Intn(somaGenes)
	animal.dir = angle(animal.genes, x)
}

func (animal *Animal) Eat(arena *Arena) {
	_, ok := arena.plants[Pos{animal.x, animal.y}]
	if ok {
		animal.energy += arena.plantEnergy
		delete(arena.plants, Pos{animal.x, animal.y})
	}
}

func (animal *Animal) Reproduce(arena *Arena) {
	if animal.energy >= animal.reproductionEnergy {
		animal.energy = animal.energy / 2
		animalnu := *animal
		x := rand.Intn(8)
		ngen := animal.genes[x] + rand.Intn(2)
		if ngen > 1 {
			animalnu.genes[x] = ngen
		} else {
			animalnu.genes[x] = 1
		}
		arena.animals = append(arena.animals, animalnu)
	}
}

func (a *Arena) UpdateWorld() {
	ans := make([]Animal, 0)
	for _, val := range a.animals {
		if val.energy > 0 {
			ans = append(ans, val)
		}
	}
	a.animals = ans
	for i, _ := range a.animals {
		a.animals[i].Turn()
		a.animals[i].Move(a)
		a.animals[i].Eat(a)
		a.animals[i].Reproduce(a)
	}
	a.AddPlants()
}

func (a *Arena) DrawWorldString() string {
	ret := ""
	for i := 0; i < a.height; i++ {
		fmt.Print("|")
		for j := 0; j < a.width; j++ {
			desenhado := false
			pos := Pos{j, i}
			for _, elem := range a.animals {
				if elem.x == j && elem.y == i && desenhado == false {
					fmt.Print("M")
					desenhado = true
				}
			}
			if a.plants[pos] && desenhado == false {
				fmt.Print("*")
				desenhado = true
			}
			if !desenhado {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("Number of animals:")
	fmt.Print(len(a.animals))
	fmt.Print("\nNumber of trees")
	fmt.Print(len(a.plants))
	fmt.Print("\n")
	return ret
}

func (a *Arena) DrawWorld() {
	fmt.Print(a.DrawWorldString())
}

func (a *Arena) Evolution(n int) {
	for i := 0; i < n; i++ {
		a.UpdateWorld()
	}
}

func NewArena() *Arena {
	a := Arena{}
	a.width = 100
	a.height = 30
	a.jungleX = 45
	a.jungleY = 10
	a.jungleWidth = 10
	a.jungleHeight = 10
	a.plantEnergy = 200
	a.plants = make(map[Pos]bool)
	a.animals = make([]Animal, 0)
	a.moveDecr = 10

	return &a
}

func NewArenaSample() *Arena {
	a := Arena{}
	a.width = 100
	a.height = 30
	a.jungleX = 45
	a.jungleY = 10
	a.jungleWidth = 10
	a.jungleHeight = 10
	a.plantEnergy = 200
	a.plants = make(map[Pos]bool)
	a.animals = make([]Animal, 0)
	a.moveDecr = 50

	b := Animal{}
	b.genes = [8]int{2, 2, 2, 2, 2, 2, 2, 2}
	b.x = 20
	b.y = 15
	b.energy = 2000
	b.reproductionEnergy = 1000
	a.animals = append(a.animals, b)

	return &a
}
