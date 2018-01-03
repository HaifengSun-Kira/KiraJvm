package classfile

type SoureceFileAttribute struct {
	cp				ConstantPool
	sourceFileIndex uint16
}

func (self *SoureceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SoureceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}

