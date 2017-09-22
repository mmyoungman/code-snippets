    #include <stdlib.h>
    #include <stdio.h>
    #include <dlfcn.h>

    int main() {
        void *handle;
        double (*cosine)(double);
        char *error;

        handle = dlopen ("/lib/libm.so.6", RTLD_LAZY);
        if (!handle) {
            fputs (dlerror(), stderr);
            exit(1);
        }

        cosine = (double (*)(double)) dlsym(handle, "cos");
        if ((error = dlerror()) != NULL)  {
            fputs(error, stderr);
            exit(1);
        }

        printf ("%f\n", (*cosine)(2.0));
        dlclose(handle);
    }