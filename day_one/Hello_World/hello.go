package main

import "fmt"

func Hello(name string) string  {
  if name==""{
    return "Hello, World"
  }else{
    return "Hello, "+name
  }
}

func main()  {
  fmt.Println(Hello("Enrinal"))
  fmt.Println(Hello(""))
}
