# Google Calendar Link Generator

This project is a web application that generates Google Calendar links from natural language input. It uses OpenAI's GPT model to parse the input and extract event details.

## Features

- Generates Google Calendar links from natural language input
- Handles events with or without specific times
- Includes original input in the event description for reference
- Adds local date and time information to the event

## Installation

### 1. Install Go

To install Go on Mac:

1. Visit the official Go download page: https://go.dev/dl/
2. Download the macOS installer package (e.g., "go1.17.7.darwin-amd64.pkg" or the latest version available for macOS).
3. Open the downloaded package and follow the installation wizard.
4. Verify the installation by opening Terminal and running:

   go version

### 2. Set up OpenAI API Key

1. Sign up for an OpenAI account and obtain an API key.
2. Set the API key as an environment variable:

   export OPENAI_API_KEY=your_api_key_here

   Replace your_api_key_here with your actual OpenAI API key.

### 3. Install Required Modules

In your project directory, run:

go mod init calendar-link-generator
go get github.com/sashabaranov/go-openai

## Running the Application

1. Clone this repository or copy the main.go and index.html files to your local machine.
2. In the project directory, run:

   go run main.go

3. Open a web browser and navigate to http://localhost:8080

## Usage

1. Once the application is running, you can paste text directly onto the page or enter it into the textarea.
2. The application will automatically generate a Google Calendar link based on the input.

Example input to paste:

Meeting with John at Starbucks on 5th Avenue tomorrow at 3pm to discuss the new project proposal

## Building an Executable

To create a standalone executable:

1. In the project directory, run:

   go build -o calendar-link-generator

2. This will create an executable file named calendar-link-generator in your current directory.

3. To run the executable:

   ./calendar-link-generator

   The application will start and be accessible at http://localhost:8080

## Notes

- If no end time is specified for an event, it will default to a 1-hour duration.
- If no time is specified at all, it will be created as an all-day event.
- The application includes the local date, time, and time zone information in the event description.
