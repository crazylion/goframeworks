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
func CountColor(ih,is,iv interface{}) (rr,rg,rb uint8){
    var r,g,b float64;
    if currentColorMode == HSB {
        h:=anyToFloat64(ih)
        s:=anyToFloat64(is)
        v:=anyToFloat64(iv)
        c:= v*s
        h= h/60
        x:=c*(1-math.Abs( float64(math.Mod(h,2)  -1 )  ))
        hplus := math.Floor(h)
        fmt.Println("hplus=",hplus)
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
        fmt.Println("m=",m)
        rr =uint8((r+m)*255)
        rg =uint8((g+m)*255)
        rb =uint8((b+m)*255)
        fmt.Println("input:r=",ih,",",is,",",iv)
        fmt.Println("result:r=",rr,",",rg,",",rb)
    }else{
        //rgb
        rr=anyToUint8(ih)
        rg=anyToUint8(is)
        rb=anyToUint8(iv)
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
