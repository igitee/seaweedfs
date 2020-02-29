package shell

import (
	"fmt"
	"io"
	"strings"

	"github.com/chrislusf/seaweedfs/weed/filer2"
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
)

func init() {
	Commands = append(Commands, &commandFsTree{})
}

type commandFsTree struct {
}

func (c *commandFsTree) Name() string {
	return "fs.tree"
}

func (c *commandFsTree) Help() string {
	return `recursively list all files under a directory

	fs.tree http://<filer_server>:<port>/dir/
`
}

func (c *commandFsTree) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {

	filerServer, filerPort, path, err := commandEnv.parseUrl(findInputDirectory(args))
	if err != nil {
		return err
	}

	dir, name := filer2.FullPath(path).DirAndName()

	dirCount, fCount, terr := treeTraverseDirectory(writer, commandEnv.getFilerClient(filerServer, filerPort), filer2.FullPath(dir), name, newPrefix(), -1)

	if terr == nil {
		fmt.Fprintf(writer, "%d directories, %d files\n", dirCount, fCount)
	}

	return terr

}

func treeTraverseDirectory(writer io.Writer, filerClient filer2.FilerClient, dir filer2.FullPath, name string, prefix *Prefix, level int) (directoryCount, fileCount int64, err error) {

	prefix.addMarker(level)

	err = filer2.ReadDirAllEntries(filerClient, dir, name, func(entry *filer_pb.Entry, isLast bool) {
		if level < 0 && name != "" {
			if entry.Name != name {
				return
			}
		}

		fmt.Fprintf(writer, "%s%s\n", prefix.getPrefix(level, isLast), entry.Name)

		if entry.IsDirectory {
			directoryCount++
			subDir := dir.Child(entry.Name)
			dirCount, fCount, terr := treeTraverseDirectory(writer, filerClient, subDir, "", prefix, level+1)
			directoryCount += dirCount
			fileCount += fCount
			err = terr
		} else {
			fileCount++
		}

	})
	return
}

type Prefix struct {
	markers map[int]bool
}

func newPrefix() *Prefix {
	return &Prefix{
		markers: make(map[int]bool),
	}
}
func (p *Prefix) addMarker(marker int) {
	p.markers[marker] = true
}
func (p *Prefix) removeMarker(marker int) {
	delete(p.markers, marker)
}
func (p *Prefix) getPrefix(level int, isLastChild bool) string {
	var sb strings.Builder
	if level < 0 {
		return ""
	}
	for i := 0; i < level; i++ {
		if _, ok := p.markers[i]; ok {
			sb.WriteString("│")
		} else {
			sb.WriteString(" ")
		}
		sb.WriteString("   ")
	}
	if isLastChild {
		sb.WriteString("└──")
		p.removeMarker(level)
	} else {
		sb.WriteString("├──")
	}
	return sb.String()
}
