package pkg

/*
#cgo CFLAGS: -std=c11

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <linux/input.h>

static inline char* getDeviceName(int fd)
{
	size_t buffer_size = 255;
	char *buffer = (char*)malloc(buffer_size);
	int ret;
	memset(buffer, 0, buffer_size);
	if((ret = ioctl(fd, EVIOCGNAME(buffer_size), buffer)) == -1)
	{
		buffer = (char*) realloc(buffer, 1);
		memset(buffer, 0, 1);
	}

	return buffer;
}

static inline int openDeviceReadOnly(const char* dev)
{
	return open (dev, O_RDONLY, 0);
}

static inline unsigned int startReadingInput(int deviceFD)
{
	int rd, event_len = 64;
	struct input_event event[event_len];

	// vamos a iterar hasta obtener el
	// primer key, del evento del teclado.
	while(1)
	{
		if((rd = read(deviceFD, event, sizeof(struct input_event) * event_len)) == -1)
		{
			return -1;
	  }

	  for (int i = 0; i < rd / sizeof(struct input_event); i++)
	  {
	  	if(event[i].value == EV_KEY)
	  	{
				return (unsigned int) event[i].code;
	  	}
	  }

	}

  return -1;
}
*/
import "C"

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"unsafe"
)

var devicesDir = "/dev/input"

func OpenKeyboardDevice() (dev *Device, err error) {
	filepath.Walk(devicesDir, func(path string, f os.FileInfo, _ error) error {
		if dev != nil {
			goto endWalk
		}

		if !f.IsDir() {
			cstr_path := C.CString(path)
			fd := C.openDeviceReadOnly(cstr_path)
			C.free(unsafe.Pointer(cstr_path))

			if fd != -1 {
				cstr_devicename := C.getDeviceName(fd)

				devicename := C.GoString(cstr_devicename)
				C.free(unsafe.Pointer(cstr_devicename))

				if match, _ := regexp.MatchString(`keyboard`, devicename); match {
					dev = &Device{}
					dev.fd = int(fd)
					dev.Name = devicename
					dev.Path = path
				} else {
					C.close(fd)
				}
			}
		}

	endWalk:
		return nil
	})

	if dev == nil {
		if os.Getuid() != 0 {
			err = errors.New("should be root user to open keyboard device")
		} else {
			err = errors.New("keyboard device not found")
		}
	}

	return
}

type Device struct {
	fd         int
	Name, Path string
}

func (d Device) Close() bool {
	return C.close(C.int(d.fd)) != -1
}

func (d Device) StartReadingInput(fn func(string, error)) {
	for {
		fn(getKey(uint(C.startReadingInput(C.int(d.fd)))))
	}
}
