using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Services;
using Microsoft.AspNetCore.Mvc;

namespace EntityFrameworkWebAPI.Controllers;

[ApiController]
[Route("api")]
public class WeatherForecastController : ControllerBase
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
    public async Task<WeatherForecastView> Get(int id)
    {
        return await _weatherForecastService.Get(id);
    }

    [HttpPost("weather-forecast")]
    public async Task<WeatherForecastView> Create(WeatherForecastRequest request)
    {
        return await _weatherForecastService.Create(request);
    }
}
