package main

import (
	"fmt"
	"math/rand"
  "sort"
	"bufio"
  "os"
  "time"
  "strconv"
  "strings"
)
func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  
  game_total := read_total_arrays()

  for i := 0; i < game_total; i++ {
    fmt.Printf("Jogo %v \n", i + 1)

    game_list := create_random_list()
    
    time.Sleep(1 * time.Second)

    fmt.Println(game_list)
  }
}

func create_random_list() []int {
  var list []int 
  
  countNumber := 0

  for {
    num := rand.Intn(25)
    if search_number(num, list) == false{
      if num != 0 {
        list = append(list,num)
        countNumber+=1
      }
    }
    if countNumber >= 15 {
      break
    }
  }  
  sort.Ints(list)
  
  return list
}

func search_number(number int, argument_list []int) bool{
  for _, item := range argument_list {
    if item == number {
      return true
    }
  }
  return false
}

func read_total_arrays() int{
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("SORTEADOR LOTOFÁCIL, por João Pedro Resende.")

  fmt.Println("Quantos jogos você quer gerar? ")
  fmt.Print(">_ ")
  text, _ := reader.ReadString('\n')
  // convert CRLF to LF
  text = strings.Replace(text, "\n", "", -1)


  total, err := strconv.Atoi(text)

  if err != nil {
    fmt.Println("Por favor, coloque um valor válido.")
  }

  return total
}
