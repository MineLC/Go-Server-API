package util

func Has(mask, field byte) bool {
	return mask&field != 0
}

func Set(mask *byte, field byte, when bool) {
	if when {
		*mask |= field
	}
}
