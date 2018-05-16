package main
import "fmt"
import "bufio"
import "os"

func main(){
   m :=make(map [string]int)
   m["Real Madrid"]=12
   m["Barcelona"]=5
   m["AC Milan"]=7
   scanner := bufio.NewScanner(os.Stdin)
   fmt.Println("Enter 'Refalona' to exit :")
   var str string
   for{
     fmt.Println("Enter club name :")
     scanner.Scan()
     str=scanner.Text()
     if(str=="Refalona"){
       break
     }
     a,b :=m[str]
     if(!b){
       fmt.Println("Refalona Refalon Refalo Refal Refa Ref Re R")
       continue
     }

     fmt.Println(str," has ",a," UCL trophies")

   }
}
