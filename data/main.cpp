#include<iostream>
#include<stdio.h>
#include<cstring>
using namespace std;

int main() {
	char ch[] = { 'a', 'b', '\0' };
	printf("%send\n", ch);
        printf("length: %d\n", strlen(ch));
	printf("end\n");
}
