using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Models.View;
using EntityFrameworkWebAPI.Services;
using Microsoft.AspNetCore.Mvc;

namespace EntityFrameworkWebAPI.Controllers;

[Route("api")]
public class WeatherForecastController : ApiControllerBase
{
    private readonly IWeatherForecastService _weatherForecastService;

    public WeatherForecastController(
        IWeatherForecastService weatherForecastService)
    {
        _weatherForecastService = weatherForecastService;
    }

    [HttpGet("weather-forecast")]
    public async Task<IEnumerable<WeatherForecastView>> List()
    {
        return await _weatherForecastService.List();
    }

    [HttpGet("weather-forecast/{id:int}")]
    public async Task<ActionResult<WeatherForecastView>> Get(int id)
    {
        var result = await _weatherForecastService.Get(id);

        return result.Match<ActionResult<WeatherForecastView>>(
            forecast => Ok(forecast),
            errors => Problem(errors));
    }

    [HttpPost("weather-forecast")]
    public async Task<ActionResult<WeatherForecastView>> Create(WeatherForecastRequest request)
    {
        var result = await _weatherForecastService.Create(request);

        return result.Match<ActionResult<WeatherForecastView>>(
            forecast => Ok(forecast),
            errors => Problem(errors));
    }
}
