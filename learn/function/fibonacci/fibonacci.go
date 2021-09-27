package fibonacci

import "fmt"

func main() {
	result := 0
	for i := 0;i<=10;i++{
		result = fb(i)
		fmt.Printf("fibonacci (%d) is :%d\n",i,result)
	}
}

func fb(nr int) ( res int) {
	if nr <= 1 {
		res = 1
	}else {
		res = fb(nr -1) + fb(nr -2)
	}
	return
}
