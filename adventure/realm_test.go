package gomud

import "testing"

func TestBuildRealm(t *testing.T) {
	expectedAreaCount := 44
	r, err := BuildRealm(AREAS_PATH)
	if err != nil {
		t.Error(err)
	}
	areaCount := 0
	cur := r.Areas.Head
	for cur != nil {
		areaCount = areaCount + 1
		cur = cur.Next
	}
	if areaCount != expectedAreaCount {
		t.Errorf("BuildRealm area count mismatch.")
	}
}

func TestParseAreaData(t *testing.T) {
	expectedRoomsBytes := 846882
	expectedMobilesBytes := 191145
	expectedObjectsBytes := 137701
	realm := TestRealm
	if len(*realm.Rooms.Data) != expectedRoomsBytes {
		t.Errorf("ParseAreaData rooms bytes mismatch.")
	}
	if len(*realm.Mobiles.Data) != expectedMobilesBytes {
		t.Errorf("ParseAreaData mobiles bytes mismatch.")
	}
	if len(*realm.Objects.Data) != expectedObjectsBytes {
		t.Errorf("ParseAreaData objects bytes mismatch.")
	}
}

func TestParseObjects(t *testing.T) {
	realm := TestRealm
	TestRealm = realm
	realm.ParseObjects([]string{})
	objs := realm.Objects
	expectedId := `#1`
	expectedShort := "a mug of coffee"
	expectedLong := "A mug of coffee lies here."
	cur := realm.Objects.Head
	if cur == nil {
		t.Skip("realm.Objects.Head is nil.")
	}
	if cur.Name != expectedShort {
		t.Errorf("ParseObjects name mismatch.")
	}

	obj := objs.FindObjectById(expectedId)
	if obj == nil {
		t.Errorf("failed to find ID: %s\n", expectedId)
	}
	objId, objShort, objLong := obj.GetObjectData()
	if objId != expectedId {
		t.Errorf("object ID mismatch: %s\n", objId)
	}
	if objShort != expectedShort {
		t.Errorf("object Short mismatch: %s\n", objShort)
	}
	if objLong != expectedLong {
		t.Errorf("object Long mismatch: %s\n", objLong)
	}
	if obj.Name != expectedShort {
		t.Errorf("object Name mismatch: %s\n", obj.Name)
	}
}
