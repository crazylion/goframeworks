package goframework
import (
    "github.com/crazylion/go-ui/ui"
    "image/color"
)

func Fill(r uint8, g uint8, b uint8){
    if Brush==nil {
        Brush=ui.NewBrush()
    }
    Brush.SetStyle(ui.SolidPattern)
    Brush.SetColor(color.RGBA{r,g,b,255})
}

func Background(r uint8,g uint8,b uint8){
    Brush := ui.NewBrush()
    Brush.SetStyle(ui.SolidPattern)
    Brush.SetColor(color.RGBA{r,g,b, 255})
    Painter.SetBrush(Brush)
    Painter.DrawRect(ui.Rect{0,0,windowWidth,windowHeight})
}

func Stroke(r uint8,gb ...uint8){
    Pen := ui.NewPen()
    g:=r
    b:=r
    if gb != nil {
        g=gb[0]
        b=gb[1]

    }
    Pen.SetColor(color.RGBA{r,g,b, 0})
    Painter.SetPen(Pen)
}

func Line(x1,y1,x2,y2 int){
    Painter.DrawLine(ui.Point{x1,y1},ui.Point{x2,y2})
}

func Rect(x int , y int , width int, height int){
    if Brush!=nil {
        Painter.SetBrush(Brush) 
    }
    Painter.DrawRect(ui.Rect{x,y,width,height})
}


func Ellipse(x int,y int,width int, height int){
    if Brush!=nil {
        Painter.SetBrush(Brush)
    }
    Painter.DrawEllipse(ui.Rect{x-width/2,y-height/2,width,height})
}
