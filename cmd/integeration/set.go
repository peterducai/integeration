package integeration

// Set is a well-defined collection of distinct objects.
// In mathematics, a set is a well-defined collection of distinct objects,
// considered as an object in its own right.
// For example, the numbers 2, 4, and 6 are distinct objects when considered separately,
// but when they are considered collectively they form a single set of size three, written {2, 4, 6}.
// One of Cantor’s basic concepts is the notion of the size or cardinality of a
// set M, denoted by |M|.
type Set struct {
	rational    bool
	cardinality uint64
}

// //The set Q of rational numbers is countable.
// //The set R of real numbers is not countable

// // Add inserts an element into the set.
// func Add(e int) {

// }

// // Has reports the existence of the element in the set.
// func Has(e int) bool {

// 	return ok
// }

// // Remove deletes the specified element from the set.
// func Remove(e int) {
// 	delete(s, e)
// }

// // Count reports the number of elements stored in the set.
// func Count() int {
// 	return len(s)
// }

// // IntsEqual reports set equality between the parameters. Sets are equal if
// // and only if they have the same elements.
// func IntsEqual(a, b Int64) bool {
// 	if intsSame(a, b) {
// 		return true
// 	}

// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for e := range a {
// 		if _, ok := b[e]; !ok {
// 			return false
// 		}
// 	}

// 	return true
// }

// // Add inserts an element into the set.
// func Add(e int64) {

// }

// // Has reports the existence of the element in the set.
// func Has(e int64) bool {
// 	_, ok := s[e]
// 	return ok
// }

// // Remove deletes the specified element from the set.
// func Remove(e int64) {
// 	delete(s, e)
// }

// // Count reports the number of elements stored in the set.
// func Count() int {
// 	return len(s)
// }

// // Int64sEqual reports set equality between the parameters. Sets are equal if
// // and only if they have the same elements.
// func Int64sEqual(a, b Int64s) bool {
// 	if int64sSame(a, b) {
// 		return true
// 	}

// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for e := range a {
// 		if _, ok := b[e]; !ok {
// 			return false
// 		}
// 	}

// 	return true
// }

// // Equal reports set equality between the parameters. Sets are equal if
// // and only if they have the same elements.
// func Equal(a, b Nodes) bool {

// 	return true
// }

// // Union takes the union of a and b, and stores it in dst.
// //
// // The union of two sets, a and b, is the set containing all the
// // elements of each, for instance:
// //
// //     {a,b,c} UNION {d,e,f} = {a,b,c,d,e,f}
// //
// // Since sets may not have repetition, unions of two sets that overlap
// // do not contain repeat elements, that is:
// //
// //     {a,b,c} UNION {b,c,d} = {a,b,c,d}
// //
// func  Union(a, b Nodes) Nodes {

// 	return dst
// }

// // Intersect takes the intersection of a and b, and stores it in dst.
// //
// // The intersection of two sets, a and b, is the set containing all
// // the elements shared between the two sets, for instance:
// //
// //     {a,b,c} INTERSECT {b,c,d} = {b,c}
// //
// // The intersection between a set and itself is itself, and thus
// // effectively a copy operation:
// //
// //     {a,b,c} INTERSECT {a,b,c} = {a,b,c}
// //
// // The intersection between two sets that share no elements is the empty
// // set:
// //
// //     {a,b,c} INTERSECT {d,e,f} = {}
// //
// func (dst Nodes) Intersect(a, b Nodes) Nodes {
// 	var swap Nodes

// 	if same(a, b) {
// 		return dst.Copy(a)
// 	}
// 	if same(a, dst) {
// 		swap = b
// 	} else if same(b, dst) {
// 		swap = a
// 	} else {
// 		dst.clear

// 		if len(a) > len(b) {
// 			a, b = b, a
// 		}

// 		for e, n := range a {
// 			if _, ok := b[e]; ok {
// 				dst[e] = n
// 			}
// 		}

// 		return dst
// 	}

// 	for e := range dst {
// 		if _, ok := swap[e]; !ok {
// 			delete(dst, e)
// 		}
// 	}

// 	return dst
// }
