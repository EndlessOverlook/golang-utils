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
	proximalWorkload := math.Inf(1)
	proximalP5DeveloperDayCount, proximalP4DeveloperDayCount, proximalP3DeveloperDayCount, proximalP2DeveloperDayCount := 0, 0, 0, 0

	for i := 0; i <= int(totalWorkload/p5DeveloperCostPerDay); i++ {
		for j := 0; j <= int(totalWorkload/p4DeveloperCostPerDay); j++ {
			for k := 0; k <= int(totalWorkload/p3DeveloperCostPerDay); k++ {
				for l := 0; l <= int(totalWorkload/p2DeveloperCostPerDay); l++ {
					tempWorkload := p5DeveloperCostPerDay*float64(i) + p4DeveloperCostPerDay*float64(j) + p3DeveloperCostPerDay*float64(k) + p2DeveloperCostPerDay*float64(l)
					if tempWorkload <= totalWorkload {
						workloadDiff := totalWorkload - tempWorkload
						if workloadDiff < proximalWorkload {
							proximalWorkload = workloadDiff
							proximalP5DeveloperDayCount, proximalP4DeveloperDayCount, proximalP3DeveloperDayCount, proximalP2DeveloperDayCount = i, j, k, l
						}
					}
				}
			}
		}
	}

	bestWorkload := p5DeveloperCostPerDay*float64(proximalP5DeveloperDayCount) + p4DeveloperCostPerDay*float64(proximalP4DeveloperDayCount) + p3DeveloperCostPerDay*float64(proximalP3DeveloperDayCount) + p2DeveloperCostPerDay*float64(proximalP2DeveloperDayCount)

	fmt.Println("=====[工作量分配计算-开始]=====")
	fmt.Printf("目标工作量值为[%.2f]\n", totalWorkload)
	fmt.Printf("最接近的结果为[%.2f] = %.2f * %d(P5) + %.2f * %d(P4) + %.2f * %d(P3) + %.2f * %d(P2)\n", bestWorkload, p5DeveloperCostPerDay, proximalP5DeveloperDayCount, p4DeveloperCostPerDay, proximalP4DeveloperDayCount, p3DeveloperCostPerDay, proximalP3DeveloperDayCount, p2DeveloperCostPerDay, proximalP2DeveloperDayCount)
	fmt.Println("=====[工作量分配计算-结束]=====")
}
