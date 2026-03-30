func searchMatrix(matrix [][]int, target int) bool {
 found := false
 for _, value := range matrix{
    for _, value2 := range value{
        if value2 == target{
            found = true
            break 
        }
    }
    if found == true {
        break
    }
 }
 return found
}
