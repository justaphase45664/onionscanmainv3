package utils

import (
	"testing"
)

func TestIsOnion(t *testing.T) {

	test := IsOnion("trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	if !test {
		t.Errorf("%s should have matched!", "trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	}

	test = IsOnion("subdomain.lots.of.trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	if !test {
		t.Errorf("%s should have matched!", "subdomain.lots.of.trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	}

	test = IsOnion("trl4l2alh4biqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	if test {
		t.Errorf("%s should not have matched!", "trl4l2alh4biqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	}

	test = IsOnion("trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	if test {
		t.Errorf("%s should not have matched!", "trl4l2alembiqvixi3yna53c4d3du5ecuevrueoh3vytrkgh6fjsyqid.onion")
	}

	test = IsOnion("www.thisisnotanonionitistoolong.onion")
	if test {
		t.Errorf("%s should not have matched!", "www.thisisnotanonionitistoolong.onion")
	}

}
