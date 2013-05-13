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
    "fmt"
    "../../../goframework"
)
var width,height int = 200,200
var barWidth int =5
var lastBar int = -1
var coswave []float64 =make([]float64,width )
func Draw(){
    for i:=0;i<100;i++ {
        for j:=0;j<100;j++ {
/*             goframework.Stroke(i,j,100) */
/*             goframework.Fill(i,j,100) */
/*             goframework.StrokeWeight(2) */
/*             goframework.Ellipse(i,j,1,1) */
/*             goframework.Point(j,i) */
         } 
     } 
/*             goframework.Fill(81,90,74) */
/*             goframework.StrokeWeight(2) */
/*             goframework.Ellipse(50,50,50,50) */
/*    whichBar  := int(goframework.MouseX() / barWidth)
    if whichBar != lastBar {
        barX := whichBar * barWidth;
        goframework.Fill(barX, 100, goframework.MouseY())
        goframework.Rect(barX, 0, barWidth, height);
        lastBar = whichBar
    }*/
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(0)
    goframework.ColorMode(goframework.HSB)
/*     goframework.CountColor(66,0.43,0.9) */
/*     goframework.CountColor(0,1,1) */
/*     goframework.Fill(66,0.43,0.90) */
/*     goframework.Fill(0,1,1) */
/*     goframework.Background(255) */
/*     goframework.NoStroke() */
/*     r,g,b:=goframework.CountColor(180,1,1) */
/*     r,g,b:=goframework.FromHSV(180,1,1) */
/*     fmt.Printf("rgb=%d,%d,%d",r,g,b) */

}
func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
