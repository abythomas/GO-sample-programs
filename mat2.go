package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "log"

func getMoves(mat [5][5]string,curr string, bcount int, bloc []string) int{
  var mov int
  mov=0
  var px,py int
  px=-2
  py=-2
  var qx []int
  var qy []int

  for t:=0;t<bcount;t++{
    qx=append(qx,6)
    qy=append(qy,6)
  }

  for p:=0;p<5;p++{
    for q:=0;q<5;q++{
      if(mat[p][q]==curr){
        px=p
        py=q
      }
      for t:=0;t<bcount;t++{
      if(mat[p][q]==bloc[t]){
        qx[t]=p
        qy[t]=q
      }
    }
    }
  }

  fmt.Println(qx,qy)

  if(px==-2){
    return 0
  }
  if(px>0 && py>0 && px<4 && py<4){
    mov=8
  }else if((px+py)%4==0 && (px==0 || py==0 || px==4 || py==4)){
    mov=3
  }else{
    mov=5
  }
  for t:=0;t<bcount;t++{
  if(((px-qx[t])==1||(qx[t]-px)==1)&&((py-qy[t])==0||(py-qy[t])==1||(qy[t]-py)==1)){
    mov=mov-1
  }else if(((py-qy[t])==1||(qy[t]-py)==1)&&((px-qx[t])==0||(px-qx[t])==1||(qx[t]-px)==1)){
    mov=mov-1
  }
  }
  return mov
}

func main(){
  var mat =[5][5]string{{"A","B","C","D","E"},{"F","G","H","I","J"},{"K","L","M","N","O"},
           {"P","Q","R","S","T"},{"U","V","W","X","Y"}}

  scanner := bufio.NewScanner(os.Stdin)
  for t:=0;t<3;t++{
  fmt.Println("Enter current value :")
  scanner.Scan()
  curr:=scanner.Text()
  fmt.Println("Enter blocked count :")
  scanner.Scan()
  bcount,err := strconv.Atoi(scanner.Text())
  if err!=nil{
    log.Println("error")
  }
  var bloc []string
  for t:=0;t<bcount;t++{
  fmt.Println("Enter blocked value :")
  scanner.Scan()
  bloc =append(bloc,scanner.Text())
  }
  i := getMoves(mat,curr,bcount,bloc)
  fmt.Println(i)
  }

}
