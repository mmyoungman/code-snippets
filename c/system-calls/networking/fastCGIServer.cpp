#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <string.h> // for memset

int main() {
    addrinfo hints;
    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_UNSPEC; // IPv4 or v6, no matter
    hints.ai_socktype = SOCK_STREAM; // TCP
    hints.ai_flags = AI_PASSIVE; // use my IP

    addrinfo *servinfo;
    int rv = getaddrinfo(NULL, "9000", &hints, &servinfo);
    if(rv < 0) {
        // getaddrinfo error
        return 1;
    }

    addrinfo *p;
    for(p = servinfo; p != NULL; p = p->ai_next) {
        int sockfd = socket(p->ai_family, p->ai_socktype, p->ai_protocol);
        if(sockfd < 0) {
            // error with socket()
            continue;
        }

        int yes = 1;
        int result;

        result = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
        if(result < 0) {
            // error with setsockopt
            return -1;
        }

        result = bind(sockfd, p->ai_addr, p->ai_addrlen);
        if(result < 0) {
            close(sockfd);
            continue;
        }

        break;
    }

    freeaddrinfo(servinfo);

}
