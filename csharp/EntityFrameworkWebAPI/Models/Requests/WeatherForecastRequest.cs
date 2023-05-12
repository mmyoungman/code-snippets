using System.ComponentModel.DataAnnotations;

namespace EntityFrameworkWebAPI.Models.Requests;

public class WeatherForecastRequest : IValidatableObject
{
    public int TemperatureC;

    public string? Summary;

    public IEnumerable<ValidationResult> Validate(ValidationContext validationContext)
    {
        var results = new List<ValidationResult>();

        if (Summary != null && !WeatherForecast.Summaries.Contains(Summary))
        {
            results.Add(new ValidationResult(
                $"Must be one of these: {string.Join(' ', WeatherForecast.Summaries)}",
                new List<string> { nameof(Summary), }));
        }

        return results;
    }
}
