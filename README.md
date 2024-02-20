# weather-api
WeatherService

## How it works!
Weather API uses Data recevied from OpenWeatherAPI to generate intelligent responses. To use the repo, add your personal API_KEY generated via OpenWeather Website into config.yml file against API_KEY environment variable.

### Request format
It accepts requests as /weather?lat={lat}&lon={lon} and accepts floating points values for latitude and longitude for which current weather conditions are sought. The default units queried are in metrics system.

### Response format
* If the temperature feels to be less than 15C, it shows cold weather.
* If the temperature feels to be between 15C and 30C, it returns moderate as response.
* If the temperature feels to be more than 30C, weather is deemed to be hot .


