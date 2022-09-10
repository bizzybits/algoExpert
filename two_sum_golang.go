package main
import ("fmt")

func TwoNumberSum(array []int, target int) []int {
    var targetArray []int
    // var emptyArray []int
	for i := 0; i < len(array); i++ {
       fmt.Printf("\ntarget is %v", target)
        fmt.Printf("\n array[i] is %v", array[i])
        for j := 0; j < len(array); j++ {
            if i == j {
                continue
            }
            fmt.Printf("\n array[j] is %v", array[j])
            sum := array[i]+array[j]
            fmt.Printf("\nSum of array[i] (%v)and array[j](%v) is %v", array[i], array[j], sum)
            if sum == target {
                fmt.Printf("\n***sum == target***")
                targetArray = append(targetArray, array[i], array[j])
                fmt.Printf("%v\n", targetArray)
                return targetArray
            } 
        }
    } 
    emptyArray := make([]int, 0)
	return emptyArray
}

