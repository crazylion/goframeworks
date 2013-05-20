package goframework
import(
    "math"
    "math/rand"
)
var TwoPi  = math.Pi*2
func Map(value , start1, stop1, start2, stop2 float64) float64{
    return (value-start1)*(stop2 -start2) / (stop1-start1) + start2
}

/*Calculates the Euclidean distance between two points  */
func Dist(x1,y1,x2,y2 int) float64{
    return math.Sqrt(float64((x1-x2)* (x1-x2) + (y1-y2)*(y1-y2)))
}

func Random(low,hight float64) float64{
    return Map(rand.Float64(),0,1,low,hight)
}
