
package utils
import (
 	"io/ioutil"
 	"strings"
	"strconv"
)

func ReadIntArrayFromFile(fname string) (nums []int, err error) {
	//Credit- https://stackoverflow.com/a/9863218/6642224
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(string(b), "\n")
    // Assign cap to avoid resize on every append.
    nums = make([]int, 0, len(lines))

    for _, l := range lines {
        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }
        // Atoi better suits the job when we know exactly what we're dealing
        // with. Scanf is the more general option.
        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }
        nums = append(nums, n)
    }

    return nums, nil
}
func ReadStringArrayFromFile(fname string) (lines []string, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines = strings.Split(string(b), "\n")
    // Assign cap to avoid resize on every append.
    return lines, nil

   
}