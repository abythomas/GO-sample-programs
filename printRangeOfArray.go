package main
import "fmt"


func main(){

  var i []int
  var start,end int
  fmt.Println(i,"\nLen : ",len(i),"\nCapacity : ",cap(i),"\n")
  i=append(i,1)
  fmt.Println(i,"\nLen : ",len(i),"\nCapacity : ",cap(i),"\n")
  i=append(i,2)
  fmt.Println(i,"\nLen : ",len(i),"\nCapacity : ",cap(i),"\n")
  i=append(i,3)
  fmt.Println(i,"\nLen : ",len(i),"\nCapacity : ",cap(i),"\n")
  i=append(i,4,5)
  fmt.Println(i,"\nLen : ",len(i),"\nCapacity : ",cap(i),"\n")

  for{
  fmt.Println("Enter start index :")
  fmt.Scanf("%d",&start)
  if(start<0||start>=len(i)){
    fmt.Println("Invalid start index")
    continue
  }
    for{
      fmt.Println("Enter end index :")
      fmt.Scanf("%d",&end)
      if(end>=start && end<=len(i)){
        break
      }
      fmt.Println("Invalid end index")
    }
    break
  }
  fmt.Println(i[start:end])
}
