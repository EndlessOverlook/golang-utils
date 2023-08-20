package collectionutils

import (
    "fmt"
    "sync"
)

func useSyncMap() {
    // sync.Map不能使用make创建
    var scene sync.Map
    // 将键值对保存到sync.Map
    // sync.Map将键和值以interface {}类型进行保存
    scene.Store("greece", 97)
    scene.Store("london", 100)
    scene.Store("egypt", 200)
    // 从sync.Map中根据键取值
    fmt.Println(scene.Load("london"))
    // 根据键删除对应的键值对
    scene.Delete("london")
    // 遍历所有sync.Map中的键值对
    // 遍历需要提供一个匿名函数，参数为k, v，类型为interface {}，每次Range()在遍历一个元素时，都会调用这个匿名函数把结果返回
    scene.Range(func(k, v interface{}) bool {
        fmt.Println("iterate:", k, v)
        return true
    })
}
