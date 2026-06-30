// package main
package free_base

import (
    "fmt"
    "sync"
)

//修改自 _src_map_textfile_ios\class_obj_new.go //目标是心智负担最低的对象手工释放实现


//--------------------------------------------------------
//正式的安全实现

var g_objs_mutex sync.Mutex

// 使用 map 存储存活的窗体
var g_objs = make(map[interface{}]struct{});
// var g_objs = make(map[*interface{}]struct{}); //这里有个坑，据说这样是无法持有对象的指针的

//如果 obj 是结构体的话有个坑，如果 new_() 之后结构体改变了，就会找不到原来的 key !!!
//所以实际上传入的必须是指针（或者类似指针的引用类型，不能是单纯的 struct）
func new_(obj interface{}) { //是不能写成 new_(obj * interface{}) 的
   	g_objs_mutex.Lock()
    
    defer g_objs_mutex.Unlock()

	//--------------------------------------------------------
	//理论上说可以判断是否传入的是指针，不过 ai 认为对性能影响太大，不建议。然后 golang 目前又没有办法从语法上进行限制，所以只能是这样了
	//参考 https://chat.z.ai/c/04d537a4-ea73-4d22-a4db-967bc1c04d37
	// ⚠️ 仅在开发调试时开启，生产环境建议注释掉，以免影响性能
    // if reflect.TypeOf(obj).Kind() != reflect.Ptr {
    //     // 这里可以 panic，也可以只打印 log
    //     fmt.Printf("[ERROR] new_() 强烈建议传入指针，你传入的是值类型: %v\n", reflect.TypeOf(obj))
    //     // panic("new_() requires a pointer") 
    // }
	//--------------------------------------------------------

    // 将窗体放入 map 管理
    g_objs[obj] = struct{}{}
    
    //fmt.Printf("[INFO] 窗体已创建并注册，当前存活数量: %d\n", len(g_forms))
}//

//obj 可以是指针，只要和 new_() 的时候的一致就行了
//如果 obj 是结构体的话有个坑，如果 new_() 之后结构体改变了，就会找不到原来的 key !!!
//所以实际上传入的必须是指针 
func delete_(obj interface{}) {  //这里倒是可以写成 delete_(obj * interface{})，不过并不好

    g_objs_mutex.Lock()
    defer g_objs_mutex.Unlock()

    // 检查 map 中是否存在
    if _, exists := g_objs[obj]; exists {
        // 从 map 中移除，解除引用
        delete(g_objs, obj)

        
        //fmt.Printf("[INFO] 窗体已手动销毁，当前存活数量: %d\n", len(g_forms))
    } else {
        fmt.Println("[WARN] delete_() 尝试销毁一个不存在的或已销毁的窗体")
    }
}//

