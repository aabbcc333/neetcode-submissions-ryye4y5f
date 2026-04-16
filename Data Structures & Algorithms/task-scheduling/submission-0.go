func leastInterval(tasks []byte, n int) int {
freq := make([]int,26)

for _, t := range tasks{
	freq[t-'A']++
}

maxFreq := 0
for _,f := range freq{
	if f > maxFreq{
		maxFreq = f
	}
}

countMax := 0 
for _, f := range freq {
	if f == maxFreq{
		countMax ++ 
	}
}

partCount := maxFreq - 1
partLength := n + 1
time := partCount * partLength + countMax

if time < len(tasks){
	return len(tasks)
}

return time
}
