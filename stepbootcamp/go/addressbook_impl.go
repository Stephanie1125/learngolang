package addressbook

// bookImpl is a concrete implementation of AddressBook
type bookImpl []Contact

func init() {
	New = func() AddressBook {
		return &bookImpl{}
	}
}

// Assert that bookImpl does actually implement AddressBook
var _ AddressBook = &bookImpl{}

// TODO(step students): write the remainder of this file.
// このファイルの残りの中身を書いてください！

func (b *bookImpl) Insert (c Contact) bool{
	for i := range *b {
		if ((*b)[i].Name == c.Name) {
			return false
		}
		if ((*b)[i].Email == c.Email){
			return false
		}
	}
	*b = append(*b, c)
	return true
}

func (b *bookImpl) FindName(name string) *Contact{
	for i := range *b {
		if ((*b)[i].Name == name){
			return &(*b)[i]
		}
	}
	return nil
}

// FindEmail looks up a contact by the specified email address.  If
// no such email address is found, it returns nil. The pointer
// returned may be to a copy of the original.

func (b *bookImpl) FindEmail(email string) *Contact{
	for i := range *b {
		if ((*b)[i].Email == email){
			return &(*b)[i]
		}
	}
	return nil
}

// FindNext returns a pointer to the contact that comes
// lexicographically after the given contact. It returns nil if
// there are no more contacts.
func (b *bookImpl) FindNext(c Contact) *Contact{
	for i := range *b{
		if ((*b)[i] == c && i != len(*b)-1){
			return &(*b)[i+1]
		}
	}
	return nil
}

// Delete deletes the specified contact. Returns true iff the given
// contact was actually deleted.
func (b *bookImpl) Delete(c Contact) bool{
	for i := range *b{
		if ((*b)[i] == c){
			*b = append((*b)[:i],(*b)[i+1:]...)
			return true
		}
	}
	return false
}