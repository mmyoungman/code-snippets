using System.Net;
using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Mvc.Infrastructure;
using EntityFrameworkWebAPI.Exceptions;

namespace EntityFrameworkWebAPI.Services;

public interface IWeatherForecastService
{
    Task<IEnumerable<WeatherForecast>> List();

    Task<WeatherForecast> Get(int id);

    Task<WeatherForecast> Create(WeatherForecastRequest request);
}

public class WeatherForecastService(
    IActionContextAccessor actionContextAccessor,
    WeatherForecastContext context
        ) : IWeatherForecastService
{
    private readonly IActionContextAccessor _actionContextAccessor = actionContextAccessor;
    private readonly WeatherForecastContext _context = context;

    public async Task<IEnumerable<WeatherForecast>> List()
    {
        return await _context.Forecasts.ToArrayAsync();
    }

    public async Task<WeatherForecast> Get(int id)
    {
        var result = await _context.Forecasts
            .SingleOrDefaultAsync(forecast => forecast.WeatherForecastId == id);

        if (id == 2)
        {
            var modelState = _actionContextAccessor.ActionContext.ModelState;
            modelState.AddModelError("YouDidThisWrong", "You did this wrong!");
            modelState.AddModelError("YouDidThisWrong", "And this!");
            modelState.AddModelError("YouDidThatWrong", "You did that wrong!");

            throw new HttpValidationResponseException(HttpStatusCode.BadRequest);
        }

        if (id > 3)
        {
            throw new HttpResponseException(HttpStatusCode.BadRequest, "Invalid forecast id");
        }

        if (result == null)
        {
            throw new HttpResponseException(HttpStatusCode.NotFound, "Forecast not found");
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
