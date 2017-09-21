package stringOps

import "fmt"


type VersionString interface {
	ToStr() string
	ToInts() [3]uint
	IsValid() bool
}


type SemanticVersion struct {
	major uint
	minor uint
	patch uint
}


func (v SemanticVersion) ToStr() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}


func (v SemanticVersion) ToInts() [3]uint {
	var nums [3]uint
	nums[0] = v.major
	nums[1] = v.minor
	nums[2] = v.patch
	return nums
}


func (v SemanticVersion) IsValid() bool {
	if v.major == 0 && v.minor == 0 && v.patch == 0 {
		return false
	}
	return true
}