package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "log"

func getMoves(mat [5][5]string,curr string, bcount int, bloc []string) int{
  var px,py int
  px=-2
  py=-2
  var qx []int
  var qy []int
  var intMat =[5][5]int{{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0}}

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
        intMat[p][q]=1
      }
    }
    }
  }

  fmt.Println(qx,qy)

  if(px==-2){
    return 0
  }

  var fin int
  fin=0
  if(px>0 && py>0 && px<4 && py<4){
    fin=0
  }else if((px+py)%4==0 && (px==0 || py==0 || px==4 || py==4)){

    fin=5
  }else{

    fin=3
  }
  var ans int
  ans=0
  var tempx,tempy int

    tempx=px+1
    var flag int
    flag=0
    for (tempx<=4){
      if(intMat[tempx][py]==1){
        flag=1
        break
      }
      tempx++
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempx=px-1
    for (tempx>=0){
      if(intMat[tempx][py]==1){
        flag=1
        break
      }
      tempx--
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempy=py+1
    for (tempy<=4){
      if(intMat[px][tempy]==1){
        flag=1
        break
      }
      tempy++
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempy=py-1
    for (tempy>=0){
      if(intMat[px][tempy]==1){
        flag=1
        break
      }
      tempy--
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempx=px+1
    tempy=py+1
    for (tempy<=4 && tempx<=4){
      if(intMat[tempx][tempy]==1){
        flag=1
        break
      }
      tempx++
      tempy++
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempx=px+1
    tempy=py-1
    for (tempy>=0 && tempx<=4){
      if(intMat[tempx][tempy]==1){
        flag=1
        break
      }
      tempx++
      tempy--
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempx=px-1
    tempy=py+1
    for (tempy<=4 && tempx>=0){
      if(intMat[tempx][tempy]==1){
        flag=1
        break
      }
      tempx--
      tempy++
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
    tempx=px-1
    tempy=py-1
    for (tempy>=0 && tempx>=0){
      if(intMat[tempx][tempy]==1){
        flag=1
        break
      }
      tempx--
      tempy--
    }
    if(flag==0){
      ans=ans+1
    }
    flag=0
  return ans-fin
}

func main(){
  var mat =[5][5]string{{"A","B","C","D","E"},{"F","G","H","I","J"},{"K","L","M","N","O"},
           {"P","Q","R","S","T"},{"U","V","W","X","Y"}}
  scanner := bufio.NewScanner(os.Stdin)
  var bloc []string
  var bcount int
  bcount =0
  for t:=0;t<3;t++{
  fmt.Println("Enter current value :")
  scanner.Scan()
  curr:=scanner.Text()
  fmt.Println("Enter blocked count :")
  scanner.Scan()
  bcount1,err := strconv.Atoi(scanner.Text())
  if err!=nil{
    log.Println("error")
  }
  bcount=bcount+bcount1
  for t:=0;t<bcount1;t++{
  fmt.Println("Enter blocked value :")
  scanner.Scan()
  bloc =append(bloc,scanner.Text())
  }
  i := getMoves(mat,curr,bcount,bloc)
  fmt.Println(i)
  }
}
