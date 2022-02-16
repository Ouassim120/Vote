package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	fmt.Println("Moving Mouse")
	robotgo.MoveMouseSmooth(37, 333, 0.0, 0.0) //chrome in taskbar
	robotgo.Click()
	time.Sleep(2 * time.Second)
	robotgo.MoveMouseSmooth(445, 84, 0.0, 0.0) // Arcadia Tab
	robotgo.Click()
	time.Sleep(2 * time.Second)
	robotgo.MoveMouseSmooth(608, 604, 0.0, 0.0) // CLICK sur Voter
	robotgo.Click()
	time.Sleep(25 * time.Second)                // Solve captcha auto
	robotgo.MoveMouseSmooth(950, 725, 0.0, 0.0) // confirm Vote in serveurprivé.net
	robotgo.Click()
	time.Sleep(2 * time.Second)
	robotgo.MoveMouseSmooth(124, 10, 0.0, 0.0) // click on Tab
	robotgo.Click()
	time.Sleep(2 * time.Second)
	robotgo.MoveMouseSmooth(578, 891, 0.0, 0.0) // Butser captcha
	robotgo.Click()
	time.Sleep(1 * time.Second)
	robotgo.MoveMouseSmooth(409, 546, 0.0, 0.0) // j'ai bien éffectué mon vote
	robotgo.Click()
	time.Sleep(1 * time.Second)
	robotgo.MoveMouseSmooth(469, 702, 0.0, 0.0) // Valider
	robotgo.Click()
	time.Sleep(1 * time.Second)
	robotgo.MoveMouseSmooth(230, 15, 0.0, 0.0) // Close Tab
	robotgo.Click()
	time.Sleep(1 * time.Second)
	robotgo.Click()

}
