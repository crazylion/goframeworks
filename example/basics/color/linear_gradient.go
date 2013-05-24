/**
 * Simple Linear Gradient 
 * by Ira Greenberg. 
 * 
 * Using the convenient red(), green() 
 * and blue() component functions,
 * generate some linear gradients.
 * porting by crazylion
*/

package main
import (
    "fmt"
    "../../../goframework"
    "image/color"
)

var width,height int = 200,200
var Y_AXIS int= 1;
var X_AXIS int= 2;

func Draw(){
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(0)
    b1 := color.RGBA{190, 190, 190,0}
    b2 := color.RGBA{20, 20, 20,0}
    setGradient(0, 0, float64(width),float64(height), b1, b2, Y_AXIS);
    c1 := color.RGBA{255, 120, 0,0}
    c2 := color.RGBA{10, 45, 255,0}
    c3 := color.RGBA{10, 255, 15,0}
    c4 := color.RGBA{125, 2, 140,0}
    c5 := color.RGBA{255, 255, 0,0}
    c6 := color.RGBA{25, 255, 200,0}
    setGradient(25, 25, 75, 75, c1, c2, Y_AXIS);
    setGradient(100, 25, 75, 75, c3, c4, X_AXIS);
    setGradient(25, 100, 75, 75, c2, c5, X_AXIS);
    setGradient(100, 100, 75, 75, c4, c6, Y_AXIS);

}

func setGradient(x float64, y float64,w float64,h float64,c1 color.RGBA,c2 color.RGBA,axis int){
 // calculate differences between color components 
 deltaR := float64(goframework.Red(c2)-goframework.Red(c1));
 deltaG := float64(goframework.Green(c2)-goframework.Green(c1));
 deltaB := float64(goframework.Blue(c2)-goframework.Blue(c1));
  if axis == Y_AXIS {
    /*nested for loops set pixels
     in a basic table structure */
    // column
    for i:=x; i<=(x+(w)); i++ {
      // row
      for  j := y; j<=(y+(h)); j++ {
          newR := uint8(float64(goframework.Red(c1))+(j-y)*(deltaR/h))
          newG := uint8(float64(goframework.Green(c1))+(j-y)*(deltaG/h))
          newB := uint8(float64(goframework.Blue(c1))+(j-y)*(deltaB/h))
          c := color.RGBA{newR,newG,newB,0}
          goframework.Set(i,j, c);
      }
    }
  } else if axis == X_AXIS {

      for i:=y; i<=(y+h); i++ {
      // row
      for j := x; j<=(x+w); j++ {
          newR := uint8(float64(goframework.Red(c1))+(j-x)*(deltaR/h))
          newG := uint8(float64(goframework.Green(c1))+(j-x)*(deltaG/h))
          newB := uint8(float64(goframework.Blue(c1))+(j-x)*(deltaB/h))
          c := color.RGBA{ newR,newG,newB,0}
          goframework.Set(j, i, c);
      }
    }

  
  
  }
}

func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
