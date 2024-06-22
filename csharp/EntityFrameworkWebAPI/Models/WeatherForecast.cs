using EntityFrameworkWebAPI.Models.Requests;

namespace EntityFrameworkWebAPI.Models;

public class WeatherForecast
{
    public static readonly string[] Summaries = new[]
    {
        "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    };

    public int WeatherForecastId { get; set; }

    public DateTime Date { get; set; }

    public int TemperatureC { get; set; }

    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);

    public string? Summary { get; set; }
}

public static class WeatherForecastMappingExtensions
{
    public static WeatherForecastView AsView(this WeatherForecast forecast)
    {
        return new()
        {
            WeatherForecastId = forecast.WeatherForecastId,
            Date = forecast.Date,
            TemperatureC = forecast.TemperatureC,
            Summary = forecast.Summary,
        };
    }
}