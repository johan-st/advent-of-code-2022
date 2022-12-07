package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	// Read input
	input := strings.Split(u.Load("input.txt"), "\r\n")

	// parse into filesystem
	fs := newDir("/")
	pos := []string{}
	for _, line := range input {
		// split line into path and size
		words := strings.Split(line, " ")

		if words[0] == "$" && words[1] == "cd" && words[2] == "/" {
			pos = []string{}
		} else if words[0] == "$" && words[1] == "cd" && words[2] == ".." {
			pos = pos[:len(pos)-1]
		} else if words[0] == "$" && words[1] == "cd" {
			pos = append(pos, words[2])
		} else if words[0] == "$" && words[1] == "ls" {
			continue
		} else if words[0] == "dir" {
			fs.at(pos).addDir(words[1])
		} else {
			// is file
			size, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}
			fs.at(pos).addFile(words[1], size)
		}

	}

	// Print filesystem
	// fmt.Printf("%vTOTAL SIZE: %d\n", fs, fs.size())

	// Find total size of directories less than 100000 in size
	fmt.Printf("SIZE PART 1: %d\n", part1(fs))

	// part 2
	diskSize := 70000000
	diskNeeded := 30000000
	diskUsed := fs.size()
	diskFree := diskSize - diskUsed
	diskToFree := diskNeeded - diskFree

	// fmt.Printf("free on disk: %d\nneeded: %d\nspace to free: %d\n", diskFree, diskNeeded, diskToFree)
	// fmt.Printf("smallest dir larger than %d\n", diskToFree)
	fmt.Printf("SIZE PART 2: %d\n", part2(fs, diskToFree))

}

func part1(fs dir) int {
	size := 0
	for _, d := range fs.dirs {
		if d.size() < 100000 {
			size += d.size()
		}
		size += part1(d)
	}
	return size
}

func part2(d dir, sizeNeeded int) int {
	sizes := []int{}
	for _, d := range d.dirs {
		if d.size() >= sizeNeeded {
			sizes = append(sizes, d.size())
		}
		sizes = append(sizes, part2(d, sizeNeeded))
	}
	if len(sizes) > 0 {

		sort.Ints(sizes)
		return sizes[0]
	} else {
		return 70000000
	}
}

func (d dir) at(pos []string) dir {
	current := d
	for _, p := range pos {
		current = current.dirs[p]
	}

	return current
}

func newDir(name string) dir {
	return dir{name: name, dirs: map[string]dir{}, files: map[string]int{}}
}

// kind is either "dir" or "file"
type dir struct {
	name  string
	dirs  map[string]dir
	files map[string]int
}

func (d dir) addDir(name string) {
	d.dirs[name] = newDir(name)
}

func (d dir) addFile(name string, size int) {
	d.files[name] = size
}

func (d dir) size() int {
	size := 0
	for _, s := range d.files {
		size += int(s)
	}
	for _, d := range d.dirs {
		size += d.size()
	}
	return size
}

func (d dir) String() string {
	return d.toString(0)
}

func (d dir) toString(indent int) string {
	s := fmt.Sprintf("%s- %s (dir)\n", indentString(indent), d.name)
	for name, size := range d.files {
		s += fmt.Sprintf("%s- %s (file, size=%d)\n", indentString(indent+1), name, size)
	}
	for _, d := range d.dirs {
		s += d.toString(indent + 1)
	}
	return s
}

func indentString(indent int) string {
	s := ""
	for i := 0; i < indent; i++ {
		s += "  "
	}
	return s
}
