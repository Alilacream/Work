package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func Cal(t [][]int, a,b int) [][]int {
	return t[a][b]
}
func main() {
	if (len(os.Args)!= 3) {
		log.Fatal("Need three arguments")
	}
	line := os.Args[1]
	row := os.Args[2]
	line_int, err:= strconv.Atoi(line)
	row_int , err1:= strconv.Atoi(row)
	if err != nil  || err1 != nil{
		log.Fatal("Either the number of Row or Line is invalid")
	}
	scanner := bufio.NewScanner(os.Stdin)

	Table := make([][]int , line_int)
	for i := range Table{
		Table[i] = make([]int , row_int)
	}
	fmt.Println("Print the number for your matrice")
	for i:= 0 ; i <line_int;i++ {
		for j :=0; j<row_int;j++ {
			if scanner.Scan() {
				fmt.Printf("For Line <%d> - Row <%d>: ")
				numstr := scanner.Text()
				number , err := strconv.Atoi(numstr)
				if err != nil {
					log.Fatal("Invalid number")
				}
				Table[i][j] = number
			}
			fmt.Println()
		}	 
	}
	fmt.Println("Choose  what you want to calculate in this: 1-Determinant\n 2-Transpose\n 3-inverse\n 4-Con")
	choose := scanner.Text()
	choice , err := strconv.Atoi(choose)

	switch choice {
	case 1 :
		milieu : rows_int-1
		determinant := Table[1][1]*(Cal(Table),) // on progress 
		break
	}
	
	
}