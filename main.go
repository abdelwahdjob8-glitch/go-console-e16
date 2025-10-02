package main

// Importing necessary packages
import (
	"bufio"   // for reading input from the console
	"fmt"     // for formatted I/O (printing)
	"log"     // for logging errors
	"os"      // for OS-related functions (like reading stdin)
	"strings" // for string manipulation functions

	"github.com/abdelwahdjob8-glitch/go-console-e16/weatherApi" // custom package to fetch weather data
)

func main() {
	// Create a reader to read user input from the console
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter a city
	fmt.Print("👉 Entrez le nom d'une ville : ")

	// Read the input until the user presses Enter
	city, _ := reader.ReadString('\n')

	// Remove extra spaces, newlines (\n), or carriage returns (\r)
	city = strings.TrimSpace(city) // <-- removes \r, \n, and extra spaces

	// Call the GetWeather function from the weatherApi package
	weather, err := weatherApi.GetWeather(city)
	if err != nil {
		// If there's an error fetching the weather, log it and exit
		log.Fatal(err)
	}

	// Check if the Weather slice is empty
	if len(weather.Weather) == 0 {
		// No weather data found for the given city
		log.Fatalf("❌ Pas de données météo trouvées pour %s", city)
	}

	// Print a professional, user-friendly weather report
	fmt.Println("\n🌍 Rapport météo")
	fmt.Println("──────────────────────────────")
	fmt.Printf("📍 Ville       : %s\n", weather.Name)                     // City name
	fmt.Printf("🌡️  Température : %.1f °C\n", weather.Main.Temp)         // Temperature
	fmt.Printf("💧 Humidité    : %d %%\n", weather.Main.Humidity)         // Humidity
	fmt.Printf("🔽 Pression    : %d hPa\n", weather.Main.Pressure)        // Pressure
	fmt.Printf("🌬️ Vent        : %.1f m/s\n", weather.Wind.Speed)        // Wind speed
	fmt.Printf("☁️  Ciel        : %s\n", weather.Weather[0].Description) // Weather description
	fmt.Println("──────────────────────────────")
}
