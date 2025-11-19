package indexdb

type IndexDB interface {
	AddDocument(string, any) error
	//Search()
}

func MergeIndexDBresults(vrs []Result, brs []BM25Result) map[string]float64 {
	merged := make(map[string]float64)

	//Ranks scores from vector results
	for rank, vr := range vrs {
		score := 1.0 / float64(rank+1)
		merged[vr.Content] += score
	}

	//Ranks scores from BM25 results
	for rank, br := range brs {
		score := 1.0 / float64(rank+1)
		merged[br.Document["content"].(string)] += score
	}

	return merged

}
