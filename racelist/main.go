package main

import (
	"fmt"
	"time"
)

func main() {
	a1 := newAlbe1()
	a2 := newAlbe2()

	lists.add(
		a1,
	)
	lists.add(
		a2,
	)

	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G1")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G2")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G3")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G4")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G5")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G6")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G7")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G8")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G9")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()
	go func() {
		// for  {
		able3 := newAlbe3(
			func() {
				fmt.Println("G10")
			},
		)
		lists.add(able3)
		lists.notify()
		lists.del(able3)
		// }
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("+++++++++++++")
	lists.notify()

	ablef := newAlbe3(
		func() {
			fmt.Println("G6")
		},
	)

	fmt.Println(ablef == ablef)
}
