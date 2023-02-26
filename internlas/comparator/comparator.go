// Package comparator provides ...
package comparator

// compareMaps compare two maps checking whether the two maps have exactly the same values.
// If one of the maps has the same values than the other, but also other values which
// the first one doesn't, it'll return false.
// For example:
//
//	map1 := map[string]any{
//		fruit:"banana",
//		price: 12.3
//	}
//
//	map2 := map[string]any{
//		fruit:"banana",
//		price: 12.3
//	}
//
//	map3 := map[string]any{
//		fruit:"banana",
//		price: 12.3,
//	 color: yellow
//	}
//
// compareMaps(map1,map2) would return true
//
// compareMaps(map1,map3) would return false
func compareMaps(map1, map2 map[string]any) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val := range map1 {
		if map2[key] != val {
			return false
		}
	}

	return true
}

// CompareMapSlices compares two slices of maps and checks whether or not they have the same values.
// If one of the maps has the same elements than the other, but also other elements which
// the first one doesn't, it'll return false.
// For example:
//
//	s1 := []map[stirng]any{
//		{
//			fruit:"banana",
//			price: 12.3
//		}
//	}
//
//	s2 := []map[stirng]any{
//		{
//			fruit:"banana",
//			price: 12.3
//		}
//	}
//
//	s3 := []map[stirng]any{
//		{
//			fruit:"apple",
//			price: 12.3
//		}
//	}
//
// CompareMapSlices(s1,s2) [true]
//
// CompareMapSlices(s1,s3) [false]
func CompareMapSlices(slice1, slice2 []map[string]any) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !compareMaps(slice1[i], slice2[i]) {
			return false
		}
	}
	return true
}
