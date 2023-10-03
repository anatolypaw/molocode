package storage

import "log"

type fake struct {
	i int
}

func NewFake(s string) (*fake, error) {
	const op = "storage.fakedb.NewFake"

	log.Printf("%s: %s", op, s)
	return &fake{i: len(s)}, nil
}

func (f *fake) Close() error {
	const op = "storage.fakedb.close"

	log.Printf("%s: %v", op, f.i)

	log.Prefix()
	return nil
}
