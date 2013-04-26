package goframework
import (
    "github.com/visualfc/go-ui/ui"
    "fmt"
    "image/color"
)
var w *ui.Widget
var interval int =30
var timerid int
var drawFunc func()
var painter *ui.Painter
var brush   *ui.Brush

var windowWidth,windowHeight int = 400,400

func Win(width int,height int){
    fmt.Println("create windows")
    ui.Main(func(){
        w = ui.NewWidget()
        w.SetSize(ui.Size{width, height})
        w.OnTimerEvent(func(e *ui.TimerEvent) {
            if e.TimerId() == timerid {
                w.Update()
                fmt.Println("on timer")
            }
        })
        w.OnPaintEvent(func(e *ui.PaintEvent) {
            fmt.Println("print") 
            painter = ui.NewPainter()
            defer painter.Close()
            painter.Begin(w)
            drawFunc()
/*             painter.DrawRect(ui.Rect{10, 10, 100, 100}) */
            painter.End()
        })
        w.Show()
        timerid = w.StartTimer(interval)
        ui.Run()
    })

}

func FrameRate(rate int){
    interval=1000/rate
}

func Draw(draw func()){
    fmt.Println("set draw")
   drawFunc = draw 
}
func Rect(x int , y int , width int, height int){
    painter.DrawRect(ui.Rect{x,y,width,height})
}

func Fill(r int, g int, b int){
    if brush {
     
    }else{
        brush=ui.NewBrush()
    }
}

func Background(r uint8,g uint8,b uint8){
    brush := ui.NewBrush()
    brush.SetStyle(ui.SolidPattern)
    brush.SetColor(color.RGBA{r,g,b, 255})
    painter.SetBrush(brush)
    painter.DrawRect(ui.Rect{0,0,windowWidth,windowHeight})
}

