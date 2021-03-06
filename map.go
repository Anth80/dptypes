package main

import (
	"bytes"
	"html/template"
)

type mapParams struct {
	Name     string
	Pkg      string
	DataType string
}

func genMap(name, pkg, dataType string) ([]byte, error) {
	t := template.Must(template.New("map").Parse(`package {{.Pkg}}
	
// Code generated by dptypes. DO NOT EDIT.

import (
	"io/ioutil"
	"math"
	"syscall"
	"unsafe"
)

type {{.Name}} struct {
	fd      int
	buf     []byte
	bufSize int64
	lut     map[string]int64
	size    int
}

func New{{.Name}}() *{{.Name}} {
	x := {{.Name}}{}
	var fd int
	if file, err := ioutil.TempFile("/tmp", "dpmap.*.mmap"); err != nil {
		panic(err)
	} else {
		fd = int(file.Fd())
	}
	return &{{.Name}}{
		fd:   fd,
		buf:  make([]byte, 0),
		lut:  make(map[string]int64, 0),
		size: int(math.Ceil(float64(unsafe.Sizeof(x))/8)*8),
	}
}

func (x *{{.Name}}) Set(key string, value {{.DataType}}) {

}

func (x *{{.Name}}) Get(key string) (*{{.DataType}}, bool) {
	if val, ok := x.lut[key]; !ok {
		return nil, false
	} else {
		buf, err := syscall.Mmap(x.fd, val*8, x.size, syscall.PROT_READ, syscall.MAP_SHARED)
		if err != nil {
			panic(err)
		}
		var x {{.DataType}} = *(*{{.DataType}})(unsafe.Pointer(&buf))
		syscall.Munmap(buf)
		return &x, true
	}
}
	`))

	w := bytes.NewBuffer(nil)
	p := mapParams{Name: name, Pkg: pkg, DataType: dataType}
	if err := t.Execute(w, p); err != nil {
		return []byte{}, err
	}

	return w.Bytes(), nil
}
