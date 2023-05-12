using EntityFrameworkWebAPI;
using EntityFrameworkWebAPI.Services;
using EntityFrameworkWebAPI.Services.Interfaces;

var builder = WebApplication.CreateBuilder(args);

var services = builder.Services;
services
    .AddEntityFrameworkSqlite()
    .AddDbContext<WeatherForecastContext>();
services.AddControllers();
services.AddTransient<IWeatherForecastService, WeatherForecastService>();

var app = builder.Build();
if (app.Environment.IsDevelopment())
{
}
app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();
app.Run();
