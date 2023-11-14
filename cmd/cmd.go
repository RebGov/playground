/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
*/
package cmd

import (
	"context"

	"github.com/RebGov/playground/server"
)

// Execute starts server and provides shutdown
func Execute() {
	_, cancel := start()
	defer server.Shutdown(cancel)
	server.WaitShutdown()

}

// start starts the Server
func start() (ctx context.Context, cancel context.CancelFunc) {
	// Utilize this function and context to start the server; in this way the context and server can be stopped
	ctx, cancel = context.WithCancel(context.Background())
	server.StartServer()
	return
}
