package classfile

/**
ConstantValue_attribute {
	u2 attribute_name_index;
	u4 attribute_length;  // must equal to 2
	u2 constantvalue_index;}
 */
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
