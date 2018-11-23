package main
import (
        "fmt";
        "math";
)
func main() {
        var sum float64;
        for i:=0; i<1E6; i++ {
                sum += math.Sin(float64(i)/0.1);
                fmt.Println(i,sum);
        }
        fmt.Println(sum);
}
