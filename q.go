package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
WAP to print the following:
i.   No. of cars in each Year, like 2020: 22, 2018: 8 ...
ii.  No. of cars in each Category, like Sedan: 22, SUV: 17, Coupe: 5 ….
iii. No. of cars in each Make, like Tesla: 4, Honda: 9, Hyundai: 8 …
iv. No. of cars having more than 1 Category.
*/
func main() {
	data, err := os.ReadFile("cars.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	yearCount(result)
	categoryCount(result)
	makeCount(result)
	multiCategoryCount(result)
}

func yearCount(result map[string]interface{}) {
	ans := make(map[string]int)
	for _, val := range result {
		for y, vals := range val.(map[string]interface{}) {
			for _, v := range vals.(map[string]interface{}) {
				ans[y] += len(v.([]interface{}))
			}
		}
	}
	fmt.Println("Year counts:")
	for y, c := range ans {
		fmt.Printf("%s: %d\n", y, c)
	}
	fmt.Println()
}

func categoryCount(result map[string]interface{}) {
	ans := make(map[string]int)
	for cat, val := range result {
		for _, vals := range val.(map[string]interface{}) {
			for _, v := range vals.(map[string]interface{}) {
				ans[cat] += len(v.([]interface{}))
			}
		}
	}
	fmt.Println("Category counts:")
	for c, n := range ans {
		fmt.Printf("%s: %d\n", c, n)
	}
	fmt.Println()
}

func makeCount(result map[string]interface{}) {
	ans := make(map[string]int)
	for _, val := range result {
		for _, vals := range val.(map[string]interface{}) {
			for mk, v := range vals.(map[string]interface{}) {
				ans[mk] += len(v.([]interface{}))
			}
		}
	}
	fmt.Println("Make counts:")
	for m, n := range ans {
		fmt.Printf("%s: %d\n", m, n)
	}
	fmt.Println()
}

func multiCategoryCount(result map[string]interface{}) {
	modelCat := make(map[string]map[string]bool)
	for cat, val := range result {
		for _, vals := range val.(map[string]interface{}) {
			for _, v := range vals.(map[string]interface{}) {
				for _, m := range v.([]interface{}) {
					ms := m.(string)
					if modelCat[ms] == nil {
						modelCat[ms] = make(map[string]bool)
					}
					modelCat[ms][cat] = true
				}
			}
		}
	}
	cnt := 0
	for _, v := range modelCat {
		if len(v) > 1 {
			cnt++
		}
	}
	fmt.Printf("Cars with more than 1 category: %d\n", cnt)
}
