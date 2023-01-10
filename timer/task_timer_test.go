package timer

import (
	"fmt"
	"testing"
	"time"
)

func now() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05:000"))
}

func TestTimer(t *testing.T) {
	timer := NewTask(time.Second*1, now)
	timer1 := NewTask(time.Second*2, now)
	timer.Start()
	timer1.Start()

	time.Sleep(time.Second * 10)

	timer.Pause()
	time.Sleep(time.Second * 5)
	timer.Restart()
	time.Sleep(time.Second * 5)
	timer.Stop()

}

func task() {
	for i := 0; i < 10; i++ {
		fmt.Println("-->", i)
		time.Sleep(time.Second)
	}
}

func TestTimerStop(t *testing.T) {
	timer := NewTask(time.Second*1, task)
	timer.Start()
	fmt.Println("--0>>", timer.Running())
	time.Sleep(time.Second * 5)
	timer.Stop()
	fmt.Println("--1>>", timer.Running())
	time.Sleep(time.Second * 10)
	fmt.Println("--2>>", timer.Running())

	/* print
	--0>> true
	--> 0
	--> 1
	--> 2
	--> 3
	--1>> false
	--> 4
	--> 5
	--> 6
	--> 7
	--> 8
	--> 9
	--2>> false
	*/
}
