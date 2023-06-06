package main
func main(){
  strings := []string{"hello", "world"}
  for i, s := range strings {
	  fmt.Println(i, s)
  }
}
