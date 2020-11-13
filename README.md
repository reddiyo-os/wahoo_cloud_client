# Wahoo API Client

This reposity contains a Golang Client Library for accessing and using Wahoo's Api.  The API is relatively new and this client will evolve.

There is no guarantees of BW incompatability changes as of right now.

## How to Download

Easy to install from the command line or import directly

Command Line:

    go get github.com/mornindew/wahoo_client/pkg

Import:

    import(
    "github.com/mornindew/wahoo_client/pkg"
    )

## How to Use

1. Get a Client Secret and Client ID
2. Construct a Client

    This can be done globally or per each thread.  Your use case will drive how you need it to be used.

    Construct A Client:

        client, err := ConstructClient(clientSecret, clientID, redirectURI, useProduction)
        if err != nil {
            t.Error(err.Error())
            return 
            }

3. Use the Methods

    Once a client object is constructed then you can acces the endpoint through each of its methods

        workouts, err := client.GetAllWorkouts(accessToken, 1, 45)
        if err != nil {
            t.Error("Error Getting the Token. " + err.Error())
            return
        }

## Methods

### Authorization

- GetOauthToken - will return the oauth token
- RefreshToken - will GET a new oauth token from the refresh token

### User

- GetUserData  - Will GET the user data
- UpdateUserData - Will PUT new user data on the user

### Workout

- GetAllWorkouts - Will GET all the workouts for a user
- GetWorkoutSummary - Will GET a summary from a specific workout
- GetSpecificWorkout - Will GET a specific workout (and it's summary)
- DeleteSpecificWorkout - Will DELETE a specific workout
- UpdateSpecificWorkout - Will UPDATE a specific workout

### Heart Rate Zones

- GetHeartRateZones - Will GET a specific users Heart Rate Zones
- UpdateHeartRateZone - Will PUT data on a specific users Heart Rate Zones

### Power Zones

- GetPowerZones - Will GET the power zones for a user
- UpdatePowerZones - Will PUT data on a users specific Power Zones

## Full Working Examples

Refer to the [unit tests](https://github.com/mornindew/wahoo_client/tree/main/test) to see many more full working examples.

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

## Issues

If you have any issues please open a github issue
