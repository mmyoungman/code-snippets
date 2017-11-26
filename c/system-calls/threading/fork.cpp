#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
// #include <errno.h>

int main() {

    pid_t pid;
    int status = 0;

    // signal(SIGCHLD, SIG_IGN);
    switch(pid = fork()) {
        case -1:
            perror("fork");
            return 1;
        case 0:
            printf("This is the child process! My PID: %d! Fork returned PID: %d! Parent PID: %d!\n", getpid(), pid, getppid());
            return 0;
        default:
            printf("This is the parent process! My PID: %d! Fork returned PID: %d! Parent PID: %d!\n", getpid(), pid, getppid());
            printf("Parent process waiting for child %d to exit...\n", pid);
            while(pid != waitpid(pid, &status, WNOHANG)); // while the child process hasn't exited, wait to reap it
            printf("Parent process exiting...\n");
            return 0;
    }
}