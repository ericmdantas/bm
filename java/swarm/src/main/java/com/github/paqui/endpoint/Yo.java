package com.github.paqui.endpoint;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/")
public class Yo {

	@GET
	@Produces(value=MediaType.TEXT_PLAIN)
	public Response getYo () {
		return Response.ok("yo!").build();
	}
}
