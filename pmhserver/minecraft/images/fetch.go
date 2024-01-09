package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var PROJECT_NAME string
var VERSION_NAME string
var BUILD_ID string
var DOWNLOAD_NAME string

var GITHUB_OUTPUT string
var DISABLE_VERSION_UPDATE_CHECK bool

func init() {
	project_name, isProjectNameProvided := os.LookupEnv("PROJECT_NAME")
	PROJECT_NAME = project_name

	log.Printf("PROJECT_NAME: %s\n", PROJECT_NAME)

	if !isProjectNameProvided {
		log.Fatalln("Required environment variable, PROJECT_NAME does not provided.")		
	}

	github_output, isGithubOutputProvided := os.LookupEnv("GITHUB_OUTPUT")
	GITHUB_OUTPUT = github_output
	DISABLE_VERSION_UPDATE_CHECK = !isGithubOutputProvided

	if DISABLE_VERSION_UPDATE_CHECK {
		log.Println("Version update check disabled")
	}
}

func main() {
	log.Println("Fetching VERSION_NAME list...")

	url := fmt.Sprintf("https://api.papermc.io/v2/projects/%s", PROJECT_NAME)
	resp, err := http.Get(url)
	handleAPIError(resp, err)

	var data map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	log.Printf("Response: %s", body)

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	versions := data["versions"].([]interface{})
	latest_version := versions[len(versions) - 1].(string)

	VERSION_NAME = latest_version

	log.Printf("Fetched VERSION_NAME: %s\n", VERSION_NAME)
	log.Println("Fetching BUILD_ID list...")

	url = fmt.Sprintf("%s/versions/%s", url, VERSION_NAME)
	resp, err = http.Get(url)
	handleAPIError(resp, err)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	log.Printf("Response: %s", body)

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	builds := data["builds"].([]interface{})
	latest_build_id := builds[len(builds) - 1].(float64)

	BUILD_ID = fmt.Sprint(latest_build_id)
	log.Printf("Fetched BUILD_ID: %s\n", BUILD_ID)
	log.Println("Fetching DOWNLOAD_NAME list...")

	url = fmt.Sprintf("%s/builds/%s", url, BUILD_ID)
	resp, err = http.Get(url)
	handleAPIError(resp, err)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	log.Printf("Response: %s\n", body)

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("PaperMC API returns an ambiguous data")
	}

	build_channel := data["channel"].(string)
	is_experimental_build := "false"

	if build_channel == "experimental" {
		log.Println("Experimental build detected")
		is_experimental_build = "true"
	}

	downloads := data["downloads"].(map[string]interface{})
	download_data := downloads["application"].(map[string]interface{})
	download_name := download_data["name"].(string)

	DOWNLOAD_NAME = download_name
	log.Printf("Fetched DOWNLOAD_NAME: %s\n", DOWNLOAD_NAME)

	if DISABLE_VERSION_UPDATE_CHECK {
		return
	}

	log.Println("Checking version history...")

	body, err = os.ReadFile("previous_args.json")
	if err != nil {
		log.Fatalf("Version history file, previous_args.json is not readable: %s\n", err.Error())
	}

	log.Printf("Previous args: %s", body)
	log.Println("Calculating NEEDS_UPDATE")

	var file map[string]interface{}
	if err := json.Unmarshal(body, &file); err != nil {
		log.Fatalln("Version history file, previous_args.json is corrupted")
	}

	needs_update := false
	
	previous_data := file[PROJECT_NAME]
	if previous_data == nil {
		needs_update = true
	}

	if previous_data != nil {
		previous_data := previous_data.(map[string]interface{})

		previous_VERSION_NAME := previous_data["VERSION_NAME"].(string)
		if VERSION_NAME != previous_VERSION_NAME {
			needs_update = true
		}

		previous_BUILD_ID := previous_data["BUILD_ID"].(string)
		if BUILD_ID != previous_BUILD_ID {
			needs_update = true
		}

		previous_DOWNLOAD_NAME := previous_data["DOWNLOAD_NAME"].(string)
		if DOWNLOAD_NAME != previous_DOWNLOAD_NAME {
			needs_update = true
		}
	}
	
	if !needs_update {
		os.WriteFile(GITHUB_OUTPUT, []byte("NEEDS_UPDATE=false"), 0666)
		log.Println("NEEDS_UPDATE: false")
		return
	}

	output_body := fmt.Sprintf(
		"NEEDS_UPDATE=true\nVERSION_NAME=%s\nBUILD_ID=%s\nDOWNLOAD_NAME=%s\nIS_EXPERIMENTAL_BUILD=%s\n",
		VERSION_NAME, BUILD_ID, DOWNLOAD_NAME, is_experimental_build)

	os.WriteFile(GITHUB_OUTPUT, []byte(output_body), 0666)
	log.Println("NEEDS_UPDATE: true")
		
	new_data := file
	new_data[PROJECT_NAME] = map[string]string{}
	new_data[PROJECT_NAME].(map[string]string)["VERSION_NAME"] = VERSION_NAME
	new_data[PROJECT_NAME].(map[string]string)["BUILD_ID"] = BUILD_ID
	new_data[PROJECT_NAME].(map[string]string)["DOWNLOAD_NAME"] = DOWNLOAD_NAME

	new_body, _ := json.Marshal(new_data)

	os.WriteFile("./previous_args.json", new_body, 0666)
}

func handleAPIError (resp *http.Response, err error) {
	if err != nil {
		log.Fatalf("PaperMC API call failed: %s\n", err.Error())
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("PaperMC API returns an ambiguous status code: %d\n", resp.StatusCode)
		}

		log.Fatalf("PaperMC API returns an ambiguous error: %s\n", string(body))
	}
}
