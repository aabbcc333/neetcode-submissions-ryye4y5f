

func carFleet(target int, position []int, speed []int) int {
    n := len(position)

    cars := make([][2]float64, n)

    // build cars array
    for i := 0; i < n; i++ {
        time := float64(target-position[i]) / float64(speed[i])
        cars[i] = [2]float64{float64(position[i]), time}
    }

    // sort by position
    sort.Slice(cars, func(i, j int) bool {
        return cars[i][0] < cars[j][0]
    })

    stack := []float64{}

    // traverse from right → left
    for i := n - 1; i >= 0; i-- {
        time := cars[i][1]

        if len(stack) == 0 || time > stack[len(stack)-1] {
            stack = append(stack, time)
        }
    }

    return len(stack)
}