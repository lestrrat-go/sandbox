package reflectelect

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"testing"
)

func setupChannels(ctx context.Context) (chan int, chan float64, chan byte) {
	a := make(chan int)
	b := make(chan float64)
	c := make(chan byte)

	go func(ctx context.Context, ch chan int) {
		var i int
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
				if i == math.MaxInt64 {
					i = 0
				} else {
					i++
				}
			}
		}
	}(ctx, a)

	go func(ctx context.Context, ch chan float64) {
		var f float64
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- f
				if f >= math.MaxFloat64 {
					f = 0
				} else {
					f += 1
				}
			}
		}
	}(ctx, b)

	go func(ctx context.Context, ch chan byte) {
		var i uint8
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- byte(i)
				if i == math.MaxUint8 {
					i = 0
				} else {
					i++
				}
			}
		}
	}(ctx, c)

	return a, b, c
}

func BenchmarkReflectSelect(b *testing.B) {
	b.StopTimer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1, ch2, ch3 := setupChannels(ctx)

	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch1),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch2),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch3),
		},
	}

	buf := ioutil.Discard

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chosen, v, ok := reflect.Select(cases)
		if !ok {
			b.Errorf("prematurely closed channel")
		}

		switch chosen {
		case 0:
			fmt.Fprintf(buf, "received %d\n", v.Int())
		case 1:
			fmt.Fprintf(buf, "received %f\n", v.Float())
		case 2:
			fmt.Fprintf(buf, "receibved %c\n", byte(v.Uint()))
		}
	}
}

func BenchmarkRawSelect(b *testing.B) {
	b.StopTimer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch1, ch2, ch3 := setupChannels(ctx)
	buf := ioutil.Discard

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		select {
		case v := <-ch1:
			fmt.Fprintf(buf, "received %d\n", v)
		case v := <-ch2:
			fmt.Fprintf(buf, "received %f\n", v)
		case v := <-ch3:
			fmt.Fprintf(buf, "receibved %c\n", v)
		}
	}
}

