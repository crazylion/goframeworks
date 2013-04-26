package main
import (
    "fmt"
    "../goframework"
)
func Draw(){
/*     fmt.Println("draw") */
    goframework.Fill(0,255,255)
    if goframework.MousePressed() {
        goframework.Fill(0,0,0)
    }else{
        goframework.Fill(255,255,255)
    }
    goframework.Ellipse(goframework.MouseX(),goframework.MouseY(),80,80)

}
func main(){
    fmt.Println("start")
    goframework.FrameRate(60)
    goframework.Draw(Draw)
    goframework.Win(400,400)
}
