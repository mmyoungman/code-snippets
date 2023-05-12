using System.Net;
using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Services.Interfaces;
using EntityFrameworkWebAPI.Utils;
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
    public async Task<IEnumerable<WeatherForecast>> List()
    {
        return await _weatherForecastService.List();
    }

    [HttpGet("weather-forecast/{id:int}")]
    public async Task<ActionResult<WeatherForecast>> Get(int id)
    {
        var result = await _weatherForecastService.Get(id);

        return Result<WeatherForecast>.GetResponse(result, HttpContext, ProblemDetailsFactory);
    }

    [HttpPost("weather-forecast")]
    public async Task<ActionResult<WeatherForecast>> Create(WeatherForecastRequest request)
    {
        return await _weatherForecastService.Create(request);
    }
}
