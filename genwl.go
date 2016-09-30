package main

import "fmt"
import "os"
import "strings"
import "math/rand"
//import "time"

func main() {
	argsWithoutProg := os.Args[1:]
        var str string = argsWithoutProg[0]
	var wl string = ""
	v := 0
	for e,_ := range str {
		for f,_ := range str {
			for g,_ := range str {
				for h,_ := range str {
					for i,_ := range str {
						for j,_ := range str {
							for k, elem := range str {
								//ra := rand.New(rand.NewSource(time.Now().UnixNano()+int64(g)))
								x := rand.Intn(10)
								v++
								if k < i {
									if k > j {
										if x > 5 {
											wl += strings.ToUpper(string(elem))
											x=g
											x=h
											x=f
											x=e
										}else{								
											wl += string(elem)
										}
									}else{
										y := rand.Intn(10)
										if y > 5 {
											wl += strings.ToUpper(string(elem))
										}else{								
											wl += string(elem)
										}
									}
								}else{
									z := rand.Intn(10)
									if z > 0 {
										wl += strings.ToUpper(string(elem))
									}else{								
										wl += string(elem)
									}
								}
							}
							fmt.Printf("%s\n", wl)
							wl = ""
						}
					}
				}
			}
		}
	}
}

