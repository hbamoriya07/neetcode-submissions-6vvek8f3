func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    cars := make([][2]float64, n)

    for i := 0 ; i < n; i++ {
        cars[i] = [2]float64{float64(position[i]),float64(speed[i])}
    }

    sort.Slice(cars, func(i,j int) bool {
        return cars[i][0] > cars [j][0]
    })

    fleets := 0 
    maxTime := 0.0

    for _, car := range cars {
        time := (float64(target) - car[0]) / car[1]

        if time > maxTime {
            fleets++
            maxTime = time
        }
    }

    return fleets
}

// pos, speed ,  target 
