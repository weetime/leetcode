package math

import "testing"

/**
*	快速乘方法：
* 从右往左推导，采用分治的方法
* x的65次方=x的32次方*x的32次方再*x，x的32次方=x的16次方*x的16次方 依次类推
* 其中x的n次方，就是x的2分之n次方*x的2分之n次方，如果n是奇数，再额外乘以一个x即可
* 这用用递归的方式就能求解出来
 */
func myPow(x float64, n int) float64 {
	var quickMul func(x float64, n int) float64
	quickMul = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		y := quickMul(x, n/2)
		if n%2 == 0 {
			return y * y
		}
		return y * y * x
	}
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, n)
}

func TestMyPow(t *testing.T) {
	res := myPow(2, 3)
	t.Log(res)

	res = myPow(2, -2)
	t.Log(res)
}
