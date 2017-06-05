package handlers

import (
	"net/http"
)

/*
 * HTTP Handlers are implemented in this package following
 * a simple function pattern as seen below.
 *
 * Service instances can be passed in as variables to the
 * top level function as shown:
 *
 * func GETBase(service pkg.ServiceType) {
 *   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 *     service.DoSomething()
 *   })
 * }
 */

// GETBase is a simple example http.Handler to show how all
// handler functions should be designed.
func GETBase() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
