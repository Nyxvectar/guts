/**
 * Author:  Nyxvectar Yan
 * Repo:    guts
 * Created: 07/30/2025
 */

package _map

import (
	"fmt"
	"unsafe"
)

func GutsMap() {
	type mapextra unsafe.Pointer
	// go语言映射的实质就是哈希表
	var a = map[string]string{}
	// a是一个指向hmap结构体的指针, 通过*hmap可知:
	type hmap struct {
		count      int   // 键值对数目
		flags      uint8 // 2^B
		B          uint8
		noverflow  uint16
		hash0      uint32
		buckets    unsafe.Pointer // 桶
		oldbuckets unsafe.Pointer // 旧桶
		nevacuate  uintptr        // 即将迁移的旧桶
		extra      *mapextra
	}
	// 但是这并不是桶的结构, 而是整体的结构
	// 通的结构为bmap, 即buckets map.
	fmt.Println(a)
}
