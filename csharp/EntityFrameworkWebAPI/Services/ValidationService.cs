using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace EntityFrameworkWebAPI.Services;

public interface IValidationService
{
    void AddModelError(string key, string errorMessage);
}

public class ValidationService : IValidationService
{
    public ValidationService(IActionContextAccessor actionContextAccessor)
    {
        _actionContextAccessor = actionContextAccessor;
    }

    private readonly IActionContextAccessor _actionContextAccessor;

    public void AddModelError(string key, string errorMessage)
    {
        _actionContextAccessor.ActionContext!.ModelState.AddModelError(key, errorMessage);
    }
}
