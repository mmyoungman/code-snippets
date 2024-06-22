using EntityFrameworkWebAPI;
using EntityFrameworkWebAPI.ExceptionFilters;
using EntityFrameworkWebAPI.Services;
using Microsoft.AspNetCore.Mvc.Infrastructure;

var builder = WebApplication.CreateBuilder(args);

var services = builder.Services;

services
    .AddEntityFrameworkSqlite()
    .AddDbContext<WeatherForecastContext>();

services.AddControllers(options => 
{
    options.Filters.Add<HttpResponseExceptionFilter>();
});

services.AddSingleton<IActionContextAccessor, ActionContextAccessor>(); // for ValidationService
services.AddSingleton<IValidationService, ValidationService>();

services.AddTransient<IWeatherForecastService, WeatherForecastService>();

var app = builder.Build();
if (app.Environment.IsDevelopment())
{
}
app.UseHttpsRedirection();
//app.UseAuthorization();
app.MapControllers();
app.Run();
