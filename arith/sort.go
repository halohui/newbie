package arith

import (
	"math/rand"
)

/****************************************************************************
* 功能描述: 直接插入排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 时间复杂度为O(n^2),空间复杂度为O(1)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 16:22:13          0.00           cxh                创建
*
*****************************************************************************/
func InsertSort(data []int) {
	var i, j, cur int
	length := len(data)

	for i = 1; i < length; i++ {
		cur = data[i]
		for j = i - 1; j >= 0; j-- {
			if data[j] > cur {
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = cur
	}
}

/****************************************************************************
* 功能描述: 折半插入排序算法
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 时间复杂度为O(n^2),空间复杂度为O(1)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 16:27:55          0.00           cxh                创建
*
*****************************************************************************/
func BinInsertSort(data []int) {
	var cur, start, mid, end int
	length := len(data)
	for i := 1; i < length; i++ {
		cur = data[i]
		start = 0
		end = i - 1
		for start <= end {
			mid = start + (end-start)/2
			if data[mid] == cur {
				end = mid
				break
			} else if data[mid] > cur {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		for j := i - 1; j >= end+1; j-- {
			data[j+1] = data[j]
		}
		data[end+1] = cur
	}
}

/****************************************************************************
* 功能描述: shell排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 16:34:54          0.00           cxh                创建
*
*****************************************************************************/
func shellCore(data [] int, step int) {
	var i, j, cur int
	length := len(data)
	for i = step; i < length; i++ {
		cur = data[i]
		for j = i - step; j >= 0; j -= step {
			if data[j] > cur {
				data[j+step] = data[j]
			} else {
				break
			}
		}
		data[j+step] = cur
	}
}

func ShellSort(data [] int) {
	step := len(data) / 2
	for step > 0 {
		shellCore(data, step)
		step /= 2
	}
}

/****************************************************************************
* 功能描述: 冒泡排序,大数沉底
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 时间复杂度为O(n^2),空间复杂度为O(1)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 17:15:34          0.00           cxh                创建
*
*****************************************************************************/
func BubbleSortBig(data []int) {
	length := len(data)

	for i := 1; i < length; i++ {
		flag := false
		for j := 0; j < length-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

/****************************************************************************
* 功能描述: 冒泡排序,小数上浮
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 时间复杂度为O(n^2),空间复杂度为O(1)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 17:25:25          0.00           cxh                创建
*
*****************************************************************************/
func BubbleSortSmall(data [] int) {
	length := len(data)
	for i := 1; i < length; i++ {
		flag := false
		for j := length - 1; j >= i; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

/****************************************************************************
* 功能描述: 快速排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 不稳定排序算法，时间复杂度为O(nlog(n))，空间复杂度为log(n)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 17:30:10          0.00           cxh                创建
*
*****************************************************************************/
func partition(data [] int, start int, end int) int {
	small := start - 1
	index := rand.Intn(end-start) + start
	data[index], data[end] = data[end], data[index]

	for i := start; i < end; i++ {
		if data[i] < data[end] {
			small++
			if small < i {
				data[small], data[i] = data[i], data[small]
			}
		}
	}
	small++
	data[small], data[end] = data[end], data[small]
	return small
}

func QuickSort(data []int,start int,end int) {
	if start<end{
		pivot :=partition(data,start,end)

		if pivot>start{
			QuickSort(data,start,pivot-1)
		}

		if end > pivot {
			QuickSort(data,pivot+1,end)
		}
	}
}

/****************************************************************************
* 功能描述: 选择排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明: 时间复杂度为O(n^2),空间复杂度为O(1)

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-11-14 18:56:46          0.00           cxh                创建
*
*****************************************************************************/
