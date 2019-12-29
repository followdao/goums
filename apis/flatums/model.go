package flatums

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

func (t *TerminalProfileT) Byte() []byte {
	b := flatbuffers.NewBuilder(0)
	b.Finish(TerminalProfilePack(b, t))
	return b.FinishedBytes()
}

// ByteToStbList  to list
func ByteTerminalProfileT(b []byte) *TerminalProfileT {
	return GetRootAsTerminalProfile(b, 0).UnPack()
}
