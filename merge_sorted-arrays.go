package main
import "fmt"
import "bufio"
import "os"
import "strings"

func merge(arr1 []string, arr2 []string) []string{
  var count1,count2 int
  count1=0
  count2=0
   var arr []string
   for (count1<len(arr1) && count2<len(arr2)){
    a := arr1[count1]
    b := arr2[count2]
    if(a < b){
      arr=append(arr,arr1[count1])
      count1++
    } else{
      arr=append(arr,arr2[count2])
      count2++
    }
  }
  for count1<len(arr1){
    arr=append(arr,arr1[count1])
    count1++
  }
  for count2<len(arr2){
    arr=append(arr,arr2[count2])
    count2++
  }
  return arr
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter first sorted array")
    scanner.Scan()
    str :=scanner.Text()
    arr1:= strings.Split(str," ")
    fmt.Println("Enter second sorted array")
    scanner.Scan()
    str=scanner.Text()
    arr2:= strings.Split(str," ")
    fmt.Println("Merging...")
    arr := merge(arr1,arr2)
    fmt.Println("Merged array is : ",arr)
}
