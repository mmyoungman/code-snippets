using System.Net;
using EntityFrameworkWebAPI.Exceptions;
using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using Microsoft.EntityFrameworkCore;

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
            throw new HttpResponseException(HttpStatusCode.BadRequest, "Invalid forecast id");
            // Exception caught by HttpResponseExceptionFilter

        if (id == 3)
        {
            Random random = new();

            if(random.Next() < Int32.MaxValue / 2)
                _validationService.AddModelError("YouDidThisWrong", "You did this wrong!");
            if(random.Next() < Int32.MaxValue / 2)
                _validationService.AddModelError("YouDidThisWrong", "And this!");

            // If you cannot continue execution if errors exist, then do this
            if(!_validationService.ModelIsValid())
                throw new HttpResponseException(HttpStatusCode.BadRequest);
            
            if(random.Next() < Int32.MaxValue / 2)
                _validationService.AddModelError("YouDidThatWrong", "You did that wrong!");
            if(random.Next() < Int32.MaxValue / 2)
                _validationService.AddModelError("YouDidThatWrong", "And that!");
            
            // If you can continue execution after errors, they will still be caught by InvalidModelStateFilter
        }

        if (forecast == null)
            throw new HttpResponseException(HttpStatusCode.NotFound, "Forecast not found");

        return forecast.AsView();
    }

    public async Task<WeatherForecastView> Create(WeatherForecastRequest request)
    {
        var newForecast = new WeatherForecast
        {
            Date = DateTime.UtcNow,
            TemperatureC = request.TemperatureC!.Value,
            Summary = request.Summary,
        };
        await _context.Forecasts.AddAsync(newForecast);
        await _context.SaveChangesAsync();

        return newForecast.AsView();
    }
}
