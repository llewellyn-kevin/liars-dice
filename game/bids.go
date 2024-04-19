package game

// TODO: Test that all these results show up when expected.
func (h *Hand) value() ValidBid {
	switch true {
	case h.isFiveOfAKind():
		return ValidBid_FIVE_OF_A_KIND
	case h.isFourOfAKind():
		return ValidBid_FOUR_OF_A_KIND
	case h.isFullHouse():
		return ValidBid_FULL_HOUSE
	case h.isHighStraight():
		return ValidBid_HIGH_STRAIGHT
	case h.isLowStraight():
		return ValidBid_LOW_STRAIGHT
	case h.isThreeOfAKind():
		return ValidBid_THREE_OF_A_KIND
	case h.isTwoPair():
		return ValidBid_TWO_PAIR
	case h.isOnePair():
		return ValidBid_ONE_PAIR
	default:
		return ValidBid_SINGLE
	}
}

func (h *Hand) isFiveOfAKind() bool {
	return h.highestInstanceCount() == 5
}

func (h *Hand) isFourOfAKind() bool {
	return h.highestInstanceCount() == 4
}

func (h *Hand) isFullHouse() bool {
	counts := h.instanceCount()
	hasThreeOfAKind := false
	hasPair := false
	for _, i := range counts {
		if i == 3 {
			hasThreeOfAKind = true
		} else if i == 2 {
			hasPair = true
		}
	}
	return hasThreeOfAKind && hasPair
}

func (h *Hand) isHighStraight() bool {
	return countsEqual(h.instanceCount(), []uint32{0, 1, 1, 1, 1, 1})
}

func (h *Hand) isLowStraight() bool {
	return countsEqual(h.instanceCount(), []uint32{1, 1, 1, 1, 1, 0})
}

func (h *Hand) isThreeOfAKind() bool {
	return h.highestInstanceCount() == 3
}

func (h *Hand) isTwoPair() bool {
	counts := h.instanceCount()
	pairCount := 0
	for _, v := range counts {
		if v == 2 {
			pairCount += 1
		}
	}
	return pairCount == 2
}

func (h *Hand) isOnePair() bool {
	return h.highestInstanceCount() == 2 && !h.isTwoPair()
}

func (h *Hand) instanceCount() []uint32 {
	var counts []uint32 = []uint32{0, 0, 0, 0, 0, 0}
	for _, i := range h.Values {
		counts[i-1] += 1
	}
	return counts
}

func (h *Hand) highestInstanceCount() uint32 {
	var highest uint32 = 0
	for _, i := range h.instanceCount() {
		if i > highest {
			highest = i
		}
	}
	return highest
}

func countsEqual(a []uint32, b []uint32) bool {
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}
	return true
}
