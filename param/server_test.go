package server_test

import (
	"goTest/server"
	"testing"
)

func Test_Options(t *testing.T) {
	var name = "test1"
	var tag = "tag1"
	t.Logf("In name[%s] tag[%s]", name, tag)

	srv := server.NewServer(server.Name(name), server.Tag(tag))
	NewName := srv.GetName()
	NewTag := srv.GetTag()

	t.Logf("Out name[%s] tag[%s]", NewName, NewTag)

	if NewName != name {
		t.Fatalf("src name[%s] != dst name[%s]", name, NewName)
	}
	if NewTag != tag {
		t.Fatalf("src name[%s] != dst name[%s]", tag, NewTag)
	}
}
