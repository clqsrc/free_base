package free_base

import (
	// "fmt"
	// "sync"
	"runtime/debug"
	"fmt"
	// "os"
	"runtime"
)

//手动管理 gc 的辅助函数，作者对之也理解不深，仅供参考

//--------------------------------------------------------


//永远不进行 GC 操作 //建议只在调试或者 dll 导出函数中使用
func StopGC() { 
    
    //永远不进行 GC 操作
    var old = debug.SetGCPercent(-1);

    //打印一下原来的默认值
    fmt.Println("StopGC()", old);

}//

//手动调用 GC 以释放内存
func GC_(obj interface{}) {  
    
    runtime.GC();

}//


//一个对象被释放时被调用 //其实几乎所有的自动释放语言都有类似功能，例如 js/java
func SetOnFree(obj any, onfree func (interface{})){

    //当 GC 发现 obj 已经不可达（没有变量引用它）时，会在回收之前调用 finalizer(obj)，然后才释放内存。
    //这里比较容易出错，大家可以自己问下 ai 【finalizer interface{} 感觉很容易误导，为何不是 finalizer func(any)】
    //runtime.SetFinalizer(ui_obj, func (interface{})  {

    runtime.SetFinalizer(obj, onfree);

}//

// 用泛型强制指针，编译期类型安全
func SetOnFreeT[T any](obj *T, callback func(*T)) {
    runtime.SetFinalizer(obj, callback)
}



