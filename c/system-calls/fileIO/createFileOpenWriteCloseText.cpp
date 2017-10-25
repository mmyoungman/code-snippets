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
    //int fd = open("data", O_RDWR);
    flags = O_RDWR | O_APPEND | O_CREAT; // r+w permissions, write appends, and create if doesn't already exist
    mode  = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH; // u+rw, g+r, o+r
    int fd = open("data.txt", flags, mode);
    if(fd < 0) {
        // an error occurred with open
        return 1;
    }

    int result;
    // write "Hello world!" to file -- it's 12 characters!
    for(int i = 0; i < 3; i++) {
        result = write(fd, "Hello world!", 12);
        if(result < 0) {
            // an error occurred with write
            return 1;
        }
    }

    result = fsync(fd); // ensure data is written to disk
    if(result < 0) {
        // an error occurred with fsync
        return 1;
    }
    // fsync of dir is needed to ensure it is written to disk
    flags = O_RDONLY | O_DIRECTORY; // read only, must be a directory
    int fdDir = open(".", flags);
    result = fsync(fdDir);
    if(result < 0) {
        // an error occurred with fsync
        return 1;
    }

    result = close(fdDir); // close dir
    if(result < 0) {
        // an error occurred with close
        return 1;
    }
    result = close(fd); // close file
    if(result < 0) {
        // an error occurred with close
        return 1;
    }

    //unlink("data"); // unlink deletes a file
}
