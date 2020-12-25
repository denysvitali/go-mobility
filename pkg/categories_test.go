package mobility

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseCategories(t *testing.T) {
	f, err := os.Open("../test/categories.json")
	if err != nil {
		t.Fatal(err)
	}

	var categories CategoriesResp
	jd := json.NewDecoder(f)
	err = jd.Decode(&categories)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range categories.Categories {
		fmt.Printf("%s\n", v.Name)
	}
}
