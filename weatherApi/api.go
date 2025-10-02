package weatherApi

// Import necessary packages
import (
	"encoding/json" // for parsing JSON responses
	"fmt"           // for string formatting
	"io/ioutil"     // for reading response body
	"net/http"      // for making HTTP requests
	"os"            // for reading environment variables

	"github.com/joho/godotenv" // for loading .env files
)

// WeatherResponse struct represents the structure of the JSON returned by OpenWeather API
type WeatherResponse struct {
	Cod     int    `json:"cod"`     // response code from API
	Message string `json:"message"` // error message if any
	Name    string `json:"name"`    // city name
	Main    struct {
		Temp     float64 `json:"temp"`     // current temperature
		Pressure int     `json:"pressure"` // atmospheric pressure
		Humidity int     `json:"humidity"` // humidity percentage
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"` // weather description (e.g., "clear sky")
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"` // wind speed in m/s
	} `json:"wind"`
}

// GetWeather fetches weather data for a given city from OpenWeather API
func GetWeather(city string) (*WeatherResponse, error) {
	// Load environment variables from .env file
	_ = godotenv.Load()

	// Get API key from environment variables
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	// Build the URL for the API request
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=fr",
		city, apiKey,
	)

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err // return error if request fails
	}
	defer resp.Body.Close() // close response body when done

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Parse the JSON response into WeatherResponse struct
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err // return error if JSON parsing fails
	}

	// Return the weather data
	return &weather, nil
}
