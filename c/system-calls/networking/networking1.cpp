int main() {
    // Set details of server we're going to set up
    addrinfo hints;
    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_UNSPEC; // IPv4 or IPv6
    hints.ai_socktype = SOCK_STREAM; // TCP
    hints.ai_flags = AI_PASSIVE; // use my IP

    // Try to set up server structures
    addrinfo *servinfo;
    int rv;
    if((rv = getaddrinfo(NULL, "9000", &hints, &servinfo) != 0) {
       fprintf(stderr, "getaddrinfo error: %s\n", gai_strerror(rv));
       return 1;
    }

    // Cycle through linked list of sockaddrs that getaddrinfo returned
    addrinfo *p;
    int sockfd;
    int yes = 1;
    for(p = servinfo; p != NULL; p = p->ai_next) {
        // Try to get file descriptor
        if((sockfd = socket(p->ai_family, p->ai_socktype, p->ai_protocol)) == -1) {
            perror("server: socket");
            continue; // if there is an error, try next entry in linked list
        }

        // Lose "Address already in use error message"
        if(setsockopt(sockfd, SOL_SOCKET, SOREUSEADDR, &yes, sizeof(int)) == -1) {
            perror("setsockopt");
            return 1;
        }

        // Associate socket with port
        if(bind(sockfd, p->ai_addr, p->ai_addrlen) == -1) {
            close(sockfd);
            perror("server: bind");
            continue;
        }
        break; // if we reach this, we have successfully binded
    }

    freeaddrinfo(servinfo); // no longer needed?

    // If failed to bind, exit
    if(p == NULL) {
        fprintf(stderr, "server: failed to bind\n");
        return 1;
    }


}
