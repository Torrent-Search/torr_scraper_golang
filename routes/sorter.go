package routes

import (
	"sort"
	"strconv"
)

func SortInfos(in *TorrentRepo, seeders bool, leechers bool, seedAscend bool, leechAsend bool) {
	if seeders {
		if seedAscend {
			sort.SliceStable(in.Data, func(i, j int) bool { return getInt((*in).Data[i].Seeders) > getInt((*in).Data[j].Seeders) })
		} else {
			sort.SliceStable(in.Data, func(i, j int) bool { return getInt((*in).Data[i].Seeders) < getInt((*in).Data[j].Seeders) })
		}
	}
	if leechers {
		if seedAscend {
			sort.SliceStable(in, func(i, j int) bool { return getInt((*in).Data[i].Leechers) > getInt((*in).Data[j].Leechers) })
		} else {
			sort.SliceStable(in, func(i, j int) bool { return getInt((*in).Data[i].Leechers) < getInt((*in).Data[j].Leechers) })
		}
	}
}

func (s *TorrentRepo) Len() int {
	return len((*s).Data)
}

func (s *TorrentRepo) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func getInt(in string) int {
	i, _ := strconv.Atoi(in)
	return i
}
