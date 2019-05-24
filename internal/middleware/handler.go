package middleware

import (
	"fmt"
	"github.com/stevenandrewcarter/gouter/configs"
	"github.com/stevenandrewcarter/gouter/internal/config"
	"html"
	"log"
	"net/http"
)

type Middleware struct {
	Config configs.Config
	Routes config.Routes
}

func (m *Middleware) Load() error {
	content, err := m.Config.GetConfig()
	if err != nil {
		return err
	}
	return m.Routes.Load(content)
}

// Handles the requests to the router
func (m *Middleware) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// startTime := time.Now()
	// from := html.EscapeString(r.URL.Path)
	// to := ""
	// responseCode := 0
	// processingTime := 0.0
	log.Printf("Handling Request from: '%v'", html.EscapeString(r.URL.Path))
	// route, err := models.FindRouteByFrom(from)
	// if err == nil {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// 	client := &http.Client{
	// 		CheckRedirect: nil,
	// 	}
	// 	req, err := http.NewRequest("GET", route.To, nil)
	// 	if err == nil {
	// 		req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// 		resp, err := client.Do(req)
	// 		if err != nil { log.Fatal(err) }
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		if err != nil { log.Fatal(err) }
	// 		w.Write(body)
	// 		responseCode = resp.StatusCode
	// 	}
	// 	to = route.To
	// } else {
	fmt.Fprintf(w, "Not found, %q. Error: %v", html.EscapeString(r.URL.Path), "")
	// responseCode = 404
	// to = "Not Found"
	// }
	// processingTime = time.Now().Sub(startTime).Seconds()
	// models.CreateLog(models.Log{CreatedAt: startTime.Format(time.RFC3339), From: r.URL.Path, To: to, ProcessingTimeSec: processingTime, ResponseCode: responseCode})
}
