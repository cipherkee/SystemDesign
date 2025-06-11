package main

import "time"

type UserMembership struct {
	readers []*Reader
}

func (u *UserMembership) AddMember(r *Reader) error { return nil }

func (u *UserMembership) IsReaderAMember(r *Reader) bool { return true }

func (u *UserMembership) ExtendMembership(r *Reader, t time.Time) error { return nil }
