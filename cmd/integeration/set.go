package integeration

// Set is a well-defined collection of distinct objects.
// In mathematics, a set is a well-defined collection of distinct objects,
// considered as an object in its own right.
// For example, the numbers 2, 4, and 6 are distinct objects when considered separately,
// but when they are considered collectively they form a single set of size three, written {2, 4, 6}.
//
// Sets are conventionally denoted with capital letters.
// Sets A and B are equal if and only if they have precisely the same elements
//
// If B is a set and x is one of the objects of B, this is denoted as x ∈ B, and is read as "x is an element of B",
// as "x belongs to B", or "x is in B"
//
// One of Cantor’s basic concepts is the notion of the size or cardinality of a
// set M, denoted by |M|.
type Set struct {
	rational    bool
	cardinality uint32
}

// Special sets of numbers include

// the empty set, denoted { } or ∅
// A set with exactly one element, x, is a unit set, or singleton, {x}
// P or ℙ, denoting the set of all primes: P = {2, 3, 5, 7, 11, 13, 17, ...}.[49]
// N or {\displaystyle \mathbb {N} }\mathbb {N} , denoting the set of all natural numbers: N = {0, 1, 2, 3, ...} (sometimes defined excluding 0).[48]
// Z or {\displaystyle \mathbb {Z} }\mathbb {Z} , denoting the set of all integers (whether positive, negative or zero): Z = {..., −2, −1, 0, 1, 2, ...}.[48]
// Q or ℚ, denoting the set of all rational numbers (that is, the set of all proper and improper fractions): Q = {a/b | a, b ∈ Z, b ≠ 0}. For example, 1/4 ∈ Q and 11/6 ∈ Q. All integers are in this set since every integer a can be expressed as the fraction a/1 (Z ⊊ Q).[48]
// R or {\displaystyle \mathbb {R} }\mathbb {R} , denoting the set of all real numbers. This set includes all rational numbers, together with all irrational numbers (that is, algebraic numbers that cannot be rewritten as fractions such as √2, as well as transcendental numbers such as π, e).[48]
// C or ℂ, denoting the set of all complex numbers: C = {a + bi | a, b ∈ R}. For example, 1 + 2i ∈ C.[48]
// H or ℍ, denoting the set of all quaternions: H = {a + bi + cj + dk | a, b, c, d ∈ R}. For example, 1 + i + 2j − k ∈ H.[50]

// //The set Q of rational numbers is countable.
// //The set R of real numbers is not countable

// Union - Two sets can be "added" together. The union of A and B,
// denoted by A ∪ B, is the set of all things that are members of either A or B.
func Union() {

}

//Intersection - A new set can also be constructed by determining which members two sets have "in common".
// The intersection of A and B, denoted by A ∩ B, is the set of all things that are members of both A and B.
func Intersection() {

}

//IsSetIntersectionDisjoint - If A ∩ B = ∅, then A and B are said to be disjoint.
func IsSetIntersectionDisjoint() {

}

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
