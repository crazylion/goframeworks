package main
import (
    "fmt"
    "../../../goframework"
)
var width,height int = 200,200
var distances [][]float64 = make([][]float64,width*height)
func Draw(){
    for i:=0; i<height; i+=2 {
        for j:=0; j<width; j+=2 {
            goframework.StrokeWeight(2)
            goframework.Stroke(distances[j][i])
            goframework.Point(j, i)
        }
    }
}
func main(){
    fmt.Println("start")
    var maxDistance float64= goframework.Dist(width/2, height/2, width, height);
    for i:=  range distances{
        distances[i] =make([]float64,width)
    }
    for i:=0; i<height; i++ {
        for j:=0; j<width; j++ {
            dist := goframework.Dist(width/2, height/2, j, i)
            distances[j][i] = dist/maxDistance * 255
        }
    }
    goframework.FrameRate(1)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
