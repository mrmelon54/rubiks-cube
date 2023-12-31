package rubiks_cube

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

//go:generate stringer -type Move

var ErrInvalidMove = errors.New("invalid move")

type Move byte

const (
	Up Move = iota
	Down
	Front
	Back
	Right
	Left
	UpPrime
	DownPrime
	FrontPrime
	BackPrime
	RightPrime
	LeftPrime
)

func (m Move) Reverse() Move {
	return (m + 6) % 6
}

func (m Move) Prime() bool {
	return m >= UpPrime
}

func (m Move) Valid() bool {
	return m <= LeftPrime
}

type MoveScanner struct {
	b           *bufio.Reader
	err         error
	currentMove Move
}

func NewMoveScanner(r io.Reader) *MoveScanner {
	return &MoveScanner{b: bufio.NewReader(r)}
}

func (s *MoveScanner) Scan() bool {
	readByte, err := s.b.ReadByte()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false
		}
		s.err = err
		return false
	}
	switch readByte {
	case 'U':
		s.currentMove = Up
	case 'D':
		s.currentMove = Down
	case 'F':
		s.currentMove = Front
	case 'B':
		s.currentMove = Back
	case 'R':
		s.currentMove = Right
	case 'L':
		s.currentMove = Left
	default:
		s.err = ErrInvalidMove
		return false
	}
	readByte, err = s.b.ReadByte()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false
		}
		s.err = err
		return false
	}
	fmt.Println(s.currentMove, rune(readByte))
	if readByte == '\'' {
		s.currentMove = s.currentMove.Reverse()
	} else {
		s.err = s.b.UnreadByte()
		if s.err != nil {
			return false
		}
	}
	fmt.Println(s.currentMove)
	return true
}

func (s *MoveScanner) Current() Move {
	return s.currentMove
}

func (s *MoveScanner) Err() error {
	return s.err
}
