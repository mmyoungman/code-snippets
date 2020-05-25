// I *think* this is libcfree?
// gcc -nostdlib -nodefaultlibs -nostartfiles -static stubstart.S -o main main.c

int main() {
	char *str = "Hello world!\n";
	return 0;
}
