package goframework

func Map(value , start1, stop1, start2, stop2 float64) float64{
    return (value-start1)*(stop2 -start2) / (stop1-start1) + start2
}
