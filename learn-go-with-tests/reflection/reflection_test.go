package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			} {"Gabriel"},
			[]string{"Gabriel"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				Surname string
			} {"Gabriel", "Chica"},
			[]string{"Gabriel", "Chica"},
		},
		{
			"Struct with string and int field",
			struct {
				Name string
				Age int
			} {"Gabriel", 25},
			[]string{"Gabriel"},
		},
		{
			"Struct with string and int field",
			Person {
				"Gabriel",
				Profile {25, "Dublin"},
			},
			[]string{"Gabriel", "Dublin"},
		},
		{
			"Pointers to things",
			&Person {
				"Gabriel",
				Profile {25, "Dublin"},
			},
			[]string{"Gabriel", "Dublin"},
		},
		{
			"Slices",
			[]Profile {{10, "Dublin"}, {20, "Madrid"}},
			[]string{"Dublin", "Madrid"},
		},
		{
			"Arrays",
			[2]Profile {{10, "Dublin"}, {20, "Madrid"}},
			[]string{"Dublin", "Madrid"},
		},
	}

	for _, test := range(cases) {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Expected %v, got %v", test.ExpectedCalls, got)
			}
		})
	}

	t.Run("Test maps", func(t *testing.T) {
		aMap := map[string]string {"1": "Dublin", "2": "Madrid"}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Dublin")
		assertContains(t, got, "Madrid")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	for _, x := range haystack {
		if x == needle {
			return
		}
	}
	t.Errorf("Expected to find %q but found nothing", needle)
}