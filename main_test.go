package main

import (
	"fmt"
	"testing"
)

type ImagesTestSet struct {
	gen    uint
	images []Image
}

func Test3Images(t *testing.T) {
	images1 := []Image{
		{3, "image-003"},
		{2, "image-002"},
		{4, "image-004"},
	}

	images20 := []Image{
		{4, "image-004"},
		{5, "image-005"},
		{6, "image-006"},
		{7, "image-007"},
		{8, "image-010"},
		{9, "image-011"},
		{10, "image-012"},
		{11, "image-013"},
		{12, "image-014"},
		{13, "image-015"},
		{14, "image-016"},
		{15, "image-017"},
		{16, "image-020"},
		{17, "image-021"},
		{18, "image-022"},
		{19, "image-023"},
		{20, "image-024"},
		{21, "image-025"},
		{22, "image-026"},
		{23, "image-027"},
	}

	var goodSets = []ImagesTestSet{
		{3, images1},
		{2, images1},
		{4, images1},
		{1, images20},
		{10, images20},
	}

	var badSets = []ImagesTestSet{
		{0, images1},
	}

	for _, set := range goodSets {
		checkImages(set, t)
	}
}

func checkImages(set ImagesTestSet, t *testing.T) {
	out := rotate(set.gen, set.images)
	fmt.Println(out)

	inLen := int(set.gen)
	outLen := len(out)
	if outLen > inLen {
		t.Errorf("output length: %v, must be same %v\n", outLen, inLen)
	}

	if out[len(out)-1].Id != 100 {
		t.Errorf("last Image id must be %v\n", 100)
	}
}
