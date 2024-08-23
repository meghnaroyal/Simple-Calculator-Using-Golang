// server.go

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection established")

	reader := bufio.NewReader(conn)
	for {
		// Read the expression from the client
		expression, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		expression = strings.TrimSpace(expression)

		// Replace all spaces in the expression
		expression = strings.ReplaceAll(expression, " ", "")

		// Split the expression into operator and operands
		parts := strings.Split(expression, "+")
		if len(parts) != 2 {
			parts = strings.Split(expression, "-")
		}
		if len(parts) != 2 {
			parts = strings.Split(expression, "*")
		}
		if len(parts) != 2 {
			parts = strings.Split(expression, "/")
		}

		if len(parts) != 2 {
			fmt.Fprintln(conn, "Invalid expression format. Please provide in the format 'operand operator operand'")
			continue
		}

		operand1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Fprintln(conn, "Invalid operand")
			continue
		}

		operand2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Fprintln(conn, "Invalid operand")
			continue
		}

		var result float64
		operator := expression[len(parts[0])]
		switch operator {
		case '+':
			result = operand1 + operand2
		case '-':
			result = operand1 - operand2
		case '*':
			result = operand1 * operand2
		case '/':
			if operand2 == 0 {
				fmt.Fprintln(conn, "Division by zero is not allowed")
				fmt.Printf("Division by zero not allowed:\n")
				continue
			}
			result = operand1 / operand2
		default:
			fmt.Fprintln(conn, "Invalid operator")
			continue
		}

		// Print the result on the server side
		fmt.Printf("Result: %f\n", result)

		// Send the result back to the client
		fmt.Fprintf(conn, "Result: %f\n", result)
	}
}

func main() {
	fmt.Println("Server is running...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
