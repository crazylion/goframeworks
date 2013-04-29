package main
import (
    "fmt"
    "math"
    "../../../goframework"
)
var width,height int = 200,200
var coswave []float64 =make([]float64,width )
func Draw(){
    for i:=0; i < width; i++ {
        var amount = goframework.Map(float64(i), 0, float64(width), 0, math.Pi)
          coswave[i] = math.Abs(math.Cos(amount))
      }
      for i := 0; i < width; i++ {
          goframework.Stroke(uint8(coswave[i]*255))
          goframework.Line(i, 0, i, height/3)
      }
      for i := 0; i < width; i++ {
          goframework.Stroke(uint8(coswave[i]*255 / 4))
          goframework.Line(i, int(height/3), i,int(height/3*2))
      }
      for i := 0; i < width; i++ {
          goframework.Stroke(uint8(255 - coswave[i]*255))
          goframework.Line(i, height/3*2, i, height)
      }
  }
  func main(){
      fmt.Println("start")
      goframework.FrameRate(60)
      goframework.Draw(Draw)
      goframework.Win(width,height)
  }
