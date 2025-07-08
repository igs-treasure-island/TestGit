// Package main demonstrates a simple function that returns a greeting message.
package main

func main() {
	// Call the SayHello function and print the result

	println("Hello from test.go!")
	testMessage := testSweep()
	println(testMessage)
}

func testSweep() string {
	// This function returns a greeting message
	return "Hello from test.go!"
}
