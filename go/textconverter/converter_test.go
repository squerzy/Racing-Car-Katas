package textconverter

import (
	"reflect"
	"testing"
)

func TestConverter(t *testing.T) {
	t.Run("Do something", func(t *testing.T) {
		c := Converter{}
		res, err := c.ConvertToHTML("./some-file.txt")

		expected := []string{
			"a b c d e f",
			"<br />",
			"&gt; gjask",
			"<br />",
			"&amp; lkasj &lt; dlkjasd",
			"<br />",
		}

		isEqual := reflect.DeepEqual(res, expected)

		// for i := range res {
		// 	if res[i] != expected[i] {
		// 		t.Fatalf("Exepected no error, but got one\nExpected: %+v\nResult:   %+v\n", expected[i], res[i])
		// 	}
		// }

		if !isEqual {
			t.Fatalf("Exepected no error, but got one\nExpected: %+v\nResult:   %+v\n", expected, res)
		}

		if err != nil {
			t.Fatal(err)
		}

	})
}
