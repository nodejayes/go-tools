package list

import (
	"fmt"
	"strings"
)

// Join List Elements into a string seperated by separator
func (l List[T]) Join(separator string, formatter ...func(T) string) string {
	return strings.Join(Map(l, func(item T, _ int, _ List[T]) string {
		if len(formatter) < 1 {
			return fmt.Sprintf("%v", item)
		}
		return formatter[0](item)
	}).GetItems(), separator)
}
