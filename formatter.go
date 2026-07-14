package main

import "fmt"

func PrintResults(host string, ports []int) {

	fmt.Println("===========================")
	fmt.Println("PORT SCANNER")
	fmt.Println("===========================")

	fmt.Println()

	fmt.Println("Target:", host)

	fmt.Println()

	if len(ports) == 0 {

		fmt.Println("No open ports found.")

		return

	}

	fmt.Println("Open ports:")

	for _, port := range ports {

		fmt.Printf(" - %d\n", port)

	}

	fmt.Println()

	fmt.Printf("Total: %d\n", len(ports))

}
