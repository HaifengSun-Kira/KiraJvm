package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// find and load class file
	// className e.x. java/lang/Object.class
	// return:
	// []byte : byte data of class file
	// Entry : the Entry of the class file
	// error : error message
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	} 
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	   strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		   return newZipEntry(path)
	}
	return newDirEntry(path)
}
