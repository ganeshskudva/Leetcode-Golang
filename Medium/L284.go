package Medium

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	nextInt int
	iter    *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	pk := &PeekingIterator{
		nextInt: -1,
		iter:    iter,
	}

	if iter.hasNext() {
		pk.nextInt = iter.next()
	}

	return pk
}

func (this *PeekingIterator) hasNext() bool {
	return this.nextInt != -1
}

func (this *PeekingIterator) next() int {
	res := this.nextInt

	if this.iter.hasNext() {
		this.nextInt = this.iter.next()
	} else {
		this.nextInt = -1
	}

	return res
}

func (this *PeekingIterator) peek() int {
	return this.nextInt
}
