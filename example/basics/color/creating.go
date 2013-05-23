/**
 porting by crazylion
*/

package main
import (
    "fmt"
    "image/color"
    "../../../goframework"
)

var width,height int = 200,200
var inside  = color.RGBA{204,102,0,0}
var middle  = color.RGBA{204,153,0,0}
var outside = color.RGBA{153,51,0,0}
func Draw(){
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(0)
    goframework.NoStroke()
    goframework.Fill(outside)
    goframework.Rect(0,0,200,200)
    goframework.Fill(middle)
    goframework.Rect(40, 60, 120, 120)
    goframework.Fill(inside)
    goframework.Rect(60, 90, 80, 80)

}
func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
