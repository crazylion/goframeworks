/*
 * All the color method here
 *
 */
package goframework
import (
    "fmt"
    "math"
)
const RGB int=0
const HSB int=1 
var currentColorMode int
func ColorMode(mode int){
    currentColorMode=mode 
}


/**
 if colorMode ==HSB  translte x,y,z,a value to hsb
*/
func CountColor(ih,is,iv interface{}) (r,g,b uint8){
    if currentColorMode == HSB {
        h:=anyToFloat64(ih)
        s:=anyToFloat64(is)
        v:=anyToFloat64(iv)
        h /= 60.0
        i := math.Floor(h)
        f := h - i
        p := v*(1 - s)
        q := v*(1 - s * f)
        t := v*(1 - s * (1 - f))
        switch(i) {
        case 0:
            r = uint8(v)
            g = uint8(t)
            b = uint8(p)
            break
        case 1:
            r = uint8(q)
            g = uint8(v)
            b = uint8(p)
            break
        case 2:
            r = uint8(p)
            g = uint8(v)
            b = uint8(t)
            break
        case 3:
            r = uint8(p)
            g = uint8(q)
            b = uint8(v)
            break
        case 4:
            r = uint8(t)
            g = uint8(p)
            b = uint8(v)
            break
        default:// case 5:
            r = uint8(v)
            g = uint8(p)
            b = uint8(q)
            break
        }
    }else{
        //rgb
        r=anyToUint8(r)
        g=anyToUint8(g)
        b=anyToUint8(b)
    }
    return

}

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
