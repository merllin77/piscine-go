package main

import (
	"fmt"
)

func QuadA(x, y int) {
	chky := 1                               // metavliti elegxou an vriskomaste se "moni" i "zigi" grammi (an einai 1 = print"/" an einai 2 = print:"\")
	if x > 0 && y > 0 && y != 1 || x != 1 { // elegxos an einai oi times einai thetikes kai AN oi grammes (y) i oi stiles (x) einai perisoteres apo mia
		for i := 1; i != y && y != 1; i++ { // loop gia to poses grammes (katheta) na ftiaxei GIA OSO DEN EINAI STIN TELEUTAI GRAMMI!
			for j := 1; j <= x; j++ { // loop gia to poses stiles (orizontia/ mikos tetragonou) tha fitaxei
				if (j == 1 || j == x) && chky == 1 { // elegxos kai prosthisi "/" an vriskomaste se "moni" grammi stin arxi i sto telos
					if j == x {
						fmt.Println("\\")
						chky = 2
					} else {
						fmt.Print("/")
					}
				} else if (j == 1 || j == x) && chky == 2 { /// elegxos kai tiposi "\" an vriskomaste se "zygi" grammi stin arxi i sto telos
					if j == x {
						fmt.Println("*")
						//	chky = 2
					} else {
						fmt.Print("*")
					}
				} else if chky == 1 { // tiposi "*" an eimaste se "moni" grammi
					fmt.Print("*")
				} else if i != y {
					fmt.Print(" ")
				} else {
					fmt.Print("A ") // tiposi " " (kena) an eimaste se "zygi grammi"
				}
			}
		}
		if x == 1 { // ti tiponei an exoume mono mia stili ???
			for i := 1; i <= y; i++ {
				if i == y {
					fmt.Println("\\")
				} else {
					if chky == 1 {
						fmt.Println("/")
						chky = 2
					} else {
						fmt.Println("*")
					}
				}
			}
		} else {
			//	fmt.Println("EDO")
			fmt.Print("\\") // ti tiponei stin teleutai grammi
			for j := 1; j < x-1; j++ {
				fmt.Print("*")
			}
			fmt.Println("/")
		}
	} else if y == 1 {
		for j := 1; j <= x; j++ { // loop gia to poses stiles (orizontia/ mikos tetragonou) tha fitaxei
			if (j == 1 || j == x) && chky == 1 { // elegxos kai prosthisi "/" an vriskomaste se "moni" grammi stin arxi i sto telos
				if j == x {
					fmt.Println("/")
				} else {
					fmt.Println("\\")
				}
			}
		}
	}
}

func main() {
	QuadA(5, 3)
}
