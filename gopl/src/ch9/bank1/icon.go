package bank1

import "image"

/** 
* Created by wanjx in 2018/12/1 22:50
**/

var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	return image.White
}

func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon

}
