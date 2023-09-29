package main

import "fmt"

const waterBoilingPointF float64 = 373.2

func main() {
	waterBoilingPointC := waterBoilingPointF - 273

	fmt.Printf("Ponto de ebulição da água em graus Celsius é %.2f", waterBoilingPointC)
}
