package goframework
import (
    "github.com/crazylion/go-ui/ui"
    "image/color"
    "math"
)

var _strokeWeight int = 1
var penStyle = ui.SolidLine
var penColor = color.RGBA{0,0,0,0}
var brushColor =color.RGBA{0,0,0,0}

const(
    CENTER =1
    RADIUS =2
    CORNEOR =3
    CORNERS =4
)
var ellipseMode=CENTER



func Fill(ir interface{}, gb ...interface{}){
    if Brush==nil {
        Brush=ui.NewBrush()
    }
    r:=anyToUint8(ir)
    g:=r
    b:=r
    if gb != nil {
        g=anyToUint8(gb[0])
        b=anyToUint8(gb[1])

    }
    r,g,b=CountColor(r,g,b)
    Brush.SetStyle(ui.SolidPattern)
    Brush.SetColor(color.RGBA{r,g,b,0})
    Painter.SetBrush(Brush)
    setPen()
}

func Background(ir interface{},gb ... interface{}){
/*     Brush := ui.NewBrush() */
/*     Brush.SetStyle(ui.SolidPattern) */
/*     Brush.SetColor(color.RGBA{r,g,b, 255}) */
/*     Painter.SetBrush(Brush) */
    r:=int(anyToUint8(ir))
    g:=r
    b:=r
    if gb != nil {
        g=int(anyToUint8(gb[0]))
        b=int(anyToUint8(gb[1]))
    }
    Fill(r,g,b)
    Rect(0,0,windowWidth,windowHeight)
}

func StrokeWeight(weight int){
    _strokeWeight=weight
    setPen()
}

/**
allow int,uint,float
*/
func Stroke(ir interface{},gb ...interface{}){
    if Pen == nil {
        Pen = ui.NewPen()
    }
    r:=anyToUint8(ir)
    g:=r
    b:=r
    if gb != nil {
        g=anyToUint8(gb[0])
        b=anyToUint8(gb[1])

    }
    r,g,b = CountColor(r,g,b)
    penColor = color.RGBA{r,g,b,255 }
    setPen()
}

func NoStroke(){
    penStyle= ui.NoPen
    setPen()
}

func setPen(){
    if Pen == nil {
        Pen = ui.NewPen()
    }
    Pen.SetColor(penColor)
    Pen.SetStyle(penStyle)
    Pen.SetWidth(_strokeWeight)
    Painter.SetPen(Pen)
    Pen.SetStyle(penStyle)
}

// base 

func Point(x1,y2 int){
    Painter.DrawPoint(ui.Point{x1,y2})
}

func Line(ix1,iy1,ix2,iy2 interface{}){
    x1:=int(anyToUint8(ix1))
    x2:=int(anyToUint8(ix2))
    y1:=int(anyToUint8(iy1))
    y2:=int(anyToUint8(iy2))
    Painter.DrawLine(ui.Point{x1,y1},ui.Point{x2,y2})
}

func Rect(x int , y int , width int, height int){
    Painter.DrawRect(ui.Rect{x,y,width,height})
}


func Ellipse(x int,y int,width int, height int){
    Painter.DrawEllipse(ui.Rect{x-width/2,y-height/2,width,height})
}

func EllipseMode(mode int){
    ellipseMode=mode
}

/* start,stop  specified in radians */
func Arc(x int, y int, width int, height int, start float64, stop float64) {
    var startAngle int =  int(start *180 /math.Pi)
    var stopAngle int =  int(stop *180 /math.Pi ) - startAngle 
    startAngle*=-16
    stopAngle*=-16
    if ellipseMode== CORNERS {
        width-=x
        height-=y
    }else if ellipseMode == RADIUS{
        x-=width
        y-=height
        width*=2
        height*=2
    }else if ellipseMode ==CENTER{
        x-= width/2
        y-= height/2
    }
    Painter.DrawPie(ui.Rect{x,y,width,height},startAngle,stopAngle)
}
