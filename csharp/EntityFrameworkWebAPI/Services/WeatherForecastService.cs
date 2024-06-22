using System.Net;
using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using Microsoft.EntityFrameworkCore;
using EntityFrameworkWebAPI.Exceptions;

namespace EntityFrameworkWebAPI.Services;

public interface IWeatherForecastService
{
    Task<IEnumerable<WeatherForecastView>> List();

    Task<WeatherForecastView> Get(int id);

    Task<WeatherForecastView> Create(WeatherForecastRequest request);
}

public class WeatherForecastService(
    WeatherForecastContext context,
    IValidationService validationService
        ) : IWeatherForecastService
{
    private readonly WeatherForecastContext _context = context;
    private readonly IValidationService _validationService = validationService;

    public async Task<IEnumerable<WeatherForecastView>> List()
    {
        return await _context.Forecasts
            .Select(f => f.AsView())
            .ToArrayAsync();
    }

    public async Task<WeatherForecastView> Get(int id)
    {
        var forecast = await _context.Forecasts
            .SingleOrDefaultAsync(forecast => forecast.WeatherForecastId == id);

        if (id == 2)
        {
            throw new HttpResponseException(HttpStatusCode.BadRequest, "Invalid forecast id");
        }

        if (id == 3)
        {
            _validationService.AddModelError("YouDidThisWrong", "You did this wrong!");
            _validationService.AddModelError("YouDidThisWrong", "And this!");
            _validationService.AddModelError("YouDidThatWrong", "You did that wrong!");

            throw new HttpResponseException(HttpStatusCode.BadRequest);
        }

        if (forecast == null)
        {
            throw new HttpResponseException(HttpStatusCode.NotFound, "Forecast not found");
        }

        return forecast.AsView();
    }

    public async Task<WeatherForecastView> Create(WeatherForecastRequest request)
    {
        var newForecast = new WeatherForecast
        {
            Date = DateTime.UtcNow,
            TemperatureC = request.TemperatureC,
            Summary = request.Summary,
        };
        await _context.Forecasts.AddAsync(newForecast);
        await _context.SaveChangesAsync();

        return newForecast.AsView();
    }
}
