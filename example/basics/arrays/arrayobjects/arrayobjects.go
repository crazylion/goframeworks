package main
import (
    "fmt"
    "../../../../goframework"
    "time"
)
var width,height int = 320,240
var Unit,count int=40,0
var mods []Module
var totalCount int
func Draw(){
    for i:=0;i<totalCount;i++{
        v := &mods[i]
        v.Update()
        x,y := v.Position()
        goframework.StrokeWeight(2);
        t:=time.Now()
        col:=int(t.Second()*4%255)
        goframework.Stroke(col);
        goframework.Point(x,y)
    }
}

func main(){
    fmt.Println("start")
    var widthCount int = width/Unit;
    var heightCount int= height/Unit;
    totalCount  = heightCount*widthCount
    mods = make([]Module,heightCount*widthCount)
    index :=0
    for y:=0;y<heightCount;y++ {
        for x:=0;x<widthCount;x++ {
            mods[index]=NewModule(x*Unit, y*Unit, float64(Unit/2), float64(Unit/2), goframework.Random(0.05, 0.8))
            index++
        }
    }
    goframework.Draw(Draw)
    goframework.Win(width,height)
}
