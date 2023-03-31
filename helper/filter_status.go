package helper

func StatusDiscountProposalByLevel(level string) (string, error) {
	var status string
	switch level {
	case "MR":
		status = " TEMPORARY,INPUT,CONFIRM "
	case "SPV":
		status = "CONFIRM"
	case "ASM":
		status = "CONFIRM SPV"
	case "FSM":
		status = "CONFIRM ASM"
	case "NSM":
		status = "CONFIRM FSM"
	case "GSM":
		status = "CONFIRM NSM"
	}
	return status, nil
}
