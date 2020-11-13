package wahoo

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	wahoo "github.com/mornindew/wahoo_client/pkg"
)

var clientSecret string
var clientID string
var accessToken string
var useProduction bool
var redirectURI string

func init() {

	clientSecret = os.Getenv("WAHOO_CLIENT_SECRET")
	clientID = os.Getenv("WAHOO_CLIENT_ID")
	accessToken = "M_Ypa_O05izTArhpl6q4kjvsMMXcxtATfRK4laEG8OU"
	useProduction = true
	redirectURI = "https://www.reddiyo.com"

}
func TestWahooGetToken(t *testing.T) {
	//Get an oath code - https://api.wahooligan.com/oauth/authorize?client_id=<YOUR CLIENT ID>&redirect_uri=<YOUR REDIRECT URI>&response_type=code&scope=<YOUR SCOPES>
	code := "" //Plug in your token that you got from the oauth flow

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//get the token
	token, err := client.GetOauthToken(code)
	if err != nil {
		t.Error("ERror Getting the Token. " + err.Error())
		return
	}
	if token == nil {
		t.Error("Empty Token: ")
		return
	}

	//Get user data
	user, err := client.GetUserData(token.AccessToken)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if user == nil {
		t.Error("Empty User")
	}
}

func TestWahooGetWorkouts(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//get the workouts
	workouts, err := client.GetAllWorkouts(accessToken, 1, 45)
	if err != nil {
		t.Error("Error Getting the Token. " + err.Error())
		return
	}
	if len(workouts) == 0 {
		t.Error("Empty Workouts")
		return
	}

	for _, workout := range workouts {
		fmt.Println("Workout ID:" + strconv.Itoa(workout.ID) + " Start Time: " + workout.Starts.String())
	}
}

func TestWahooGetSpecificWorkoutSummary(t *testing.T) {
	workoutID := 55833989
	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	workout, err := client.GetWorkoutSummary(accessToken, workoutID)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if workout == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestWahooGetSpecificWorkout(t *testing.T) {
	workoutID := 55695256

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	workout, err := client.GetSpecificWorkout(accessToken, workoutID)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if workout == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestGetHeartRateZone(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	workout, err := client.GetHeartRateZones(accessToken)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if workout == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestSetHeartRateZone(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Heart rate zone
	zoneOne := 70
	zoneTwo := 85
	zoneThree := 100
	zoneFour := 125
	zoneFive := 150
	maximum := 180
	resting := 70

	heartRateZone := &wahoo.HeartRateZone{}
	heartRateZone.Zone1 = &zoneOne
	heartRateZone.Zone2 = &zoneTwo
	heartRateZone.Zone3 = &zoneThree
	heartRateZone.Zone4 = &zoneFour
	heartRateZone.Zone5 = &zoneFive
	heartRateZone.Maximum = &maximum
	heartRateZone.Resting = &resting
	//Get the workout data
	err = client.UpdateHeartRateZone(accessToken, heartRateZone)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetPowerZone(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	zones, err := client.GetPowerZones(accessToken)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if zones == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestSetPowerZones(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Heart rate zone
	zoneOne := 70
	zoneTwo := 85
	zoneThree := 100
	zoneFour := 125
	zoneFive := 150
	zoneSix := 200
	zoneSeven := 280
	ftp := 320

	powerZone := &wahoo.PowerZone{}
	powerZone.Zone1 = &zoneOne
	powerZone.Zone2 = &zoneTwo
	powerZone.Zone3 = &zoneThree
	powerZone.Zone4 = &zoneFour
	powerZone.Zone5 = &zoneFive
	powerZone.Zone6 = &zoneSix
	powerZone.Zone7 = &zoneSeven
	powerZone.Ftp = &ftp

	//Get the workout data
	err = client.UpdatePowerZones(accessToken, powerZone)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetUser(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	user, err := client.GetUserData(accessToken)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if user == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestUpdateUser(t *testing.T) {

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//UserData
	user := &wahoo.User{}
	val := "blah@blah.com"
	user.Email = &val
	//Get the workout data
	err = client.UpdateUserData(accessToken, user)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if user == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestWahooUpdateSpecificWorkout(t *testing.T) {
	workoutID := 55695256
	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Wahoo Workout
	workoutToSend := &wahoo.Workout{}
	workoutToSend.ID = workoutID
	summary := &wahoo.WorkoutSummary{}
	workoutToSend.WorkoutSummary = summary
	value := 543.21
	workoutToSend.WorkoutSummary.CaloriesAccum = &value

	//Get the workout data
	err = client.UpdateSpecificWorkout(accessToken, workoutToSend)
	if err != nil {
		t.Error(err.Error())
		return
	}

	//Get the same workout
	workout, err := client.GetSpecificWorkout(accessToken, workoutID)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if workout == nil {
		t.Error("Empty Workout: ")
		return
	}
}

func TestWahooDELETESpecificWorkout(t *testing.T) {
	workoutID := 55833989
	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//Get the workout data
	err = client.DeleteSpecificWorkout(accessToken, workoutID)
	if err != nil {
		t.Error(err.Error())
		return
	}

}
func TestWahooRefreshToken(t *testing.T) {
	refreshToken := "" //Put in your refresh token

	client, err := wahoo.ConstructClient(clientSecret, clientID, redirectURI, useProduction)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//get the token
	token, err := client.RefreshToken(refreshToken)
	if err != nil {
		t.Error("ERror Getting the Token. " + err.Error())
		return
	}

	if token == nil {
		t.Error("Empty Token: ")
		return
	}
}
