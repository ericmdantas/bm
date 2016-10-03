namespace webapi
{
    using Microsoft.AspNetCore.Mvc;

    [Route("/")]
    public class IndexController : Controller
    {
        [HttpGet]
        public string Get()
        {
            return "yo";
        }
    }
}
