using FluentValidation;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Infrastructure;
using Microsoft.AspNetCore.Mvc.ModelBinding;

namespace EntityFrameworkWebAPI.Filters;

public class FluentValidationFilter(
    IServiceProvider serviceProvider,
    ProblemDetailsFactory problemDetailsFactory) : IAsyncActionFilter
{
    public async Task OnActionExecutionAsync(ActionExecutingContext context, ActionExecutionDelegate next)
    {
        var modelState = new ModelStateDictionary();

        foreach (var argument in context.ActionArguments.Values)
        {
            if (argument is null)
                continue;

            var validatorType = typeof(IValidator<>).MakeGenericType(argument.GetType());
            if (serviceProvider.GetService(validatorType) is not IValidator validator)
                continue;

            var result = await validator.ValidateAsync(
                new ValidationContext<object>(argument),
                context.HttpContext.RequestAborted);

            foreach (var failure in result.Errors)
                modelState.AddModelError(failure.PropertyName, failure.ErrorMessage);
        }

        if (!modelState.IsValid)
        {
            var problemDetails = problemDetailsFactory.CreateValidationProblemDetails(
                context.HttpContext,
                modelState,
                statusCode: StatusCodes.Status400BadRequest);

            context.Result = new ObjectResult(problemDetails)
            {
                StatusCode = StatusCodes.Status400BadRequest,
            };

            return;
        }

        await next();
    }
}
