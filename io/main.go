package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	/*
		스트림을 읽을 때 더 이상 읽어들일 인풋이 없을 경우 발생하는 에러
		Read 메서드는 반드시 이 값을 리턴해야 함
		인풋이 안정적으로 끝났을 떄 EOF가 리턴되어야 한다
		그런 경우가 아니면 ErrUnexpectedEOF 또는 그 외의 다른 에러가 리턴됨
	*/
	fmt.Println(io.EOF)

	fmt.Println(io.ErrUnexpectedEOF)

	/*
		2가지 타입: Reader, Writer 인터페이스
		reader는 스트림에서 바이트를 읽어들이는 함수를 제공
		writer는 그 반대의 역할: 내부의 바이트 스트림에 데이터를 기록

		Reader: From whom I can copy data
		Writer: To whom I can write data to
	*/

	/*데이터 읽어오기: 앞으로 자주 쓰일 내부적인 메서드 입니다*/
	r := strings.NewReader("Hello, world")

	b, err := io.ReadAll(r) //r은 io.Reader인터페이스 구현체, Read 메서드를 제공한다
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	/*데아터 쓰기. 작동 원리까지만 알고 있으면 될것같습니다*/
	r = strings.NewReader("hello world")

	var w bytes.Buffer
	r.WriteTo(&w)

	fmt.Println(w.String())

	/*io 패키지 예제*/
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}

	// Create a new Person object
	person := Person{Name: "John", Age: 30, Address: "123 Main St."}

	// Marshal the Person object to JSON
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP request with the JSON body:
	//요청 바디는 Read 메서드를 제공하므로 서버도 이를 이용해 데이터를 읽는다.
	req, err := http.NewRequest("POST", "https://example.com/api/person", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	//simple http server.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read the request body*************************이 부분에 주목해주세요
		body, err := ioutil.ReadAll(r.Body) // Read 메서드 제공하는 io.readCloser 타입. io.Reader의 구현체
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Unmarshal the request body to a Person object
		var person Person
		err = json.Unmarshal(body, &person)
		if err != nil {
			http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
			return
		}

		// Print the Person object to the console
		fmt.Printf("Received Person: %+v\n", person)

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Person received"))
	})

	// Start the HTTP server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
