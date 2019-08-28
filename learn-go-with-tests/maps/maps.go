package maps

const (
	NoWordInDictionaryError = DictionaryErr("The dictionary does not have that word")
	ExistingWordError = DictionaryErr("The dictionary already has that word")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	var err error = nil
	_, err = Search(d, key)
	switch(err) {
		case NoWordInDictionaryError:
			d[key] = value
			err = nil
		case nil:
			err = ExistingWordError			
	}
	return err
}

func (d Dictionary) Update(key, value string) error {
	_, err := Search(d, key)
	
	if err == NoWordInDictionaryError {
		return NoWordInDictionaryError
	}
	d[key] = value

	return nil

}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func Search(aMap Dictionary, aKey string) (string, error) {
	definition, ok := aMap[aKey]
	if !ok {
		return "", NoWordInDictionaryError
	}
	return definition, nil
}
