namespace WebApi.Owin.SelfHosting
{
    using System.Web.Http;

    using global::Owin;

    /// <summary>
    /// The startup class.
    /// </summary>
    public class Startup
    {
        /// <summary>
        /// Configures the application.
        /// </summary>
        /// <param name="appBuilder">
        /// The app builder.
        /// </param>
        public void Configuration(IAppBuilder appBuilder)
        {
            var config = new HttpConfiguration();
            config.Routes.MapHttpRoute("DefaultApi", "{controller}");

            appBuilder.UseWebApi(config);
        }
    }
}