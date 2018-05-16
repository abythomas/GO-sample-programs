package main

import "fmt"


func swap(x int , y int) (int ,int){
  return y,x
}

func partition(arr []int, l int, r int) int{
  pivot := arr[r]
  i := l-1
  for k := l ; k <= r-1 ;k++ {
    if(arr[k]<pivot){
      i=i+1
      arr[k],arr[i]=swap(arr[k],arr[i])
    }
  }
  arr[i+1],arr[r]=swap(arr[i+1],arr[r])
  return i+1
}

func qsort(arr []int, l int, r int){
  if(l<r){
    p := partition(arr,l,r)
    qsort(arr,l,p-1)
    qsort(arr,p+1,r)
  }
}

func main(){
  var arr =[]int{4,1,6,24,7,89,32,90,0,11}
  qsort(arr,0,len(arr)-1)
  fmt.Println(arr)
}
