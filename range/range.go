package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}
	fmt.Println(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 30 {
			fmt.Println("found", num, "at index:", i)
		}
	}
	kvs:=map[string]string{"a":"apple","b":"ball","c":"cat"}
	for k,v:=range kvs{
		fmt.Println(k ,"for",v)
	}
	for k := range kvs {
		fmt.Println("key of map kvs:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
