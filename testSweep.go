// Package main demonstrates a simple function that returns a greeting message.
package main

const helloMessage = "Hello from test.go!"

func main() {
	// Call the SayHello function and print the result

	println(helloMessage)
	testMessage := testSweep()
	println(testMessage)
}

func testSweep() string {
	// This function returns a greeting message
	return helloMessage
}
