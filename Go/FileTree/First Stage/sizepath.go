package main

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"
)

type ContainerType struct {
	SizePath  map[int64][]string
	SizeSlice []int64
}

var Container ContainerType = ContainerType{
	SizePath: make(map[int64][]string),
}

func AddPath(path string, info fs.FileInfo, err error) error {
	if info.IsDir() {
		return err
	}
	size := info.Size()
	Container.SizePath[size] = append(Container.SizePath[size], path)
	return err
}
func PrintContainer() {
	fmt.Printf("%v", Container.SizePath)
	fmt.Printf("%v", Container.SizeSlice)
}

func (c *ContainerType) createSliceOfSizes() {
	for index := range c.SizePath {
		c.SizeSlice = append(c.SizeSlice, index)
	}
}

func (c *ContainerType) Order(option int) {
	c.createSliceOfSizes()
	if option == 1 {
		sort.Slice(c.SizeSlice, func(i, j int) bool {
			return c.SizeSlice[i] > c.SizeSlice[j]
		})
	} else if option == 2 {
		sort.Slice(c.SizeSlice, func(i, j int) bool {
			return c.SizeSlice[i] < c.SizeSlice[j]
		})
	}
}
func (c *ContainerType) PrintInfo(fileType string) {

	for _, size := range c.SizeSlice {

		for i, path := range c.SizePath[size] {
			if fileType != "" && !strings.HasSuffix(path, fileType) {
				continue
			}
			if i == 0 {
				fmt.Println()
				fmt.Printf("%d bytes\n", size)
			}
			fmt.Println(path)
		}
	}
}
