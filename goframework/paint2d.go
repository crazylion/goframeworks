package goframework
import (
    "github.com/crazylion/go-ui/ui"
    "image/color"
    "reflect"
)

var _strokeWeight int = 1

func Fill(ir uint8, gb ...uint8){
    if Brush==nil {
        Brush=ui.NewBrush()
    }
    r:=translateToUint8(ir)
    g:=r
    b:=r
    if gb != nil {
        g=translateToUint8(gb[0])
        b=translateToUint8(gb[1])

    }
    Brush.SetStyle(ui.SolidPattern)
    Brush.SetColor(color.RGBA{r,g,b,0})
    Painter.SetBrush(Brush)
}

func Background(r uint8,g uint8,b uint8){
/*     Brush := ui.NewBrush() */
/*     Brush.SetStyle(ui.SolidPattern) */
/*     Brush.SetColor(color.RGBA{r,g,b, 255}) */
/*     Painter.SetBrush(Brush) */
    Painter.DrawRect(ui.Rect{0,0,windowWidth,windowHeight})
}

func StrokeWeight(weight int){
    _strokeWeight=weight
}

/**
allow int,uint,float
*/
func Stroke(ir interface{},gb ...interface{}){
    if Pen == nil {
        Pen = ui.NewPen()
    }
    r:=translateToUint8(ir)
    g:=r
    b:=r
    if gb != nil {
        g=translateToUint8(gb[0])
        b=translateToUint8(gb[1])

    }
    Pen.SetColor(color.RGBA{r,g,b,0 })
    Pen.SetWidth(_strokeWeight)
    Painter.SetPen(Pen)
}

// base 

func translateToUint8(a interface{}) uint8{
    t := reflect.TypeOf(a)
    kind := t.Kind()
    v :=reflect.ValueOf(a)
    if kind ==reflect.Uint8 {
        return uint8(v.Int())
    } else if kind == reflect.Float64 {
        return uint8(v.Float())
    }
    //assume a is a int
    return uint8(v.Int())
}

func Point(x1,y2 int){
    Painter.DrawPoint(ui.Point{x1,y2})
}

func Line(ix1,iy1,ix2,iy2 interface{}){
    x1:=int(translateToUint8(ix1))
    x2:=int(translateToUint8(ix2))
    y1:=int(translateToUint8(iy1))
    y2:=int(translateToUint8(iy2))
    Painter.DrawLine(ui.Point{x1,y1},ui.Point{x2,y2})
}

func Rect(x int , y int , width int, height int){
    Painter.DrawRect(ui.Rect{x,y,width,height})
}


func Ellipse(x int,y int,width int, height int){
    Painter.DrawEllipse(ui.Rect{x-width/2,y-height/2,width,height})
}
