/**
 * Subtractive Color Wheel 
 * by Ira Greenberg. 
 * 
 * The primaries are red, yellow, and blue. The secondaries are green, 
 * purple, and orange. The tertiaries are  yellow-orange, red-orange, 
 * red-purple, blue-purple, blue-green, and yellow-green.
 * 
 * Create a shade or tint of the subtractive color wheel using
 * SHADE or TINT parameters.
 *
 * Updated 26 February 2010.
 porting by crazylion
 */

package main
import (
    "fmt"
    "math"
    "image/color"
    "../../../goframework"
)
/*var width,height int = 200,200
var barWidth int =5
var lastBar int = -1
var coswave []float64 =make([]float64,width )*/

var width,height int = 200,200
var segs,steps int = 12,6
var rotAdjust float64 = goframework.TwoPi / float64(segs) / 2;
var radius,segWidth float64;
var interval float64 = goframework.TwoPi / float64(segs);
func drawShadeWheel(){
    var cols  []color.RGBA 
    for  j:= 0; j < steps; j++ {
        cols = []color.RGBA{ color.RGBA{uint8(255-(255/steps)*j), uint8(255-(255/steps)*j), 0,0} ,
                    color.RGBA{uint8(255-(255/steps)*j), uint8((255/1.5)-((255/1.5)/steps)*j), 0,0} ,
                    color.RGBA{uint8(255-(255/steps)*j), uint8((255/2)-((255/2)/steps)*j), 0,0} ,
                    color.RGBA{uint8(255-(255/steps)*j), uint8((255/2.5)-((255/2.5)/steps)*j), 0,0} ,
                    color.RGBA{uint8(255-(255/steps)*j), 0, 0,0} ,
                    color.RGBA{uint8(255-(255/steps)*j), 0,uint8((255/2)-((255/2)/steps)*j),0} ,
                    color.RGBA{uint8(255-(255/steps)*j), 0, uint8(255-(255/steps)*j),0} ,
                    color.RGBA{uint8((255/2)-((255/2)/steps)*j), 0, uint8(255-(255/steps)*j),0} ,
                    color.RGBA{0, 0, uint8(255-(255/steps)*j),0} ,
                    color.RGBA{0, uint8(255-(255/steps)*j), 0,0} ,
                    color.RGBA{0, uint8( 255-(255/steps)*j), 0,0} ,
                    color.RGBA{uint8((255/2)-((255/2)/steps)*j), uint8(255-(255/steps)*j), 0,0} }

        for i:= 0; i < segs; i++ {
/*             for i:= 0; i < 1; i++ { */
            goframework.Fill(cols[i].R,cols[i].G,cols[i].B)
            goframework.Stroke(cols[i].R,cols[i].G,cols[i].B)
/*             fmt.Println("fill",cols[i]) */
/*             goframework.Rect(40,40,50,50) */
            goframework.Arc(width/2, height/2,int(radius),int(radius), interval*float64(i)+rotAdjust, interval*float64(i+1)+rotAdjust) 
/*             goframework.Arc(width/2, height/2,int(radius),int(radius), interval*float64(i)+rotAdjust, interval*float64(i+1)+rotAdjust)  */
/*             fmt.Println("draw",width/2,height/2,radius,interval*float64(i)+rotAdjust, interval*float64(i+1)+rotAdjust) */
        }
         radius -= segWidth;
    }
}
func Draw(){
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(127)
/*     goframework.NoStroke() */
    radius = math.Min(float64(width),float64(height)) * 0.45
    segWidth = radius / float64(steps)
    goframework.EllipseMode(goframework.RADIUS)
    drawShadeWheel()

}
func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
