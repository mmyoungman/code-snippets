using System.Net;
using Microsoft.AspNetCore.Mvc.ModelBinding;

namespace EntityFrameworkWebAPI.Exceptions;

public class HttpResponseException : Exception
{
    public HttpResponseException(HttpStatusCode statusCode, string? detail = null, ModelStateDictionary? modelState = null)
    {
        this.StatusCode = statusCode;
        this.Detail = detail;
    }

    public HttpStatusCode StatusCode { get; }

    public string? Detail { get; }

    public ModelStateDictionary? ModelState { get; }
}

public class HttpValidationResponseException : Exception
{
    public HttpValidationResponseException(HttpStatusCode statusCode, string? detail = null)
    {
        this.StatusCode = statusCode;
        this.Detail = detail;
    }

    public HttpStatusCode StatusCode { get; }

    public string? Detail { get; }
}