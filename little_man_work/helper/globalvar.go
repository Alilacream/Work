package Mat 

import(
	"fmt"
)

var (
	a , b , c int
	d , e , f int
	g , h , i int
)

func assignABCtoI(m [][]int) {
	a = m[0][0]
    b = m[0][1]
    c = m[0][2]
    d = m[1][0]
    e = m[1][1]
    f = m[1][2]
    g = m[2][0]
    h = m[2][1]
    i = m[2][2]
}
func Printmatrice(m [][]int) {
	for _, row:= range m {
		for _, value := range row{
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func Determinant3x3(m [][]int) int {
		assignABCtoI(m)
		return a*(e*i - f*h) - b*(d*i - f*g) + c*(d*h - e*g)
}
// this function purpose is only for the 3x3 dimensionale matrices ONLY!!
func Con3x3(m [][]int, line,row int) [][]int {
	assignABCtoI(m)	// made it a two dimension,with range of lines = line 
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

