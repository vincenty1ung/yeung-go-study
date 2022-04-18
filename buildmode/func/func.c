
// file func.c
#include <stdio.h>
#include "libfunc.h"

int main() {
  printf("This is a C Application.\n");
  GoString name = {(char*)"Jane", 4};
  SayHello(name);
  GoSlice buf = {(void*)"Jane", 4, 4};
  SayHelloByte(buf);
  SayBye();
  Print((char*)"Jane");
  //Add(1,1);
  int i  = Add(1,1);
  printf("%2d",i);
  return 0;
 }

