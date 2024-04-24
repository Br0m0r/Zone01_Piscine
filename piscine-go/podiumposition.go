package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i, j := 0, len(podium)-1; j > i; i, j = i+1, j-1 {
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
