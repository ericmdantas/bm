package undertow.servlet.io;

import io.undertow.Undertow;
import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.Headers;

public class Main {

	public static void main(final String[] args) {
        Undertow server = Undertow.builder()
                .addHttpListener(9090, "0.0.0.0")
                .setHandler(new HttpHandler() {
                    @Override
                    public void handleRequest(final HttpServerExchange exchange) throws Exception {
                        exchange.getResponseHeaders().put(Headers.CONTENT_TYPE, "application/json; charset=UTF-8");
                        exchange.getResponseSender().send("Yo!");
                    }
                }).build();
        server.start();
    }
}
