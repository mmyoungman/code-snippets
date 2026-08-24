using Microsoft.AspNetCore.Diagnostics;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Exceptions;

public class GlobalExceptionHandler(
    ProblemDetailsFactory problemDetailsFactory,
    ILogger<GlobalExceptionHandler> logger) : IExceptionHandler
{
    public async ValueTask<bool> TryHandleAsync(
        HttpContext httpContext,
        Exception exception,
        CancellationToken cancellationToken)
    {
        logger.LogError(exception, "Unhandled exception for {Method} {Path}",
            httpContext.Request.Method, httpContext.Request.Path);

        var problemDetails = problemDetailsFactory.CreateProblemDetails(
            httpContext,
            statusCode: StatusCodes.Status500InternalServerError,
            title: "An unexpected error occurred");

        httpContext.Response.StatusCode = StatusCodes.Status500InternalServerError;
        await httpContext.Response.WriteAsJsonAsync(problemDetails, cancellationToken);

        return true;
    }
}
