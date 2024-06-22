using EntityFrameworkWebAPI.Exceptions;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.ExceptionFilters;

public class HttpResponseExceptionFilter : IActionFilter, IOrderedFilter
{
    private readonly ProblemDetailsFactory _problemDetailsFactory;

    public int Order => int.MaxValue - 10;
    
    public HttpResponseExceptionFilter(
        ProblemDetailsFactory problemDetailsFactory)
        => _problemDetailsFactory = problemDetailsFactory;


    public void OnActionExecuting(ActionExecutingContext context) { }

    public void OnActionExecuted(ActionExecutedContext context)
    {
        if (context.Exception is HttpResponseException exception)
        {
            var problemDetails = _problemDetailsFactory.CreateProblemDetails(
                context.HttpContext,
                statusCode: (int)exception.StatusCode,
                detail: exception.Detail
            );
            context.Result = new ObjectResult(problemDetails);

            context.ExceptionHandled = true;
        }
    }
}   

public class HttpValidationResponseExceptionFilter : IActionFilter, IOrderedFilter
{
    private readonly ProblemDetailsFactory _problemDetailsFactory;

    public int Order => int.MaxValue - 10;
    
    public HttpValidationResponseExceptionFilter(ProblemDetailsFactory problemDetailsFactory)
        {
            _problemDetailsFactory = problemDetailsFactory;
        }


    public void OnActionExecuting(ActionExecutingContext context) { }

    public void OnActionExecuted(ActionExecutedContext context)
    {
        if (context.Exception is HttpValidationResponseException exception)
        {
            var problemDetails = _problemDetailsFactory.CreateValidationProblemDetails(
                context.HttpContext,
                context.ModelState,
                statusCode: (int)exception.StatusCode,
                detail: exception.Detail
            );
            context.Result = new ObjectResult(problemDetails);

            context.ExceptionHandled = true;
        }
    }
}   