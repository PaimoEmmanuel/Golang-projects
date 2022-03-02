package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"white": "#ffffff",
	// }
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
}

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here
    division := []int32{}
    noOfPossibleDivision := int32(0)
	index := 0

    for i, number := range s {
		fmt.Println(i, "index")
        division = append(division, number)
        if(int32(len(division)) == m) {
			fmt.Println(division)
            sum := int32(0)
            for _, divisionItem := range division {
                sum += divisionItem
            }
            if(sum == d) {
                noOfPossibleDivision++
            }
        } else if(int32(len(division)) > m) {
            division = []int32{}
			index++
			i = index
        }
    } 
    return noOfPossibleDivision
}