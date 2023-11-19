package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateLambdaFunction makes a request to AWS Lambda API to create a new function.
func CreateLambdaFunction(awsAccessKey, awsSecretKey, awsRegion, functionName, roleARN, zipFilePath string) error {
	// Read the Go binary (ZIP file) of your Lambda function
	zipFile, err := ioutil.ReadFile(zipFilePath)
	if err != nil {
		return fmt.Errorf("error reading ZIP file: %v", err)
	}

	// Set up the HTTP request
	url := fmt.Sprintf("https://lambda.%s.amazonaws.com/2015-03-31/functions", awsRegion)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(zipFile))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/zip")
	req.Header.Set("x-amz-Invocation-Type", "RequestResponse") // Adjust as needed

	// Set AWS credentials
	req.SetBasicAuth(awsAccessKey, awsSecretKey)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(body))
	return nil
}
