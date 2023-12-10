package workload

import (
	"fmt"
	"math"
)

const (
	p5DeveloperCostPerDay = 1554.71
	p4DeveloperCostPerDay = 1198.62
	p3DeveloperCostPerDay = 899.72
	p2DeveloperCostPerDay = 767.17
)

// CalculateWorkloadDistribution 计算工作量分配
func CalculateWorkloadDistribution(totalWorkload float64) {
	bestResult := math.Inf(1)
	bestP5DeveloperDayCount, bestP4DeveloperDayCount, bestP3DeveloperDayCount, bestP2DeveloperDayCount := 0.0, 0.0, 0.0, 0.0

	for i := 0.0; i <= totalWorkload/p5DeveloperCostPerDay; i++ {
		for j := 0.0; j <= totalWorkload/p4DeveloperCostPerDay; j++ {
			for k := 0.0; k <= totalWorkload/p3DeveloperCostPerDay; k++ {
				for l := 0.0; l <= totalWorkload/p2DeveloperCostPerDay; l++ {
					tempWorkload := p5DeveloperCostPerDay*i + p4DeveloperCostPerDay*j + p3DeveloperCostPerDay*k + p2DeveloperCostPerDay*l
					if tempWorkload <= totalWorkload {
						workloadDiff := totalWorkload - tempWorkload
						if workloadDiff < bestResult {
							bestResult = workloadDiff
							bestP5DeveloperDayCount, bestP4DeveloperDayCount, bestP3DeveloperDayCount, bestP2DeveloperDayCount = i, j, k, l
						}
					}
				}
			}
		}
	}

	fmt.Println("=====[工作量分配计算]=====")
	fmt.Printf("目标工作量值为%.2f, 最接近的结果为%.2f * %.f(P5) + %.2f * %.f(P4) + %.2f * %.f(P3) + %.2f * %.f(P2) = %.2f\n", totalWorkload, p5DeveloperCostPerDay, bestP5DeveloperDayCount, p4DeveloperCostPerDay, bestP4DeveloperDayCount, p3DeveloperCostPerDay, bestP3DeveloperDayCount, p2DeveloperCostPerDay, bestP2DeveloperDayCount, p5DeveloperCostPerDay*bestP5DeveloperDayCount+p4DeveloperCostPerDay*bestP4DeveloperDayCount+p3DeveloperCostPerDay*bestP3DeveloperDayCount+p2DeveloperCostPerDay*bestP2DeveloperDayCount)
}
