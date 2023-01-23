// Package helpers contains helpers functions.
package helpers

import "reflect"

// IsNil return true if the interface or it's content is nil, otherwise return
// false.
func IsNil(a interface{}) bool {
    return a == nil || reflect.ValueOf(a).IsNil()
}
