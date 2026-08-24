using EntityFrameworkWebAPI;
using EntityFrameworkWebAPI.Exceptions;
using EntityFrameworkWebAPI.Filters;
using EntityFrameworkWebAPI.Services;
using FluentValidation;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

var services = builder.Services;

services.AddDbContext<WeatherForecastContext>();

services.AddControllers(options =>
{
    options.Filters.Add<FluentValidationFilter>();
});

services.AddValidatorsFromAssemblyContaining<Program>();

services.AddProblemDetails();
services.AddExceptionHandler<GlobalExceptionHandler>();

services.AddScoped<IWeatherForecastService, WeatherForecastService>();

var app = builder.Build();

using (var scope = app.Services.CreateScope())
{
    var context = scope.ServiceProvider.GetRequiredService<WeatherForecastContext>();
    await context.Database.MigrateAsync();
}

app.UseExceptionHandler();
app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();
app.Run();
