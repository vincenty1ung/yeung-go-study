> 生成 c文件

> go build -trimpath -buildmode=c-archive -o libfunc.a(Go语言导出C静态库)
> go build -trimpath -buildmode=c-shared -o libfunc.so(Go语言导出C动态库)



> gcc -o func func.c func.a -pthread
>
> cd ./func
> gcc -c -o func.o func.c
> ar rcs libfunc.a func.o
> gcc -shared -o libfunc.so number.c