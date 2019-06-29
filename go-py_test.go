package GoPy

import "testing"


func TestGoPy(t *testing.T) {
	defer AtExit()
	Exec("from time import time,ctime\n")
	Exec("print('Today is', ctime(time()))\n")
}
