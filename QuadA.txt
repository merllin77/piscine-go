package main

import "fmt"

func QuadA(x, y int) {
	chky := 1            // metavliti elegxou an vriskomaste se "moni" i "zigi" grammi (an einai 1 = print"o" an einai 2 = print:"|")
	if x > 0 && y >= 1 { // elegxos an einai oi times einai thetikes
		for i := 0; i < y; i++ { // loop gia to poses grammes na (katheta) na ftiaxei
			for j := 1; j <= x; j++ { // loop gia to poses stiles (orizontia/ mikos tetragonou) tha fitaxei
				if (j == 1 || j == x) && chky == 1 { // elegxos kai prosthisi "o" an vriskomaste se "moni" grammi stin arxi i sto telos
					if j == x {
						fmt.Println("o")
						chky = 2
					} else {
						fmt.Print("o")
					}
				} else if (j == 1 || j == x) && chky == 2 { /// elegxos kai tiposi "|" an vriskomaste se "zygi" grammi stin arxi i sto telos
					if j == x {
						fmt.Println("|")
						chky = 1
					} else {
						fmt.Print("|")
					}
				} else if chky == 1 { // tiposi "-" an eimaste se "moni" grammi
					fmt.Print("-")
				} else {
					fmt.Print(" ") // i tiposi " " (kena) an eimaste se "zygi grammi"
				}
			}
		}
	}
}

func main() {
	QuadA(5, 3)
}
