package main

import (
	"fmt"
	"os"
)

func fileOpener() {
  file, err := os.OpenFile("myfile", os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    panic(err)
  }
  defer file.Close() 

  data := []byte("This is the text written after opening myfile\n")
  _, err = file.Write(data)
  if err != nil {
    panic(err)
  }
}

func main() {

  file, err := os.Create("myfile")
  if err != nil {
    panic(err)
  }
  defer file.Close() 
  initialData := []byte("Initial data in the file\n")
  _, err = file.Write(initialData)
  if err != nil {
    panic(err)
  }

  fileOpener()

  fmt.Println("Operations performed")
}
