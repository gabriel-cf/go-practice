package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "a Test"}

		got, _ := Search(dictionary, "test")
		want := "a Test"

		AssertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "a Test"}

		_, err := Search(dictionary, "noWord")
		want := NoWordInDictionaryError

		AssertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "a Test")
		AssertNoError(t, err)
		AssertDefinition(t, dictionary, "test", "a Test")
	})

	t.Run("Add existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "a Test")
		err := dictionary.Add("test", "a Test")
		AssertError(t, err, ExistingWordError)
		AssertDefinition(t, dictionary, "test", "a Test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		dictionary.Add(word, "a Test")
		err := dictionary.Update(word, "another Test")
		AssertNoError(t, err)
		AssertDefinition(t, dictionary, word, "another Test")
	})

	t.Run("Attempt to update non-exsiting word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, "another Test")
		AssertError(t, err, NoWordInDictionaryError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete an existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "a Test"}
		dictionary.Delete(word)
		_, err := Search(dictionary, word)
		if err != NoWordInDictionaryError {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func AssertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func AssertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but got none")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q, want None", got)
	}
}

func AssertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := Search(dictionary, "test")
	if err != nil {
		t.Fatal("Should not have thrown an error")
	}
	AssertStrings(t, got, definition)
}