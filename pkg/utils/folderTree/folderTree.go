package folderTree

import (
	"VulTracks/pkg/interfaces"
	"sort"
)

type FolderTreeNode struct {
	Folder   interfaces.CountPerFolderInterface
	Children []FolderTreeNode
}

func BuildTree(folders []interfaces.CountPerFolderInterface, parentID string) []FolderTreeNode {
	result := make([]FolderTreeNode, 0)

	for _, folder := range folders {
		if folder.ParentId.String == parentID {
			node := FolderTreeNode{
				Folder: folder,
			}
			children := BuildTree(folders, folder.Id)
			sort.Slice(children, func(i, j int) bool {
				return children[i].Folder.Name < children[j].Folder.Name
			})
			if len(children) > 0 {
				node.Children = children
			}
			result = append(result, node)
		}
	}

	return result
}
