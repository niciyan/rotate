package main

import (
	"fmt"
	"sort"
	"strings"
)

type Image struct {
	Id   uint64
	Name string
}

type ByName []Image

func (images ByName) Len() int { return len(images) }
func (images ByName) Less(i, j int) bool {
	return strings.Compare(images[i].Name, images[j].Name) == -1
}
func (images ByName) Swap(i, j int) { images[i], images[j] = images[j], images[i] }

// func for each image
type DoImage func(id uint64)

func main() {
	var gen uint = 3

	images1 := []Image{
		Image{3, "image-003"},
		Image{2, "image-002"},
		Image{4, "image-004"},
	}

	images2 := []Image{
		Image{4, "image-004"},
		Image{3, "image-003"},
	}

	images3 := []Image{
		Image{5, "image-005"},
		Image{2, "image-002"},
		Image{4, "image-004"},
		Image{3, "image-003"},
	}

	fmt.Println(rotate(gen, images1))
	fmt.Println(rotate(gen, images2))
	fmt.Println(rotate(gen, images3))
}

var newImage = Image{
	Id:   100,
	Name: "image-100",
}

// You add 1 element and sort the list
// When list is larger than generation number, remove young element
func rotate(gen uint, images []Image) []Image {
	images = append(images, newImage)

	sort.Sort(ByName(images))

	// exit when images length < gen
	// in other words, no need to remove some element..
	if len(images) <= int(gen) {
		return images
	}

	bound := len(images) - int(gen)
	rest := images[bound:]
	return rest
}
