using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Utils;

public class Result<T>
{
    public readonly T? Value;
    public readonly ResultError? Error;

    public static implicit operator Result<T>(T value) => new(value);
    public static implicit operator Result<T>(ResultError error) => new(error);

    public Result(T result)
    {
        Value = result;
        Error = null;
    }

    public Result(ResultError resultError)
    {
        Value = default;
        Error = resultError;
    }

    public static ObjectResult GetResponse(Result<T> result, HttpContext context, ProblemDetailsFactory? factory)
    {
        if (result.Error.HasValue)
        {
            var (httpStatusCode, detail) = ResultConstants.ErrorMap[result.Error.Value];
            return new ObjectResult(CreateProblemDetails(httpStatusCode, detail, context, factory));
        }

        if (result.Value is null)
        {
            throw new Exception("result.Value should be set if there is no error");
        }

        return new ObjectResult(result.Value);
    }

    private static ProblemDetails CreateProblemDetails(
        HttpStatusCode statusCode,
        string detail,
        HttpContext context,
        ProblemDetailsFactory? factory)
    {
        if (factory == null)
        {
            return new ProblemDetails
            {
                Detail = detail,
                Instance = null,
                Status = (int)statusCode,
                Title = null,
                Type = null,
            };
        }

        return factory.CreateProblemDetails(
            context,
            statusCode: (int)statusCode,
            detail: detail);
    }
}

public enum ResultError
{
    ForecastNotFound,
    InvalidForecastId,
}

public static class ResultConstants
{
    public static readonly Dictionary<ResultError, (HttpStatusCode, string)> ErrorMap =
        new ()
        {
            { ResultError.ForecastNotFound,
                new (HttpStatusCode.NotFound, "Forecast not found")},
            { ResultError.InvalidForecastId,
                new (HttpStatusCode.BadRequest, "Invalid forecast Id")},
        };
}
