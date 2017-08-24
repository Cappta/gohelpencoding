package gohelpencoding

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Cappta/gofixture"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBase64StdDecodeBatch(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a length higher than 2 and lower than 10", func() {
			length := gofixture.AnyIntBetween(2, 10)
			Convey(fmt.Sprintf("Given %d byte slices with any length between 10 and 1000", length), func() {
				bytes := make([][]byte, length)
				for i := 0; i < length; i++ {
					bytes[i] = gofixture.AnyBytes(gofixture.AnyIntBetween(10, 1000))
				}
				Convey("When converting the byte slice to a Base64 string slice", func() {
					base64Slice := make([]string, length)
					for i := 0; i < length; i++ {
						base64Slice[i] = base64.StdEncoding.EncodeToString(bytes[i])
					}
					Convey("When creating a new Base64DecodedBatch from the Base64 string slice", func() {
						batch, err := NewBase64StdDecodeBatch(base64Slice...)
						Convey("Then batch should not be nil", func() {
							So(batch, ShouldNotBeNil)
						})
						Convey("Then error should be nil", func() {
							So(err, ShouldBeNil)
						})
						for i := 0; i < length; i++ {
							nextValue := batch.Next()
							Convey(fmt.Sprintf("Then bytes from index %d should match next value on batch", i), func() {
								So(nextValue, ShouldResemble, bytes[i])
							})
						}
					})
				})
			})
		})
	})
	Convey("Given an invalid base64 string", t, func() {
		invalidString := gofixture.String(256)
		Convey("When creating a new Base64DecodedBatch", func() {
			batch, err := NewBase64StdDecodeBatch(invalidString)
			Convey("Then batch should be nil", func() {
				So(batch, ShouldBeNil)
			})
			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
