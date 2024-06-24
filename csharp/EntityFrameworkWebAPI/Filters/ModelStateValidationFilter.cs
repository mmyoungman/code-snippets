using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Filters;

public class ModelStateValidationFilter : IActionFilter, IOrderedFilter
{
    private readonly ProblemDetailsFactory _problemDetailsFactory;
    public int Order => int.MaxValue - 10;

    public ModelStateValidationFilter(
        ProblemDetailsFactory problemDetailsFactory)
        => _problemDetailsFactory = problemDetailsFactory;

    public void OnActionExecuting(ActionExecutingContext context) {}

    public void OnActionExecuted(ActionExecutedContext context) 
    {
        // This filter is only necessary because ApiControllerAttribute's 
        // ModelStateInvalidFilter only executes if context.Result == null (I believe).
        // We want to return a BadRequest any time there is a model state error.
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
