package main

func main(){
	// Call SetupServer()
	// Call Begin Server

	// Graceful shudown ?
}


/*
-- channels(buffered) graceful shutdown
OS signal 
OS interrupt
signal.notify

-- server 
will hold: 
Echo Framework object (struct or interface)
handler struct
Will have a router path which will be weather and it would call the handler that has a function called weather

OPTIONAL
May want to start server in a separate goroutine (way to learn about conccurency)

-- handler -> https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
create a client
weather.go (gonna be part of the handler package)
-> func(h *Handler) Weather(c echo.Context) error{...} 
-> split apart your function into smaller functions that do specific jobs. 
	-> getOpenStreetMapLink
	-> openStreetMapData
	-> ... and more
	-> func (h *Handler) aejhlikhfehfes(){...}

-- Tests
write tests and do BDD(behavioural driven development) -> https://go.dev/src/net/http/httptest/example_test.go
<filename>_test.go -> e.g. server_test.go
*/

// your next session is going to be about learning about in memory databases, mongodb and redis (interface)




package main

import (
	"os"
	"os/signal"

	"github.com/jazaltron10/goAPI/weatherAPI/api"
	"github.com/jazaltron10/goAPI/weatherAPI/internal/server"
	"github.com/labstack/echo/v4"
)

func init(){
	// Logging settings
	// logrus
}

func main() {
	// create channel 
	// to gracefully shutdown you application

	gcQuit := make(chan os.Signal, 1)
	signal.Notify(gcQuit, os.Interrupt)

	s :=server.SetupServer()
	s.BeginServer(gcQuit)

	// In the internal that has a file called server
	// server will Begin(gcQuit)
	// go routine -> start the server
	// <-gcQuit
	// gracefulShutdown()... context timeout ... 
	// shutdown your server
	/*
	e := echo.New()

	// Initialize routes
	api.InitializeRoutes(e)

	// Start the server on port 1323.
	e.Logger.Fatal(e.Start(":1323"))
	*/ 

}
