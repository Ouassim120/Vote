package Loop

import "fmt"

func Loop() {
Loop:
	for h := 0; h < 60; h++ {
		for m := 0; m < 60; m++ {
			for s := 0; s < 60; s++ {
				// time.Sleep(1 * time.Second)
				fmt.Printf("%vh %vm %vs\n", h, m, s)
				if h == 1 && m == 30 {
					break Loop
				}
			}
		}
	}

}
