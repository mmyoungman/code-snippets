#include <stdio.h>
#include <iostream>

#include <ixwebsocket/IXNetSystem.h>
#include <ixwebsocket/IXUserAgent.h>
#include <ixwebsocket/IXWebSocket.h>


// TODO: Create python script for wss local server for testing
// https://github.com/dpallot/simple-websocket-server#tlsssl-example
#define SERVER_URL "wss://ws.ifelse.io"

int main(int argc, char *argv[]) {
  ix::initNetSystem();

  ix::WebSocket webSocket;

  ix::SocketTLSOptions options;
  options.tls = true;
  options.certFile = "";
  options.keyFile = "";
  options.caFile = "NONE";
  webSocket.setTLSOptions(options);

  std::string url(SERVER_URL);
  webSocket.setUrl(url);

  printf("Connecting to %s...\n", url.c_str());

  webSocket.setOnMessageCallback([](const ix::WebSocketMessagePtr &msg) {
    if (msg->type == ix::WebSocketMessageType::Message) {
      printf("Recieved message: %s\n> ", msg->str.c_str());
      fflush(stdout);
    } else if (msg->type == ix::WebSocketMessageType::Open) {
      printf("Connection established\n> ");
      fflush(stdout);
    } else if (msg->type == ix::WebSocketMessageType::Error) {
      // Maybe SSL is not configured properly
      printf("Connection error: %s\n> ", msg->errorInfo.reason.c_str());
      fflush(stdout);
    }
  });

  // Now that our callback is setup, we can start our background thread and
  // receive messages
  webSocket.start();

  // Display a prompt
  printf("> ");
  fflush(stdout);

  std::string text;
  // Read text from the console and send messages in text mode.
  // Exit with Ctrl-D on Unix or Ctrl-Z on Windows.
  while (std::getline(std::cin, text)) {
    webSocket.send(text);
    printf("> ");
    fflush(stdout);
  }

  return 0;
}
