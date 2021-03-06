package kverrors

import (
	"fmt"
)

type KeyNotFoundError struct {
	Key     interface{}
	Closest interface{}
}

func (err *KeyNotFoundError) Error() string {
	return fmt.Sprintf("key %v not found", err.Key)
}

type InsertionError struct {
	Type     interface{}
	Value    interface{}
	Size     interface{}
	Position int
	Capacity int
}

func (err *InsertionError) Error() string {
	return fmt.Sprintf("could not insert %v with value %v at position %d in slice of size %d/%d", err.Type, err.Value, err.Position, err.Size, err.Capacity)
}

type DeletionError struct {
	Type     interface{}
	Value    interface{}
	Size     interface{}
	Position int
	Capacity int
}

func (err *DeletionError) Error() string {
	return fmt.Sprintf("could not delete %v with value %v at position %d in slice of size %d/%d", err.Type, err.Value, err.Position, err.Size, err.Capacity)
}

type OverflowError struct {
	Type   interface{}
	Max    interface{}
	Actual interface{}
}

func (err *OverflowError) Error() string {
	return fmt.Sprintf("the size of the slice \"%v\" exceeds its supposed bound %v/%v", err.Type, err.Max, err.Actual)
}

type IllegalValueError struct {
	Value interface{}
	Type  interface{}
}

func (err *IllegalValueError) Error() string {
	return fmt.Sprintf("illegal value %v cannot be used as type %v", err.Value, err.Type)
}

type BufferOverflowError struct {
	Max    interface{}
	Cursor interface{}
}

func (err *BufferOverflowError) Error() string {
	return fmt.Sprintf("buffer overflow: max %v, cursor %v", err.Max, err.Cursor)
}

type InvalidSizeError struct {
	Got    interface{}
	Should interface{}
}

func (err *InvalidSizeError) Error() string {
	return fmt.Sprintf("invalid size for data, got %v expected %v", err.Got, err.Should)
}

type IndexOutOfRangeError struct {
	Index  interface{}
	Length interface{}
}

func (err *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("index %v out of range, length %v", err.Index, err.Length)
}

type InvalidNodeSizeError struct {
	NumberOfChildren interface{}
	NumberOfEntries  interface{}
}

func (err *InvalidNodeSizeError) Error() string {
	return fmt.Sprintf("invalid node size: children (%d) and entries (%d) cannot both be superior to 0", err.NumberOfChildren, err.NumberOfEntries)
}

type InvalidNodeError struct{}

func (err *InvalidNodeError) Error() string {
	return fmt.Sprintf("pager returned an invalid node")
}

type UnregisteredError struct{}

func (err *UnregisteredError) Error() string {
	return "the frame id provided doesn't match any registered frame, please register first"
}

type PartialWriteError struct {
	Total   int
	Written int
}

func (err *PartialWriteError) Error() string {
	return fmt.Sprintf("partial write: %d/%d", err.Written, err.Total)
}

type PartialReadError struct {
	Total int
	Read  int
}

func (err *PartialReadError) Error() string {
	return fmt.Sprintf("partial read: %d/%d", err.Read, err.Total)
}

type FrameOverflowError struct {
	Max interface{}
}

func (err *FrameOverflowError) Error() string {
	return fmt.Sprintf("frame overflow: max %v", err.Max)
}

type InvalidFrameIdError struct{}

func (err *InvalidFrameIdError) Error() string {
	return "the frame id provided is invalid"
}

type InvalidMetadataError struct {
	Root interface{}
	Size interface{}
}

func (err *InvalidMetadataError) Error() string {
	return fmt.Sprintf("invalid metadata: root %v, size %v", err.Root, err.Size)
}

type BufferPoolLimitError struct {
	Limit interface{}
}

func (err *BufferPoolLimitError) Error() string {
	return fmt.Sprintf("buffer pool limit reached: %v", err.Limit)
}

type InvalidNodeIOError struct {
	Node   interface{}
	Cursor interface{}
}

func (err *InvalidNodeIOError) Error() string {
	return fmt.Sprintf("invalid node io: node %v, cursor %v", err.Node, err.Cursor)
}

type UnspecifiedFileError struct{}

func (err *UnspecifiedFileError) Error() string {
	return "unspecified file: please specify a file"
}

type OutsideOfRangeError struct {
	From   interface{}
	To     interface{}
	Actual interface{}
}

func (err *OutsideOfRangeError) Error() string {
	return fmt.Sprintf("value %v is outside of range %v-%v", err.Actual, err.From, err.To)
}
