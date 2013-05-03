package main

type Module struct {
    mx,my int
    x,y,speed,big float64
    xdir,ydir int
}

func NewModule(mx ,my int,x,y,speed float64 ) Module{
    return  Module{mx:mx,my:my,x:x,y:y,speed:speed,big:float64(Unit),xdir:1,ydir:1}
}


func (m *Module)Update(){
    m.x += (m.speed*float64(m.xdir))
    if m.x>= m.big  || m.x<=0 {
        m.xdir *=-1
        m.x=m.x+float64(m.xdir)
        m.y=m.y+float64(m.ydir)
    }

    if m.y>=m.big || m.y<=0 {
        m.ydir *= -1
       m.y=m.y+float64(m.ydir)
   }
}

func (m *Module)Position() (rx,ry int){
    rx = m.mx+int(m.x)-1
    ry = m.my+int(m.y)-1
    return 
}
