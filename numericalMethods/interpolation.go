package numericalMethods

import "errors"

// NewtonForwardDividedDifference is for calculating the coefficients for newton's forward divided-difference interpolating polynomial
func NewtonForwardDividedDifference(xValues []float64, functionValues []float64) ([]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	previousValue := float64(0)
	tempValue := float64(0)

	for i := 1; i < size; i++ {
		previousValue = functionValues[i-1]
		for j := i; j < size; j++ {
			tempValue = functionValues[j]
			functionValues[j] = (functionValues[j] - previousValue) / (xValues[j] - xValues[j-i])
			previousValue = tempValue
		}
	}

	return functionValues, nil
}

// NewtonDividedDifference is for calculating the coefficients for newton's divided-difference interpolating polynomial
func NewtonDividedDifference(xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	tableValues := make([][]float64, size)

	for i := 0; i < size; i++ {
		tableValues[i] = make([]float64, i+1)
		tableValues[i][0] = functionValues[i]
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= i; j++ {
			tableValues[i][j] = (tableValues[i][j-1] - tableValues[i-1][j-1]) / (xValues[i] - xValues[i-j])
		}
	}

	return tableValues, nil
}

// NevilleIterated is for determining the table values of neville iterated interpolation
func NevilleIterated(valueToApprox float64, xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	tableValues := make([][]float64, size)

	for i := 0; i < size; i++ {
		tableValues[i] = make([]float64, i+1)
		tableValues[i][0] = functionValues[i]
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= i; j++ {
			tableValues[i][j] = ((valueToApprox-xValues[i-j])*tableValues[i][j-1] - (valueToApprox-xValues[i])*tableValues[i-1][j-1]) / (xValues[i] - xValues[i-j])
		}
	}

	return tableValues, nil
}

// Hermite is for determining the coefficients of the hermite interpolation polynomial
func Hermite(xValues []float64, functionValues []float64, dfunctionValues []float64) ([]float64, error) {
	size := len(xValues)

	if size != len(functionValues) || size != len(dfunctionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	valueDoubleSet := make([]float64, 2*size)

	tableValues := make([][]float64, 2*size)

	for i := 0; i < size; i++ {
		tableValues[2*i] = make([]float64, 2*i+1)
		tableValues[2*i+1] = make([]float64, 2*(i+1))
	}

	valueDoubleSet[0] = xValues[0]
	valueDoubleSet[1] = xValues[0]

	tableValues[0][0] = functionValues[0]
	tableValues[1][0] = functionValues[0]
	tableValues[1][1] = dfunctionValues[0]

	for i := 1; i < size; i++ {
		valueDoubleSet[2*i] = xValues[i]
		valueDoubleSet[2*i+1] = xValues[i]

		tableValues[2*i][0] = functionValues[i]
		tableValues[2*i+1][0] = functionValues[i]
		tableValues[2*i+1][1] = dfunctionValues[i]

		tableValues[2*i][1] = (tableValues[2*i][0] - tableValues[2*i-1][0]) / (valueDoubleSet[2*i] - valueDoubleSet[2*i-1])
	}

	for i := 2; i < 2*size; i++ {
		for j := 2; j <= i; j++ {
			tableValues[i][j] = (tableValues[i][j-1] - tableValues[i-1][j-1]) / (valueDoubleSet[i] - valueDoubleSet[i-j])
		}
	}

	solutionSet := make([]float64, 2*size)

	for i := 0; i < 2*size; i++ {
		solutionSet[i] = tableValues[i][i]
	}

	return solutionSet, nil
}

// NaturalCubicSpline is used for finding the coefficients solution set of the natural cubic spline
func NaturalCubicSpline(xValues []float64, functionValues []float64) ([][]float64, error) {
	size := len(xValues)

	if size != len(functionValues) {
		return nil, errors.New("Length of x values array and function values array does not match")
	}

	stepLengthSet := make([]float64, size-1)

	for i := 0; i < size-1; i++ {
		stepLengthSet[i] = xValues[i+1] - xValues[i]
	}

	alpha := make([]float64, size-1)

	for i := 1; i < size-1; i++ {
		// fmt.Printf("%v", i)
		alpha[i] = 3.0*(functionValues[i+1]-functionValues[i])/stepLengthSet[i] - 3.0*(functionValues[i]-functionValues[i-1])/stepLengthSet[i-1]
	}

	solvingSetA := make([]float64, size)
	solvingSetB := make([]float64, size-1)
	solvingSetC := make([]float64, size)

	solvingSetA[0] = 1
	solvingSetB[0] = 0
	solvingSetC[0] = 0

	for i := 1; i < size-1; i++ {
		solvingSetA[i] = 2.0*(xValues[i+1]-xValues[i-1]) - stepLengthSet[i-1]*solvingSetB[i-1]
		solvingSetB[i] = stepLengthSet[i] / solvingSetA[i]
		solvingSetC[i] = (alpha[i] - stepLengthSet[i-1]*solvingSetC[i-1]) / solvingSetA[i]
	}

	solvingSetA[size-1] = 1
	solvingSetC[size-1] = 0

	var solutionSetA []float64
	solutionSetA = functionValues[:size-1]
	solutionSetB := make([]float64, size-1)
	solutionSetC := make([]float64, size)
	solutionSetD := make([]float64, size-1)

	solutionSetC[size-1] = 0

	sizeSolutionSet := len(solutionSetB)

	for i := sizeSolutionSet - 1; i >= 0; i-- {
		solutionSetC[i] = solvingSetC[i] - solvingSetB[i]*solutionSetC[i+1]
		solutionSetB[i] = (functionValues[i+1]-functionValues[i])/stepLengthSet[i] - stepLengthSet[i]*(solutionSetC[i+1]+2*solutionSetC[i])/3.0
		solutionSetD[i] = (solutionSetC[i+1] - solutionSetC[i]) / (3.0 * stepLengthSet[i])
	}

	solutionSetC = solutionSetC[:size-1]

	solutionTable := [][]float64{}

	solutionTable = append(solutionTable, solutionSetA)
	solutionTable = append(solutionTable, solutionSetB)
	solutionTable = append(solutionTable, solutionSetC)
	solutionTable = append(solutionTable, solutionSetD)

	return solutionTable, nil
}