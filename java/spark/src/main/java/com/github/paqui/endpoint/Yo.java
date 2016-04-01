package com.github.paqui.endpoint;

import static spark.Spark.get;
import static spark.Spark.port;

public class Yo {
	public static void main(String[] args) {
		port(configPort(args));
        get("/", (req, res) -> "yo!");
    }

	private static int configPort(String[] args) {
		int port = 8080;
		for (int i = 0; i < args.length; i++) {
			if ("-port".equals(args[i])) {
				port = new Integer(args[i+1]);
			}
		}
		return port;
	}
}
