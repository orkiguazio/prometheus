/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logger

import (
    "context"
)

// Level defines a log level
type Level uint8

const (
    LevelDebug Level = iota
    LevelInfo
    LevelWarn
    LevelError
)

// Logger allows outputting logs to various logger sinks
type Logger interface {

    // emit a log entry of a given verbosity. the first argument may be an object, a string
    // or a format string. in case of the latter, the following varargs are passed
    // to a formatter (e.g. fmt.Sprintf)

    // Error emits an unstructured error log
    Error(format interface{}, vars ...interface{})

    // Warn emits an unstructured warning log
    Warn(format interface{}, vars ...interface{})

    // Info emits an unstructured informational log
    Info(format interface{}, vars ...interface{})

    // Debug emits an unstructured debug log
    Debug(format interface{}, vars ...interface{})

    // ErrorCtx emits an unstructured error log with context
    ErrorCtx(ctx context.Context, format interface{}, vars ...interface{})

    // WarnCtx emits an unstructured warning log with context
    WarnCtx(ctx context.Context, format interface{}, vars ...interface{})

    // InfoCtx emits an unstructured informational log with context
    InfoCtx(ctx context.Context, format interface{}, vars ...interface{})

    // DebugCtx emits an unstructured debug log with context
    DebugCtx(ctx context.Context, format interface{}, vars ...interface{})

    // emit a structured log entry. example:
    //
    // l.InfoWith("The message",
    //  "first-key", "first-value",
    //  "second-key", 2)
    //

    // ErrorWith emits a structured error log
    ErrorWith(format interface{}, vars ...interface{})

    // WarnWith emits a structured warning log
    WarnWith(format interface{}, vars ...interface{})

    // InfoWith emits a structured info log
    InfoWith(format interface{}, vars ...interface{})

    // DebugWith emits a structured debug log
    DebugWith(format interface{}, vars ...interface{})

    // ErrorWithCtx emits a structured error log with context
    ErrorWithCtx(ctx context.Context, format interface{}, vars ...interface{})

    // WarnWithCtx emits a structured warning log with context
    WarnWithCtx(ctx context.Context, format interface{}, vars ...interface{})

    // InfoWithCtx emits a structured info log with context
    InfoWithCtx(ctx context.Context, format interface{}, vars ...interface{})

    // DebugWithCtx emits a structured debug log with context
    DebugWithCtx(ctx context.Context, format interface{}, vars ...interface{})

    // Flush flushes buffered logs, if applicable
    Flush()

    // GetChild returns a child logger, if underlying logger supports hierarchal logging
    GetChild(name string) Logger
}
