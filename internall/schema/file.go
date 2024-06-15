package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type File struct {
	file      *os.File
	Data      []byte
	Dir       string
	Version   string
	FileName  string
	Ext       string
	Print     bool
	CreatedAt time.Time
}

type Version struct {
	Format    string
	Direction string
	Time      time.Time
}

func (f *File) TimeVersion(format string) error {
	switch format {
	case "unix":
		f.Version = strconv.FormatInt(f.CreatedAt.Unix(), 10)
	case "unixNano":
		f.Version = strconv.FormatInt(f.CreatedAt.UnixNano(), 10)
	default:
		f.Version = f.CreatedAt.Format(format)
	}

	return nil
}

func (f *File) Write(data any) error {
	if _, err := f.file.Write([]byte(fmt.Sprintf("%v \n", data))); err != nil {
		return err
	}

	return nil
}

func (f *File) Unmarshal(data any) error {
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(f.Data, &data); err != nil {
		return err
	}

	return nil
}

func (f *File) WithCreate(w func(w io.Writer) error) (err error) {
	f.file, err = os.OpenFile(filepath.Join(f.Dir, f.FileName), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}

	if err := w(f.file); err != nil {
		if cerr := f.file.Close(); cerr != nil {
			// optionally, return something like
			// "error encountered while handling other error: ..."
			return cerr
		}
		return err
	}

	return f.file.Close()
}

func (f *File) CreateFile() (err error) {
	// create exclusive (fails if file already exists)
	// os.Create() specifies 0666 as the FileMode, so we're doing the same
	f.file, err = os.OpenFile(filepath.Join(f.Dir, f.FileName), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return nil
}

func LoadFile(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &File{
		Data:      data,
		Dir:       filepath.Dir(path),
		FileName:  filepath.Base(path),
		Ext:       filepath.Ext(path),
		CreatedAt: time.Now(),
	}, nil
}

func NewFile(dir string, table *Table, version *Version, ext string) (*File, error) {
	direction := ""

	if version != nil {
		direction = version.Direction
	}

	f := &File{
		Data:      []byte{},
		Dir:       filepath.Clean(dir),
		Version:   "",
		FileName:  strings.ReplaceAll(table.Name, "_", "-"),
		Ext:       ext,
		CreatedAt: table.CreatedAt,
	}

	if version != nil {
		if err := f.TimeVersion(version.Format); err != nil {
			return nil, err
		}

		versionGlob := filepath.Join(dir, f.Version+"_*."+direction+f.Ext)
		matches, err := filepath.Glob(versionGlob)
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			return nil, fmt.Errorf("duplicate migration version: %s", f.Version)
		}
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	var filename strings.Builder

	if version != nil {
		filename.WriteString(fmt.Sprintf("%s_%s.%s.%s", f.Version, table.Name, direction, f.Ext))
	} else {
		filename.WriteString(fmt.Sprintf("%s.%s", f.FileName, f.Ext))
	}
	f.FileName = filename.String()

	defer f.file.Close()

	return f, nil
}
