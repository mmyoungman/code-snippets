#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* read_file(FILE *file)
{
    char *str = (char *)malloc(4096);
    char *s = str;
    int len = 0;

    while(fgets(s, len, file))
    {
         len += strlen(s);
         str = (char *)realloc(str, len+4096);
         s = str + len;
    }

    return s;
}

int main(int argc, const char *argv[])
{
    FILE *file;
    file = fopen("fileio.txt", "a+");

    if(file == NULL)
    {
         return 1;
    }

    char *buffer = read_file(file);

    // Read must come first, or fprintf appends to EOF
    //int c;
    //while(1)
    //{
    //    c = fgetc(file);
    //    if(feof(file))
    //        break;
    //    printf("%c", c);
    //}

    //printf("%c", buffer[0]);
    //printf("%s", buffer);

    fprintf(file, "This is a test write.\n");

    fclose(file);

    return 0;
}
