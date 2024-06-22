using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Filters;

public class InvalidModelStateFilter : IActionFilter, IOrderedFilter
{
    private readonly ProblemDetailsFactory _problemDetailsFactory;
    public int Order => int.MaxValue - 10;

    public InvalidModelStateFilter(
        ProblemDetailsFactory problemDetailsFactory)
        => _problemDetailsFactory = problemDetailsFactory;

    public void OnActionExecuting(ActionExecutingContext context) {}

    public void OnActionExecuted(ActionExecutedContext context) 
    {
        if (!context.ModelState.IsValid)
        {
            var validationProblemDetails = _problemDetailsFactory.CreateValidationProblemDetails(
                context.HttpContext,
                context.ModelState,
                statusCode: (int)HttpStatusCode.BadRequest);

            context.Result = new ObjectResult(validationProblemDetails);
        }
    }
}
