package model

import (
	"fmt"
	"testing"
)

func TestGetCompanyTypeCorrect(t *testing.T) {
	c := Company{ID: 12345,
		Name:    "ABCD.LTD",
		Country: "China",
	}

	companyType := c.GetCompanyType()
	fmt.Println(companyType)

	if companyType != "Limited Liability Company" {
		t.Errorf("Test failed")
	}
}
