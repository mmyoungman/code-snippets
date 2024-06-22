using System.ComponentModel.DataAnnotations;

namespace EntityFrameworkWebAPI.Models.Requests;

public class WeatherForecastRequest : IValidatableObject
{
    [Required]
    public int? TemperatureC { get; set; }

    public string? Summary { get; set; }

    public IEnumerable<ValidationResult> Validate(ValidationContext validationContext)
    {
        var validationErrors = new List<ValidationResult>();

        if (TemperatureC < -273)
        {
            validationErrors.Add(new ValidationResult(
                "Cannot be below absolute zero",
                new List<string> { nameof(TemperatureC), }));
        }

        if (Summary != null && !WeatherForecast.Summaries.Contains(Summary))
        {
            validationErrors.Add(new ValidationResult(
                $"Must be one of these: {string.Join(' ', WeatherForecast.Summaries)}",
                new List<string> { nameof(Summary), }));
        }

        return validationErrors;
    }
}
