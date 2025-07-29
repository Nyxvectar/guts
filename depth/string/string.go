/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/29/2025
 */

package string

func GutsString() {
	// 为了判定字符，给每个字符设置一个编码。
	// 出于读取的需要，派定一个不同区间适用的格式，一般标头。
	// 于是得到诸如UTF-8编码，这也是Go语言的默认编码格式。
	// 但是如此并解决不了识别结尾的问题，故存在字节数尾。
	var stringDemo = "Hello, Golang!"
	// 在Go语言当中的字符串默认是不会被修改的，
	// 上述这种方法所得到的string将分配进入只读内存段，
	// 这也是为什么我们不可以直接 sD[] = ...修改其值的缘由。
	// 为了避免公用内存段修改造成不可预估的影响，
	// Go语言中string值的修改应该是整体赋予新的值，
	// 例如下面这一行：
	stringDemo = "Hello, GNU!"
	// 这样并不会修改原来的内存，而是赋予了新的内存段
	// 除此以外，还有另外一种修改的方式，强制转换为slice：
	var stringExam = ([]byte)(stringDemo)
	stringExam[2] = 'a'
	// 这里不用“”,需要rune(符文)类型
	// 此时可以实现同样的效果
	print(stringDemo, stringExam)
}
