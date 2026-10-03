package hamming
import "errors"
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("String lengths do not match.")
    }
    
    counter := 0
	for index := range a {
        if a[index] != b[index] {
            counter++
        }
    }
        return counter, nil
}
