func isValid(s string) bool {
    stack := []rune{}

    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else {
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]

            if (ch == ')' && top != '(') ||
               (ch == '}' && top != '{') ||
               (ch == ']' && top != '[') {
                return false
            }

            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}