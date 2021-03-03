// Copyright (c) 2021 John Dewey <john@dewey.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// Debug starts a new message with debug level.
func Debug() *zerolog.Event {
	return log.Logger.Debug()
}

// Info starts a new message with info level.
func Info() *zerolog.Event {
	return log.Logger.Info()
}

// Warn starts a new message with warn level.
func Warn() *zerolog.Event {
	return log.Logger.Warn()
}

// Error starts a new message with error level.
func Error() *zerolog.Event {
	return log.Logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
func Fatal() *zerolog.Event {
	return log.Logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
func Panic() *zerolog.Event {
	return log.Logger.Panic()
}
