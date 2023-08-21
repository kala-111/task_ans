//create a struct for weathersample.json and print it

package main

type weather_sample struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat:`
}
type Weather_info struct {
	Temperature_2m string `json:"temp"`
	Time           string `json:"time"`
	Main           struct {
		Text string  `json:"time"`
		Icon string  `json:"icon"`
		Code int     `json:"code"`
		Temp float64 `json:"temp"`
	} `json:"main"`
}

//sorting of numbers in asc and desc

/*package main

import (
	"fmt"
	"sort"
)/*

func main() {
	var a = []int{3, 6, 78, 8, 4, 45, 56}
	sort.Slice(a, func(i, j int) bool {
		return a[j] < a[i]
	})
	for _, v := range a {
		fmt.Println(v)
	}
}

/*sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	for _, v := range a {
		fmt.Println(v)
	}
}*/

//findout maximum and minimum

/*package main

import (
	"fmt"
)

func main() {
	values := []int{3, 6, 78, 8, 4, 45, 56}
min := values[0]
for _, v := range values {
	if v < min {
		min = v
	}

}
fmt.Println(min)
max := values[0]
	for _, v := range values {

		if v > max {
			max = v

		}
	}
	fmt.Println(max)
}*/
