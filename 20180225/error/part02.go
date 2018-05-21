package main

import "fmt"

func main() {
	result, err := Divide(100, 10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

	result, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

type DivideError struct {
	dividee int
	divider int
}

/*
自定义错误
 */
func (d *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, d.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg error) {
	if varDivider == 0 {
		dData := DivideError{varDividee, varDivider}
		errorMsg = &dData
		return
	} else {
		return varDividee / varDivider, errorMsg
	}
}
