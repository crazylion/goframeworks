package goframework
import (
    "github.com/crazylion/go-ui/ui"
    "fmt"
)
var w *ui.Widget
var interval int =30
var timerid int
var drawFunc func()
var setupFunc func() = nil
var Painter *ui.Painter
var painter *ui.Painter
var Brush   *ui.Brush
var Pen     *ui.Pen
var _mousePressed = false
var _mouseX,_mouseY int = 0,0
var isInit bool = false
var mainCanvas *ui.Image  //main canvas

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
        mainCanvas:= ui.NewImageWithSize(width,height)
        defer mainCanvas.Close()
        Painter= ui.NewPainterWithImage(mainCanvas)
/*         w.SetAttr() */
        w.OnTimerEvent(func(e *ui.TimerEvent) {
            if e.TimerId() == timerid {
                if !isInit {
                    isInit=true
                    if setupFunc != nil {
                       setupFunc() 
                    }
                }
                if drawFunc != nil {
                    drawFunc()
                }
                w.Update()
            }
        })
        w.OnPaintEvent(func(e *ui.PaintEvent) {
            painter = ui.NewPainter()
            defer painter.Close()
            painter.Begin(w)
                painter.DrawImageEx(ui.Point{0, 0}, mainCanvas, mainCanvas.Rect())
            painter.End()
        })
        w.OnMousePressEvent(func(e *ui.MouseEvent){
            _mousePressed=true
        })
        w.OnMouseMoveEvent(func(e *ui.MouseEvent){
            _mouseX = e.Pos().X
            _mouseY = e.Pos().Y
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

func Setup(setup func()){
    setupFunc = setup
}


/** shape  **/


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

