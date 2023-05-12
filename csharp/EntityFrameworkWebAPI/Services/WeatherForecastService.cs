using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Services.Interfaces;
using EntityFrameworkWebAPI.Utils;
using Microsoft.EntityFrameworkCore;

namespace EntityFrameworkWebAPI.Services;

public class WeatherForecastService : IWeatherForecastService
{
    private readonly WeatherForecastContext _context;

    public WeatherForecastService(
        WeatherForecastContext context)
    {
        _context = context;
    }

    public async Task<IEnumerable<WeatherForecast>> List()
    {
        return await _context.Forecasts.ToArrayAsync();
    }

    public async Task<Result<WeatherForecast>> Get(int id)
    {
        var result = await _context.Forecasts
            .SingleOrDefaultAsync(forecast => forecast.WeatherForecastId == id);

        // should be in WeatherForecastRequest#Validate, but
        // simulating need for multiple errors to be generated in a service
        if (id > 3)
        {
            return ResultError.InvalidForecastId;
        }

        if (result == null)
        {
            return ResultError.ForecastNotFound;
        }

        return result;
    }

    public async Task<WeatherForecast> Create(WeatherForecastRequest request)
    {
        var newForecast = new WeatherForecast
        {
            Date = DateTime.UtcNow,
            TemperatureC = request.TemperatureC,
            Summary = request.Summary,
        };
        await _context.Forecasts.AddAsync(newForecast);
        await _context.SaveChangesAsync();

        return newForecast;
    }
}
