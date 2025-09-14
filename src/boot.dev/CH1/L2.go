package main

import "fmt"

func main() {
	var smsSendingLimit int = 25
	var costPerSMS float64 = 0.02
	var hasPermission bool = true
	var username string = "Greendizer"

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
