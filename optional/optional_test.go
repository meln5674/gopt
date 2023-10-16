package optional_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/meln5674/gopt/optional"
)

func addOne(x int) int {
	return x + 1
}

func addOnePtr(x *int) int {
	return *x + 1
}

func addOneOptional(x int) optional.Optional[int] {
	return optional.Of(x + 1)
}

func addOneOptionalPtr(x *int) optional.Optional[int] {
	return optional.Of(*x + 1)
}

func alwaysEmpty(x int) optional.Optional[int] {
	return optional.Empty[int]()
}

func alwaysEmptyPtr(x *int) optional.Optional[int] {
	return optional.Empty[int]()
}

var _ = Describe("Optional", func() {
	Describe("Of", func() {
		It("Should return a present value", func() {
			x := optional.Of(1)
			Expect(x.Present()).To(BeTrue())
		})
	})
	Describe("Empty", func() {
		It("Should return a missing value", func() {
			x := optional.Empty[string]()
			Expect(x.Present()).To(BeFalse())
		})
	})
	Describe("OfPointer", func() {
		It("Should return a present value for a non-nil pointer", func() {
			x := 1
			y := optional.OfPointer(&x)
			Expect(y.Present()).To(BeTrue())
		})
		It("Should return a missing value for a nil pointer", func() {
			x := optional.OfPointer[int](nil)
			Expect(x.Present()).To(BeFalse())
		})
	})
	Describe("GetOrPanic", func() {
		It("Should return a present value", func() {
			x := optional.Of(1)
			Expect(x.GetOrPanic()).To(Equal(1))
		})
		It("Should panic for a missing value", func() {
			x := optional.Empty[string]()
			Expect(func() { x.GetOrPanic() }).To(Panic())
		})
	})
	Describe("GetOrDefault", func() {
		It("Should return a present value", func() {
			x := optional.Of(1)
			Expect(x.GetOrDefault(2)).To(Equal(1))
		})
		It("Should return the default for a missing value", func() {
			x := optional.Empty[int]()
			Expect(x.GetOrDefault(2)).To(Equal(2))
		})
	})
	Describe("AsPointer", func() {
		It("Should return the address of a present value", func() {
			x := optional.Of(1)
			ptr := x.AsPointer()
			Expect(*ptr).To(Equal(1))
			*ptr = 2
			Expect(x.GetOrPanic()).To(Equal(2))
		})
		It("Should return nil for a missing value", func() {
			x := optional.Empty[int]()
			Expect(x.AsPointer()).To(BeNil())
		})
	})
	Describe("AsCopyPointer", func() {
		It("Should return the address of a copy of a present value", func() {
			x := optional.Of(1)
			ptr := x.AsCopyPointer()
			Expect(*ptr).To(Equal(1))
			*ptr = 2
			Expect(x.GetOrPanic()).To(Equal(1))
		})
		It("Should return nil for a missing value", func() {
			x := optional.Empty[int]()
			Expect(x.AsCopyPointer()).To(BeNil())
		})
	})

	Describe("Apply", func() {
		It("Should return the result of the function for a present value", func() {
			x := optional.Of(1)
			y := optional.Apply(addOne, x)
			Expect(y.Present()).To(BeTrue())
			Expect(y.GetOrPanic()).To(Equal(2))
		})
		It("Should return a missing value for a missing value", func() {
			x := optional.Empty[int]()
			y := optional.Apply(addOne, x)
			Expect(y.Present()).To(BeFalse())
		})
	})

	Describe("ApplyPtr", func() {
		It("Should return the result of the function for a present value", func() {
			x := optional.Of(1)
			y := optional.ApplyPtr(addOnePtr, &x)
			Expect(y.Present()).To(BeTrue())
			Expect(y.GetOrPanic()).To(Equal(2))
		})
		It("Should return a missing value for a missing value", func() {
			x := optional.Empty[int]()
			y := optional.ApplyPtr(addOnePtr, &x)
			Expect(y.Present()).To(BeFalse())
		})
	})

	Describe("Map", func() {
		It("Should return the result of the function for a present value", func() {
			x := optional.Of(1)
			y := optional.Map(addOneOptional, x)
			Expect(y.Present()).To(BeTrue())
			Expect(y.GetOrPanic()).To(Equal(2))
		})
		It("Should return a missing value if the function returns a missing value", func() {
			x := optional.Of(1)
			y := optional.Map(alwaysEmpty, x)
			Expect(y.Present()).To(BeFalse())
		})
		It("Should return a missing value for a missing value", func() {
			x := optional.Empty[int]()
			y := optional.Map(addOneOptional, x)
			Expect(y.Present()).To(BeFalse())
		})
	})

	Describe("MapPtr", func() {
		It("Should return the result of the function for a present value", func() {
			x := optional.Of(1)
			y := optional.MapPtr(addOneOptionalPtr, &x)
			Expect(y.Present()).To(BeTrue())
			Expect(y.GetOrPanic()).To(Equal(2))
		})
		It("Should return a missing value if the function returns a missing value", func() {
			x := optional.Of(1)
			y := optional.MapPtr(alwaysEmptyPtr, &x)
			Expect(y.Present()).To(BeFalse())
		})
		It("Should return a missing value for a missing value", func() {
			x := optional.Empty[int]()
			y := optional.MapPtr(addOneOptionalPtr, &x)
			Expect(y.Present()).To(BeFalse())
		})
	})
})
