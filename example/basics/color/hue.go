/**
 porting by crazylion
*/

package main
import (
    "fmt"
    "../../../goframework"
)

var width,height int = 200,200
var barWidth int = 5;
var hue []int;
func Draw(){
    j := 0;
    for i:=0; i<=(width-barWidth); i+=barWidth {
        if (goframework.MouseX() > i) && (goframework.MouseX() < i+barWidth) {
            hue[j] = goframework.MouseY()
        }
        goframework.Fill(hue[j],float64(height)/1.2,float64(height)/1.2)
        goframework.Rect(i, 0, barWidth, height)
        j++;
    }
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(0)
    goframework.ColorMode(goframework.HSB,360,height,width)
    hue = make([]int,width/barWidth)
    goframework.NoStroke()

}
func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
