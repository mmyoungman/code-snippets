// Created by Mark Youngman 29 November 2017
// Opens a port at localhost:9393
// Can connect to it as follows: "telnet --no-login localhost 9393" and send a string of text

#include <stdio.h>
#include <string.h> // for memset()
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h> // for close()

#define MYPORT "9393"
#define BACKLOG 10

int main() {

    int status;
    int sockfd;
    struct addrinfo hints;
    struct addrinfo *servinfo;

    int newfd;
    struct sockaddr_storage clientaddr;
    socklen_t clientaddrsize;

    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_UNSPEC;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_flags = AI_PASSIVE;

    status = getaddrinfo(NULL, MYPORT, &hints, &servinfo);
    if(status != 0) {
        printf("getaddrinfo() error\n");
        printf("%s\n", gai_strerror(status));
        return 1;
    }

    sockfd = socket(servinfo->ai_family, servinfo->ai_socktype, servinfo->ai_protocol);
    if(sockfd < 0) {
        printf("socket() error\n");
        perror("socket");
        return 1;
    }

    // use setsockopt() here to handle "Address already in use" case when bind() is called
    // int yes = 1;
    // status = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(yes));
    // if(status < 0) {
    //     printf("setsockopt() error\n");
    //     return 1;
    // }

    status = bind(sockfd, servinfo->ai_addr, servinfo->ai_addrlen);
    if(status < 0) {
        perror("bind");
        return 1;
    }

    status = listen(sockfd, BACKLOG);
    if(status < 0) {
        printf("listen() error\n");
        perror("listen");
        return 1;
    }
    printf("Waiting for incoming connections...\n");

    newfd = accept(sockfd, (struct sockaddr*)&clientaddr, &clientaddrsize); 
    if(newfd < 0) {
        printf("accept() error\n");
        perror("accept");
        return 1;
    }
    printf("Connection accepted!\n");

    int bufferLen = 256;
    char buffer[bufferLen];
    status = recv(newfd, &buffer, bufferLen, 0);
    if(status < 0) {
        perror("recv");
    }

    buffer[255] = '\0'; // Prevent buffer overflow
    printf("Telnet input: %s\n", buffer);

    //shutdown(sockfd, 2);
    close(sockfd);
    freeaddrinfo(servinfo);
}