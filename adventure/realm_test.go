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

	realm, err := BuildRealm(AREAS_PATH)
	if err != nil {
		t.Error(err)
	}
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
