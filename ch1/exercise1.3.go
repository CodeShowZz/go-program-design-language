package concat_test

import (
 "strings"
 "testing"
)

var args  = []string{"hi","three","buddy","boy","5","6","7","8","9"}

func concat(args []string) {
 r,sep := "",""
 for _, a:= range args {
    r += sep + a	
    sep = " "
}
}

    



