package free_base

import (
    // "fmt"
    // "sync"
)

//这个文件只是 class_obj_new.go 的导出封装。
//如果只是自己用的个人项目，其实直接复制 class_obj_new.go 到项目下改下包名是最好的。性能和函数名都最优 :)
//不过我自己还是比较严格用独立的包实现，因为我的项目太多了，复制出去反而难维护


//--------------------------------------------------------
//正式的安全实现




//如果 obj 是结构体的话有个坑，如果 new_() 之后结构体改变了，就会找不到原来的 key !!!
//所以实际上传入的必须是指针（或者类似指针的引用类型，不能是单纯的 struct）
func New_(obj interface{}) { //是不能写成 new_(obj * interface{}) 的

    new_(obj);

}//

//据说 go 1.18 之后的泛型可以解决这个问题
func New_T[T any](obj *T){

    // new_(obj);
    New_(obj);

}//

//obj 可以是指针，只要和 new_() 的时候的一致就行了
//如果 obj 是结构体的话有个坑，如果 new_() 之后结构体改变了，就会找不到原来的 key !!!
//所以实际上传入的必须是指针 
func Delete_(obj interface{}) {  //这里倒是可以写成 delete_(obj * interface{})，不过并不好

    delete_(obj);

}//

//据说 go 1.18 之后的泛型可以解决这个问题
func Delete_T[T any](obj *T){

    // new_(obj);
    Delete_(obj);

}//


//程序结束时可以看下还有哪些没释放
func DebugObjList_() any {  //这里倒是可以写成 delete_(obj * interface{})，不过并不好

    return g_objs;

}//

