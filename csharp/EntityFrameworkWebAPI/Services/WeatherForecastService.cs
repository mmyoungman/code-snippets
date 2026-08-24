using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Models.View;
using ErrorOr;
using Microsoft.EntityFrameworkCore;

namespace EntityFrameworkWebAPI.Services;

public interface IWeatherForecastService
{
    Task<IEnumerable<WeatherForecastView>> List();
    Task<ErrorOr<WeatherForecastView>> Get(int id);
    Task<ErrorOr<WeatherForecastView>> Create(WeatherForecastRequest request);
}

public class WeatherForecastService(WeatherForecastContext context) : IWeatherForecastService
{
    private readonly WeatherForecastContext _context = context;

    public async Task<IEnumerable<WeatherForecastView>> List()
    {
        return await _context.Forecasts
            .Select(f => f.AsView())
            .ToArrayAsync();
    }

    public async Task<ErrorOr<WeatherForecastView>> Get(int id)
    {
        var forecast = await _context.Forecasts
            .SingleOrDefaultAsync(forecast => forecast.WeatherForecastId == id);

        if (id == 3)
        {
            List<Error> errors = [];

            if (Random.Shared.Next(2) == 0)
                errors.Add(Error.Validation("YouDidThisWrong", "You did this wrong!"));
            if (Random.Shared.Next(2) == 0)
                errors.Add(Error.Validation("YouDidThisWrong", "And this!"));

            // If you cannot continue execution while these errors exist, return here.
            if (errors.Count > 0)
                return errors;

            if (Random.Shared.Next(2) == 0)
                errors.Add(Error.Validation("YouDidThatWrong", "You did that wrong!"));
            if (Random.Shared.Next(2) == 0)
                errors.Add(Error.Validation("YouDidThatWrong", "And that!"));

            // If you can continue, keep accumulating and return the whole list at the end.
            if (errors.Count > 0)
                return errors;
        }

        if (forecast is null)
            return Error.NotFound(description: "Forecast not found");

        return forecast.AsView();
    }

    public async Task<ErrorOr<WeatherForecastView>> Create(WeatherForecastRequest request)
    {
        List<Error> errors = [];

        var today = DateTime.UtcNow.Date;
        if (await _context.Forecasts.AnyAsync(f => f.Date.Date == today))
            errors.Add(Error.Validation("Date", "A forecast already exists for today"));

        if (request.Summary is not null
            && !SummaryMatchesTemperature(request.Summary, request.TemperatureC!.Value))
        {
            errors.Add(Error.Validation(
                "Summary",
                $"'{request.Summary}' does not match {request.TemperatureC}C"));
        }

        if (errors.Count > 0)
            return errors;

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

    private static bool SummaryMatchesTemperature(string summary, int temperatureC) => summary switch
    {
        "Freezing" => temperatureC <= 0,
        "Scorching" => temperatureC >= 35,
        _ => true,
    };
}
