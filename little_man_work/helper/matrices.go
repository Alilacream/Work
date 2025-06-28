package Mat 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func Printmatrice(m [][]int) {
	for _, row:= range m {
		for _, value := range row{
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func Determinant3x3(m [][]int) int {
	a := m[0][0]
	b := m[0][1]
	c := m[0][2]
	d := m[1][0]
	e := m[1][1]
	f := m[1][2]
	g := m[2][0]
	h := m[2][1]
	i := m[2][2]

	return a*(e*i - f*h) - b*(d*i - f*g) + c*(d*h - e*g)
}
// this function purpose is only for the 3x3 dimensionale matrices ONLY!!
func Con3x3(m [][]int, line,row int) [][]int {
	a := m[0][0]
	b := m[0][1]
	c := m[0][2]
	d := m[1][0]
	e := m[1][1]
	f := m[1][2]
	g := m[2][0]
	h := m[2][1]
	i := m[2][2]

	// making another matrice because i cannot change the values in the m directly , or asign it as a pointer, might also change the values of m
	// made it a two dimension,with range of lines = line 
	m_Con := make([][]int, line)
	//and asigned each row
	for i := range m_Con{
		m_Con[i] = make([]int, row)
	}
	// first line 
	m_Con[0][0] = (e*i - f*h)
	m_Con[0][1] = -(d*i - f*g)
	m_Con[0][2] = (d*h - e*g)
	// second line
	m_Con[1][0] = -(b*i-h*c)
	m_Con[1][1] = (a*i - c*g)
	m_Con[1][2] = -(a*h - b*g)
	// third line
	m_Con[2][0] = (b*f - c*e)
	m_Con[2][1] = -(a*f - c*d)
	m_Con[2][2] = (a*e - b*d)
	
 
	return m_Con
}


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
	fmt.Println("Print the number for your matrice")
	for i:= 0 ; i <line_int;i++ {
		for j :=0; j<row_int;j++ {
			if scanner.Scan() {
				fmt.Printf("For Line <%d> - Row <%d>: ", i , j)
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
	fmt.Println("Choose  what you want to calculate in this: 1-Determinant\n 2-Transpose\n 3-inverse\n 4-Con\n 5-Read_Matrice ")
	if scanner.Scan() {
	choose := scanner.Text()
	choice , err := strconv.Atoi(choose)
		if err != nil {
		log.Fatal("Invalid Input number")
			switch choice {
			case 1 :
			determinant := Determinant3x3(Table) 
			fmt.Printf("The determinant of the matrice is %d\n", determinant)
			case 4:
			Table_con := Con3x3(Table, line_int, row_int)
			fmt.Println("The con of the Table:")
			Printmatrice(Table_con)
			case 5:
				Printmatrice(Table)
		default:
			fmt.Println("Not implemented yet")
			}
		}
	}	
}