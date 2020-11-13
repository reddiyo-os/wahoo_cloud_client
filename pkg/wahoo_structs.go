package wahoo

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

var wahooDateString = "2006-01-02T15:04:05.000Z"
var wahooYearDateFormat = "2006-01-02"

//Token - the actual token
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}

//User - the User Data
type User struct {
	ID            int            `json:"id"`
	Height        *float64       `json:"height"`
	Weight        *float64       `json:"weight"`
	First         *string        `json:"first"`
	Last          *string        `json:"last"`
	Email         *string        `json:"email"`
	Mobile        *string        `json:"mobile"`
	Birth         *time.Time     `json:"birth"`
	Gender        *int           `json:"gender"`
	HeartRateZone *HeartRateZone `json:"heart_rate_zone"`
	PowerZone     *PowerZone     `json:"power_zone"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
}

/*
convertUserToFormField - method that will take values from a user zone and convert them

It will only convert the values that are able to be set on the PUT operations

*/
func (v *User) convertUserToFormField(writer *multipart.Writer) {

	if v.Email != nil {
		_ = writer.WriteField("user[email]", *v.Email)
	}
	if v.First != nil {
		_ = writer.WriteField("user[first]", *v.First)
	}
	if v.Last != nil {
		_ = writer.WriteField("user[last]", *v.Last)
	}
	if v.Mobile != nil {
		_ = writer.WriteField("user[mobile]", *v.Mobile)
	}
	if v.Height != nil {
		_ = writer.WriteField("user[last]", fmt.Sprintf("%f", *v.Height))
	}
	if v.Weight != nil {
		_ = writer.WriteField("user[last]", fmt.Sprintf("%f", *v.Weight))
	}
	if v.Birth != nil {
		_ = writer.WriteField("user[birth]", v.Birth.Format(wahooYearDateFormat))
	}
	if v.Gender != nil {
		_ = writer.WriteField("user[gender]", strconv.Itoa(*v.Gender))
	}
}

/*
UnmarshalJSON - custom json unmarshaller because the API appears to use an explicit null (e.g. passes null instead of nothing)
for empty values and the data types aren't correct (e.g. numbers coming across as strings)
*/
func (v *User) UnmarshalJSON(data []byte) error {

	//mapToStoreValues := make(map[string]interface{})
	var mapToStoreValues map[string]*json.RawMessage

	if err := json.Unmarshal(data, &mapToStoreValues); err != nil {
		return err
	}

	value, exists := mapToStoreValues["id"]
	if exists && value != nil {
		var valueToSet int
		err := json.Unmarshal(*value, &valueToSet)
		if err != nil {
			fmt.Println("Error marhsalling id: " + err.Error())
		}
		v.ID = valueToSet
	}

	value, exists = mapToStoreValues["birth"]
	if exists && value != nil {
		var stringVal string
		err := json.Unmarshal(*value, &stringVal)
		if err != nil {
			fmt.Println("Error marhsalling birtdate: " + err.Error())
		} else {
			t, err := time.Parse(wahooYearDateFormat, stringVal)
			if err != nil {
				fmt.Println("Error Converting birtdate: " + err.Error())
			} else {
				v.Birth = &t
			}
		}
	}

	//Height string      `json:"height"`
	value, exists = mapToStoreValues["height"]
	if exists && value != nil {
		var stringValue string
		err := json.Unmarshal(*value, &stringValue)
		if err != nil {
			fmt.Println("Error marhsalling height: " + err.Error())
		} else {
			tempVal, _ := strconv.ParseFloat(stringValue, 64)
			v.Height = &tempVal
		}
	}

	// //Weight string      `json:"weight"`
	value, exists = mapToStoreValues["weight"]
	if exists && value != nil {
		var stringValue string
		err := json.Unmarshal(*value, &stringValue)
		if err != nil {
			fmt.Println("Error marhsalling weight: " + err.Error())
		} else {
			tempWeight, _ := strconv.ParseFloat(stringValue, 64)
			v.Weight = &tempWeight
		}
	}

	// //First  string      `json:"first"`
	value, exists = mapToStoreValues["first"]
	if exists && value != nil {
		var tempValue string
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling name: " + err.Error())
		} else {
			v.First = &tempValue
		}
	}
	// //Last   string      `json:"last"`
	value, exists = mapToStoreValues["last"]
	if exists && value != nil {
		var tempValue string
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling last name: " + err.Error())
		} else {
			v.Last = &tempValue
		}
	}
	// //Email  string      `json:"email"`
	value, exists = mapToStoreValues["email"]
	if exists && value != nil {
		var tempValue string
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling name: " + err.Error())
		} else {
			v.Email = &tempValue
		}
	}
	// //Mobile interface{} `json:"mobile"`
	value, exists = mapToStoreValues["mobile"]
	if exists && value != nil {
		var tempValue string
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling mobile: " + err.Error())
		} else {
			v.Mobile = &tempValue
		}
	}
	// //Gender int         `json:"gender"`
	value, exists = mapToStoreValues["gender"]
	if exists && value != nil {
		var tempValue int
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling gender: " + err.Error())
		} else {
			v.Gender = &tempValue
		}
	}

	// CreatedAt     time.Time     `json:"created_at"`
	value, exists = mapToStoreValues["created_at"]
	if exists && value != nil {
		var tempValue time.Time
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling Created At: " + err.Error())
		} else {
			v.CreatedAt = &tempValue
		}
	}

	// UpdatedAt     time.Time     `json:"updated_at"`
	value, exists = mapToStoreValues["updated_at"]
	if exists && value != nil {
		var tempValue time.Time
		err := json.Unmarshal(*value, &tempValue)
		if err != nil {
			fmt.Println("Error marhsalling Updated At: " + err.Error())
		} else {
			v.UpdatedAt = &tempValue
		}
	}

	//Marshall the Heart Rate Zone
	value, exists = mapToStoreValues["heart_rate_zone"]
	if exists && value != nil {
		heartRateZones := &HeartRateZone{}
		err := json.Unmarshal(*value, heartRateZones)
		if err != nil {
			fmt.Println("Error Marshalling Heart Rate: " + err.Error())
		}
		v.HeartRateZone = heartRateZones
	}

	//Marshall the Power Zone
	value, exists = mapToStoreValues["power_zone"]
	if exists && value != nil {
		powerZone := &PowerZone{}
		err := json.Unmarshal(*value, powerZone)
		if err != nil {
			fmt.Println("Error Marshalling Power : " + err.Error())
		}
		v.PowerZone = powerZone
	}

	return nil
}

//HeartRateZone - heart raze zone
type HeartRateZone struct {
	ID        int       `json:"id"`
	Zone1     *int      `json:"zone_1"`
	Zone2     *int      `json:"zone_2"`
	Zone3     *int      `json:"zone_3"`
	Zone4     *int      `json:"zone_4"`
	Zone5     *int      `json:"zone_5"`
	Resting   *int      `json:"resting"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Maximum   *int      `json:"maximum"`
}

/*
convertHeartRateZonesToFormFields - method that will take values from a heartrate zone and convert them

It will only convert the values that are able to be set on the PUT operations

*/
func (v *HeartRateZone) convertHeartRateZonesToFormFields(writer *multipart.Writer) {

	if v.Zone1 != nil {
		_ = writer.WriteField("heart_rate_zone[zone_1]", strconv.Itoa(*v.Zone1))
	}
	if v.Zone2 != nil {
		_ = writer.WriteField("heart_rate_zone[zone_2]", strconv.Itoa(*v.Zone2))
	}
	if v.Zone3 != nil {
		_ = writer.WriteField("heart_rate_zone[zone_3]", strconv.Itoa(*v.Zone3))
	}
	if v.Zone4 != nil {
		_ = writer.WriteField("heart_rate_zone[zone_4]", strconv.Itoa(*v.Zone4))
	}
	if v.Zone5 != nil {
		_ = writer.WriteField("heart_rate_zone[zone_5]", strconv.Itoa(*v.Zone5))
	}
	if v.Resting != nil {
		_ = writer.WriteField("heart_rate_zone[resting]", strconv.Itoa(*v.Resting))
	}
	if v.Maximum != nil {
		_ = writer.WriteField("heart_rate_zone[maximum]", strconv.Itoa(*v.Maximum))
	}
}

//PowerZone - the power zone
type PowerZone struct {
	ID        int       `json:"id"`
	Zone1     *int      `json:"zone_1"`
	Zone2     *int      `json:"zone_2"`
	Zone3     *int      `json:"zone_3"`
	Zone4     *int      `json:"zone_4"`
	Zone5     *int      `json:"zone_5"`
	Zone6     *int      `json:"zone_6"`
	Zone7     *int      `json:"zone_7"`
	Ftp       *int      `json:"ftp"`
	ZoneCount *int      `json:"zone_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*
convertHeartRateZonesToFormFields - method that will take values from a heartrate zone and convert them

It will only convert the values that are able to be set on the PUT operations

*/
func (v *PowerZone) convertPowerZonesToFormFields(writer *multipart.Writer) {

	if v.Zone1 != nil {
		_ = writer.WriteField("power_zone[zone_1]", strconv.Itoa(*v.Zone1))
	}
	if v.Zone2 != nil {
		_ = writer.WriteField("power_zone[zone_2]", strconv.Itoa(*v.Zone2))
	}
	if v.Zone3 != nil {
		_ = writer.WriteField("power_zone[zone_3]", strconv.Itoa(*v.Zone3))
	}
	if v.Zone4 != nil {
		_ = writer.WriteField("power_zone[zone_4]", strconv.Itoa(*v.Zone4))
	}
	if v.Zone5 != nil {
		_ = writer.WriteField("power_zone[zone_5]", strconv.Itoa(*v.Zone5))
	}
	if v.Zone6 != nil {
		_ = writer.WriteField("power_zone[zone_6]", strconv.Itoa(*v.Zone6))
	}
	if v.Zone7 != nil {
		_ = writer.WriteField("power_zone[zone_7]", strconv.Itoa(*v.Zone7))
	}
	if v.Ftp != nil {
		_ = writer.WriteField("power_zone[ftp]", strconv.Itoa(*v.Ftp))
	}
	if v.ZoneCount != nil {
		_ = writer.WriteField("power_zone[zone_count]", strconv.Itoa(*v.ZoneCount))
	}
}

//GetAllWorkoutsResponse - the response to get all the workouts response
type GetAllWorkoutsResponse struct {
	Workouts []*Workout `json:"workouts"`
	Total    int        `json:"total"`
	Page     int        `json:"page"`
	PerPage  int        `json:"per_page"`
	Order    string     `json:"order"`
	Sort     string     `json:"sort"`
}

//Workout - the workout data that is returned
type Workout struct {
	ID             int             `json:"id"`
	Starts         *time.Time      `json:"starts"`
	Minutes        *int            `json:"minutes"`
	Name           *string         `json:"name"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	PlanID         *string         `json:"plan_id"`
	WorkoutToken   *string         `json:"workout_token"`
	WorkoutTypeID  *int            `json:"workout_type_id"`
	WorkoutSummary *WorkoutSummary `json:"workout_summary"`
}

/*
convertWorkoutToFormFields - method that will take values from a workout and convert them

It will only convert the values that are able to be set on the PUT operations

*/
func (v *Workout) convertWorkoutToFormFields(writer *multipart.Writer) {

	if v.Starts != nil && !v.Starts.IsZero() {
		//Convert the time to string
		value := v.Starts.Format(wahooDateString)
		_ = writer.WriteField("workout[starts]", value)
	}

	// Minutes        int            `json:"minutes"`
	if v.Minutes != nil {
		_ = writer.WriteField("workout[minutes]", strconv.Itoa(*v.Minutes))
	}

	// Name           string         `json:"name"`
	if v.Name != nil {
		_ = writer.WriteField("workout[name]", *v.Name)
	}

	//PlanID         string         `json:"plan_id"`
	if v.PlanID != nil {
		_ = writer.WriteField("workout[plan_id]", *v.PlanID)

	}
	// WorkoutToken   string         `json:"workout_token"`
	if v.WorkoutToken != nil {
		_ = writer.WriteField("workout[workout_token]", *v.WorkoutToken)
	}
	// WorkoutTypeID  int            `json:"workout_type_id"`
	if v.WorkoutTypeID != nil {
		_ = writer.WriteField("workout[type_id]", strconv.Itoa(*v.WorkoutTypeID))
	}

	// WorkoutSummary WorkoutSummary `json:"workout_summary"`
	if v.WorkoutSummary != nil {
		// HeartRateAvg        *float64  `json:"heart_rate_avg"`
		if v.WorkoutSummary.HeartRateAvg != nil {
			_ = writer.WriteField("workout[workout_summary][heart_rate_avg]", fmt.Sprintf("%f", *v.WorkoutSummary.HeartRateAvg))
		}
		// CaloriesAccum       *float64  `json:"calories_accum"`
		if v.WorkoutSummary.CaloriesAccum != nil {
			_ = writer.WriteField("workout[workout_summary][calories_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.CaloriesAccum))
		}
		// PowerAvg            *float64  `json:"power_avg"`
		if v.WorkoutSummary.PowerAvg != nil {
			_ = writer.WriteField("workout[workout_summary][power_avg]", fmt.Sprintf("%f", *v.WorkoutSummary.PowerAvg))
		}
		// DistanceAccum       *float64  `json:"distance_accum"`
		if v.WorkoutSummary.DistanceAccum != nil {
			_ = writer.WriteField("workout[workout_summary][distance_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.DistanceAccum))
		}
		// CadenceAvg          *float64  `json:"cadence_avg"`
		if v.WorkoutSummary.CadenceAvg != nil {
			_ = writer.WriteField("workout[workout_summary][cadence_avg]", fmt.Sprintf("%f", *v.WorkoutSummary.CadenceAvg))
		}
		// AscentAccum         *float64  `json:"ascent_accum"`
		if v.WorkoutSummary.AscentAccum != nil {
			_ = writer.WriteField("workout[workout_summary][ascent_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.AscentAccum))
		}
		// DurationActiveAccum *float64  `json:"duration_active_accum"`
		if v.WorkoutSummary.DurationActiveAccum != nil {
			_ = writer.WriteField("workout[workout_summary][duration_active_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.DurationActiveAccum))
		}
		// DurationPausedAccum *float64  `json:"duration_paused_accum"`
		if v.WorkoutSummary.DurationPausedAccum != nil {
			_ = writer.WriteField("workout[workout_summary][duration_paused_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.DurationPausedAccum))
		}
		// DurationTotalAccum  *float64  `json:"duration_total_accum"`
		if v.WorkoutSummary.DurationTotalAccum != nil {
			_ = writer.WriteField("workout[workout_summary][duration_total_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.DurationTotalAccum))
		}
		// PowerBikeNpLast     *float64  `json:"power_bike_np_last"`
		if v.WorkoutSummary.PowerBikeNpLast != nil {
			_ = writer.WriteField("workout[workout_summary][power_bike_np_last]", fmt.Sprintf("%f", *v.WorkoutSummary.PowerBikeNpLast))
		}
		// PowerBikeTssLast    *float64  `json:"power_bike_tss_last"`
		if v.WorkoutSummary.PowerBikeTssLast != nil {
			_ = writer.WriteField("workout[workout_summary][power_bike_tss_last]", fmt.Sprintf("%f", *v.WorkoutSummary.PowerBikeTssLast))
		}
		// SpeedAvg            *float64  `json:"speed_avg"`
		if v.WorkoutSummary.SpeedAvg != nil {
			_ = writer.WriteField("workout[workout_summary][speed_avg]", fmt.Sprintf("%f", *v.WorkoutSummary.SpeedAvg))
		}
		// WorkAccum           *float64  `json:"work_accum"`
		if v.WorkoutSummary.WorkAccum != nil {
			_ = writer.WriteField("workout[workout_summary][work_accum]", fmt.Sprintf("%f", *v.WorkoutSummary.WorkAccum))
		}
		// File                File      `json:"file"`
		if v.WorkoutSummary.File != nil {
			if v.WorkoutSummary.File.URL != "" {
				_ = writer.WriteField("workout[workout_summary][file][url]", fmt.Sprintf("%f", *v.WorkoutSummary.WorkAccum))
			}
		}
	}
	return
}

//File - the location of the fit file
type File struct {
	URL string `json:"url"`
}

//WorkoutSummary - the workoutSummary
type WorkoutSummary struct {
	ID                  int       `json:"id"`
	HeartRateAvg        *float64  `json:"heart_rate_avg"`
	CaloriesAccum       *float64  `json:"calories_accum"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	PowerAvg            *float64  `json:"power_avg"`
	DistanceAccum       *float64  `json:"distance_accum"`
	CadenceAvg          *float64  `json:"cadence_avg"`
	AscentAccum         *float64  `json:"ascent_accum"`
	DurationActiveAccum *float64  `json:"duration_active_accum"`
	DurationPausedAccum *float64  `json:"duration_paused_accum"`
	DurationTotalAccum  *float64  `json:"duration_total_accum"`
	PowerBikeNpLast     *float64  `json:"power_bike_np_last"`
	PowerBikeTssLast    *float64  `json:"power_bike_tss_last"`
	SpeedAvg            *float64  `json:"speed_avg"`
	WorkAccum           *float64  `json:"work_accum"`
	File                *File     `json:"file"`
}

/*
UnmarshalJSON - custom json unmarshaller because the API appears to use an explicit null (e.g. passes null instead of nothing)
for empty values and the data types aren't correct (e.g. numbers coming across as strings)
*/
func (v *WorkoutSummary) UnmarshalJSON(data []byte) error {

	mapToStoreValues := make(map[string]interface{})
	if err := json.Unmarshal(data, &mapToStoreValues); err != nil {
		return err
	}

	value, exists := mapToStoreValues["id"]
	if exists && value != nil {
		intValue, ok := value.(float64)
		if ok {
			v.ID = int(intValue)
		}
	}

	//Heart Rate
	value, exists = mapToStoreValues["heart_rate_avg"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.HeartRateAvg = &floatValue
		}
	}

	//Calories Accum
	value, exists = mapToStoreValues["calories_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.CaloriesAccum = &floatValue
		}
	}

	//Created At
	value, exists = mapToStoreValues["created_at"]
	if exists && value != nil {
		stringValue, ok := value.(string)
		if ok {
			//Convert the string to time
			timeVal, err := time.Parse(wahooDateString, stringValue)
			if err != nil {
				fmt.Println("Error Parsing Created At time.  " + err.Error())
			} else {
				v.CreatedAt = timeVal
			}
		}
	}

	//Updated At
	value, exists = mapToStoreValues["updated_at"]
	if exists && value != nil {
		stringValue, ok := value.(string)
		if ok {
			//Convert the string to time
			timeVal, err := time.Parse(wahooDateString, stringValue)
			if err != nil {
				fmt.Println("Error Parsing updated At time.  " + err.Error())
			} else {
				v.UpdatedAt = timeVal
			}
		}
	}

	//Do the power average
	value, exists = mapToStoreValues["power_avg"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.PowerAvg = &floatValue
		}
	}

	// DistanceAccum       string    `json:"distance_accum"`
	value, exists = mapToStoreValues["distance_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.DistanceAccum = &floatValue
		}
	}

	// CadenceAvg          string    `json:"cadence_avg"`
	value, exists = mapToStoreValues["cadence_avg"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.CadenceAvg = &floatValue
		}
	}

	// AscentAccum         string    `json:"ascent_accum"`
	value, exists = mapToStoreValues["ascent_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.AscentAccum = &floatValue
		}
	}

	// DurationActiveAccum string    `json:"duration_active_accum"`
	value, exists = mapToStoreValues["duration_active_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.DurationActiveAccum = &floatValue
		}
	}

	// DurationPausedAccum string    `json:"duration_paused_accum"`
	value, exists = mapToStoreValues["duration_paused_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.DurationPausedAccum = &floatValue
		}
	}

	// DurationTotalAccum  string    `json:"duration_total_accum"`
	value, exists = mapToStoreValues["duration_total_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.DurationTotalAccum = &floatValue
		}
	}

	// PowerBikeNpLast     int       `json:"power_bike_np_last"`
	value, exists = mapToStoreValues["power_bike_np_last"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.PowerBikeNpLast = &floatValue
		}
	}

	// PowerBikeTssLast    float32   `json:"power_bike_tss_last"`
	value, exists = mapToStoreValues["power_bike_tss_last"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.PowerBikeTssLast = &floatValue
		}
	}

	// SpeedAvg            float32   `json:"speed_avg"`
	value, exists = mapToStoreValues["speed_avg"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.SpeedAvg = &floatValue
		}
	}

	// WorkAccum           int       `json:"work_accum"`
	value, exists = mapToStoreValues["work_accum"]
	if exists && value != nil {
		if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
			v.WorkAccum = &floatValue
		}
	}

	// File                File      `json:"file"`
	value, exists = mapToStoreValues["file"]
	if exists && value != nil {
		//Confirm that it is a map[string]interface{} - since this is a nested object
		mapObject, ok := value.(map[string]interface{})
		if ok {
			file := &File{}
			//Get the URL
			value, ok := mapObject["url"]
			if ok {
				stringValue, ok := value.(string)
				if ok {
					file.URL = stringValue
				}
			}
			v.File = file
		}
	}
	return nil
}
