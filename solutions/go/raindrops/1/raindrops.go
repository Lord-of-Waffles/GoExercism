package raindrops
import (
    "strconv"
)
func Convert(number int) string {
	//panic("Please implement the Convert function")
    var result string 
    var divisible3 bool = false
	var divisible5 bool = false
    var divisible7 bool = false
    
    if number % 3 == 0 {
        result += "Pling"
    } else {
        divisible3 = true
    }
    if number % 5 == 0 {
        result += "Plang"
    } else {
        divisible5 = true
    }
    if number % 7 == 0 {
        result += "Plong"
    } else {
        divisible7 = true
    }

    if divisible3 == true && divisible5 == true && divisible7 == true {
        result = strconv.Itoa(number)
        return result
    } else {
        return result
    }
    
}
