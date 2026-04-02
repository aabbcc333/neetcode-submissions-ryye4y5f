func largestRectangleArea(heights []int) int {
stack := []int{}
n := len(heights)
maxArea := 0

for i:=0 ; i<= n ; i++{
	var currHeight int
	if i == n{
      currHeight = 0
	}else{
		currHeight = heights[i]
	}

	for len(stack) > 0 && currHeight < heights[stack[len(stack)-1]]{
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		height := heights[top]

		var width int 
		if len(stack) == 0 {
			width = i 
		} else{
			width = i - stack[len(stack)-1]-1
		}

		area := height * width 
		if area > maxArea {
			maxArea = area 
		}

	}
	stack = append(stack,i)
}
return maxArea
}
