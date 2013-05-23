/*
 * All the color method here
 *
 */
package goframework
import (
    "fmt"
    "math"
    "image/color"
)
const RGB int=0
const HSB int=1
var colorRange1,colorRange2,colorRange3 float64 = 255,255,255
var currentColorMode int
func ColorMode(mode int,ranges ... interface{}){
    currentColorMode=mode
    if HSB==mode {
        colorRange1=360
        colorRange2=1
        colorRange2=1
    }else{
        //reset to back
        colorRange1=255
        colorRange2=255
        colorRange2=255
    }
    if ranges != nil && len(ranges)>0 {
        switch(len(ranges)){
            case 1: colorRange1 = anyToFloat64(ranges[0])
            case 2: colorRange1 = anyToFloat64(ranges[0]); colorRange2 = anyToFloat64(ranges[1])
            case 3:colorRange1 = anyToFloat64(ranges[0]); colorRange2 = anyToFloat64(ranges[1]);colorRange3 = anyToFloat64(ranges[2])
        }
    }
}


/**
 if colorMode ==HSB  translte x,y,z,a value to hsb
*/
func CountColor(ih,is,iv interface{}) (rr,rg,rb uint8){
    var r,g,b float64;
    h:=anyToFloat64(ih)
    s:=anyToFloat64(is)
    v:=anyToFloat64(iv)
    if currentColorMode == HSB {
        //mapping to range
        h= Map(h,0,colorRange1,0,360)
        s= Map(s,0,colorRange2,0,1)
        v= Map(v,0,colorRange3,0,1)

        c:= v*s
        h= h/60
        x:=c*(1-math.Abs( float64(math.Mod(h,2)  -1 )  ))
        hplus := math.Floor(h)
        switch(hplus){
            case 0:
                r=c
                g=x
                b=0
            case 1:
                r=x
                g=c
                b=0
            case 2:
                r=0
                g=c
                b=x
            case 3:
                r=0
                g=x
                b=c
            case 4:
                r=x
                g=0
                b=c
            case 5:
                r=c
                g=0
                b=x
        }
        m:=v-c
        rr =uint8((r+m)*255)
        rg =uint8((g+m)*255)
        rb =uint8((b+m)*255)
/*         fmt.Println("input:r=",ih,",",is,",",iv) */
/*         fmt.Println("result:r=",rr,",",rg,",",rb) */
    }else{
        //mapping to range
        h:= Map(h,0,colorRange1,0,255)
        s:= Map(s,0,colorRange2,0,255)
        v:= Map(v,0,colorRange3,0,255)
        //rgb
        rr=anyToUint8(h)
        rg=anyToUint8(s)
        rb=anyToUint8(v)
    }
    return

}

/*
 translate the params to color.RGBA

*/

func anyToUint8(any interface{}) uint8{
    switch any := any.(type) {
    default:
        fmt.Printf("unexpected type %T", any)       // %T prints whatever type t has
    case float64:
        return uint8(any)
    case uint8:
        return uint8(any)
    case int:
        return uint8(any)
    }
    return 0
}

func anyToColor(ir interface{}, gb []interface{})(r,g,b uint8){
    isColor :=false
    switch ir := ir.(type) {
    default:
        fmt.Printf("unexpected type %T", ir)       // %T prints whatever type t has
    case float64:
        r=uint8(ir)
    case uint8:
        r=uint8(ir)
    case int:
        r=uint8(ir)
    case color.RGBA:
        r= ir.R
        g=ir.G
        b=ir.B
        isColor=true
    }
    if !isColor {
        if gb!=nil {
            g=anyToUint8(gb[0])
            b=anyToUint8(gb[1])
        }else{
            g=r
            b=r
        }
    }

    return

}

func anyToFloat64(any interface{}) float64{
    switch any := any.(type) {
    default:
        fmt.Printf("unexpected type %T", any)       // %T prints whatever type t has
    case float64:
        return float64(any)
    case uint8:
        return float64(any)
    case int:
        return float64(any)
    }
    return 0
}
