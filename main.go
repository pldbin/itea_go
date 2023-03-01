package main

import (
	"fmt"
	"sort"
)

type Horse struct {
	power int
	ID    int
}
type Horses struct {
	horses []Horse
}

var horses Horses

func (h *Horses) addhorse(power int) {
	horse := Horse{}
	horse.power = power
	horse.ID = len(h.horses)
	h.horses = append(h.horses, horse)
}

func soft() {
	var horses int
	fmt.Scanf("%d", &horses)
	strengths := make([]int, 0, horses)
	for i := 0; i < horses; i++ {
		var strength int
		fmt.Scanf("%d", &strength)
		strengths = append(strengths, strength)
	}
	fmt.Println("str", strengths)
	sort.Ints(strengths)
	fmt.Println("str2", strengths)
	for i := 1; i < horses; i++ {
		minstr := (strengths[i] - strengths[i-1])
		fmt.Printf("%d\n", minstr)
		min := minstr
		for _, element := range strengths {
			if element < minstr {
				continue
				//min = element
			}
		}
		fmt.Println("min", min)
	}

}

/*
потрібно знайти мінімальну різницю
брати коня 0 і коня 1 і взяти їх різницю,
кінь 1 і кінь 2 і порівняти різниця їх сили меньше чим різниця сили першої пари якщо так
то берем цю різницю як мінімальною. І так по всім парам поки не знайдемо мінімальну різницю пар.
*/
func main() {
	soft()
}
