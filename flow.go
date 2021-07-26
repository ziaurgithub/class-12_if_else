package main
import "fmt"
func main(){

/*
fmt.Print("enter your age: ")
var age int
fmt.Scanf("%d", &age)
*/

/* 
0-2 = infant
   2-12 = children
   12-19 = teen age
   19 = adult 
*/

/* 

                               // if else
if age < 3 {
fmt.Println("infant")


}else if age >2 && age <13 {
fmt.Println("children")

}else if age >12 && age <= 19 {
fmt.Println("teen age")

}else{
fmt.Println("adult")
}
*/


                              //switch case
                              
 /*                             
switch age {
case 2:
fmt.Println("infant")
fallthrough
case 3,4,5,6,7,8,9,10,11,12:
fmt.Println("children")
fallthrough
case 13,14,15,16,17,18,19:
fmt.Println("teen age")
fallthrough
default:
fmt.Println("adult")
}
*/

                                   //for loop
/*
for i:=1; i<=9; i++ {
fmt.Println(i)
}
*/

                                   //for range loop

/*
students := []string{"ziaur","zahid","rafia"}
for i, val := range students {
fmt.Println(i, val)
}
*/

                                   //While loop
/*                                  
for true{
fmt.Println("ziaur")
}
*/
i:=0
for i<50 {
fmt.Println(i,"ziaur")
i++
}





}


//< >