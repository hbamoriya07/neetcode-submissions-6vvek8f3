func canCompleteCircuit(gas []int, cost []int) int {
    totalGas := 0
    totalCost := 0
    currentTank := 0
    startIndex := 0

    for i := 0; i < len(gas); i++ {
        totalGas += gas[i]
        totalCost += cost[i]
        
        // Add the net gas (gain/loss) for this station to our current tank
        currentTank += gas[i] - cost[i]
        
        // If our tank runs dry, this station (and all before it) are invalid starts
        if currentTank < 0 {
            // Reset the starting candidate to the next station
            startIndex = i + 1
            // Reset our running tank for the new candidate
            currentTank = 0
        }
    }

    // Rule 1: If total gas is less than total cost, it's impossible.
    if totalGas < totalCost {
        return -1
    }

    // Otherwise, the last candidate we found is guaranteed to be correct.
    return startIndex
}