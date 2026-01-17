package stringtointeger

func myAtoi(s string) int {
    res := 0
    sign := 1
    flag := false
    started := false
    for _, c := range s {
        switch {
            case c >= 'a' && c <= 'z':
                return sign*res
            case c >= 'A' && c <= 'Z':
                return sign*res
            case c >= '0' && c <= '9':
                res = res*10 + int(c - '0')
                if res > 1<<32 - 1 || res < -1<<31 {
                    return round(res*sign)
                }
                flag = true
                started = true
            case c == '-' && !flag:
                sign = -1
                flag = true
                started = true
            case c == '+' && !flag:
                flag = true
                started = true
            case c == '.':
                return round(sign*res)
            default:
                if (started) {
                    return round(sign*res)
                }
        }
    }
    return round(sign*res)
}

func round(num int) int {
    if num > 1<<31 - 1 {
        return 1<<31 - 1
    }
    if num < -1<<31 {
        return -1<<31
    }
    return num
}