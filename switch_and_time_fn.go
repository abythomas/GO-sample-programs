package main
import "fmt"
import "time"

func main(){

  t := time.Now()
  fmt.Println("Current time : ", t)
  switch t.Weekday(){                      /*normal switch usage*/
  case time.Saturday, time.Sunday:
    fmt.Println("Weekend")
  default :
    fmt.Println("Weekday")
  }

  switch{                                 /*switch like if*/
  case t.Hour()>12 :
    fmt.Println("After noon")
  default :
    fmt.Println("Before noon")
  }

}
