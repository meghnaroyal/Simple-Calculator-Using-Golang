package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    serverAddr := "192.168.108.7:8080" 
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Choose an option:")
        fmt.Println("1. Enter expression")
        fmt.Println("Enter your choice:")
        var choice string
        fmt.Scanln(&choice)

        switch choice {
        case "1":
  
            fmt.Print("Enter expression (e.g., 1+2, 1-2, 1*2, 1/2): ")
            expression, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading:", err)
                return
            }

          
            _, err = fmt.Fprintf(conn, "%s\n", expression)
            if err != nil {
                fmt.Println("Error sending data:", err)
                return
            }

      
            result, err := bufio.NewReader(conn).ReadString('\n')
            if err != nil {
                fmt.Println("Error receiving result:", err)
                return
            }


            fmt.Println("Result:", result)

        case "2":
            fmt.Print("Enter IP address to match: ")
            ipToMatch, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading:", err)
                return
            }
            ipToMatch = strings.TrimSpace(ipToMatch)

       
            _, err = fmt.Fprintf(conn, "MATCH_IP:%s\n", ipToMatch)
            if err != nil {
                fmt.Println("Error sending data:", err)
                return
            }


            result, err := bufio.NewReader(conn).ReadString('\n')
            if err != nil {
                fmt.Println("Error receiving result:", err)
                return
            }

        
            fmt.Println("Result:", result)

        default:
            fmt.Println("Invalid choice. Please choose again.")
        }
    }
}
