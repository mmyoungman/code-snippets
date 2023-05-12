using EntityFrameworkWebAPI.Models;
using Microsoft.EntityFrameworkCore;

namespace EntityFrameworkWebAPI;

public class WeatherForecastContext : DbContext
{
    public DbSet<WeatherForecast> Forecasts { get; set; } = null!;

    public string DbPath { get; }

    public WeatherForecastContext()
    {
        DbPath = "weather.db";
    }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
        => options.UseSqlite($"Data Source={DbPath}");
}
