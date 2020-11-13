package wahoo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

//Client - the client object that makes the calls
type Client struct {
	baseURL      string
	clientSecret string
	redirectURI  string
	clientID     string
}

/*
ConstructClient -


*/
func ConstructClient(wahooClientSecret, wahooClientID, redirectURI string, useProduction bool) (*Client, error) {

	clientToReturn := &Client{
		clientSecret: wahooClientSecret,
		redirectURI:  redirectURI,
		clientID:     wahooClientID,
	}

	//Toggle off of production or sandbox
	if useProduction {
		clientToReturn.baseURL = "api.wahooligan.com"
	} else {
		clientToReturn.baseURL = "developers.staging.wahooligan.com"
	}

	return clientToReturn, nil
}

//AUTHORIZATION ENDPOINTS

//GetOauthToken - Function that will get an Oauth Token from the code provided
func (v *Client) GetOauthToken(code string) (*Token, error) {

	url := "https://" + v.baseURL + "/oauth/token?client_secret=" + v.clientSecret + "&code=" + code + "&redirect_uri=" + v.redirectURI + "&grant_type=authorization_code&client_id=" + v.clientID
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert to a token
	token, err := convertJSONResponseToOauthToken(body)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//RefreshToken - will take the refresh token and get a new oauth token
func (v *Client) RefreshToken(refreshToken string) (*Token, error) {
	if refreshToken == "" {
		return nil, errors.New("Missing Mandatory Value")
	}
	url := "https://" + v.baseURL + "/oauth/token?client_secret=" + v.clientSecret + "&client_id=" + v.clientID + "&grant_type=refresh_token&refresh_token=" + refreshToken
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert to a token
	token, err := convertJSONResponseToOauthToken(body)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//USER ENDPOINTS

//GetUserData - gets the user data
func (v *Client) GetUserData(accessToken string) (*User, error) {

	if accessToken == "" {
		return nil, errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/user"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	user, err := convertJSONResponseToUserToken(body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//UpdateUserData - sets the user data
func (v *Client) UpdateUserData(accessToken string, newUserData *User) error {

	if accessToken == "" || newUserData == nil {
		return errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/user"
	method := "PUT"

	client := &http.Client{}

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	newUserData.convertUserToFormField(writer)

	err := writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	req.Header.Add("Authorization", "Bearer "+accessToken)
	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return constructWahooErrorFromResponse(res.StatusCode)
	}
	return nil
}

//WORKOUT ENDPOINTS

//GetAllWorkouts - Method to get all the workouts
func (v *Client) GetAllWorkouts(accessToken string, pageNumber, resultsPerPage int) ([]*Workout, error) {
	if accessToken == "" {
		return nil, errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/workouts?"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	//Add the query params
	q := req.URL.Query()
	if pageNumber > 0 {
		q.Add("page", strconv.Itoa(pageNumber))
	}
	if resultsPerPage > 0 {
		q.Add("per_page", strconv.Itoa(resultsPerPage))
	}
	//Encode the query params
	req.URL.RawQuery = q.Encode()
	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert the body to the slice to return
	workoutSlice, err := convertJSONResponseToWorkoutArray(body)
	if err != nil {
		return nil, err
	}
	return workoutSlice, nil
}

//GetWorkoutSummary - Method to get a specific workoutSummary
func (v *Client) GetWorkoutSummary(accessToken string, workoutID int) (*WorkoutSummary, error) {
	if accessToken == "" || workoutID == 0 {
		return nil, errors.New("Missing Mandatory Value")
	}
	url := "https://" + v.baseURL + "/v1/workouts/" + strconv.Itoa(workoutID) + "/workout_summary"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert the body to the slice to return
	workout, err := convertJSONResponseToWorkoutSummary(body)
	if err != nil {
		return nil, err
	}
	return workout, nil
}

//GetSpecificWorkout - Method to get a specific workout
func (v *Client) GetSpecificWorkout(accessToken string, workoutID int) (*Workout, error) {
	if accessToken == "" || workoutID == 0 {
		return nil, errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/workouts/" + strconv.Itoa(workoutID)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert the body to the slice to return
	workout, err := convertJSONResponseToWorkout(body)
	if err != nil {
		return nil, err
	}
	return workout, nil
}

//DeleteSpecificWorkout - Method to get a specific workout
func (v *Client) DeleteSpecificWorkout(accessToken string, workoutID int) error {
	if accessToken == "" || workoutID == 0 {
		return errors.New("Missing Mandatory Value")
	}
	url := "https://" + v.baseURL + "/v1/workouts/" + strconv.Itoa(workoutID)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return constructWahooErrorFromResponse(res.StatusCode)
	}
	return nil
}

//UpdateSpecificWorkout - Method to update data on a workout
func (v *Client) UpdateSpecificWorkout(accessToken string, workout *Workout) error {
	//Check that obth workoutID and workout is set
	if workout == nil || workout.ID == 0 {
		return errors.New("Missing Mandatory Value")
	}
	url := "https://" + v.baseURL + "/v1/workouts/" + strconv.Itoa(workout.ID)
	method := "PUT"
	client := &http.Client{}

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	workout.convertWorkoutToFormFields(writer)

	err := writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return constructWahooErrorFromResponse(res.StatusCode)
	}
	return nil
}

//Heart Rate zones Endpoint

//GetHeartRateZones - gets the heart rate zones
func (v *Client) GetHeartRateZones(accessToken string) (*HeartRateZone, error) {

	if accessToken == "" {
		return nil, errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/heart_rate_zone"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert the body to the slice to return
	heartRate, err := convertJSONResponseToHeartRate(body)
	if err != nil {
		return nil, err
	}
	return heartRate, nil
}

//UpdateHeartRateZone - sets the heart rate zone
func (v *Client) UpdateHeartRateZone(accessToken string, newZonesData *HeartRateZone) error {

	if accessToken == "" || newZonesData == nil {
		return errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/heart_rate_zone"
	method := "PUT"

	client := &http.Client{}

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	newZonesData.convertHeartRateZonesToFormFields(writer)

	err := writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	req.Header.Add("Authorization", "Bearer "+accessToken)
	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return constructWahooErrorFromResponse(res.StatusCode)
	}
	return nil
}

//POWER ZONES

//GetPowerZones - gets the Power zones
func (v *Client) GetPowerZones(accessToken string) (*PowerZone, error) {

	if accessToken == "" {
		return nil, errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/power_zone"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return nil, constructWahooErrorFromResponse(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//Convert the body to the slice to return
	powerZone, err := convertJSONResponseToPower(body)
	if err != nil {
		return nil, err
	}
	return powerZone, nil
}

//UpdatePowerZones - sets the power Zone
func (v *Client) UpdatePowerZones(accessToken string, newZonesData *PowerZone) error {

	if accessToken == "" || newZonesData == nil {
		return errors.New("Missing Mandatory Value")
	}

	url := "https://" + v.baseURL + "/v1/power_zone"
	method := "PUT"

	client := &http.Client{}

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	newZonesData.convertPowerZonesToFormFields(writer)

	err := writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	req.Header.Add("Authorization", "Bearer "+accessToken)
	//Do the http request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Handle anything above 299
	if res.StatusCode >= 300 {
		return constructWahooErrorFromResponse(res.StatusCode)
	}
	return nil
}

//Helper funciton
func convertJSONResponseToOauthToken(token []byte) (*Token, error) {
	tokenToReturn := &Token{}
	err := json.Unmarshal(token, tokenToReturn)
	if err != nil {
		return nil, err
	}
	return tokenToReturn, nil
}

func convertJSONResponseToUserToken(userResponse []byte) (*User, error) {
	userToReturn := &User{}
	err := json.Unmarshal(userResponse, userToReturn)
	if err != nil {
		return nil, err
	}
	return userToReturn, nil
}

func convertJSONResponseToWorkoutSummary(data []byte) (*WorkoutSummary, error) {
	response := &WorkoutSummary{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func convertJSONResponseToWorkout(data []byte) (*Workout, error) {
	response := &Workout{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func convertJSONResponseToHeartRate(data []byte) (*HeartRateZone, error) {
	response := &HeartRateZone{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func convertJSONResponseToPower(data []byte) (*PowerZone, error) {
	response := &PowerZone{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func convertJSONResponseToWorkoutArray(data []byte) ([]*Workout, error) {
	response := &GetAllWorkoutsResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Workouts, nil
}
