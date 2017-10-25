#include <unistd.h>
#include <stdio.h>

int main() {

    int pid = fork();

    if(pid == 0) {
        printf("This is the parent process! My PID: %d! Parent PID: %d!\n", getpid(), pid);    
    } else {
        printf("This is the child process! My PID: %d! Parent PID: %d!\n", getpid(), pid);
    }

    return 0;
}