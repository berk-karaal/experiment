package app

import (
	dummylogic "demo/internal/dummyLogic"
	"fmt"
)

func main() {
	fmt.Println("Starting the app")
	dummylogic.DummyLogic(5)
}
