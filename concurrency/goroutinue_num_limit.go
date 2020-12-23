package concurrency

import "fmt"

func LimitNGoroutinue(n int) {
	ch := make(chan bool, n)
	list := []string{
		"1",
		"2",
	}
	for _, v := range list {
		ch <- true
		go func(v string, ch chan bool) {
			Do(v)
		}(v, ch)

	}
}

func Do(v string	)  {
	fmt.Println(v)
}