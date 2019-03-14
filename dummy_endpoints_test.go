package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestApi1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(api1)

	handler.ServeHTTP(resp, req)

	// Проверяем HTTP код
	if resp.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK) // API works in general
	}
}

func TestApi2(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:8080/api2", strings.NewReader(`{"name":"test name","text":"text text"}`))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(api2)

	handler.ServeHTTP(resp, req)

	// Проверяем HTTP код ответа
	if resp.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK) // API works in general
	}

	headerType := resp.Header().Get("Content-Type")
	wantHeaderType := "application/json"

	// Проверяем заголовок
	if headerType != wantHeaderType {
		t.Errorf("content type header does not match: got %v want %v", headerType, wantHeaderType)
	}
}

func TestApi3(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:8080/api3", strings.NewReader(`{"name":"test name","text":"text text"}`))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(api3)

	handler.ServeHTTP(resp, req)

	// Проверяем HTTP код ответа
	if resp.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK) // API works in general
	}

	headerType := resp.Header().Get("Content-Type")
	wantHeaderType := "application/vnd.api+json"

	// Проверяем заголовок
	if headerType != wantHeaderType {
		t.Errorf("content type header does not match: got %v want %v", headerType, wantHeaderType)
	}
}

func TestApi1Less10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 5

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			res, err := http.Get("http://localhost:8080/api1") //выполнение GET-запроса
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi1Equal10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 10

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			res, err := http.Get("http://localhost:8080/api1") //выполнение GET-запроса
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi1More10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 11

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			res, err := http.Get("http://localhost:8080/api1") //выполнение GET-запроса
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi2Less10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 5
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api2", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi2Equal10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 10
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api2", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi2More10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 11
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api2", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi3Less10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 5
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api3", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi3Equal10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 10
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api3", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}

func TestApi3More10Hits(t *testing.T) {
	wg := &sync.WaitGroup{}
	hits := 11
	jsonStr := `{"name":"test name","text":"text text"}`

	for i := 0; i < hits; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res, err := http.Post("http://localhost:8080/api3", "application/json", strings.NewReader(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			// Проверяем HTTP код
			if res.StatusCode == http.StatusTooManyRequests {
				t.Errorf("Too many requests")
				return
			}
		}()
	}

	wg.Wait()
}
