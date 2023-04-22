package vector

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func Options() {
	//Initialize vectors
	vectorA := []float64{11.0, 5.2, -6.3}
	vectorB := []float64{-5.9, 4.0, 1.2}

	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("Dot Product result:%f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("1.5 * VA:%v\n", vectorA)

	normB := floats.Norm(vectorB, 2)
	fmt.Printf("Norm/length of B:%f", normB)
}
