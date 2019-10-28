package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)
var text string
var i,bases,count,max,min,medio, q20, q30 int
func main(){
	args := os.Args[1:]
	fileIn, err := os.Open(args[0])
	if err != nil{
		fmt.Println(err)
		return	
	}
	sc := bufio.NewScanner(fileIn)
	for sc.Scan(){
		text = sc.Text()	
		if count==1 {
			if i==1 {
				max = len(text)
				min = len(text)
			}else{
				if len(text) > max {
					max = len(text)
				}
				if len(text) < min {
					min = len(text)
				}
			}
			bases = bases+len(text)
		}
		if count == 3 {
			for _, r := range text{
				if r > 20{
					q20 = q20+1
				}
				if r > 30{
					q30 = q30+1
				}
			}
			count= -1
		}
		count++
		i++
	}
	total := i/4
	medio = bases/total
	fileIn.Close()
	fileOut, err := os.Create(args[1])	
	result := strconv.Itoa(total)+"\t"+strconv.Itoa(bases)+"\t"+strconv.Itoa(q30)+"\t"+strconv.Itoa(q20)+"\t"+strconv.Itoa(medio)+"\t"+strconv.Itoa(max)+"\t"+strconv.Itoa(min)
	n,_ := fileOut.WriteString(result)
	fmt.Println("wrote ",n)
}

#####
