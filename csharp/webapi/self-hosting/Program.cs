namespace WebApi.Owin.SelfHosting
{
    using System;

    using Microsoft.Owin.Hosting;

    /// <summary>
    /// The program.
    /// </summary>
    internal class Program
    {
        /// <summary>
        /// The main.
        /// </summary>
        /// <param name="args">
        /// The args.
        /// </param>
        private static void Main(string[] args)
        {
            const string ADDRESS = "http://localhost:3000/";

            using (WebApp.Start<Startup>(ADDRESS))
            {
                Console.WriteLine("Your application is running on " + ADDRESS);
                Console.WriteLine("Press any [Enter] to close the host.");
                Console.ReadLine();
            }

        }
    }
}