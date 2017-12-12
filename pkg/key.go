package pkg

/*
#include <linux/input-event-codes.h>
*/
import "C"
import "errors"

var keys = map[uint]string{
	C.KEY_Q: "q",
	C.KEY_W: "w",
	C.KEY_E: "e",
	C.KEY_R: "r",
	C.KEY_T: "t",
	C.KEY_Y: "y",
	C.KEY_U: "u",
	C.KEY_I: "i",
	C.KEY_O: "o",
	C.KEY_P: "p",

	C.KEY_A: "a",
	C.KEY_S: "s",
	C.KEY_D: "d",
	C.KEY_F: "f",
	C.KEY_G: "g",
	C.KEY_H: "h",
	C.KEY_J: "j",
	C.KEY_K: "k",
	C.KEY_L: "l",
	C.KEY_Z: "z",

	C.KEY_X: "x",
	C.KEY_C: "c",
	C.KEY_V: "v",
	C.KEY_B: "b",
	C.KEY_N: "n",
	C.KEY_M: "m",

	C.KEY_1: "1",
	C.KEY_2: "2",
	C.KEY_3: "3",
	C.KEY_4: "4",
	C.KEY_5: "5",
	C.KEY_6: "6",
	C.KEY_7: "7",
	C.KEY_8: "8",
	C.KEY_9: "9",
	C.KEY_0: "0",

	C.KEY_SPACE:     " ",
	C.KEY_DOT:       ".",
	C.KEY_COMMA:     ",",
	C.KEY_SEMICOLON: ";",
}

func getKey(c uint) (ret string, err error) {
	ret = keys[c]
	if ret == "" {
		ret = ""
		err = errors.New("unknow key")
		goto end
	}

end:
	return
}
