
# Weather Texter

![weather-texter sample image](https://private-user-images.githubusercontent.com/194737880/428422904-1be90f42-32c7-421e-b4f5-0c2a99007957.JPG?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NDMzNjM3NjcsIm5iZiI6MTc0MzM2MzQ2NywicGF0aCI6Ii8xOTQ3Mzc4ODAvNDI4NDIyOTA0LTFiZTkwZjQyLTMyYzctNDIxZS1iNGY1LTBjMmE5OTAwNzk1Ny5KUEc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjUwMzMwJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI1MDMzMFQxOTM3NDdaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT1iZDE0MTZiY2JlNjg4MTBmMDVhOThhY2ZmZDZkNzdmYzcwYTg5YmE2NmM5ZWJkZTQyOWMwOGU0ZGQyYmIwZTM3JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.2WkDH6i25JjN9KChr3xbPO7VaQaZsmB3gym6umwFv5M)

A simple daily weather notification service designed for dumbphones. 

This Go program fetches weather data from OpenWeatherMap and sends it as a text message through email-to-SMS gateways provided by mobile carriers. This avoids expensive sms services like twilio.

## Features

- Simple daily weather texts
- Completely free to use (uses free tiers of services)
- Works with any phone that can receive SMS

## Prerequisites

- Go (1.16+)
- A free OpenWeatherMap API key
- A Gmail account with 2-Factor Authentication enabled and an App Password
- A mobile phone with a carrier that supports email-to-SMS gateways

## Setup Instructions

### 1. Get an OpenWeatherMap API Key

1. Sign up for a free account at [OpenWeatherMap](https://openweathermap.org/)
2. Navigate to API Keys in your account
3. Create a new free API key or copy your existing one

### 2. Create a Gmail App Password

1. Go to your [Google Account](https://myaccount.google.com/)
2. Navigate to Security
3. Enable 2-Step Verification if not already enabled
4. Go to "App passwords" (under "Signing in to Google")
5. Select "Mail" and "Other" (give it a name like "Weather Texter")
6. Copy the generated app password (16 characters)
If you have trouble finding this setting, [follow this link.](https://myaccount.google.com/apppasswords) 

### 3. Clone and Configure the Repository

```bash
# Clone the repository
git clone https://github.com/knobby-tires/weather-texter.git
cd weather-texter

# Edit the configuration values in main.go
# Replace all template values (YOUR_*) with your actual information
```
### 4. Installation and Configuration

#### FreeBSD

```bash
# Install Go 
pkg install go

# Build the executable
cd /path/to/weather-texter
go build -o weather-texter main.go

# Make it executable
chmod +x weather-texter

# Test run
./weather-texter

# Set up cron job to run daily at at your desired time
# paste this line and replace "30" with desired minute digits and "7" with desired hour digit
crontab -e
# Add the following line:
30 7 * * * cd /path/to/weather-texter && ./weather-texter
```
#### Linux

```bash
# Install Go (Debian/Ubuntu)
sudo apt update
sudo apt install golang-go

# Install Go (Fedora/RHEL)
sudo dnf install golang

# Install Go (Arch Linux)
sudo pacman -S go

# Build the executable
cd /path/to/weather-texter
go build -o weather-texter main.go

# Make it executable
chmod +x weather-texter

# Test run
./weather-texter

# Set up cron job to run daily at at your desired time
# paste this line and replace "30" with desired minute digits and "7" with desired hour digit
crontab -e
# Add the following line:
30 7 * * * cd /path/to/weather-texter && ./weather-texter

# If cron is not installed
# Debian/Ubuntu:
sudo apt install cron
sudo systemctl enable cron
sudo systemctl start cron

# Fedora/RHEL:
sudo dnf install cronie
sudo systemctl enable crond
sudo systemctl start crond

# Arch Linux:
sudo pacman -S cronie
sudo systemctl enable cronie
sudo systemctl start cronie
```

#### macOS

```bash
# Install Go using Homebrew
brew install go

# Or install Go using the official installer
# Download from https://golang.org/dl/

# Build the executable
cd /path/to/weather-texter
go build -o weather-texter main.go

# Make it executable
chmod +x weather-texter

# Test run
./weather-texter

# Set up cron job to run daily at at your desired time
# paste this line and replace "30" with desired minute digits and "7" with desired hour digit
crontab -e
# Add the following line:
30 7 * * * cd /path/to/weather-texter && ./weather-texter
```
## Carrier Email-to-SMS Gateways

Replace `YOUR_CARRIER_DOMAIN` with the appropriate domain for your carrier:

| Carrier | Email-to-SMS Gateway Domain |
|---------|----------------------------|
| Verizon | vtext.com                  |
| AT&T    | txt.att.net                |
| T-Mobile| tmomail.net                |
| Sprint  | messaging.sprintpcs.com    |
| Virgin Mobile | vmobl.com            |
| Metro PCS | mymetropcs.com           |
| Boost Mobile | sms.myboostmobile.com |
| Cricket | sms.cricketwireless.net    |
| Republic Wireless | text.republicwireless.com |
| Google Fi | msg.fi.google.com        |
| Xfinity Mobile | vtext.com           |
| Spectrum Mobile | vtext.com          |
| Tracfone | mmst5.tracfone.com        |
| Mint Mobile | mailmymobile.net       |

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- OpenWeatherMap for their free weather API
- Email-to-SMS gateways provided by mobile carriers

