namespace WebApi.Owin.SelfHosting
{
    using System.Web.Http;

    /// <summary>
    /// The index controller.
    /// </summary>
    public class IndexController : ApiController
    {
        /// <summary>
        /// Gets yo.
        /// </summary>
        /// <returns>
        /// yo string.
        /// </returns>
        public string Get()
        {
            return "yo";
        }
    }
}