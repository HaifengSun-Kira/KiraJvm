package classfile

/**
SourceFile_attribute {
	u2 attribute_name_index;
	u4 attribute_length;   must equal to 2
	u2 sourcefile_index;
}
 */
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

