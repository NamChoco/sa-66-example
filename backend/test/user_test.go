package test

import (
	"fmt"
	"testing"

	"github.com/NamChoco/sa-66-example/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestData(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Username is true`, func(t *testing.T) {
		user := entity.Member{
			Username: "nam", 
			Password: "123456",
			Email: "nam@gmail.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestMember(t *testing.T) {

	g := NewGomegaWithT(t)
 
	t.Run(`Username is required`, func(t *testing.T) {
		user := entity.Member{
			Username: "", //ผิด
			Password: "123456",
			Email: "nam@gmail.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeNil())
		g.Expect(err).NotTo(BeNil())

		// g.Expect(err.Error()).To(Equal("Username is required"))
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Username is required")))
	})
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Email is invalid`, func(t *testing.T) {
		user := entity.Member{
			Username: "nam",
			Password: "123456",
			Email: "namThanawat", //ผิด
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Email is invalid")))
	})
}