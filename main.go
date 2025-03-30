package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/smtp"
    "time"
)

// Weather response struct
type WeatherResponse struct {
    Weather []struct {
        Description string `json:"description"`
        Main        string `json:"main"`
    } `json:"weather"`
    Main struct {
        Temp     float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
        TempMin  float64 `json:"temp_min"`
        TempMax  float64 `json:"temp_max"`
    } `json:"main"`
    Wind struct {
        Speed float64 `json:"speed"`
    } `json:"wind"`
    Rain struct {
        OneHour float64 `json:"1h,omitempty"`
    } `json:"rain,omitempty"`
    Snow struct {
        OneHour float64 `json:"1h,omitempty"`
    } `json:"snow,omitempty"`
    Name string `json:"name"`
}

func main() {
    // Configuration - replace these values with your own
    weatherAPIKey := "YOUR_OPENWEATHERMAP_API_KEY"
    city := "YOUR_CITY_NAME"
    units := "imperial" // use "metric" for Celsius
    
    // Gmail configuration.
    from := "YOUR_GMAIL_ADDRESS"
    password := "YOUR_GMAIL_APP_PASSWORD"
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    
    // Carrier SMS gateway.
    phoneNumber := "YOUR_PHONE_NUMBER" // 10 digits, no spaces or symbols. Example: (1234567890)
    carrierDomain := "YOUR_CARRIER_DOMAIN" // e.g., vtext.com for Verizon
    to := phoneNumber + "@" + carrierDomain
    
    // Get weather data.
    weatherData, err := getWeather(weatherAPIKey, city, units)
    if err != nil {
        fmt.Printf("Error getting weather: %v\n", err)
        return
    }
    
    // Format message
    message := buildWeatherMessage(weatherData)
    
    // Send email-to-SMS
    err = sendEmailToSMS(from, password, to, smtpHost, smtpPort, message)
    if err != nil {
        fmt.Printf("Error sending email: %v\n", err)
        return
    }
    
    fmt.Println("Weather text sent successfully.")
}

func buildWeatherMessage(weather *WeatherResponse) string {
    // Get current time
    currentTime := time.Now().Format("3:04pm")
    
    // Basic info
    message := fmt.Sprintf("%s Today:\nHigh: %.0fF at %s\nLow: %.0fF at %s\n", 
        weather.Name, 
        weather.Main.TempMax,
        currentTime,
        weather.Main.TempMin,
        currentTime)
    
    // Add wind info if greater than 8mph
    if weather.Wind.Speed > 8 {
        message += fmt.Sprintf("%.0f mph wind at %s\n", weather.Wind.Speed, currentTime)
    }
    
    // Add rain info if available
    if weather.Rain.OneHour > 0 {
        message += fmt.Sprintf("%.1f inches rain at %s\n", weather.Rain.OneHour, currentTime)
    }
    
    // Add snow info if available
    if weather.Snow.OneHour > 0 {
        message += fmt.Sprintf("%.1f inches snow at %s\n", weather.Snow.OneHour, currentTime)
    }
    
    // Check if weather is stormy (they are cozy)
    weatherType := weather.Weather[0].Main
    if weatherType == "Thunderstorm" {
        message += fmt.Sprintf("Thunderstorms at %s\n", currentTime)
    }
    
    return message
}

func getWeather(apiKey, city, units string) (*WeatherResponse, error) {
    url := fmt.Sprintf(
        "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s",
        city, apiKey, units,
    )
    
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("weather API returned status code %d", resp.StatusCode)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    var weatherResp WeatherResponse
    err = json.Unmarshal(body, &weatherResp)
    if err != nil {
        return nil, err
    }
    
    return &weatherResp, nil
}

func sendEmailToSMS(from, password, to, smtpHost, smtpPort, body string) error {
    // Authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)
    
    // Compose email
    subject := "" // Keep subject empty for cleaner SMS!
    mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
    msg := []byte("To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        mime + "\r\n" +
        body + "\r\n")
    
    // Send email - done :)
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
    if err != nil {
        return err
    }
    
    return nil
}