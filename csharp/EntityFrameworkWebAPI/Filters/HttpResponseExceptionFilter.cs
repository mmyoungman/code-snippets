using EntityFrameworkWebAPI.Exceptions;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Filters;

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
            if (context.ModelState.IsValid) // i.e. no ValidationService.AddModelError calls
            {
                var problemDetails = _problemDetailsFactory.CreateProblemDetails(
                    context.HttpContext,
                    statusCode: (int)exception.StatusCode,
                    detail: exception.Detail);

                context.Result = new ObjectResult(problemDetails);

                context.ExceptionHandled = true;

                return;
            }

            var validationProblemDetails = _problemDetailsFactory.CreateValidationProblemDetails(
                context.HttpContext,
                context.ModelState,
                statusCode: (int)exception.StatusCode,
                detail: exception.Detail
            );
            context.Result = new ObjectResult(validationProblemDetails);

            context.ExceptionHandled = true;
        }
    }
}   