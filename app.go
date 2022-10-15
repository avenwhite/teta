package main

import (
	"fmt"
	"sort"
	"strings"
)
func strIsExistMap (item rune,m *map[rune]int) bool{
	for i,_:=range *m {
		if i == item {
			return true
		}
	}
	return false
}
func strMutate(m *map[rune]int) string  {
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	res:=""
	for _, k := range keys {
		r := []rune(k)
		res=res+fmt.Sprint(k, (*m)[r[0]])
	}
	return res
}
func strCompress(str string)string {
	m:=make(map[rune]int)
	symArr:=[]rune(str)
	for _,itemin:=range symArr {
		if !strIsExistMap(itemin,&m) {
			for _,itemout:=range symArr {
				if itemout==itemin {
					m[itemout]= strings.Count(str, string(itemout))
				}
			}
		}
	}
	return strMutate(&m)
}
func main() {
	var str string = "zzzzcccUUUuu"
	fmt.Println(strCompress(str))
}
