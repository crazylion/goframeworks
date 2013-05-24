/**
 * Simple Radial Gradient 
 * by Ira Greenberg. 
 * 
 * Using the convenient red(), green() 
 * and blue() component functions,
 * generate an array of radial gradients.
 * porting by crazylion
 * very very very slow at this version
*/

package main
import (
    "fmt"
    "math"
    "../../../goframework"
    "image/color"
)

var width,height float64 = 200,200

func Draw(){
}

func Setup(){
    goframework.FrameRate(30)
    goframework.Background(0)
    columns := 4.0
    radius := float64((width/columns)/2)
    // create some gradients
    for i:=radius; i< float64(width); i+=float64(radius)*2 {
        for j :=radius; j< float64(height); j+=float64(radius)*2 {
            sColor := color.RGBA{uint8(goframework.Random(255)), uint8(goframework.Random(255)), uint8(goframework.Random(255)),0}
            eColor :=color.RGBA{uint8(goframework.Random(255)), uint8(goframework.Random(255)), uint8(goframework.Random(255)),0}
            createGradient(float64(i), float64(j), float64(radius),sColor ,eColor )
        }
    }
}

func createGradient(x float64,y float64,radius float64 ,c1 color.RGBA, c2 color.RGBA){
    var px,py,angle float64 = 0,0,0

  // calculate differences between color components 
  deltaR := float64(goframework.Red(c2)-goframework.Red(c1))
  deltaG := float64(goframework.Green(c2)-goframework.Green(c1))
  deltaB := float64(goframework.Blue(c2)-goframework.Blue(c1))
  // hack to ensure there are no holes in gradient
  // needs to be increased, as radius increases
  gapFiller := 8.0

  for i:=0.0; i< radius; i++ {
      for j:=0.0; j<360; j+=1.0/gapFiller {
      px = x+math.Cos(goframework.Radians(angle))*i;
      py = y+math.Sin(goframework.Radians(angle))*i;
      angle+=1.0/gapFiller;
      r:=uint8(float64(goframework.Red(c1))+(i)*(deltaR/radius))
      g:=uint8(float64(goframework.Green(c1))+(i)*(deltaG/radius))
      b:=uint8(float64(goframework.Blue(c1))+(i)*(deltaB/radius))
      c := color.RGBA{r,g ,b,0}
      goframework.Set(int(px), int(py), c);
    }
  }
  // adds smooth edge 
  // hack anti-aliasing
  goframework.NoFill();
  goframework.StrokeWeight(3);
  goframework.Ellipse(x, y, radius*2, radius*2);
}

func main(){
    fmt.Println("start")
    goframework.Setup(Setup)
    goframework.Draw(Draw)
    goframework.Win(int(width),int(height))
}
