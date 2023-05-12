using EntityFrameworkWebAPI.Models;
using EntityFrameworkWebAPI.Models.Requests;
using EntityFrameworkWebAPI.Utils;

namespace EntityFrameworkWebAPI.Services.Interfaces;

public interface IWeatherForecastService
{
    Task<IEnumerable<WeatherForecast>> List();

    Task<Result<WeatherForecast>> Get(int id);

    Task<WeatherForecast> Create(WeatherForecastRequest request);
}
