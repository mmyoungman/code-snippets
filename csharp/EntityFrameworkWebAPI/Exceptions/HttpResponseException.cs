using System.Net;

namespace EntityFrameworkWebAPI.Exceptions;

public class HttpResponseException : Exception
{
    public HttpResponseException(HttpStatusCode statusCode, string? detail = null)
    {
        this.StatusCode = statusCode;
        this.Detail = detail;
    }

    public HttpStatusCode StatusCode { get; }

    public string? Detail { get; }
}