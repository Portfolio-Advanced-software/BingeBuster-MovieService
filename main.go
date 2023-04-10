package main

import "github.com/Portfolio-Advanced-software/BingeBuster-MovieService/messaging"

func main() {
	messaging.ProduceMessage("delete movie 1 from user 5", "user_movie")
	messaging.ConsumeMessage("user_movie")
}
