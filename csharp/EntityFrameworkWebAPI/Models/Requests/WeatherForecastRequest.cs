using FluentValidation;

namespace EntityFrameworkWebAPI.Models.Requests;

public class WeatherForecastRequest
{
    public int? TemperatureC { get; set; }

    public string? Summary { get; set; }

    public class Validator : AbstractValidator<WeatherForecastRequest>
    {
        public Validator()
        {
            RuleFor(request => request.TemperatureC)
                .Cascade(CascadeMode.Stop)
                .NotNull()
                .GreaterThanOrEqualTo(-273).WithMessage("Cannot be below absolute zero");

            RuleFor(request => request.Summary)
                .Must(summary => summary is null || WeatherForecast.Summaries.Contains(summary))
                .WithMessage($"Must be one of these: {string.Join(' ', WeatherForecast.Summaries)}");
        }
    }
}
