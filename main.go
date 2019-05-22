package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
	"time"
)

// getOperands extracts operands for calculator and converts to floats.
func getOperands(i []js.Value) (float64, float64, error) {
	s1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	s2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		return 0, 0, err
	}
	f2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return 0, 0, err
	}
	return f1, f2, nil
}

// calc performs arithmetic calculations.
func calc(_ js.Value, i []js.Value) interface{} {
	n1, n2, err := getOperands(i)
	if err != nil {
		js.Global().Get("document").Call("getElementById", "calcResult").Set("innerHTML", err.Error())
		return nil
	}
	result := 0.0
	switch i[2].Int() {
	case 1:
		result = n1 + n2
	case 2:
		result = n1 - n2
	case 3:
		result = n1 * n2
	case 4:
		result = n1 / n2
	}
	js.Global().Get("document").Call("getElementById", "calcResult").Set("innerHTML", result)
	return nil
}

// fib calculates the nth Fibonacci number.
func fib(n int) int {
	prev, result := 1, 1
	for i := 2; i < n; i++ {
		prev, result = result, result+prev
	}
	return result
}

// getFib uses fib on the input field's value, converted to int first.
func getFib(_ js.Value, i []js.Value) interface{} {
	s := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	n, err := strconv.Atoi(s)
	if err != nil {
		js.Global().Get("document").Call("getElementById", "fibResult").Set("innerHTML", err.Error())
		return nil
	}
	start := time.Now()
	result := fib(n)
	duration := time.Since(start)
	js.Global().Get("document").Call("getElementById", "fibResult").Set("innerHTML", result)
	js.Global().Get("document").Call("getElementById", "fibDuration").Set("innerHTML", duration.String())
	return nil
}

// fac calculates the factorial for n.
func fac(n int64) *big.Int {
	var f big.Int
	f.MulRange(1, n)
	return &f
}

// getFac uses fac on the input value, converted to int first.
func getFac(_ js.Value, i []js.Value) interface{} {
	s := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	n, err := strconv.Atoi(s)
	if err != nil {
		js.Global().Get("document").Call("getElementById", "facResult").Set("innerHTML", err.Error())
		return nil
	}
	start := time.Now()
	result := fac(int64(n))
	duration := time.Since(start)
	js.Global().Get("document").Call("getElementById", "facResult").Set("innerHTML", result.String())
	js.Global().Get("document").Call("getElementById", "facDuration").Set("innerHTML", duration.String())
	return nil
}

// registerCallbacks prepares the Go functions to be exposed to Javascript.
func registerCallbacks() {
	js.Global().Set("calc", js.FuncOf(calc))
	js.Global().Set("fib", js.FuncOf(getFib))
	js.Global().Set("fac", js.FuncOf(getFac))
}

func main() {
	// This prints in the JS console.
	fmt.Println("Go WASM! Initialized")

	// Fade-in animation via Go for loop.
	for i := 0.0; i <= 10.0; i++ {
		bs := js.Global().Get("document").Call("getElementById", "body").Get("style")
		bs.Set("opacity", i/10.0)
		<-time.After(30 * time.Millisecond)
	}

	// Setup Go to JS functions.
	registerCallbacks()

	// Wati forever.
	select {}
}
