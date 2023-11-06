package main

import (
	"fmt"

	"github.com/lindell/go-burner-email-providers/burner"
)

func main() {
	isBurnerEmail := burner.IsBurnerEmail("wyc42262@cdfaq.com")
	fmt.Println(isBurnerEmail) // true

}
