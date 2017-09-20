#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h> // for creat and open
#include <unistd.h> // for close and unlink
#include <stdio.h>

int main() {
    int flags = 0;
    int mode = 0;
    // creat will overwrite a preexisting file and open it
    //int fd = creat("data", 0666);

    // open opens file, in an access mode: O_RDONLY, O_WRONLY or O_RDWR
    flags = O_RDONLY;
    mode  = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH; // u+rw, g+r, o+r
    int fd = open("data.txt", flags);
    if(fd < 0) {
        // an error occurred
        return -1;
    }

    char string[13]; // read the file in 13 character chunks
    
    int bytesRead;
    while(1) {
        bytesRead = read(fd, &string, 13);
        if(bytesRead < 0) {
            // an error occurred with read
            return -1;
        }
        if(bytesRead == 0) {
            break;
        }
        else if(bytesRead == 13) {
            printf("%d: %s\n", bytesRead, string);
        }
        else { // i.e. remaining < 13 && remaining != 0
            string[bytesRead] = '\0';
            printf("%d: %s\n", bytesRead, string);
        }
    }

    int result = close(fd); // close file
    if(result < 0) {
        // an error occurred with close
        return -1;
    }

    //unlink("data"); // unlink deletes a file
}
