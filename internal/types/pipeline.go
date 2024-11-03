package types

type PipelineData struct {
	Value        any
	SourcePath   string
	Directory    string
	BaseName     string
	DeleteOrigin bool
	TargetExt    string
}
