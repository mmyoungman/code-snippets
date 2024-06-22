using System.ComponentModel.DataAnnotations;

namespace EntityFrameworkWebAPI.Models.Requests;

public class WeatherForecastRequest : IValidatableObject
{
    [Required]
    public int TemperatureC;

    public string? Summary;

    public IEnumerable<ValidationResult> Validate(ValidationContext validationContext)
    {
        var validationErrors = new List<ValidationResult>();

        if (Summary != null && !WeatherForecast.Summaries.Contains(Summary))
        {
            validationErrors.Add(new ValidationResult(
                $"Must be one of these: {string.Join(' ', WeatherForecast.Summaries)}",
                new List<string> { nameof(Summary), }));
        }

        if (TemperatureC < -273)
        {
            validationErrors.Add(new ValidationResult(
                "TemperatureC cannot be below absolute zero",
                new List<string> { nameof(TemperatureC), }));
        }

        return validationErrors;
    }
}
