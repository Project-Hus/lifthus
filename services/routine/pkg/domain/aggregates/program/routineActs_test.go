package program

func getValidRoutineActs() RoutineActs {
	ras := RoutineActs{}
	for ord := 1; ord < 7; ord++ {
		ra := CreateRoutineActWithoutDailyRoutine(
			RoutineActOrder(ord),
			"ABCDEF12",
			MainStage,
			10,
			0.7,
		)
		ras = append(ras, ra)
	}
	return ras
}

func getInvalidRoutineActsSets() []RoutineActs {
	rass := []RoutineActs{}
	invalidOrderSet := [][]RoutineActOrder{
		[]RoutineActOrder{2, 3, 4},
		[]RoutineActOrder{1, 2, 4},
		[]RoutineActOrder{1, 2, 3, 3, 4},
	}
	for _, orders := range invalidOrderSet {
		ras := getRoutineActsWithOrders(orders)
		rass = append(rass, ras)
	}
	return rass
}

func getRoutineActWithOrder(order RoutineActOrder) *RoutineAct {
	return CreateRoutineActWithoutDailyRoutine(
		order,
		"ABCDEF12",
		MainStage,
		10,
		0.7,
	)
}

func getRoutineActsWithOrders(orders []RoutineActOrder) RoutineActs {
	ras := RoutineActs{}
	for _, ord := range orders {
		ra := getRoutineActWithOrder(ord)
		ras = append(ras, ra)
	}
	return ras
}
