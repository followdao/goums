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

func (in *AccessResultT) Builder() *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	b.Finish(AccessResultPack(b, in))
	return b
}

func (in *TerminalListT) Builder() *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	b.Finish(TerminalListPack(b, in))
	return b
}

func (in *TerminalListT) Byte() []byte {
	b := flatbuffers.NewBuilder(0)
	b.Finish(TerminalListPack(b, in))
	return b.FinishedBytes()
}

func (in *TerminalRequestT) Builder() *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	b.Finish(TerminalRequestPack(b, in))
	return b
}

func (in *TerminalRequestT) Byte() []byte {
	b := flatbuffers.NewBuilder(0)
	b.Finish(TerminalRequestPack(b, in))
	return b.FinishedBytes()
}

// ResultBuild   grpc return message to flat
func ResultBuilder(tid, code int64, msg string) *flatbuffers.Builder {
	re := &ResultT{
		Tid:     tid,
		Code:    code,
		Message: msg,
	}

	b := flatbuffers.NewBuilder(0)
	b.Finish(ResultPack(b, re))
	return b
}
