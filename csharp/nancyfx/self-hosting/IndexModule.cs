// --------------------------------------------------------------------------------------------------------------------
// <copyright file="IndexModule.cs" company="bm">
//     bm.
// </copyright>
// <summary>
// The index module.
// </summary>
// --------------------------------------------------------------------------------------------------------------------

namespace NancyApplication.SelfHosting
{
    using Nancy;

    /// <summary>
    /// The index module.
    /// </summary>
    public class IndexModule : NancyModule
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="IndexModule"/> class.
        /// </summary>
        public IndexModule()
        {
            this.Get["/"] = parameters => "yo";
        }
    }
}