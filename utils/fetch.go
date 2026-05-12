package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wvtrserv/logger"
)

func Give[T any](obj T, w http.ResponseWriter, log bool) {
	strsend, err := json.Marshal(obj)
	toSend := string(strsend)
	if err != nil {
		logger.ErrLog.Println("Problem, can't decode object: ", err)
		fmt.Fprintf(w, "%s", "{}")
	}
	if log {
		logger.DumpLog.Println("Give : ", toSend)
	}
	fmt.Fprintf(w, "%s", toSend)
}

func GetParamInt(paramName string, r *http.Request) int {
	s := r.PathValue(paramName)
	res, _ := strconv.Atoi(s)
	return res
}

func DecodeJson[T any](obj *T, reader io.Reader) *T {
	err := json.NewDecoder(reader).Decode(obj)
	if err != nil {
		logger.ErrLog.Println("Problem while trying to decode: ", err)
		return nil
	}
	return obj
}

func ReadResponse(response *http.Response) []byte {
	// Read and print response
	logger.DumpLog.Println(response)
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}
	}
	//logger.DumpLog.Println(string(resp))
	return resp
}

func Fetch(reqURL string, method string, paramsJSON string, header []string) *http.Response {
	// Create a new HTTP client
	client := &http.Client{
		Timeout: time.Second * 60, // Timeout each requests
	}

	req := CreateRequest(reqURL, method, paramsJSON, header)

	//Execute the request using the custom HTTP client
	response, err := client.Do(req)
	if err != nil {
		logger.ErrLog.Println("Error making request:", err)
		return nil
	}

	return response
}

func CreateRequest(reqURL string, method string, paramsJSON string, header []string) *http.Request {

	// logger.DumpLog.Println(reqURL)
	// logger.DumpLog.Println(header)
	// logger.DumpLog.Println(method)
	// logger.DumpLog.Println(params.Encode())

	req, err := http.NewRequest(method, reqURL, strings.NewReader(paramsJSON))

	//logger.DumpLog.Println("request body : ", req.Body)

	if err != nil {
		logger.ErrLog.Println("Error creating request:", err)
		return nil
	}

	if len(header)%2 != 0 {
		logger.ErrLog.Println("Error creating header: the number of parameters don't match")
		return nil
	}

	// Set headers
	for i := 0; i < len(header); i += 2 {
		req.Header.Add(header[i], header[i+1])
	}

	return req
}
