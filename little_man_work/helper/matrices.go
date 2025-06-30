package Mat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Matrice() {
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
	fmt.Println("Print the number for your matrice:")
	for i:= 0 ; i <line_int;i++ {
		for j :=0; j<row_int;j++ {
			fmt.Printf("For Line <%d> - Row <%d>: ", i , j)
			if scanner.Scan() {
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
	fmt.Println("Choose  what you want to calculate in this:\n 1-Determinant\n 2-Transpose\n 3-inverse\n 4-Con\n 5-Read_Matrice ")
	fmt.Print("Your choice: ")
	if scanner.Scan() {
		
	choose := scanner.Text()
	choice , err := strconv.Atoi(choose)
		if err != nil {
			log.Fatal("Invalid Input number")
		}
			switch choice {

			case 1 :
			determinant := Determinant3x3(Table) 
			fmt.Printf("The Determinant of the matrice is %d\n", determinant)

			case 2:
				transpose := Transpose3x3(Table)
				fmt.Println("Le Transpose de matrice est:")
				Printmatrice(transpose)
			case 3:
				Coematrice_transopose := Transpose3x3((Con3x3(Table,line_int,row_int)))
				determinant := (Determinant3x3(Table))
				if determinant == 0 {
					log.Fatal("The determinant is equal to 0, this matrix cannot be reversed")
				}
				determinant_str := strconv.Itoa(determinant)
				fmt.Println("The Inverse:")
				fmt.Printf("%s\n", determinant_str)
				fmt.Println("Mutliplied by:")
				Printmatrice(Coematrice_transopose)
			case 4:
			Table_con := Con3x3(Table, line_int, row_int)
			fmt.Println("The Coematrice of the Table:")
			Printmatrice(Table_con)

			case 5:
				fmt.Println("The Matrice you assigned:")
				Printmatrice(Table)

			default:
			fmt.Println("Not implemented yet")
			
			}
	}	
}