using System.Text.Json.Serialization;

namespace EntityFrameworkWebAPI.Models.Requests;

public class WeatherForecastView
{
    public int WeatherForecastId { get; set; }

    public DateTime Date { get; set; }

    public int TemperatureC { get; set; }

    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);

    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public string? Summary { get; set; }
}
