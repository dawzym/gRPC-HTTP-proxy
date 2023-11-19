package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Define structs to represent the structure of your JSON data
type Code struct {
	ImageUri        string `json:"ImageUri"`
	S3Bucket        string `json:"S3Bucket"`
	S3Key           string `json:"S3Key"`
	S3ObjectVersion string `json:"S3ObjectVersion"`
	ZipFile         string `json:"ZipFile"`
}

type DeadLetterConfig struct {
	TargetArn string `json:"TargetArn"`
}

type FileSystemConfig struct {
	Arn            string `json:"Arn"`
	LocalMountPath string `json:"LocalMountPath"`
}

type ImageConfig struct {
	Command          []string `json:"Command"`
	EntryPoint       []string `json:"EntryPoint"`
	WorkingDirectory string   `json:"WorkingDirectory"`
}

type Environment struct {
	Variables map[string]string `json:"Variables"`
}

type LoggingConfig struct {
	ApplicationLogLevel string `json:"ApplicationLogLevel"`
	LogFormat           string `json:"LogFormat"`
	LogGroup            string `json:"LogGroup"`
	SystemLogLevel      string `json:"SystemLogLevel"`
}

type SnapStart struct {
	ApplyOn string `json:"ApplyOn"`
}

type VpcConfig struct {
	Ipv6AllowedForDualStack bool     `json:"Ipv6AllowedForDualStack"`
	SecurityGroupIds        []string `json:"SecurityGroupIds"`
	SubnetIds               []string `json:"SubnetIds"`
}

type TracingConfig struct {
	Mode string `json:"Mode"`
}

type LambdaFunction struct {
	Architectures        []string           `json:"Architectures"`
	Code                 Code               `json:"Code"`
	CodeSigningConfigArn string             `json:"CodeSigningConfigArn"`
	DeadLetterConfig     DeadLetterConfig   `json:"DeadLetterConfig"`
	Description          string             `json:"Description"`
	Environment          Environment        `json:"Environment"`
	EphemeralStorage     map[string]int     `json:"EphemeralStorage"`
	FileSystemConfigs    []FileSystemConfig `json:"FileSystemConfigs"`
	FunctionName         string             `json:"FunctionName"`
	Handler              string             `json:"Handler"`
	ImageConfig          ImageConfig        `json:"ImageConfig"`
	KMSKeyArn            string             `json:"KMSKeyArn"`
	Layers               []string           `json:"Layers"`
	LoggingConfig        LoggingConfig      `json:"LoggingConfig"`
	MemorySize           int                `json:"MemorySize"`
	PackageType          string             `json:"PackageType"`
	Publish              bool               `json:"Publish"`
	Role                 string             `json:"Role"`
	Runtime              string             `json:"Runtime"`
	SnapStart            SnapStart          `json:"SnapStart"`
	Tags                 map[string]string  `json:"Tags"`
	Timeout              int                `json:"Timeout"`
	TracingConfig        TracingConfig      `json:"TracingConfig"`
	VpcConfig            VpcConfig          `json:"VpcConfig"`
}

func loadDataFromJSONFile(filePath string) (LambdaFunction, error) {
	// Read the JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return LambdaFunction{}, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Create an instance of the struct to hold the data
	var lambdaFunction LambdaFunction

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &lambdaFunction)
	if err != nil {
		return LambdaFunction{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return lambdaFunction, nil
}

func main() {
	// Specify the path to your JSON file
	jsonFilePath := "data.json"

	// Load data from JSON file
	lambdaFunction, err := loadDataFromJSONFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Print the loaded data
	fmt.Printf("Architectures: %v\n", lambdaFunction.Architectures)
	fmt.Printf("FunctionName: %s\n", lambdaFunction.FunctionName)
	// Add more fields as needed...

	// Print specific fields as needed...
	fmt.Printf("Description: %s\n", lambdaFunction.Description)
	fmt.Printf("Environment Variables: %v\n", lambdaFunction.Environment.Variables)

	// Create Function
	awsAccessKey := "YOUR_AWS_ACCESS_KEY"
    awsSecretKey := "YOUR_AWS_SECRET_KEY"
    awsRegion := "your-region"
    functionName := "YourFunctionName"
    roleARN := "arn:aws:iam::your-account-id:role/YourLambdaRole"
    zipFilePath := "your-lambda-function.zip"

    if err := CreateLambdaFunction(awsAccessKey, awsSecretKey, awsRegion, functionName, roleARN, zipFilePath); err != nil {
        log.Fatal("Error creating Lambda function:", err)
}
