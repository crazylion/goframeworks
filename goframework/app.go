package goframework
import (
    "github.com/crazylion/go-ui/ui"
    "fmt"
    "image/color"
)
var w *ui.Widget
var interval int =30
var timerid int
var drawFunc func()
var painter *ui.Painter
var brush   *ui.Brush
var _mousePressed = false
var _mouseX,_mouseY int = 0,0

var windowWidth,windowHeight int = 400,400

/**
    main window 

*/
func Win(width int,height int){
    fmt.Println("create windows")
    ui.Main(func(){
        w = ui.NewWidget()
        w.SetSize(ui.Size{width, height})
        w.SetOpAquepaintevent(true)
        w.SetMouseTracking(true)
/*         w.SetAttr() */
        w.OnTimerEvent(func(e *ui.TimerEvent) {
            if e.TimerId() == timerid {
                w.Update()
            }
        })
        w.OnPaintEvent(func(e *ui.PaintEvent) {
            painter = ui.NewPainter()
            defer painter.Close()
            painter.Begin(w)
            if drawFunc != nil {
                drawFunc()
            }
            painter.End()
        })
        w.OnMousePressEvent(func(e *ui.MouseEvent){
            fmt.Println("press")
            _mousePressed=true
        })
        w.OnMouseMoveEvent(func(e *ui.MouseEvent){
            fmt.Println("move")
            _mouseX = e.Pos().X
            _mouseY = e.Pos().Y
            fmt.Println(e.Pos())
        })
        w.OnMouseReleaseEvent(func(e *ui.MouseEvent){
            _mousePressed=false
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
   drawFunc = draw 
}
func Rect(x int , y int , width int, height int){
    if brush!=nil {
        painter.SetBrush(brush) 
    }
    painter.DrawRect(ui.Rect{x,y,width,height})
}

func Fill(r uint8, g uint8, b uint8){
    if brush==nil {
        brush=ui.NewBrush()
    }
    brush.SetStyle(ui.SolidPattern)
    brush.SetColor(color.RGBA{r,g,b,255})
}

func Background(r uint8,g uint8,b uint8){
    brush := ui.NewBrush()
    brush.SetStyle(ui.SolidPattern)
    brush.SetColor(color.RGBA{r,g,b, 255})
    painter.SetBrush(brush)
    painter.DrawRect(ui.Rect{0,0,windowWidth,windowHeight})
}


/** shape  **/

func Ellipse(x int,y int,width int, height int){
    if brush!=nil {
        painter.SetBrush(brush) 
    }
    
    painter.DrawEllipse(ui.Rect{x-width/2,y-height/2,width,height})
}

/* mouse event  */

func MousePressed() bool{
   return _mousePressed
}

func MouseX() int {
    return _mouseX
}

func MouseY() int {
    return _mouseY
}

