/**
 * source from processing
 * Brightness 
 * by Rusty Robison. 
 * 
 * Brightness is the relative lightness or darkness of a color.
 * Move the cursor vertically over each bar to alter its brightness. 
 * 
 * Updated 28 February 2010.
 * porting by crazylion
 */
package main

import (
	"../../../goframework"
	"fmt"
)

var width, height int = 200, 200
var barWidth int = 5
var lastBar int = -1
var coswave []float64 = make([]float64, width)

func Draw() {
	whichBar := int(goframework.MouseX() / barWidth)
	if whichBar != lastBar {
		barX := whichBar * barWidth
		goframework.Fill(barX, 100, goframework.MouseY())
		goframework.Rect(barX, 0, barWidth, height)
		lastBar = whichBar
	}
}

func Setup() {
	goframework.FrameRate(60)
	goframework.ColorMode(goframework.HSB)
	goframework.Background(0)

}
func main() {
	fmt.Println("start")
	goframework.Setup(Setup)
	goframework.Draw(Draw)
	goframework.Win(width, height)
}
