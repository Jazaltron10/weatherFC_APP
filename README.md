
---

# Weather Forecast API Project

## Objective
The noble endeavor at hand is the creation of a dainty web service in the illustrious Golang kingdom, employing the Echo framework. This esteemed service shall graciously receive a list of cities as its offering and, in return, bestow upon the inquirer a detailed 2-day weather forecast for each city in the realm.

## Key Components

### 1. Endpoint
The portal to this celestial service is accessed through the `/weather` endpoint, where a humble query parameter named `city` awaits, bearing the burden of a comma-separated list of city names.

### 2. Coordinate Conversion
To unveil the celestial secrets, the service employs the arcane arts of a cartographer at `https://nominatim.openstreetmap.org`. This enigmatic service translates the mere names of cities into coordinates, unlocking the gateways to the celestial heavens.

### 3. Weather Forecast API
An emissary is dispatched to the mystical realms of `https://api.weather.gov/points`, seeking the prophecies foretold by the stars. A URL, hidden in the properties of the response, serves as a map to the actual forecast.

### 4. Filtering Forecast
Navigating the cosmic currents, the service must sift through the myriad "periods" in the forecast. Only those within the next 72 hours shall be deemed worthy, discarding the rest as distant echoes of a future unknown.

### 5. Response JSON
The celestial response, elegantly composed in the sacred JSON script, unfolds the details for each city: the commencement and conclusion times, and the ethereal descriptions of the imminent weather for the upcoming two days.

## Implementation Details

### 1. Client
An envoy, perhaps an HTTP acolyte, embarks on quests to distant realms, making entreaties to the revered external services.

### 2. JSON to Structs
The esoteric JSON responses are transmuted into the sacred glyphs of Go structs, utilizing the mystic JSON tags and the revered `encoding/json` grimoire.

### 3. Caching
Whispers of a file-based cache system echo through the code, a sanctuary where fragments of data may linger, sparing the emissaries from redundant pilgrimages to distant APIs.

### 4. Echo Framework
The resplendent Echo framework orchestrates the creation of a celestial web server, a stage where the cosmic ballet unfolds. It conducts the symphony of HTTP requests and gracefully returns the weather forecasts in response to the supplications of the clients.

### 5. Testing
The implementation is not merely a conjuration but a meticulously crafted spell, encompassing tests to ascertain the fidelity of the summoned functionalities.

## Response.json
A parchment of example response JSON is presented, depicting the forecasted destinies of the cities, inscribed with the arcana of start times, end times, and ethereal descriptions.

## Additional Notes
- The service is designed for ease of invocation, ideally through the revered Docker or the venerable command line.
- The code is expected to exude wisdom in design and be fortified by the bastions of comprehensive testing.
- A file-based cache system is to be woven into the fabric of the code, abstaining from reliance on external frameworks.

In summation, the code endeavors to be a Golang marvel, a web service draped in the Echo framework, delivering unto the querent a meticulously procured 2-day weather forecast for an assemblage of cities. Should further elucidation be sought or specific queries arise, thy humble servant stands ready for inquiry, My Lord.

## File Structure For the Project
            goAPI/
            |-- cmd/
            |   |-- main.go
            |-- configs/
            |   |-- config_test.go
            |   |-- config.go
            |-- internal/
            |   |-- cache/
            |       |-- mocks/
            |           |-- mock_cache.go
            |       |-- cache.go
            |   |-- forecast/
            |       |-- forecast.go
            |   |-- handler/
            |       |-- handler.go
            |       |-- weather.go
            |   |-- server/
            |       |-- server.go
            |-- vendor/
            |-- tests/
            |-- Dockerfile
            |-- go.mod
            |-- go.sum
            |-- Notes.md
            |-- README.md

---