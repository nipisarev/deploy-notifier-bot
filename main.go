package deploy_notifier_bot

import (
	"log"
)

func main() {
	router := NewRouter()

	log.Fatal(router.Run(":8181"))
}
