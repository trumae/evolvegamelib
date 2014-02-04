package evolvegamelib

import "testing"

func TestAngle(t *testing.T) {
	n := angle([8]int{2, 2, 2, 2, 2, 2, 2, 2}, 1)
	if n != 0 {
		t.Error("Erro na funcao angle")
	}
	n = angle([8]int{2, 2, 2, 2, 2, 2, 2, 2}, 3)
	if n != 1 {
		t.Error("Erro na funcao angle")
	}
	n = angle([8]int{2, 2, 2, 2, 2, 2, 2, 2}, 5)
	if n != 2 {
		t.Error("Erro na funcao angle")
	}
	n = angle([8]int{2, 2, 2, 2, 2, 2, 2, 2}, 7)
	if n != 3 {
		t.Error("Erro na funcao angle")
	}
	n = angle([8]int{2, 2, 2, 2, 2, 2, 2, 2}, 9)
	if n != 4 {
		t.Error("Erro na funcao angle")
	}
}
