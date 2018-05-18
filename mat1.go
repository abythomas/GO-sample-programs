package main

import "fmt"
import "bufio"
import "os"

func main(){
  var mat =[5][5]string{{"A","B","C","D","E"},{"F","G","H","I","J"},{"K","L","M","N","O"},
           {"P","Q","R","S","T"},{"U","V","W","X","Y"}}
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Enter current value :")
  scanner.Scan()
  curr:=scanner.Text()
  fmt.Println("Enter blocked value :")
  scanner.Scan()
  bloc:=scanner.Text()

  var mov int
  mov=0
  var px,qx,py,qy int
  px=-2
  py=-2
  qx=6
  qy=6

  for p:=0;p<5;p++{
    for q:=0;q<5;q++{
      if(mat[p][q]==curr){
        px=p
        py=q
      }
      if(mat[p][q]==bloc){
        qx=p
        qy=q
      }
    }
  }

  if(px==-2){
    return
  }
  if(px>0 && py>0 && px<4 && py<4){
    mov=8
  }else if((px+py)%4==0 && (px==0 || py==0 || px==4 || py==4)){
    mov=3
  }else{
    mov=5
  }
  if(((px-qx)==1||(qx-px)==1)&&((py-qy)==0||(py-qy)==1||(qy-py)==1)){
    mov=mov-1
  }else if(((py-qy)==1||(qy-py)==1)&&((px-qx)==0||(px-qx)==1||(qx-px)==1)){
    mov=mov-1
  }
  fmt.Println(mov)
}
