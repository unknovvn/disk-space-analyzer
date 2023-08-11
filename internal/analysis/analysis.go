package analysis

import (
	"io/fs"
	"path"
	"sort"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

type Record struct {
	Name       string
	Size_bytes int64
}

type record_list []Record

func (r record_list) Len() int           { return len(r) }
func (r record_list) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r record_list) Less(i, j int) bool { return r[i].Size_bytes > r[j].Size_bytes }

func Run(current_dir string, entries []fs.DirEntry) (record_list, error) {
	records := make(record_list, len(entries))
	for idx, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if fileInfo.IsDir() {
			dir_size, err := getFolderSize.Invoke(path.Join(current_dir, fileInfo.Name()))
			if err != nil {
				return nil, err
			}

			records[idx] = Record{Name: fileInfo.Name(), Size_bytes: dir_size}
		} else {
			records[idx] = Record{Name: fileInfo.Name(), Size_bytes: fileInfo.Size()}
		}
	}

	sort.Sort(records)

	return records, nil
}
