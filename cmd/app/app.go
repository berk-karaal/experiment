package app

import (
	"demo/internal/dummylogic"
	"fmt"
)

func main() {
	fmt.Println("Starting the app")
	dummylogic.DummyLogic(5)
}
