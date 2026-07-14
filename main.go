package main

func main() {

	open := Scan(
		TargetHost,
		Ports,
	)

	PrintResults(
		TargetHost,
		open,
	)

}
